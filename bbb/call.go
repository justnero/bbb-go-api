package bbb

import (
	"encoding/xml"
	"github.com/valyala/fasthttp"
)

func (api *Api) Get(method string, params interface{}, body interface{}) error {
	url := api.GenerateUrl(method, params)
	println(url)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)

	if err := api.http.Do(req, resp); err != nil {
		return err
	}

	return xml.Unmarshal(resp.Body(), body)
}
