package client

import (
	"fmt"
	upper_soap "git.cloud.top/go/go-soap/pkg/myservice/upper"
	"log"
	"git.cloud.top/go/go-soap/ctx"
)

type Client struct {

}

func Start(flag,ip,port,secKey,arg0 string) {
	log.Print("Start client service...\r\n")
	//todo get report data
	//url := "http://localhost:10080/WSSmCommUpper/WSSmCommUpper?wsdl"
	url := fmt.Sprintf("http://%s:%s/WSSmCommUpper/WSSmCommUpper?wsdl", ip, port)
	log.Printf("The addr is = %s\r\n", url)
	//todo 启动客户端
	c := upper_soap.NewWSSmCommUpper(url, false, nil)
	if flag == "Report_Config_1.00" {
		opcodePlainText := "SystemInfo_View_1.00"
		opcodeEncText := ctx.AES128CBCEnc_WithIV(string(opcodePlainText), secKey)

		//获取系统视图信息
		sysInfo := ReportSystemInfo()
		sysInfoEncText := ctx.AES128CBCEnc_WithIV(sysInfo, secKey)
		//log.Printf("opcode enc data = %s\r\n",opcodeEncText)
		//log.Printf(sysInfoEncText)
		r, err := c.ReportView(&upper_soap.ReportView{
			Arg0:arg0,
			Arg1: opcodeEncText,
			Arg2: sysInfoEncText,
		})

		if err != nil {
			fmt.Println(err)
			return
		}
		log.Printf("ReportView response is = %s\r\n",r)
	} else {

	}
	log.Printf("client process finish.\r\n")

}

func init(){
	log.SetFlags(log.Lshortfile)
}