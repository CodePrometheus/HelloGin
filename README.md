## 总结Gin框架实战

### demo内容
路由解析和路由分组，表单数据实体绑定，模板文件与模板语言的使用，以及企业项目开发的需求分析，功能开发和RESTful接口的开发，mysql数据库的操作和使用，表操作，session的使用，文件上传和下载

#### 框架下载和搭建
可以通过go get命令进行gin框架源码的下载：

```go
go  get -u github.com/gin-gonic/gin
```

在$GOPATH变量下的src/github.com/gin-gonic目录中，存放了gin框架的源码。

#### Gin构建HTTP服务
Gin框架有两种方式创建渲染引擎，分别是gin.Default()和gin.New()。
* gin.New()：使用该方法可以创建一个gin引擎。
* gin.Default()：使用gin.New()方法创建一个gin引擎，同时创建gin引擎时，默认会创建和使用Logger和Recovery中间件组件功能。

通常情况下推荐使用gin.Default()方式进行创建。

构建的gin引擎，可以处理HTTP的请求，处理方式与go语言原生的http服务解析方式类似，通过默认的路由组件接收并处理http请求。
```go
func (group *RouterGroup) Handle(httpMethod, relativePath string,handlers ...HandlerFunc) IRouters{
    ...
}
```

上述的Handle方法是处理所有HTTP请求的通用方法，具体到GET、POST等类型的请求，可以使用诸如:

```go
app := gin.Default()
app.GET()
或者
app.POST()
或者
app.DELETE()
等
```
在路由组中封装了对应到具体HTTP请求类型的方法，方便HTTP接口解析处理。

#### 请求与返回数据格式
gin框架支持多种数据格式如：Form表单字段，File，JSON，XML等多种数据格式。

* ShouldBindQuery：使用该方法绑定go语言结构体对象，用于解析请求参数。
* ShouldBind：同理使用该方法绑定go语言结构体指针对象，用于请求参数的自动解析。
* ShouldBindJson：使用该方法绑定go语言结构体指针，用于json格式的请求参数的解析。

除去以上的常用方法外，还可以使用gin框架提供的其他的参数解析的方法。

请求处理结束后的数据格式返回，在gin中也同样是支持多种数据格式的返回。
```go
//byte切片数据
context.Writer.Write([]byte("hello world"))
//string字符串
context.Writer.WriteString("Hello world")
//HTML页面和
context.HTML(http.StatusOK,htmlName,paramter)
//JSON格式
context.JSON(http.StatusOK,map[string]interface{}{...})
```

#### 路由组的定义和使用
在使用gin.Default()创建gin引擎时，使用的是默认的路由组。路由组是router.Group中的一个方法，用于对请求进行分发处理，Group返回的是一个RouterGroup指针对象，而RouterGroup是gin框架中的一个路由组结构体定义。如下所示：
```go
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}
```
RouterGroup实现了IRoutes中定义的方法，诸如GET、POST、DELETE等方法。

#### 中间件的编写和使用
在使用gin进行项目功能开发时，为了更好的梳理系统架构，可以将一些通用业务单独抽离并进行开发，然后以插件化的形式进行对接使用，这种方式称之为中间件开发。gin框架中中间件是一个类型为HandleFunc的函数类型。

gin中使用Use方法设置使用中间件。比如
```go
func CustomMiddler() gin.HandlerFunc {
   ...
}
app := gin.Default()
app.Use(CustomMiddler())
...
```

#### 数据库的配置和访问
使用gin连接和操作mysql数据库，需要下载mysql驱动，并进行包引入。同时，链接数据库需要配置连接信息，包括：url，用户名，密码。

除了可以在代码中配置连接数据库的参数外，还可以将数据库参数配置在配置文件中。在本课程的项目开发中，就采用第二种方式，在config.json文件中进行数据库参数配置。

### 项目功能开发
在学习完了gin的基础功能后，开始进行一个接口项目的功能开发，从零开始编写一个go语言项目。

#### 前端项目的使用
在本课程中开发实现的是一个接口项目，配合接口功能进行调试和测试的前端项目功能使用vue进行开发。在课程中，我们关注的是前端项目的使用和阅读，不侧重于编写和功能实现。

了解和熟悉使用命令进行前端项目的编译和运行，以及接口调试即可。

#### 第三方SDK的使用
在实际的开发中，第三方SDK的使用时非常常见的情况。在接口项目的功能开发过程中，接入的是阿里云的短信SDK发送短信的功能。有些项目在开发中，可能还会接入支付SDK，文件存储SDK等一些第三方的sdk，这些均属于第三方SDK的接入范畴。

#### 分布式文件系统
在接口项目的功能开发过程中，涉及到了文件上传的功能。不同于以前的web框架项目教程，本课程中讲解了如何搭建分布式文件存储系统FastDFS。并讲解了FastDFS的工作原理和代码实现步骤。

FastDFS由Tracker、Server和Client三个组件构成。Tracker是跟踪服务器，负责系统整体的调度，Server是真正的文件存储服务器，Client则负责与Tracker和Server的连接。

#### 跨域访问
在接口项目的功能开发过程中，前端框架项目和后端项目代码分离开发和运行的，在前端项目执行网络请求时，会涉及到跨域访问的问题。我们采取的方案是在接口项目服务端，编写跨域访问的中间件，并设置生效，从而解决跨域访问的问题。跨域访问的中间件编程实现如下所示：
```go
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range context.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}
```

设置跨域访问中间件生效的代码：
```go
...
app := gin.Default()
app.Use(Cors())
...
```

#### xorm框架
接口项目发开中，数据库的操作使用的是xorm框架与gin框架结合实现。xorm框架实现从Go语言结构体自动映射成为数据库表的操作，包括数据库表结构，字段约束等操作。同时，还可以使用go语言结构体表示出表结构实体的关系。

xorm框架使用前,同样需要先下载并引入后使用。

#### Redis缓存集成
在接口项目功能开发过程中，还集成了Redis缓存的功能，通过config.json配置文件，对Redis相关的集成环境进行配置，并通过封装工具方法，实现对Redis缓存的初始化。具体的代码编写已经在项目开发过程中进行了实现。

数据库的操作，redis缓存的操作，fastdfs配置的操作等均是作为全局配置进行封装的，可以实现跨项目的使用和集成。


















