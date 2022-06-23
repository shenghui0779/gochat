package school

import (
	"encoding/json"
	"strconv"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type Parent struct {
	ParentUserID   string   `json:"parent_userid"`
	Relation       string   `json:"relation"`
	Mobile         string   `json:"mobile"`
	IsSubscribe    int      `json:"is_subscribe"`
	ExternalUserID string   `json:"external_userid"`
	Children       []*Child `json:"children"`
}

type Child struct {
	StudentUserID string `json:"student_userid"`
	Relation      string `json:"relation"`
	Name          string `json:"name"`
}

type Student struct {
	StudentUserID string    `json:"student_userid"`
	Name          string    `json:"name"`
	Department    []int64   `json:"department"`
	Parents       []*Parent `json:"parents"`
}

type ResultUserGet struct {
	UserType int      `json:"user_type"`
	Student  *Student `json:"student"`
	Parent   *Parent  `json:"parent"`
}

// GetUser 读取学生或家长
func GetUser(userID string, result *ResultUserGet) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolUserGet,
		wx.WithQuery("userid", userID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ResultUserList struct {
	Students []*Student `json:"students"`
}

// ListUser 获取部门成员详情
func ListUser(departmentID int64, fetchChild int, result *ResultUserList) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolUserList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithQuery("fetch_child", strconv.Itoa(fetchChild)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsArchSyncModeSet struct {
	ArchSyncMode int `json:"arch_sync_mode"`
}

// SetArchSyncMode 设置家校通讯录自动同步模式
func SetArchSyncMode(mode int) wx.Action {
	params := &ParamsArchSyncModeSet{
		ArchSyncMode: mode,
	}

	return wx.NewPostAction(urls.CorpSchoolSetArchSyncMode,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ResultParentList struct {
	Parents []*Parent `json:"parents"`
}

// ListParent 获取部门家长详情
func ListParent(departmentID int64, result *ResultParentList) wx.Action {
	return wx.NewGetAction(urls.CorpSchoolParentList,
		wx.WithQuery("department_id", strconv.FormatInt(departmentID, 10)),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
