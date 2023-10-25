package mode

type UploadReq struct {
	UserID       int    `json:"userId"`
	VideoName    string `json:"videoName"`
	Class        string `json:"class"`
	AcademicYear string `json:"academicYear"`
	Subject      string `json:"subject"`
	Duration     int    `json:"duration"`
}
type UploadRes struct {
	VideoId int `json:"videoId"`
}
