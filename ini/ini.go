package ini

import (
	"github.com/abulo/go-esd/config"
	"github.com/abulo/go-esd/config/ini/parser"
)

// Decoder ini内容解码器
var Decoder config.Decoder = parser.Decode

// Encoder ini内容编码器
var Encoder config.Encoder = func(ptr interface{}) (out []byte, err error) {
	return parser.Encode(ptr)
}

// Driver ini 的实例
var Driver = &iniDriver{config.Ini}

// iniDriver 格式内容
type iniDriver struct {
	name string
}

// Name
func (d *iniDriver) Name() string {
	return d.name
}

// GetDecoder 解码
func (d *iniDriver) GetDecoder() config.Decoder {
	return Decoder
}

// GetEncoder 编码
func (d *iniDriver) GetEncoder() config.Encoder {
	return Encoder
}
