/*
@Time : 2021/8/16 5:09 下午
@Author : 21
@File : consts
@Software: GoLand
*/
package urls

const  (
	BaseUrl = "https://mp.weixin.qq.com"
)

const (
	// 获取令牌
	ComponentApiComponentTokenUrl = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	// 获取预授权码
	ComponentApiCreatePreAuthCode = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode"
	// 使用授权码获取授权信息
	ComponentApiQueryAuthUrl = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth"
	// 获取授权方的帐号基本信息
	ComponentApiGetAuthorizerInfoUrl = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info"
	// 获取/刷新接口调用令牌
	ComponentApiGetAuthorizerTokenUrl = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token"
)

const   (
	//关联小程序
	WxopenWxamplinkUrl = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplink"
	//获取公众号关联的小程序
	WxopenWxamplinkGetUrl = "https://api.weixin.qq.com/cgi-bin/wxopen/wxamplinkget"
)

const  (
	// 图文消息内的图片获取URL
	OaMediaUpload = "https://api.weixin.qq.com/cgi-bin/media/upload"
	// 图文永久素材
	OaAddMaterial = "https://api.weixin.qq.com/cgi-bin/media/add_material"

)