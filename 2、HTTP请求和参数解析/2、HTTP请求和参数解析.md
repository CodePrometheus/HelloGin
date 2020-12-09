## gin网络请求与路由处理

第一节课我们介绍了Gin框架，并做了Gin框架的安装，完成了第一个Gin工程的创建。从本节课开始，我们继续深入学习Gin框架。

### 创建Engine
在gin框架中，Engine被定义成为一个结构体，Engine代表gin框架的一个结构体定义，其中包含了路由组、中间件、页面渲染接口、框架配置设置等相关内容。默认的Engine可以通过gin.Default进行创建，或者使用gin.New()同样可以创建。两种方式如下所示：

```go
engine1 = gin.Default()
engine2 = gin.New()
```

gin.Default()和gin.New()的区别在于gin.Default也使用gin.New()创建engine实例，但是会默认使用Logger和Recovery中间件。

Logger是负责进行打印并输出日志的中间件，方便开发者进行程序调试；Recovery中间件的作用是如果程序执行过程中遇到panic中断了服务，则Recovery会恢复程序执行，并返回服务器500内部错误。通常情况下，我们使用默认的gin.Default创建Engine实例。

### 处理HTTP请求
在创建的engine实例中，包含很多方法可以直接处理不同类型的HTTP请求。

#### HTTP请求类型
http协议中一共定义了八种方法或者称之为类型来表明对请求网络资源（Request-URI）的不同的操作方式，分别是：OPTIONS、HEAD、GET、POST、PUT、DELETE、TRACE、CONNECT。

虽然一共有八种请求操作类型，但是实际开发中常用的就：GET、POST、DELETE等几种。

#### 通用处理
engine中可以直接进行HTTP请求的处理，在engine中使用Handle方法进行http请求的处理。Handle方法包含三个参数，具体如下所示：

```go
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) IRoutes
```

* httpMethod：第一个参数表示要处理的HTTP的请求类型，是GET、POST、DELETE等8种请求类型中的一种。

* relativePath：第二个参数表示要解析的接口，由开发者进行定义。

* handlers：第三个参数是处理对应的请求的代码的定义。

举例如下：

##### Handle处理GET请求
```go
...
engine.Handle("GET", "/hello", func(context *gin.Context) {
	//获取请求接口
	fmt.Println(context.FullPath())
	//获取字符串参数
	name := context.DefaultQuery("name", "")
	fmt.Println(name)

	//输出
	context.Writer.Write([]byte("Hello ," + name))
})
...
```

通过Handle方法第一个参数指定处理GET类型的请求，解析的接口是/hello。

Context是gin框架中封装的一个结构体，这是gin框架中最重要，最基础的一个结构体对象。该结构体可以提供我们操作请求，处理请求，获取数据等相关的操作，通常称之为上下文对象，简单说为我们提供操作环境。

可以通过context.Query和context.DefaultQuery获取GET请求携带的参数。

可以通过context.Writer.Write向请求发起端返回数据。

##### Handle处理POST请求
```go
...
engine.Handle("POST", "/login", func(context *gin.Context) {

	fmt.Println(context.FullPath())
	//userName
	username := context.PostForm("username")
	fmt.Println(username)

	//passWord
	password := context.PostForm("pwd")
	fmt.Println(password)

	context.Writer.Write([]byte("User login"))
})
...
```

如上的案例，通过第一个参数指定了解析POST类型的请求，第二个参数指定解析的接口为/login。

POST请求是以form表单的方式提交数据的，可以通过context.PostForm获取表单中提交的数据字段。

其他类型的HTTP请求也可以通过Handle方法处理对应类型的请求。

#### 分类处理
除了engine中包含的通用的处理方法以外，engine还可以按类型进行直接解析。engine中包含有get方法、post方法、delete方法等与http请求类型对应的方法。

##### engine.GET()处理GET请求
engine中包含GET方法处理HTTP的GET类型的请求。engine的GET方法包含两个参数，编程使用如下所示：
```go
...
engine.GET("/hello", func(context *gin.Context) {
	fmt.Println(context.FullPath())

	username := context.Query("name")
	fmt.Println(username)

	context.Writer.Write([]byte("Hello," + username))
})
...
```

**context.DefaultQuery：** 除了context.DefaultQuery方法获取请求携带的参数数据以外，还可以使用context.Query方法来获取Get请求携带的参数。

##### engine.POST()处理POST请求
```go
...
engine.POST("/login", func(context *gin.Context) {
	
	fmt.Println(context.FullPath())
	username, exist := context.GetPostForm("username")
	if exist {
		fmt.Println(username)
	}

	password, exists := context.GetPostForm("pwd")
	if exists {
		fmt.Println(password)
	}
	
	context.Writer.Write([]byte("Hello , " + username))
})
...
```

**context.GetPostForm获取表单数据：**POST请求以表单的形式提交数据,除了可以使用context.PostForm获取表单数据意外，还可以使用context.GetPostForm来获取表单数据。

##### engine.DELETE()处理DELETE请求
在项目开发中，通常都是遵循RESTful标准进行接口开发。除了GET、POST以外，还会有DELETE等操作。

比如要执行某个删除操作，会发送DELETE类型的请求，同时需要携带一些操作的参数。比如要删除用户，按照RESTful标准会进行如下所示：
```go
...
engine.DELETE("/user/:id", DeleteHandle)
func DeleteHandle(context *gin.Context) {
	fmt.Println(context.FullPath())

	userID := context.Param("id")

	fmt.Println(userID)

	context.Writer.Write([]byte("Delete user's id : " + userID))
}
...
```

**context.Param获取请求参数**  

客户端的请求接口是DELETE类型，请求url为：[http://localhost:9000/user/1](http://localhost:9000/user/1)。最后的1是要删除的用户的id，是一个变量。因此在服务端gin中，通过路由的:id来定义一个要删除用户的id变量值，同时使用context.Param进行获取。

### RouterGroup
之所以engine中包含通用型的Handle和分类处理的GET、POST等类型的方法，是因为Engine中有RouterGroup作为匿名字段。

RouteGroup可以称之为路由集合，在gin中定义为结构体：
```go
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine
	root     bool
}
```

RouteGroup的作用就是为每一个服务请求提供解析功能，并指定每一个请求对应的处理程序。




















