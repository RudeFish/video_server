package dbops

import (
	"log"
	"database/sql"
	"imooc/Go语言实战流媒体视频网站/video_server/api/defs"
	"imooc/Go语言实战流媒体视频网站/video_server/api/utils"
	"time"
)


// 用户部分
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil{
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil{
		log.Printf("获取用户失败：%s\n", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	// 当QueryRow(loginName)无法查询到值时会抛出sql.ErrNoRows错误。
	if err !=nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmrDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil{
		log.Printf("删除用户失败：%s\n", err)
		return err
	}
	_,err = stmrDel.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	defer stmrDel.Close()
	return nil
}


// 视频部分
// 添加视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// 创建一个 uuid
	vid, err := utils.NewUUID()
	if err != nil{
		return nil, err
	}

	t := time.Now()
	// 视频id
	ctime := t.Format("Jan 02 2006, 15:04:05") // format里面的时间不能改
	stmrAdd, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES (?,?,?,?)")
	if err != nil {
		log.Printf("添加视频失败：%s\n", err.Error())
		return nil ,err
	}
	_, err = stmrAdd.Exec(vid, aid, name, ctime)
	if err != nil{
		log.Printf("添加视频失败：%s\n", err.Error())
		return nil, err
	}

	res := &defs.VideoInfo{vid, aid, name, ctime}
	defer stmrAdd.Close()
	return res, nil
}

// 获取视频
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmrOut, err := dbConn.Prepare("SELECT author_id , name, display_ctime  FROM video_info WHERE ID = ?")
	//if err != nil {
	//	log.Printf("获取视频失败：%s\n", err.Error())
	//	return nil ,err
	//}
	var aid int
	var name string
	var ctime string

	err = stmrOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		//log.Printf("获取视频失败：%s\n", err.Error())
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmrOut.Close()
	res := &defs.VideoInfo{vid, aid, name, ctime}
	return res, nil
}

// 删除视频
func DeleteVideoInfo(vid string)  error {
	stmrOut, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("获取视频失败：%s\n", err.Error())
		return err
	}

	_, err = stmrOut.Exec(vid)
	if err != nil{
		return err
	}
	defer stmrOut.Close()
	return nil
}


// 评论部分
func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil{
		return err
	}
	stmtIns, err :=  dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) value (?, ?, ?, ?)")
	if err != nil {
		log.Printf("添加评论失败：%s\n", err)
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		log.Printf("添加评论失败：%s\n", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 参数视频id，起始和结束时间
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	var res []*defs.Comment
	var ComId, name, content string

	stmtOut, err := dbConn.Prepare(`select  c.id, u.login_name, c.content
									from comments c inner join users u
									ON u.id = c.author_id 
									where c.video_id = ? and c.time > FROM_UNIXTIME(?) and c.time <= FROM_UNIXTIME(?)`)
	// 此处返回不是一条，需要迭代出来
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil{
		log.Printf("获取评论失败：%s\n", err)
		return nil, err
	}
	for rows.Next() {
		if err = rows.Scan(&ComId, &name, &content); err != nil{
			return nil, err
		}
		c := defs.Comment{Id: ComId, VideoId: vid, AuthorName: name, Content: content}
		res = append(res, &c)
	}
	defer stmtOut.Close()
	return res, err
}