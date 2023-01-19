package constants

const (
	LoggerServerCode = "s_code"
)

const (
	ServiceCode = "30800"
	ServiceName = "version-console"
)

//header field
const (
	HeaderXRequestID = "X-Request-ID"
	HeaderAppID      = "Fotoable-App-ID"
	HeaderSdkVer     = "Fotoable-Sdk-Version"
	HeaderAppVer     = "Fotoable-App-Version"
	HeaderError      = "Fotoable-Error"
	HeaderSign       = "sign"
	HeaderTimestamp  = "timestamp"
)

const (
	InitZero    = 0
	EmptyString = ""
	Point       = "."
	Underline   = "_"
	MiddleLine  = "-"
	Comma       = ","
	Colon       = ":"
)

const (
	TimeYMDHMM   = "2006-01-02 15:04:05"
	TimeYMD      = "2006-01-02"
	TimeYMDH     = "2006-01-02 15"
	TimeYM       = "2006-01"
	TimeCNYM     = "01月02日"
	TimeYYYYMMDD = "20060102"
	TimeYMDZ     = "2006-01-02T15:04:05Z0700"
)

const (
	DBIsNotDeleted = 0
	DBIsDeleted    = 1
)

const (
	RpcOKCode = 0
)

const (
	TopicAddPersonGroupTypeAdd = "add"
	TopicAddPersonGroupTypeDel = "del"
	TopicAddPersonGroup        = "user_grouping"
)

const (
	url         = "https://oapi.dingtalk.com/robot/send?access_token=88d05964cfbea4879309d721f11c16e15f5cea2b529fabc22c7e836108b08e0d"
	DingTalkURL = "配置自己的钉钉url"
)

// SendManageEnvType 服务器类型
var SendManageEnvType = map[int32]string{
	1: "prod",
	2: "test",
}

const (
	NumberZero  = 0
	NumberOne   = 1
	NumberTwo   = 2
	NumberThree = 3
	NumberFour  = 4
	NumberFive  = 5
	NumberSix   = 6
	NumberSeven = 7
	NumberEight = 8
	NumberNine  = 9
)
