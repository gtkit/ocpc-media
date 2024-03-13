// @Author 2024/1/18 15:17:00
package ocpc_media

const (
	hwCallbackUrl = "https://ppscrowd-drcn.op.cloud.huawei.com/action-lib-track/hiad/v2/actionupload"
	bdCallbackUrl = "http://ocpc.baidu.com/ocpcapi/cb/actionCb"
	xmCallbackUrl = "http://trail.e.mi.com/api/callback"
)

type ActionBody struct {
	Custom
	Common
}

type Common struct {
	TrackingEnabled string `json:"tracking_enabled"` // '0': 不允许跟踪，'1' 允许跟踪。没有此字段的话则使用空字符串
}

/**
 * 应用转化类别
 * activate: 激活
 * browse: 浏览
 * collection: 收藏
 * addToCart: 加入购物车
 * preOrder: 下单
 * register: 注册
 * retain: 次日留存
 * paid: 付费
 * custom: 自定义
 * form_submit: 表单提交
 * consult: 有效咨询
 * custom_acquisit: 有效获客
 * book: 有效预定
 * custom_ landingpage: 自定义
 */

type Custom struct {
	Oaid            string `json:"oaid" binding:"required"`     // 设备标识符，明文，没有传空字符
	ConversionType  string `json:"conversion_type"`             // 转化事件的类型
	ConversionCount string `json:"conversion_count,omitempty"`  // 转化数量
	ConversionTime  string `json:"conversion_time"`             // 转化事件发生的时间，Unix 时间戳，单位秒
	ConversionPrice string `json:"conversion_price,omitempty"`  // 转化价格，单位：元，保留两位小数
	ContentID       string `json:"content_id,omitempty"`        // 素材id，与该条转化行为匹配的、广告主接收到素材id
	AdvertiserID    string `json:"advertiser_id,omitempty"`     // 广告主id
	CampaignID      string `json:"campaign_id,omitempty"`       // 计划id
	Timestamp       string `json:"timestamp"`                   // 本请求发起的时间戳，Unix时间戳，单位毫秒
	UserAgent       string `json:"user_agent,omitempty"`        // 用户设备的 User-Agent 信息
	Callback        string `json:"callback" binding:"required"` // Callback 字段和AdvertiserID字段必须传一个
}

type HwResp struct {
	ResultCode    int    `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

type BdResp struct {
	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
	Error     string `json:"error,omitempty"`
	Status    int    `json:"status,omitempty"`
}

type AqyResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
