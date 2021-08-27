/*
@Time : 2021/8/16 5:05 下午
@Author : 21
@File : authorize
@Software: GoLand
*/
package oplatform

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// 获取令牌 https://api.weixin.qq.com/cgi-bin/component/api_component_token
type ComponentAccessToken struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
	// 返回值
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

// 获取预授权码 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
type PreAuthCode struct {
	ComponentAppid       string `json:"component_appid"`
	ComponentAccessToken string `json:"component_access_token"`
	PreAuthCode          string `json:"pre_auth_code"`
	ExpiresIn            int    `json:"expires_in"`
}

// 构建移动端授权链接 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/Authorization_Process_Technical_Description.html
type SafeBindComponent struct {
	ComponentAppid string `json:"component_appid"`
	PreAuthCode    string `json:"pre_auth_code"`
	RedirectUri    string `json:"redirect_uri"`
	AuthType       int    `json:"auth_type"`
	BizAppId       string `json:"biz_app_id"`
}

// 使用授权码获取授权信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/authorization_info.html#%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
type ComponentApiQueryAuth struct {
	ComponentAccessToken string             `json:"component_access_token"`
	ComponentAppid       string             `json:"component_appid"`
	AuthorizationCode    string             `json:"authorization_code"`
	AuthorizationInfo    *AuthorizationInfo `json:"authorization_info"`
}

