package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Resp struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var csrfToken string

// func TestCsrfToken(t *testing.T) {
// 	router := GetApp()
// 	resp := Resp{}

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/csrftoken", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	respStr := w.Body.String()
// 	fmt.Println(respStr)
// 	_ = json.Unmarshal([]byte(respStr), &resp)
// 	assert.Equal(t, 0, resp.Ret)
// 	csrfToken = resp.Data.(map[string]interface{})["X-CSRF-TOKEN"].(string)
// 	fmt.Println(csrfToken)
// }

func getCsrfToken(t *testing.T) string {
	router := GetApp()
	resp := Resp{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/csrftoken", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	respStr := w.Body.String()
	fmt.Println(respStr)
	_ = json.Unmarshal([]byte(respStr), &resp)
	assert.Equal(t, 0, resp.Ret)
	csrfToken = resp.Data.(map[string]interface{})["X-CSRF-TOKEN"].(string)
	fmt.Println(csrfToken)
	return csrfToken
}

func TestRegist(t *testing.T) {
	router := GetApp()
	resp := Resp{}
	jsonByte, _ := json.Marshal(map[string]string{"username": "test3", "password": "1234", "email": "chshbox@gmail.com"})

	req, _ := http.NewRequest("POST", "/api/regist", bytes.NewReader(jsonByte))
	req.Header.Set("X-CSRF-TOKEN", getCsrfToken(t))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	respStr := w.Body.String()
	fmt.Println(respStr)
	_ = json.Unmarshal([]byte(respStr), &resp)
	assert.Equal(t, 0, resp.Ret)
}
