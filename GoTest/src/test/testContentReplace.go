package test

import "path/filepath"

//替换文件内容
type TestContentReplace struct {
	filePath, oldText, newText string
}

func NewTestContentReplace(filePath, oldText, newText string) {
	return &TestContentReplace{
		filePath: filePath,
		oldText:  oldText,
		newText:  newText,
	}
}

func (this *TestContentReplace) Do() {
	dir := filepath.Dir(this.filePath)
	filepath.Walk()
}
