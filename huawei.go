package ocpc

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/gtkit/encry/hmac"
	"github.com/gtkit/logger"
)

const (
	ACTIVATE = "activate"
	REGISTER = "register"
	RETAIN   = "retain"
)

func (hw *HuaWei) Report(ctx context.Context, payload []byte) (string, error) {
	hwBodyJSON := string(payload)
	logger.Info("华为 Report query post-body: ", hwBodyJSON)

	hash256 := Hmac256(hw.SecretKey, hwBodyJSON)

	res, err := R().
		SetContext(ctx).
		SetHeader("Authorization", AuthHeader(hash256)).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetBody(hwBodyJSON).
		Post(hwCallbackUrl)
	if err != nil {
		return "", err
	}
	respStr := res.String()

	var resp HwResp
	if err = json.Unmarshal([]byte(respStr), &resp); err != nil {
		return respStr, err
	}

	if resp.ResultCode != 0 || resp.ResultMessage != "success" {
		return respStr, errors.New(resp.ResultMessage)
	}

	return respStr, nil
}

func HWJSONString(bj Custom) string {
	hwdody := ActionBody{
		Custom: Custom{
			Oaid:            bj.Oaid,
			ConversionType:  bj.ConversionType,
			ConversionCount: bj.ConversionCount,
			ConversionTime:  bj.ConversionTime,
			ContentID:       bj.ContentID,
			CampaignID:      bj.CampaignID,
			Timestamp:       bj.Timestamp,
			UserAgent:       bj.UserAgent,
			Callback:        bj.Callback,
		},
		Common: Common{
			TrackingEnabled: "1",
		},
	}
	jsonStr, _ := json.Marshal(hwdody)

	return string(jsonStr)
}

func HWJSONByte(bj Custom) []byte {
	hwdody := ActionBody{
		Custom: Custom{
			Oaid:            bj.Oaid,
			ConversionType:  bj.ConversionType,
			ConversionCount: bj.ConversionCount,
			ConversionTime:  bj.ConversionTime,
			ContentID:       bj.ContentID,
			CampaignID:      bj.CampaignID,
			Timestamp:       bj.Timestamp,
			UserAgent:       bj.UserAgent,
			Callback:        bj.Callback,
		},
		Common: Common{
			TrackingEnabled: "1",
		},
	}
	jsonStr, _ := json.Marshal(hwdody)

	return jsonStr
}

func AuthHeader(hash string) string {
	milltime := time.Now().UnixMilli()
	timeStr := strconv.FormatInt(milltime, 10)
	return "Digest validTime=\"" + timeStr + "\",response=\"" + hash + "\""
}

func Hmac256(key, body string) string {
	return hmac.Sha256ToHex(key, body)
}

// ConversionType 转换类型.
func ConversionType(conversionType string) string {
	var ct string
	switch conversionType {
	case ACTIVATE, REGISTER: // 激活, 注册
		ct = conversionType
	case "day1retention": // 留存
		ct = RETAIN
	default:
		ct = ACTIVATE
	}
	return ct
}
