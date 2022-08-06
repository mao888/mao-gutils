package app

type ResCode int64

const (
	// 成功（默认返回状态码）
	CodeSuccess ResCode = 0
	// 全局未知异常
	CodeSeverError ResCode = 500
	// 请求失败（一般前端处理，不常用）
	CodeBadRequest ResCode = 400
	// 请求资源不存在（静态资源不存在，不常用）
	CodeDataNotFount ResCode = 404
	// 登录、权限认证异常
	CodeLoginExpire ResCode = 401
	// 权限不足
	CodeIdentityNotRow ResCode = 403
)

/*
	通用业务
*/
const (
	/*
	   1001-1010 通用操作相关
	*/
	// 操作失败
	CodeOperationFail ResCode = 1001 + iota
	// 查询操作失败
	CodeSelectOperationFail
	// 更新操作失败
	CodeUpdateOperationFail
	// 删除操作失败
	CodeDeleteOperationFail
	// 新增操作失败
	CodeInsertOperationFail
	// 新增考评重复
	CodeInsertRepeat

	/*
	   1011-1050 登录注册相关
	*/
	// 登录失败，账号或者密码错误
	CodeLoginFailResCode ResCode = 1011 + iota
	// 登录失败，请重试
	CodeLoginFailReLogin
	// 验证码错误
	CodeLoginFailCode
	// 无效token
	CodeInvalidToken
	// 用户不存在
	CodeNoUser
	// 注册失败，手机号已经存在
	CodeRegisterFail
	// 认证失败，手机号不存在
	CodeNoUserPhone
	// 请求参数不能为空
	CodeParamsNotNull
	// 用户未激活
	CodeUserIsNotEnabled
	// 角色名字已存在
	CodeRoleNameExist
	// 用户名字已存在
	CodeUserNameExist
	// 用户信息丢失
	CodeUserNotFind
	// wx或QQ登陆过期
	CodeWXOrQQLoginLandExpiration
	// 获取用户信息失败
	CodeFailedGetUserInfo

	/*
	   1051-1070 短信业务相关
	*/
	// 短信发送失败
	CodeSMSNotSend ResCode = 1051 + iota
	// 短信验证码失效
	CodeSMSCodeExpire
	// 短信验证码验证失败
	CodeSMSVerityFail

	/*
	   1071-1100 文件、资源相关
	*/
	// 文件超出规定大小
	CodeFileOverstepSize ResCode = 1071 + iota
	// 文件上传失败
	CodeFileUploadFail
	// 文件不存在，加载失败
	CodeFileLoadingFail
	// 文件类型不支持查看
	CodeFileRequestFail
	// 图片不能为空
	CodeImageIsNotNull
	// 请上传图片类型的文件
	CodeFileImageFail
	//文件导入错误
	CodeFileImportFail

	/*
	   1101-1199 请求参数相关
	*/
	// 参数无效
	CodeParamIsInvalid ResCode = 1101 + iota
	// 参数为空
	CodeParamIsBlank
	// 参数类型错误
	CodeParamTypeBindError
	// 参数缺失
	CodeParamNotComplete
	CodeActContentIsExist
	// 重定向失败
	CodeRedirectionFailed
)

/*
   -----------go_api 业务相关（2xxx）------------
*/
const (
	CodeArticleIsNotApply ResCode = 2001 + iota
	CodeMarchIsNotApply

	CodeReplaceDevice
	FailedAccessFailure
	FailedToGetMyRanking
	FailedGetRanking
	FailDepMasterDelete
	FailStudentIsNotAdmin

	// 获取用户部门id失败
	CodeUserDepartmentInfoFail
	// 用户已经被添加
	CodeFailedStudentHasExisted
)

/*
   第三方相关（3xxx）
*/
const (
	/*
	   3001-3020 微信公众号
	*/
	// 微信公众号JSSDK获取access_token失败
	CodeWxGzhAccessTokenFail = 3001 + iota
	// 微信公众号JSSDK获取jsapi_ticket失败
	CodeWxGzhJsApiTicketFail
	// 微信公众号JSSDK获取SIGN失败
	CodeWxGzhSignFail
	// 微信wxCode为空
	CodeWxEmpty
	// 微信wxCode失效或不正确请重新获取
	CodeWxOuttime
	// 微信生成二维码失败
	CodeWxTickerFail
	//微信登录用户为扫码
	UserWxTickerFail
	// 获取微信用户Openid 失败
	CodeWxOpenidFail
	// 添加微信用户信息失败
	CodeAddWXUserInfoFail
	// 获取微信用户信息失败
	CodeWXInfoFail
	// 创建http get请求失败
	CodeWxGetCreateFailure
	// Code解码错误
	CodeWxDeCodeFailure
	// 微信接口返回异常
	CodeWxInterfaceReturnException
)

