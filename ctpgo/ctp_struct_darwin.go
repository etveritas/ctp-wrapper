/*
 * @Author: Wavy Wei
 * @Date: 2022-12-12 20:41:51
 * @Last Modified by: Wavy Wei
 * @Last Modified time: 2025-04-06 20:52:56
 */

package ctpgo

// 信息分发
type CThostFtdcDisseminationField struct {
	// 序列系列号
	SequenceSeries TThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcReqUserLoginField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 密码
	Password TThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 动态密码
	OneTimePassword TThostFtdcPasswordType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// 终端IP端口
	ClientIPPort TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcRspUserLoginField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 登录成功时间
	LoginTime TThostFtdcTimeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 交易系统名称
	SystemName TThostFtdcSystemNameType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 最大报单引用
	MaxOrderRef TThostFtdcOrderRefType
	// 上期所时间
	SHFETime TThostFtdcTimeType
	// 大商所时间
	DCETime TThostFtdcTimeType
	// 郑商所时间
	CZCETime TThostFtdcTimeType
	// 中金所时间
	FFEXTime TThostFtdcTimeType
	// 能源中心时间
	INETime TThostFtdcTimeType
	// 后台版本信息
	SysVersion TThostFtdcSysVersionType
	// 广期所时间
	GFEXTime TThostFtdcTimeType
}

// 信息分发
type CThostFtdcUserLogoutField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcForceUserLogoutField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcReqAuthenticateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 认证码
	AuthCode TThostFtdcAuthCodeType
	// App代码
	AppID TThostFtdcAppIDType
}

// 信息分发
type CThostFtdcRspAuthenticateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// App代码
	AppID TThostFtdcAppIDType
	// App类型
	AppType TThostFtdcAppTypeType
}

// 信息分发
type CThostFtdcAuthenticationInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 认证信息
	AuthInfo TThostFtdcAuthInfoType
	// 是否为认证结果
	IsResult TThostFtdcBoolType
	// App代码
	AppID TThostFtdcAppIDType
	// App类型
	AppType TThostFtdcAppTypeType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcRspUserLogin2Field struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 登录成功时间
	LoginTime TThostFtdcTimeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 交易系统名称
	SystemName TThostFtdcSystemNameType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 最大报单引用
	MaxOrderRef TThostFtdcOrderRefType
	// 上期所时间
	SHFETime TThostFtdcTimeType
	// 大商所时间
	DCETime TThostFtdcTimeType
	// 郑商所时间
	CZCETime TThostFtdcTimeType
	// 中金所时间
	FFEXTime TThostFtdcTimeType
	// 能源中心时间
	INETime TThostFtdcTimeType
	// 随机串
	RandomString TThostFtdcRandomStringType
}

// 信息分发
type CThostFtdcTransferHeaderField struct {
	// 版本号，常量，1.0
	Version TThostFtdcVersionType
	// 交易代码，必填
	TradeCode TThostFtdcTradeCodeType
	// 交易日期，必填，格式：yyyymmdd
	TradeDate TThostFtdcTradeDateType
	// 交易时间，必填，格式：hhmmss
	TradeTime TThostFtdcTradeTimeType
	// 发起方流水号，N/A
	TradeSerial TThostFtdcTradeSerialType
	// 期货公司代码，必填
	FutureID TThostFtdcFutureIDType
	// 银行代码，根据查询银行得到，必填
	BankID TThostFtdcBankIDType
	// 银行分中心代码，根据查询银行得到，必填
	BankBrchID TThostFtdcBankBrchIDType
	// 操作员，N/A
	OperNo TThostFtdcOperNoType
	// 交易设备类型，N/A
	DeviceID TThostFtdcDeviceIDType
	// 记录数，N/A
	RecordNum TThostFtdcRecordNumType
	// 会话编号，N/A
	SessionID TThostFtdcSessionIDType
	// 请求编号，N/A
	RequestID TThostFtdcRequestIDType
}

// 信息分发
type CThostFtdcTransferBankToFutureReqField struct {
	// 期货资金账户
	FutureAccount TThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag TThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd TThostFtdcFutureAccPwdType
	// 转账金额
	TradeAmt TThostFtdcMoneyType
	// 客户手续费
	CustFee TThostFtdcMoneyType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode TThostFtdcCurrencyCodeType
}

// 信息分发
type CThostFtdcTransferBankToFutureRspField struct {
	// 响应代码
	RetCode TThostFtdcRetCodeType
	// 响应信息
	RetInfo TThostFtdcRetInfoType
	// 资金账户
	FutureAccount TThostFtdcAccountIDType
	// 转帐金额
	TradeAmt TThostFtdcMoneyType
	// 应收客户手续费
	CustFee TThostFtdcMoneyType
	// 币种
	CurrencyCode TThostFtdcCurrencyCodeType
}

// 信息分发
type CThostFtdcTransferFutureToBankReqField struct {
	// 期货资金账户
	FutureAccount TThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag TThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd TThostFtdcFutureAccPwdType
	// 转账金额
	TradeAmt TThostFtdcMoneyType
	// 客户手续费
	CustFee TThostFtdcMoneyType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode TThostFtdcCurrencyCodeType
}

// 信息分发
type CThostFtdcTransferFutureToBankRspField struct {
	// 响应代码
	RetCode TThostFtdcRetCodeType
	// 响应信息
	RetInfo TThostFtdcRetInfoType
	// 资金账户
	FutureAccount TThostFtdcAccountIDType
	// 转帐金额
	TradeAmt TThostFtdcMoneyType
	// 应收客户手续费
	CustFee TThostFtdcMoneyType
	// 币种
	CurrencyCode TThostFtdcCurrencyCodeType
}

// 信息分发
type CThostFtdcTransferQryBankReqField struct {
	// 期货资金账户
	FutureAccount TThostFtdcAccountIDType
	// 密码标志
	FuturePwdFlag TThostFtdcFuturePwdFlagType
	// 密码
	FutureAccPwd TThostFtdcFutureAccPwdType
	// 币种：RMB-人民币 USD-美圆 HKD-港元
	CurrencyCode TThostFtdcCurrencyCodeType
}

// 信息分发
type CThostFtdcTransferQryBankRspField struct {
	// 响应代码
	RetCode TThostFtdcRetCodeType
	// 响应信息
	RetInfo TThostFtdcRetInfoType
	// 资金账户
	FutureAccount TThostFtdcAccountIDType
	// 银行余额
	TradeAmt TThostFtdcMoneyType
	// 银行可用余额
	UseAmt TThostFtdcMoneyType
	// 银行可取余额
	FetchAmt TThostFtdcMoneyType
	// 币种
	CurrencyCode TThostFtdcCurrencyCodeType
}

// 信息分发
type CThostFtdcTransferQryDetailReqField struct {
	// 期货资金账户
	FutureAccount TThostFtdcAccountIDType
}

// 信息分发
type CThostFtdcTransferQryDetailRspField struct {
	// 交易日期
	TradeDate TThostFtdcDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 交易代码
	TradeCode TThostFtdcTradeCodeType
	// 期货流水号
	FutureSerial TThostFtdcTradeSerialNoType
	// 期货公司代码
	FutureID TThostFtdcFutureIDType
	// 资金帐号
	FutureAccount TThostFtdcFutureAccountType
	// 银行流水号
	BankSerial TThostFtdcTradeSerialNoType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行账号
	BankAccount TThostFtdcBankAccountType
	// 证件号码
	CertCode TThostFtdcCertCodeType
	// 货币代码
	CurrencyCode TThostFtdcCurrencyCodeType
	// 发生金额
	TxAmount TThostFtdcMoneyType
	// 有效标志
	Flag TThostFtdcTransferValidFlagType
}

// 信息分发
type CThostFtdcRspInfoField struct {
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcExchangeField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所名称
	ExchangeName TThostFtdcExchangeNameType
	// 交易所属性
	ExchangeProperty TThostFtdcExchangePropertyType
}

// 信息分发
type CThostFtdcProductField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 产品名称
	ProductName TThostFtdcProductNameType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品类型
	ProductClass TThostFtdcProductClassType
	// 合约数量乘数
	VolumeMultiple TThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick TThostFtdcPriceType
	// 市价单最大下单量
	MaxMarketOrderVolume TThostFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume TThostFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume TThostFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume TThostFtdcVolumeType
	// 持仓类型
	PositionType TThostFtdcPositionTypeType
	// 持仓日期类型
	PositionDateType TThostFtdcPositionDateTypeType
	// 平仓处理类型
	CloseDealType TThostFtdcCloseDealTypeType
	// 交易币种类型
	TradeCurrencyID TThostFtdcCurrencyIDType
	// 质押资金可用范围
	MortgageFundUseRange TThostFtdcMortgageFundUseRangeType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 合约基础商品乘数
	UnderlyingMultiple TThostFtdcUnderlyingMultipleType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 交易所产品代码
	ExchangeProductID TThostFtdcInstrumentIDType
	// 开仓量限制粒度
	OpenLimitControlLevel TThostFtdcOpenLimitControlLevelType
	// 报单频率控制粒度
	OrderFreqControlLevel TThostFtdcOrderFreqControlLevelType
}

// 信息分发
type CThostFtdcInstrumentField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约名称
	InstrumentName TThostFtdcInstrumentNameType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldInstrumentIDType
	// 产品类型
	ProductClass TThostFtdcProductClassType
	// 交割年份
	DeliveryYear TThostFtdcYearType
	// 交割月
	DeliveryMonth TThostFtdcMonthType
	// 市价单最大下单量
	MaxMarketOrderVolume TThostFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume TThostFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume TThostFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume TThostFtdcVolumeType
	// 合约数量乘数
	VolumeMultiple TThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick TThostFtdcPriceType
	// 创建日
	CreateDate TThostFtdcDateType
	// 上市日
	OpenDate TThostFtdcDateType
	// 到期日
	ExpireDate TThostFtdcDateType
	// 开始交割日
	StartDelivDate TThostFtdcDateType
	// 结束交割日
	EndDelivDate TThostFtdcDateType
	// 合约生命周期状态
	InstLifePhase TThostFtdcInstLifePhaseType
	// 当前是否交易
	IsTrading TThostFtdcBoolType
	// 持仓类型
	PositionType TThostFtdcPositionTypeType
	// 持仓日期类型
	PositionDateType TThostFtdcPositionDateTypeType
	// 多头保证金率
	LongMarginRatio TThostFtdcRatioType
	// 空头保证金率
	ShortMarginRatio TThostFtdcRatioType
	// 是否使用大额单边保证金算法
	MaxMarginSideAlgorithm TThostFtdcMaxMarginSideAlgorithmType
	// 保留的无效字段
	reserve4 TThostFtdcOldInstrumentIDType
	// 执行价
	StrikePrice TThostFtdcPriceType
	// 期权类型
	OptionsType TThostFtdcOptionsTypeType
	// 合约基础商品乘数
	UnderlyingMultiple TThostFtdcUnderlyingMultipleType
	// 组合类型
	CombinationType TThostFtdcCombinationTypeType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 基础商品代码
	UnderlyingInstrID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcBrokerField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司简称
	BrokerAbbr TThostFtdcBrokerAbbrType
	// 经纪公司名称
	BrokerName TThostFtdcBrokerNameType
	// 是否活跃
	IsActive TThostFtdcBoolType
}

// 信息分发
type CThostFtdcTraderField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 密码
	Password TThostFtdcPasswordType
	// 安装数量
	InstallCount TThostFtdcInstallCountType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 撤单时选择席位算法
	OrderCancelAlg TThostFtdcOrderCancelAlgType
}

// 信息分发
type CThostFtdcInvestorField struct {
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者名称
	InvestorName TThostFtdcPartyNameType
	// 证件类型
	IdentifiedCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 是否活跃
	IsActive TThostFtdcBoolType
	// 联系电话
	Telephone TThostFtdcTelephoneType
	// 通讯地址
	Address TThostFtdcAddressType
	// 开户日期
	OpenDate TThostFtdcDateType
	// 手机
	Mobile TThostFtdcMobileType
	// 手续费率模板代码
	CommModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 是否频率控制
	IsOrderFreq TThostFtdcEnumBoolType
	// 是否开仓限制
	IsOpenVolLimit TThostFtdcEnumBoolType
}

// 信息分发
type CThostFtdcTradingCodeField struct {
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 是否活跃
	IsActive TThostFtdcBoolType
	// 交易编码类型
	ClientIDType TThostFtdcClientIDTypeType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 业务类型
	BizType TThostFtdcBizTypeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcPartBrokerField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 是否活跃
	IsActive TThostFtdcBoolType
}

// 信息分发
type CThostFtdcSuperUserField struct {
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户名称
	UserName TThostFtdcUserNameType
	// 密码
	Password TThostFtdcPasswordType
	// 是否活跃
	IsActive TThostFtdcBoolType
}

// 信息分发
type CThostFtdcSuperUserFunctionField struct {
	// 用户代码
	UserID TThostFtdcUserIDType
	// 功能代码
	FunctionCode TThostFtdcFunctionCodeType
}

// 信息分发
type CThostFtdcInvestorGroupField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组名称
	InvestorGroupName TThostFtdcInvestorGroupNameType
}

// 信息分发
type CThostFtdcTradingAccountField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage TThostFtdcMoneyType
	// 上次信用额度
	PreCredit TThostFtdcMoneyType
	// 上次存款额
	PreDeposit TThostFtdcMoneyType
	// 上次结算准备金
	PreBalance TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin TThostFtdcMoneyType
	// 利息基数
	InterestBase TThostFtdcMoneyType
	// 利息收入
	Interest TThostFtdcMoneyType
	// 入金金额
	Deposit TThostFtdcMoneyType
	// 出金金额
	Withdraw TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 期货结算准备金
	Balance TThostFtdcMoneyType
	// 可用资金
	Available TThostFtdcMoneyType
	// 可取资金
	WithdrawQuota TThostFtdcMoneyType
	// 基本准备金
	Reserve TThostFtdcMoneyType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 信用额度
	Credit TThostFtdcMoneyType
	// 质押金额
	Mortgage TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin TThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin TThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance TThostFtdcMoneyType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn TThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut TThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable TThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin TThostFtdcMoneyType
	// 业务类型
	BizType TThostFtdcBizTypeType
	// 延时换汇冻结金额
	FrozenSwap TThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap TThostFtdcMoneyType
}

