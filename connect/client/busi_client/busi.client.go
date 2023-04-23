package busi_client

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type AuthBo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type AuthVo struct {
	UserId int64 `json:"user_id"`
}

type AuthResp struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data *AuthVo `json:"data"`
}

const (
	authUrl = "http://127.0.0.1:9099/api/user/auth"
)

func Auth(bo *AuthBo) *AuthVo {
	data, err := json.Marshal(bo)
	if err != nil {
		zap.S().Errorf("用户%v鉴权失败%v", bo.UserId, err.Error())
		return nil
	}
	resp, err := http.Post(authUrl, "application/json", bytes.NewReader(data))
	if err != nil {
		zap.S().Errorf("用户%v鉴权失败%v", bo.UserId, err.Error())
		return nil
	}
	defer resp.Body.Close()

	allBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.S().Errorf("用户%v鉴权失败%v", bo.UserId, err.Error())
		return nil
	}

	vo := &AuthResp{}
	err = json.Unmarshal(allBytes, &vo)
	if err != nil {
		zap.S().Errorf("用户%v鉴权失败%v", bo.UserId, err.Error())
		return nil
	}

	if vo.Code != 0 {
		zap.S().Errorf("用户%v鉴权失败%v", bo.UserId, err.Error())
		return nil
	}

	return vo.Data
}
