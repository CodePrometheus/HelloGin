## Gin框架介绍及环境搭建

## 背景
在之前的课程中，我们已经学习了常用的beego框架、iris框架。在实际开发中，框架是不可或缺的可以提高开发效率开发者效率的工具。

大家如果没有学习过的同学，可以去学习另外两门go web的课程，课程链接：

* 《Beego框架项目实战》：[https://ke.qq.com/course/411509?tuin=335a96b2](https://ke.qq.com/course/411509?tuin=335a96b2)

* 《Iris框架项目开发实战》：[https://ke.qq.com/course/407365?tuin=335a96b2](https://ke.qq.com/course/407365?tuin=335a96b2)

从本节课开始，我们来学习go语言web应用开发中另外一个比较常用，也比较知名的框架：Gin框架。

## Gin简介
Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，已经发布了1.0版本。具有快速灵活，容错方便等特点。其实对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错。框架更像是一些常用函数或者工具的集合。借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范。

Gin框架是开源的，可以在github上下载其源码库，查看相应的说明。Gin源码库地址：[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

Gin框架有一个官方网站，有对Gin相关的介绍和学习资料。官方网站：[https://gin-gonic.com/](https://gin-gonic.com/)

## Gin特点和特性
* 速度：Gin之所以被很多企业和团队所采用，第一个原因是因为其速度快，性能表现出众。

* 中间件：和iris类型，gin在处理请求时，支持中间件操作，方便编码处理。

* 路由：在gin中可以非常简单的实现路由解析的功能，并包含路由组解析功能。

* 内置渲染：Gin支持JSON、XML和HTML等多种数据格式的渲染，并提供了方便的操作API。

## 学习文档
Gin官方为我们开发者提供了参考文档。学习文档地址如下：[https://gin-gonic.com/zh-cn/docs/](https://gin-gonic.com/zh-cn/docs/)

## Gin开发环境搭建
### 环境要求
gin框架需要go语言版本在1.6及以上。可以通过go version查看自己的go语言版本是否符合要求。

### 安装gin框架库
通过go get命令安装gin框架：
```go
go get -u github.com/gin-gonic/gin
```
安装完毕后，可以在当前系统的$GOPATH目录下的src/github.com目录中找到gin-gonic目录，该目录下存放的就是gin框架的源码。

安装完毕后，我们可以使用gin来写一个简单的Hello world程序。使用一下gin。
## Hello Wrold
### 编码
新创建一个项目GinHelloWrold,并创建main.go程序,内容如下:
```go
func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

### 运行
执行运行命令：
```go
go run main.go
```

## 指定端口号运行
```go
server := gin.Default()
...
server.Run(":8090")
```



