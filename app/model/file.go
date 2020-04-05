package model

import (
	"strings"
)

// A File is a file, along with a URL and last modified time.
type File struct {
	Key   string `xml:"Key"`   // Object的Key
	Size  string `xml:"Size"`  // Object的长度字节数
	Utime string `xml:"Utime"` // Object上传时间
	Uname string `xml:"Uname"` // 文件所属人
}

var (
	Accesskey string
	Secretkey string
	Bucket    string
	Endpoint  string
)

var OssFilePath = "http://ygjs-static-hz.oss-cn-beijing.aliyuncs.com/dingo"

// CheckSafe checks if the directory is a child directory of base, to make sure
// that GetFileList won't read any folder other than the upload folder.

// GetFileList returns a slice of all Files in the given directory.

// RemoveFile removes a file with the given path.

// CreateFilePath creates a filepath from the given directory and name,
// returning the name of the newly created filepath.
func CreateFilePath(name string) string {
	return strings.Join([]string{OssFilePath, name}, "/")
}
