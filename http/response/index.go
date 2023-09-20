package response

type IndexResponse struct {
	Code    int    `json:"code" xml:"code" yaml:"code" example:"200"`
	Message string `json:"message" xml:"message" yaml:"message" example:"操作成功 or 操作失败"`
	Data    string `json:"data" xml:"data" yaml:"data" example:"success or fail"`
}
