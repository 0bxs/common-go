package dto

import (
	"common/src/utils/strs"
)

type PageDto struct {
	PageNum   int    // 当前页码
	PageSize  int    `binding:"required,gte=1,lte=100"`  // 每页展示数量
	OrderBy   string `binding:"required"`                // 排序字段
	OrderType string `binding:"required,oneof=DESC ASC"` // 排序方式 DESC/ASC
}

func (dto *PageDto) GetOffset() int {
	return dto.PageNum * dto.PageSize
}

func (dto *PageDto) GetOrder() string {
	return strs.Join(dto.OrderBy, " ", dto.OrderType)
}
