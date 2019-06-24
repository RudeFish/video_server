package dbops

import (
	"strconv"
	"imooc/Go语言实战流媒体视频网站/video_server/api/defs"
	"sync"
	"log"
)

func InserSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUSE (?, ?, ?, ?)")
	if err != nil{
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss :=  &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT login_name, TTL FROM sessions WHERE session_id = ?")
	if err != nil{
		return nil, err
	}
	var ttl, uname string
	stmtOut.QueryRow(sid).Scan(&uname, &ttl)
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil{
		ss.TTL = res
		ss.Username = uname
	}else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	m :=  &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil{
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		var id, ttlstr, uname string
		if rows.Scan(&id, &ttlstr, &uname); err != nil{
			log.Printf("retrive session err: %s\n", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{Username: uname, TTL: ttl}
			m.Store(m, ss)
			log.Printf("session id: %s, ttl: %s\n", id, ss.TTL)
		}
	}

	defer stmtOut.Close()
	return m, nil
}

func DelteSession(sid string) error {
	stmrOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("删除session失败：%s\n", err.Error())
		return err
	}

	_, err = stmrOut.Exec(sid)
	if err != nil{
		return err
	}
	defer stmrOut.Close()
	return nil
}