package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int
	Data interface{}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	mp := make(map[string]string)
	mp["id"] = "226"
	mp["name"] = "许磊"

	res := &Response{
		Code: 200,
		Data: mp,
	}

	msg, err := json.Marshal(res)
	if err != nil {
		_, err := fmt.Fprint(w, "内部错误")
		if err != nil {
			return
		}
		return
	}

	_, _ = w.Write(msg)
}
