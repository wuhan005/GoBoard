package main

import (
	"net/http"
)

func (s *Service)errorMsg(code int, msg string, httpStatus int)(int, interface{}){
	return httpStatus, map[string]interface{}{"code": code, "msg": msg}
}

func (s *Service)successMsg(code int, msg string, data interface{}) (int, interface{}){
	return http.StatusOK, map[string]interface{}{"code": code, "msg": msg, "data": data}
}