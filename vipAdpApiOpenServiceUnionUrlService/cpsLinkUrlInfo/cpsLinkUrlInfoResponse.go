package cpsLinkUrlInfo

type UrlInfoList struct {
	NoEvokeUrl     string `json:"noEvokeUrl"`     //CPS短链接：不唤起快应用
	VipQuickAppUrl string `json:"vipQuickAppUrl"` //唯品会快应用链接
	VipWxUrl       string `json:"vipWxUrl"`       //唯品会小程序链接：仅在根据商品id获取时返回
	VipWxCode      string `json:"vipWxCode"`      //唯品会小程序码：仅在根据商品id获取时返回，需获取小程序码高级权限
	DeeplinkUrl    string `json:"deeplinkUrl"`    //CPS Deeplink链接
	LongUrl        string `json:"longUrl"`        //CPS长连接
	Source         string `json:"source"`         //链接生成的数据源： 如果根据商品id生成链接，该值商品id， 如果根据链接生成链接，该值为唯品会链接
	UlUrl          string `json:"ulUrl"`          //CPS通用连接
	Url            string `json:"url"`            //CPS短链接
	TraFrom        string `json:"traFrom"`        //小程序CPS参数：通用小程序跟单参数
	NoEvokeLongUrl string `json:"noEvokeLongUrl"` //CPS长链接：不唤起快应用
}
