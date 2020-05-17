package data

type Meetings struct {
	Meeting []Meeting `xml:"meeting"`
}

type Meeting struct {
	MeetingName           string        `xml:"meetingName"`
	MeetingID             string        `xml:"meetingID"`
	InternalMeetingID     string        `xml:"internalMeetingID"`
	CreateTime            string        `xml:"createTime"`
	CreateDate            string        `xml:"createDate"`
	VoiceBridge           string        `xml:"voiceBridge"`
	DialNumber            string        `xml:"dialNumber"`
	AttendeePW            string        `xml:"attendeePW"`
	ModeratorPW           string        `xml:"moderatorPW"`
	Running               bool          `xml:"running"`
	Duration              int           `xml:"duration"`
	HasUserJoined         bool          `xml:"hasUserJoined"`
	Recording             bool          `xml:"recording"`
	HasBeenForciblyEnded  bool          `xml:"hasBeenForciblyEnded"`
	StartTime             string        `xml:"startTime"`
	EndTime               string        `xml:"endTime"`
	ParticipantCount      int           `xml:"participantCount"`
	ListenerCount         int           `xml:"listenerCount"`
	VoiceParticipantCount int           `xml:"voiceParticipantCount"`
	VideoCount            int           `xml:"videoCount"`
	MaxUsers              int           `xml:"maxUsers"`
	ModeratorCount        int           `xml:"moderatorCount"`
	Attendees             Attendees     `xml:"attendees"`
	Metadata              StringMap     `xml:"metadata"`
	IsBreakout            bool          `xml:"isBreakout"`
	ParentMeetingID       string        `xml:"parentMeetingID"`
	Sequence              int           `xml:"sequence"`
	IsFreeJoin            bool          `xml:"freeJoin"`
	BreakoutRooms         BreakoutRooms `xml:"breakoutRooms"`
}
