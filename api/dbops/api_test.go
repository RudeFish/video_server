package dbops

import (
	"testing"
	"strconv"
	"time"
	"fmt"
)

// 由于test里面没法传参，这里用一个全局变量记录
var tempVid string

// 删除数据->跑参数数据->清除数据
// 清除数据
func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T)  {
	err := AddUserCredential("silas", "123456")
	if err != nil{
		t.Errorf("Error of Adduser: %v\n", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("silas")
	if pwd != "123456" || err != nil{
		t.Errorf("Error of Getuser")
	}
}

func testDeleteUser(t *testing.T)  {
	err := DeleteUser("silas", "123456")
	if err != nil{
		t.Errorf("Error of Deluser: %v\n", err)
	}
}

// 查看del是否正确
func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("silas")
	if err != nil{
		t.Errorf("Error of Regetuser:%v\n", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed\n")
	}
}


// video部分测试
func TestVideoWorkInfo(t *testing.T)  {
	clearTables()
	t.Run("PreparUser", testAddUser) // 由于videoInfo里面有用户id
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDelVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T)  {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil{
		t.Errorf("Error of AddNewVideo: %s\n", err)
	}
	// 获取vid
	tempVid = vi.Id
}

func testGetVideoInfo(t *testing.T)  {
	_, err :=  GetVideoInfo(tempVid)
	if err != nil{
		t.Errorf("Error of GetVideoInfo: %s\n", err)
	}
}

func testDelVideoInfo(t *testing.T)  {
	err :=  DeleteVideoInfo(tempVid)
	if err != nil{
		t.Errorf("Error of DeleteVideoInfo: %s\n", err)
	}
}

func testRegetVideoInfo(t *testing.T)  {
	vi , err :=  GetVideoInfo(tempVid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %s\v", err)
	}
}

// 评论部分测试
func TestComments(t *testing.T)  {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
	//t.Run("RegetListComments", testRegetListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	err := AddNewComments(vid, aid, "这是一条测试评论！")
	if err != nil{
		t.Errorf("Error of AddComments err: %s\n", err)
	}
}

func testListComments(t *testing.T)  {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10) ) // 纳秒转秒除十亿
	comments, err := ListComments(vid, from, to)
	if err != nil{
		t.Errorf("Error of ListComments err: %s\n", err)
	}
	for i, v := range comments {
		fmt.Printf("comment: %d, %v \n", i, v)
	}
}