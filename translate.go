package baidufanyi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (r *Fanyi) Translate(text string, from, to Language) ([]*TranslateResult, error) {
	uri := "https://fanyi-api.baidu.com/api/trans/vip/translate"
	salt := r.salt()
	u := url.Values{}
	u.Add("q", text)
	u.Add("from", string(from))
	u.Add("to", string(to))
	u.Add("appid", r.AppID)
	u.Add("salt", salt)
	u.Add("sign", r.sign(text, salt))

	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(u.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ts := new(translateResp)
	if err = json.Unmarshal(bs, ts); err != nil {
		return nil, err
	}
	if ts.ErrorCode != "" {
		return nil, fmt.Errorf(ts.ErrorMsg)
	}
	return ts.Result, nil
}

type translateResp struct {
	From      string             `json:"from"`
	To        string             `json:"to"`
	Result    []*TranslateResult `json:"trans_result"`
	ErrorCode string             `json:"error_code"`
	ErrorMsg  string             `json:"error_msg"`
}

type TranslateResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func (r *Fanyi) salt() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 10)
	randIns := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		result[i] = bytes[randIns.Intn(len(bytes))]
	}
	return string(result)
}

func (r *Fanyi) sign(text, salt string) string {
	txt := r.AppID + text + salt + r.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(txt))
	return hex.EncodeToString(ctx.Sum(nil))
}
