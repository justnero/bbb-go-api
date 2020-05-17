package data

type Response struct {
	ReturnCode string `xml:"returncode"`
	MessageKey string `xml:"messageKey"`
	Message    string `xml:"message"`
}

type GetMeetingsResponse struct {
	Meetings Meetings `xml:"meetings"`
	Response
}

type GetMeetingInfoResponse struct {
	Meeting
	Response
}

type IsMeetingRunningResponse struct {
	Running bool `xml:"running"`
	Response
}

type JoinResponse struct {
	MeetingId    string `xml:"meeting_id"`
	UserId       string `xml:"user_id"`
	AuthToken    string `xml:"auth_token"`
	SessionToken string `xml:"session_token"`
	GuestStatus  string `xml:"guestStatus"`
	Url          string `xml:"url"`
	Response
}
