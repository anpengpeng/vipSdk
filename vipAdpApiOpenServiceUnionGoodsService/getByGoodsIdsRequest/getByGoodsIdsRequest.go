package getByGoodsIdsRequest

import (
	"encoding/json"
	"github.com/anpengpeng/vipSdk"
	"github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionGoodsService/getByGoodsIdsResponse"
)

type UnionGoodsServiceGetByGoodsIdsRequest struct {
	vipSdk.BaseVipApiRequest
	Param *UnionGoodsServiceGetByGoodsIdsParam
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetMethodName() string {
	return "getByGoodsIds"
}

func New(config *vipSdk.VipBaseConfig) *UnionGoodsServiceGetByGoodsIdsRequest {
	request := UnionGoodsServiceGetByGoodsIdsRequest{
		Param: &UnionGoodsServiceGetByGoodsIdsParam{},
	}
	request.SetConfig(config)
	request.SetClient(vipSdk.DefaultVipApiClient)

	return &request
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) Execute() (*getByGoodsIdsResponse.UnionGoodsServiceGetByGoodsIdsResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &getByGoodsIdsResponse.UnionGoodsServiceGetByGoodsIdsResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetParams() *UnionGoodsServiceGetByGoodsIdsParam {
	return v.Param
}

type UnionGoodsServiceGetByGoodsIdsParam struct {
	GoodsIdList []string `json:"goodsIdList"`
	ChanTag     string   `json:"chanTag"`   //渠道标识
	RequestId   string   `json:"requestId"` //请求id：UUID
}
