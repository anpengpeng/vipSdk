#   About vip-sdk
唯品会开放平台golang版本sdk

### 唯品会api地址
<https://vop.vip.com/home#/api/service/list>


##  demo

    var (
        AppKey      string = "Your AppKey"
        AppSecret   string = "Your AppSecret"
        AccessToken string = "Your accessToken"
    )

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

### ps

-   目前仅只支持了几个api，其他的需要自己手动接入，仿照写即可。