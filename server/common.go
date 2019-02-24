package server

import (
	"os"
	"log"
	"net/http"
	"bytes"
	"io"
)

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDirExist(path string) error {
	if exist, err := IsExist(path); err == nil {
		if exist == true {
			log.Printf(" Dir:%s is exist.",path)
		}else{
			//log.Printf(" Dir:%s is not exist.\r\n",path)
			err := os.MkdirAll(path, 0666);
			if err == nil {
				log.Printf("Make dir:%s success.",path)
				return err
			}else{
			}
		}
	}
	return nil
}

func ReadBody(data *http.Request) []byte {
	buffer := bytes.NewBuffer(make([]byte, 0, 512))
	io.Copy(buffer, data.Body)
	temp := buffer.Bytes()
	length := len(temp)
	var body []byte

	if cap(temp) > (length + length / 10) {
		body = make([]byte, length)
		copy(body, temp)
	} else {
		body = temp
	}
	data.Body.Close()
	return body
}
