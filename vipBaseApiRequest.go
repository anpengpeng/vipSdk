package vipSdk

type VipBaseApiRequest interface {
	GetConfig() *VipBaseConfig
	SetConfig(*VipBaseConfig)
	GetParamsObject() interface{}
	GetServiceName() string
	GetMethodName() string
}

type BaseVipApiRequest struct {
	config *VipBaseConfig
	client *VipClient
}

func (b *BaseVipApiRequest) GetConfig() *VipBaseConfig {
	return b.config
}

func (b *BaseVipApiRequest) SetConfig(config *VipBaseConfig) {
	b.config = config
}

func (b *BaseVipApiRequest) GetClient() *VipClient {
	return b.client
}

func (b *BaseVipApiRequest) SetClient(client *VipClient) {
	b.client = client
}
