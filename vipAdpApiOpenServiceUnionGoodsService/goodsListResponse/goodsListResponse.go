package goodsListResponse

import "github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionGoodsService/getByGoodsIdsResponse"

type UnionGoodsServiceGoodsListResponse struct {
	ReturnCode string                                   `json:"returnCode"`
	Result     UnionGoodsServiceGoodsListResponseResult `json:"result"`
}

type UnionGoodsServiceGoodsListResponseResult struct {
	Total         int                                  `json:"total"`
	GoodsInfoList []getByGoodsIdsResponse.VipGoodsInfo `json:"goodsInfoList"`
	PageSize      int                                  `json:"pageSize"`
	SortFields    []struct {
		FieldName string `json:"fieldName"`
		FieldDesc string `json:"fieldDesc"`
	} `json:"sortFields"`
	Page int `json:"page"`
}
