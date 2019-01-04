package BeegoAccessLog

import (
	"net/http"
	"strconv"
)

func FormatAccessLog(statusCode int64,beegoRequest *http.Request) (accessLog string) {

	accessLog = "StatusCode:"+strconv.FormatInt(statusCode,10)+ "-"+ "Method:"+beegoRequest.Method + "-" +
		"RequestURI:"+ beegoRequest.RequestURI + "-" + "Host:" + beegoRequest.Host + "-" + "UserAgent" + beegoRequest.UserAgent()

	return
}
