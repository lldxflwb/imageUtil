package infos

import (
	"DealImages/executor"
	"DealImages/recored"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//var FaterPath = "/Volumes/192.168.3.6/1069boys/m/G"
var FaterPath = "/Users/chenjiale/Desktop/a"
var VPath = "/Users/chenjiale/Desktop/images/v"
var HPath = "/Users/chenjiale/Desktop/images/h"
var EXTS = []string{".png", ".jpg", ".JPG"}
var Files []string = []string{}
var Types []string = []string{}

func inArray(str string, array []string) bool {
	for _, v := range array {
		if v == str {
			return true
		}
	}
	return false
}

func addAndMoveFile(path string) error {
	// 如果该文件类型不在类型列表中，则添加
	if !inArray(filepath.Ext(path), Types) {
		Types = append(Types, filepath.Ext(path))
		fmt.Println(recored.GetProgress(), filepath.Ext(path))
	}
	// 如果文件类型属于EXTS
	if inArray(filepath.Ext(path), EXTS) {
		// 判断是否是竖图
		result, result_err := executor.IsVertical(path)
		if result_err != nil {
			fmt.Println(recored.GetProgress(), result_err)
		}
		if result {
			err := executor.CopyFile(path, VPath)
			if err != nil {
				fmt.Println(recored.GetProgress(), err)
			}
			return nil
		}
		err := executor.CopyFile(path, HPath)
		if err != nil {
			fmt.Println(recored.GetProgress(), err)
		}
	}
	return nil
}

func scanf() {
	// 先遍历目录 统计进度
	filepath.WalkDir(FaterPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		recored.Progresss.Total++
		return nil
	})
	filepath.Walk(FaterPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			fmt.Println(recored.GetProgress(), "dir:", path)
			return nil
		}
		// 如果该文件不在文件列表中，则添加
		recored.Progresss.Done++
		if !inArray(path, Files) {
			Files = append(Files, path)
			return addAndMoveFile(path)
		}
		return nil

	})
}

func Init() {
	scanf()
	fmt.Println(Types)
}
