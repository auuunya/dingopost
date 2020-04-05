**This project is under heavy development and is not ready for use. However, we'd love to have your contribution.**

Dingopost is a full-featured blog engine written in Go.

Initial framework source: [dingoblog](https://github.com/dingoblog/dingo) [hux-theme](https://github.com/Kaijun/hexo-theme-huxblog)

![Dingopost backend](https://github.com/mu1er/pictures/blob/master/%E9%80%89%E5%8C%BA_001.png?raw=true)
![Dingopost fronted](https://github.com/mu1er/pictures/blob/master/%E9%80%89%E5%8C%BA_002.png?raw=true)

## Main Features

- **Blog Comments**: Dingo has a built-in comment system.
- **Markdown Editor**: You can write your post in markdown format, with a beautiful markdown editor.
- **Powerful Admin Panel**: Dingo has a powerful dashboard, in which you can view various information about your blog.

## Installation

```
$ go get github.com/mu1er/dingopost
```

## Run the Server

```
$ cd $GOPATH/src/github.com/mu1er/dingopost
$ go run main.go --port 8000
```

## Contributing

To contribute, please take a look at our [roadmap](https://github.com/mu1er/dingopost) to find the issue that you would like to work on.

To read the source code, please start from the [URL endpoints](https://github.com/mu1er/dingopost/app/app.go#L71)

## Admin Panel

Plase visit [http://localhost:8000/signup/](http://localhost:8000/signup/) to register a new user and [http://localhost:8000/login/](http://localhost:8000/login/) to log into the admin panel.

## LICENSE

[MIT LICENSE](/LICENSE)
