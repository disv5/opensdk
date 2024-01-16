package report

type GetFileResponse struct {
	ReportFileType ReportFileType `json:"reportFileType"`
}

type ReportFileType struct {
	TaskID int64 `json:"taskId"`
	Ready  bool  `json:"ready"`
}
