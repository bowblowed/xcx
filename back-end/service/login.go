package service

import (
	"back-end/config"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type WXLoginResponse struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func WxLogin(code string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", config.Wx.AppiD, config.Wx.Appsecret, code)
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("wx server connect error")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	var wxResp WXLoginResponse
	if err := json.Unmarshal(body, &wxResp); err != nil {
		return "", errors.New("failed to parse WeChat response")
	}

	if wxResp.ErrCode != 0 {
		return "", errors.New(wxResp.ErrMsg)
	}
	user, err := GetUserByOpenId(wxResp.Openid)
	if err != nil {
		CreateWxUserDefault(wxResp.Openid)
	}
	return GenerateToken(fmt.Sprintf("%d", user.ID))
}
