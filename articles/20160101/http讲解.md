# HTTP学习（一）

## HTTP简介


### HTTP报文结构
![Alt text](./http1.png)
HTTP报文分为以下三个部分：

* 起始行
* 首部
* 消息主体

其中首部包含Content-type和Content-length，分别指定了主题的数据格式和数据长度。主体部分包含由任意数据组成的数据块。并不是所有的报文都包含实体部分的，有时，报文只是以一个CRLF结束。
```http
<method>
	<request-url>
		<version>
			<headers>
				<entity-body></entity-body>
			</headers>
		</version>
	</request-url>
</method>
```
### 报文的语法
所有的HTTP报文都可以分为两类：
- 请求报文（request message）
- 响应报文（response message）
### HTTP请求方法
- get            请求获取request-uri所标识的资源
- put            请求服务器存储一个资源，并用request-uri作为其标识
- post          在Request-URI所标识的资源后附加新的数据
- delete       请求服务器删除Request-URI所标识的资源
- options     请求查询服务器的性能，或者查询与资源相关的选项和需求
- trace         请求服务器回送收到的请求信息，主要用于测试或诊断
- connect    保留将来使用
### 请求URL
命令了所请求的资源

### 状态码
使用三位数字表示请求过程中所发生的情况。每个状态码的第一位数字都用来描述状态码的一般类别（成功，出错）。
常见的状态码有：

- 404
- 505
- 200



