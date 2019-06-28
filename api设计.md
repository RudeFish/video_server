# API部分

## api部分

![1560890345506](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560890345506.png)



![1560892006059](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560892006059.png)

![1560892076287](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560892076287.png)

![1560892255944](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560892255944.png)

## session

![1561384761170](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1561384761170.png)

1、服务启动从DB拉取session到cache

2、用户登录产生session

3、判断session是否过期



这里session用sync.map保存

流程



### api前端部分

main->middleware -> defs(message, err)->handlers->dbops->response



# stream部分

## limiter（流控机制

> 1、新建一个ConnLimiter使用chan来控制连接数量
>
> ```go
> // 当bucker满的情况下
> if len(cl.bucket) >= cl.concurrentConn {
>    log.Printf("Reached the rate limitation.")
>    return false
> }
> ```
>
> 2、在主函数中添加一个NewMiddleWareHandle，把limiter校验放到函数中。
>
> ```go
> type midelWareHandler struct {
>    r *httprouter.Router
>    l *ConnLimiter
> }
> 
> 
> func NewMiddleWareHandle(r *httprouter.Router, cc int) http.Handler {
>    m := midelWareHandler{}
>    m.r = r
>    m.l = NewConnLimiter(cc)
>    return m
> }
> 
> func (m  midelWareHandler)ServeHTTP (w http.ResponseWriter, r *http.Request)  {
>    // 判断如果超过流控值
>    if !m.l.GetConn() {
>       sendErrorResponse(w, http.StatusTooManyRequests, "Too Many Requests")
>       return
>    }
> 
>    m.r.ServeHTTP(w, r)
>    // 释放token
>    defer m.l.ReleaseConn()
> }
> ```

midelWareHandler变成http.Handler需先实现ServeHTTP，这里将判断放到ServeHTTP函数中。

## Handler

返回的头文件中加入视频强制格式

> StreamHandler (读取文件产生到页面

```go
// 加入header视频文件强制提醒
w.Header().Set("Content-type", "video/mp4")
// 传输二进制流 播放视频
	http.ServeContent(w, r, "", time.Now(), video)
```

> uploadHandler（页面上传文件到服务端

```go
// 限定上传文件的大小
r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAN_SIZE)
if err := r.ParseMultipartForm(MAX_UPLOAN_SIZE); err != nil{
   log.Printf("Error when try upload file: %v\n", err)
   sendErrorResponse(w, http.StatusBadRequest, "File is too big!")
   return
}
```

> testPageHandler(文件上传页面

