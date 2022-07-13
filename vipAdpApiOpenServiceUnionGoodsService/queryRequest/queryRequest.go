package queryRequest

import (
	"encoding/json"
	"github.com/anpengpeng/vipSdk"
	"github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionGoodsService/queryResponse"
)

type UnionGoodsServiceQueryRequest struct {
	vipSdk.BaseVipApiRequest
	Param *UnionGoodsServiceQueryParam
}

func (v *UnionGoodsServiceQueryRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

func (v *UnionGoodsServiceQueryRequest) GetMethodName() string {
	return "query"
}

func New(config *vipSdk.VipBaseConfig) *UnionGoodsServiceQueryRequest {
	request := UnionGoodsServiceQueryRequest{
		Param: &UnionGoodsServiceQueryParam{},
	}
	request.SetConfig(config)
	request.SetClient(vipSdk.DefaultVipApiClient)

	return &request
}

func (v *UnionGoodsServiceQueryRequest) Execute() (*queryResponse.UnionGoodsServiceQueryResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &queryResponse.UnionGoodsServiceQueryResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionGoodsServiceQueryRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionGoodsServiceQueryRequest) GetParams() *UnionGoodsServiceQueryParam {
	return v.Param
}

type UnionGoodsServiceQueryParam struct {
	QueryRequest QueryRequest `json:"request"`
}

type QueryRequest struct {
	Keyword   string `json:"keyword"`
	FieldName string `json:"fieldName"` //排序字段
	Order     int    `json:"order"`     //排序顺序：0-正序，1-逆序，默认正序
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`  //页面大小：默认20,最大50
	RequestId string `json:"requestId"` //请求id：UUID
}
