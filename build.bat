goto web_end
cd .\web
go build
copy web.exe D:\work\imooc.com\video_server\video_server_web_ui\web.exe
cd ..
xcopy .\template D:\work\imooc.com\video_server\video_server_web_ui\template /E/I/D/Y
:web_end

:goto api_end
cd .\api
go build
copy api.exe D:\work\imooc.com\video_server\video_server_api\api.exe
cd ..
:api_end
