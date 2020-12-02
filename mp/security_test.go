package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestImageSecCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ImageSecCheck("test.jpg"))

	assert.Nil(t, err)
}

func TestMediaCheckAsync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"trace_id": "967e945cd8a3e458f3c74dcb886068e9"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MediaCheckAsyncResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", MediaCheckAsync(SecCheckMediaImage, "test.jpg", dest))

	assert.Nil(t, err)
	assert.Equal(t, "967e945cd8a3e458f3c74dcb886068e9", dest.TraceID)
}

func TestMsgSecCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", MsgSecCheck("hello world!"))

	assert.Nil(t, err)
}
