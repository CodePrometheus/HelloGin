## MySQL数据库安装

### 一、下载MySQL

* 1、访问MySQL的官方网站：[https://www.mysql.com/](https://www.mysql.com/)
* 2、在页面中找到Downloads进入下载页面：[https://www.mysql.com/downloads/](https://www.mysql.com/downloads/)
* 3、选择MySQL Community
  Server进入下载详情页面：[https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)。页面会自动根据系统类型进行匹配，如下图所示：

![MySQL下载](./img/WX20191008-144205@2x.png)

等待下载结束，下载过程中保持网络环境稳定。

如果MySQL官方网站访问和下载速度过慢，可以访问国内的镜像源，如搜狐的镜像源如下所示：[http://mirrors.sohu.com/](http://mirrors.sohu.com/)
![MySQL的国内镜像源](./img/WX20191008-150927@2x.png)

### 二、安装MySQL

打开MySQL的安装包：
![mysql安装包](./img/WX20191008-151411@2x.png)

选择强密码：
![选择强密码](./img/WX20191008-151512@2x.png)

输入默认用户名root的密码(至少8位）：
![输入8位密码](./img/WX20191008-151614@2x.png)

安装完毕后，mysql就会自动启动了。在macOS中，可以到系统偏好设置里面的服务列表中查看macOS已经启动的服务,如下图所示：

![服务列表里的mysql](./img/WX20191008-152058@2x.png)

三、mysql的安装路径和配置 在macOS系统下，默认的mySQL的安装路径是：/usr/local/ 在windows系统下，默认的mySQL的安装路径在C盘：

进入到8.0的安装目录如下所示：
![macOS系统下的mySQL安装目录](./img/WX20191008-153513@2x.png)

用户登录 安装好了mySQL，并自动启动起来。用户如果想要操作mySQL，只能通过命令行进行登录。进入到mysql的安装目录中的bin目录下,执行登录命令：

```
./mysql -u root -p
```

输入安装mysql时填写的密码（至少8位），点击回车，即可登录mysql。登录成功后，光标变为由mysql>开始，如下所示：
![用户命令行登录](./img/WX20191008-153158@2x.png)

用户退出 执行exit命令退出当前用户登录状态。

```
exit
```

配置mysql环境变量 将mysql可执行文件所在的目录作为环境变量进行全局配置。
macOS系统下：/usr/local/mysql-8.0.15-macos10.14-x86_64/bin。将该路径配置到全局的配置文件.bash_profile中，并起别名mysql。如下图所示：
![配置mysql环境变量](./img/WX20191008-155430@2x.png)

window下的配置环境变量，可以同过图形化界面进行操作，将mysql的可执行文件目录添加到path目录中。

Window目录下的配置 1、下载windows5.7.x压缩包 下载链接：mysql官网或者国内镜像。
mysql官网：[https://downloads.mysql.com/archives/community/](https://downloads.mysql.com/archives/community/)
镜像：[http://mirrors.sohu.com/](http://mirrors.sohu.com/)

2、解压安装包并配置环境变量
![解压缩目录](./img/WX20191014-002041@2x.png)
右击我的电脑/计算机，点击属性，打开高级系统设置，点击环境变量。 变量名：MYSQL_HOME 变量值：C:\Program Files\MySQL

path里添加：%MYSQL_HOME%\bin

3、创建数据库及配置文件 此版本MySQL并没有创建data目录及my.ini。在MYSQL_HOME目录下创建data目录，建议将data目录设为C:\Program
Files\MySQL\data。另外，创建Uploads目录作为MySQL导入导出的目录。my.ini建议放在MYSQL_HOME目录下，简单配置可参考：

```
[mysqld]
port=3306
character_set_server=utf8
basedir=C:\Program Files\MySQL
datadir=C:\Program Files\MySQL\data
server-id=1
sql_mode=NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
lower_case_table_names=1
innodb_file_per_table = 1
log_timestamps=SYSTEM

log-error = error.log
slow_query_log = 1
slow_query_log_file = slow.log
long_query_time = 5
log-bin = binlog
binlog_format = row
expire_logs_days = 15
log_bin_trust_function_creators = 1
secure-file-priv=C:\Program Files\MySQL\Uploads

[client]   
default-character-set=utf8
```

注意：将basedir、datadir和secure-file-priv三个变量，替换成自己的安装目录和文件所在的目录。basedir指的是解压缩的mysql的目录，datadir指的是创建的data目录的目录，secure-file-priv指的是uploads的所在目录。

4、初始化数据库 cmd命令行进入C:\Program Files\MySQL\bin目录，执行命令:

```
mysqld --initialize-insecure
```

执行完毕之后，在data目录下会生成很多文件。
![初始化数据库](./img/WX20191014-003025@2x.png)

5、注册并启动MySQL服务 执行mysqld –install MySQL57安装服务（install后面是服务的名字，我们这里以MySQL57作为mysql5.7的服务名）net start
MySQL57启动MySQL服务。具体命令为,在mysql的bin目录下：

 ```
 mysqld -install MySQL57
 ```

将mysql安装到系统服务，注册到注册表。 随后，启动mySQL：

 ```
 net start MySQL57
 ```

如果遇到报错，使用系统管理员角色运行该命令。


