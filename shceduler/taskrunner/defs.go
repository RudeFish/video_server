package taskrunner

// 定义controlChan中的消息体
const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "./videos/"
)

// 控制部分
type controlChan chan string

// 数据channel，由于数据类型不确定使用interface
type dataChan chan interface{}

// 生产消费函数
type fn func(dc dataChan) error
