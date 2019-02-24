package server

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	lower_soap "git.cloud.top/go/go-soap/pkg/myservice/lower"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"git.cloud.top/go/go-soap/client"
)

const (
	USER_ECHO_DEBUG = true

)
func ServerStart() {
	log.Printf("go-soap server is running...\r\n")
	if USER_ECHO_DEBUG{
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.HideBanner = true
		e.POST("/go-soap/sever",soapServerHandler)
		e.Logger.Fatal(e.Start(":8001"))

	}else{
		//todo no echo
		s := NewSOAPServer("0.0.0.0:8001")
		log.Fatal(s.ListenAndServe())
	}
}

func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", soapHandler)
	return mux
}

func NewSOAPServer(addr string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}
	return server
}

//todo Use Echo
func soapServerHandler(c echo.Context)error {
	rawBody, err := ioutil.ReadAll(c.Request().Body)
	var res interface{}
	config := regexp.MustCompile(`<ns2:fillConfig xmlns:ns2=`)
	if config.MatchString(string(rawBody)) {
		//下发参数配置
		res = processConfig(rawBody)
	} else {
		res = lower_soap.FillConfigResponse{
			Return_: 1,
		}
		fmt.Println("the method requested is not available")
	}
	v := lower_soap.SOAPEnvelope{
		Body: lower_soap.SOAPBody{
			Content: res,
		},
	}

	log.Printf("process finished.\r\n")
	//todo 根据需要，可以启动soap的客户端
	//todo 在服务端完成HTTP通信，返回请求后，可以启动soap客户端，完成soap的单向通信
	{
		defer client.Start()
	}
	c.Response().Writer.Header().Set("Content-Type", "text/xml")
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return nil
	}
	c.Response().WriteHeader(http.StatusOK)
	c.Response().Write(x)

	return c.XML(http.StatusOK, &x)
}

//todo No ECHO
func soapHandler(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// match method
	var res interface{}
	config := regexp.MustCompile(`<ns2:fillConfig xmlns:"`)
	if config.MatchString(string(rawBody)) {
		//下发参数配置
		res = processConfig(rawBody)
	}else{
		res = nil
		fmt.Println("the method requested is not available")
	}
	log.Printf("process finished.\r\n")
	v := lower_soap.SOAPEnvelope{
		Body: lower_soap.SOAPBody{
			Content: res,
		},
	}
	w.Header().Set("Content-Type", "text/xml")
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(x)
	return
}

func init(){
	log.SetFlags(log.Lshortfile)
}