package admin

import (
	"admin/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {

	user := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	valid := validation.Validation{}
	valid.Required(user.Account, "account").Message("account is required")
	valid.Required(user.Password, "password").Message("password is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}

	au, e := services.GetAdminUserByAccountAndPassword(user.Account, user.Password)
	if e != nil {
		this.json(500, "login fail,check your account or password", "")
	}

	j := services.JWT{}
	c := services.Claims{
		ID:     au.ID,
		RoleId: au.RoleId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //过期时间
			Issuer:    "amdin",                     // 签名的发行者
		},
	}
	s, err := j.CreateToken(c)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	this.json(0, "ok", map[string]interface{}{"token": s})
}

func (this *LoginController) Info() {
	token := this.Ctx.Request.Header.Get("X-token")
	if token == "" {
		this.json(500, "the token required", "")
	}

	j := services.JWT{}
	claims, e := j.ParseToken(token)
	if e != nil {
		this.json(500, e.Error(), "")
	}

	ams, e := services.GetAdminMenuByRoleId(claims.RoleId, false)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	au, e := services.GetAdminUserById(int(claims.ID))
	if e != nil {
		this.json(500, e.Error(), "")
	}
	data := map[string]interface{}{
		"user": au,
		"menu": ams,
	}
	this.json(0, "ok", data)
}

func (this *LoginController) LogOut() {
	token := this.Ctx.Request.Header.Get("X-token")
	j := services.JWT{}
	_, e := j.RefreshToken(token, true)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", "")

}

func (this *LoginController) json(code int, msg string, data interface{}) {

	returnJson := returnDataStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(returnJson)
	if err != nil {
		this.Data["json"] = err.Error()
		this.ServeJSON()
	}
	_, _ = this.Ctx.ResponseWriter.Write(ret)
	this.StopRun()
}