// 信息分发
type CThostFtdcInvestorPositionField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection TThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate TThostFtdcPositionDateType
	// 上日持仓
	YdPosition TThostFtdcVolumeType
	// 今日持仓
	Position TThostFtdcVolumeType
	// 多头冻结
	LongFrozen TThostFtdcVolumeType
	// 空头冻结
	ShortFrozen TThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType
	// 开仓量
	OpenVolume TThostFtdcVolumeType
	// 平仓量
	CloseVolume TThostFtdcVolumeType
	// 开仓金额
	OpenAmount TThostFtdcMoneyType
	// 平仓金额
	CloseAmount TThostFtdcMoneyType
	// 持仓成本
	PositionCost TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin TThostFtdcMoneyType
	// 占用的保证金
	UseMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 开仓成本
	OpenCost TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin TThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition TThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen TThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen TThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 今日持仓
	TodayPosition TThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 执行冻结
	StrikeFrozen TThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount TThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen TThostFtdcVolumeType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 持仓成本差值
	PositionCostOffset TThostFtdcMoneyType
	// tas持仓手数
	TasPosition TThostFtdcVolumeType
	// tas持仓成本
	TasPositionCost TThostFtdcMoneyType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInstrumentMarginRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative TThostFtdcBoolType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInstrumentCommissionRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 业务类型
	BizType TThostFtdcBizTypeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcDepthMarketDataField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 最新价
	LastPrice TThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 昨收盘
	PreClosePrice TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest TThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice TThostFtdcPriceType
	// 最高价
	HighestPrice TThostFtdcPriceType
	// 最低价
	LowestPrice TThostFtdcPriceType
	// 数量
	Volume TThostFtdcVolumeType
	// 成交金额
	Turnover TThostFtdcMoneyType
	// 持仓量
	OpenInterest TThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice TThostFtdcPriceType
	// 昨虚实度
	PreDelta TThostFtdcRatioType
	// 今虚实度
	CurrDelta TThostFtdcRatioType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TThostFtdcMillisecType
	// 申买价一
	BidPrice1 TThostFtdcPriceType
	// 申买量一
	BidVolume1 TThostFtdcVolumeType
	// 申卖价一
	AskPrice1 TThostFtdcPriceType
	// 申卖量一
	AskVolume1 TThostFtdcVolumeType
	// 申买价二
	BidPrice2 TThostFtdcPriceType
	// 申买量二
	BidVolume2 TThostFtdcVolumeType
	// 申卖价二
	AskPrice2 TThostFtdcPriceType
	// 申卖量二
	AskVolume2 TThostFtdcVolumeType
	// 申买价三
	BidPrice3 TThostFtdcPriceType
	// 申买量三
	BidVolume3 TThostFtdcVolumeType
	// 申卖价三
	AskPrice3 TThostFtdcPriceType
	// 申卖量三
	AskVolume3 TThostFtdcVolumeType
	// 申买价四
	BidPrice4 TThostFtdcPriceType
	// 申买量四
	BidVolume4 TThostFtdcVolumeType
	// 申卖价四
	AskPrice4 TThostFtdcPriceType
	// 申卖量四
	AskVolume4 TThostFtdcVolumeType
	// 申买价五
	BidPrice5 TThostFtdcPriceType
	// 申买量五
	BidVolume5 TThostFtdcVolumeType
	// 申卖价五
	AskPrice5 TThostFtdcPriceType
	// 申卖量五
	AskVolume5 TThostFtdcVolumeType
	// 当日均价
	AveragePrice TThostFtdcPriceType
	// 业务日期
	ActionDay TThostFtdcDateType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 上带价
	BandingUpperPrice TThostFtdcPriceType
	// 下带价
	BandingLowerPrice TThostFtdcPriceType
}

// 信息分发
type CThostFtdcInstrumentTradingRightField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易权限
	TradingRight TThostFtdcTradingRightType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcBrokerUserField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户名称
	UserName TThostFtdcUserNameType
	// 用户类型
	UserType TThostFtdcUserTypeType
	// 是否活跃
	IsActive TThostFtdcBoolType
	// 是否使用令牌
	IsUsingOTP TThostFtdcBoolType
	// 是否强制终端认证
	IsAuthForce TThostFtdcBoolType
}

// 信息分发
type CThostFtdcBrokerUserPasswordField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 密码
	Password TThostFtdcPasswordType
	// 上次修改时间
	LastUpdateTime TThostFtdcDateTimeType
	// 上次登陆时间
	LastLoginTime TThostFtdcDateTimeType
	// 密码过期时间
	ExpireDate TThostFtdcDateType
	// 弱密码过期时间
	WeakExpireDate TThostFtdcDateType
}

// 信息分发
type CThostFtdcBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 经纪公司功能代码
	BrokerFunctionCode TThostFtdcBrokerFunctionCodeType
}

// 信息分发
type CThostFtdcTraderOfferField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 交易所交易员连接状态
	TraderConnectStatus TThostFtdcTraderConnectStatusType
	// 发出连接请求的日期
	ConnectRequestDate TThostFtdcDateType
	// 发出连接请求的时间
	ConnectRequestTime TThostFtdcTimeType
	// 上次报告日期
	LastReportDate TThostFtdcDateType
	// 上次报告时间
	LastReportTime TThostFtdcTimeType
	// 完成连接日期
	ConnectDate TThostFtdcDateType
	// 完成连接时间
	ConnectTime TThostFtdcTimeType
	// 启动日期
	StartDate TThostFtdcDateType
	// 启动时间
	StartTime TThostFtdcTimeType
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 本席位最大成交编号
	MaxTradeID TThostFtdcTradeIDType
	// 本席位最大报单备拷
	MaxOrderMessageReference TThostFtdcReturnCodeType
	// 撤单时选择席位算法
	OrderCancelAlg TThostFtdcOrderCancelAlgType
}

// 信息分发
type CThostFtdcSettlementInfoField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 消息正文
	Content TThostFtdcContentType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcInstrumentMarginRateAdjustField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative TThostFtdcBoolType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 跟随交易所投资者多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 跟随交易所投资者多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 跟随交易所投资者空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 跟随交易所投资者空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 交易所多头保证金率
	ExchLongMarginRatioByMoney TThostFtdcRatioType
	// 交易所多头保证金费
	ExchLongMarginRatioByVolume TThostFtdcMoneyType
	// 交易所空头保证金率
	ExchShortMarginRatioByMoney TThostFtdcRatioType
	// 交易所空头保证金费
	ExchShortMarginRatioByVolume TThostFtdcMoneyType
	// 不跟随交易所投资者多头保证金率
	NoLongMarginRatioByMoney TThostFtdcRatioType
	// 不跟随交易所投资者多头保证金费
	NoLongMarginRatioByVolume TThostFtdcMoneyType
	// 不跟随交易所投资者空头保证金率
	NoShortMarginRatioByMoney TThostFtdcRatioType
	// 不跟随交易所投资者空头保证金费
	NoShortMarginRatioByVolume TThostFtdcMoneyType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 源币种
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源币种单位数量
	FromCurrencyUnit TThostFtdcCurrencyUnitType
	// 目标币种
	ToCurrencyID TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate TThostFtdcExchangeRateType
}

// 信息分发
type CThostFtdcSettlementRefField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
}

// 信息分发
type CThostFtdcCurrentTimeField struct {
	// 当前交易日
	CurrDate TThostFtdcDateType
	// 当前时间
	CurrTime TThostFtdcTimeType
	// 当前时间（毫秒）
	CurrMillisec TThostFtdcMillisecType
	// 自然日期
	ActionDay TThostFtdcDateType
}

// 信息分发
type CThostFtdcCommPhaseField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 通讯时段编号
	CommPhaseNo TThostFtdcCommPhaseNoType
	// 系统编号
	SystemID TThostFtdcSystemIDType
}

// 信息分发
type CThostFtdcLoginInfoField struct {
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 登录日期
	LoginDate TThostFtdcDateType
	// 登录时间
	LoginTime TThostFtdcTimeType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// 系统名称
	SystemName TThostFtdcSystemNameType
	// 密码,已弃用
	PasswordDeprecated TThostFtdcPasswordType
	// 最大报单引用
	MaxOrderRef TThostFtdcOrderRefType
	// 上期所时间
	SHFETime TThostFtdcTimeType
	// 大商所时间
	DCETime TThostFtdcTimeType
	// 郑商所时间
	CZCETime TThostFtdcTimeType
	// 中金所时间
	FFEXTime TThostFtdcTimeType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 动态密码
	OneTimePassword TThostFtdcPasswordType
	// 能源中心时间
	INETime TThostFtdcTimeType
	// 查询时是否需要流控
	IsQryControl TThostFtdcBoolType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// 密码
	Password TThostFtdcPasswordType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcLogoutAllField struct {
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 系统名称
	SystemName TThostFtdcSystemNameType
}

// 信息分发
type CThostFtdcFrontStatusField struct {
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 上次报告日期
	LastReportDate TThostFtdcDateType
	// 上次报告时间
	LastReportTime TThostFtdcTimeType
	// 是否活跃
	IsActive TThostFtdcBoolType
}

// 信息分发
type CThostFtdcUserPasswordUpdateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 原来的口令
	OldPassword TThostFtdcPasswordType
	// 新的口令
	NewPassword TThostFtdcPasswordType
}

// 信息分发
type CThostFtdcInputOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal TThostFtdcVolumeType
	// 有效期类型
	TimeCondition TThostFtdcTimeConditionType
	// GTD日期
	GTDDate TThostFtdcDateType
	// 成交量类型
	VolumeCondition TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume TThostFtdcVolumeType
	// 触发条件
	ContingentCondition TThostFtdcContingentConditionType
	// 止损价
	StopPrice TThostFtdcPriceType
	// 强平原因
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TThostFtdcBoolType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose TThostFtdcBoolType
	// 互换单标志
	IsSwapOrder TThostFtdcBoolType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal TThostFtdcVolumeType
	// 有效期类型
	TimeCondition TThostFtdcTimeConditionType
	// GTD日期
	GTDDate TThostFtdcDateType
	// 成交量类型
	VolumeCondition TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume TThostFtdcVolumeType
	// 触发条件
	ContingentCondition TThostFtdcContingentConditionType
	// 止损价
	StopPrice TThostFtdcPriceType
	// 强平原因
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TThostFtdcBoolType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 报单来源
	OrderSource TThostFtdcOrderSourceType
	// 报单状态
	OrderStatus TThostFtdcOrderStatusType
	// 报单类型
	OrderType TThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded TThostFtdcVolumeType
	// 剩余数量
	VolumeTotal TThostFtdcVolumeType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 委托时间
	InsertTime TThostFtdcTimeType
	// 激活时间
	ActiveTime TThostFtdcTimeType
	// 挂起时间
	SuspendTime TThostFtdcTimeType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID TThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 用户强评标志
	UserForceClose TThostFtdcBoolType
	// 操作用户代码
	ActiveUserID TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq TThostFtdcSequenceNoType
	// 相关报单
	RelativeOrderSysID TThostFtdcOrderSysIDType
	// 郑商所成交数量
	ZCETotalTradedVolume TThostFtdcVolumeType
	// 互换单标志
	IsSwapOrder TThostFtdcBoolType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExchangeOrderField struct {
	// 报单价格条件
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal TThostFtdcVolumeType
	// 有效期类型
	TimeCondition TThostFtdcTimeConditionType
	// GTD日期
	GTDDate TThostFtdcDateType
	// 成交量类型
	VolumeCondition TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume TThostFtdcVolumeType
	// 触发条件
	ContingentCondition TThostFtdcContingentConditionType
	// 止损价
	StopPrice TThostFtdcPriceType
	// 强平原因
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TThostFtdcBoolType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 报单来源
	OrderSource TThostFtdcOrderSourceType
	// 报单状态
	OrderStatus TThostFtdcOrderStatusType
	// 报单类型
	OrderType TThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded TThostFtdcVolumeType
	// 剩余数量
	VolumeTotal TThostFtdcVolumeType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 委托时间
	InsertTime TThostFtdcTimeType
	// 激活时间
	ActiveTime TThostFtdcTimeType
	// 挂起时间
	SuspendTime TThostFtdcTimeType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID TThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExchangeOrderInsertErrorField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcInputOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量变化
	VolumeChange TThostFtdcVolumeType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量变化
	VolumeChange TThostFtdcVolumeType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExchangeOrderActionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量变化
	VolumeChange TThostFtdcVolumeType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExchangeOrderActionErrorField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcExchangeTradeField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 成交编号
	TradeID TThostFtdcTradeIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 交易角色
	TradingRole TThostFtdcTradingRoleType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 价格
	Price TThostFtdcPriceType
	// 数量
	Volume TThostFtdcVolumeType
	// 成交时期
	TradeDate TThostFtdcDateType
	// 成交时间
	TradeTime TThostFtdcTimeType
	// 成交类型
	TradeType TThostFtdcTradeTypeType
	// 成交价来源
	PriceSource TThostFtdcPriceSourceType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 成交来源
	TradeSource TThostFtdcTradeSourceType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcTradeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 成交编号
	TradeID TThostFtdcTradeIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 交易角色
	TradingRole TThostFtdcTradingRoleType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 价格
	Price TThostFtdcPriceType
	// 数量
	Volume TThostFtdcVolumeType
	// 成交时期
	TradeDate TThostFtdcDateType
	// 成交时间
	TradeTime TThostFtdcTimeType
	// 成交类型
	TradeType TThostFtdcTradeTypeType
	// 成交价来源
	PriceSource TThostFtdcPriceSourceType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 经纪公司报单编号
	BrokerOrderSeq TThostFtdcSequenceNoType
	// 成交来源
	TradeSource TThostFtdcTradeSourceType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcUserSessionField struct {
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 登录日期
	LoginDate TThostFtdcDateType
	// 登录时间
	LoginTime TThostFtdcTimeType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryMaxOrderVolumeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume TThostFtdcVolumeType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 确认日期
	ConfirmDate TThostFtdcDateType
	// 确认时间
	ConfirmTime TThostFtdcTimeType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcSyncDepositField struct {
	// 出入金流水号
	DepositSeqNo TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 入金金额
	Deposit TThostFtdcMoneyType
	// 是否强制进行
	IsForce TThostFtdcBoolType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 是否是个股期权内转
	IsFromSopt TThostFtdcBoolType
	// 资金密码
	TradingPassword TThostFtdcPasswordType
}

// 信息分发
type CThostFtdcSyncFundMortgageField struct {
	// 货币质押流水号
	MortgageSeqNo TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID TThostFtdcCurrencyIDType
	// 质押金额
	MortgageAmount TThostFtdcMoneyType
	// 目标币种
	ToCurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcBrokerSyncField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcSyncingInvestorField struct {
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者名称
	InvestorName TThostFtdcPartyNameType
	// 证件类型
	IdentifiedCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 是否活跃
	IsActive TThostFtdcBoolType
	// 联系电话
	Telephone TThostFtdcTelephoneType
	// 通讯地址
	Address TThostFtdcAddressType
	// 开户日期
	OpenDate TThostFtdcDateType
	// 手机
	Mobile TThostFtdcMobileType
	// 手续费率模板代码
	CommModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 是否频率控制
	IsOrderFreq TThostFtdcEnumBoolType
	// 是否开仓限制
	IsOpenVolLimit TThostFtdcEnumBoolType
}

// 信息分发
type CThostFtdcSyncingTradingCodeField struct {
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 是否活跃
	IsActive TThostFtdcBoolType
	// 交易编码类型
	ClientIDType TThostFtdcClientIDTypeType
}

// 信息分发
type CThostFtdcSyncingInvestorGroupField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者分组代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 投资者分组名称
	InvestorGroupName TThostFtdcInvestorGroupNameType
}

// 信息分发
type CThostFtdcSyncingTradingAccountField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage TThostFtdcMoneyType
	// 上次信用额度
	PreCredit TThostFtdcMoneyType
	// 上次存款额
	PreDeposit TThostFtdcMoneyType
	// 上次结算准备金
	PreBalance TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin TThostFtdcMoneyType
	// 利息基数
	InterestBase TThostFtdcMoneyType
	// 利息收入
	Interest TThostFtdcMoneyType
	// 入金金额
	Deposit TThostFtdcMoneyType
	// 出金金额
	Withdraw TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 期货结算准备金
	Balance TThostFtdcMoneyType
	// 可用资金
	Available TThostFtdcMoneyType
	// 可取资金
	WithdrawQuota TThostFtdcMoneyType
	// 基本准备金
	Reserve TThostFtdcMoneyType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 信用额度
	Credit TThostFtdcMoneyType
	// 质押金额
	Mortgage TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin TThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin TThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance TThostFtdcMoneyType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn TThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut TThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable TThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin TThostFtdcMoneyType
	// 延时换汇冻结金额
	FrozenSwap TThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap TThostFtdcMoneyType
}

// 信息分发
type CThostFtdcSyncingInvestorPositionField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection TThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate TThostFtdcPositionDateType
	// 上日持仓
	YdPosition TThostFtdcVolumeType
	// 今日持仓
	Position TThostFtdcVolumeType
	// 多头冻结
	LongFrozen TThostFtdcVolumeType
	// 空头冻结
	ShortFrozen TThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType
	// 开仓量
	OpenVolume TThostFtdcVolumeType
	// 平仓量
	CloseVolume TThostFtdcVolumeType
	// 开仓金额
	OpenAmount TThostFtdcMoneyType
	// 平仓金额
	CloseAmount TThostFtdcMoneyType
	// 持仓成本
	PositionCost TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin TThostFtdcMoneyType
	// 占用的保证金
	UseMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 开仓成本
	OpenCost TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin TThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition TThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen TThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen TThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 今日持仓
	TodayPosition TThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 执行冻结
	StrikeFrozen TThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount TThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen TThostFtdcVolumeType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 持仓成本差值
	PositionCostOffset TThostFtdcMoneyType
	// tas持仓手数
	TasPosition TThostFtdcVolumeType
	// tas持仓成本
	TasPositionCost TThostFtdcMoneyType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSyncingInstrumentMarginRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative TThostFtdcBoolType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSyncingInstrumentCommissionRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSyncingInstrumentTradingRightField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易权限
	TradingRight TThostFtdcTradingRightType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryTradeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 成交编号
	TradeID TThostFtdcTradeIDType
	// 开始时间
	TradeTimeStart TThostFtdcTimeType
	// 结束时间
	TradeTimeEnd TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryInvestorPositionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryTradingAccountField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 业务类型
	BizType TThostFtdcBizTypeType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
}

