package tools

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

type ParamsScheduleAttendee struct {
	UserID string `json:"userid"`
}

type ParamsScheduleReminders struct {
	IsRemind              int   `json:"is_remind,omitempty"`
	RemindBeforeEventSecs int   `json:"remind_before_event_secs,omitempty"`
	IsRepeat              int   `json:"is_repeat,omitempty"`
	RepeatType            int   `json:"repeat_type,omitempty"`
	RepeatUntil           int64 `json:"repeat_until,omitempty"`
	IsCustomRepeat        int   `json:"is_custom_repeat,omitempty"`
	RepeatInterval        int   `json:"repeat_interval,omitempty"`
	RepeatDayOfWeek       []int `json:"repeat_day_of_week,omitempty"`
	RepeatDayOfMonth      []int `json:"repeat_day_of_month,omitempty"`
	Timezone              int   `json:"timezone,omitempty"`
}

type ParamsScheduleAdd struct {
	Schedule *ScheduleAddData `json:"schedule"`
	AgentID  int64            `json:"agentid,omitempty"`
}

type ScheduleAddData struct {
	Organizer   string                    `json:"organizer"`
	StartTime   int64                     `json:"start_time"`
	EndTime     int64                     `json:"end_time"`
	Attendees   []*ParamsScheduleAttendee `json:"attendees,omitempty"`
	Summary     string                    `json:"summary,omitempty"`
	Description string                    `json:"description,omitempty"`
	Reminders   *ParamsScheduleReminders  `json:"reminders,omitempty"`
	Location    string                    `json:"location,omitempty"`
	CalID       string                    `json:"cal_id,omitempty"`
}

type ResultScheduleAdd struct {
	ScheduleID string `json:"schedule_id"`
}

// AddSchedule 创建日程
func AddSchedule(params *ParamsScheduleAdd, result *ResultScheduleAdd) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsScheduleUpdate struct {
	Schedule *ScheduleUpdateData `json:"schedule"`
}

type ScheduleUpdateData struct {
	Organizer     string                    `json:"organizer"`
	ScheduleID    string                    `json:"schedule_id"`
	StartTime     int64                     `json:"start_time"`
	EndTime       int64                     `json:"end_time"`
	Attendees     []*ParamsScheduleAttendee `json:"attendees,omitempty"`
	Summary       string                    `json:"summary,omitempty"`
	Description   string                    `json:"description,omitempty"`
	Reminders     *ParamsScheduleReminders  `json:"reminders,omitempty"`
	Location      string                    `json:"location,omitempty"`
	SkipAttendees bool                      `json:"skip_attendees"`
}

// UpdateSchedule 更新日程
func UpdateSchedule(params *ParamsScheduleUpdate) wx.Action {
	return wx.NewPostAction(urls.CorpToolsScheduleUpdate,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type Schedule struct {
	ScheduleID  string              `json:"schedule_id"`
	Organizer   string              `json:"organizer"`
	Attendees   []*ScheduleAttendee `json:"attendees"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	Reminders   *ScheduleReminders  `json:"reminders"`
	Location    string              `json:"location"`
	CalID       string              `json:"cal_id"`
	StartTime   int64               `json:"start_time"`
	EndTime     int64               `json:"end_time"`
	Status      int                 `json:"status"`
}

type ScheduleAttendee struct {
	UserID         string `json:"userid"`
	ResponseStatus int    `json:"response_status,omitempty"`
}

type ScheduleReminders struct {
	IsRemind              int                    `json:"is_remind"`
	IsRepeat              int                    `json:"is_repeat"`
	RemindBeforeEventSecs int                    `json:"remind_before_event_secs"`
	RemindTimeDiffs       []int                  `json:"remind_time_diffs"`
	RepeatType            int                    `json:"repeat_type"`
	RepeatUntil           int64                  `json:"repeat_until"`
	IsCustomRepeat        int                    `json:"is_custom_repeat"`
	RepeatInterval        int                    `json:"repeat_interval"`
	RepeatDayOfWeek       []int                  `json:"repeat_day_of_week"`
	RepeatDayOfMonth      []int                  `json:"repeat_day_of_month"`
	Timezone              int                    `json:"timezone,omitempty"`
	ExcludeTimeList       []*ScheduleExcludeTime `json:"exclude_time_list"`
}

type ScheduleExcludeTime struct {
	StartTime int64 `json:"start_time"`
}

type ParamsScheduleGet struct {
	ScheduleIDList []string `json:"schedule_id_list"`
}

type ResultScheduleGet struct {
	ScheduleList []*Schedule `json:"schedule_list"`
}

// GetSchedule 获取日程详情
func GetSchedule(scheduleIDs []string, result *ResultScheduleGet) wx.Action {
	params := &ParamsScheduleGet{
		ScheduleIDList: scheduleIDs,
	}

	return wx.NewPostAction(urls.CorpToolsScheduleGet,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsScheduleDelete struct {
	ScheduleID string `json:"schedule_id"`
}

// DeleteSchedule 取消日程
func DeleteSchedule(scheduleID string) wx.Action {
	params := &ParamsScheduleDelete{
		ScheduleID: scheduleID,
	}

	return wx.NewPostAction(urls.CorpToolsScheduleDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsScheduleGetByCalendar struct {
	CalID  string `json:"cal_id"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type ResultScheduleGetByCalendar struct {
	ScheduleList []*Schedule `json:"schedule_list"`
}

// GetScheduleByCalendar 获取日历下的日程列表
func GetScheduleByCalendar(calID string, offset, limit int, result *ResultScheduleGetByCalendar) wx.Action {
	params := &ParamsScheduleGetByCalendar{
		CalID:  calID,
		Offset: offset,
		Limit:  limit,
	}

	return wx.NewPostAction(urls.CorpToolsScheduleGetByCalendar,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsScheduleAttendeeOpt struct {
	ScheduleID string                    `json:"schedule_id"`
	Attendees  []*ParamsScheduleAttendee `json:"attendees"`
}

// AddScheduleAttendee 新增日程参与者
func AddScheduleAttendee(scheduleID string, userIDs ...string) wx.Action {
	params := &ParamsScheduleAttendeeOpt{
		ScheduleID: scheduleID,
		Attendees:  make([]*ParamsScheduleAttendee, 0, len(userIDs)),
	}

	for _, v := range userIDs {
		params.Attendees = append(params.Attendees, &ParamsScheduleAttendee{
			UserID: v,
		})
	}

	return wx.NewPostAction(urls.CorpToolsScheduleAttendeeAdd,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

// DeleteScheduleAttendee 删除日程参与者
func DeleteScheduleAttendee(scheduleID string, userIDs ...string) wx.Action {
	params := &ParamsScheduleAttendeeOpt{
		ScheduleID: scheduleID,
		Attendees:  make([]*ParamsScheduleAttendee, 0, len(userIDs)),
	}

	for _, v := range userIDs {
		params.Attendees = append(params.Attendees, &ParamsScheduleAttendee{
			UserID: v,
		})
	}

	return wx.NewPostAction(urls.CorpToolsScheduleAttendeeDelete,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
