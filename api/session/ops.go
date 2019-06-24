package session

import (
	"sync"
	"imooc/Go语言实战流媒体视频网站/video_server/api/dbops"
	"imooc/Go语言实战流媒体视频网站/video_server/api/defs"
	"imooc/Go语言实战流媒体视频网站/video_server/api/utils"
	"time"
)

var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

// 获取当前时间
func nowInMilli() int64 {
	return time.Now().UnixNano()/100000
}

// 删除过期session
func deleteExpiredSession(sid string) {
	// 将map和DB中的session删除
	sessionMap.Delete(sid)
	dbops.DelteSession(sid)
}

// 从数据库加载session到全局map中
func LoadSessionFromDB()  {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

// 创建一个session
func GenerateNewSessionId(username string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli() // 精确到毫秒
	ttl := ct + 30 * 60 * 1000 // serverside session vaild time: 30min

	// 将session加到map中
	ss := &defs.SimpleSession{Username: username, TTL: ttl}
	sessionMap.Store(id, ss)
	// 将session加到DB中
	dbops.InserSession(id, ttl, username)

	return id
}

// 判断session是否过期
func IsSessionExpired(sid string) (string, bool) {
	if ss, ok := sessionMap.Load(sid); ok {
		// 如果已经过期
		if ss.(*defs.SimpleSession).TTL < nowInMilli() {
			// delete expired session；
			deleteExpiredSession(sid)
			return "", true
		}
		// 未过期返回username
		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
