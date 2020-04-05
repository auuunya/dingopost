package handler

import (
	"dingo/app/model"
	"dingo/app/utils"
	"io"
	"io/ioutil"
	"log"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dinever/golf"
)

func OssSetting() *oss.Bucket {
	var ss *model.Oss
	ss = model.GetOssSetting()
	client, err := oss.New(ss.Endpoint, ss.Accesskey, ss.Secretkey)
	if err != nil {
		log.Fatal("OSS设置错误", err)
	}
	bucket, err := client.Bucket(ss.Bucket)
	if err != nil {
		log.Fatal("空间打开失败", err)
	}
	return bucket
}

func FileViewHandler(ctx *golf.Context) {
	user, _ := ctx.Session.Get("user")
	ctx.Request.ParseForm()
	bucket := OssSetting()
	lsRes, err := bucket.ListObjects(oss.Prefix("dingo/"), oss.Delimiter("/"))
	if err != nil {
		log.Fatal("打开存储错误...")
	}
	var files []model.File
	for _, common := range lsRes.Objects {
		props, err := bucket.GetObjectDetailedMeta(common.Key)
		if err != nil {
			// HandleError(err)
		}
		name := strings.Split(common.Key, "/")
		files = append(files, model.File{name[1], utils.FileSize(common.Size), props.Get("X-Oss-Meta-Utime"), props.Get("X-Oss-Meta-Uname")})
	}
	Files := files[1:]
	ctx.Loader("admin").Render("files.html", map[string]interface{}{
		"Title": "文件",
		"Files": Files,
		"User":  user,
	})
}

//从aliyun删除文件
func delete_aliyun(filename string) {
	bucket := OssSetting()
	err := bucket.DeleteObject(path.Join("dingo/", filename))
	if err != nil {
		log.Fatal("Delete File Error :", err)
	}
}
func FileRemoveHandler(ctx *golf.Context) {
	name := ctx.Request.FormValue("name")
	user, _ := ctx.Session.Get("user")
	u := user.(*model.User)
	if u.Role == false {
		ctx.JSON(map[string]interface{}{
			"status": "error",
			"msg":    "Insufficient Privileges",
		})
	} else {
		delete_aliyun(name)
	}
}

//上传aliyun
func upload_aliyun(filename string, body io.Reader, args []oss.Option) error {
	bucket := OssSetting()
	err := bucket.PutObject(path.Join("dingo/", filename), body, args...)
	if err != nil {
		return err
	}
	return nil
}

func FileUploadHandler(ctx *golf.Context) {
	req := ctx.Request
	req.ParseMultipartForm(32 << 20)
	userObj, _ := ctx.Session.Get("user")
	u := userObj.(*model.User)
	f, h, e := req.FormFile("file")
	if e != nil {
		ctx.JSON(map[string]interface{}{
			"status": "error",
			"msg":    e.Error(),
		})
		return
	}
	NowTime := utils.Now()
	time := utils.DateFormat(NowTime, "%Y-%m-%d %H:%M")
	options := []oss.Option{
		oss.ContentDisposition(h.Filename),
		oss.Meta("uName", u.Name),
		oss.Meta("uTime", time),
	}

	err := upload_aliyun(h.Filename, f, options)
	if err != nil {
		ctx.JSON(map[string]interface{}{
			"status": "error",
			"msg":    "OSS Not Setting.",
		})
		return
	}
	data, _ := ioutil.ReadAll(f)
	maxSize, _ := ctx.App.Config.GetInt("app.upload_size", 1024*1024*10)
	defer func() {
		f.Close()
		data = nil
		h = nil
	}()
	if len(data) >= maxSize {
		ctx.JSON(map[string]interface{}{
			"status": "error",
			"msg":    "File size should be smaller than 10MB.",
		})
		delete_aliyun(h.Filename)
		return
	}
	fileExt, _ := ctx.App.Config.GetString("app.upload_files", ".rar,.jpg,.png,.gif,.jpeg,.zip,.txt,.doc,.docx,.xls,.xlsx,.ppt,.pptx")
	if !strings.Contains(fileExt, path.Ext(h.Filename)) {
		ctx.JSON(map[string]interface{}{
			"status": "error",
			"msg":    "Only supports documents, images and zip files.",
		})
		delete_aliyun(h.Filename)
		return
	}
	Url := model.CreateFilePath(h.Filename)
	ctx.JSON(map[string]interface{}{
		"status": "success",
		"file": map[string]interface{}{
			"url":   Url,
			"name":  h.Filename,
			"time":  time,
			"uName": u.Name,
		},
	})
}
