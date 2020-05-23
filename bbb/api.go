package bbb

import "github.com/valyala/fasthttp"

type Api struct {
	base string
	salt string
	http *fasthttp.Client
}

func CreateApi(base string, salt string) Api {
	return Api{
		base: base,
		salt: salt,
		http: &fasthttp.Client{},
	}
}
