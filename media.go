package ocpc_media

import (
	"encoding/json"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/gtkit/encry/md5"
)

var restyClient *resty.Client

func New(resty *resty.Client) {
	if resty == nil {
		panic("resty client is nil")
	}
	restyClient = resty
}

// R 返回一个新的请求对象.
func R() *resty.Request {
	return restyClient.R()
}

// VerifySign 验证sign.
func VerifySign(uri, sec string, domain ...string) bool {
	var url string
	urisp := strings.Split(uri, "&sign=")
	if len(urisp) != 2 {
		return false
	}

	if len(domain) > 0 {
		url = domain[0] + urisp[0] + sec
	} else {
		url = urisp[0] + sec
	}
	versign := md5.New(url)
	sign := urisp[1]

	if sign != versign {
		return false
	}
	return true
}

func ReportJSONByte(bj Custom) []byte {
	c := Custom{
		Oaid:            bj.Oaid,
		ConversionType:  bj.ConversionType,
		ConversionCount: bj.ConversionCount,
		ConversionTime:  bj.ConversionTime,
		ContentID:       bj.ContentID,
		CampaignID:      bj.CampaignID,
		Timestamp:       bj.Timestamp,
		UserAgent:       bj.UserAgent,
		Callback:        bj.Callback,
	}
	jsonStr, _ := json.Marshal(c)

	return jsonStr
}
