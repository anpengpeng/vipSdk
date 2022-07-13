package genByGoodsIdRequest

import (
	"encoding/json"
	"github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionUrlService/genByGoodsIdResponse"
)

type UnionUrlServiceGenByGoodsIdRequest struct {
	vipSdk.BaseVipApiRequest
	Param *UnionUrlServiceGenByGoodsIdParam
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetMethodName() string {
	return "genByGoodsId"
}

func New(config *vipSdk.VipBaseConfig) *UnionUrlServiceGenByGoodsIdRequest {
	request := UnionUrlServiceGenByGoodsIdRequest{
		Param: &UnionUrlServiceGenByGoodsIdParam{},
	}
	request.SetConfig(config)
	request.SetClient(vipSdk.DefaultVipApiClient)

	return &request
}

func (v *UnionUrlServiceGenByGoodsIdRequest) Execute() (*genByGoodsIdResponse.GenByGoodsIdResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &genByGoodsIdResponse.GenByGoodsIdResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetParams() *UnionUrlServiceGenByGoodsIdParam {
	return v.Param
}

type UnionUrlServiceGenByGoodsIdParam struct {
	GoodsIdList []string `json:"goodsIdList"`
	ChanTag     string   `json:"chanTag"`     //渠道标识
	StatParam   string   `json:"statParam"`   //自定义渠道统计参数
	RequestId   string   `json:"requestId"`   //请求id：UUID
	GenShortUrl bool     `json:"genShortUrl"` //是否压缩生成的链接,默认false(理论上生成的链接无需压缩也能满足大部分推广情景，如非必要，请勿选择压缩，压缩的链接有有效期(目前是一个月))
}
