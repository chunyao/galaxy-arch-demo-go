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

	Success             = "200" // 成功
	Fail                = "500" // 成功
	Unauthorized        = "401" // token校验错误
	InternalServerError = "500" // 服务器内部错误

	UserUnAuthIdentityErr = "-20001" // 该用户未实名认证
	AgeRefusePayErr       = "-20002" // 年龄在[0,8) 不提供充值服务

)

var resultMsg = map[string]string{
	Success: "成功",
	Fail:    "失败",
}

// GetResultMsg 获取错误描述
func GetResultMsg(code string) string {
	return resultMsg[code]
}