// 信息分发
type CThostFtdcQryInvestorField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcQryTradingCodeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 交易编码类型
	ClientIDType TThostFtdcClientIDTypeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcQryInvestorGroupField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcQryInstrumentMarginRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryInstrumentTradingRightField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryBrokerField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcQryTraderField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
}

// 信息分发
type CThostFtdcQrySuperUserFunctionField struct {
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcQryUserSessionField struct {
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcQryPartBrokerField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
}

// 信息分发
type CThostFtdcQryFrontStatusField struct {
	// 前置编号
	FrontID TThostFtdcFrontIDType
}

// 信息分发
type CThostFtdcQryExchangeOrderField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcQryOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcQryExchangeOrderActionField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
}

// 信息分发
type CThostFtdcQrySuperUserField struct {
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcQryExchangeField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcQryProductField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 产品类型
	ProductClass TThostFtdcProductClassType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryInstrumentField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryDepthMarketDataField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryBrokerUserField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcQryBrokerUserFunctionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcQryTraderOfferField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
}

// 信息分发
type CThostFtdcQrySyncDepositField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 出入金流水号
	DepositSeqNo TThostFtdcDepositSeqNoType
}

// 信息分发
type CThostFtdcQrySettlementInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易日
	TradingDay TThostFtdcDateType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcQryExchangeMarginRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryExchangeMarginRateAdjustField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryExchangeRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 源币种
	FromCurrencyID TThostFtdcCurrencyIDType
	// 目标币种
	ToCurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcQrySyncFundMortgageField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 货币质押流水号
	MortgageSeqNo TThostFtdcDepositSeqNoType
}

// 信息分发
type CThostFtdcQryHisOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd TThostFtdcTimeType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcOptionInstrMiniMarginField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 单位（手）期权合约最小保证金
	MinMargin TThostFtdcMoneyType
	// 取值方式
	ValueMethod TThostFtdcValueMethodType
	// 是否跟随交易所收取
	IsRelative TThostFtdcBoolType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcOptionInstrMarginAdjustField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机空头保证金调整系数
	SShortMarginRatioByMoney TThostFtdcRatioType
	// 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType
	// 保值空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType
	// 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType
	// 套利空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType
	// 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType
	// 是否跟随交易所收取
	IsRelative TThostFtdcBoolType
	// 做市商空头保证金调整系数
	MShortMarginRatioByMoney TThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcOptionInstrCommRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney TThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume TThostFtdcRatioType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 期权合约保证金不变部分
	FixedMargin TThostFtdcMoneyType
	// 期权合约最小保证金
	MiniMargin TThostFtdcMoneyType
	// 期权合约权利金
	Royalty TThostFtdcMoneyType
	// 交易所期权合约保证金不变部分
	ExchFixedMargin TThostFtdcMoneyType
	// 交易所期权合约最小保证金
	ExchMiniMargin TThostFtdcMoneyType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryOptionInstrTradeCostField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 期权合约报价
	InputPrice TThostFtdcPriceType
	// 标的价格,填0则用昨结算价
	UnderlyingPrice TThostFtdcPriceType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcIndexPriceField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 指数现货收盘价
	ClosePrice TThostFtdcPriceType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInputExecOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 执行宣告引用
	ExecOrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 执行类型
	ActionType TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcInputExecOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef TThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExecOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 执行宣告引用
	ExecOrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 执行类型
	ActionType TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 本地执行宣告编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 执行宣告提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 执行宣告编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 执行结果
	ExecResult TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerExecOrderSeq TThostFtdcSequenceNoType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExecOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef TThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地执行宣告编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 执行类型
	ActionType TThostFtdcActionTypeType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExecOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行宣告编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 开始时间
	InsertTimeStart TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd TThostFtdcTimeType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeExecOrderField struct {
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 执行类型
	ActionType TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 本地执行宣告编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 执行宣告提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 执行宣告编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 执行结果
	ExecResult TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExchangeExecOrderField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcQryExecOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcExchangeExecOrderActionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地执行宣告编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 执行类型
	ActionType TThostFtdcActionTypeType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 数量
	Volume TThostFtdcVolumeType
	// IP地址
	IPAddress TThostFtdcIPAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcQryExchangeExecOrderActionField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
}

// 信息分发
type CThostFtdcErrExecOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 执行宣告引用
	ExecOrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 执行类型
	ActionType TThostFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TThostFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记,该字段已废弃
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TThostFtdcExecOrderCloseFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryErrExecOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 执行宣告操作引用
	ExecOrderActionRef TThostFtdcOrderActionRefType
	// 执行宣告引用
	ExecOrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行宣告操作编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryErrExecOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcOptionInstrTradingRightField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 交易权限
	TradingRight TThostFtdcTradingRightType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryOptionInstrTradingRightField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInputForQuoteField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 询价引用
	ForQuoteRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcForQuoteField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 询价引用
	ForQuoteRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 本地询价编号
	ForQuoteLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 询价状态
	ForQuoteStatus TThostFtdcForQuoteStatusType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID TThostFtdcUserIDType
	// 经纪公司询价编号
	BrokerForQutoSeq TThostFtdcSequenceNoType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryForQuoteField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 开始时间
	InsertTimeStart TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeForQuoteField struct {
	// 本地询价编号
	ForQuoteLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 询价状态
	ForQuoteStatus TThostFtdcForQuoteStatusType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExchangeForQuoteField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcInputQuoteField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报价引用
	QuoteRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 卖价格
	AskPrice TThostFtdcPriceType
	// 买价格
	BidPrice TThostFtdcPriceType
	// 卖数量
	AskVolume TThostFtdcVolumeType
	// 买数量
	BidVolume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag TThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType
	// 衍生卖报单引用
	AskOrderRef TThostFtdcOrderRefType
	// 衍生买报单引用
	BidOrderRef TThostFtdcOrderRefType
	// 应价编号
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
	// 被顶单编号
	ReplaceSysID TThostFtdcOrderSysIDType
}

// 信息分发
type CThostFtdcInputQuoteActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报价操作引用
	QuoteActionRef TThostFtdcOrderActionRefType
	// 报价引用
	QuoteRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQuoteField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报价引用
	QuoteRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 卖价格
	AskPrice TThostFtdcPriceType
	// 买价格
	BidPrice TThostFtdcPriceType
	// 卖数量
	AskVolume TThostFtdcVolumeType
	// 买数量
	BidVolume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag TThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType
	// 本地报价编号
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报价提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 报价提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 报价编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 报价状态
	QuoteStatus TThostFtdcOrderStatusType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 卖方报单编号
	AskOrderSysID TThostFtdcOrderSysIDType
	// 买方报单编号
	BidOrderSysID TThostFtdcOrderSysIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID TThostFtdcUserIDType
	// 经纪公司报价编号
	BrokerQuoteSeq TThostFtdcSequenceNoType
	// 衍生卖报单引用
	AskOrderRef TThostFtdcOrderRefType
	// 衍生买报单引用
	BidOrderRef TThostFtdcOrderRefType
	// 应价编号
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
	// 被顶单编号
	ReplaceSysID TThostFtdcOrderSysIDType
}

// 信息分发
type CThostFtdcQuoteActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报价操作引用
	QuoteActionRef TThostFtdcOrderActionRefType
	// 报价引用
	QuoteRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报价编号
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryQuoteField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报价编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd TThostFtdcTimeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeQuoteField struct {
	// 卖价格
	AskPrice TThostFtdcPriceType
	// 买价格
	BidPrice TThostFtdcPriceType
	// 卖数量
	AskVolume TThostFtdcVolumeType
	// 买数量
	BidVolume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 卖开平标志
	AskOffsetFlag TThostFtdcOffsetFlagType
	// 买开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TThostFtdcHedgeFlagType
	// 买投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType
	// 本地报价编号
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报价提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 报价提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 报价编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 报价状态
	QuoteStatus TThostFtdcOrderStatusType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 卖方报单编号
	AskOrderSysID TThostFtdcOrderSysIDType
	// 买方报单编号
	BidOrderSysID TThostFtdcOrderSysIDType
	// 应价编号
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExchangeQuoteField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcQryQuoteActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcExchangeQuoteActionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报价操作编号
	QuoteSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报价编号
	QuoteLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExchangeQuoteActionField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
}

