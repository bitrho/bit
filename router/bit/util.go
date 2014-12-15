package bit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ajaxReturn(msg string, redirect string, w http.ResponseWriter) {
	data := make(map[string]string)
	if msg != "" {
		data["msg"] = msg
	}
	if redirect != "" {
		data["redirect"] = redirect
	}
	re, _ := json.Marshal(data)
	fmt.Fprintln(w, string(re))
	panic(msg)
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
