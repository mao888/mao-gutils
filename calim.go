package utils

import (
	"Project/model/frontdesk"
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	CtxUserIdAndName = "user"
	CtxUserIDKey     = "user_id"
	CtxLibraryIdKey  = "library_id"
	CtxUserInfoKey   = "info"
	CtxUserOnline    = "user_online"
	CtxDevice        = "device"
)

type UserMessage struct {
	UserId    int64
	LibraryId int64
}

var (
	ErrorUserNotLogin     = errors.New("用户未登录")
	InvalidUseOfEquipment = errors.New("登陆设备无效")
)

// GetCurrentUserId 获取当前登录的用户ID
func GetCurrentUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetCurrentLibraryId 获取当前登录的图书馆ID
func GetCurrentLibraryId(c *gin.Context) (LibraryId int64, err error) {
	uid, ok := c.Get(CtxLibraryIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	LibraryId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

//GetUserMessage 获取当前登录的用户ID和图书馆id
func GetUserMessage(c *gin.Context) (*UserMessage, error) {
	res, ok := c.Get(CtxUserIdAndName)
	if !ok {
		err := ErrorUserNotLogin
		return nil, err
	}
	userMessage := res.(*UserMessage)
	return userMessage, nil
}

//GetUserInfo 获取当前登录的用户所有信息
func GetUserInfo(c *gin.Context) (*frontdesk.OnlineUser, error) {
	res, ok := c.Get(CtxUserInfoKey)
	if !ok {
		err := ErrorUserNotLogin
		return nil, err
	}
	userInfo := res.(*frontdesk.OnlineUser)
	return userInfo, nil
}

// GetDevice 获取上下文中的device
func GetDevice(c *gin.Context) (string, error) {
	deviceInto, exists := c.Get(CtxDevice)
	if !exists {
		err := InvalidUseOfEquipment
		return "", err
	}
	device := deviceInto.(string)
	return device, nil
}
