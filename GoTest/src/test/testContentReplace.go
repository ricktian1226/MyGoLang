package test

import (
	"fmt"
	"io/ioutil"
	"log"
	//	"os"
	"path/filepath"
	"strings"
	"time"
)

//替换文件内容
type TestContentReplace struct {
	dir, fileName, oldText, newText string
}

func NewTestContentReplace(dir, fileName, oldText, newText string) *TestContentReplace {
	return &TestContentReplace{
		dir:      dir,
		fileName: fileName,
		oldText:  oldText,
		newText:  newText,
	}
}

func (this *TestContentReplace) Do() {
	var (
		err                                              error
		fileName, fileNameOnly, oldFileName, newFileName string
		now                                              = time.Now()
		buf                                              []byte
	)

	//当前目录
	fileName = filepath.Base(this.fileName)
	fileNameOnly = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	oldFileName = this.dir + "/" + this.fileName
	newFileName = fmt.Sprintf("%s/%s%d%d%d%s", this.dir, fileNameOnly, now.Hour(), now.Minute(), now.Second(), filepath.Ext(fileName))
	log.Printf("newFileName %s", newFileName)

	if buf, err = ioutil.ReadFile(oldFileName); err != nil {
		log.Fatalf("读取文件 %s 失败, %v\n", err)
		return
	} else {
		//替换文件内容
		newContent := strings.Replace(string(buf), this.oldText, this.newText, -1)
		//保存到新文件中
		if err = ioutil.WriteFile(newFileName, []byte(newContent), 0); err != nil {
			log.Fatalf("写文件 %s 失败, %v\n", err)
			return
		}
	}
}
