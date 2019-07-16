package hcl

import (
	"errors"

	"github.com/hashicorp/hcl"
	"github.com/pm-esd/config"
)

// Decoder hcl内容解码器
var Decoder config.Decoder = hcl.Unmarshal

// Encoder hcl内容编码器
var Encoder config.Encoder = func(ptr interface{}) (out []byte, err error) {
	err = errors.New("HCL: is not support encode data to HCL")
	return
}

// Driver hcl的实例
var Driver = &hclDriver{config.Hcl}

// hclDriver 格式内容
type hclDriver struct {
	name string
}

// Name
func (d *hclDriver) Name() string {
	return d.name
}

// GetDecoder 解码
func (d *hclDriver) GetDecoder() config.Decoder {
	return Decoder
}

// GetEncoder 编码
func (d *hclDriver) GetEncoder() config.Encoder {
	return Encoder
}
