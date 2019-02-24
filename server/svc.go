package server
import (
	"encoding/xml"
	"fmt"
	lower_soap "git.cloud.top/go/go-soap/pkg/myservice/lower"
)

func processConfig(body []byte) *lower_soap.FillConfigResponse {
	envlop := &lower_soap.SOAPEnvelope{
		Body: lower_soap.SOAPBody{
			Content: &lower_soap.FillConfig{},
		},
	}
	err := xml.Unmarshal(body, envlop)
	if err != nil {
		fmt.Println("xml Unmarshal error:", err)
		return nil
	}
	_, ok := envlop.Body.Content.(*lower_soap.FillConfig)
	if !ok {
		return nil
	}else{
		return &lower_soap.FillConfigResponse{
			Return_:0,
		}
	}
	//config不返回数据
	return nil
}