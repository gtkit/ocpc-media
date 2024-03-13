package ocpc

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/gtkit/encry/md5"
	"github.com/gtkit/logger"
)

func (a *AiQiYi) Report(ctx context.Context, payload []byte) (string, error) {
	var (
		custom Custom
		resp   AqyResp
		cbkurl string
		err    error
	)
	if err = json.Unmarshal(payload, &custom); err != nil {
		return "", err
	}
	if cbkurl, err = url.QueryUnescape(custom.Callback); err != nil {
		return "", err
	}

	sign := md5.New(cbkurl + a.SecretKey)
	reqQuery := cbkurl + "&event_type=" + a.eventType(custom.ConversionType) + "&sign=" + sign

	logger.Info("【 *爱奇艺 】Report query: ", reqQuery)

	res, err := R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		Get(reqQuery)
	if err != nil {
		return "", err
	}
	respStr := res.String()

	if err = json.Unmarshal([]byte(respStr), &resp); err != nil {
		return "", err
	}
	if resp.Status > 200 || strings.Contains(resp.Message, "event_type") {
		return "", errors.New("status code is: " + strconv.Itoa(resp.Status) + ", message is: " + resp.Message)
	}

	return respStr, nil
}

func (a *AiQiYi) eventType(ct string) string {
	switch ct {
	case "activate":
		return "0" // 激活
	case "retain":
		return "3" // 次日留存
	default:
		return "0" // 日活
	}
}
