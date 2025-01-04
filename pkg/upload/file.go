package upload

import (
	"github.com/testercc/blog-service/global"
	"github.com/testercc/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

//  FileType 为 int 的类型别名，利用 FileType 作为类别标识的基础类型
type FileType int

// Go 语言中 iota 相当于是一个 const 的常量计数器，你也可以理解为枚举值，第一个声明的 iota 的值为 0，在新的一行被使用时，它的值都会自动递增。
const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// 获取文件名称，先是通过获取文件后缀并筛出原始文件名进行 MD5 加密，最后返回经过加密处理后的文件名。
func GetFileName(name string) string {

	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 获取文件后缀，主要是通过调用 path.Ext 方法进行循环查找”.“符号，最后通过切片索引返回对应的文件后缀名称。
func GetFileExt(name string) string {

	return path.Ext(name)
}

// 获取文件保存地址，这里直接返回配置中的文件保存目录即可，也便于后续的调整。
func GetSavePath() string {

	return global.AppSetting.UploadSavePath
}

// 检查文件的相关方法
// 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查文件后缀是否包含在约定的后缀配置项中，需要的是所上传的文件的后缀有可能是大写、小写、大小写等
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

// 检查文件大小是否超出最大大小限制。
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建在上传文件时所使用的保存目录
func CreateSavePath(dst string, perm os.FileMode) error {

	// 该方法将会以传入的 os.FileMode 权限位去递归创建所需的所有目录结构，若涉及的目录均已存在，则不会进行任何操作，直接返回 nil。
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// 保存所上传的文件，该方法主要是通过调用 os.Create 方法创建目标地址的文件，再通过 file.Open 方法打开源地址的文件，
// 结合 io.Copy 方法实现两者之间的文件内容拷贝
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
