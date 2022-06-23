package school

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type PaymentInfo struct {
	StudentUserID     string `json:"student_userid"`
	TradeState        int    `json:"trade_state"`
	TradeNO           string `json:"trade_no"`
	PayerParentUserID string `json:"payer_parent_userid"`
}

type ParamsPaymentGet struct {
	PaymentID string `json:"payment_id"`
}

type ResultPaymentGet struct {
	ProjectName   string         `json:"project_name"`
	Amount        int            `json:"amount"`
	PaymentResult []*PaymentInfo `json:"payment_result"`
}

// GetPaymentResult 获取学生付款结果
func GetPaymentResult(paymentID string, result *ResultPaymentGet) wx.Action {
	params := &ParamsPaymentGet{
		PaymentID: paymentID,
	}

	return wx.NewPostAction(urls.CorpSchoolGetPaymentResult,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsTradeGet struct {
	PaymentID string `json:"payment_id"`
	TradeNO   string `json:"trade_no"`
}

type ResultTradeGet struct {
	TransactionID string `json:"transaction_id"`
	PayTime       int64  `json:"pay_time"`
}

// GetTrade 获取订单详情
func GetTrade(paymentID, tradeNO string, result *ResultTradeGet) wx.Action {
	params := &ParamsTradeGet{
		PaymentID: paymentID,
		TradeNO:   tradeNO,
	}

	return wx.NewPostAction(urls.CorpSchoolGetTrade,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
