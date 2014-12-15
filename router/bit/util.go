package bit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func AjaxReturn(code int, result interface{}, w http.ResponseWriter) {
	data := make(map[string]interface{})
	msg := SysStatus(code)
	data["status"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	data["result"] = result
	re, _ := json.Marshal(data)
	fmt.Fprintln(w, string(re))
	panic(msg)
}

func SysStatus(code int) (msg string) {
	msg, ok := Stat[code]
	if !ok {
		msg = Stat[4004]
	}
	return
}

func ReadFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		log.Println(string(err.Error()))
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	return string(fd)
}

func IsDir(fileName string) bool {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsFile(fileName string) bool {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func ScanDir(directory string) []string {
	file, err := os.Open(directory)
	if err != nil {
		return []string{}
	}
	names, err := file.Readdirnames(-1)
	if err != nil {
		return []string{}
	}
	return names
}
