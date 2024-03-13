package ocpc

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/gtkit/logger"
)

func (x *XiaoMi) Report(ctx context.Context, payload []byte) (string, error) {
	var (
		custom Custom
		resp   AqyResp
		err    error
	)
	if err = json.Unmarshal(payload, &custom); err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.WriteString(xmCallbackUrl + "?callback=")
	builder.WriteString(custom.Callback)
	builder.WriteString("&oaid=")
	builder.WriteString(custom.Oaid)
	builder.WriteString("&conv_time=")
	builder.WriteString(custom.Timestamp)
	builder.WriteString("&convType=")
	builder.WriteString(x.eventType(custom.ConversionType))
	query := builder.String()

	logger.Info("【 *小米 】Report query: ", query)

	res, err := R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		Get(query)
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

func (x *XiaoMi) eventType(ct string) string {
	switch ct {
	case "activate":
		return "APP_ACTIVE" // 激活
	case "retain":
		return "APP_RETENTION" // 次日留存
	default:
		return "APP_ACTIVE" // 激活
	}
}
