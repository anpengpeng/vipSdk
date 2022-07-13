package test

import (
	"fmt"
	"github.com/anpengpeng/vipSdk"
	"github.com/anpengpeng/vipSdk/vipAdpApiOpenServiceUnionGoodsService/queryRequest"
	"testing"
)

const (
	AppKey      string = "Your AppKey"
	AppSecret   string = "Your AppSecret"
	AccessToken string = "Your accessToken"
)

func TestVipSdk(t *testing.T) {
	config := &vipSdk.VipBaseConfig{
		AppKey:         AppKey,
		AppSecret:      AppSecret,
		ApiVersion:     "1.0.0",
		AccessToken:    AccessToken,
		BaseUrl:        "https://vop.vipapis.com",
		ConnectTimeout: 3000,
	}

	request := queryRequest.New(config)
	params := request.GetParams()
	param := queryRequest.QueryRequest{
		Keyword:   "牛奶",
		RequestId: "requestId:1", //官方建议uuid
	}
	params.QueryRequest = param
	execute, err := request.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(execute)
}
