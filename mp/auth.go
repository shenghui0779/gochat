package mp

import (
	"net/url"

	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// AuthSession 小程序授权Session
type AuthSession struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
}

// AccessToken 小程序access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthInfo 小程序授权信息
type AuthInfo interface {
	AppID() string
}

// UserInfo 用户信息
type UserInfo struct {
	OpenID    string    `json:"openId"`
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	AvatarURL string    `json:"avatarUrl"`
	NickName  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	UnionID   string    `json:"unionId"`
	WaterMark WaterMark `json:"watermark"`
}

func (u *UserInfo) AppID() string {
	return u.WaterMark.AppID
}

// PhoneInfo 用户手机号绑定信息
type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	WaterMark       WaterMark `json:"watermark"`
}

func (p *PhoneInfo) AppID() string {
	return p.WaterMark.AppID
}

// WaterMark 水印
type WaterMark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

// PaidUnionID 支付用户unionid
type PaidUnionID struct {
	UnionID string
}

// GetPaidUnionIDByTransactionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func GetPaidUnionIDByTransactionID(dest *PaidUnionID, openid, transactionID string) wx.Action {
	query := url.Values{}

	query.Set("openid", openid)
	query.Set("transaction_id", transactionID)

	return wx.NewOpenGetAPI(PaidUnionURL, query, func(resp []byte) error {
		dest.UnionID = gjson.GetBytes(resp, "unionid").String()

		return nil
	})
}

// GetPaidUnionIDByOutTradeNO 用户支付完成后，获取该用户的 UnionId，无需用户授权
func GetPaidUnionIDByOutTradeNO(dest *PaidUnionID, openid, mchid, outTradeNO string) wx.Action {
	query := url.Values{}

	query.Set("openid", openid)
	query.Set("mch_id", mchid)
	query.Set("out_trade_no", outTradeNO)

	return wx.NewOpenGetAPI(PaidUnionURL, query, func(resp []byte) error {
		dest.UnionID = gjson.GetBytes(resp, "unionid").String()

		return nil
	})
}