// 信息分发
type CThostFtdcOptionInstrDeltaField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// Delta值
	Delta TThostFtdcRatioType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcForQuoteRspField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 询价编号
	ForQuoteSysID TThostFtdcOrderSysIDType
	// 询价时间
	ForQuoteTime TThostFtdcTimeType
	// 业务日期
	ActionDay TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcStrikeOffsetField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 执行偏移值
	Offset TThostFtdcMoneyType
	// 执行偏移类型
	OffsetType TThostFtdcStrikeOffsetTypeType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryStrikeOffsetField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInputBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef TThostFtdcOrderActionRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef TThostFtdcOrderActionRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcExchangeBatchOrderActionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryBatchOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	//
	GuarantRatio TThostFtdcRatioType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryCombInstrumentGuardField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInputCombActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 组合引用
	CombActionRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 数量
	Volume TThostFtdcVolumeType
	// 组合指令方向
	CombDirection TThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcCombActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 组合引用
	CombActionRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 数量
	Volume TThostFtdcVolumeType
	// 组合指令方向
	CombDirection TThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 本地申请组合编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 组合状态
	ActionStatus TThostFtdcOrderActionStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 组合编号
	ComTradeID TThostFtdcTradeIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryCombActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeCombActionField struct {
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 数量
	Volume TThostFtdcVolumeType
	// 组合指令方向
	CombDirection TThostFtdcCombDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 本地申请组合编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 组合状态
	ActionStatus TThostFtdcOrderActionStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 组合编号
	ComTradeID TThostFtdcTradeIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExchangeCombActionField struct {
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcProductExchRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报价币种类型
	QuoteCurrencyID TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate TThostFtdcExchangeRateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryProductExchRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryForQuoteParamField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcForQuoteParamField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 最新价
	LastPrice TThostFtdcPriceType
	// 价差
	PriceInterval TThostFtdcPriceType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcMMOptionInstrCommRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney TThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume TThostFtdcRatioType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryMMOptionInstrCommRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcMMInstrumentCommissionRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryMMInstrumentCommissionRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInstrumentOrderCommRateField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 报单手续费
	OrderCommByVolume TThostFtdcRatioType
	// 撤单手续费
	OrderActionCommByVolume TThostFtdcRatioType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 报单手续费
	OrderCommByTrade TThostFtdcRatioType
	// 撤单手续费
	OrderActionCommByTrade TThostFtdcRatioType
}

// 信息分发
type CThostFtdcQryInstrumentOrderCommRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcTradeParamField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 参数代码
	TradeParamID TThostFtdcTradeParamIDType
	// 参数代码值
	TradeParamValue TThostFtdcSettlementParamValueType
	// 备注
	Memo TThostFtdcMemoType
}

// 信息分发
type CThostFtdcInstrumentMarginRateULField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcFutureLimitPosiParamField struct {
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 当日投机开仓数量限制
	SpecOpenVolume TThostFtdcVolumeType
	// 当日套利开仓数量限制
	ArbiOpenVolume TThostFtdcVolumeType
	// 当日投机+套利开仓数量限制
	OpenVolume TThostFtdcVolumeType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcLoginForbiddenIPField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcIPListField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 是否白名单
	IsWhite TThostFtdcBoolType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcInputOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcInputOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef TThostFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 期权自对冲编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 自对冲结果
	ExecResult TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 操作用户代码
	ActiveUserID TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOptionSelfCloseSeq TThostFtdcSequenceNoType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef TThostFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryOptionSelfCloseField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 期权自对冲编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart TThostFtdcTimeType
	// 结束时间
	InsertTimeEnd TThostFtdcTimeType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcExchangeOptionSelfCloseField struct {
	// 数量
	Volume TThostFtdcVolumeType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 期权自对冲编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 插入时间
	InsertTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 自对冲结果
	ExecResult TThostFtdcExecResultType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryOptionSelfCloseActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcExchangeOptionSelfCloseActionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType
	// IP地址
	IPAddress TThostFtdcIPAddressType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcSyncDelaySwapField struct {
	// 换汇流水号
	DelaySwapSeqNo TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源金额
	FromAmount TThostFtdcMoneyType
	// 源换汇冻结金额(可用冻结)
	FromFrozenSwap TThostFtdcMoneyType
	// 源剩余换汇额度(可提冻结)
	FromRemainSwap TThostFtdcMoneyType
	// 目标币种
	ToCurrencyID TThostFtdcCurrencyIDType
	// 目标金额
	ToAmount TThostFtdcMoneyType
	// 是否手工换汇
	IsManualSwap TThostFtdcBoolType
	// 是否将所有外币的剩余换汇额度设置为0
	IsAllRemainSetZero TThostFtdcBoolType
}

// 信息分发
type CThostFtdcQrySyncDelaySwapField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 延时换汇流水号
	DelaySwapSeqNo TThostFtdcDepositSeqNoType
}

// 信息分发
type CThostFtdcInvestUnitField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 投资者单元名称
	InvestorUnitName TThostFtdcPartyNameType
	// 投资者分组代码
	InvestorGroupID TThostFtdcInvestorIDType
	// 手续费率模板代码
	CommModelID TThostFtdcInvestorIDType
	// 保证金率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcQryInvestUnitField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcSecAgentCheckModeField struct {
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 币种
	CurrencyID TThostFtdcCurrencyIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID TThostFtdcAccountIDType
	// 是否需要校验自己的资金账户
	CheckSelfAccount TThostFtdcBoolType
}

// 信息分发
type CThostFtdcSecAgentTradeInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID TThostFtdcAccountIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 二级代理商姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcMarketDataField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 最新价
	LastPrice TThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 昨收盘
	PreClosePrice TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest TThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice TThostFtdcPriceType
	// 最高价
	HighestPrice TThostFtdcPriceType
	// 最低价
	LowestPrice TThostFtdcPriceType
	// 数量
	Volume TThostFtdcVolumeType
	// 成交金额
	Turnover TThostFtdcMoneyType
	// 持仓量
	OpenInterest TThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice TThostFtdcPriceType
	// 昨虚实度
	PreDelta TThostFtdcRatioType
	// 今虚实度
	CurrDelta TThostFtdcRatioType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TThostFtdcMillisecType
	// 业务日期
	ActionDay TThostFtdcDateType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcMarketDataBaseField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 昨收盘
	PreClosePrice TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest TThostFtdcLargeVolumeType
	// 昨虚实度
	PreDelta TThostFtdcRatioType
}

// 信息分发
type CThostFtdcMarketDataStaticField struct {
	// 今开盘
	OpenPrice TThostFtdcPriceType
	// 最高价
	HighestPrice TThostFtdcPriceType
	// 最低价
	LowestPrice TThostFtdcPriceType
	// 今收盘
	ClosePrice TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 今虚实度
	CurrDelta TThostFtdcRatioType
}

// 信息分发
type CThostFtdcMarketDataLastMatchField struct {
	// 最新价
	LastPrice TThostFtdcPriceType
	// 数量
	Volume TThostFtdcVolumeType
	// 成交金额
	Turnover TThostFtdcMoneyType
	// 持仓量
	OpenInterest TThostFtdcLargeVolumeType
}

// 信息分发
type CThostFtdcMarketDataBestPriceField struct {
	// 申买价一
	BidPrice1 TThostFtdcPriceType
	// 申买量一
	BidVolume1 TThostFtdcVolumeType
	// 申卖价一
	AskPrice1 TThostFtdcPriceType
	// 申卖量一
	AskVolume1 TThostFtdcVolumeType
}

// 信息分发
type CThostFtdcMarketDataBid23Field struct {
	// 申买价二
	BidPrice2 TThostFtdcPriceType
	// 申买量二
	BidVolume2 TThostFtdcVolumeType
	// 申买价三
	BidPrice3 TThostFtdcPriceType
	// 申买量三
	BidVolume3 TThostFtdcVolumeType
}

// 信息分发
type CThostFtdcMarketDataAsk23Field struct {
	// 申卖价二
	AskPrice2 TThostFtdcPriceType
	// 申卖量二
	AskVolume2 TThostFtdcVolumeType
	// 申卖价三
	AskPrice3 TThostFtdcPriceType
	// 申卖量三
	AskVolume3 TThostFtdcVolumeType
}

// 信息分发
type CThostFtdcMarketDataBid45Field struct {
	// 申买价四
	BidPrice4 TThostFtdcPriceType
	// 申买量四
	BidVolume4 TThostFtdcVolumeType
	// 申买价五
	BidPrice5 TThostFtdcPriceType
	// 申买量五
	BidVolume5 TThostFtdcVolumeType
}

// 信息分发
type CThostFtdcMarketDataAsk45Field struct {
	// 申卖价四
	AskPrice4 TThostFtdcPriceType
	// 申卖量四
	AskVolume4 TThostFtdcVolumeType
	// 申卖价五
	AskPrice5 TThostFtdcPriceType
	// 申卖量五
	AskVolume5 TThostFtdcVolumeType
}

// 信息分发
type CThostFtdcMarketDataUpdateTimeField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TThostFtdcMillisecType
	// 业务日期
	ActionDay TThostFtdcDateType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcMarketDataBandingPriceField struct {
	// 上带价
	BandingUpperPrice TThostFtdcPriceType
	// 下带价
	BandingLowerPrice TThostFtdcPriceType
}

// 信息分发
type CThostFtdcMarketDataExchangeField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcSpecificInstrumentField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInstrumentStatusField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 结算组代码
	SettlementGroupID TThostFtdcSettlementGroupIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 合约交易状态
	InstrumentStatus TThostFtdcInstrumentStatusType
	// 交易阶段编号
	TradingSegmentSN TThostFtdcTradingSegmentSNType
	// 进入本状态时间
	EnterTime TThostFtdcTimeType
	// 进入本状态原因
	EnterReason TThostFtdcInstStatusEnterReasonType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryInstrumentStatusField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldExchangeInstIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
}

// 信息分发
type CThostFtdcInvestorAccountField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcPositionProfitAlgorithmField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 盈亏算法
	Algorithm TThostFtdcAlgorithmType
	// 备注
	Memo TThostFtdcMemoType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcDiscountField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 资金折扣比例
	Discount TThostFtdcRatioType
}

// 信息分发
type CThostFtdcQryTransferBankField struct {
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID TThostFtdcBankBrchIDType
}

// 信息分发
type CThostFtdcTransferBankField struct {
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行名称
	BankName TThostFtdcBankNameType
	// 是否活跃
	IsActive TThostFtdcBoolType
}

// 信息分发
type CThostFtdcQryInvestorPositionDetailField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInvestorPositionDetailField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 买卖
	Direction TThostFtdcDirectionType
	// 开仓日期
	OpenDate TThostFtdcDateType
	// 成交编号
	TradeID TThostFtdcTradeIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 开仓价
	OpenPrice TThostFtdcPriceType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 成交类型
	TradeType TThostFtdcTradeTypeType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 逐日盯市平仓盈亏
	CloseProfitByDate TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 逐日盯市持仓盈亏
	PositionProfitByDate TThostFtdcMoneyType
	// 逐笔对冲持仓盈亏
	PositionProfitByTrade TThostFtdcMoneyType
	// 投资者保证金
	Margin TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin TThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 昨结算价
	LastSettlementPrice TThostFtdcPriceType
	// 结算价
	SettlementPrice TThostFtdcPriceType
	// 平仓量
	CloseVolume TThostFtdcVolumeType
	// 平仓金额
	CloseAmount TThostFtdcMoneyType
	// 先开先平剩余数量
	TimeFirstVolume TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 特殊持仓标志
	SpecPosiType TThostFtdcSpecPosiTypeType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 组合合约代码
	CombInstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcTradingAccountPasswordField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 密码
	Password TThostFtdcPasswordType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcMDTraderOfferField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 交易所交易员连接状态
	TraderConnectStatus TThostFtdcTraderConnectStatusType
	// 发出连接请求的日期
	ConnectRequestDate TThostFtdcDateType
	// 发出连接请求的时间
	ConnectRequestTime TThostFtdcTimeType
	// 上次报告日期
	LastReportDate TThostFtdcDateType
	// 上次报告时间
	LastReportTime TThostFtdcTimeType
	// 完成连接日期
	ConnectDate TThostFtdcDateType
	// 完成连接时间
	ConnectTime TThostFtdcTimeType
	// 启动日期
	StartDate TThostFtdcDateType
	// 启动时间
	StartTime TThostFtdcTimeType
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 本席位最大成交编号
	MaxTradeID TThostFtdcTradeIDType
	// 本席位最大报单备拷
	MaxOrderMessageReference TThostFtdcReturnCodeType
	// 撤单时选择席位算法
	OrderCancelAlg TThostFtdcOrderCancelAlgType
}

// 信息分发
type CThostFtdcQryMDTraderOfferField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
}

// 信息分发
type CThostFtdcQryNoticeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcNoticeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 消息正文
	Content TThostFtdcContentType
	// 经纪公司通知内容序列号
	SequenceLabel TThostFtdcSequenceLabelType
}

// 信息分发
type CThostFtdcUserRightField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 客户权限类型
	UserRightType TThostFtdcUserRightTypeType
	// 是否禁止
	IsForbidden TThostFtdcBoolType
}

// 信息分发
type CThostFtdcQrySettlementInfoConfirmField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcLoadSettlementInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcBrokerWithdrawAlgorithmField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 可提资金算法
	WithdrawAlgorithm TThostFtdcAlgorithmType
	// 资金使用率
	UsingRatio TThostFtdcRatioType
	// 可提是否包含平仓盈利
	IncludeCloseProfit TThostFtdcIncludeCloseProfitType
	// 本日无仓且无成交客户是否受可提比例限制
	AllWithoutTrade TThostFtdcAllWithoutTradeType
	// 可用是否包含平仓盈利
	AvailIncludeCloseProfit TThostFtdcIncludeCloseProfitType
	// 是否启用用户事件
	IsBrokerUserEvent TThostFtdcBoolType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 货币质押比率
	FundMortgageRatio TThostFtdcRatioType
	// 权益算法
	BalanceAlgorithm TThostFtdcBalanceAlgorithmType
}

// 信息分发
type CThostFtdcTradingAccountPasswordUpdateV1Field struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 原来的口令
	OldPassword TThostFtdcPasswordType
	// 新的口令
	NewPassword TThostFtdcPasswordType
}

// 信息分发
type CThostFtdcTradingAccountPasswordUpdateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 原来的口令
	OldPassword TThostFtdcPasswordType
	// 新的口令
	NewPassword TThostFtdcPasswordType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcQryCombinationLegField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 单腿编号
	LegID TThostFtdcLegIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 组合合约代码
	CombInstrumentID TThostFtdcInstrumentIDType
	// 单腿合约代码
	LegInstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQrySyncStatusField struct {
	// 交易日
	TradingDay TThostFtdcDateType
}

// 信息分发
type CThostFtdcCombinationLegField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 单腿编号
	LegID TThostFtdcLegIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 单腿乘数
	LegMultiple TThostFtdcLegMultipleType
	// 派生层数
	ImplyLevel TThostFtdcImplyLevelType
	// 组合合约代码
	CombInstrumentID TThostFtdcInstrumentIDType
	// 单腿合约代码
	LegInstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSyncStatusField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 数据同步状态
	DataSyncStatus TThostFtdcDataSyncStatusType
}

