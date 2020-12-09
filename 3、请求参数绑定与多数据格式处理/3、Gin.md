## 请求参数绑定与多数据格式处理

上节课我们学习了使用Gin框架的Engine的默认路由功能解析HTTP请求。本节课我们来学习gin框架的参数绑定操作和请求结果返回格式。

### 内容回顾
在上节课处理POST请求时，使用context.PostForm或者context.DefaultPostForm获取客户端表单提交的数据。
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

像上述这种只有username和password两个字段的表单数据进行提交时，可以使用context.PostForm和context.GetPostForm获取。但是如果表单数据较多时，使用PostForm和GetPostForm一次获取一个表单数据，开发效率较慢。

Gin框架提供给开发者表单实体绑定的功能，可以将表单数据与结构体绑定。

###  表单实体绑定
使用PostForm这种单个获取属性和字段的方式，代码量较多，需要一个一个属性进行获取。而表单数据的提交，往往对应着完整的数据结构体定义，其中对应着表单的输入项。gin框架提供了数据结构体和表单提交数据绑定的功能，提高表单数据获取的效率。如下所示：

以一个用户注册功能来进行讲解表单实体绑定操作。用户注册需要提交表单数据，假设注册时表单数据包含三项，分别为：username、phone和password。
```go
type UserRegister struct {
	Username string `form:"username" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}
```
创建了UserRegister结构体用于接收表单数据，通过tag标签的方式设置每个字段对应的form表单中的属性名，通过binding属于设置属性是否是必须。

#### ShouldBindQuery
使用ShouldBindQuery可以实现Get方式的数据请求的绑定。具体实现如下：
```go
func main() {

	engine := gin.Default()

	// http://localhost:8080/hello?name=davie&classes=软件工程
	engine.GET("/hello", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		context.Writer.Write([]byte("hello," + student.Name))

	})

	engine.Run()
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
```

#### ShouldBind
使用ShouldBind可以实现Post方式的提交数据的绑定工作。具体编程如下所示：
```go
func main() {

	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		context.Writer.Write([]byte(register.UserName + " Register "))

	})

	engine.Run()
}

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"pwd"`
}
```

#### ShouldBindJson
当客户端使用Json格式进行数据提交时，可以采用ShouldBindJson对数据进行绑定并自动解析,如下所示：
```go
func main() {

	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println("姓名：" + person.Name)
		fmt.Println("年龄：", person.Age)
		context.Writer.Write([]byte(" 添加记录：" + person.Name))
	})

	engine.Run()
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}
```

当然，除了本案例讲解的三种数据绑定方式外，gin还支持其他的方式，也提供了相应的api供开发者学习和使用。在实际的开发过程中，大家可以慢慢学习。

