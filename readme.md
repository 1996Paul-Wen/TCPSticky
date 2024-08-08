本项目演示了所谓的tcp"粘包"问题

server.go启动一个tcp listener，client.go启动一个client连接tcp server，在建立的tcp连接中loop发送字符串"Hello! How are you?"

演示命令：`go run main.go server.go client.go`

tcp server收到的消息可能如下
```plaintext
read 19 bytes: Hello! How are you?
read 38 bytes: Hello! How are you?Hello! How are you?
read 171 bytes: Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?
read 19 bytes: Hello! How are you?
read 114 bytes: Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?Hello! How are you?
read 19 bytes: Hello! How are you?
```
可以看到多段消息可能会产生粘连

既然 TCP 协议是基于字节流的，这其实就意味着应用层协议要自己划分消息的边界

如果我们能在应用层协议中定义消息的边界，那么无论 TCP 协议如何对应用层协议的数据包进程拆分和重组，接收方都能根据协议的规则恢复对应的消息。在应用层协议中，最常见的两种解决方案就是基于长度或者基于终结符（Delimiter）:

client和server需要约定好协议，通过加入长度header或者加入Delimiter的方式，定义具体的encoding和decoding实现，从而在应用层正确分离消息单元

refer to https://draveness.me/whys-the-design-tcp-message-frame/