# go-soap
Soap server is built by go.and soap client request the post.Achieving the soap one-way communication.Refer only.

# use echo
        e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.HideBanner = true
		e.POST("/go-soap/sever",soapServerHandler)

# server start client to return request with you want to post
use defer client.Start()
##
eg:
        //todo 根据需要，可以启动soap的客户端
    	//todo 在服务端完成HTTP通信，返回请求后，可以启动soap客户端，完成soap的单向通信
    	{
    		defer client.Start()
    	}

# pkg为根据wsdl描述文件生成的soap通信代码
use it,you can use little time to establish the soap frame.

Finally.Refer to my code,you should rebuild project according to needing.
Happy enjoy it.






