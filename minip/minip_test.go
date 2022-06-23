package minip

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/chenghonour/gochat/mock"
	"github.com/chenghonour/gochat/wx"
)

func TestAccount(t *testing.T) {
	mp := New("wx1def0e9e5891b338", "192006250b4c09247ec02edce69f6a2d")

	assert.Equal(t, "wx1def0e9e5891b338", mp.AppID())
	assert.Equal(t, "192006250b4c09247ec02edce69f6a2d", mp.AppSecret())
}

func TestCode2Session(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"openid": "OPENID",
	"session_key": "SESSION_KEY",
	"unionid": "UNIONID",
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=APPSECRET&js_code=JSCODE&grant_type=authorization_code", nil).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	authSession, err := mp.Code2Session(context.TODO(), "JSCODE")

	assert.Nil(t, err)
	assert.Equal(t, &AuthSession{
		SessionKey: "SESSION_KEY",
		OpenID:     "OPENID",
		UnionID:    "UNIONID",
	}, authSession)
}

func TestAccessToken(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(bytes.NewReader([]byte(`{
	"access_token": "ACCESS_TOKEN",
	"expires_in": 7200,
	"errcode": 0,
	"errmsg": "ok"
}`))),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockHTTPClient(ctrl)

	client.EXPECT().Do(gomock.AssignableToTypeOf(context.TODO()), http.MethodGet, "https://api.weixin.qq.com/cgi-bin/token?appid=APPID&secret=APPSECRET&grant_type=client_credential", nil).Return(resp, nil)

	mp := New("APPID", "APPSECRET")
	mp.SetClient(wx.WithHTTPClient(client))

	accessToken, err := mp.AccessToken(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, &AccessToken{
		Token:     "ACCESS_TOKEN",
		ExpiresIn: 7200,
	}, accessToken)
}

func TestVerifyEventSign(t *testing.T) {
	mp := New("APPID", "APPSECRET")
	mp.SetServerConfig("2faf43d6343a802b6073aae5b3f2f109", "jxAko083VoJ3lcPXJWzcGJ0M1tFVLgdD6qAq57GJY1U")

	assert.True(t, mp.VerifyEventSign("ffb882ae55647757d3b807ff0e9b6098dfc2bc57", "1606902086", "1246833592"))
}

func TestDecryptEventMessage(t *testing.T) {
	mp := New("wx1def0e9e5891b338", "APPSECRET")
	mp.SetServerConfig("2faf43d6343a802b6073aae5b3f2f109", "jxAko083VoJ3lcPXJWzcGJ0M1tFVLgdD6qAq57GJY1U")

	msg, err := mp.DecryptEventMessage("GmSmP2C7QlatlbnrXhJHweW5JsW2F1Fr/xmoMBIJNGnZcN/1PoOySJOJNYEC9ttFhaqDrkznaMkDs7s9u7/eOpvqqRn144EBkLdBLxcNbjLRoF4lD3zBGqjPUS9k/U0x/lET35SkYi+ZwRvuSJSzVEfaRmixYep+JmzIYf5k2qT8113wg2tI68+3gUaKZQqq5W/jC7tbWjWX67XgzMW2JdQOs9VnTjJJO292PWkNZxbhzudrvj2Up8NdJbmaDw93Jz/Kcf7qRfdh5h0GFtOoVh7M4bVwTJf94iZU4ZDx1r8/xDxDINRWGJou4Er72cDBCVBK1TUrtwdmb8eWNJ1gSvw53LckULci98+peaSnTFYuaNhgRQqpVQ+CqVjT0+ASRdyMmDomRyUmhBqSsdrGae9pRfP+Dq4tiRoub87T0gGkFTxAXbUZ0ZPxme67ddreWKFCN/V5ypCynDbjkgpIgfPAFpk017ShXc30RRq4qPvPvN/6XUi1HVXSJq8AkgSQ")

	assert.Nil(t, err)
	assert.Equal(t, wx.WXML{
		"ToUserName":   "gh_3ad31c0ba9b5",
		"FromUserName": "oB4tA6ANthOfuQ5XSlkdPsWOVUsY",
		"CreateTime":   "1606902602",
		"MsgType":      "text",
		"MsgId":        "10086",
		"Content":      "ILoveGochat",
		"URL":          "http://182.92.100.180/webhook",
	}, msg)
}
