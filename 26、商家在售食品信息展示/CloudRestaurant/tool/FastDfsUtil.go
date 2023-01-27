package tool

import (
	"bufio"
	"fmt"
	"github.com/tedcy/fdfs_client"
	"os"
	"strings"
)

/**
 * 上传文件到fastDFS系统
 */
func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return fileId
}

/**
 * 从配置文件中读取文件服务器的ip和端口相关配置
 */
func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_server_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