// TODO 授权之后的用户 信息 取消 func_info 暂时没时间补充
type AuthorizationInfo struct {
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int64  `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

// 获取授权方的帐号基本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html#%E8%AF%B7%E6%B1%82%E5%9C%B0%E5%9D%80
type ComponentApiGetAuthorizerInfo struct {
	ComponentAccessToken string             `json:"component_access_token"`
	ComponentAppid       string             `json:"component_appid"`
	AuthorizerAppid      string             `json:"authorizer_appid"`
	AuthorizerInfo       *AuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo    *AuthorizationInfo `json:"authorization_info"`
}

// TODO AuthorizerInfo 信息不完全
type AuthorizerInfo struct {
	NickName        string           `json:"nick_name"`
	HeadImg         string           `json:"head_img"`
	UserName        string           `json:"user_name"`
	PrincipalName   string           `json:"principal_name"`
	Alias           string           `json:"alias"`
	QrcodeUrl       string           `json:"qrcode_url"`
	ServiceTypeInfo *ServiceTypeInfo `json:"service_type_info"`
	VerifyTypeInfo  *VerifyTypeInfo  `json:"verify_type_info"`
}

type ServiceTypeInfo struct {
	Id int64 `json:"id"`
}

type VerifyTypeInfo struct {
	Id int64 `json:"id"`
}

// 获取刷新接口调用令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_authorizer_token.html
type ComponentApiAuthorizerToken struct {
	ComponentAccessToken   string `json:"component_access_token"`
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	ResponseInfo           *ApiAuthorizerTokenResp
}

type ApiAuthorizerTokenResp struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

// 获取令牌
func GetApiComponentToken(data *ComponentAccessToken) wx.Action {
	return wx.NewPostAction(urls.ComponentApiComponentTokenUrl,
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ComponentAccessToken = gjson.GetBytes(resp, "component_access_token").String()
			data.ExpiresIn = int(gjson.GetBytes(resp, "expires_in").Int())
			return nil
		}),
	)
}

// 获取预授权码
func GetPreAuthCode(data *PreAuthCode) wx.Action {
	return wx.NewPostAction(urls.ComponentApiCreatePreAuthCode,
		wx.WithQuery("component_access_token", data.ComponentAccessToken),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.PreAuthCode = gjson.GetBytes(resp, "pre_auth_code").String()
			data.ExpiresIn = int(gjson.GetBytes(resp, "expires_in").Int())
			return nil
		}),
	)
}

// 授权码获取授权信息
func GetComponentApiQueryAuth(data *ComponentApiQueryAuth) wx.Action {
	return wx.NewPostAction(urls.ComponentApiQueryAuthUrl,
		wx.WithQuery("component_access_token",data.ComponentAccessToken),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.AuthorizationInfo = &AuthorizationInfo{}
			data.AuthorizationInfo.AuthorizerAppid = gjson.GetBytes(resp, "authorization_info.authorizer_appid").String()
			data.AuthorizationInfo.AuthorizerAccessToken = gjson.GetBytes(resp, "authorization_info.authorizer_access_token").String()
			data.AuthorizationInfo.ExpiresIn = gjson.GetBytes(resp, "authorization_info.authorizer_access_token").Int()
			data.AuthorizationInfo.AuthorizerRefreshToken = gjson.GetBytes(resp, "authorization_info.authorizer_refresh_token").String()
			return nil
		}),
	)
}

// 获取授权方的帐号基本信息
func GetComponentApiGetAuthorizerInfo(data *ComponentApiGetAuthorizerInfo) wx.Action {
	return wx.NewPostAction(urls.ComponentApiGetAuthorizerInfoUrl,
		wx.WithQuery("component_access_token", data.ComponentAccessToken),

		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.AuthorizerInfo = &AuthorizerInfo{}
			data.AuthorizationInfo = &AuthorizationInfo{}

			err := json.Unmarshal(resp, &data)
			//data.AuthorizerInfo.ServiceTypeInfo = &ServiceTypeInfo{Id: gjson.GetBytes(resp, "authorizer_info.service_type_info.id").Int()}
			//data.AuthorizerInfo.VerifyTypeInfo = &VerifyTypeInfo{Id: gjson.GetBytes(resp, "authorizer_info.verify_type_info.id").Int()}
			//data.AuthorizerInfo.NickName = gjson.GetBytes(resp, "authorizer_info.nick_name").String()
			//data.AuthorizerInfo.HeadImg = gjson.GetBytes(resp, "authorizer_info.head_img").String()
			//data.AuthorizerInfo.UserName = gjson.GetBytes(resp, "authorizer_info.user_name").String()
			//data.AuthorizerInfo.PrincipalName = gjson.GetBytes(resp, "authorizer_info.principal_name").String()
			//data.AuthorizerInfo.Alias = gjson.GetBytes(resp, "authorizer_info.alias").String()
			//data.AuthorizerInfo.QrcodeUrl = gjson.GetBytes(resp, "authorizer_info.qrcode_url").String()
			//
			//data.AuthorizationInfo.AuthorizerAppid = gjson.GetBytes(resp, "authorization_info.authorizer_appid").String()
			//data.AuthorizationInfo.AuthorizerAccessToken = gjson.GetBytes(resp, "authorization_info.authorizer_access_token").String()
			//data.AuthorizationInfo.ExpiresIn = gjson.GetBytes(resp, "authorization_info.authorizer_access_token").Int()
			//data.AuthorizationInfo.AuthorizerRefreshToken = gjson.GetBytes(resp, "authorization_info.authorizer_refresh_token").String()
			return err
		}),
	)
}

// 获取/刷新接口调用令牌
func GetComponentApiAuthorizertoken(data *ComponentApiAuthorizerToken) wx.Action {
	return wx.NewPostAction(urls.ComponentApiGetAuthorizerTokenUrl,
		wx.WithQuery("component_access_token", data.ComponentAccessToken),
		wx.WithBody(func() ([]byte, error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {
			data.ResponseInfo = &ApiAuthorizerTokenResp{}
			data.ResponseInfo.AuthorizerRefreshToken = gjson.GetBytes(resp, "authorizer_refresh_token").String()
			data.ResponseInfo.AuthorizerAccessToken = gjson.GetBytes(resp, "authorizer_access_token").String()
			data.ResponseInfo.ExpiresIn = int(gjson.GetBytes(resp, "expires_in").Int())
			return nil
		}),
	)
}
