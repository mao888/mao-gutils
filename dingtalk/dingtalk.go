package gutil

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mao888/go-utils/http"

	"github.com/mao888/go-utils/constants"
)

const (
	template = `{"msgtype":"text","text":{"content":"%s"}}`
)

func DingTalkAlarm(serverName, message string) bool {
	hostname, _ := os.Hostname()
	message = fmt.Sprintf("[业务异常][服务：%s][主机名：%s] %s", serverName, hostname, message)
	httpStatus, _, _ := gutil.HttpPostJson(constants.DingTalkURL, []byte(fmt.Sprintf(template, message)), nil)
	return httpStatus == http.StatusOK
}

func DingTalkAlarmUrl(url, serverName, message string) bool {
	if len(url) == 0 {
		return false
	}
	hostname, _ := os.Hostname()
	message = fmt.Sprintf("[服务：%s][主机名：%s] %s", serverName, hostname, message)
	httpStatus, _, _ := gutil.HttpPostJson(url, []byte(fmt.Sprintf(template, message)), nil)
	return httpStatus == http.StatusOK
}
