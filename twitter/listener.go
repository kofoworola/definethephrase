package twitter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type crcResponse struct {
	ResponseToken string `json:"response_token"`
}

func CrcCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type","application/json")
	fmt.Printf("Crc check occurred at %s\n",time.Now().String())
	token := request.URL.Query()["crc_token"]
	if len(token) < 1 {
		fmt.Println("No crc_token given")
		fmt.Fprintf(writer, "No token given")
		return
	}
	h := hmac.New(sha256.New, []byte(os.Getenv("CONSUMER_SECRET")))
	h.Write([]byte(token[0]))
	encoded := base64.StdEncoding.EncodeToString(h.Sum(nil))
	response, _ := json.Marshal(crcResponse{ResponseToken: "sha256=" + encoded})
	fmt.Fprintf(writer, string(response))
}

func WebhookHandler(writer http.ResponseWriter, request *http.Request){
	body,_ := ioutil.ReadAll(request.Body)
	fmt.Println(string(body));
}