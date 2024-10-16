package controller

type Response struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data"`
}
