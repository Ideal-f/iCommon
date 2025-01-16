package common

type PageInfo struct {
	PageSize     int   `json:"page_size"`
	PageIndex    int   `json:"page_index"`
	TotalNumber  int64 `json:"total_number"`
	Limit        int   `json:"limit"`
	Offset       int   `json:"offset"`
	CountUseless bool  `json:"count_useless"`
}
