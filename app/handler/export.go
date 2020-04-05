package handler

import (
	"bufio"
	"dingo/app/model"
	"dingo/app/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dinever/golf"
)

type FuncMap map[string]interface{}

// ExportMarkdownAction exports articles as markdown zip file.
func ExportGetHandler(ctx *golf.Context) {
	userObj, _ := ctx.Session.Get("user")
	u := userObj.(*model.User)
	p := ctx.Request.FormValue("page")
	var page int
	if p == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(p)
	}
	posts := new(model.Posts)
	var err error
	if u.Role == true {
		_, err = posts.GetPostList(int64(page), 10, false, false, "created_at DESC")
	} else {
		_, err = posts.GetPostAuthor(u.Name, int64(page), 10, false, false, "created_at DESC")
	}
	if err != nil {
		panic(err)
	}

	for _, post := range *posts {
		err = WriteIndexHTML(post.Slug, post)
		if err != nil {
			ctx.Redirect("/admin/posts/")
		}
	}
	ctx.Redirect("/admin/posts/")
}

func getTemplate() (*template.Template, error) {
	funcMap := template.FuncMap{"DateFormat": utils.DateFormat, "Md2html": utils.Markdown2HtmlTemplate}
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	templatePath := filepath.Join(path, "staticPost", "template.html")
	t := template.New("template.html").Funcs(funcMap)
	t, err = t.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("error reading template  %v", err)
	}
	return t, nil
}

func WriteIndexHTML(name string, post *model.Post) error {
	t, err := getTemplate()
	path, _ := os.Getwd()
	err = os.MkdirAll(filepath.Join(path, "staticPost", "PostsView", name), 0777)
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(path, "staticPost", "PostsView", name, "index.html")
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filePath, err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	td := map[string]interface{}{
		"Title":   post.Title,
		"Post":    post,
		"Content": post,
	}
	if err := t.Execute(w, td); err != nil {
		return fmt.Errorf("error executing template %s: %v", filePath, err)
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("error writing file %s: %v", filePath, err)
	}
	return nil
}
