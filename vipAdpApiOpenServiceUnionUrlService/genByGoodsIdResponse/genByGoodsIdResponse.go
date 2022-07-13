package genByGoodsIdResponse

import "github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionUrlService/cpsLinkUrlInfo"

type GenByGoodsIdResponse struct {
	ReturnCode string        `json:"returnCode"`
	Result     UrlInfoResult `json:"result"`
}

type UrlInfoResult struct {
	UrlInfoLists []cpsLinkUrlInfo.UrlInfoList `json:"urlInfoList"`
}
