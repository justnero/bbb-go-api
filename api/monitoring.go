package api

import (
	"github.com/justnero/bbb-go-api/data"
)

func (api *Api) GetMeetings() (data.GetMeetingsResponse, error) {
	var req data.GetMeetingsRequest
	var resp data.GetMeetingsResponse

	err := api.Get("getMeetings", req, &resp)

	return resp, err
}

func (api *Api) GetMeetingInfo(meetingId string) (data.GetMeetingInfoResponse, error) {
	req := data.GetMeetingInfoRequest{
		MeetingId: meetingId,
	}
	var resp data.GetMeetingInfoResponse

	err := api.Get("getMeetingInfo", req, &resp)

	return resp, err
}

func (api *Api) IsMeetingRunning(meetingId string) (data.IsMeetingRunningResponse, error) {
	req := data.IsMeetingRunningRequest{
		MeetingId: meetingId,
	}
	var resp data.IsMeetingRunningResponse

	err := api.Get("isMeetingRunning", req, &resp)

	return resp, err
}
