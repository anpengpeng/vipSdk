package genByVipUrlRequest

import (
	"encoding/json"
	"github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionUrlService/genByVipUrlResponse"
)

type UnionUrlServiceGenByVipUrlRequest struct {
	vipSdk.BaseVipApiRequest
	Param *UnionUrlServiceGenByVipUrlParam
}

func (v *UnionUrlServiceGenByVipUrlRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

func (v *UnionUrlServiceGenByVipUrlRequest) GetMethodName() string {
	return "genByVIPUrl"
}

func New(config *vipSdk.VipBaseConfig) *UnionUrlServiceGenByVipUrlRequest {
	request := UnionUrlServiceGenByVipUrlRequest{
		Param: &UnionUrlServiceGenByVipUrlParam{},
	}
	request.SetConfig(config)
	request.SetClient(vipSdk.DefaultVipApiClient)

	return &request
}

func (v *UnionUrlServiceGenByVipUrlRequest) Execute() (*genByVipUrlResponse.GenByVipUrlResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &genByVipUrlResponse.GenByVipUrlResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionUrlServiceGenByVipUrlRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionUrlServiceGenByVipUrlRequest) GetParams() *UnionUrlServiceGenByVipUrlParam {
	return v.Param
}

type UnionUrlServiceGenByVipUrlParam struct {
	UrlList       []string      `json:"urlList"`
	ChanTag       string        `json:"chanTag"`       //渠道标识
	StatParam     string        `json:"statParam"`     //自定义渠道统计参数
	RequestId     string        `json:"requestId"`     //请求id：UUID
	UrlGenRequest UrlGenRequest `json:"urlGenRequest"` //是否压缩生成的链接,默认false(理论上生成的链接无需压缩也能满足大部分推广情景，如非必要，请勿选择压缩，压缩的链接有有效期(目前是一个月))
}

type UrlGenRequest struct {
	TargetType      string   `json:"targetType"`      //落地页类型
	TargetValueList []string `json:"targetValueList"` //落地页信息
	EvokeQuickApp   bool     `json:"evokeQuickApp"`   //是否自动唤起快应用,默认false
	GenShortUrl     bool     `json:"genShortUrl"`     //是否压缩生成的链接,默认false(理论上生成的链接无需压缩也能满足大部分推广情景，如非必要，请勿选择压缩，压缩的链接有有效期(目前是一个月))
}
