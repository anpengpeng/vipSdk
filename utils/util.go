package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func CreateRequestSign(access_token, appKey, format, language, method, service, timestamp, version, busiParams, appSecret string) string {
	signStr := ""
	if access_token != "" {
		signStr += "accessToken" + access_token
	}
	signStr += "appKey" + appKey + "format" + format
	if language != "" {
		signStr += "language" + language
	}
	signStr += "method" + method + "service" + service + "timestamp" + timestamp + "version" + version + busiParams

	return Hmac(signStr, appSecret)
}

func GetQueryString(maps map[string]string) string {
	queryString := ""
	i := 0
	for k, v := range maps {
		if i == 0 {
			queryString += k + "=" + v
		} else {
			queryString += "&" + k + "=" + v
		}
		i++
	}

	return queryString
}

func Hmac(data, key string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))

	return strings.ToUpper(hex.EncodeToString(hmac.Sum([]byte(""))))
}
