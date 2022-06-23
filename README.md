# gochat
# Forked from shenghui0779/gochat, thinks

[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org) [![GitHub release](https://img.shields.io/github/release/chenghonour/gochat.svg)](https://github.com/chenghonour/gochat/releases/latest) [![pkg.go.dev](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/chenghonour/gochat) [![Apache 2.0 license](http://img.shields.io/badge/license-Apache%202.0-brightgreen.svg)](http://opensource.org/licenses/apache2.0)

📦 微信 Go SDK

| 模块             | 功能                                                                                          |
| --------------- | --------------------------------------------------------------------------------------------- |
| 支付 &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;\|&nbsp; mch      | 下单 \| 支付 \| 退款 \| 查询 \| 委托代扣 \| 红包 \| 企业付款 \| 账单 \| 评价数据 \| 验签 \| 解密       |
| 公众号 &nbsp;\|&nbsp; offia  | 授权 \| 用户 \| 消息 \| 素材 \| 菜单 \| 发布能力 \| 草稿箱 \| 客服 \| 二维码 \| OCR \| 回复 \| 事件处理 |
| 小程序 &nbsp;\|&nbsp; minip  | 授权 \| 解密 \| 二维码 \| 消息 \| 客服 \| 素材 \| 插件 \| URL Scheme \| URL Link \| OCR \| 事件处理  |
| 企业微信 \| corp | 支持几乎所有服务端API                                                                             |

## 获取

```sh
go get -u github.com/chenghonour/gochat
```

## 使用须知

- 微信API被封装成 `Action` 接口（授权 和 AccessToken 等部分API除外）
- 每个API对应一个 `Action`，统一由 `Do` 方法执行
- 除支付（mch）外，返回结果均以 `Result` 为前缀的结构体指针接收
- 对于微信支付的回调通知处理，提供了两个方法：
  - 验签 - `VerifyWXMLResult`
  - 解密 - `DecryptWithAES256ECB` (退款)
- 对于微信推送的事件消息处理，提供了三个方法：
  - 验签 - `VerifyEventSign`
  - 解密 - `DecryptEventMessage`
  - 回复 - `Reply`
- 企业微信按照不同功能模块划分了相应的目录，根据URL可以找到对应的目录和文件
- 所有API均采用Mock单元测试（Mock数据来源于官方文档，如遇问题，欢迎提[Issue](https://github.com/chenghonour/gochat/issues)）

> - 执行单元测试时，有些不能通过（比如：因时间戳导致等），需要看代码注释说明
> - 执行 `mch` 单元测试时，还需要使用用于单元测试的 `FormatMap2XML`

## 支付

> - 版本：V2
> - 模式：普通商户直连模式

```go
import (
    "github.com/chenghonour/gochat"
    "github.com/chenghonour/gochat/wx"
    "github.com/chenghonour/gochat/mch"
)

// 创建实例
pay := gochat.NewMch("mchid", "apikey", tls.Certificate...)

// 设置 debug 模式（支持自定义日志）
pay.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 统一下单 -------------------------------
action := mch.UnifyOrder("appid", &mch.ParamsUnifyOrder{...})
result, err := pay.Do(ctx, action)

if err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 拉起支付 -------------------------------

// APP支付
pay.APPAPI("appid", "prepayID")

// JSAPI/小程序支付
pay.JSAPI("appid", "prepayID")

// 小程序红包
pay.MinipRedpackJSAPI("appid", "package")
```

## 公众号

```go
import (
    "github.com/chenghonour/gochat"
    "github.com/chenghonour/gochat/wx"
    "github.com/chenghonour/gochat/offia"
)

// 创建实例
oa := gochat.NewOffia("appid", "appsecret")

// 设置服务器配置
oa.SetServerConfig("token", "encodingAESKey")

// 设置 debug 模式（支持自定义日志）
oa.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 生成网页授权URL -------------------------------

url := oa.OAuth2URL(offia.ScopeSnsapiBase, "redirectURL", "state")

fmt.Println(url)

// --------- 获取网页授权Token -------------------------------

result, err := oa.Code2OAuthToken(ctx, "code")

if err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 获取AccessToken -------------------------------

result, err := oa.AccessToken(ctx)

if err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 获取关注的用户列表 -------------------------------

result := new(offia.ResultUserList)
action := offia.GetUserList("nextOpenID", result)

if err := oa.Do(ctx, action); err != nil {
    log.Println(err)

    return
}

fmt.Println(result)
```

## 小程序

```go
import (
    "github.com/chenghonour/gochat"
    "github.com/chenghonour/gochat/wx"
    "github.com/chenghonour/gochat/minip"
)

// 创建实例
mp := gochat.NewMinip("appid", "appsecret")

// 设置服务器配置
mp.SetServerConfig("token", "encodingAESKey")

// 设置 debug 模式（支持自定义日志）
mp.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 获取授权SessionKey -------------------------------

result, err := mp.Code2Session(ctx, "code")

if err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 获取AccessToken -------------------------------

result, err := mp.AccessToken(ctx)

if err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 解密授权的用户信息 -------------------------------

result := new(minip.UserInfo)

if err := mp.DecryptAuthInfo("sessionKey", "iv", "encryptedData", result); err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 获取用户手机号 -------------------------------

result := new(minip.ResultPhoneNumber)

if err := minip.GetPhoneNumber("code", result); err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 创建小程序二维码 -------------------------------

qrcode := new(minip.QRCode)
action := minip.CreateQRCode("pagepath", 120, qrcode)

if err := minip.Do(ctx, action); err != nil {
    log.Println(err)

    return
}

fmt.Println(base64.StdEncoding.EncodeToString(qrcode.Buffer))
```

## 企业微信

```go
import (
    "github.com/chenghonour/gochat"
    "github.com/chenghonour/gochat/wx"
    "github.com/chenghonour/gochat/corp"
    "github.com/chenghonour/gochat/corp/addrbook"
)

// 创建实例
cp := gochat.NewCorp("corpid")

// 设置服务器配置
cp.SetServerConfig("token", "encodingAESKey")

// 设置 debug 模式（支持自定义日志）
cp.SetClient(wx.WithDedug(), wx.WithLogger(wx.Logger))

// --------- 生成网页授权URL -------------------------------

url := cp.OAuth2URL(corp.ScopeSnsapiBase, "redirectURL", "state")

fmt.Println(url)

// --------- 生成扫码授权URL -------------------------------

url := cp.QRCodeAuthURL("agentID", "redirectURL", "state")

fmt.Println(url)

// --------- 获取AccessToken -------------------------------

result, err := cp.AccessToken(ctx, "secret")

if err != nil {
    log.Println(err)

    return
}

fmt.Println(result)

// --------- 获取部门列表 -------------------------------

result := new(addrbook.ResultDepartmentList)
action := addrbook.ListDepartment(0, result)

if err := cp.Do(ctx, action); err != nil {
    log.Println(err)

    return
}

fmt.Println(result)
```

## 说明

- [API Reference](https://pkg.go.dev/github.com/chenghonour/gochat)
- 注意：因 `access_token` 每日获取次数有限且含有效期，故服务端应妥善保存 `access_token` 并定时刷新
- 配合 [yiigo](https://github.com/shenghui0779/yiigo) 使用，可以更方便的操作 `MySQL`、`MongoDB` 与 `Redis` 等

**Enjoy 😊**
