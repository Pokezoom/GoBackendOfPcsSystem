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
type DelVideoRes struct {
	UserID  int `json:"userId"`
	VideoId int `json:"videoId"`
}
type VideoListReq struct {
	PageSize  int    `json:"pageSize"`
	PageNum   int    `json:"pageNum"`
	Name      string `json:"name,omitempty"`
	Class     string `json:"class,omitempty"`
	Subject   string `json:"subject,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}
