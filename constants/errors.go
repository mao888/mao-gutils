package constants

import gerrors "github.com/mao888/go-errors"

const (
	ErrorCreateGameSQLCode = 306001
	ErrorCreateGameSQLMsg  = "创建游戏失败"

	ErrorFindGameSQLCode = 306002
	ErrorFindGameSQLMsg  = "查找游戏失败"

	ErrorUpdateGameSQLCode = 306003
	ErrorUpdateGameMsg     = "更新游戏失败"

	ErrorDeleteGameSQL    = 306004
	ErrorDeleteGameSQLMsg = "删除游戏失败"

	ErrorCreateAppSQLCode = 306005
	ErrorCreateAppSQLMsg  = "创建应用失败"

	ErrorFindAppSQLCode = 306006
	ErrorFindAppSQLMsg  = "查找应用失败"

	ErrorUpdateAppSQLCode = 306007
	ErrorUpdateAppMsg     = "更新应用失败"

	ErrorDeleteAppSQL    = 306008
	ErrorDeleteAppSQLMsg = "删除应用失败"

	ErrorCreateFBAccountSQLCode = 306009
	ErrorCreateFBAccountSQLMsg  = "创建Facebook账户失败"

	ErrorFindFBAccountSQLCode = 3060010
	ErrorFindFBAccountSQLMsg  = "查找Facebook账户失败"

	ErrorUpdateFBAccountSQLCode = 306011
	ErrorUpdateFBAccountMsg     = "更新Facebook账户失败"

	ErrorDeleteFBAccountSQL    = 306012
	ErrorDeleteFBAccountSQLMsg = "删除Facebook账户失败"

	ErrorCreateGoogleAccountSQLCode = 306013
	ErrorCreateGoogleAccountSQLMsg  = "创建Google账户失败"

	ErrorFindGoogleAccountSQLCode = 3060014
	ErrorFindGoogleAccountSQLMsg  = "查找Google账户失败"

	ErrorUpdateGoogleAccountSQLCode = 306015
	ErrorUpdateGoogleAccountMsg     = "更新Google账户失败"

	ErrorDeleteGoogleAccountSQLCode = 306016
	ErrorDeleteGoogleAccountSQLMsg  = "删除Google账户失败"

	ErrorDeleteGameAllAuthCode = 306017
	ErrorDeleteGameAllAuthMsg  = "删除所有所属游戏权限"

	ErrorCreateGameNameIsNil    = 306020
	ErrorCreateGameNameIsNilMsg = "游戏name为必传字段，不能为空"

	ErrorCreateGameRepeatCode = 306021
	ErrorCreateGameRepeatMsg  = "游戏ID或名称重复, 已存在"

	ErrorCreateAppRepeatCode = 306022
	ErrorCreateAppRepeatMsg  = "应用ID或名称重复, 已存在"

	ErrorAppAssociationGameCode = 306023
	ErrorAppAssociationGameMsg  = "找不到APP所关联的Game"

	ErrorDelFBVerifyCode = 306024
	ErrorDelFBVerifyMsg  = "此应用拥有facebook账户的关联，请先删除关联账户"

	ErrorDelGoogleVerifyCode = 306025
	ErrorDelGoogleVerifyMsg  = "此应用拥有google账户的关联，请先删除关联账户"

	ErrorFBIsLinkedCode = 306026
	ErrorFBIsLinkedMsg  = "facebook投放账户或投放账户名称已存在"

	ErrorGGIsLinkedCode = 306027
	ErrorGGIsLinkedMsg  = "google投放账户或投放账户名称已存在"

	ErrorGameHasAppCode = 306028
	ErrorGameHasAppMsg  = "此游戏下有尚未删除的应用，请先删除应用"

	ErrorMyGameAuthCode = 306030
	ErrorMyGameAuthMsg  = "暂无游戏权限，请尝试联系管理员添加权限"

	ErrorUnifiedConfigParamCode = 306040
	ErrorUnifiedConfigParamMsg  = "请确保参数统一秘钥配置至少有一项"

	ErrorUnifiedConfigEncryptionCode = 306041
	ErrorUnifiedConfigEncryptionMsg  = "游戏的公钥和私钥不存在"

	ErrorUnifiedConfigAccountGameIDCode = 306042
	ErrorUnifiedConfigAccountGameIDMsg  = "统一密钥配置贝塔账号游戏id已存在"

	ErrorUnifiedKeyDelRedisCode = 306043
	ErrorUnifiedKeyDelRedisMsg  = "删除统一秘钥公钥私钥失败，请检查参数"
)

const (
	ErrorConfigAccountIDParamCode = 306050
	ErrorConfigAccountIDParamMsg  = "配置账户ID参数错误，请检查参数"
	ErrorGameIDParamCode          = 306051
	ErrorGameIDParamMsg           = "游戏ID参数错误，请检查参数"
)

//	礼包码业务错误
const (
	ErrNotFondGiftCodeIDCode = 500001
	ErrNotFondGiftCodeIDMsg  = "找不到礼包码ID"
	ErrSameGiftCodeCode      = 500002
	ErrSameGiftCodeMsg       = "已存在相同礼包码"
)

var (
	ErrPersonGroupAddDupName      = gerrors.New(306061, "新增失败，同游戏下分群名称不能重复")
	ErrPersonGroupDelReferInvalid = gerrors.New(306062, "删除失败， 该人群包有实验在引用")
)

var (
	ErrHttpPostProjectUrlFailed = gerrors.New(306071, "查询失败,URL无法访问或者玩家ID不存在")
	ErrPlayerIdNotFound         = gerrors.New(306072, "此玩家ID不存在")
)

// ErrQueueNameIsNil 邮件管理错误
var (
	ErrQueueNameIsNil   = gerrors.New(30673, "PushToReadyQueue 队列名称为空")
	ErrSendMailFailed   = gerrors.New(30674, "邮件发送失败")
	ErrDeleteMailFailed = gerrors.New(30675, "删除邮件失败，邮件ID不存在")
)

// 活动管理业务错误
var (
	ErrMustCompanyParams   = gerrors.New(306080, "company_id is null")
	ErrMustGameParams      = gerrors.New(306081, "company_id is null")
	ErrRequestServer       = gerrors.New(306082, "URL无法访问、推送失败")
	ErrActivityTitleRepeat = gerrors.New(306083, "此活动名称已存在")
	ErrRequestCallback     = gerrors.New(306084, "callback url is null")
)
