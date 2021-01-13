/**
 * @Author: li.zhang
 * @Description:
 * @File:  http
 * @Version: 1.0.0
 * @Date: 2020/12/11 下午3:27
 */
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/zlsmd/zchat/server/model"
	"io"
	"net/http"
)

type BaseRespDao struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Form)
	username := r.Form.Get("username")
	if username == "" {
		jsonReturn(w, BaseRespDao{
			Code: 10001,
			Msg:  "username is empty",
			Data: nil,
		})
		return
	}

	password := r.Form.Get("password")
	if password == "" {
		jsonReturn(w, BaseRespDao{
			Code: 10002,
			Msg:  "password is empty",
			Data: nil,
		})
		return
	}

	userMod := model.User{}
	ok, token, _ := userMod.CheckLogin(username, password)
	if ok {
		jsonReturn(w, BaseRespDao{
			Code: 0,
			Msg:  "ok",
			Data: token,
		})
	} else {
		jsonReturn(w, BaseRespDao{
			Code: 10000,
			Msg:  "fail",
			Data: nil,
		})
	}
	return
}

func jsonReturn(w io.Writer, value interface{}) {
	json.NewEncoder(w).Encode(value)
}
