package bcGolang

import (
	"fmt"
	"reflect"
	"time"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"math/rand"
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
)


func AttachAppSign(reqParams *BCReqParams, reqestType BCReqestType, bcApp BCApp) {
	if StrEmpty(bcApp.AppId) {
		// raise error. Maybe panic
		fmt.Println("exception: app_id is empty")
	}
	reqParams.AppId = bcApp.AppId
	reqParams.Timestamp = time.Time.UnixNano()/1000000 // 毫秒数

	if bcApp.IsTestMode {
		if StrEmpty(bcApp.TestSecret) {
			fmt.Println("exception: test secret empty")
		} else {
			reqParams.AppSign = getMd5SignString(bcApp.AppId + strconv.FormatInt(reqParams.Timestamp, 10) + bcApp.TestSecret)
		}
	} else {
		if (reqestType == TRANSFER) || (reqestType == REFUND) {
			if StrEmpty(bcApp.MasterSecret) {
				fmt.Println("exception: master secret empty")
			} else {
				reqParams.AppSign = getMd5SignString(bcApp.AppId + strconv.FormatInt(reqParams.Timestamp, 10) + bcApp.MasterSecret)
			}
		} else {
			// QUERY or PAY
			if StrEmpty(bcApp.AppSecret) {
				fmt.Println("exception: app_secret empty")
			} else {
				reqParams.AppSign = getMd5SignString(bcApp.AppId + strconv.FormatInt(reqParams.Timestamp, 10) + bcApp.AppSecret)
			}
		}
	}
}

func getMd5SignString(origText string) string{
	srcdata := []byte(origText)
	hash := md5.New()
	hash.Write(srcdata)
	cipherText := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText)
	fmt.Println("md5 encrypto is ", string(hexText))
	return string(hexText)
}

func GetRandomHost() string {
	rand.Seed(time.Now().UnixNano())
	return BEECLOUD_HOSTS[rand.Intn(4)] + BEECLOUD_RESTFUL_VERSION
}

func HttpPost(reqUrl string, para MapObject) (content []byte, ok bool) {
	body, _ = json.Marshal(para)
	response, err1 := http.Post(reqUrl, "application/json", strings.NewReader(string(body)))
	defer response.Body.Close()
	if err1 != nil {
		fmt.Println("exception: http post error")
	}
	content, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("exception: http content read error")
		// should panic exception
	}
	if 200 != response.StatusCode {
		return content, false
	}
	return content, true
}

func NotSupportedTestError(methodName string) BCResult {
	var errResult BCResult
	errResult.ResultCode = NOT_SUPPORTED_CODE
	errResult.ResultMsg = NOT_SUPPORTED_NAME
	errResult.ErrDetail = methodName + "is not supported in test mode currently!"
	return errResult
}

func HandleInvalidResp(content []byte) BCResult {
	return BCResult{
		result_code: NETWORK_ERROR_CODE,
		result_msg: resp.Status,
		err_detail: string(content)}
}
// private methods
