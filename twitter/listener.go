package twitter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var TWITTER_SECRET = os.Getenv("TWITTER_SECRET")
var TWITTER_API_KEY = os.Getenv("TWITTER_API_KEY")

type crcResponse struct {
	ResponseToken string `json:"response_token"`
}
func CrcCheck(writer http.ResponseWriter, request *http.Request) {
	token := request.URL.Query()["crc_token"]
	if len(token) < 1 {
		fmt.Println("No crc_token given")
		fmt.Fprintf(writer, "No token given")
		return
	}
	h := hmac.New(sha256.New,[]byte(TWITTER_SECRET))
	h.Write([]byte(token[0]))
	encoded := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil))))
	response,_ := json.Marshal(crcResponse{ResponseToken:"sha256="+encoded})
	fmt.Fprintf(writer,string(response))
}
