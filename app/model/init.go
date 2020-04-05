package model

import (
	"database/sql"
	"dingo/app/utils"
)

var db *sql.DB

const samplePostContent = `
Welcome to Dingo! This is your first post. You can find it in the [admin panel](/admin/).

此文本框支持Markdown语法:

# 标题

## 副标题

### 其他标题

文字样式 _italic_,
**bold**, ` + "`" + `monospace` + "`" + `.

图片:

![Dingo Logo](https://cloud.githubusercontent.com/assets/1311594/14765969/bc14bafc-09c7-11e6-92f8-d69774cca249.png)

无序列表:

  * apples
  * oranges
  * pears

有序列表:

  1. apples
  2. oranges
  3. pears


引用:

> Sportsman delighted improving dashwoods gay instantly happiness six. Ham now amounted absolute not mistaken way pleasant whatever. At an these still no dried folly stood thing. Rapid it on hours hills it seven years. If polite he active county in spirit an. Mrs ham intention promotion engrossed assurance defective. Confined so graceful building opinions whatever trifling in. Insisted out differed ham man endeavor expenses. At on he total their he songs. Related compact effects is on settled do.

代码块:

` + "```" + `go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
` + "```" + `

超链接:

An [example link](http://example.com).

表格:

|        | Cost to x | Cost to y | Cost to z |
|--------|-----------|-----------|-----------|
| From x | 0         | 3         | 4         |
| From y | 3         | 0         | 6         |
| From z | 4         | 6         | 0         |
`

// A Row contains data that can be Scanned into a variable.
type Row interface {
	Scan(dest ...interface{}) error
}

func Initialize() error {
	if err := initConnection(); err != nil {
		return err
	}
	if err := createTableIfNotExist(); err != nil {
		return err
	}
	//	if err := createWelcomeData(); err != nil {
	//		return err
	//	}
	//DbExists()
	if !DbExists() {
		if err := createWelcomeData(); err != nil {
			return err
		}
	}
	return nil
}

// Initialize sets up the DB by creaing a new connection, creating tables if
// they don't exist yet, and creates the welcome data.
//func Initialize(dbPath string, dbExists bool) error {
//	if err := initConnection(dbPath); err != nil {
//		return err
//	}

//	if err := createTableIfNotExist(); err != nil {
//		return err
//	}

//	if !dbExists {
//		if err := createWelcomeData(); err != nil {
//			return err
//		}
//	}

//	return nil
//}

func initConnection() error {
	var err error
	db, err = Conn()
	if err != nil {
		return err
	}
	return nil
}

func createTableIfNotExist() error {
	for i := 0; i < len(TableList); i++ {
		if _, err := db.Exec(TableList[i]); err != nil {
			return err
		}
	}
	checkBlogSettings()
	return nil
}

func checkBlogSettings() {
	SetSettingIfNotExists("theme", "hux-theme", "blog")
	SetSettingIfNotExists("title", "My Blog", "blog")
	SetSettingIfNotExists("description", "Awesome blog created by Dingo.", "blog")
}

func createWelcomeData() error {
	var err error
	p := NewPost()
	p.Title = "Welcome to Dingo!"
	p.Slug = "welcome-to-dingo"
	p.Markdown = samplePostContent
	p.Html = utils.Markdown2Html(p.Markdown)
	p.AllowComment = true
	p.Category = ""
	p.CreatedBy = 0
	p.UpdatedBy = 0
	p.IsPublished = true
	p.IsPage = false
	if err != nil {
		return err
	}

	c := NewComment()
	c.Author = "Shawn Ding"
	c.Email = "luohao.brian@gmail.com"
	c.Website = "https://github.com/luohao-brian/dingo"
	c.Content = "Welcome to Dingo! This is your first comment."
	c.Avatar = utils.Gravatar(c.Email, "50")
	c.PostId = p.Id
	c.Parent = int64(0)
	c.Ip = "127.0.0.1"
	c.UserAgent = "Mozilla"
	c.UserId = 0
	c.Approved = true
	c.Save()
	return nil
}
