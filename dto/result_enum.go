package dto

const (
	SignatureBuildErr         = "-10000" // 构建签名失败
	ApiClientGetErr           = "-10001" // 获取APIClient失败
	CallParamsBuildErr        = "-10002" // 获取组装和签名后的请求参数失败
	AccessedOrderServiceErr   = "-10003" // 访问支付下单接口失败
	PrepayIdGetErr            = "-10004" // 获取预支付Id失败
	ThirdPartyResponseErr     = "-10005" // 服务器对第三方支付响应结果验证不通过
	PlatformCertificateGetErr = "-10006" // 获取第三方支付平台证书失败
	SignatureVerifyErr        = "-10007" // 验证第三方签名失败
	JwtUserNotMatchErr        = "-10008" // Jwt校验用户不一致

	Success             = "0"   // 成功
	Unauthorized        = "401" // token校验错误
	InternalServerError = "500" // 服务器内部错误

	UserUnAuthIdentityErr           = "-20001" // 该用户未实名认证
	AgeRefusePayErr                 = "-20002" // 年龄在[0,8) 不提供充值服务
	OrderNotPayErr                  = "-20003" // 订单未支付成功
	OrderDeliveredErr               = "-20004" // 订单已处于发货状态
	ObtainPayStrategyErr            = "-20005" // 获取未成年人支付校验策略失败
	GuestNotPayErr                  = "-20006" // 不支持游客使用支付服务
	OverMaxCharge8And16PerTimeErr   = "-20007" // 年龄在[8，16），单笔充值金额超过充值上限
	OverMaxCharge16And18PerTimeErr  = "-20008" // 年龄在[16，18），单笔充值金额超过充值上限
	OverMaxCharge8And16PerMonthErr  = "-20009" // 年龄在[8，16），单月充值金额超过充值上限
	OverMaxCharge16And18PerMonthErr = "-20010" // 年龄在[16，18），单月充值金额超过充值上限
	ObtainDeliveryStrategyErr       = "-20011" // 获取通知游戏厂商发货策略失败
	NotifyCPErr                     = "-20012" // 通知游戏厂商发放道具请求失败

	QueryAgeServerErr = "-30001" // 查询年龄的服务不可访问
	DingTalkRobotErr  = "-30002" // 钉钉机器人服务不可用
)

var resultMsg = map[string]string{
	Success:                         "成功",
	SignatureBuildErr:               "构建签名失败",
	ApiClientGetErr:                 "获取APIClient失败",
	CallParamsBuildErr:              "获取组装和签名后的请求参数失败",
	AccessedOrderServiceErr:         "访问支付下单接口失败",
	PrepayIdGetErr:                  "获取预支付Id失败",
	ThirdPartyResponseErr:           "服务器对第三方支付响应结果验证不通过",
	PlatformCertificateGetErr:       "获取第三方支付平台证书失败",
	SignatureVerifyErr:              "验证第三方签名失败",
	InternalServerError:             "服务器内部异常",
	Unauthorized:                    "Token校验失败",
	JwtUserNotMatchErr:              "Token校验用户不一致",
	UserUnAuthIdentityErr:           "该用户未实名认证",
	AgeRefusePayErr:                 "年龄在0~8岁,不提供充值服务",
	ObtainPayStrategyErr:            "获取未成年人支付校验策略失败",
	QueryAgeServerErr:               "查询年龄的服务不可访问",
	GuestNotPayErr:                  "不支持游客使用支付服务",
	OverMaxCharge8And16PerTimeErr:   "年龄在[8，16），单笔充值金额超过充值上限",
	OverMaxCharge16And18PerTimeErr:  "年龄在[16，18），单笔充值金额超过充值上限",
	OverMaxCharge8And16PerMonthErr:  "年龄在[8，16），单月充值金额超过充值上限",
	OverMaxCharge16And18PerMonthErr: "年龄在[16，18），单月充值金额超过充值上限",
	OrderNotPayErr:                  "订单未支付成功",
	OrderDeliveredErr:               "订单已处于发货状态",
	ObtainDeliveryStrategyErr:       "获取通知游戏厂商发货策略失败",
	NotifyCPErr:                     "通知游戏厂商发放道具请求失败",
	DingTalkRobotErr:                "钉钉机器人服务不可用",
}

// GetResultMsg 获取错误描述
func GetResultMsg(code string) string {
	return resultMsg[code]
}
