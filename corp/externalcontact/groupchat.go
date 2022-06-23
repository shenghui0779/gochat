package externalcontact

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type GroupChatOwnerFilter struct {
	UserIDList []string `json:"userid_list"`
}

type ParamsGroupChatList struct {
	StatusFilter int                   `json:"status_filter,omitempty"`
	OwnerFilter  *GroupChatOwnerFilter `json:"owner_filter,omitempty"`
	Cursor       string                `json:"cursor,omitempty"`
	Limit        int                   `json:"limit"`
}

type GroupChatListData struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

type ResultGroupChatList struct {
	GroupChatList []*GroupChatListData `json:"group_chat_list"`
	NextCursor    string               `json:"next_cursor"`
}

// ListGroupChat 获取客户群列表
func ListGroupChat(params *ParamsGroupChatList, result *ResultGroupChatList) wx.Action {
	return wx.NewPostAction(urls.CorpExternalContactGroupChatList,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsGroupChatGet struct {
	ChatID   string `json:"chat_id"`
	NeedName int    `json:"need_name,omitempty"`
}

type GroupChat struct {
	ChatID     string             `json:"chat_id"`
	Name       string             `json:"name"`
	Owner      string             `json:"owner"`
	CreateTime int64              `json:"create_time"`
	Notice     string             `json:"notice"`
	MemberList []*GroupChatMember `json:"member_list"`
	AdminList  []*GroupChatAdmin  `json:"admin_list"`
}

type GroupChatAdmin struct {
	UserID string `json:"userid"`
}

type GroupChatInvitor struct {
	UserID string `json:"userid"`
}

type GroupChatMember struct {
	UserID        string            `json:"userid"`
	Type          int64             `json:"type"`
	Unionid       string            `json:"unionid"`
	JoinTime      int64             `json:"join_time"`
	JoinScene     int64             `json:"join_scene"`
	Invitor       *GroupChatInvitor `json:"invitor"`
	GroupNickname string            `json:"group_nickname"`
	Name          string            `json:"name"`
}

type ResultGroupChatGet struct {
	GroupChat *GroupChat `json:"group_chat"`
}

// GetGroupChat 获取客户群详情
func GetGroupChat(chatID string, needName int, result *ResultGroupChatGet) wx.Action {
	params := &ParamsGroupChatGet{
		ChatID:   chatID,
		NeedName: needName,
	}

	return wx.NewPostAction(urls.CorpExternalContactGroupChatGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsOpenGIDToChatID struct {
	OpenGID string `json:"opengid"`
}

type ResultOpenGIDToChatID struct {
	ChatID string `json:"chat_id"`
}

// OpenGIDToChatID 客户群opengid转换
func OpenGIDToChatID(opengid string, result *ResultOpenGIDToChatID) wx.Action {
	params := &ParamsOpenGIDToChatID{
		OpenGID: opengid,
	}

	return wx.NewPostAction(urls.CorpExternalContactOpenGIDToChatID,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
