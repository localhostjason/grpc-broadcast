# grpc-broadcast
grpc 广播实时通信



### Server端

1.  支持 tcp+grpc 一起共用一个端口
2.  有gin+gorm 一系列 web 开发 组件



### Client端

1.  主要通过grpc 跟server 端通信。数据流 ：C==>S 单向传输
2.  服务器推送消息至客户端，让客户端主动去做一些事件，需要一条双向数据流
3. 服务器挂了，客户端会定时尝试去连接。



### 运行

````shell
#server
go run main.go -x

#client
go run main.go -x
````

