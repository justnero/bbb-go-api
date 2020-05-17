package data

type GetMeetingsRequest struct {
}

type GetMeetingInfoRequest struct {
	MeetingId string `url:"meetingID"`
}

type IsMeetingRunningRequest struct {
	MeetingId string `url:"meetingID"`
}

/**
 * Be sure making Redirect true if you need it
 */
type JoinRequest struct {
	FullName      string `url:"fullName"`
	MeetingId     string `url:"meetingID"`
	Password      string `url:"password"`
	CreateTime    string `url:"createTime,omitempty"`
	UserId        string `url:"userID,omitempty"`
	WebVoiceConf  string `url:"webVoiceConf,omitempty"`
	ConfigToken   string `url:"configToken,omitempty"`
	DefaultLayout string `url:"defaultLayout,omitempty"`
	AvatarUrl     string `url:"avatarURL,omitempty"`
	Redirect      bool   `url:"redirect"`
	ClientUrl     string `url:"clientURL,omitempty"`
	JoinViaHtml5  bool   `url:"joinViaHtml5,omitempty"`
	Guest         bool   `url:"guest,omitempty"`
}
