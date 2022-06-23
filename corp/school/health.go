package school

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type ParamsHealthReportStat struct {
	Date string `json:"date"`
}

type ResultHealthReportStat struct {
	PV int `json:"pv"`
	UV int `json:"uv"`
}

// GetHealthReportStat 获取健康上报使用统计
func GetHealthReportStat(date string, result *ResultHealthReportStat) wx.Action {
	params := &ParamsHealthReportStat{
		Date: date,
	}

	return wx.NewPostAction(urls.CorpSchoolGetHealthReportStat,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsHealthReportJobIDs struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type ResultHealthReportJobIDs struct {
	Ending int      `json:"ending"`
	JobIDs []string `json:"jobids"`
}

// GetHealthReportJobIDs 获取健康上报任务ID列表
func GetHealthReportJobIDs(offset, limit int, result *ResultHealthReportJobIDs) wx.Action {
	params := &ParamsHealthReportJobIDs{
		Offset: offset,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpSchoolGetHealthReportJobIDs,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthReportJobInfo struct {
	Title             string                    `json:"title"`
	Creator           string                    `json:"creator"`
	Type              int                       `json:"type"`
	ApplyRange        *HealthReportApplyRange   `json:"apply_range"`
	ReportTo          *HealthReportTo           `json:"report_to"`
	ReportType        int                       `json:"report_type"`
	SkipWeekend       int                       `json:"skip_weekend"`
	FinishCnt         int                       `json:"finish_cnt"`
	QuestionTemplates []*HealthQuestionTemplate `json:"question_templates"`
}

type HealthReportApplyRange struct {
	UserIDs  []string `json:"userids"`
	PartyIDs []int64  `json:"partyids"`
}

type HealthReportTo struct {
	UserIDs []string `json:"userids"`
}

type HealthQuestionTemplate struct {
	QuestionID   int64                   `json:"question_id"`
	Title        string                  `json:"title"`
	QuestionType int                     `json:"question_type"`
	IsRequired   int                     `json:"is_required"`
	OptionList   []*HealthQuestionOption `json:"option_list"`
}

type HealthQuestionOption struct {
	OptionID   int64  `json:"option_id"`
	OptionText string `json:"option_text"`
}

type ParamsHealthReportJobInfo struct {
	JobID string `json:"jobid"`
	Date  string `json:"date"`
}

type ResultHealthReportJobInfo struct {
	JobInfo *HealthReportJobInfo `json:"job_info"`
}

// GetHealthReportJobInfo 获取健康上报任务详情
func GetHealthReportJobInfo(jobID, date string, result *ResultHealthReportJobInfo) wx.Action {
	params := &ParamsHealthReportJobInfo{
		JobID: jobID,
		Date:  date,
	}

	return wx.NewPostAction(urls.CorpSchoolGetHealthReportJobInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthReportAnswer struct {
	IDType        int                  `json:"id_type"`
	UserID        string               `json:"userid"`
	StudentUserID string               `json:"student_userid"`
	ParentUserID  string               `json:"parent_userid"`
	ReportTime    int64                `json:"report_time"`
	ReportValues  []*HealthReportValue `json:"report_values"`
}

type HealthReportValue struct {
	QuestionID   int64    `json:"question_id"`
	SingleChoice int      `json:"single_choice"`
	MultiChoice  []int    `json:"multi_choice"`
	Text         string   `json:"text"`
	FileID       []string `json:"fileid"`
}

type ParamsHealthReportAnswer struct {
	JobID  string `json:"jobid"`
	Date   string `json:"date"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultHealthReportAnswer struct {
	Answers []*HealthReportAnswer `json:"answers"`
}

// GetHealthReportAnswer 获取用户填写答案
func GetHealthReportAnswer(jobID, date string, offset, limit int, result *ResultHealthReportAnswer) wx.Action {
	params := &ParamsHealthReportAnswer{
		JobID:  jobID,
		Date:   date,
		Offset: offset,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpSchoolGetHealthReportAnswer,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type CustomizeHealthQuestionTemplate struct {
	QuestionID   int64                   `json:"question_id"`
	Title        string                  `json:"title"`
	QuestionType int                     `json:"question_type"`
	IsMustFill   int                     `json:"is_must_fill"`
	IsNotDisplay int                     `json:"is_not_display"`
	OptionList   []*HealthQuestionOption `json:"option_list"`
}

type CustomizeHealthReportValue struct {
	QuestionID  int64  `json:"question_id"`
	SingleChose int    `json:"single_chose"`
	Text        string `json:"text"`
}

type CustomizeHealthInfo struct {
	UserID             string                        `json:"userid"`
	HealthQRCodeStatus int                           `json:"health_qrcode_status"`
	SelfSubmit         int                           `json:"self_submit"`
	ReportValues       []*CustomizeHealthReportValue `json:"report_values"`
}

type ParamsCustomizeHealthInfo struct {
	Date    string `json:"date"`
	NextKey string `json:"next_key"`
	Limit   int    `json:"limit"`
}

type ResultCustomizeHealthInfo struct {
	HealthInfos       []*CustomizeHealthInfo             `json:"health_infos"`
	QuestionTemplates []*CustomizeHealthQuestionTemplate `json:"question_templates"`
	TemplateID        string                             `json:"template_id"`
	Ending            int                                `json:"ending"`
	NextKey           string                             `json:"next_key"`
}

// GetTeacherCustomizeHealthInfo 获取老师健康信息
func GetTeacherCustomizeHealthInfo(date, nextKey string, limit int, result *ResultCustomizeHealthInfo) wx.Action {
	params := &ParamsCustomizeHealthInfo{
		Date:    date,
		NextKey: nextKey,
		Limit:   limit,
	}

	return wx.NewPostAction(urls.CorpSchoolGetTeacherCustomizeHealthInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetStudentCustomizeHealthInfo 获取学生健康信息
func GetStudentCustomizeHealthInfo(date, nextKey string, limit int, result *ResultCustomizeHealthInfo) wx.Action {
	params := &ParamsCustomizeHealthInfo{
		Date:    date,
		NextKey: nextKey,
		Limit:   limit,
	}

	return wx.NewPostAction(urls.CorpSchoolGetStudentCustomizeHealthInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type HealthQRCode struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	UserID     string `json:"userid"`
	QRCodeData string `json:"qrcode_data"`
}

type ParamsHealthQRCode struct {
	Type    int      `json:"type"`
	UserIDs []string `json:"userids"`
}

type ResultHealthQRCode struct {
	ResultList []*HealthQRCode `json:"result_list"`
}

// GetHealthQRCode 获取师生健康码
func GetHealthQRCode(objType int, userIDs []string, result *ResultHealthQRCode) wx.Action {
	params := &ParamsHealthQRCode{
		Type:    objType,
		UserIDs: userIDs,
	}

	return wx.NewPostAction(urls.CorpSchoolGetHealthQRCode,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
