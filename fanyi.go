package baidufanyi

import (
	"net/http"
)

type Fanyi struct {
	AppID     string
	SecretKey string
	client    *http.Client
}

type ClientOptionFunc func(*Fanyi)

func WithCredential(appID, appSecret string) ClientOptionFunc {
	return func(r *Fanyi) {
		r.AppID = appID
		r.SecretKey = appSecret
	}
}

func New(options ...ClientOptionFunc) *Fanyi {
	return newClient(options)
}

func newClient(options []ClientOptionFunc) *Fanyi {
	r := &Fanyi{}
	for _, v := range options {
		if v != nil {
			v(r)
		}
	}

	if r.client == nil {
		r.client = &http.Client{}
	}

	return r
}
