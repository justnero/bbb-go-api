package data

type Recordings struct {
	Recording []Recording `xml:"recording"`
}

type Recording struct {
	RecordID     string   `xml:"recordID"`
	MeetingID    string   `xml:"meetingID"`
	Name         string   `xml:"name"`
	Published    string   `xml:"published"`
	State        string   `xml:"state"`
	StartTime    string   `xml:"startTime"`
	EndTime      string   `xml:"endTime"`
	Participants string   `xml:"participants"`
	MetaData     Metadata `xml:"metadata"`
	Playback     Playback `xml:"playback"`
}

type Playback struct {
	Format []Format `xml:"format"`
}

type Format struct {
	Type   string   `xml:"type"`
	Url    string   `xml:"url"`
	Length string   `xml:"length"`
	Images []string `xml:"preview>images>image"`
}

type Metadata struct {
	Title       string `xml:"title"`
	Subject     string `xml:"subject"`
	Description string `xml:"description"`
	Creator     string `xml:"creator"`
	Contributor string `xml:"contributor"`
	Language    string `xml:"language"`
}
