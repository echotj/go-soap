package client

import (
	"fmt"
	upper_soap "git.cloud.top/go/go-soap/pkg/myservice/upper"
	"log"
)

func Start() {
	log.Print("Start client service...\r\n")
	//todo 需要修改自己想发送请求的地址+port和URI
	ip := "localhost"
	port := "8009"
	url := fmt.Sprintf("http://%s:%s/clientup", ip, port)
	//todo 启动客户端
	c := upper_soap.NewWSSmCommUpper(url, false, nil)
		r, err := c.ReportView(&upper_soap.ReportView{
			Arg0:"This is demo",
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	log.Printf("client process finish.\r\n")
}