const (
	/*
	   3021 - 3040 QQ公众号
	*/
	// QQ Code为空
	JSCodeQQEmpty = 3021 + iota
	// QQwxCode失效或不正确请重新获取
	CodeQQOuttime
	// 获取QQ用户信息失败
	CodeQQInfoFail
)

const (
	/*
	   3041- QQ小程序
	*/
	// 添加QQ用户信息失败
	CodeAddQQUserInfoFail = 3041 + iota
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:             "success",
	CodeSeverError:          "服务器繁忙请重试",
	CodeBadRequest:          "请求失败",
	CodeDataNotFount:        "未找到资源",
	CodeLoginExpire:         "请登录后重试",
	CodeIdentityNotRow:      "权限不足",
	CodeUserNameExist:       "用户名字已存在",
	CodeOperationFail:       "操作失败",
	CodeSelectOperationFail: "查询操作失败！",
	CodeUpdateOperationFail: "更新操作失败！",
	CodeDeleteOperationFail: "删除操作失败！",
	CodeInsertOperationFail: "新增操作失败！",
	CodeInsertRepeat:        "新增考评重复",

	CodeLoginFailResCode:          "登录失败，账号或者密码错误",
	CodeLoginFailReLogin:          "登录失败，请重试",
	CodeLoginFailCode:             "验证码错误",
	CodeInvalidToken:              "无效的token",
	CodeNoUser:                    "用户不存在",
	CodeRegisterFail:              "注册失败，手机号已经存在",
	CodeNoUserPhone:               "认证失败，手机号不存在",
	CodeParamsNotNull:             "请求参数不能为空",
	CodeUserIsNotEnabled:          "用户未激活",
	CodeRoleNameExist:             "角色名字已存在",
	CodeUserNotFind:               "用户数据丢失",
	CodeWXOrQQLoginLandExpiration: "登陆过期",
	CodeFailedGetUserInfo:         "获取用户信息失败",

	CodeSMSNotSend:    "短信发送失败",
	CodeSMSCodeExpire: "短信验证码失效",
	CodeSMSVerityFail: "短信验证码验证失败",

	CodeFileOverstepSize: "文件超出规定大小",
	CodeFileUploadFail:   "文件上传失败",
	CodeFileLoadingFail:  "文件不存在，加载失败",
	CodeFileRequestFail:  "文件类型不支持查看",
	CodeImageIsNotNull:   "图片不能为空",
	CodeFileImageFail:    "请上传图片类型的文件",

	CodeParamIsInvalid:     "参数无效",
	CodeParamIsBlank:       "参数为空",
	CodeParamTypeBindError: "参数类型错误",
	CodeParamNotComplete:   "参数缺失",
	CodeRedirectionFailed:  "重定向失败",

	CodeArticleIsNotApply: "文章不是发布未审核状态",
	CodeMarchIsNotApply:   "该用户未申请三月圈",

	CodeWxGzhAccessTokenFail: "微信公众号JSSDK获取access_token失败",
	CodeWxGzhJsApiTicketFail: "微信公众号JSSDK获取jsapi_ticket失败",
	CodeWxGzhSignFail:        "微信公众号JSSDK获取SIGN失败",
	CodeWxEmpty:              "微信wxCode为空",
	CodeWxOuttime:            "微信wxCode失效或不正确请重新获取",
	CodeWxOpenidFail:         "获取微信用户信息失败",
	CodeAddWXUserInfoFail:    "添加微信用户信息失败",
	CodeWXInfoFail:           "获取微信用户信息失败",

	CodeAddQQUserInfoFail: "添加QQ用户信息失败",
	JSCodeQQEmpty:         "QQ登陆Code为空",
	CodeQQInfoFail:        "获取QQ用户信息失败",

	CodeActContentIsExist: "活动内容已存在",

	CodeReplaceDevice:    "请更换设备重试",
	FailedToGetMyRanking: "获取我的名次失败",
	FailedAccessFailure:  "访问数据失败",
	FailedGetRanking:     "获取排名失败",
	FailDepMasterDelete:  "管理员不能被删除",

	CodeUserDepartmentInfoFail:  "请使用学生会成员账号登陆",
	CodeFailedStudentHasExisted: "用户已经被添加",
	FailStudentIsNotAdmin:       "该学生不是管理员",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeSeverError]
	}
	return msg
}
