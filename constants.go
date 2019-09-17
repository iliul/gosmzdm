package main

import "encoding/json"

type Config struct {
	Cookie string `json:"cookie"`
	SCKey  string `json:"sc_key"`
	FakeIP string `json:"fake_id"`
}

const (
	UserAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36"
	Referer    = "https://www.smzdm.com"
	CheckInURL = "https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"
)

type CheckInResult struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	Data      struct {
		AddPoint   int             `json:"add_point"`
		CheckInNum json.RawMessage `json:"checkin_num"`
		Point      int             `json:"point"`
		Exp        int             `json:"exp"`
		Gold       int             `json:"gold"`
		Prestige   int             `json:"prestige"`
		Rank       int             `json:"rank"`
	} `json:"data"`
}

type NotifyResult struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
}
