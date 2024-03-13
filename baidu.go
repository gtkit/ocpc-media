package ocpc_media

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"

	"github.com/gtkit/encry/md5"
	"github.com/gtkit/logger"
)

func (bd *BaiDu) Report(ctx context.Context, payload []byte) (string, error) {
	var (
		bdresp BdResp
		custom Custom
		err    error
	)

	if err = json.Unmarshal(payload, &custom); err != nil {
		return "", err
	}

	query := BDQueryString(custom)

	sign := md5.New(bdCallbackUrl + "?" + query + bd.SecretKey)
	reqQuery := query + "&sign=" + sign

	logger.Info("【 *百度 】Report query: ", bdCallbackUrl+"?"+reqQuery)

	res, err := R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetQueryString(reqQuery).
		Get(bdCallbackUrl)
	if err != nil {
		return "", err
	}
	respStr := res.String()

	if err = json.Unmarshal([]byte(respStr), &bdresp); err != nil {
		return "", err
	}
	if bdresp.ErrorCode != 0 || bdresp.Status > 200 {
		return "", errors.New(bdresp.ErrorMsg + bdresp.Error)
	}

	return respStr, nil
}

func BDQueryString(bj Custom) string {
	urlval := url.Values{}
	urlval.Add("a_type", bj.ConversionType)
	urlval.Add("a_value", "0")
	urlval.Add("actType", "2")
	urlval.Add("join_type", "oaid")
	urlval.Add("oaid", bj.Oaid)
	urlval.Add("ext_info", bj.Callback)

	return urlval.Encode()
}
