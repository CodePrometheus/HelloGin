module CloudRestaurant

go 1.16

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1089
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-xorm/xorm v0.7.9
	github.com/kr/pretty v0.2.1 // indirect
	github.com/mojocn/base64Captcha v1.3.4
	github.com/tedcy/fdfs_client v0.0.0-20200106031142-21a04994525a
	github.com/wonderivan/logger v1.0.0
)

replace github.com/mojocn/base64Captcha v1.3.4 => github.com/mojocn/base64Captcha v1.2.2
