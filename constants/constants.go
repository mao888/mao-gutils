package constants

const (
	LoggerServerCode = "s_code"
)

const (
	ServiceCode = "30600"
	ServiceName = "application-console"
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
	DBApplicationConsole = "application_console"
	DBPg                 = "fotoabledb"
)

const (
	RedisName                  = "redis"
	RedisApplicationSecret     = "application_console:secret_key"
	RedisApplicationSecretSize = 2048
)

const (
	AuthCodeGame = "Game" // 获取游戏权限标识
)

// 统一配置
const (
	SensorsCloseTestUrl = "https://pay.ftstats.com/"
	SensorsCloseProdUrl = "https://pay.ftstats.com/"

	SensorsOpenTestUrl = "https://cdh-sc.nuclearport.com/sa?project=%s"
	SensorsOpenProdUrl = "https://cdh-sc.nuclearport.com/sa?project=%s_test"

	DefaultPage  = 1
	DefaultLimit = 10
)

const (
	TopicAddPersonGroupTypeAdd = "add"
	TopicAddPersonGroupTypeDel = "del"
	TopicAddPersonGroup        = "user_grouping"
)

// GM管理 (邮件管理、活动管理)
const (
	EnvProd = 1
	EnvTest = 2

	GMConfigMailType         = 1
	GMConfigActivityType     = 2
	GMConfigPlayerSearchType = 3
	GMConfigCommandType      = 4

	MailAttributesCarryItems   = "carry_items"
	MailAttributesActionButton = "action_button"
	MailAttributesRedirectURI  = "redirect_uri"

	GMConfigCallbackTestUrl       = "application_console:gm:callback:test:%s"
	GMConfigCallbackProdUrl       = "application_console:gm:callback:prod:%s"
	GMConfigCallbackMailField     = "mail_url"
	GMConfigCallbackActivityField = "activity_url"

	MailsBucket           = "gm:mails:bucket"
	MailQueue             = "gm:mails:queue"
	MailLock              = "gm:mails:lock"
	DefaultMailLockExpire = 25

	DefaultMailStatus = "pending"
	MailStatusFinish  = "finish"
	MailStatusPass    = "pass"
	MailStatusTiming  = "timing"
	MailStatusRevoke  = "revoke"
	MailStatusReject  = "reject"

	MailTypeNowSend    = 1 // 邮件类型立即发送
	MailTypeTimingSend = 1 // 邮件类型定时发送

	SyncServerCallbackUrl = "application_console:gm:callback:" // redis 获取回调的Key 格式:gm:callback:{test:prod}:{游戏ID}
)

// SendManageEnvType 服务器类型
var SendManageEnvType = map[int32]string{
	1: "prod",
	2: "test",
}
