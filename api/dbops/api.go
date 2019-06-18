package dbops

import "log"

func AddUSerCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil{
		return err
	}

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
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
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmrDel, err := dbConn.Prepare("DELETE FROM uesrs WHERE login_name = ? AND pwd = ?")
	if err != nil{
		log.Printf("删除用户失败：%s\n", err)
		return err
	}
	stmrDel.Exec(loginName, pwd)
	stmrDel.Close()
	return nil
}