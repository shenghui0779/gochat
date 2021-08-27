/*
@Time : 2021/8/17 5:34 下午
@Author : 21
@File : oa
@Software: GoLand
*/
package oplatform

import (
	"encoding/json"
	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// 关联小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Official__Accounts/Mini_Program_Management_Permission.html
type WxopenWxamplink struct {
	AccessToken string `json:"access_token"`
	Appid       string `json:"appid"`
	NotifyUsers string `json:"notify_users"`
	ShowProfile string `json:"show_profile"`
}

// 获取公众号关联的小程序
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Official__Accounts/Mini_Program_Management_Permission.html
type WxopenWxamplinkget struct {
	AccessToken string `json:"access_token"`
	// 列表
	Items []*WxopensItems
}

type WxopensItems struct {
	Status              int    `json:"status"`
	Username            string `json:"username"`
	Appid               string `json:"appid"`
	Source              string `json:"source"`
	Nickname            string `json:"nickname"`
	Selected            int    `json:"selected"`
	NearbyDisplayStatus int    `json:"nearby_display_status"`
	Released            int    `json:"released"`
	HeadimgURL          string `json:"headimg_url"`
	FuncInfos           []struct {
		Status int    `json:"status"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
	} `json:"func_infos"`
	CopyVerifyStatus int    `json:"copy_verify_status"`
	Email            string `json:"email"`
}

func SetWxopenWxamplink(data *WxopenWxamplink) wx.Action {
	return wx.NewPostAction(urls.WxopenWxamplinkUrl,
		wx.WithQuery("access_token", data.AccessToken),
		wx.WithBody(func() (bytes []byte, e error) {
			return json.Marshal(data)
		}),
		wx.WithDecode(func(resp []byte) error {

			return nil
		},
		))
}

func GetWxampLink(data *WxopenWxamplinkget) wx.Action {
	return wx.NewPostAction(urls.WxopenWxamplinkGetUrl,
		wx.WithQuery("access_token", data.AccessToken),
		wx.WithDecode(func(resp []byte) error {
			data.Items = []*WxopensItems{}
			jsonStr := gjson.GetBytes(resp, "wxopens.items").String()
			err := json.Unmarshal([]byte(jsonStr), &data.Items)
			if err != nil {
				return err
			}
			return nil
		}, ),
	)
}
