package vo

import (
	"app/src/model"
)

type ReportSubscribeVo struct {
	SubscribeData model.ReportSubscribe `json:"subscribeData"` // 内容
	Status        bool                  `json:"status"`        // 是否订阅
}
