package getByGoodsIdsResponse

type UnionGoodsServiceGetByGoodsIdsResponse struct {
	ReturnCode string         `json:"returnCode"`
	Result     []VipGoodsInfo `json:"result"`
}

type VipGoodsInfo struct {
	MarketPrice            string     `json:"marketPrice"`    //市场价(元)
	CommissionRate         string     `json:"commissionRate"` //佣金比例（%）
	WhiteImage             string     `json:"whiteImage"`     //商品透明底图
	GoodsId                string     `json:"goodsId"`
	Discount               string     `json:"discount"`
	GoodsCarouselPictures  []string   `json:"goodsCarouselPictures"` //商品轮播图
	GoodsDetailPictures    []string   `json:"goodsDetailPictures"`   //商品详情图片
	CategoryName           string     `json:"categoryName"`
	HaiTao                 int        `json:"haiTao"`
	PrepayInfo             PrepayInfo `json:"prepayInfo"` //预售信息
	Cat2NdName             string     `json:"cat2ndName"`
	IsSubsidyActivityGoods bool       `json:"isSubsidyActivityGoods"`
	Cat1StName             string     `json:"cat1stName"`
	VipPrice               string     `json:"vipPrice"`   //唯品价(元)
	Commission             string     `json:"commission"` //佣金金额（元）
	Sn                     string     `json:"sn"`
	Cat1StId               int        `json:"cat1stId"`
	GoodsName              string     `json:"goodsName"`
	StoreServiceCapability struct {   //商品所属店铺服务能力评价
		StoreScore    string `json:"storeScore"`
		StoreRankRate string `json:"storeRankRate"`
	} `json:"storeServiceCapability"`
	BrandName       string   `json:"brandName"`
	BrandLogoFull   string   `json:"brandLogoFull"`
	BrandStoreSn    string   `json:"brandStoreSn"`
	SellTimeFrom    int64    `json:"sellTimeFrom"`
	SchemeStartTime int64    `json:"schemeStartTime"`
	CommentsInfo    struct { //商品评价信息
		Comments          int    `json:"comments"`
		GoodCommentsShare string `json:"goodCommentsShare"`
	} `json:"commentsInfo"`
	SaleStockStatus int    `json:"saleStockStatus"` //商品库存状态：1-已抢光，2-有库存，3-有机会,当列表查询库存或者查询商品详情时返回
	SchemeEndTime   int64  `json:"schemeEndTime"`
	SourceType      int    `json:"sourceType"` //0-自营，1-MP
	SellTimeTo      int64  `json:"sellTimeTo"`
	BrandId         int    `json:"brandId"`
	GoodsThumbUrl   string `json:"goodsThumbUrl"` //商品缩略图
	Cat2NdId        int    `json:"cat2ndId"`
	SpuId           string `json:"spuId"`
	StoreInfo       struct {
		StoreName string `json:"storeName"`
		StoreId   string `json:"storeId"`
	} `json:"storeInfo"`
	GoodsMainPicture string                 `json:"goodsMainPicture"`
	DestUrl          string                 `json:"destUrl"`
	CategoryId       int                    `json:"categoryId"`
	Status           int                    `json:"status"`          //商品状态：0-下架，1-上架
	ExclusiveCoupon  ChannelExclusiveCoupon `json:"exclusiveCoupon"` //渠道专属红包：目前仅开放单品券，没有则返回空
	CouponInfo       PMSCouponInfo          `json:"couponInfo"`      //红包信息
}

type PrepayInfo struct {
	IsPrepay             int    `json:"isPrepay"`             //是否预付商品:0-否，1-是
	PrepayPrice          string `json:"prepayPrice"`          //预付到手价：元
	FirstAmount          string `json:"firstAmount"`          //预付首款金额：元
	LastAmount           string `json:"lastAmount"`           //预付尾款金额：元
	PrepayFavAmount      string `json:"prepayFavAmount"`      //预付优惠金额: 元
	DeductionPrice       string `json:"deductionPrice"`       //抵扣价格(首款+优惠金额)：元
	PrepayDiscount       string `json:"prepayDiscount"`       //预付折扣：(唯品价-优惠金额)/唯品价 保留两位小数的数字字符串
	PrepayFirstStartTime int    `json:"prepayFirstStartTime"` //首款支付开始时间:时间戳,单位毫秒
	PrepayFirstEndTime   int    `json:"prepayFirstEndTime"`   //首款支付结束时间:时间戳,单位毫秒
	PrepayLastStartTime  int    `json:"prepayLastStartTime"`  //尾款支付开始时间:时间戳,单位毫秒
	PrepayLastEndTime    int    `json:"prepayLastEndTime"`    //尾款支付结束时间:时间戳,单位毫秒
}

type ChannelExclusiveCoupon struct {
	CouponNo          string `json:"couponNo"`          //优惠券批次号
	CouponName        string `json:"couponName"`        //优惠劵名称
	Fav               string `json:"fav"`               //优惠金额：单位-元,查询详情时返回
	Buy               string `json:"buy"`               //使用门槛：单位-元，查询详情时返回
	ActivateBeginTime int    `json:"activateBeginTime"` //券激活开始时间,毫秒级时间戳
	ActivateEndTime   int    `json:"activateEndTime"`   //是券激活结束时间,毫秒级时间戳
	UseBeginTime      int    `json:"useBeginTime"`      //使用开始时间，毫秒级时间戳
	UseEndTime        int    `json:"useEndTime"`        //使用结束时间，毫秒级时间戳
	TotalAmount       int    `json:"totalAmount"`       //生成劵的总量：查询详情时返回
	ActivedAmount     int    `json:"activedAmount"`     //劵已激活的数量：查询详情时返回
	ReceiveUrl        string `json:"receiveUrl"`        //专属领券页地址
}

type PMSCouponInfo struct {
	CouponNo          string `json:"couponNo"`          //优惠券批次号
	CouponName        string `json:"couponName"`        //优惠劵名称
	Fav               string `json:"fav"`               //优惠金额：单位-元,查询详情时返回
	Buy               string `json:"buy"`               //使用门槛：单位-元，查询详情时返回
	ActivateBeginTime int    `json:"activateBeginTime"` //券激活开始时间,毫秒级时间戳
	ActivateEndTime   int    `json:"activateEndTime"`   //是券激活结束时间,毫秒级时间戳
	UseBeginTime      int    `json:"useBeginTime"`      //使用开始时间，毫秒级时间戳
	UseEndTime        int    `json:"useEndTime"`        //使用结束时间，毫秒级时间戳
	TotalAmount       int    `json:"totalAmount"`       //生成劵的总量：查询详情时返回
	ActivedAmount     int    `json:"activedAmount"`     //劵已激活的数量：查询详情时返回
	CouponType        int    `json:"couponType"`        //券类型,(1:买赠 2:满减 3:折扣 4:免邮 5:多减多减)
}
