package testservice

import (
	"encoding/json"
	"fullstack_toko/backend/model/web"
	"io"
	"log"
	"net/http/httptest"
)

// if result its nil ,mean its got error
func WebResponseUnmarshal(recorder *httptest.ResponseRecorder, data any) (result *web.WebResponse) {
	bytes, err := io.ReadAll(recorder.Body)
	if err != nil {
		log.Println("error inside func WebResponseUnmarshal,error whil read the recorder body,message:", err)
		result = nil
		return
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		if syntaxErr, ok := err.(*json.SyntaxError); ok {
			log.Printf("Syntax error at byte offset %d", syntaxErr.Offset)
		}
		log.Println("error inside func WebResponseUnmarshal,error whil unmarshal json,message:", err)
		result = nil
		return
	}

	result = data.(*web.WebResponse)

	return
}
