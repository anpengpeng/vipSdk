package vipSdk

import (
	"encoding/json"
	"fmt"
	"github.com/anpengpeng/vipSdk/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type VipClient struct {
}

func NewVipClient() *VipClient {
	return &VipClient{}
}

func (j *VipClient) Request(params VipBaseApiRequest, isNeedAuth bool) (string, error) {
	timeStr := fmt.Sprintf("%d", time.Now().Unix())
	service := params.GetServiceName()
	method := params.GetMethodName()
	format := "json"
	apiVersion := params.GetConfig().ApiVersion
	appKey := params.GetConfig().AppKey
	appSecret := params.GetConfig().AppSecret
	timeout := params.GetConfig().ConnectTimeout

	busiParams, _ := json.Marshal(params.GetParamsObject())

	accessToken := ""
	if isNeedAuth {
		accessToken = params.GetConfig().AccessToken
	}
	sign := utils.CreateRequestSign(accessToken, appKey, format, "", method, service, timeStr, apiVersion, string(busiParams), appSecret)
	fields := map[string]string{
		"appKey":    params.GetConfig().AppKey,
		"format":    format,
		"method":    method,
		"service":   service,
		"sign":      sign,
		"timestamp": timeStr,
		"version":   apiVersion,
	}
	if accessToken != "" {
		fields["accessToken"] = accessToken
	}

	return HttpPost(params.GetConfig().BaseUrl+"?"+utils.GetQueryString(fields), string(busiParams), timeout)
}

func HttpPost(urls string, data string, timeout int64) (string, error) {
	do := &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}

	resp, err := do.Post(urls, "application/json,charset=utf-8",
		strings.NewReader(data))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

var DefaultVipApiClient *VipClient = NewVipClient()
