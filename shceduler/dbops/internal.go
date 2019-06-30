package dbops

import "log"

// 读取要删除的数据,批量拿取
func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	var ids []string
	if err != nil{
		log.Printf("Prepare ReadVideoDeletionRecord error: %v\n", err)
		return nil, err
	}

	rows, err := stmtOut.Query(count)
	if err != nil{
		log.Printf("Query ReadVideoDeletionRecord error: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil{
			log.Printf("Scan ReadVideoDeletionRecord error: %v\n", err)
			return nil, err
		}
		log.Printf("get del video id: %s\n", id)
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}


// 删除
func DelVideoDeltionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id = ?")
	if err != nil{
		log.Printf("Prepare DelVideoDeltionRecord error: %v\n", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil{
		log.Printf("Exec DelVideoDeltionRecord error: %v\n", err)
		return err
	}

	defer stmtDel.Close()
	return nil
}