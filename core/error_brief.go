package core

import "time"

// Brief 异常简要信息类
type Brief struct {
	// Success 是否成功
	Success bool `json:"success"`
	// HTTP状态码
	Status int32 `json:"-"`
	// 状态码
	Code int32 `json:"code"`
	// Msg 异常信息
	Msg string `json:"msg"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// GmtCreate 时间字符串
	GmtCreate string `json:"gmtCreate"`
}

func (b *Brief) GetStatus() int32 {
	return b.Status
}

func (b *Brief) SetStatus(status int32) {
	b.Status = status
}

func (b *Brief) GetCode() int32 {
	return b.Code
}

func (b *Brief) SetCode(code int32) {
	b.Code = code
}

func (b *Brief) GetMsg() string {
	return b.Msg
}

func (b *Brief) SetMsg(msg string) {
	b.Msg = msg
}

// Verify 检验是否成功
func (b *Brief) Verify() {
	b.Success = (b.Status == 0 || b.Status == 200) && (b.Code == 0 || b.Code == 200)
}

// _setTime 设置时间
func (b *Brief) _setTime() {
	current := time.Now()
	b.GmtCreate = current.Format("2006-01-02 15:04:00.000")
	b.Timestamp = current.UnixMilli()
}
