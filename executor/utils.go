package executor

import (
	"DealImages/recored"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
)
import (
	_ "image/jpeg"
	_ "image/png"
)

// 方法 判断图片是竖着还是横着
func IsVertical(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return false, err
	}
	//fmt.Println(img.Bounds().Dx(), img.Bounds().Dy())
	return img.Bounds().Dx() < img.Bounds().Dy(), nil
}

// 拷贝文件 将文件从源路径拷贝到目标路径之下
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	// 获取文件名
	dst = dst + "/" + filepath.Base(src)
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	fmt.Println(recored.GetProgress(), " 拷贝文件: ", src, " 到: ", dst)
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}