// 信息分发
type CThostFtdcQryLinkManField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcLinkManField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 联系人类型
	PersonType TThostFtdcPersonTypeType
	// 证件类型
	IdentifiedCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 名称
	PersonName TThostFtdcPartyNameType
	// 联系电话
	Telephone TThostFtdcTelephoneType
	// 通讯地址
	Address TThostFtdcAddressType
	// 邮政编码
	ZipCode TThostFtdcZipCodeType
	// 优先级
	Priority TThostFtdcPriorityType
	// 开户邮政编码
	UOAZipCode TThostFtdcUOAZipCodeType
	// 全称
	PersonFullName TThostFtdcInvestorFullNameType
}

// 信息分发
type CThostFtdcQryBrokerUserEventField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户事件类型
	UserEventType TThostFtdcUserEventTypeType
}

// 信息分发
type CThostFtdcBrokerUserEventField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户事件类型
	UserEventType TThostFtdcUserEventTypeType
	// 用户事件序号
	EventSequenceNo TThostFtdcSequenceNoType
	// 事件发生日期
	EventDate TThostFtdcDateType
	// 事件发生时间
	EventTime TThostFtdcTimeType
	// 用户事件信息
	UserEventInfo TThostFtdcUserEventInfoType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryContractBankField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID TThostFtdcBankBrchIDType
}

// 信息分发
type CThostFtdcContractBankField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分中心代码
	BankBrchID TThostFtdcBankBrchIDType
	// 银行名称
	BankName TThostFtdcBankNameType
}

// 信息分发
type CThostFtdcInvestorPositionCombineDetailField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 开仓日期
	OpenDate TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 组合编号
	ComTradeID TThostFtdcTradeIDType
	// 撮合编号
	TradeID TThostFtdcTradeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 买卖
	Direction TThostFtdcDirectionType
	// 持仓量
	TotalAmt TThostFtdcVolumeType
	// 投资者保证金
	Margin TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin TThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 单腿编号
	LegID TThostFtdcLegIDType
	// 单腿乘数
	LegMultiple TThostFtdcLegMultipleType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 成交组号
	TradeGroupID TThostFtdcTradeGroupIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 组合持仓合约编码
	CombInstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcParkedOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal TThostFtdcVolumeType
	// 有效期类型
	TimeCondition TThostFtdcTimeConditionType
	// GTD日期
	GTDDate TThostFtdcDateType
	// 成交量类型
	VolumeCondition TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume TThostFtdcVolumeType
	// 触发条件
	ContingentCondition TThostFtdcContingentConditionType
	// 止损价
	StopPrice TThostFtdcPriceType
	// 强平原因
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TThostFtdcBoolType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose TThostFtdcBoolType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 预埋报单编号
	ParkedOrderID TThostFtdcParkedOrderIDType
	// 用户类型
	UserType TThostFtdcUserTypeType
	// 预埋单状态
	Status TThostFtdcParkedOrderStatusType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder TThostFtdcBoolType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量变化
	VolumeChange TThostFtdcVolumeType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 预埋撤单单编号
	ParkedOrderActionID TThostFtdcParkedOrderActionIDType
	// 用户类型
	UserType TThostFtdcUserTypeType
	// 预埋撤单状态
	Status TThostFtdcParkedOrderStatusType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryParkedOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcRemoveParkedOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 预埋报单编号
	ParkedOrderID TThostFtdcParkedOrderIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcRemoveParkedOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 预埋撤单编号
	ParkedOrderActionID TThostFtdcParkedOrderActionIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcInvestorWithdrawAlgorithmField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 可提资金比例
	UsingRatio TThostFtdcRatioType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 货币质押比率
	FundMortgageRatio TThostFtdcRatioType
}

// 信息分发
type CThostFtdcQryInvestorPositionCombineDetailField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 组合持仓合约编码
	CombInstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcMarketDataAveragePriceField struct {
	// 当日均价
	AveragePrice TThostFtdcPriceType
}

// 信息分发
type CThostFtdcVerifyInvestorPasswordField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 密码
	Password TThostFtdcPasswordType
}

// 信息分发
type CThostFtdcUserIPField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
	// IP地址掩码
	IPMask TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcTradingNoticeInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 发送时间
	SendTime TThostFtdcTimeType
	// 消息正文
	FieldContent TThostFtdcContentType
	// 序列系列号
	SequenceSeries TThostFtdcSequenceSeriesType
	// 序列号
	SequenceNo TThostFtdcSequenceNoType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcTradingNoticeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 序列系列号
	SequenceSeries TThostFtdcSequenceSeriesType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 发送时间
	SendTime TThostFtdcTimeType
	// 序列号
	SequenceNo TThostFtdcSequenceNoType
	// 消息正文
	FieldContent TThostFtdcContentType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcQryTradingNoticeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcQryErrOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcErrOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal TThostFtdcVolumeType
	// 有效期类型
	TimeCondition TThostFtdcTimeConditionType
	// GTD日期
	GTDDate TThostFtdcDateType
	// 成交量类型
	VolumeCondition TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume TThostFtdcVolumeType
	// 触发条件
	ContingentCondition TThostFtdcContingentConditionType
	// 止损价
	StopPrice TThostFtdcPriceType
	// 强平原因
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TThostFtdcBoolType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 用户强评标志
	UserForceClose TThostFtdcBoolType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder TThostFtdcBoolType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易编码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcErrorConditionalOrderField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 报单价格条件
	OrderPriceType TThostFtdcOrderPriceTypeType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 组合开平标志
	CombOffsetFlag TThostFtdcCombOffsetFlagType
	// 组合投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量
	VolumeTotalOriginal TThostFtdcVolumeType
	// 有效期类型
	TimeCondition TThostFtdcTimeConditionType
	// GTD日期
	GTDDate TThostFtdcDateType
	// 成交量类型
	VolumeCondition TThostFtdcVolumeConditionType
	// 最小成交量
	MinVolume TThostFtdcVolumeType
	// 触发条件
	ContingentCondition TThostFtdcContingentConditionType
	// 止损价
	StopPrice TThostFtdcPriceType
	// 强平原因
	ForceCloseReason TThostFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TThostFtdcBoolType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldExchangeInstIDType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 报单提交状态
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TThostFtdcSequenceNoType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 报单来源
	OrderSource TThostFtdcOrderSourceType
	// 报单状态
	OrderStatus TThostFtdcOrderStatusType
	// 报单类型
	OrderType TThostFtdcOrderTypeType
	// 今成交数量
	VolumeTraded TThostFtdcVolumeType
	// 剩余数量
	VolumeTotal TThostFtdcVolumeType
	// 报单日期
	InsertDate TThostFtdcDateType
	// 委托时间
	InsertTime TThostFtdcTimeType
	// 激活时间
	ActiveTime TThostFtdcTimeType
	// 挂起时间
	SuspendTime TThostFtdcTimeType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 撤销时间
	CancelTime TThostFtdcTimeType
	// 最后修改交易所交易员代码
	ActiveTraderID TThostFtdcTraderIDType
	// 结算会员编号
	ClearingPartID TThostFtdcParticipantIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 用户强评标志
	UserForceClose TThostFtdcBoolType
	// 操作用户代码
	ActiveUserID TThostFtdcUserIDType
	// 经纪公司报单编号
	BrokerOrderSeq TThostFtdcSequenceNoType
	// 相关报单
	RelativeOrderSysID TThostFtdcOrderSysIDType
	// 郑商所成交数量
	ZCETotalTradedVolume TThostFtdcVolumeType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 互换单标志
	IsSwapOrder TThostFtdcBoolType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 资金账号
	AccountID TThostFtdcAccountIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 保留的无效字段
	reserve3 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryErrOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcErrOrderActionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 报单操作引用
	OrderActionRef TThostFtdcOrderActionRefType
	// 报单引用
	OrderRef TThostFtdcOrderRefType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 前置编号
	FrontID TThostFtdcFrontIDType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 报单编号
	OrderSysID TThostFtdcOrderSysIDType
	// 操作标志
	ActionFlag TThostFtdcActionFlagType
	// 价格
	LimitPrice TThostFtdcPriceType
	// 数量变化
	VolumeChange TThostFtdcVolumeType
	// 操作日期
	ActionDate TThostFtdcDateType
	// 操作时间
	ActionTime TThostFtdcTimeType
	// 交易所交易员代码
	TraderID TThostFtdcTraderIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 本地报单编号
	OrderLocalID TThostFtdcOrderLocalIDType
	// 操作本地编号
	ActionLocalID TThostFtdcOrderLocalIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 客户代码
	ClientID TThostFtdcClientIDType
	// 业务单元
	BusinessUnit TThostFtdcBusinessUnitType
	// 报单操作状态
	OrderActionStatus TThostFtdcOrderActionStatusType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 状态信息
	StatusMsg TThostFtdcErrorMsgType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 营业部编号
	BranchID TThostFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryExchangeSequenceField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcExchangeSequenceField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 序号
	SequenceNo TThostFtdcSequenceNoType
	// 合约交易状态
	MarketStatus TThostFtdcInstrumentStatusType
}

// 信息分发
type CThostFtdcQryMaxOrderVolumeWithPriceField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 开平标志
	OffsetFlag TThostFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 最大允许报单数量
	MaxVolume TThostFtdcVolumeType
	// 报单价格
	Price TThostFtdcPriceType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
}

// 信息分发
type CThostFtdcBrokerTradingParamsField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保证金价格类型
	MarginPriceType TThostFtdcMarginPriceTypeType
	// 盈亏算法
	Algorithm TThostFtdcAlgorithmType
	// 可用是否包含平仓盈利
	AvailIncludeCloseProfit TThostFtdcIncludeCloseProfitType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 期权权利金价格类型
	OptionRoyaltyPriceType TThostFtdcOptionRoyaltyPriceTypeType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
}

// 信息分发
type CThostFtdcQryBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcBrokerTradingAlgosField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 持仓处理算法编号
	HandlePositionAlgoID TThostFtdcHandlePositionAlgoIDType
	// 寻找保证金率算法编号
	FindMarginRateAlgoID TThostFtdcFindMarginRateAlgoIDType
	// 资金处理算法编号
	HandleTradingAccountAlgoID TThostFtdcHandleTradingAccountAlgoIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQueryBrokerDepositField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcBrokerDepositField struct {
	// 交易日期
	TradingDay TThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 会员代码
	ParticipantID TThostFtdcParticipantIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 上次结算准备金
	PreBalance TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 期货结算准备金
	Balance TThostFtdcMoneyType
	// 入金金额
	Deposit TThostFtdcMoneyType
	// 出金金额
	Withdraw TThostFtdcMoneyType
	// 可提资金
	Available TThostFtdcMoneyType
	// 基本准备金
	Reserve TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
}

// 信息分发
type CThostFtdcQryCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcCFMMCBrokerKeyField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID TThostFtdcParticipantIDType
	// 密钥生成日期
	CreateDate TThostFtdcDateType
	// 密钥生成时间
	CreateTime TThostFtdcTimeType
	// 密钥编号
	KeyID TThostFtdcSequenceNoType
	// 动态密钥
	CurrentKey TThostFtdcCFMMCKeyType
	// 动态密钥类型
	KeyKind TThostFtdcCFMMCKeyKindType
}

// 信息分发
type CThostFtdcCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID TThostFtdcParticipantIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 密钥编号
	KeyID TThostFtdcSequenceNoType
	// 动态密钥
	CurrentKey TThostFtdcCFMMCKeyType
}

// 信息分发
type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcBrokerUserOTPParamField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 动态令牌提供商
	OTPVendorsID TThostFtdcOTPVendorsIDType
	// 动态令牌序列号
	SerialNumber TThostFtdcSerialNumberType
	// 令牌密钥
	AuthKey TThostFtdcAuthKeyType
	// 漂移值
	LastDrift TThostFtdcLastDriftType
	// 成功值
	LastSuccess TThostFtdcLastSuccessType
	// 动态令牌类型
	OTPType TThostFtdcOTPTypeType
}

// 信息分发
type CThostFtdcManualSyncBrokerUserOTPField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 动态令牌类型
	OTPType TThostFtdcOTPTypeType
	// 第一个动态密码
	FirstOTP TThostFtdcPasswordType
	// 第二个动态密码
	SecondOTP TThostFtdcPasswordType
}

// 信息分发
type CThostFtdcCommRateModelField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID TThostFtdcInvestorIDType
	// 模板名称
	CommModelName TThostFtdcCommModelNameType
}

// 信息分发
type CThostFtdcQryCommRateModelField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 手续费率模板代码
	CommModelID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcMarginModelField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID TThostFtdcInvestorIDType
	// 模板名称
	MarginModelName TThostFtdcCommModelNameType
}

// 信息分发
type CThostFtdcQryMarginModelField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 保证金率模板代码
	MarginModelID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcEWarrantOffsetField struct {
	// 交易日期
	TradingDay TThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 数量
	Volume TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryEWarrantOffsetField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryInvestorProductGroupMarginField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 品种/跨品种标示
	ProductGroupID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcInvestorProductGroupMarginField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 多头冻结的保证金
	LongFrozenMargin TThostFtdcMoneyType
	// 空头冻结的保证金
	ShortFrozenMargin TThostFtdcMoneyType
	// 占用的保证金
	UseMargin TThostFtdcMoneyType
	// 多头保证金
	LongUseMargin TThostFtdcMoneyType
	// 空头保证金
	ShortUseMargin TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin TThostFtdcMoneyType
	// 交易所多头保证金
	LongExchMargin TThostFtdcMoneyType
	// 交易所空头保证金
	ShortExchMargin TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 折抵总金额
	OffsetAmount TThostFtdcMoneyType
	// 多头折抵总金额
	LongOffsetAmount TThostFtdcMoneyType
	// 空头折抵总金额
	ShortOffsetAmount TThostFtdcMoneyType
	// 交易所折抵总金额
	ExchOffsetAmount TThostFtdcMoneyType
	// 交易所多头折抵总金额
	LongExchOffsetAmount TThostFtdcMoneyType
	// 交易所空头折抵总金额
	ShortExchOffsetAmount TThostFtdcMoneyType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 品种/跨品种标示
	ProductGroupID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
}

