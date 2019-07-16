// Package json use the https://github.com/json-iterator/go for parse json
package json

import (
	"github.com/pm-esd/config"

	jsoniter "github.com/json-iterator/go"
)

var parser = jsoniter.ConfigCompatibleWithStandardLibrary

// Decoder for json
var Decoder config.Decoder = func(data []byte, v interface{}) (err error) {
	s := config.StripJSONComments(string(data))
	return parser.Unmarshal([]byte(s), v)
}

// Encoder for json
var Encoder config.Encoder = parser.Marshal

// Driver for json
var Driver = &jsonDriver{config.JSON}

// jsonDriver for json format content
type jsonDriver struct {
	name string
}

// Name get name
func (d *jsonDriver) Name() string {
	return d.name
}

// GetDecoder for json
func (d *jsonDriver) GetDecoder() config.Decoder {
	return Decoder
}

// GetEncoder for json
func (d *jsonDriver) GetEncoder() config.Encoder {
	return Encoder
}
