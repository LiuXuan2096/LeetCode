package Go_LeetCode

import (
	"os"
	"path/filepath"
)

/*
给定一个文件目录的路径，
写一个函数统计这个目录下所有的文件数量并返回，
隐藏文件也算，但是文件夹不算
*/
func GetFileNumber(folderpath string) int {
	root, err := os.Stat(folderpath)
	if err != nil || (!root.IsDir() && !root.Mode().IsRegular()) {
		return 0
	}
	if root.Mode().IsRegular() {
		return 1
	}
	var stack []string
	stack = append(stack, folderpath)
	files := 0
	for len(stack) > 0 {
		folder := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fileInfo, err := os.ReadDir(folder)
		if err != nil {
			continue
		}
		for _, file := range fileInfo {
			if file.IsDir() {
				stack = append(stack, filepath.Join(folder, file.Name()))
			} else {
				files++
			}
		}
	}
	return files
}
