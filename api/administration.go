package api

import "github.com/justnero/bbb-go-api/data"

func (api *Api) Join(req data.JoinRequest) (data.JoinResponse, error) {
	var resp data.JoinResponse

	err := api.Get("join", req, &resp)

	return resp, err
}