// 信息分发
type CThostFtdcCFMMCTradingAccountTokenField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 经纪公司统一编码
	ParticipantID TThostFtdcParticipantIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 密钥编号
	KeyID TThostFtdcSequenceNoType
	// 动态令牌
	Token TThostFtdcCFMMCTokenType
}

// 信息分发
type CThostFtdcQryProductGroupField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcProductGroupField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 保留的无效字段
	reserve2 TThostFtdcOldInstrumentIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 产品组代码
	ProductGroupID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcBulletinField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 交易日
	TradingDay TThostFtdcDateType
	// 公告编号
	BulletinID TThostFtdcBulletinIDType
	// 序列号
	SequenceNo TThostFtdcSequenceNoType
	// 公告类型
	NewsType TThostFtdcNewsTypeType
	// 紧急程度
	NewsUrgency TThostFtdcNewsUrgencyType
	// 发送时间
	SendTime TThostFtdcTimeType
	// 消息摘要
	Abstract TThostFtdcAbstractType
	// 消息来源
	ComeFrom TThostFtdcComeFromType
	// 消息正文
	Content TThostFtdcContentType
	// WEB地址
	URLLink TThostFtdcURLLinkType
	// 市场代码
	MarketID TThostFtdcMarketIDType
}

// 信息分发
type CThostFtdcQryBulletinField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 公告编号
	BulletinID TThostFtdcBulletinIDType
	// 序列号
	SequenceNo TThostFtdcSequenceNoType
	// 公告类型
	NewsType TThostFtdcNewsTypeType
	// 紧急程度
	NewsUrgency TThostFtdcNewsUrgencyType
}

// 信息分发
type CThostFtdcMulticastInstrumentField struct {
	// 主题号
	TopicID TThostFtdcInstallIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约编号
	InstrumentNo TThostFtdcInstallIDType
	// 基准价
	CodePrice TThostFtdcPriceType
	// 合约数量乘数
	VolumeMultiple TThostFtdcVolumeMultipleType
	// 最小变动价位
	PriceTick TThostFtdcPriceType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryMulticastInstrumentField struct {
	// 主题号
	TopicID TThostFtdcInstallIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldInstrumentIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcAppIDAuthAssignField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// App代码
	AppID TThostFtdcAppIDType
	// 交易中心代码
	DRIdentityID TThostFtdcDRIdentityIDType
}

// 信息分发
type CThostFtdcReqOpenAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 交易ID
	TID TThostFtdcTIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcReqCancelAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 交易ID
	TID TThostFtdcTIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcReqChangeAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 新银行帐号
	NewBankAccount TThostFtdcBankAccountType
	// 新银行密码
	NewBankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易ID
	TID TThostFtdcTIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcReqTransferField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message TThostFtdcAddInfoType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 转账交易状态
	TransferStatus TThostFtdcTransferStatusType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcRspTransferField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message TThostFtdcAddInfoType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 转账交易状态
	TransferStatus TThostFtdcTransferStatusType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcReqRepealField struct {
	// 冲正时间间隔
	RepealTimeInterval TThostFtdcRepealTimeIntervalType
	// 已经冲正次数
	RepealedTimes TThostFtdcRepealedTimesType
	// 银行冲正标志
	BankRepealFlag TThostFtdcBankRepealFlagType
	// 期商冲正标志
	BrokerRepealFlag TThostFtdcBrokerRepealFlagType
	// 被冲正平台流水号
	PlateRepealSerial TThostFtdcPlateSerialType
	// 被冲正银行流水号
	BankRepealSerial TThostFtdcBankSerialType
	// 被冲正期货流水号
	FutureRepealSerial TThostFtdcFutureSerialType
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message TThostFtdcAddInfoType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 转账交易状态
	TransferStatus TThostFtdcTransferStatusType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcRspRepealField struct {
	// 冲正时间间隔
	RepealTimeInterval TThostFtdcRepealTimeIntervalType
	// 已经冲正次数
	RepealedTimes TThostFtdcRepealedTimesType
	// 银行冲正标志
	BankRepealFlag TThostFtdcBankRepealFlagType
	// 期商冲正标志
	BrokerRepealFlag TThostFtdcBrokerRepealFlagType
	// 被冲正平台流水号
	PlateRepealSerial TThostFtdcPlateSerialType
	// 被冲正银行流水号
	BankRepealSerial TThostFtdcBankSerialType
	// 被冲正期货流水号
	FutureRepealSerial TThostFtdcFutureSerialType
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount TThostFtdcTradeAmountType
	// 期货可取金额
	FutureFetchAmount TThostFtdcTradeAmountType
	// 费用支付标志
	FeePayFlag TThostFtdcFeePayFlagType
	// 应收客户费用
	CustFee TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee TThostFtdcFutureFeeType
	// 发送方给接收方的消息
	Message TThostFtdcAddInfoType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 转账交易状态
	TransferStatus TThostFtdcTransferStatusType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcReqQueryAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcRspQueryAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 银行可用金额
	BankUseAmount TThostFtdcTradeAmountType
	// 银行可取金额
	BankFetchAmount TThostFtdcTradeAmountType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcFutureSignIOField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
}

// 信息分发
type CThostFtdcRspFutureSignInField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// PIN密钥
	PinKey TThostFtdcPasswordKeyType
	// MAC密钥
	MacKey TThostFtdcPasswordKeyType
}

// 信息分发
type CThostFtdcReqFutureSignOutField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
}

// 信息分发
type CThostFtdcRspFutureSignOutField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcReqQueryTradeResultBySerialField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 流水号
	Reference TThostFtdcSerialType
	// 本流水号发布者的机构类型
	RefrenceIssureType TThostFtdcInstitutionTypeType
	// 本流水号发布者机构编码
	RefrenceIssure TThostFtdcOrganCodeType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount TThostFtdcTradeAmountType
	// 摘要
	Digest TThostFtdcDigestType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcRspQueryTradeResultBySerialField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 流水号
	Reference TThostFtdcSerialType
	// 本流水号发布者的机构类型
	RefrenceIssureType TThostFtdcInstitutionTypeType
	// 本流水号发布者机构编码
	RefrenceIssure TThostFtdcOrganCodeType
	// 原始返回代码
	OriginReturnCode TThostFtdcReturnCodeType
	// 原始返回码描述
	OriginDescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 转帐金额
	TradeAmount TThostFtdcTradeAmountType
	// 摘要
	Digest TThostFtdcDigestType
}

// 信息分发
type CThostFtdcReqDayEndFileReadyField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 文件业务功能
	FileBusinessCode TThostFtdcFileBusinessCodeType
	// 摘要
	Digest TThostFtdcDigestType
}

// 信息分发
type CThostFtdcReturnResultField struct {
	// 返回代码
	ReturnCode TThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType
}

// 信息分发
type CThostFtdcVerifyFuturePasswordField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcVerifyCustInfoField struct {
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcDepositResultInformField struct {
	// 出入金流水号，该流水号为银期报盘返回的流水号
	DepositSeqNo TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 入金金额
	Deposit TThostFtdcMoneyType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 返回代码
	ReturnCode TThostFtdcReturnCodeType
	// 返回码描述
	DescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType
}

// 信息分发
type CThostFtdcReqSyncKeyField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message TThostFtdcAddInfoType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
}

// 信息分发
type CThostFtdcRspSyncKeyField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message TThostFtdcAddInfoType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcNotifyQueryAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 银行可用金额
	BankUseAmount TThostFtdcTradeAmountType
	// 银行可取金额
	BankFetchAmount TThostFtdcTradeAmountType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcTransferSerialField struct {
	// 平台流水号
	PlateSerial TThostFtdcPlateSerialType
	// 交易发起方日期
	TradeDate TThostFtdcTradeDateType
	// 交易日期
	TradingDay TThostFtdcDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 交易代码
	TradeCode TThostFtdcTradeCodeType
	// 会话编号
	SessionID TThostFtdcSessionIDType
	// 银行编码
	BankID TThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 期货公司编码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 期货公司帐号类型
	FutureAccType TThostFtdcFutureAccTypeType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 期货公司流水号
	FutureSerial TThostFtdcFutureSerialType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 交易金额
	TradeAmount TThostFtdcTradeAmountType
	// 应收客户费用
	CustFee TThostFtdcCustFeeType
	// 应收期货公司费用
	BrokerFee TThostFtdcFutureFeeType
	// 有效标志
	AvailabilityFlag TThostFtdcAvailabilityFlagType
	// 操作员
	OperatorCode TThostFtdcOperatorCodeType
	// 新银行帐号
	BankNewAccount TThostFtdcBankAccountType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcQryTransferSerialField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 银行编码
	BankID TThostFtdcBankIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcNotifyFutureSignInField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// PIN密钥
	PinKey TThostFtdcPasswordKeyType
	// MAC密钥
	MacKey TThostFtdcPasswordKeyType
}

// 信息分发
type CThostFtdcNotifyFutureSignOutField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcNotifySyncKeyField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 交易核心给银期报盘的消息
	Message TThostFtdcAddInfoType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 请求编号
	RequestID TThostFtdcRequestIDType
	// 交易ID
	TID TThostFtdcTIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcQryAccountregisterField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 银行编码
	BankID TThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID TThostFtdcBankBrchIDType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcAccountregisterField struct {
	// 交易日期
	TradeDay TThostFtdcTradeDateType
	// 银行编码
	BankID TThostFtdcBankIDType
	// 银行分支机构编码
	BankBranchID TThostFtdcBankBrchIDType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 期货公司编码
	BrokerID TThostFtdcBrokerIDType
	// 期货公司分支机构编码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 开销户类别
	OpenOrDestroy TThostFtdcOpenOrDestroyType
	// 签约日期
	RegDate TThostFtdcTradeDateType
	// 解约日期
	OutDate TThostFtdcTradeDateType
	// 交易ID
	TID TThostFtdcTIDType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcOpenAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 交易ID
	TID TThostFtdcTIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcCancelAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 汇钞标志
	CashExchangeCode TThostFtdcCashExchangeCodeType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 渠道标志
	DeviceID TThostFtdcDeviceIDType
	// 期货单位帐号类型
	BankSecuAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 期货单位帐号
	BankSecuAcc TThostFtdcBankAccountType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易柜员
	OperNo TThostFtdcOperNoType
	// 交易ID
	TID TThostFtdcTIDType
	// 用户标识
	UserID TThostFtdcUserIDType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcChangeAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 新银行帐号
	NewBankAccount TThostFtdcBankAccountType
	// 新银行密码
	NewBankPassWord TThostFtdcPasswordType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 银行密码标志
	BankPwdFlag TThostFtdcPwdFlagType
	// 期货资金密码核对标志
	SecuPwdFlag TThostFtdcPwdFlagType
	// 交易ID
	TID TThostFtdcTIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
	// 长客户姓名
	LongCustomerName TThostFtdcLongIndividualNameType
}

// 信息分发
type CThostFtdcSecAgentACIDMapField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 资金账户
	AccountID TThostFtdcAccountIDType
	// 币种
	CurrencyID TThostFtdcCurrencyIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID TThostFtdcAccountIDType
}

// 信息分发
type CThostFtdcQrySecAgentACIDMapField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 资金账户
	AccountID TThostFtdcAccountIDType
	// 币种
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcUserRightsAssignField struct {
	// 应用单元代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 交易中心代码
	DRIdentityID TThostFtdcDRIdentityIDType
}

// 信息分发
type CThostFtdcBrokerUserRightAssignField struct {
	// 应用单元代码
	BrokerID TThostFtdcBrokerIDType
	// 交易中心代码
	DRIdentityID TThostFtdcDRIdentityIDType
	// 能否交易
	Tradeable TThostFtdcBoolType
}

// 信息分发
type CThostFtdcDRTransferField struct {
	// 原交易中心代码
	OrigDRIdentityID TThostFtdcDRIdentityIDType
	// 目标交易中心代码
	DestDRIdentityID TThostFtdcDRIdentityIDType
	// 原应用单元代码
	OrigBrokerID TThostFtdcBrokerIDType
	// 目标易用单元代码
	DestBrokerID TThostFtdcBrokerIDType
}

// 信息分发
type CThostFtdcFensUserInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 登录模式
	LoginMode TThostFtdcLoginModeType
}

// 信息分发
type CThostFtdcCurrTransferIdentityField struct {
	// 交易中心代码
	IdentityID TThostFtdcDRIdentityIDType
}

