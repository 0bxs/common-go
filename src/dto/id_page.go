package dto

type IdPageDto struct {
	LastId   uint64 `json:"lastId"`
	PageSize int    `json:"pageSize" binding:"required,gte=1,lte=100"` // 每页展示数量
}
