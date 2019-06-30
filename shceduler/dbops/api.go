package dbops

import "log"

// 将要删除的vid加到待删除数据库中
func AddVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES (?)")
	if err != nil{
		log.Printf("Prepare AddVideoDeletionRecord error: %v\n", err)
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err!=nil{
		log.Printf("Exec AddVideoDeletionRecord error: %v", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}