// 信息分发
type CThostFtdcLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryLoginForbiddenUserField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcTradingAccountReserveField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 基本准备金
	Reserve TThostFtdcMoneyType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcQryLoginForbiddenIPField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryIPListField struct {
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryUserRightsAssignField struct {
	// 应用单元代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcReserveOpenAccountConfirmField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcLongIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易ID
	TID TThostFtdcTIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 期货密码
	Password TThostFtdcPasswordType
	// 预约开户银行流水号
	BankReserveOpenSeq TThostFtdcBankSerialType
	// 预约开户日期
	BookDate TThostFtdcTradeDateType
	// 预约开户验证密码
	BookPsw TThostFtdcPasswordType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcReserveOpenAccountField struct {
	// 业务功能码
	TradeCode TThostFtdcTradeCodeType
	// 银行代码
	BankID TThostFtdcBankIDType
	// 银行分支机构代码
	BankBranchID TThostFtdcBankBrchIDType
	// 期商代码
	BrokerID TThostFtdcBrokerIDType
	// 期商分支机构代码
	BrokerBranchID TThostFtdcFutureBranchIDType
	// 交易日期
	TradeDate TThostFtdcTradeDateType
	// 交易时间
	TradeTime TThostFtdcTradeTimeType
	// 银行流水号
	BankSerial TThostFtdcBankSerialType
	// 交易系统日期
	TradingDay TThostFtdcTradeDateType
	// 银期平台消息流水号
	PlateSerial TThostFtdcSerialType
	// 最后分片标志
	LastFragment TThostFtdcLastFragmentType
	// 会话号
	SessionID TThostFtdcSessionIDType
	// 客户姓名
	CustomerName TThostFtdcLongIndividualNameType
	// 证件类型
	IdCardType TThostFtdcIdCardTypeType
	// 证件号码
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType
	// 性别
	Gender TThostFtdcGenderType
	// 国家代码
	CountryCode TThostFtdcCountryCodeType
	// 客户类型
	CustType TThostFtdcCustTypeType
	// 地址
	Address TThostFtdcAddressType
	// 邮编
	ZipCode TThostFtdcZipCodeType
	// 电话号码
	Telephone TThostFtdcTelephoneType
	// 手机
	MobilePhone TThostFtdcMobilePhoneType
	// 传真
	Fax TThostFtdcFaxType
	// 电子邮件
	EMail TThostFtdcEMailType
	// 资金账户状态
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType
	// 银行帐号
	BankAccount TThostFtdcBankAccountType
	// 银行密码
	BankPassWord TThostFtdcPasswordType
	// 安装编号
	InstallID TThostFtdcInstallIDType
	// 验证客户证件号码标志
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 摘要
	Digest TThostFtdcDigestType
	// 银行帐号类型
	BankAccType TThostFtdcBankAccTypeType
	// 期货公司银行编码
	BrokerIDByBank TThostFtdcBankCodingForFutureType
	// 交易ID
	TID TThostFtdcTIDType
	// 预约开户状态
	ReserveOpenAccStas TThostFtdcReserveOpenAccStasType
	// 错误代码
	ErrorID TThostFtdcErrorIDType
	// 错误信息
	ErrorMsg TThostFtdcErrorMsgType
}

// 信息分发
type CThostFtdcAccountPropertyField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 银行统一标识类型
	BankID TThostFtdcBankIDType
	// 银行账户
	BankAccount TThostFtdcBankAccountType
	// 银行账户的开户人名称
	OpenName TThostFtdcInvestorFullNameType
	// 银行账户的开户行
	OpenBank TThostFtdcOpenBankType
	// 是否活跃
	IsActive TThostFtdcBoolType
	// 账户来源
	AccountSourceType TThostFtdcAccountSourceTypeType
	// 开户日期
	OpenDate TThostFtdcDateType
	// 注销日期
	CancelDate TThostFtdcDateType
	// 录入员代码
	OperatorID TThostFtdcOperatorIDType
	// 录入日期
	OperateDate TThostFtdcDateType
	// 录入时间
	OperateTime TThostFtdcTimeType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
}

// 信息分发
type CThostFtdcQryCurrDRIdentityField struct {
	// 交易中心代码
	DRIdentityID TThostFtdcDRIdentityIDType
}

// 信息分发
type CThostFtdcCurrDRIdentityField struct {
	// 交易中心代码
	DRIdentityID TThostFtdcDRIdentityIDType
}

// 信息分发
type CThostFtdcQrySecAgentCheckModeField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcQrySecAgentTradeInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 境外中介机构资金帐号
	BrokerSecAgentID TThostFtdcAccountIDType
}

// 信息分发
type CThostFtdcReqUserAuthMethodField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcRspUserAuthMethodField struct {
	// 当前可以用的认证模式
	UsableAuthMethod TThostFtdcCurrentAuthMethodType
}

// 信息分发
type CThostFtdcReqGenUserCaptchaField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcRspGenUserCaptchaField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 图片信息长度
	CaptchaInfoLen TThostFtdcCaptchaInfoLenType
	// 图片信息
	CaptchaInfo TThostFtdcCaptchaInfoType
}

// 信息分发
type CThostFtdcReqGenUserTextField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
}

// 信息分发
type CThostFtdcRspGenUserTextField struct {
	// 短信验证码序号
	UserTextSeq TThostFtdcUserTextSeqType
}

// 信息分发
type CThostFtdcReqUserLoginWithCaptchaField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 密码
	Password TThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// 图形验证码的文字内容
	Captcha TThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcReqUserLoginWithTextField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 密码
	Password TThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// 短信验证码文字内容
	Text TThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcReqUserLoginWithOTPField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 密码
	Password TThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// OTP密码
	OTPPassword TThostFtdcPasswordType
	// 终端IP端口
	ClientIPPort TThostFtdcIPPortType
	// 终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcReqApiHandshakeField struct {
	// api与front通信密钥版本号
	CryptoKeyVersion TThostFtdcCryptoKeyVersionType
}

// 信息分发
type CThostFtdcRspApiHandshakeField struct {
	// 握手回复数据长度
	FrontHandshakeDataLen TThostFtdcHandshakeDataLenType
	// 握手回复数据
	FrontHandshakeData TThostFtdcHandshakeDataType
	// API认证是否开启
	IsApiAuthEnabled TThostFtdcBoolType
}

// 信息分发
type CThostFtdcReqVerifyApiKeyField struct {
	// 握手回复数据长度
	ApiHandshakeDataLen TThostFtdcHandshakeDataLenType
	// 握手回复数据
	ApiHandshakeData TThostFtdcHandshakeDataType
}

// 信息分发
type CThostFtdcDepartmentUserField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 投资者范围
	InvestorRange TThostFtdcDepartmentRangeType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcQueryFreqField struct {
	// 查询频率
	QueryFreq TThostFtdcQueryFreqType
}

// 信息分发
type CThostFtdcAuthForbiddenIPField struct {
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryAuthForbiddenIPField struct {
	// IP地址
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcSyncDelaySwapFrozenField struct {
	// 换汇流水号
	DelaySwapSeqNo TThostFtdcDepositSeqNoType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 源币种
	FromCurrencyID TThostFtdcCurrencyIDType
	// 源剩余换汇额度(可提冻结)
	FromRemainSwap TThostFtdcMoneyType
	// 是否手工换汇
	IsManualSwap TThostFtdcBoolType
}

// 信息分发
type CThostFtdcUserSystemInfoField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 用户端系统内部信息长度
	ClientSystemInfoLen TThostFtdcSystemInfoLenType
	// 用户端系统内部信息
	ClientSystemInfo TThostFtdcClientSystemInfoType
	// 保留的无效字段
	reserve1 TThostFtdcOldIPAddressType
	// 终端IP端口
	ClientIPPort TThostFtdcIPPortType
	// 登录成功时间
	ClientLoginTime TThostFtdcTimeType
	// App代码
	ClientAppID TThostFtdcAppIDType
	// 用户公网IP
	ClientPublicIP TThostFtdcIPAddressType
	// 客户登录备注2
	ClientLoginRemark TThostFtdcClientLoginRemarkType
}

// 信息分发
type CThostFtdcAuthUserIDField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// App代码
	AppID TThostFtdcAppIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 校验类型
	AuthType TThostFtdcAuthTypeType
}

// 信息分发
type CThostFtdcAuthIPField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// App代码
	AppID TThostFtdcAppIDType
	// 用户代码
	IPAddress TThostFtdcIPAddressType
}

// 信息分发
type CThostFtdcQryClassifiedInstrumentField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 合约交易状态
	TradingType TThostFtdcTradingTypeType
	// 合约分类类型
	ClassType TThostFtdcClassTypeType
}

// 信息分发
type CThostFtdcQryCombPromotionParamField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcCombPromotionParamField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投机套保标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType
	// 期权组合保证金比例
	Xparameter TThostFtdcDiscountRatioType
}

// 信息分发
type CThostFtdcReqUserLoginSCField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 用户代码
	UserID TThostFtdcUserIDType
	// 密码
	Password TThostFtdcPasswordType
	// 用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	// 协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	// Mac地址
	MacAddress TThostFtdcMacAddressType
	// 动态密码
	OneTimePassword TThostFtdcPasswordType
	// 终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
	// 登录备注
	LoginRemark TThostFtdcLoginRemarkType
	// 终端IP端口
	ClientIPPort TThostFtdcIPPortType
	// 认证码
	AuthCode TThostFtdcAuthCodeType
	// App代码
	AppID TThostFtdcAppIDType
}

// 信息分发
type CThostFtdcQryRiskSettleInvstPositionField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQryRiskSettleProductStatusField struct {
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcRiskSettleInvstPositionField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 持仓多空方向
	PosiDirection TThostFtdcPosiDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 持仓日期
	PositionDate TThostFtdcPositionDateType
	// 上日持仓
	YdPosition TThostFtdcVolumeType
	// 今日持仓
	Position TThostFtdcVolumeType
	// 多头冻结
	LongFrozen TThostFtdcVolumeType
	// 空头冻结
	ShortFrozen TThostFtdcVolumeType
	// 开仓冻结金额
	LongFrozenAmount TThostFtdcMoneyType
	// 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType
	// 开仓量
	OpenVolume TThostFtdcVolumeType
	// 平仓量
	CloseVolume TThostFtdcVolumeType
	// 开仓金额
	OpenAmount TThostFtdcMoneyType
	// 平仓金额
	CloseAmount TThostFtdcMoneyType
	// 持仓成本
	PositionCost TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin TThostFtdcMoneyType
	// 占用的保证金
	UseMargin TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 开仓成本
	OpenCost TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin TThostFtdcMoneyType
	// 组合成交形成的持仓
	CombPosition TThostFtdcVolumeType
	// 组合多头冻结
	CombLongFrozen TThostFtdcVolumeType
	// 组合空头冻结
	CombShortFrozen TThostFtdcVolumeType
	// 逐日盯市平仓盈亏
	CloseProfitByDate TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 今日持仓
	TodayPosition TThostFtdcVolumeType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 执行冻结
	StrikeFrozen TThostFtdcVolumeType
	// 执行冻结金额
	StrikeFrozenAmount TThostFtdcMoneyType
	// 放弃执行冻结
	AbandonFrozen TThostFtdcVolumeType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 执行冻结的昨仓
	YdStrikeFrozen TThostFtdcVolumeType
	// 投资单元代码
	InvestUnitID TThostFtdcInvestUnitIDType
	// 持仓成本差值
	PositionCostOffset TThostFtdcMoneyType
	// tas持仓手数
	TasPosition TThostFtdcVolumeType
	// tas持仓成本
	TasPositionCost TThostFtdcMoneyType
}

// 信息分发
type CThostFtdcRiskSettleProductStatusField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品编号
	ProductID TThostFtdcInstrumentIDType
	// 产品结算状态
	ProductStatus TThostFtdcProductStatusType
}

// 信息分发
type CThostFtdcSyncDeltaInfoField struct {
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
	// 追平状态
	SyncDeltaStatus TThostFtdcSyncDeltaStatusType
	// 追平描述
	SyncDescription TThostFtdcSyncDescriptionType
	// 是否只有资金追平
	IsOnlyTrdDelta TThostFtdcBoolType
}

// 信息分发
type CThostFtdcSyncDeltaProductStatusField struct {
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 是否允许交易
	ProductStatus TThostFtdcProductStatusType
}

