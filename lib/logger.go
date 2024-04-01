package lib

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/shenghui0779/sdk-go/lib/curl"
)

// ReqLog 请求日志
type ReqLog struct {
	data map[string]string
}

// Set 设置日志K-V
func (l *ReqLog) Set(k, v string) {
	l.data[k] = v
}

// SetReqHeader 设置请求头
func (l *ReqLog) SetReqHeader(h http.Header) {
	l.data["request_header"] = HeaderEncode(h)
}

// SetBody 设置请求Body
func (l *ReqLog) SetReqBody(v string) {
	l.data["request_body"] = v
}

// SetRespHeader 设置返回头
func (l *ReqLog) SetRespHeader(h http.Header) {
	l.data["response_header"] = HeaderEncode(h)
}

// SetResp 设置返回报文
func (l *ReqLog) SetRespBody(v string) {
	l.data["response_body"] = v
}

// SetStatusCode 设置HTTP状态码
func (l *ReqLog) SetStatusCode(code int) {
	l.data["status_code"] = strconv.Itoa(code)
}

// Do 日志记录
func (l *ReqLog) Do(ctx context.Context, log func(ctx context.Context, data map[string]string)) {
	if log == nil {
		return
	}

	log(ctx, l.data)
}

// NewReqLog 生成请求日志
func NewReqLog(method, reqURL string) *ReqLog {
	return &ReqLog{
		data: map[string]string{
			"method": method,
			"url":    reqURL,
		},
	}
}

func HeaderEncode(h http.Header) string {
	var buf strings.Builder

	for k, vals := range h {
		for _, v := range vals {
			if buf.Len() > 0 {
				buf.WriteString("&")
			}

			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(v)
		}
	}

	return buf.String()
}

func HeaderToHttpOption(h http.Header) []curl.Option {
	options := make([]curl.Option, 0, len(h))
	for k, vals := range h {
		options = append(options, curl.WithHeader(k, vals...))
	}

	return options
}
