# api

![1560890345506](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560890345506.png)



![1560892006059](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560892006059.png)

![1560892076287](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560892076287.png)

![1560892255944](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1560892255944.png)

# session

![1561384761170](C:\Users\llw98\AppData\Roaming\Typora\typora-user-images\1561384761170.png)

1、服务启动从DB拉取session到cache

2、用户登录产生session

3、判断session是否过期



这里session用sync.map保存

流程



## api前端部分

main->middleware -> defs(message, err)->handlers->dbops->response