// 信息分发
type CThostFtdcSyncDeltaInvstPosDtlField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 买卖
	Direction TThostFtdcDirectionType
	// 开仓日期
	OpenDate TThostFtdcDateType
	// 成交编号
	TradeID TThostFtdcTradeIDType
	// 数量
	Volume TThostFtdcVolumeType
	// 开仓价
	OpenPrice TThostFtdcPriceType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 成交类型
	TradeType TThostFtdcTradeTypeType
	// 组合合约代码
	CombInstrumentID TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 逐日盯市平仓盈亏
	CloseProfitByDate TThostFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType
	// 逐日盯市持仓盈亏
	PositionProfitByDate TThostFtdcMoneyType
	// 逐笔对冲持仓盈亏
	PositionProfitByTrade TThostFtdcMoneyType
	// 投资者保证金
	Margin TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin TThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 昨结算价
	LastSettlementPrice TThostFtdcPriceType
	// 结算价
	SettlementPrice TThostFtdcPriceType
	// 平仓量
	CloseVolume TThostFtdcVolumeType
	// 平仓金额
	CloseAmount TThostFtdcMoneyType
	// 先开先平剩余数量
	TimeFirstVolume TThostFtdcVolumeType
	// 特殊持仓标志
	SpecPosiType TThostFtdcSpecPosiTypeType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaInvstPosCombDtlField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 开仓日期
	OpenDate TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 组合编号
	ComTradeID TThostFtdcTradeIDType
	// 撮合编号
	TradeID TThostFtdcTradeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 买卖
	Direction TThostFtdcDirectionType
	// 持仓量
	TotalAmt TThostFtdcVolumeType
	// 投资者保证金
	Margin TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin TThostFtdcMoneyType
	// 保证金率
	MarginRateByMoney TThostFtdcRatioType
	// 保证金率(按手数)
	MarginRateByVolume TThostFtdcRatioType
	// 单腿编号
	LegID TThostFtdcLegIDType
	// 单腿乘数
	LegMultiple TThostFtdcLegMultipleType
	// 成交组号
	TradeGroupID TThostFtdcTradeGroupIDType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaTradingAccountField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者帐号
	AccountID TThostFtdcAccountIDType
	// 上次质押金额
	PreMortgage TThostFtdcMoneyType
	// 上次信用额度
	PreCredit TThostFtdcMoneyType
	// 上次存款额
	PreDeposit TThostFtdcMoneyType
	// 上次结算准备金
	PreBalance TThostFtdcMoneyType
	// 上次占用的保证金
	PreMargin TThostFtdcMoneyType
	// 利息基数
	InterestBase TThostFtdcMoneyType
	// 利息收入
	Interest TThostFtdcMoneyType
	// 入金金额
	Deposit TThostFtdcMoneyType
	// 出金金额
	Withdraw TThostFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TThostFtdcMoneyType
	// 冻结的资金
	FrozenCash TThostFtdcMoneyType
	// 冻结的手续费
	FrozenCommission TThostFtdcMoneyType
	// 当前保证金总额
	CurrMargin TThostFtdcMoneyType
	// 资金差额
	CashIn TThostFtdcMoneyType
	// 手续费
	Commission TThostFtdcMoneyType
	// 平仓盈亏
	CloseProfit TThostFtdcMoneyType
	// 持仓盈亏
	PositionProfit TThostFtdcMoneyType
	// 期货结算准备金
	Balance TThostFtdcMoneyType
	// 可用资金
	Available TThostFtdcMoneyType
	// 可取资金
	WithdrawQuota TThostFtdcMoneyType
	// 基本准备金
	Reserve TThostFtdcMoneyType
	// 交易日
	TradingDay TThostFtdcDateType
	// 结算编号
	SettlementID TThostFtdcSettlementIDType
	// 信用额度
	Credit TThostFtdcMoneyType
	// 质押金额
	Mortgage TThostFtdcMoneyType
	// 交易所保证金
	ExchangeMargin TThostFtdcMoneyType
	// 投资者交割保证金
	DeliveryMargin TThostFtdcMoneyType
	// 交易所交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType
	// 保底期货结算准备金
	ReserveBalance TThostFtdcMoneyType
	// 币种代码
	CurrencyID TThostFtdcCurrencyIDType
	// 上次货币质入金额
	PreFundMortgageIn TThostFtdcMoneyType
	// 上次货币质出金额
	PreFundMortgageOut TThostFtdcMoneyType
	// 货币质入金额
	FundMortgageIn TThostFtdcMoneyType
	// 货币质出金额
	FundMortgageOut TThostFtdcMoneyType
	// 货币质押余额
	FundMortgageAvailable TThostFtdcMoneyType
	// 可质押货币金额
	MortgageableFund TThostFtdcMoneyType
	// 特殊产品占用保证金
	SpecProductMargin TThostFtdcMoneyType
	// 特殊产品冻结保证金
	SpecProductFrozenMargin TThostFtdcMoneyType
	// 特殊产品手续费
	SpecProductCommission TThostFtdcMoneyType
	// 特殊产品冻结手续费
	SpecProductFrozenCommission TThostFtdcMoneyType
	// 特殊产品持仓盈亏
	SpecProductPositionProfit TThostFtdcMoneyType
	// 特殊产品平仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType
	// 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType
	// 特殊产品交易所保证金
	SpecProductExchangeMargin TThostFtdcMoneyType
	// 延时换汇冻结金额
	FrozenSwap TThostFtdcMoneyType
	// 剩余换汇额度
	RemainSwap TThostFtdcMoneyType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaInitInvstMarginField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 追平前总风险保证金
	LastRiskTotalInvstMargin TThostFtdcMoneyType
	// 追平前交易所总风险保证金
	LastRiskTotalExchMargin TThostFtdcMoneyType
	// 本次追平品种总保证金
	ThisSyncInvstMargin TThostFtdcMoneyType
	// 本次追平品种交易所总保证金
	ThisSyncExchMargin TThostFtdcMoneyType
	// 本次未追平品种总保证金
	RemainRiskInvstMargin TThostFtdcMoneyType
	// 本次未追平品种交易所总保证金
	RemainRiskExchMargin TThostFtdcMoneyType
	// 追平前总特殊产品风险保证金
	LastRiskSpecTotalInvstMargin TThostFtdcMoneyType
	// 追平前总特殊产品交易所风险保证金
	LastRiskSpecTotalExchMargin TThostFtdcMoneyType
	// 本次追平品种特殊产品总保证金
	ThisSyncSpecInvstMargin TThostFtdcMoneyType
	// 本次追平品种特殊产品交易所总保证金
	ThisSyncSpecExchMargin TThostFtdcMoneyType
	// 本次未追平品种特殊产品总保证金
	RemainRiskSpecInvstMargin TThostFtdcMoneyType
	// 本次未追平品种特殊产品交易所总保证金
	RemainRiskSpecExchMargin TThostFtdcMoneyType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaDceCombInstrumentField struct {
	// 合约代码
	CombInstrumentID TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 成交组号
	TradeGroupID TThostFtdcTradeGroupIDType
	// 投机套保标志
	CombHedgeFlag TThostFtdcHedgeFlagType
	// 组合类型
	CombinationType TThostFtdcDceCombinationTypeType
	// 买卖
	Direction TThostFtdcDirectionType
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 期权组合保证金比例
	Xparameter TThostFtdcDiscountRatioType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaInvstMarginRateField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 是否相对交易所收取
	IsRelative TThostFtdcBoolType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaExchMarginRateField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaOptExchMarginField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投机空头保证金调整系数
	SShortMarginRatioByMoney TThostFtdcRatioType
	// 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType
	// 保值空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType
	// 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType
	// 套利空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType
	// 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType
	// 做市商空头保证金调整系数
	MShortMarginRatioByMoney TThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaOptInvstMarginField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机空头保证金调整系数
	SShortMarginRatioByMoney TThostFtdcRatioType
	// 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType
	// 保值空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType
	// 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType
	// 套利空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType
	// 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType
	// 是否跟随交易所收取
	IsRelative TThostFtdcBoolType
	// 做市商空头保证金调整系数
	MShortMarginRatioByMoney TThostFtdcRatioType
	// 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaInvstMarginRateULField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 多头保证金率
	LongMarginRatioByMoney TThostFtdcRatioType
	// 多头保证金费
	LongMarginRatioByVolume TThostFtdcMoneyType
	// 空头保证金率
	ShortMarginRatioByMoney TThostFtdcRatioType
	// 空头保证金费
	ShortMarginRatioByVolume TThostFtdcMoneyType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaOptInvstCommRateField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 执行手续费率
	StrikeRatioByMoney TThostFtdcRatioType
	// 执行手续费
	StrikeRatioByVolume TThostFtdcRatioType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaInvstCommRateField struct {
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 开仓手续费率
	OpenRatioByMoney TThostFtdcRatioType
	// 开仓手续费
	OpenRatioByVolume TThostFtdcRatioType
	// 平仓手续费率
	CloseRatioByMoney TThostFtdcRatioType
	// 平仓手续费
	CloseRatioByVolume TThostFtdcRatioType
	// 平今手续费率
	CloseTodayRatioByMoney TThostFtdcRatioType
	// 平今手续费
	CloseTodayRatioByVolume TThostFtdcRatioType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaProductExchRateField struct {
	// 产品代码
	ProductID TThostFtdcInstrumentIDType
	// 报价币种类型
	QuoteCurrencyID TThostFtdcCurrencyIDType
	// 汇率
	ExchangeRate TThostFtdcExchangeRateType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaDepthMarketDataField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约在交易所的代码
	ExchangeInstID TThostFtdcExchangeInstIDType
	// 最新价
	LastPrice TThostFtdcPriceType
	// 上次结算价
	PreSettlementPrice TThostFtdcPriceType
	// 昨收盘
	PreClosePrice TThostFtdcPriceType
	// 昨持仓量
	PreOpenInterest TThostFtdcLargeVolumeType
	// 今开盘
	OpenPrice TThostFtdcPriceType
	// 最高价
	HighestPrice TThostFtdcPriceType
	// 最低价
	LowestPrice TThostFtdcPriceType
	// 数量
	Volume TThostFtdcVolumeType
	// 成交金额
	Turnover TThostFtdcMoneyType
	// 持仓量
	OpenInterest TThostFtdcLargeVolumeType
	// 今收盘
	ClosePrice TThostFtdcPriceType
	// 本次结算价
	SettlementPrice TThostFtdcPriceType
	// 涨停板价
	UpperLimitPrice TThostFtdcPriceType
	// 跌停板价
	LowerLimitPrice TThostFtdcPriceType
	// 昨虚实度
	PreDelta TThostFtdcRatioType
	// 今虚实度
	CurrDelta TThostFtdcRatioType
	// 最后修改时间
	UpdateTime TThostFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TThostFtdcMillisecType
	// 申买价一
	BidPrice1 TThostFtdcPriceType
	// 申买量一
	BidVolume1 TThostFtdcVolumeType
	// 申卖价一
	AskPrice1 TThostFtdcPriceType
	// 申卖量一
	AskVolume1 TThostFtdcVolumeType
	// 申买价二
	BidPrice2 TThostFtdcPriceType
	// 申买量二
	BidVolume2 TThostFtdcVolumeType
	// 申卖价二
	AskPrice2 TThostFtdcPriceType
	// 申卖量二
	AskVolume2 TThostFtdcVolumeType
	// 申买价三
	BidPrice3 TThostFtdcPriceType
	// 申买量三
	BidVolume3 TThostFtdcVolumeType
	// 申卖价三
	AskPrice3 TThostFtdcPriceType
	// 申卖量三
	AskVolume3 TThostFtdcVolumeType
	// 申买价四
	BidPrice4 TThostFtdcPriceType
	// 申买量四
	BidVolume4 TThostFtdcVolumeType
	// 申卖价四
	AskPrice4 TThostFtdcPriceType
	// 申卖量四
	AskVolume4 TThostFtdcVolumeType
	// 申买价五
	BidPrice5 TThostFtdcPriceType
	// 申买量五
	BidVolume5 TThostFtdcVolumeType
	// 申卖价五
	AskPrice5 TThostFtdcPriceType
	// 申卖量五
	AskVolume5 TThostFtdcVolumeType
	// 当日均价
	AveragePrice TThostFtdcPriceType
	// 业务日期
	ActionDay TThostFtdcDateType
	// 上带价
	BandingUpperPrice TThostFtdcPriceType
	// 下带价
	BandingLowerPrice TThostFtdcPriceType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaIndexPriceField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 指数现货收盘价
	ClosePrice TThostFtdcPriceType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSyncDeltaEWarrantOffsetField struct {
	// 交易日期
	TradingDay TThostFtdcTradeDateType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 买卖方向
	Direction TThostFtdcDirectionType
	// 投机套保标志
	HedgeFlag TThostFtdcHedgeFlagType
	// 数量
	Volume TThostFtdcVolumeType
	// 操作标志
	ActionDirection TThostFtdcActionDirectionType
	// 追平序号
	SyncDeltaSequenceNo TThostFtdcSequenceNoType
}

// 信息分发
type CThostFtdcSPBMFutureParameterField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
	// 期货合约因子
	Cvf TThostFtdcVolumeMultipleType
	// 阶段标识
	TimeRange TThostFtdcTimeRangeType
	// 品种保证金标准
	MarginRate TThostFtdcRatioType
	// 期货合约内部对锁仓费率折扣比例
	LockRateX TThostFtdcRatioType
	// 提高保证金标准
	AddOnRate TThostFtdcRatioType
	// 昨结算价
	PreSettlementPrice TThostFtdcPriceType
}

// 信息分发
type CThostFtdcSPBMOptionParameterField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
	// 期权合约因子
	Cvf TThostFtdcVolumeMultipleType
	// 期权冲抵价格
	DownPrice TThostFtdcPriceType
	// Delta值
	Delta TThostFtdcDeltaType
	// 卖方期权风险转换最低值
	SlimiDelta TThostFtdcDeltaType
	// 昨结算价
	PreSettlementPrice TThostFtdcPriceType
}

// 信息分发
type CThostFtdcSPBMIntraParameterField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
	// 品种内合约间对锁仓费率折扣比例
	IntraRateY TThostFtdcRatioType
}

// 信息分发
type CThostFtdcSPBMInterParameterField struct {
	// 交易日
	TradingDay TThostFtdcDateType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 优先级
	SpreadId TThostFtdcSpreadIdType
	// 品种间对锁仓费率折扣比例
	InterRateZ TThostFtdcRatioType
	// 第一腿构成品种
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSyncSPBMParameterEndField struct {
	// 交易日
	TradingDay TThostFtdcDateType
}

// 信息分发
type CThostFtdcQrySPBMFutureParameterField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQrySPBMOptionParameterField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 合约代码
	InstrumentID TThostFtdcInstrumentIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQrySPBMIntraParameterField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQrySPBMInterParameterField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 第一腿构成品种
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType
	// 第二腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcSPBMPortfDefinitionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 组合保证金套餐代码
	PortfolioDefID TThostFtdcPortfolioDefIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
	// 是否启用SPBM
	IsSPBM TThostFtdcBoolType
}

// 信息分发
type CThostFtdcSPBMInvestorPortfDefField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 组合保证金套餐代码
	PortfolioDefID TThostFtdcPortfolioDefIDType
}

// 信息分发
type CThostFtdcInvestorPortfMarginRatioField struct {
	// 投资者范围
	InvestorRange TThostFtdcInvestorRangeType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 会员对投资者收取的保证金和交易所对投资者收取的保证金的比例
	MarginRatio TThostFtdcRatioType
}

// 信息分发
type CThostFtdcQrySPBMPortfDefinitionField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 组合保证金套餐代码
	PortfolioDefID TThostFtdcPortfolioDefIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcQrySPBMInvestorPortfDefField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
}

// 信息分发
type CThostFtdcQryInvestorPortfMarginRatioField struct {
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
}

// 信息分发
type CThostFtdcInvestorProdSPBMDetailField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
	// 合约内对锁保证金
	IntraInstrMargin TThostFtdcMoneyType
	// 买归集保证金
	BCollectingMargin TThostFtdcMoneyType
	// 卖归集保证金
	SCollectingMargin TThostFtdcMoneyType
	// 品种内合约间对锁保证金
	IntraProdMargin TThostFtdcMoneyType
	// 净保证金
	NetMargin TThostFtdcMoneyType
	// 产品间对锁保证金
	InterProdMargin TThostFtdcMoneyType
	// 裸保证金
	SingleMargin TThostFtdcMoneyType
	// 附加保证金
	AddOnMargin TThostFtdcMoneyType
	// 交割月保证金
	DeliveryMargin TThostFtdcMoneyType
	// 看涨期权最低风险
	CallOptionMinRisk TThostFtdcMoneyType
	// 看跌期权最低风险
	PutOptionMinRisk TThostFtdcMoneyType
	// 卖方期权最低风险
	OptionMinRisk TThostFtdcMoneyType
	// 买方期权冲抵价值
	OptionValueOffset TThostFtdcMoneyType
	// 卖方期权权利金
	OptionRoyalty TThostFtdcMoneyType
	// 价值冲抵
	RealOptionValueOffset TThostFtdcMoneyType
	// 保证金
	Margin TThostFtdcMoneyType
	// 交易所保证金
	ExchMargin TThostFtdcMoneyType
}

// 信息分发
type CThostFtdcQryInvestorProdSPBMDetailField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 品种代码
	ProdFamilyCode TThostFtdcInstrumentIDType
}

// 信息分发
type CThostFtdcPortfTradeParamSettingField struct {
	// 交易所代码
	ExchangeID TThostFtdcExchangeIDType
	// 经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	// 投资者代码
	InvestorID TThostFtdcInvestorIDType
	// 新型组保算法
	Portfolio TThostFtdcPortfolioType
	// 撤单是否验资
	IsActionVerify TThostFtdcBoolType
	// 平仓是否验资
	IsCloseVerify TThostFtdcBoolType
}
