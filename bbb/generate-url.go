package bbb

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/google/go-querystring/query"
)

func (api *Api) GenerateUrl(method string, params interface{}) string {
	v, _ := query.Values(params)
	q := v.Encode()
	v.Set("checksum", api.sign(method, q))

	return api.base + method + "?" + v.Encode()
}

func (api *Api) sign(method string, q string) string {
	toByte := []byte(method + q + api.salt)
	checkSumString := sha1.Sum(toByte)

	return hex.EncodeToString(checkSumString[:])
}
