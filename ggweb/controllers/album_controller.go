package controllers


import (
	"fmt"
	"time"
	"os"
	"io"
	"path/filepath"
	"goWeb/ggweb/models"
	"github.com/opentracing/opentracing-go/log"
)


type AlbumController struct{
	BaseController
} 

type UploadController struct{
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		log.Error(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}

func (this *UploadController) Post() {
	fmt.Println("this is UploadController, fileupload.....")
	fileData, fileHeader, err := this.GetFile("upload")
	if err != nil {
		this.responseErr(err)
		return
	}
	fmt.Println("name:",fileHeader.Filename, fileHeader.Size)
	fmt.Println("fileData:",fileData)
	now := time.Now()
	fmt.Println("ext:",filepath.Ext(fileHeader.Filename))
	fileType := "other"
	// 判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fileExt := filepath.Ext(fileHeader.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" ||fileExt == ".jepg" {
		fileType = "img"
	}
	// 文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		this.responseErr(err)
		return
	}
	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmr.Sprintf("%d-%s",timeStamp,fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}
	// 将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		this.responseErr(err)
		return
	}
	if fileType == "img" {
		album := models.Album{0,filePathStr,fileName,0,timeStamp}
		models.InsertAlbum(album)
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message":"上传成功"}
	this.ServeJSON()
}

func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message":err}
	this.ServeJSON()
}