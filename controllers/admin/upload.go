package admin

import (
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) Upload() {
	file, header, e := this.GetFile("file")

	if e != nil {
		this.json(500, e.Error(), "")
	}
	defer file.Close()
	rand.Seed(time.Now().UnixNano())
	max := 999999
	min := 100000
	randnum := rand.Intn(max-min) + min
	s := strconv.Itoa(randnum)
	split := strings.Split(header.Filename, ".")
	ext := split[len(split)-1]
	time := time.Now().Unix()
	s1 := strconv.Itoa(int(time))
	fileName := s1 + s + "." + ext
	path := "static/upload/" + fileName
	e = this.SaveToFile("file", path)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	url := "http://" + beego.AppConfig.String("apphost") + ":" + beego.AppConfig.String("httpport") + "/" + path
	this.json(0, "ok", map[string]interface{}{
		"url": url,
	})
}
