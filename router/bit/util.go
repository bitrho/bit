package bit

import (
	"encoding/json"
	"fmt"
	"net/http"
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
