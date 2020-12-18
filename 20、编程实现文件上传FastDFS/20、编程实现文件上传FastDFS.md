## 编程实现文件上传FastDFS

### 功能介绍
在前面的三节课中，我们已经完成了文件上传保存到项目目录的功能，也搭建完成了FastDFS分布式文件系统。接下来，我们要做的就是，将在我们的项目中，使用代码编程的方式，实现将前端提交的文件，上传保存到FastDFS文件系统中，并返回文件的连接，然后保存在对应的数据库中。

### Go语言编程上传FastDFS
编程实现FastDFS上传功能可以通过以下几个步骤来完成。

#### 1、安装fastdfs的golang库
在项目中进行文件上传编程实现，需要安装一个go语言库，该库名称为fdfs_client，通过如下命令进行安装。
```go
go get github.com/tedcy/fdfs_client
```
下载后，可以在$GOPATH/src/github.com/tedcy目录下找到fdfs_client库的源文件。

#### 2、编写fdfs.conf配置文件
在fdfs_client库中，提供对文件的上传和下载方法，其中文件上传支持两种方式。

要使用文件上传功能方法，首先需要构造一个fdfsClient实例。如同我们前文讲的fastDFS的组件构成一样，client需要连接tracker服务器，然后进行上传。

在构造fdfsClient实例时，首先需要编写配置文件fdfs.conf，在fdfs.conf文件中进行配置选项的设置：
```go
tracker_server=114.246.98.91:22122
http_port=http://114.246.98.91:80
maxConns=100
```
在fdfs.conf配置文件中，配置了三个选项，分别是：
* tracker_server：跟踪服务器的ip和跟踪服务的端口
* http_port：配置了nginx服务器后的http访问地址和端口
* maxConns：最大连接数为100，默认值即可

在构造fdfsClient对象实例化时，会使用该文件。

#### 3、文件上传编程实现
将文件上传功能作为全局的一个工具函数进行定义，实现文件上传功能，并返回保存后的文件的id。编程实现如下：
```go
func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()

	if err != nil {
		fmt.Println(err.Error())
	}

	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return fileId
}
```
在自定义UploadFile函数中，通过fdfs_client.NewClientWithConfig实例化client对象，并调用UploadByFilename方法实现文件上传，该方法接收一个文件名称，该文件名称包含文件的路径，最后返回上传的fileId。

#### 4、修改Controller文件上传方法
现在，已经接入了fastDFS文件系统，因此，对MemberController的uploadAvator方法进行修改。

修改思路：将客户端上传的文件，先保存在服务器目录下的uploadfile目录中，然后将文件的路径和名称作为参数传递到UploadFile函数中，进行上传。上传成功后，将保存到本地的uploadfile文件删除，并把保存到fastDFS系统的fileId更新到对应用户记录的数据库。最后拼接文件访问的全路径，返回给客户端。

依照上述思路，修改后的uploadAvator方法逻辑实现如下所示：
```go
func (mc *MemberController) uploadAvator(context *gin.Context) {

	//1、获取上传的文件
	userId := context.PostForm("user_id") //用户id
	fmt.Println(userId)
	file, err := context.FormFile("avatar")
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//从session中获取用户信息
	var member model.Member
	sess := tool.GetSess(context, "user_"+userId)
	if sess == nil {
		tool.Failed(context, "参数不合法")
		return
	}

	json.Unmarshal(sess.([]byte), &member)

	//2、将文件保存到本地
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	fmt.Println(err)
	if err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}

	//3、将文件上传到分布式文件系统
	fileId := tool.UploadFile(fileName)

	if fileId != "" {
		//删除本地文件
		os.Remove(fileName)

		//4、将文件对应路径更新到数据库中
		memberService := impl.NewMemberService()
		path := memberService.UploadAvator(3, fileId)
		fullPath := tool.FileServerAddr() + "/" + path
		tool.Success(context, fullPath)
		return
	}
	tool.Failed(context, "上传失败")
}
```

在最后返回客户端数据前，需要对上传文件的访问全路径进行拼接。因此，提供FileServerAddr函数，用户解析fdfs.conf文件，并提取其中的http_port选项。

FileServerAddr函数的实现如下所示：
```go
func FileServerAddr() string {
	f, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
```



