package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TestJson struct {
	dir string
}

func NewTestJson(dir string) *TestJson {
	return &TestJson{
		dir: dir,
	}
}

func (this *TestJson) Do() {
	fileName := this.dir + "/" + "test.json"
	var content map[string]interface{}
	if data, err := ioutil.ReadFile(fileName); err != nil {
		log.Fatalf("读取文件 %s 失败，%v", fileName, err)
		return
	} else {
		if err = json.Unmarshal(data, &content); err != nil {
			log.Fatalf("解码内容失败 %v", err)
			return
		} else {
			for k, v := range content {
				fmt.Sprintf("%v : %v\n", k, v)
			}
		}
	}
}
