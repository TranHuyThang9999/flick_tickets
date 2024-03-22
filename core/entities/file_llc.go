package entities

type UploadResponse struct {
	Result    Result `json:"result"`
	ID        string `json:"id"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
