package vipLinkCheckResponse

type VipLinkCheckResponse struct {
	ReturnCode string                     `json:"returnCode"`
	Result     VipLinkCheckResponseResult `json:"result"`
}

type VipLinkCheckResponseResult struct {
	SuccessMap map[string]VipLinkCheckVO `json:"successMap"`
	FailMap    map[string]VipLinkCheckVO `json:"failMap"`
}

type VipLinkCheckVO struct {
	GoodsId  string `json:"goodsId"`  //商品id
	LandUrl  string `json:"landUrl"`  //落地页url
	BrandId  string `json:"brandId"`  //专场id
	LinkType string `json:"linkType"` //链接类型
}
