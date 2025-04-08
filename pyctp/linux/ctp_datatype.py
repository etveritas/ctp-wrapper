#!/usr/bin/env python3
# -*- encoding: utf-8 -*-
# LUX et VERITAS
# Create On: 2025/04/06 14:16:31

from enum import Enum
from ctypes import c_int32, c_double, c_char, c_short

# 自定义Python的Constant声明方式
class _const:
    def __setattr__(self, name, value):
        if self.__dict__.__contains__(name):
            raise Exception(f"Can't reassignment constant {name}")
        self.__dict__[name]=value

const =_const()

class THOST_TE_RESUME_TYPE(Enum):
    THOST_TERT_RESTART = 0
    THOST_TERT_RESUME = 1
    THOST_TERT_QUICK = 2
    THOST_TERT_NONE = 3


TThostFtdcExchangePropertyType = c_char
"""交易所属性类型"""

const.THOST_FTDC_EXP_Normal = "0"
"""正常"""    
const.THOST_FTDC_EXP_GenOrderByTrade = "1"
"""根据成交生成报单"""    


TThostFtdcIdCardTypeType = c_char
"""证件类型类型"""

const.THOST_FTDC_ICT_EID = "0"
"""组织机构代码"""    
const.THOST_FTDC_ICT_IDCard = "1"
"""中国公民身份证"""    
const.THOST_FTDC_ICT_OfficerIDCard = "2"
"""军官证"""    
const.THOST_FTDC_ICT_PoliceIDCard = "3"
"""警官证"""    
const.THOST_FTDC_ICT_SoldierIDCard = "4"
"""士兵证"""    
const.THOST_FTDC_ICT_HouseholdRegister = "5"
"""户口簿"""    
const.THOST_FTDC_ICT_Passport = "6"
"""护照"""    
const.THOST_FTDC_ICT_TaiwanCompatriotIDCard = "7"
"""台胞证"""    
const.THOST_FTDC_ICT_HomeComingCard = "8"
"""回乡证"""    
const.THOST_FTDC_ICT_LicenseNo = "9"
"""营业执照号"""    
const.THOST_FTDC_ICT_TaxNo = "A"
"""税务登记号/当地纳税ID"""    
const.THOST_FTDC_ICT_HMMainlandTravelPermit = "B"
"""港澳居民来往内地通行证"""    
const.THOST_FTDC_ICT_TwMainlandTravelPermit = "C"
"""台湾居民来往大陆通行证"""    
const.THOST_FTDC_ICT_DrivingLicense = "D"
"""驾照"""    
const.THOST_FTDC_ICT_SocialID = "F"
"""当地社保ID"""    
const.THOST_FTDC_ICT_LocalID = "G"
"""当地身份证"""    
const.THOST_FTDC_ICT_BusinessRegistration = "H"
"""商业登记证"""    
const.THOST_FTDC_ICT_HKMCIDCard = "I"
"""港澳永久性居民身份证"""    
const.THOST_FTDC_ICT_AccountsPermits = "J"
"""人行开户许可证"""    
const.THOST_FTDC_ICT_FrgPrmtRdCard = "K"
"""外国人永久居留证"""    
const.THOST_FTDC_ICT_CptMngPrdLetter = "L"
"""资管产品备案函"""    
const.THOST_FTDC_ICT_HKMCTwResidencePermit = "M"
"""港澳台居民居住证"""    
const.THOST_FTDC_ICT_UniformSocialCreditCode = "N"
"""统一社会信用代码"""    
const.THOST_FTDC_ICT_CorporationCertNo = "O"
"""机构成立证明文件"""    
const.THOST_FTDC_ICT_OtherCard = "x"
"""其他证件"""    


TThostFtdcInvestorRangeType = c_char
"""投资者范围类型"""

const.THOST_FTDC_IR_All = "1"
"""所有"""    
const.THOST_FTDC_IR_Group = "2"
"""投资者组"""    
const.THOST_FTDC_IR_Single = "3"
"""单一投资者"""    


TThostFtdcDepartmentRangeType = c_char
"""投资者范围类型"""

const.THOST_FTDC_DR_All = "1"
"""所有"""    
const.THOST_FTDC_DR_Group = "2"
"""组织架构"""    
const.THOST_FTDC_DR_Single = "3"
"""单一投资者"""    


TThostFtdcDataSyncStatusType = c_char
"""数据同步状态类型"""

const.THOST_FTDC_DS_Asynchronous = "1"
"""未同步"""    
const.THOST_FTDC_DS_Synchronizing = "2"
"""同步中"""    
const.THOST_FTDC_DS_Synchronized = "3"
"""已同步"""    


TThostFtdcBrokerDataSyncStatusType = c_char
"""经纪公司数据同步状态类型"""

const.THOST_FTDC_BDS_Synchronized = "1"
"""已同步"""    
const.THOST_FTDC_BDS_Synchronizing = "2"
"""同步中"""    


TThostFtdcExchangeConnectStatusType = c_char
"""交易所连接状态类型"""

const.THOST_FTDC_ECS_NoConnection = "1"
"""没有任何连接"""    
const.THOST_FTDC_ECS_QryInstrumentSent = "2"
"""已经发出合约查询请求"""    
const.THOST_FTDC_ECS_GotInformation = "9"
"""已经获取信息"""    


TThostFtdcTraderConnectStatusType = c_char
"""交易所交易员连接状态类型"""

const.THOST_FTDC_TCS_NotConnected = "1"
"""没有任何连接"""    
const.THOST_FTDC_TCS_Connected = "2"
"""已经连接"""    
const.THOST_FTDC_TCS_QryInstrumentSent = "3"
"""已经发出合约查询请求"""    
const.THOST_FTDC_TCS_SubPrivateFlow = "4"
"""订阅私有流"""    


TThostFtdcFunctionCodeType = c_char
"""功能代码类型"""

const.THOST_FTDC_FC_DataAsync = "1"
"""数据异步化"""    
const.THOST_FTDC_FC_ForceUserLogout = "2"
"""强制用户登出"""    
const.THOST_FTDC_FC_UserPasswordUpdate = "3"
"""变更管理用户口令"""    
const.THOST_FTDC_FC_BrokerPasswordUpdate = "4"
"""变更经纪公司口令"""    
const.THOST_FTDC_FC_InvestorPasswordUpdate = "5"
"""变更投资者口令"""    
const.THOST_FTDC_FC_OrderInsert = "6"
"""报单插入"""    
const.THOST_FTDC_FC_OrderAction = "7"
"""报单操作"""    
const.THOST_FTDC_FC_SyncSystemData = "8"
"""同步系统数据"""    
const.THOST_FTDC_FC_SyncBrokerData = "9"
"""同步经纪公司数据"""    
const.THOST_FTDC_FC_BachSyncBrokerData = "A"
"""批量同步经纪公司数据"""    
const.THOST_FTDC_FC_SuperQuery = "B"
"""超级查询"""    
const.THOST_FTDC_FC_ParkedOrderInsert = "C"
"""预埋报单插入"""    
const.THOST_FTDC_FC_ParkedOrderAction = "D"
"""预埋报单操作"""    
const.THOST_FTDC_FC_SyncOTP = "E"
"""同步动态令牌"""    
const.THOST_FTDC_FC_DeleteOrder = "F"
"""删除未知单"""    


TThostFtdcBrokerFunctionCodeType = c_char
"""经纪公司功能代码类型"""

const.THOST_FTDC_BFC_ForceUserLogout = "1"
"""强制用户登出"""    
const.THOST_FTDC_BFC_UserPasswordUpdate = "2"
"""变更用户口令"""    
const.THOST_FTDC_BFC_SyncBrokerData = "3"
"""同步经纪公司数据"""    
const.THOST_FTDC_BFC_BachSyncBrokerData = "4"
"""批量同步经纪公司数据"""    
const.THOST_FTDC_BFC_OrderInsert = "5"
"""报单插入"""    
const.THOST_FTDC_BFC_OrderAction = "6"
"""报单操作"""    
const.THOST_FTDC_BFC_AllQuery = "7"
"""全部查询"""    
const.THOST_FTDC_BFC_log = "a"
"""系统功能：登入/登出/修改密码等"""    
const.THOST_FTDC_BFC_BaseQry = "b"
"""基本查询：查询基础数据，如合约，交易所等常量"""    
const.THOST_FTDC_BFC_TradeQry = "c"
"""交易查询：如查成交，委托"""    
const.THOST_FTDC_BFC_Trade = "d"
"""交易功能：报单，撤单"""    
const.THOST_FTDC_BFC_Virement = "e"
"""银期转账"""    
const.THOST_FTDC_BFC_Risk = "f"
"""风险监控"""    
const.THOST_FTDC_BFC_Session = "g"
"""查询/管理：查询会话，踢人等"""    
const.THOST_FTDC_BFC_RiskNoticeCtl = "h"
"""风控通知控制"""    
const.THOST_FTDC_BFC_RiskNotice = "i"
"""风控通知发送"""    
const.THOST_FTDC_BFC_BrokerDeposit = "j"
"""察看经纪公司资金权限"""    
const.THOST_FTDC_BFC_QueryFund = "k"
"""资金查询"""    
const.THOST_FTDC_BFC_QueryOrder = "l"
"""报单查询"""    
const.THOST_FTDC_BFC_QueryTrade = "m"
"""成交查询"""    
const.THOST_FTDC_BFC_QueryPosition = "n"
"""持仓查询"""    
const.THOST_FTDC_BFC_QueryMarketData = "o"
"""行情查询"""    
const.THOST_FTDC_BFC_QueryUserEvent = "p"
"""用户事件查询"""    
const.THOST_FTDC_BFC_QueryRiskNotify = "q"
"""风险通知查询"""    
const.THOST_FTDC_BFC_QueryFundChange = "r"
"""出入金查询"""    
const.THOST_FTDC_BFC_QueryInvestor = "s"
"""投资者信息查询"""    
const.THOST_FTDC_BFC_QueryTradingCode = "t"
"""交易编码查询"""    
const.THOST_FTDC_BFC_ForceClose = "u"
"""强平"""    
const.THOST_FTDC_BFC_PressTest = "v"
"""压力测试"""    
const.THOST_FTDC_BFC_RemainCalc = "w"
"""权益反算"""    
const.THOST_FTDC_BFC_NetPositionInd = "x"
"""净持仓保证金指标"""    
const.THOST_FTDC_BFC_RiskPredict = "y"
"""风险预算"""    
const.THOST_FTDC_BFC_DataExport = "z"
"""数据导出"""    
const.THOST_FTDC_BFC_RiskTargetSetup = "A"
"""风控指标设置"""    
const.THOST_FTDC_BFC_MarketDataWarn = "B"
"""行情预警"""    
const.THOST_FTDC_BFC_QryBizNotice = "C"
"""业务通知查询"""    
const.THOST_FTDC_BFC_CfgBizNotice = "D"
"""业务通知模板设置"""    
const.THOST_FTDC_BFC_SyncOTP = "E"
"""同步动态令牌"""    
const.THOST_FTDC_BFC_SendBizNotice = "F"
"""发送业务通知"""    
const.THOST_FTDC_BFC_CfgRiskLevelStd = "G"
"""风险级别标准设置"""    
const.THOST_FTDC_BFC_TbCommand = "H"
"""交易终端应急功能"""    
const.THOST_FTDC_BFC_DeleteOrder = "J"
"""删除未知单"""    
const.THOST_FTDC_BFC_ParkedOrderInsert = "K"
"""预埋报单插入"""    
const.THOST_FTDC_BFC_ParkedOrderAction = "L"
"""预埋报单操作"""    
const.THOST_FTDC_BFC_ExecOrderNoCheck = "M"
"""资金不够仍允许行权"""    
const.THOST_FTDC_BFC_Designate = "N"
"""指定"""    
const.THOST_FTDC_BFC_StockDisposal = "O"
"""证券处置"""    
const.THOST_FTDC_BFC_BrokerDepositWarn = "Q"
"""席位资金预警"""    
const.THOST_FTDC_BFC_CoverWarn = "S"
"""备兑不足预警"""    
const.THOST_FTDC_BFC_PreExecOrder = "T"
"""行权试算"""    
const.THOST_FTDC_BFC_ExecOrderRisk = "P"
"""行权交收风险"""    
const.THOST_FTDC_BFC_PosiLimitWarn = "U"
"""持仓限额预警"""    
const.THOST_FTDC_BFC_QryPosiLimit = "V"
"""持仓限额查询"""    
const.THOST_FTDC_BFC_FBSign = "W"
"""银期签到签退"""    
const.THOST_FTDC_BFC_FBAccount = "X"
"""银期签约解约"""    


TThostFtdcOrderActionStatusType = c_char
"""报单操作状态类型"""

const.THOST_FTDC_OAS_Submitted = "a"
"""已经提交"""    
const.THOST_FTDC_OAS_Accepted = "b"
"""已经接受"""    
const.THOST_FTDC_OAS_Rejected = "c"
"""已经被拒绝"""    


TThostFtdcOrderStatusType = c_char
"""报单状态类型"""

const.THOST_FTDC_OST_AllTraded = "0"
"""全部成交"""    
const.THOST_FTDC_OST_PartTradedQueueing = "1"
"""部分成交还在队列中"""    
const.THOST_FTDC_OST_PartTradedNotQueueing = "2"
"""部分成交不在队列中"""    
const.THOST_FTDC_OST_NoTradeQueueing = "3"
"""未成交还在队列中"""    
const.THOST_FTDC_OST_NoTradeNotQueueing = "4"
"""未成交不在队列中"""    
const.THOST_FTDC_OST_Canceled = "5"
"""撤单"""    
const.THOST_FTDC_OST_Unknown = "a"
"""未知"""    
const.THOST_FTDC_OST_NotTouched = "b"
"""尚未触发"""    
const.THOST_FTDC_OST_Touched = "c"
"""已触发"""    


TThostFtdcOrderSubmitStatusType = c_char
"""报单提交状态类型"""

const.THOST_FTDC_OSS_InsertSubmitted = "0"
"""已经提交"""    
const.THOST_FTDC_OSS_CancelSubmitted = "1"
"""撤单已经提交"""    
const.THOST_FTDC_OSS_ModifySubmitted = "2"
"""修改已经提交"""    
const.THOST_FTDC_OSS_Accepted = "3"
"""已经接受"""    
const.THOST_FTDC_OSS_InsertRejected = "4"
"""报单已经被拒绝"""    
const.THOST_FTDC_OSS_CancelRejected = "5"
"""撤单已经被拒绝"""    
const.THOST_FTDC_OSS_ModifyRejected = "6"
"""改单已经被拒绝"""    


TThostFtdcPositionDateType = c_char
"""持仓日期类型"""

const.THOST_FTDC_PSD_Today = "1"
"""今日持仓"""    
const.THOST_FTDC_PSD_History = "2"
"""历史持仓"""    


TThostFtdcPositionDateTypeType = c_char
"""持仓日期类型类型"""

const.THOST_FTDC_PDT_UseHistory = "1"
"""使用历史持仓"""    
const.THOST_FTDC_PDT_NoUseHistory = "2"
"""不使用历史持仓"""    


TThostFtdcTradingRoleType = c_char
"""交易角色类型"""

const.THOST_FTDC_ER_Broker = "1"
"""代理"""    
const.THOST_FTDC_ER_Host = "2"
"""自营"""    
const.THOST_FTDC_ER_Maker = "3"
"""做市商"""    


TThostFtdcProductClassType = c_char
"""产品类型类型"""

const.THOST_FTDC_PC_Futures = "1"
"""期货"""    
const.THOST_FTDC_PC_Options = "2"
"""期货期权"""    
const.THOST_FTDC_PC_Combination = "3"
"""组合"""    
const.THOST_FTDC_PC_Spot = "4"
"""即期"""    
const.THOST_FTDC_PC_EFP = "5"
"""期转现"""    
const.THOST_FTDC_PC_SpotOption = "6"
"""现货期权"""    
const.THOST_FTDC_PC_TAS = "7"
"""TAS合约"""    
const.THOST_FTDC_PC_MI = "I"
"""金属指数"""    


TThostFtdcAPIProductClassType = c_char
"""产品类型类型"""

const.THOST_FTDC_APC_FutureSingle = "1"
"""期货单一合约"""    
const.THOST_FTDC_APC_OptionSingle = "2"
"""期权单一合约"""    
const.THOST_FTDC_APC_Futures = "3"
"""可交易期货(含期货组合和期货单一合约)"""    
const.THOST_FTDC_APC_Options = "4"
"""可交易期权(含期权组合和期权单一合约)"""    
const.THOST_FTDC_APC_TradingComb = "5"
"""可下单套利组合"""    
const.THOST_FTDC_APC_UnTradingComb = "6"
"""可申请的组合（可以申请的组合合约 包含可以交易的合约）"""    
const.THOST_FTDC_APC_AllTrading = "7"
"""所有可以交易合约"""    
const.THOST_FTDC_APC_All = "8"
"""所有合约（包含不能交易合约 慎用）"""    


TThostFtdcInstLifePhaseType = c_char
"""合约生命周期状态类型"""

const.THOST_FTDC_IP_NotStart = "0"
"""未上市"""    
const.THOST_FTDC_IP_Started = "1"
"""上市"""    
const.THOST_FTDC_IP_Pause = "2"
"""停牌"""    
const.THOST_FTDC_IP_Expired = "3"
"""到期"""    


TThostFtdcDirectionType = c_char
"""买卖方向类型"""

const.THOST_FTDC_D_Buy = "0"
"""买"""    
const.THOST_FTDC_D_Sell = "1"
"""卖"""    


TThostFtdcPositionTypeType = c_char
"""持仓类型类型"""

const.THOST_FTDC_PT_Net = "1"
"""净持仓"""    
const.THOST_FTDC_PT_Gross = "2"
"""综合持仓"""    


TThostFtdcPosiDirectionType = c_char
"""持仓多空方向类型"""

const.THOST_FTDC_PD_Net = "1"
"""净"""    
const.THOST_FTDC_PD_Long = "2"
"""多头"""    
const.THOST_FTDC_PD_Short = "3"
"""空头"""    


TThostFtdcSysSettlementStatusType = c_char
"""系统结算状态类型"""

const.THOST_FTDC_SS_NonActive = "1"
"""不活跃"""    
const.THOST_FTDC_SS_Startup = "2"
"""启动"""    
const.THOST_FTDC_SS_Operating = "3"
"""操作"""    
const.THOST_FTDC_SS_Settlement = "4"
"""结算"""    
const.THOST_FTDC_SS_SettlementFinished = "5"
"""结算完成"""    


TThostFtdcRatioAttrType = c_char
"""费率属性类型"""

const.THOST_FTDC_RA_Trade = "0"
"""交易费率"""    
const.THOST_FTDC_RA_Settlement = "1"
"""结算费率"""    


TThostFtdcHedgeFlagType = c_char
"""投机套保标志类型"""

const.THOST_FTDC_HF_Speculation = "1"
"""投机"""    
const.THOST_FTDC_HF_Arbitrage = "2"
"""套利"""    
const.THOST_FTDC_HF_Hedge = "3"
"""套保"""    
const.THOST_FTDC_HF_MarketMaker = "5"
"""做市商"""    
const.THOST_FTDC_HF_SpecHedge = "6"
"""第一腿投机第二腿套保"""    
const.THOST_FTDC_HF_HedgeSpec = "7"
"""第一腿套保第二腿投机"""    


TThostFtdcBillHedgeFlagType = c_char
"""投机套保标志类型"""

const.THOST_FTDC_BHF_Speculation = "1"
"""投机"""    
const.THOST_FTDC_BHF_Arbitrage = "2"
"""套利"""    
const.THOST_FTDC_BHF_Hedge = "3"
"""套保"""    


TThostFtdcClientIDTypeType = c_char
"""交易编码类型类型"""

const.THOST_FTDC_CIDT_Speculation = "1"
"""投机"""    
const.THOST_FTDC_CIDT_Arbitrage = "2"
"""套利"""    
const.THOST_FTDC_CIDT_Hedge = "3"
"""套保"""    
const.THOST_FTDC_CIDT_MarketMaker = "5"
"""做市商"""    


TThostFtdcOrderPriceTypeType = c_char
"""报单价格条件类型"""

const.THOST_FTDC_OPT_AnyPrice = "1"
"""任意价"""    
const.THOST_FTDC_OPT_LimitPrice = "2"
"""限价"""    
const.THOST_FTDC_OPT_BestPrice = "3"
"""最优价"""    
const.THOST_FTDC_OPT_LastPrice = "4"
"""最新价"""    
const.THOST_FTDC_OPT_LastPricePlusOneTicks = "5"
"""最新价浮动上浮1个ticks"""    
const.THOST_FTDC_OPT_LastPricePlusTwoTicks = "6"
"""最新价浮动上浮2个ticks"""    
const.THOST_FTDC_OPT_LastPricePlusThreeTicks = "7"
"""最新价浮动上浮3个ticks"""    
const.THOST_FTDC_OPT_AskPrice1 = "8"
"""卖一价"""    
const.THOST_FTDC_OPT_AskPrice1PlusOneTicks = "9"
"""卖一价浮动上浮1个ticks"""    
const.THOST_FTDC_OPT_AskPrice1PlusTwoTicks = "A"
"""卖一价浮动上浮2个ticks"""    
const.THOST_FTDC_OPT_AskPrice1PlusThreeTicks = "B"
"""卖一价浮动上浮3个ticks"""    
const.THOST_FTDC_OPT_BidPrice1 = "C"
"""买一价"""    
const.THOST_FTDC_OPT_BidPrice1PlusOneTicks = "D"
"""买一价浮动上浮1个ticks"""    
const.THOST_FTDC_OPT_BidPrice1PlusTwoTicks = "E"
"""买一价浮动上浮2个ticks"""    
const.THOST_FTDC_OPT_BidPrice1PlusThreeTicks = "F"
"""买一价浮动上浮3个ticks"""    
const.THOST_FTDC_OPT_FiveLevelPrice = "G"
"""五档价"""    


TThostFtdcOffsetFlagType = c_char
"""开平标志类型"""

const.THOST_FTDC_OF_Open = "0"
"""开仓"""    
const.THOST_FTDC_OF_Close = "1"
"""平仓"""    
const.THOST_FTDC_OF_ForceClose = "2"
"""强平"""    
const.THOST_FTDC_OF_CloseToday = "3"
"""平今"""    
const.THOST_FTDC_OF_CloseYesterday = "4"
"""平昨"""    
const.THOST_FTDC_OF_ForceOff = "5"
"""强减"""    
const.THOST_FTDC_OF_LocalForceClose = "6"
"""本地强平"""    


TThostFtdcForceCloseReasonType = c_char
"""强平原因类型"""

const.THOST_FTDC_FCC_NotForceClose = "0"
"""非强平"""    
const.THOST_FTDC_FCC_LackDeposit = "1"
"""资金不足"""    
const.THOST_FTDC_FCC_ClientOverPositionLimit = "2"
"""客户超仓"""    
const.THOST_FTDC_FCC_MemberOverPositionLimit = "3"
"""会员超仓"""    
const.THOST_FTDC_FCC_NotMultiple = "4"
"""持仓非整数倍"""    
const.THOST_FTDC_FCC_Violation = "5"
"""违规"""    
const.THOST_FTDC_FCC_Other = "6"
"""其它"""    
const.THOST_FTDC_FCC_PersonDeliv = "7"
"""自然人临近交割"""    
const.THOST_FTDC_FCC_Notverifycapital = "8"
"""风控强平不验证资金"""    


TThostFtdcOrderTypeType = c_char
"""报单类型类型"""

const.THOST_FTDC_ORDT_Normal = "0"
"""正常"""    
const.THOST_FTDC_ORDT_DeriveFromQuote = "1"
"""报价衍生"""    
const.THOST_FTDC_ORDT_DeriveFromCombination = "2"
"""组合衍生"""    
const.THOST_FTDC_ORDT_Combination = "3"
"""组合报单"""    
const.THOST_FTDC_ORDT_ConditionalOrder = "4"
"""条件单"""    
const.THOST_FTDC_ORDT_Swap = "5"
"""互换单"""    
const.THOST_FTDC_ORDT_DeriveFromBlockTrade = "6"
"""大宗交易成交衍生"""    
const.THOST_FTDC_ORDT_DeriveFromEFPTrade = "7"
"""期转现成交衍生"""    


TThostFtdcTimeConditionType = c_char
"""有效期类型类型"""

const.THOST_FTDC_TC_IOC = "1"
"""立即完成，否则撤销"""    
const.THOST_FTDC_TC_GFS = "2"
"""本节有效"""    
const.THOST_FTDC_TC_GFD = "3"
"""当日有效"""    
const.THOST_FTDC_TC_GTD = "4"
"""指定日期前有效"""    
const.THOST_FTDC_TC_GTC = "5"
"""撤销前有效"""    
const.THOST_FTDC_TC_GFA = "6"
"""集合竞价有效"""    


TThostFtdcVolumeConditionType = c_char
"""成交量类型类型"""

const.THOST_FTDC_VC_AV = "1"
"""任何数量"""    
const.THOST_FTDC_VC_MV = "2"
"""最小数量"""    
const.THOST_FTDC_VC_CV = "3"
"""全部数量"""    


TThostFtdcContingentConditionType = c_char
"""触发条件类型"""

const.THOST_FTDC_CC_Immediately = "1"
"""立即"""    
const.THOST_FTDC_CC_Touch = "2"
"""止损"""    
const.THOST_FTDC_CC_TouchProfit = "3"
"""止赢"""    
const.THOST_FTDC_CC_ParkedOrder = "4"
"""预埋单"""    
const.THOST_FTDC_CC_LastPriceGreaterThanStopPrice = "5"
"""最新价大于条件价"""    
const.THOST_FTDC_CC_LastPriceGreaterEqualStopPrice = "6"
"""最新价大于等于条件价"""    
const.THOST_FTDC_CC_LastPriceLesserThanStopPrice = "7"
"""最新价小于条件价"""    
const.THOST_FTDC_CC_LastPriceLesserEqualStopPrice = "8"
"""最新价小于等于条件价"""    
const.THOST_FTDC_CC_AskPriceGreaterThanStopPrice = "9"
"""卖一价大于条件价"""    
const.THOST_FTDC_CC_AskPriceGreaterEqualStopPrice = "A"
"""卖一价大于等于条件价"""    
const.THOST_FTDC_CC_AskPriceLesserThanStopPrice = "B"
"""卖一价小于条件价"""    
const.THOST_FTDC_CC_AskPriceLesserEqualStopPrice = "C"
"""卖一价小于等于条件价"""    
const.THOST_FTDC_CC_BidPriceGreaterThanStopPrice = "D"
"""买一价大于条件价"""    
const.THOST_FTDC_CC_BidPriceGreaterEqualStopPrice = "E"
"""买一价大于等于条件价"""    
const.THOST_FTDC_CC_BidPriceLesserThanStopPrice = "F"
"""买一价小于条件价"""    
const.THOST_FTDC_CC_BidPriceLesserEqualStopPrice = "H"
"""买一价小于等于条件价"""    


TThostFtdcActionFlagType = c_char
"""操作标志类型"""

const.THOST_FTDC_AF_Delete = "0"
"""删除"""    
const.THOST_FTDC_AF_Modify = "3"
"""修改"""    


TThostFtdcTradingRightType = c_char
"""交易权限类型"""

const.THOST_FTDC_TR_Allow = "0"
"""可以交易"""    
const.THOST_FTDC_TR_CloseOnly = "1"
"""只能平仓"""    
const.THOST_FTDC_TR_Forbidden = "2"
"""不能交易"""    


TThostFtdcOrderSourceType = c_char
"""报单来源类型"""

const.THOST_FTDC_OSRC_Participant = "0"
"""来自参与者"""    
const.THOST_FTDC_OSRC_Administrator = "1"
"""来自管理员"""    


TThostFtdcTradeTypeType = c_char
"""成交类型类型"""

const.THOST_FTDC_TRDT_SplitCombination = "#"
"""组合持仓拆分为单一持仓,初始化不应包含该类型的持仓"""    
const.THOST_FTDC_TRDT_Common = "0"
"""普通成交"""    
const.THOST_FTDC_TRDT_OptionsExecution = "1"
"""期权执行"""    
const.THOST_FTDC_TRDT_OTC = "2"
"""OTC成交"""    
const.THOST_FTDC_TRDT_EFPDerived = "3"
"""期转现衍生成交"""    
const.THOST_FTDC_TRDT_CombinationDerived = "4"
"""组合衍生成交"""    
const.THOST_FTDC_TRDT_BlockTrade = "5"
"""大宗交易成交"""    


TThostFtdcSpecPosiTypeType = c_char
"""特殊持仓明细标识类型"""

const.THOST_FTDC_SPOST_Common = "#"
"""普通持仓明细"""    
const.THOST_FTDC_SPOST_Tas = "0"
"""TAS合约成交产生的标的合约持仓明细"""    


TThostFtdcPriceSourceType = c_char
"""成交价来源类型"""

const.THOST_FTDC_PSRC_LastPrice = "0"
"""前成交价"""    
const.THOST_FTDC_PSRC_Buy = "1"
"""买委托价"""    
const.THOST_FTDC_PSRC_Sell = "2"
"""卖委托价"""    
const.THOST_FTDC_PSRC_OTC = "3"
"""场外成交价"""    


TThostFtdcInstrumentStatusType = c_char
"""合约交易状态类型"""

const.THOST_FTDC_IS_BeforeTrading = "0"
"""开盘前"""    
const.THOST_FTDC_IS_NoTrading = "1"
"""非交易"""    
const.THOST_FTDC_IS_Continous = "2"
"""连续交易"""    
const.THOST_FTDC_IS_AuctionOrdering = "3"
"""集合竞价报单"""    
const.THOST_FTDC_IS_AuctionBalance = "4"
"""集合竞价价格平衡"""    
const.THOST_FTDC_IS_AuctionMatch = "5"
"""集合竞价撮合"""    
const.THOST_FTDC_IS_Closed = "6"
"""收盘"""    


TThostFtdcInstStatusEnterReasonType = c_char
"""品种进入交易状态原因类型"""

const.THOST_FTDC_IER_Automatic = "1"
"""自动切换"""    
const.THOST_FTDC_IER_Manual = "2"
"""手动切换"""    
const.THOST_FTDC_IER_Fuse = "3"
"""熔断"""    


TThostFtdcBatchStatusType = c_char
"""处理状态类型"""

const.THOST_FTDC_BS_NoUpload = "1"
"""未上传"""    
const.THOST_FTDC_BS_Uploaded = "2"
"""已上传"""    
const.THOST_FTDC_BS_Failed = "3"
"""审核失败"""    


TThostFtdcReturnStyleType = c_char
"""按品种返还方式类型"""

const.THOST_FTDC_RS_All = "1"
"""按所有品种"""    
const.THOST_FTDC_RS_ByProduct = "2"
"""按品种"""    


TThostFtdcReturnPatternType = c_char
"""返还模式类型"""

const.THOST_FTDC_RP_ByVolume = "1"
"""按成交手数"""    
const.THOST_FTDC_RP_ByFeeOnHand = "2"
"""按留存手续费"""    


TThostFtdcReturnLevelType = c_char
"""返还级别类型"""

const.THOST_FTDC_RL_Level1 = "1"
"""级别1"""    
const.THOST_FTDC_RL_Level2 = "2"
"""级别2"""    
const.THOST_FTDC_RL_Level3 = "3"
"""级别3"""    
const.THOST_FTDC_RL_Level4 = "4"
"""级别4"""    
const.THOST_FTDC_RL_Level5 = "5"
"""级别5"""    
const.THOST_FTDC_RL_Level6 = "6"
"""级别6"""    
const.THOST_FTDC_RL_Level7 = "7"
"""级别7"""    
const.THOST_FTDC_RL_Level8 = "8"
"""级别8"""    
const.THOST_FTDC_RL_Level9 = "9"
"""级别9"""    


TThostFtdcReturnStandardType = c_char
"""返还标准类型"""

const.THOST_FTDC_RSD_ByPeriod = "1"
"""分阶段返还"""    
const.THOST_FTDC_RSD_ByStandard = "2"
"""按某一标准"""    


TThostFtdcMortgageTypeType = c_char
"""质押类型类型"""

const.THOST_FTDC_MT_Out = "0"
"""质出"""    
const.THOST_FTDC_MT_In = "1"
"""质入"""    


TThostFtdcInvestorSettlementParamIDType = c_char
"""投资者结算参数代码类型"""

const.THOST_FTDC_ISPI_MortgageRatio = "4"
"""质押比例"""    
const.THOST_FTDC_ISPI_MarginWay = "5"
"""保证金算法"""    
const.THOST_FTDC_ISPI_BillDeposit = "9"
"""结算单结存是否包含质押"""    


TThostFtdcExchangeSettlementParamIDType = c_char
"""交易所结算参数代码类型"""

const.THOST_FTDC_ESPI_MortgageRatio = "1"
"""质押比例"""    
const.THOST_FTDC_ESPI_OtherFundItem = "2"
"""分项资金导入项"""    
const.THOST_FTDC_ESPI_OtherFundImport = "3"
"""分项资金入交易所出入金"""    
const.THOST_FTDC_ESPI_CFFEXMinPrepa = "6"
"""中金所开户最低可用金额"""    
const.THOST_FTDC_ESPI_CZCESettlementType = "7"
"""郑商所结算方式"""    
const.THOST_FTDC_ESPI_ExchDelivFeeMode = "9"
"""交易所交割手续费收取方式"""    
const.THOST_FTDC_ESPI_DelivFeeMode = "0"
"""投资者交割手续费收取方式"""    
const.THOST_FTDC_ESPI_CZCEComMarginType = "A"
"""郑商所组合持仓保证金收取方式"""    
const.THOST_FTDC_ESPI_DceComMarginType = "B"
"""大商所套利保证金是否优惠"""    
const.THOST_FTDC_ESPI_OptOutDisCountRate = "a"
"""虚值期权保证金优惠比率"""    
const.THOST_FTDC_ESPI_OptMiniGuarantee = "b"
"""最低保障系数"""    


TThostFtdcSystemParamIDType = c_char
"""系统参数代码类型"""

const.THOST_FTDC_SPI_InvestorIDMinLength = "1"
"""投资者代码最小长度"""    
const.THOST_FTDC_SPI_AccountIDMinLength = "2"
"""投资者帐号代码最小长度"""    
const.THOST_FTDC_SPI_UserRightLogon = "3"
"""投资者开户默认登录权限"""    
const.THOST_FTDC_SPI_SettlementBillTrade = "4"
"""投资者交易结算单成交汇总方式"""    
const.THOST_FTDC_SPI_TradingCode = "5"
"""统一开户更新交易编码方式"""    
const.THOST_FTDC_SPI_CheckFund = "6"
"""结算是否判断存在未复核的出入金和分项资金"""    
const.THOST_FTDC_SPI_CommModelRight = "7"
"""是否启用手续费模板数据权限"""    
const.THOST_FTDC_SPI_MarginModelRight = "9"
"""是否启用保证金率模板数据权限"""    
const.THOST_FTDC_SPI_IsStandardActive = "8"
"""是否规范用户才能激活"""    
const.THOST_FTDC_SPI_UploadSettlementFile = "U"
"""上传的交易所结算文件路径"""    
const.THOST_FTDC_SPI_DownloadCSRCFile = "D"
"""上报保证金监控中心文件路径"""    
const.THOST_FTDC_SPI_SettlementBillFile = "S"
"""生成的结算单文件路径"""    
const.THOST_FTDC_SPI_CSRCOthersFile = "C"
"""证监会文件标识"""    
const.THOST_FTDC_SPI_InvestorPhoto = "P"
"""投资者照片路径"""    
const.THOST_FTDC_SPI_CSRCData = "R"
"""全结经纪公司上传文件路径"""    
const.THOST_FTDC_SPI_InvestorPwdModel = "I"
"""开户密码录入方式"""    
const.THOST_FTDC_SPI_CFFEXInvestorSettleFile = "F"
"""投资者中金所结算文件下载路径"""    
const.THOST_FTDC_SPI_InvestorIDType = "a"
"""投资者代码编码方式"""    
const.THOST_FTDC_SPI_FreezeMaxReMain = "r"
"""休眠户最高权益"""    
const.THOST_FTDC_SPI_IsSync = "A"
"""手续费相关操作实时上场开关"""    
const.THOST_FTDC_SPI_RelieveOpenLimit = "O"
"""解除开仓权限限制"""    
const.THOST_FTDC_SPI_IsStandardFreeze = "X"
"""是否规范用户才能休眠"""    
const.THOST_FTDC_SPI_CZCENormalProductHedge = "B"
"""郑商所是否开放所有品种套保交易"""    


TThostFtdcTradeParamIDType = c_char
"""交易系统参数代码类型"""

const.THOST_FTDC_TPID_EncryptionStandard = "E"
"""系统加密算法"""    
const.THOST_FTDC_TPID_RiskMode = "R"
"""系统风险算法"""    
const.THOST_FTDC_TPID_RiskModeGlobal = "G"
"""系统风险算法是否全局 0-否 1-是"""    
const.THOST_FTDC_TPID_modeEncode = "P"
"""密码加密算法"""    
const.THOST_FTDC_TPID_tickMode = "T"
"""价格小数位数参数"""    
const.THOST_FTDC_TPID_SingleUserSessionMaxNum = "S"
"""用户最大会话数"""    
const.THOST_FTDC_TPID_LoginFailMaxNum = "L"
"""最大连续登录失败数"""    
const.THOST_FTDC_TPID_IsAuthForce = "A"
"""是否强制认证"""    
const.THOST_FTDC_TPID_IsPosiFreeze = "F"
"""是否冻结证券持仓"""    
const.THOST_FTDC_TPID_IsPosiLimit = "M"
"""是否限仓"""    
const.THOST_FTDC_TPID_ForQuoteTimeInterval = "Q"
"""郑商所询价时间间隔"""    
const.THOST_FTDC_TPID_IsFuturePosiLimit = "B"
"""是否期货限仓"""    
const.THOST_FTDC_TPID_IsFutureOrderFreq = "C"
"""是否期货下单频率限制"""    
const.THOST_FTDC_TPID_IsExecOrderProfit = "H"
"""行权冻结是否计算盈利"""    
const.THOST_FTDC_TPID_IsCheckBankAcc = "I"
"""银期开户是否验证开户银行卡号是否是预留银行账户"""    
const.THOST_FTDC_TPID_PasswordDeadLine = "J"
"""弱密码最后修改日期"""    
const.THOST_FTDC_TPID_IsStrongPassword = "K"
"""强密码校验"""    
const.THOST_FTDC_TPID_BalanceMorgage = "a"
"""自有资金质押比"""    
const.THOST_FTDC_TPID_MinPwdLen = "O"
"""最小密码长度"""    
const.THOST_FTDC_TPID_LoginFailMaxNumForIP = "U"
"""IP当日最大登陆失败次数"""    
const.THOST_FTDC_TPID_PasswordPeriod = "V"
"""密码有效期"""    


TThostFtdcFileIDType = c_char
"""文件标识类型"""

const.THOST_FTDC_FI_SettlementFund = "F"
"""资金数据"""    
const.THOST_FTDC_FI_Trade = "T"
"""成交数据"""    
const.THOST_FTDC_FI_InvestorPosition = "P"
"""投资者持仓数据"""    
const.THOST_FTDC_FI_SubEntryFund = "O"
"""投资者分项资金数据"""    
const.THOST_FTDC_FI_CZCECombinationPos = "C"
"""组合持仓数据"""    
const.THOST_FTDC_FI_CSRCData = "R"
"""上报保证金监控中心数据"""    
const.THOST_FTDC_FI_CZCEClose = "L"
"""郑商所平仓了结数据"""    
const.THOST_FTDC_FI_CZCENoClose = "N"
"""郑商所非平仓了结数据"""    
const.THOST_FTDC_FI_PositionDtl = "D"
"""持仓明细数据"""    
const.THOST_FTDC_FI_OptionStrike = "S"
"""期权执行文件"""    
const.THOST_FTDC_FI_SettlementPriceComparison = "M"
"""结算价比对文件"""    
const.THOST_FTDC_FI_NonTradePosChange = "B"
"""上期所非持仓变动明细"""    


TThostFtdcFileTypeType = c_char
"""文件上传类型类型"""

const.THOST_FTDC_FUT_Settlement = "0"
"""结算"""    
const.THOST_FTDC_FUT_Check = "1"
"""核对"""    


TThostFtdcFileFormatType = c_char
"""文件格式类型"""

const.THOST_FTDC_FFT_Txt = "0"
"""文本文件(.txt)"""    
const.THOST_FTDC_FFT_Zip = "1"
"""压缩文件(.zip)"""    
const.THOST_FTDC_FFT_DBF = "2"
"""DBF文件(.dbf)"""    


TThostFtdcFileUploadStatusType = c_char
"""文件状态类型"""

const.THOST_FTDC_FUS_SucceedUpload = "1"
"""上传成功"""    
const.THOST_FTDC_FUS_FailedUpload = "2"
"""上传失败"""    
const.THOST_FTDC_FUS_SucceedLoad = "3"
"""导入成功"""    
const.THOST_FTDC_FUS_PartSucceedLoad = "4"
"""导入部分成功"""    
const.THOST_FTDC_FUS_FailedLoad = "5"
"""导入失败"""    


TThostFtdcTransferDirectionType = c_char
"""移仓方向类型"""

const.THOST_FTDC_TD_Out = "0"
"""移出"""    
const.THOST_FTDC_TD_In = "1"
"""移入"""    


TThostFtdcSpecialCreateRuleType = c_char
"""特殊的创建规则类型"""

const.THOST_FTDC_SC_NoSpecialRule = "0"
"""没有特殊创建规则"""    
const.THOST_FTDC_SC_NoSpringFestival = "1"
"""不包含春节"""    


TThostFtdcBasisPriceTypeType = c_char
"""挂牌基准价类型类型"""

const.THOST_FTDC_IPT_LastSettlement = "1"
"""上一合约结算价"""    
const.THOST_FTDC_IPT_LaseClose = "2"
"""上一合约收盘价"""    


TThostFtdcProductLifePhaseType = c_char
"""产品生命周期状态类型"""

const.THOST_FTDC_PLP_Active = "1"
"""活跃"""    
const.THOST_FTDC_PLP_NonActive = "2"
"""不活跃"""    
const.THOST_FTDC_PLP_Canceled = "3"
"""注销"""    


TThostFtdcDeliveryModeType = c_char
"""交割方式类型"""

const.THOST_FTDC_DM_CashDeliv = "1"
"""现金交割"""    
const.THOST_FTDC_DM_CommodityDeliv = "2"
"""实物交割"""    


TThostFtdcFundIOTypeType = c_char
"""出入金类型类型"""

const.THOST_FTDC_FIOT_FundIO = "1"
"""出入金"""    
const.THOST_FTDC_FIOT_Transfer = "2"
"""银期转帐"""    
const.THOST_FTDC_FIOT_SwapCurrency = "3"
"""银期换汇"""    


TThostFtdcFundTypeType = c_char
"""资金类型类型"""

const.THOST_FTDC_FT_Deposite = "1"
"""银行存款"""    
const.THOST_FTDC_FT_ItemFund = "2"
"""分项资金"""    
const.THOST_FTDC_FT_Company = "3"
"""公司调整"""    
const.THOST_FTDC_FT_InnerTransfer = "4"
"""资金内转"""    


TThostFtdcFundDirectionType = c_char
"""出入金方向类型"""

const.THOST_FTDC_FD_In = "1"
"""入金"""    
const.THOST_FTDC_FD_Out = "2"
"""出金"""    


TThostFtdcFundStatusType = c_char
"""资金状态类型"""

const.THOST_FTDC_FS_Record = "1"
"""已录入"""    
const.THOST_FTDC_FS_Check = "2"
"""已复核"""    
const.THOST_FTDC_FS_Charge = "3"
"""已冲销"""    


TThostFtdcPublishStatusType = c_char
"""发布状态类型"""

const.THOST_FTDC_PS_None = "1"
"""未发布"""    
const.THOST_FTDC_PS_Publishing = "2"
"""正在发布"""    
const.THOST_FTDC_PS_Published = "3"
"""已发布"""    


TThostFtdcSystemStatusType = c_char
"""系统状态类型"""

const.THOST_FTDC_ES_NonActive = "1"
"""不活跃"""    
const.THOST_FTDC_ES_Startup = "2"
"""启动"""    
const.THOST_FTDC_ES_Initialize = "3"
"""交易开始初始化"""    
const.THOST_FTDC_ES_Initialized = "4"
"""交易完成初始化"""    
const.THOST_FTDC_ES_Close = "5"
"""收市开始"""    
const.THOST_FTDC_ES_Closed = "6"
"""收市完成"""    
const.THOST_FTDC_ES_Settlement = "7"
"""结算"""    


TThostFtdcSettlementStatusType = c_char
"""结算状态类型"""

const.THOST_FTDC_STS_Initialize = "0"
"""初始"""    
const.THOST_FTDC_STS_Settlementing = "1"
"""结算中"""    
const.THOST_FTDC_STS_Settlemented = "2"
"""已结算"""    
const.THOST_FTDC_STS_Finished = "3"
"""结算完成"""    


TThostFtdcInvestorTypeType = c_char
"""投资者类型类型"""

const.THOST_FTDC_CT_Person = "0"
"""自然人"""    
const.THOST_FTDC_CT_Company = "1"
"""法人"""    
const.THOST_FTDC_CT_Fund = "2"
"""投资基金"""    
const.THOST_FTDC_CT_SpecialOrgan = "3"
"""特殊法人"""    
const.THOST_FTDC_CT_Asset = "4"
"""资管户"""    


TThostFtdcBrokerTypeType = c_char
"""经纪公司类型类型"""

const.THOST_FTDC_BT_Trade = "0"
"""交易会员"""    
const.THOST_FTDC_BT_TradeSettle = "1"
"""交易结算会员"""    


TThostFtdcRiskLevelType = c_char
"""风险等级类型"""

const.THOST_FTDC_FAS_Low = "1"
"""低风险客户"""    
const.THOST_FTDC_FAS_Normal = "2"
"""普通客户"""    
const.THOST_FTDC_FAS_Focus = "3"
"""关注客户"""    
const.THOST_FTDC_FAS_Risk = "4"
"""风险客户"""    


TThostFtdcFeeAcceptStyleType = c_char
"""手续费收取方式类型"""

const.THOST_FTDC_FAS_ByTrade = "1"
"""按交易收取"""    
const.THOST_FTDC_FAS_ByDeliv = "2"
"""按交割收取"""    
const.THOST_FTDC_FAS_None = "3"
"""不收"""    
const.THOST_FTDC_FAS_FixFee = "4"
"""按指定手续费收取"""    


TThostFtdcPasswordTypeType = c_char
"""密码类型类型"""

const.THOST_FTDC_PWDT_Trade = "1"
"""交易密码"""    
const.THOST_FTDC_PWDT_Account = "2"
"""资金密码"""    


TThostFtdcAlgorithmType = c_char
"""盈亏算法类型"""

const.THOST_FTDC_AG_All = "1"
"""浮盈浮亏都计算"""    
const.THOST_FTDC_AG_OnlyLost = "2"
"""浮盈不计，浮亏计"""    
const.THOST_FTDC_AG_OnlyGain = "3"
"""浮盈计，浮亏不计"""    
const.THOST_FTDC_AG_None = "4"
"""浮盈浮亏都不计算"""    


TThostFtdcIncludeCloseProfitType = c_char
"""是否包含平仓盈利类型"""

const.THOST_FTDC_ICP_Include = "0"
"""包含平仓盈利"""    
const.THOST_FTDC_ICP_NotInclude = "2"
"""不包含平仓盈利"""    


TThostFtdcAllWithoutTradeType = c_char
"""是否受可提比例限制类型"""

const.THOST_FTDC_AWT_Enable = "0"
"""无仓无成交不受可提比例限制"""    
const.THOST_FTDC_AWT_Disable = "2"
"""受可提比例限制"""    
const.THOST_FTDC_AWT_NoHoldEnable = "3"
"""无仓不受可提比例限制"""    


TThostFtdcFuturePwdFlagType = c_char
"""资金密码核对标志类型"""

const.THOST_FTDC_FPWD_UnCheck = "0"
"""不核对"""    
const.THOST_FTDC_FPWD_Check = "1"
"""核对"""    


TThostFtdcTransferTypeType = c_char
"""银期转账类型类型"""

const.THOST_FTDC_TT_BankToFuture = "0"
"""银行转期货"""    
const.THOST_FTDC_TT_FutureToBank = "1"
"""期货转银行"""    


TThostFtdcTransferValidFlagType = c_char
"""转账有效标志类型"""

const.THOST_FTDC_TVF_Invalid = "0"
"""无效或失败"""    
const.THOST_FTDC_TVF_Valid = "1"
"""有效"""    
const.THOST_FTDC_TVF_Reverse = "2"
"""冲正"""    


TThostFtdcReasonType = c_char
"""事由类型"""

const.THOST_FTDC_RN_CD = "0"
"""错单"""    
const.THOST_FTDC_RN_ZT = "1"
"""资金在途"""    
const.THOST_FTDC_RN_QT = "2"
"""其它"""    


TThostFtdcSexType = c_char
"""性别类型"""

const.THOST_FTDC_SEX_None = "0"
"""未知"""    
const.THOST_FTDC_SEX_Man = "1"
"""男"""    
const.THOST_FTDC_SEX_Woman = "2"
"""女"""    


TThostFtdcUserTypeType = c_char
"""用户类型类型"""

const.THOST_FTDC_UT_Investor = "0"
"""投资者"""    
const.THOST_FTDC_UT_Operator = "1"
"""操作员"""    
const.THOST_FTDC_UT_SuperUser = "2"
"""管理员"""    


TThostFtdcRateTypeType = c_char
"""费率类型类型"""

const.THOST_FTDC_RATETYPE_MarginRate = "2"
"""保证金率"""    


TThostFtdcNoteTypeType = c_char
"""通知类型类型"""

const.THOST_FTDC_NOTETYPE_TradeSettleBill = "1"
"""交易结算单"""    
const.THOST_FTDC_NOTETYPE_TradeSettleMonth = "2"
"""交易结算月报"""    
const.THOST_FTDC_NOTETYPE_CallMarginNotes = "3"
"""追加保证金通知书"""    
const.THOST_FTDC_NOTETYPE_ForceCloseNotes = "4"
"""强行平仓通知书"""    
const.THOST_FTDC_NOTETYPE_TradeNotes = "5"
"""成交通知书"""    
const.THOST_FTDC_NOTETYPE_DelivNotes = "6"
"""交割通知书"""    


TThostFtdcSettlementStyleType = c_char
"""结算单方式类型"""

const.THOST_FTDC_SBS_Day = "1"
"""逐日盯市"""    
const.THOST_FTDC_SBS_Volume = "2"
"""逐笔对冲"""    


TThostFtdcSettlementBillTypeType = c_char
"""结算单类型类型"""

const.THOST_FTDC_ST_Day = "0"
"""日报"""    
const.THOST_FTDC_ST_Month = "1"
"""月报"""    


TThostFtdcUserRightTypeType = c_char
"""客户权限类型类型"""

const.THOST_FTDC_URT_Logon = "1"
"""登录"""    
const.THOST_FTDC_URT_Transfer = "2"
"""银期转帐"""    
const.THOST_FTDC_URT_EMail = "3"
"""邮寄结算单"""    
const.THOST_FTDC_URT_Fax = "4"
"""传真结算单"""    
const.THOST_FTDC_URT_ConditionOrder = "5"
"""条件单"""    


TThostFtdcMarginPriceTypeType = c_char
"""保证金价格类型类型"""

const.THOST_FTDC_MPT_PreSettlementPrice = "1"
"""昨结算价"""    
const.THOST_FTDC_MPT_SettlementPrice = "2"
"""最新价"""    
const.THOST_FTDC_MPT_AveragePrice = "3"
"""成交均价"""    
const.THOST_FTDC_MPT_OpenPrice = "4"
"""开仓价"""    


TThostFtdcBillGenStatusType = c_char
"""结算单生成状态类型"""

const.THOST_FTDC_BGS_None = "0"
"""未生成"""    
const.THOST_FTDC_BGS_NoGenerated = "1"
"""生成中"""    
const.THOST_FTDC_BGS_Generated = "2"
"""已生成"""    


TThostFtdcAlgoTypeType = c_char
"""算法类型类型"""

const.THOST_FTDC_AT_HandlePositionAlgo = "1"
"""持仓处理算法"""    
const.THOST_FTDC_AT_FindMarginRateAlgo = "2"
"""寻找保证金率算法"""    


TThostFtdcHandlePositionAlgoIDType = c_char
"""持仓处理算法编号类型"""

const.THOST_FTDC_HPA_Base = "1"
"""基本"""    
const.THOST_FTDC_HPA_DCE = "2"
"""大连商品交易所"""    
const.THOST_FTDC_HPA_CZCE = "3"
"""郑州商品交易所"""    


TThostFtdcFindMarginRateAlgoIDType = c_char
"""寻找保证金率算法编号类型"""

const.THOST_FTDC_FMRA_Base = "1"
"""基本"""    
const.THOST_FTDC_FMRA_DCE = "2"
"""大连商品交易所"""    
const.THOST_FTDC_FMRA_CZCE = "3"
"""郑州商品交易所"""    


TThostFtdcHandleTradingAccountAlgoIDType = c_char
"""资金处理算法编号类型"""

const.THOST_FTDC_HTAA_Base = "1"
"""基本"""    
const.THOST_FTDC_HTAA_DCE = "2"
"""大连商品交易所"""    
const.THOST_FTDC_HTAA_CZCE = "3"
"""郑州商品交易所"""    


TThostFtdcPersonTypeType = c_char
"""联系人类型类型"""

const.THOST_FTDC_PST_Order = "1"
"""指定下单人"""    
const.THOST_FTDC_PST_Open = "2"
"""开户授权人"""    
const.THOST_FTDC_PST_Fund = "3"
"""资金调拨人"""    
const.THOST_FTDC_PST_Settlement = "4"
"""结算单确认人"""    
const.THOST_FTDC_PST_Company = "5"
"""法人"""    
const.THOST_FTDC_PST_Corporation = "6"
"""法人代表"""    
const.THOST_FTDC_PST_LinkMan = "7"
"""投资者联系人"""    
const.THOST_FTDC_PST_Ledger = "8"
"""分户管理资产负责人"""    
const.THOST_FTDC_PST_Trustee = "9"
"""托（保）管人"""    
const.THOST_FTDC_PST_TrusteeCorporation = "A"
"""托（保）管机构法人代表"""    
const.THOST_FTDC_PST_TrusteeOpen = "B"
"""托（保）管机构开户授权人"""    
const.THOST_FTDC_PST_TrusteeContact = "C"
"""托（保）管机构联系人"""    
const.THOST_FTDC_PST_ForeignerRefer = "D"
"""境外自然人参考证件"""    
const.THOST_FTDC_PST_CorporationRefer = "E"
"""法人代表参考证件"""    


TThostFtdcQueryInvestorRangeType = c_char
"""查询范围类型"""

const.THOST_FTDC_QIR_All = "1"
"""所有"""    
const.THOST_FTDC_QIR_Group = "2"
"""查询分类"""    
const.THOST_FTDC_QIR_Single = "3"
"""单一投资者"""    


TThostFtdcInvestorRiskStatusType = c_char
"""投资者风险状态类型"""

const.THOST_FTDC_IRS_Normal = "1"
"""正常"""    
const.THOST_FTDC_IRS_Warn = "2"
"""警告"""    
const.THOST_FTDC_IRS_Call = "3"
"""追保"""    
const.THOST_FTDC_IRS_Force = "4"
"""强平"""    
const.THOST_FTDC_IRS_Exception = "5"
"""异常"""    


TThostFtdcUserEventTypeType = c_char
"""用户事件类型类型"""

const.THOST_FTDC_UET_Login = "1"
"""登录"""    
const.THOST_FTDC_UET_Logout = "2"
"""登出"""    
const.THOST_FTDC_UET_Trading = "3"
"""交易成功"""    
const.THOST_FTDC_UET_TradingError = "4"
"""交易失败"""    
const.THOST_FTDC_UET_UpdatePassword = "5"
"""修改密码"""    
const.THOST_FTDC_UET_Authenticate = "6"
"""客户端认证"""    
const.THOST_FTDC_UET_SubmitSysInfo = "7"
"""终端信息上报"""    
const.THOST_FTDC_UET_Transfer = "8"
"""转账"""    
const.THOST_FTDC_UET_Other = "9"
"""其他"""    


TThostFtdcCloseStyleType = c_char
"""平仓方式类型"""

const.THOST_FTDC_ICS_Close = "0"
"""先开先平"""    
const.THOST_FTDC_ICS_CloseToday = "1"
"""先平今再平昨"""    


TThostFtdcStatModeType = c_char
"""统计方式类型"""

const.THOST_FTDC_SM_Non = "0"
"""----"""    
const.THOST_FTDC_SM_Instrument = "1"
"""按合约统计"""    
const.THOST_FTDC_SM_Product = "2"
"""按产品统计"""    
const.THOST_FTDC_SM_Investor = "3"
"""按投资者统计"""    


TThostFtdcParkedOrderStatusType = c_char
"""预埋单状态类型"""

const.THOST_FTDC_PAOS_NotSend = "1"
"""未发送"""    
const.THOST_FTDC_PAOS_Send = "2"
"""已发送"""    
const.THOST_FTDC_PAOS_Deleted = "3"
"""已删除"""    


TThostFtdcVirDealStatusType = c_char
"""处理状态类型"""

const.THOST_FTDC_VDS_Dealing = "1"
"""正在处理"""    
const.THOST_FTDC_VDS_DeaclSucceed = "2"
"""处理成功"""    


TThostFtdcOrgSystemIDType = c_char
"""原有系统代码类型"""

const.THOST_FTDC_ORGS_Standard = "0"
"""综合交易平台"""    
const.THOST_FTDC_ORGS_ESunny = "1"
"""易盛系统"""    
const.THOST_FTDC_ORGS_KingStarV6 = "2"
"""金仕达V6系统"""    


TThostFtdcVirTradeStatusType = c_char
"""交易状态类型"""

const.THOST_FTDC_VTS_NaturalDeal = "0"
"""正常处理中"""    
const.THOST_FTDC_VTS_SucceedEnd = "1"
"""成功结束"""    
const.THOST_FTDC_VTS_FailedEND = "2"
"""失败结束"""    
const.THOST_FTDC_VTS_Exception = "3"
"""异常中"""    
const.THOST_FTDC_VTS_ManualDeal = "4"
"""已人工异常处理"""    
const.THOST_FTDC_VTS_MesException = "5"
"""通讯异常 ，请人工处理"""    
const.THOST_FTDC_VTS_SysException = "6"
"""系统出错，请人工处理"""    


TThostFtdcVirBankAccTypeType = c_char
"""银行帐户类型类型"""

const.THOST_FTDC_VBAT_BankBook = "1"
"""存折"""    
const.THOST_FTDC_VBAT_BankCard = "2"
"""储蓄卡"""    
const.THOST_FTDC_VBAT_CreditCard = "3"
"""信用卡"""    


TThostFtdcVirementStatusType = c_char
"""银行帐户类型类型"""

const.THOST_FTDC_VMS_Natural = "0"
"""正常"""    
const.THOST_FTDC_VMS_Canceled = "9"
"""销户"""    


TThostFtdcVirementAvailAbilityType = c_char
"""有效标志类型"""

const.THOST_FTDC_VAA_NoAvailAbility = "0"
"""未确认"""    
const.THOST_FTDC_VAA_AvailAbility = "1"
"""有效"""    
const.THOST_FTDC_VAA_Repeal = "2"
"""冲正"""    


TThostFtdcVirementTradeCodeType = c_char
"""交易代码类型"""

const.THOST_FTDC_VTC_BankBankToFuture = "102001"
"""银行发起银行资金转期货"""    
const.THOST_FTDC_VTC_BankFutureToBank = "102002"
"""银行发起期货资金转银行"""    
const.THOST_FTDC_VTC_FutureBankToFuture = "202001"
"""期货发起银行资金转期货"""    
const.THOST_FTDC_VTC_FutureFutureToBank = "202002"
"""期货发起期货资金转银行"""    


TThostFtdcAMLGenStatusType = c_char
"""Aml生成方式类型"""

const.THOST_FTDC_GEN_Program = "0"
"""程序生成"""    
const.THOST_FTDC_GEN_HandWork = "1"
"""人工生成"""    


TThostFtdcCFMMCKeyKindType = c_char
"""动态密钥类别(保证金监管)类型"""

const.THOST_FTDC_CFMMCKK_REQUEST = "R"
"""主动请求更新"""    
const.THOST_FTDC_CFMMCKK_AUTO = "A"
"""CFMMC自动更新"""    
const.THOST_FTDC_CFMMCKK_MANUAL = "M"
"""CFMMC手动更新"""    


TThostFtdcCertificationTypeType = c_char
"""证件类型类型"""

const.THOST_FTDC_CFT_IDCard = "0"
"""身份证"""    
const.THOST_FTDC_CFT_Passport = "1"
"""护照"""    
const.THOST_FTDC_CFT_OfficerIDCard = "2"
"""军官证"""    
const.THOST_FTDC_CFT_SoldierIDCard = "3"
"""士兵证"""    
const.THOST_FTDC_CFT_HomeComingCard = "4"
"""回乡证"""    
const.THOST_FTDC_CFT_HouseholdRegister = "5"
"""户口簿"""    
const.THOST_FTDC_CFT_LicenseNo = "6"
"""营业执照号"""    
const.THOST_FTDC_CFT_InstitutionCodeCard = "7"
"""组织机构代码证"""    
const.THOST_FTDC_CFT_TempLicenseNo = "8"
"""临时营业执照号"""    
const.THOST_FTDC_CFT_NoEnterpriseLicenseNo = "9"
"""民办非企业登记证书"""    
const.THOST_FTDC_CFT_OtherCard = "x"
"""其他证件"""    
const.THOST_FTDC_CFT_SuperDepAgree = "a"
"""主管部门批文"""    


TThostFtdcFileBusinessCodeType = c_char
"""文件业务功能类型"""

const.THOST_FTDC_FBC_Others = "0"
"""其他"""    
const.THOST_FTDC_FBC_TransferDetails = "1"
"""转账交易明细对账"""    
const.THOST_FTDC_FBC_CustAccStatus = "2"
"""客户账户状态对账"""    
const.THOST_FTDC_FBC_AccountTradeDetails = "3"
"""账户类交易明细对账"""    
const.THOST_FTDC_FBC_FutureAccountChangeInfoDetails = "4"
"""期货账户信息变更明细对账"""    
const.THOST_FTDC_FBC_CustMoneyDetail = "5"
"""客户资金台账余额明细对账"""    
const.THOST_FTDC_FBC_CustCancelAccountInfo = "6"
"""客户销户结息明细对账"""    
const.THOST_FTDC_FBC_CustMoneyResult = "7"
"""客户资金余额对账结果"""    
const.THOST_FTDC_FBC_OthersExceptionResult = "8"
"""其它对账异常结果文件"""    
const.THOST_FTDC_FBC_CustInterestNetMoneyDetails = "9"
"""客户结息净额明细"""    
const.THOST_FTDC_FBC_CustMoneySendAndReceiveDetails = "a"
"""客户资金交收明细"""    
const.THOST_FTDC_FBC_CorporationMoneyTotal = "b"
"""法人存管银行资金交收汇总"""    
const.THOST_FTDC_FBC_MainbodyMoneyTotal = "c"
"""主体间资金交收汇总"""    
const.THOST_FTDC_FBC_MainPartMonitorData = "d"
"""总分平衡监管数据"""    
const.THOST_FTDC_FBC_PreparationMoney = "e"
"""存管银行备付金余额"""    
const.THOST_FTDC_FBC_BankMoneyMonitorData = "f"
"""协办存管银行资金监管数据"""    


TThostFtdcCashExchangeCodeType = c_char
"""汇钞标志类型"""

const.THOST_FTDC_CEC_Exchange = "1"
"""汇"""    
const.THOST_FTDC_CEC_Cash = "2"
"""钞"""    


TThostFtdcYesNoIndicatorType = c_char
"""是或否标识类型"""

const.THOST_FTDC_YNI_Yes = "0"
"""是"""    
const.THOST_FTDC_YNI_No = "1"
"""否"""    


TThostFtdcBanlanceTypeType = c_char
"""余额类型类型"""

const.THOST_FTDC_BLT_CurrentMoney = "0"
"""当前余额"""    
const.THOST_FTDC_BLT_UsableMoney = "1"
"""可用余额"""    
const.THOST_FTDC_BLT_FetchableMoney = "2"
"""可取余额"""    
const.THOST_FTDC_BLT_FreezeMoney = "3"
"""冻结余额"""    


TThostFtdcGenderType = c_char
"""性别类型"""

const.THOST_FTDC_GD_Unknown = "0"
"""未知状态"""    
const.THOST_FTDC_GD_Male = "1"
"""男"""    
const.THOST_FTDC_GD_Female = "2"
"""女"""    


TThostFtdcFeePayFlagType = c_char
"""费用支付标志类型"""

const.THOST_FTDC_FPF_BEN = "0"
"""由受益方支付费用"""    
const.THOST_FTDC_FPF_OUR = "1"
"""由发送方支付费用"""    
const.THOST_FTDC_FPF_SHA = "2"
"""由发送方支付发起的费用，受益方支付接受的费用"""    


TThostFtdcPassWordKeyTypeType = c_char
"""密钥类型类型"""

const.THOST_FTDC_PWKT_ExchangeKey = "0"
"""交换密钥"""    
const.THOST_FTDC_PWKT_PassWordKey = "1"
"""密码密钥"""    
const.THOST_FTDC_PWKT_MACKey = "2"
"""MAC密钥"""    
const.THOST_FTDC_PWKT_MessageKey = "3"
"""报文密钥"""    


TThostFtdcFBTPassWordTypeType = c_char
"""密码类型类型"""

const.THOST_FTDC_PWT_Query = "0"
"""查询"""    
const.THOST_FTDC_PWT_Fetch = "1"
"""取款"""    
const.THOST_FTDC_PWT_Transfer = "2"
"""转帐"""    
const.THOST_FTDC_PWT_Trade = "3"
"""交易"""    


TThostFtdcFBTEncryModeType = c_char
"""加密方式类型"""

const.THOST_FTDC_EM_NoEncry = "0"
"""不加密"""    
const.THOST_FTDC_EM_DES = "1"
"""DES"""    
const.THOST_FTDC_EM_3DES = "2"
"""3DES"""    


TThostFtdcBankRepealFlagType = c_char
"""银行冲正标志类型"""

const.THOST_FTDC_BRF_BankNotNeedRepeal = "0"
"""银行无需自动冲正"""    
const.THOST_FTDC_BRF_BankWaitingRepeal = "1"
"""银行待自动冲正"""    
const.THOST_FTDC_BRF_BankBeenRepealed = "2"
"""银行已自动冲正"""    


TThostFtdcBrokerRepealFlagType = c_char
"""期商冲正标志类型"""

const.THOST_FTDC_BRORF_BrokerNotNeedRepeal = "0"
"""期商无需自动冲正"""    
const.THOST_FTDC_BRORF_BrokerWaitingRepeal = "1"
"""期商待自动冲正"""    
const.THOST_FTDC_BRORF_BrokerBeenRepealed = "2"
"""期商已自动冲正"""    


TThostFtdcInstitutionTypeType = c_char
"""机构类别类型"""

const.THOST_FTDC_TS_Bank = "0"
"""银行"""    
const.THOST_FTDC_TS_Future = "1"
"""期商"""    
const.THOST_FTDC_TS_Store = "2"
"""券商"""    


TThostFtdcLastFragmentType = c_char
"""最后分片标志类型"""

const.THOST_FTDC_LF_Yes = "0"
"""是最后分片"""    
const.THOST_FTDC_LF_No = "1"
"""不是最后分片"""    


TThostFtdcBankAccStatusType = c_char
"""银行账户状态类型"""

const.THOST_FTDC_BAS_Normal = "0"
"""正常"""    
const.THOST_FTDC_BAS_Freeze = "1"
"""冻结"""    
const.THOST_FTDC_BAS_ReportLoss = "2"
"""挂失"""    


TThostFtdcMoneyAccountStatusType = c_char
"""资金账户状态类型"""

const.THOST_FTDC_MAS_Normal = "0"
"""正常"""    
const.THOST_FTDC_MAS_Cancel = "1"
"""销户"""    


TThostFtdcManageStatusType = c_char
"""存管状态类型"""

const.THOST_FTDC_MSS_Point = "0"
"""指定存管"""    
const.THOST_FTDC_MSS_PrePoint = "1"
"""预指定"""    
const.THOST_FTDC_MSS_CancelPoint = "2"
"""撤销指定"""    


TThostFtdcSystemTypeType = c_char
"""应用系统类型类型"""

const.THOST_FTDC_SYT_FutureBankTransfer = "0"
"""银期转帐"""    
const.THOST_FTDC_SYT_StockBankTransfer = "1"
"""银证转帐"""    
const.THOST_FTDC_SYT_TheThirdPartStore = "2"
"""第三方存管"""    


TThostFtdcTxnEndFlagType = c_char
"""银期转帐划转结果标志类型"""

const.THOST_FTDC_TEF_NormalProcessing = "0"
"""正常处理中"""    
const.THOST_FTDC_TEF_Success = "1"
"""成功结束"""    
const.THOST_FTDC_TEF_Failed = "2"
"""失败结束"""    
const.THOST_FTDC_TEF_Abnormal = "3"
"""异常中"""    
const.THOST_FTDC_TEF_ManualProcessedForException = "4"
"""已人工异常处理"""    
const.THOST_FTDC_TEF_CommuFailedNeedManualProcess = "5"
"""通讯异常 ，请人工处理"""    
const.THOST_FTDC_TEF_SysErrorNeedManualProcess = "6"
"""系统出错，请人工处理"""    


TThostFtdcProcessStatusType = c_char
"""银期转帐服务处理状态类型"""

const.THOST_FTDC_PSS_NotProcess = "0"
"""未处理"""    
const.THOST_FTDC_PSS_StartProcess = "1"
"""开始处理"""    
const.THOST_FTDC_PSS_Finished = "2"
"""处理完成"""    


TThostFtdcCustTypeType = c_char
"""客户类型类型"""

const.THOST_FTDC_CUSTT_Person = "0"
"""自然人"""    
const.THOST_FTDC_CUSTT_Institution = "1"
"""机构户"""    


TThostFtdcFBTTransferDirectionType = c_char
"""银期转帐方向类型"""

const.THOST_FTDC_FBTTD_FromBankToFuture = "1"
"""入金，银行转期货"""    
const.THOST_FTDC_FBTTD_FromFutureToBank = "2"
"""出金，期货转银行"""    


TThostFtdcOpenOrDestroyType = c_char
"""开销户类别类型"""

const.THOST_FTDC_OOD_Open = "1"
"""开户"""    
const.THOST_FTDC_OOD_Destroy = "0"
"""销户"""    


TThostFtdcAvailabilityFlagType = c_char
"""有效标志类型"""

const.THOST_FTDC_AVAF_Invalid = "0"
"""未确认"""    
const.THOST_FTDC_AVAF_Valid = "1"
"""有效"""    
const.THOST_FTDC_AVAF_Repeal = "2"
"""冲正"""    


TThostFtdcOrganTypeType = c_char
"""机构类型类型"""

const.THOST_FTDC_OT_Bank = "1"
"""银行代理"""    
const.THOST_FTDC_OT_Future = "2"
"""交易前置"""    
const.THOST_FTDC_OT_PlateForm = "9"
"""银期转帐平台管理"""    


TThostFtdcOrganLevelType = c_char
"""机构级别类型"""

const.THOST_FTDC_OL_HeadQuarters = "1"
"""银行总行或期商总部"""    
const.THOST_FTDC_OL_Branch = "2"
"""银行分中心或期货公司营业部"""    


TThostFtdcProtocalIDType = c_char
"""协议类型类型"""

const.THOST_FTDC_PID_FutureProtocal = "0"
"""期商协议"""    
const.THOST_FTDC_PID_ICBCProtocal = "1"
"""工行协议"""    
const.THOST_FTDC_PID_ABCProtocal = "2"
"""农行协议"""    
const.THOST_FTDC_PID_CBCProtocal = "3"
"""中国银行协议"""    
const.THOST_FTDC_PID_CCBProtocal = "4"
"""建行协议"""    
const.THOST_FTDC_PID_BOCOMProtocal = "5"
"""交行协议"""    
const.THOST_FTDC_PID_FBTPlateFormProtocal = "X"
"""银期转帐平台协议"""    


TThostFtdcConnectModeType = c_char
"""套接字连接方式类型"""

const.THOST_FTDC_CM_ShortConnect = "0"
"""短连接"""    
const.THOST_FTDC_CM_LongConnect = "1"
"""长连接"""    


TThostFtdcSyncModeType = c_char
"""套接字通信方式类型"""

const.THOST_FTDC_SRM_ASync = "0"
"""异步"""    
const.THOST_FTDC_SRM_Sync = "1"
"""同步"""    


TThostFtdcBankAccTypeType = c_char
"""银行帐号类型类型"""

const.THOST_FTDC_BAT_BankBook = "1"
"""银行存折"""    
const.THOST_FTDC_BAT_SavingCard = "2"
"""储蓄卡"""    
const.THOST_FTDC_BAT_CreditCard = "3"
"""信用卡"""    


TThostFtdcFutureAccTypeType = c_char
"""期货公司帐号类型类型"""

const.THOST_FTDC_FAT_BankBook = "1"
"""银行存折"""    
const.THOST_FTDC_FAT_SavingCard = "2"
"""储蓄卡"""    
const.THOST_FTDC_FAT_CreditCard = "3"
"""信用卡"""    


TThostFtdcOrganStatusType = c_char
"""接入机构状态类型"""

const.THOST_FTDC_OS_Ready = "0"
"""启用"""    
const.THOST_FTDC_OS_CheckIn = "1"
"""签到"""    
const.THOST_FTDC_OS_CheckOut = "2"
"""签退"""    
const.THOST_FTDC_OS_CheckFileArrived = "3"
"""对帐文件到达"""    
const.THOST_FTDC_OS_CheckDetail = "4"
"""对帐"""    
const.THOST_FTDC_OS_DayEndClean = "5"
"""日终清理"""    
const.THOST_FTDC_OS_Invalid = "9"
"""注销"""    


TThostFtdcCCBFeeModeType = c_char
"""建行收费模式类型"""

const.THOST_FTDC_CCBFM_ByAmount = "1"
"""按金额扣收"""    
const.THOST_FTDC_CCBFM_ByMonth = "2"
"""按月扣收"""    


TThostFtdcCommApiTypeType = c_char
"""通讯API类型类型"""

const.THOST_FTDC_CAPIT_Client = "1"
"""客户端"""    
const.THOST_FTDC_CAPIT_Server = "2"
"""服务端"""    
const.THOST_FTDC_CAPIT_UserApi = "3"
"""交易系统的UserApi"""    


TThostFtdcLinkStatusType = c_char
"""连接状态类型"""

const.THOST_FTDC_LS_Connected = "1"
"""已经连接"""    
const.THOST_FTDC_LS_Disconnected = "2"
"""没有连接"""    


TThostFtdcPwdFlagType = c_char
"""密码核对标志类型"""

const.THOST_FTDC_BPWDF_NoCheck = "0"
"""不核对"""    
const.THOST_FTDC_BPWDF_BlankCheck = "1"
"""明文核对"""    
const.THOST_FTDC_BPWDF_EncryptCheck = "2"
"""密文核对"""    


TThostFtdcSecuAccTypeType = c_char
"""期货帐号类型类型"""

const.THOST_FTDC_SAT_AccountID = "1"
"""资金帐号"""    
const.THOST_FTDC_SAT_CardID = "2"
"""资金卡号"""    
const.THOST_FTDC_SAT_SHStockholderID = "3"
"""上海股东帐号"""    
const.THOST_FTDC_SAT_SZStockholderID = "4"
"""深圳股东帐号"""    


TThostFtdcTransferStatusType = c_char
"""转账交易状态类型"""

const.THOST_FTDC_TRFS_Normal = "0"
"""正常"""    
const.THOST_FTDC_TRFS_Repealed = "1"
"""被冲正"""    


TThostFtdcSponsorTypeType = c_char
"""发起方类型"""

const.THOST_FTDC_SPTYPE_Broker = "0"
"""期商"""    
const.THOST_FTDC_SPTYPE_Bank = "1"
"""银行"""    


TThostFtdcReqRspTypeType = c_char
"""请求响应类别类型"""

const.THOST_FTDC_REQRSP_Request = "0"
"""请求"""    
const.THOST_FTDC_REQRSP_Response = "1"
"""响应"""    


TThostFtdcFBTUserEventTypeType = c_char
"""银期转帐用户事件类型类型"""

const.THOST_FTDC_FBTUET_SignIn = "0"
"""签到"""    
const.THOST_FTDC_FBTUET_FromBankToFuture = "1"
"""银行转期货"""    
const.THOST_FTDC_FBTUET_FromFutureToBank = "2"
"""期货转银行"""    
const.THOST_FTDC_FBTUET_OpenAccount = "3"
"""开户"""    
const.THOST_FTDC_FBTUET_CancelAccount = "4"
"""销户"""    
const.THOST_FTDC_FBTUET_ChangeAccount = "5"
"""变更银行账户"""    
const.THOST_FTDC_FBTUET_RepealFromBankToFuture = "6"
"""冲正银行转期货"""    
const.THOST_FTDC_FBTUET_RepealFromFutureToBank = "7"
"""冲正期货转银行"""    
const.THOST_FTDC_FBTUET_QueryBankAccount = "8"
"""查询银行账户"""    
const.THOST_FTDC_FBTUET_QueryFutureAccount = "9"
"""查询期货账户"""    
const.THOST_FTDC_FBTUET_SignOut = "A"
"""签退"""    
const.THOST_FTDC_FBTUET_SyncKey = "B"
"""密钥同步"""    
const.THOST_FTDC_FBTUET_ReserveOpenAccount = "C"
"""预约开户"""    
const.THOST_FTDC_FBTUET_CancelReserveOpenAccount = "D"
"""撤销预约开户"""    
const.THOST_FTDC_FBTUET_ReserveOpenAccountConfirm = "E"
"""预约开户确认"""    
const.THOST_FTDC_FBTUET_Other = "Z"
"""其他"""    


TThostFtdcDBOperationType = c_char
"""记录操作类型类型"""

const.THOST_FTDC_DBOP_Insert = "0"
"""插入"""    
const.THOST_FTDC_DBOP_Update = "1"
"""更新"""    
const.THOST_FTDC_DBOP_Delete = "2"
"""删除"""    


TThostFtdcSyncFlagType = c_char
"""同步标记类型"""

const.THOST_FTDC_SYNF_Yes = "0"
"""已同步"""    
const.THOST_FTDC_SYNF_No = "1"
"""未同步"""    


TThostFtdcSyncTypeType = c_char
"""同步类型类型"""

const.THOST_FTDC_SYNT_OneOffSync = "0"
"""一次同步"""    
const.THOST_FTDC_SYNT_TimerSync = "1"
"""定时同步"""    
const.THOST_FTDC_SYNT_TimerFullSync = "2"
"""定时完全同步"""    


TThostFtdcExDirectionType = c_char
"""换汇方向类型"""

const.THOST_FTDC_FBEDIR_Settlement = "0"
"""结汇"""    
const.THOST_FTDC_FBEDIR_Sale = "1"
"""售汇"""    


TThostFtdcFBEResultFlagType = c_char
"""换汇成功标志类型"""

const.THOST_FTDC_FBERES_Success = "0"
"""成功"""    
const.THOST_FTDC_FBERES_InsufficientBalance = "1"
"""账户余额不足"""    
const.THOST_FTDC_FBERES_UnknownTrading = "8"
"""交易结果未知"""    
const.THOST_FTDC_FBERES_Fail = "x"
"""失败"""    


TThostFtdcFBEExchStatusType = c_char
"""换汇交易状态类型"""

const.THOST_FTDC_FBEES_Normal = "0"
"""正常"""    
const.THOST_FTDC_FBEES_ReExchange = "1"
"""交易重发"""    


TThostFtdcFBEFileFlagType = c_char
"""换汇文件标志类型"""

const.THOST_FTDC_FBEFG_DataPackage = "0"
"""数据包"""    
const.THOST_FTDC_FBEFG_File = "1"
"""文件"""    


TThostFtdcFBEAlreadyTradeType = c_char
"""换汇已交易标志类型"""

const.THOST_FTDC_FBEAT_NotTrade = "0"
"""未交易"""    
const.THOST_FTDC_FBEAT_Trade = "1"
"""已交易"""    


TThostFtdcFBEUserEventTypeType = c_char
"""银期换汇用户事件类型类型"""

const.THOST_FTDC_FBEUET_SignIn = "0"
"""签到"""    
const.THOST_FTDC_FBEUET_Exchange = "1"
"""换汇"""    
const.THOST_FTDC_FBEUET_ReExchange = "2"
"""换汇重发"""    
const.THOST_FTDC_FBEUET_QueryBankAccount = "3"
"""银行账户查询"""    
const.THOST_FTDC_FBEUET_QueryExchDetial = "4"
"""换汇明细查询"""    
const.THOST_FTDC_FBEUET_QueryExchSummary = "5"
"""换汇汇总查询"""    
const.THOST_FTDC_FBEUET_QueryExchRate = "6"
"""换汇汇率查询"""    
const.THOST_FTDC_FBEUET_CheckBankAccount = "7"
"""对账文件通知"""    
const.THOST_FTDC_FBEUET_SignOut = "8"
"""签退"""    
const.THOST_FTDC_FBEUET_Other = "Z"
"""其他"""    


TThostFtdcFBEReqFlagType = c_char
"""换汇发送标志类型"""

const.THOST_FTDC_FBERF_UnProcessed = "0"
"""未处理"""    
const.THOST_FTDC_FBERF_WaitSend = "1"
"""等待发送"""    
const.THOST_FTDC_FBERF_SendSuccess = "2"
"""发送成功"""    
const.THOST_FTDC_FBERF_SendFailed = "3"
"""发送失败"""    
const.THOST_FTDC_FBERF_WaitReSend = "4"
"""等待重发"""    


TThostFtdcNotifyClassType = c_char
"""风险通知类型类型"""

const.THOST_FTDC_NC_NOERROR = "0"
"""正常"""    
const.THOST_FTDC_NC_Warn = "1"
"""警示"""    
const.THOST_FTDC_NC_Call = "2"
"""追保"""    
const.THOST_FTDC_NC_Force = "3"
"""强平"""    
const.THOST_FTDC_NC_CHUANCANG = "4"
"""穿仓"""    
const.THOST_FTDC_NC_Exception = "5"
"""异常"""    


TThostFtdcForceCloseTypeType = c_char
"""强平单类型类型"""

const.THOST_FTDC_FCT_Manual = "0"
"""手工强平"""    
const.THOST_FTDC_FCT_Single = "1"
"""单一投资者辅助强平"""    
const.THOST_FTDC_FCT_Group = "2"
"""批量投资者辅助强平"""    


TThostFtdcRiskNotifyMethodType = c_char
"""风险通知途径类型"""

const.THOST_FTDC_RNM_System = "0"
"""系统通知"""    
const.THOST_FTDC_RNM_SMS = "1"
"""短信通知"""    
const.THOST_FTDC_RNM_EMail = "2"
"""邮件通知"""    
const.THOST_FTDC_RNM_Manual = "3"
"""人工通知"""    


TThostFtdcRiskNotifyStatusType = c_char
"""风险通知状态类型"""

const.THOST_FTDC_RNS_NotGen = "0"
"""未生成"""    
const.THOST_FTDC_RNS_Generated = "1"
"""已生成未发送"""    
const.THOST_FTDC_RNS_SendError = "2"
"""发送失败"""    
const.THOST_FTDC_RNS_SendOk = "3"
"""已发送未接收"""    
const.THOST_FTDC_RNS_Received = "4"
"""已接收未确认"""    
const.THOST_FTDC_RNS_Confirmed = "5"
"""已确认"""    


TThostFtdcRiskUserEventType = c_char
"""风控用户操作事件类型"""

const.THOST_FTDC_RUE_ExportData = "0"
"""导出数据"""    


TThostFtdcConditionalOrderSortTypeType = c_char
"""条件单索引条件类型"""

const.THOST_FTDC_COST_LastPriceAsc = "0"
"""使用最新价升序"""    
const.THOST_FTDC_COST_LastPriceDesc = "1"
"""使用最新价降序"""    
const.THOST_FTDC_COST_AskPriceAsc = "2"
"""使用卖价升序"""    
const.THOST_FTDC_COST_AskPriceDesc = "3"
"""使用卖价降序"""    
const.THOST_FTDC_COST_BidPriceAsc = "4"
"""使用买价升序"""    
const.THOST_FTDC_COST_BidPriceDesc = "5"
"""使用买价降序"""    


TThostFtdcSendTypeType = c_char
"""报送状态类型"""

const.THOST_FTDC_UOAST_NoSend = "0"
"""未发送"""    
const.THOST_FTDC_UOAST_Sended = "1"
"""已发送"""    
const.THOST_FTDC_UOAST_Generated = "2"
"""已生成"""    
const.THOST_FTDC_UOAST_SendFail = "3"
"""报送失败"""    
const.THOST_FTDC_UOAST_Success = "4"
"""接收成功"""    
const.THOST_FTDC_UOAST_Fail = "5"
"""接收失败"""    
const.THOST_FTDC_UOAST_Cancel = "6"
"""取消报送"""    


TThostFtdcClientIDStatusType = c_char
"""交易编码状态类型"""

const.THOST_FTDC_UOACS_NoApply = "1"
"""未申请"""    
const.THOST_FTDC_UOACS_Submited = "2"
"""已提交申请"""    
const.THOST_FTDC_UOACS_Sended = "3"
"""已发送申请"""    
const.THOST_FTDC_UOACS_Success = "4"
"""完成"""    
const.THOST_FTDC_UOACS_Refuse = "5"
"""拒绝"""    
const.THOST_FTDC_UOACS_Cancel = "6"
"""已撤销编码"""    


TThostFtdcQuestionTypeType = c_char
"""特有信息类型类型"""

const.THOST_FTDC_QT_Radio = "1"
"""单选"""    
const.THOST_FTDC_QT_Option = "2"
"""多选"""    
const.THOST_FTDC_QT_Blank = "3"
"""填空"""    


TThostFtdcBusinessTypeType = c_char
"""业务类型类型"""

const.THOST_FTDC_BT_Request = "1"
"""请求"""    
const.THOST_FTDC_BT_Response = "2"
"""应答"""    
const.THOST_FTDC_BT_Notice = "3"
"""通知"""    


TThostFtdcCfmmcReturnCodeType = c_char
"""监控中心返回码类型"""

const.THOST_FTDC_CRC_Success = "0"
"""成功"""    
const.THOST_FTDC_CRC_Working = "1"
"""该客户已经有流程在处理中"""    
const.THOST_FTDC_CRC_InfoFail = "2"
"""监控中客户资料检查失败"""    
const.THOST_FTDC_CRC_IDCardFail = "3"
"""监控中实名制检查失败"""    
const.THOST_FTDC_CRC_OtherFail = "4"
"""其他错误"""    


TThostFtdcClientTypeType = c_char
"""客户类型类型"""

const.THOST_FTDC_CfMMCCT_All = "0"
"""所有"""    
const.THOST_FTDC_CfMMCCT_Person = "1"
"""个人"""    
const.THOST_FTDC_CfMMCCT_Company = "2"
"""单位"""    
const.THOST_FTDC_CfMMCCT_Other = "3"
"""其他"""    
const.THOST_FTDC_CfMMCCT_SpecialOrgan = "4"
"""特殊法人"""    
const.THOST_FTDC_CfMMCCT_Asset = "5"
"""资管户"""    


TThostFtdcExchangeIDTypeType = c_char
"""交易所编号类型"""

const.THOST_FTDC_EIDT_SHFE = "S"
"""上海期货交易所"""    
const.THOST_FTDC_EIDT_CZCE = "Z"
"""郑州商品交易所"""    
const.THOST_FTDC_EIDT_DCE = "D"
"""大连商品交易所"""    
const.THOST_FTDC_EIDT_CFFEX = "J"
"""中国金融期货交易所"""    
const.THOST_FTDC_EIDT_INE = "N"
"""上海国际能源交易中心股份有限公司"""    


TThostFtdcExClientIDTypeType = c_char
"""交易编码类型类型"""

const.THOST_FTDC_ECIDT_Hedge = "1"
"""套保"""    
const.THOST_FTDC_ECIDT_Arbitrage = "2"
"""套利"""    
const.THOST_FTDC_ECIDT_Speculation = "3"
"""投机"""    


TThostFtdcUpdateFlagType = c_char
"""更新状态类型"""

const.THOST_FTDC_UF_NoUpdate = "0"
"""未更新"""    
const.THOST_FTDC_UF_Success = "1"
"""更新全部信息成功"""    
const.THOST_FTDC_UF_Fail = "2"
"""更新全部信息失败"""    
const.THOST_FTDC_UF_TCSuccess = "3"
"""更新交易编码成功"""    
const.THOST_FTDC_UF_TCFail = "4"
"""更新交易编码失败"""    
const.THOST_FTDC_UF_Cancel = "5"
"""已丢弃"""    


TThostFtdcApplyOperateIDType = c_char
"""申请动作类型"""

const.THOST_FTDC_AOID_OpenInvestor = "1"
"""开户"""    
const.THOST_FTDC_AOID_ModifyIDCard = "2"
"""修改身份信息"""    
const.THOST_FTDC_AOID_ModifyNoIDCard = "3"
"""修改一般信息"""    
const.THOST_FTDC_AOID_ApplyTradingCode = "4"
"""申请交易编码"""    
const.THOST_FTDC_AOID_CancelTradingCode = "5"
"""撤销交易编码"""    
const.THOST_FTDC_AOID_CancelInvestor = "6"
"""销户"""    
const.THOST_FTDC_AOID_FreezeAccount = "8"
"""账户休眠"""    
const.THOST_FTDC_AOID_ActiveFreezeAccount = "9"
"""激活休眠账户"""    


TThostFtdcApplyStatusIDType = c_char
"""申请状态类型"""

const.THOST_FTDC_ASID_NoComplete = "1"
"""未补全"""    
const.THOST_FTDC_ASID_Submited = "2"
"""已提交"""    
const.THOST_FTDC_ASID_Checked = "3"
"""已审核"""    
const.THOST_FTDC_ASID_Refused = "4"
"""已拒绝"""    
const.THOST_FTDC_ASID_Deleted = "5"
"""已删除"""    


TThostFtdcSendMethodType = c_char
"""发送方式类型"""

const.THOST_FTDC_UOASM_ByAPI = "1"
"""文件发送"""    
const.THOST_FTDC_UOASM_ByFile = "2"
"""电子发送"""    


TThostFtdcEventModeType = c_char
"""操作方法类型"""

const.THOST_FTDC_EvM_ADD = "1"
"""增加"""    
const.THOST_FTDC_EvM_UPDATE = "2"
"""修改"""    
const.THOST_FTDC_EvM_DELETE = "3"
"""删除"""    
const.THOST_FTDC_EvM_CHECK = "4"
"""复核"""    
const.THOST_FTDC_EvM_COPY = "5"
"""复制"""    
const.THOST_FTDC_EvM_CANCEL = "6"
"""注销"""    
const.THOST_FTDC_EvM_Reverse = "7"
"""冲销"""    


TThostFtdcUOAAutoSendType = c_char
"""统一开户申请自动发送类型"""

const.THOST_FTDC_UOAA_ASR = "1"
"""自动发送并接收"""    
const.THOST_FTDC_UOAA_ASNR = "2"
"""自动发送，不自动接收"""    
const.THOST_FTDC_UOAA_NSAR = "3"
"""不自动发送，自动接收"""    
const.THOST_FTDC_UOAA_NSR = "4"
"""不自动发送，也不自动接收"""    


TThostFtdcFlowIDType = c_char
"""流程ID类型"""

const.THOST_FTDC_EvM_InvestorGroupFlow = "1"
"""投资者对应投资者组设置"""    
const.THOST_FTDC_EvM_InvestorRate = "2"
"""投资者手续费率设置"""    
const.THOST_FTDC_EvM_InvestorCommRateModel = "3"
"""投资者手续费率模板关系设置"""    


TThostFtdcCheckLevelType = c_char
"""复核级别类型"""

const.THOST_FTDC_CL_Zero = "0"
"""零级复核"""    
const.THOST_FTDC_CL_One = "1"
"""一级复核"""    
const.THOST_FTDC_CL_Two = "2"
"""二级复核"""    


TThostFtdcCheckStatusType = c_char
"""复核级别类型"""

const.THOST_FTDC_CHS_Init = "0"
"""未复核"""    
const.THOST_FTDC_CHS_Checking = "1"
"""复核中"""    
const.THOST_FTDC_CHS_Checked = "2"
"""已复核"""    
const.THOST_FTDC_CHS_Refuse = "3"
"""拒绝"""    
const.THOST_FTDC_CHS_Cancel = "4"
"""作废"""    


TThostFtdcUsedStatusType = c_char
"""生效状态类型"""

const.THOST_FTDC_CHU_Unused = "0"
"""未生效"""    
const.THOST_FTDC_CHU_Used = "1"
"""已生效"""    
const.THOST_FTDC_CHU_Fail = "2"
"""生效失败"""    


TThostFtdcBankAcountOriginType = c_char
"""账户来源类型"""

const.THOST_FTDC_BAO_ByAccProperty = "0"
"""手工录入"""    
const.THOST_FTDC_BAO_ByFBTransfer = "1"
"""银期转账"""    


TThostFtdcMonthBillTradeSumType = c_char
"""结算单月报成交汇总方式类型"""

const.THOST_FTDC_MBTS_ByInstrument = "0"
"""同日同合约"""    
const.THOST_FTDC_MBTS_ByDayInsPrc = "1"
"""同日同合约同价格"""    
const.THOST_FTDC_MBTS_ByDayIns = "2"
"""同合约"""    


TThostFtdcFBTTradeCodeEnumType = c_char
"""银期交易代码枚举类型"""

const.THOST_FTDC_FTC_BankLaunchBankToBroker = "102001"
"""银行发起银行转期货"""    
const.THOST_FTDC_FTC_BrokerLaunchBankToBroker = "202001"
"""期货发起银行转期货"""    
const.THOST_FTDC_FTC_BankLaunchBrokerToBank = "102002"
"""银行发起期货转银行"""    
const.THOST_FTDC_FTC_BrokerLaunchBrokerToBank = "202002"
"""期货发起期货转银行"""    


TThostFtdcOTPTypeType = c_char
"""动态令牌类型类型"""

const.THOST_FTDC_OTP_NONE = "0"
"""无动态令牌"""    
const.THOST_FTDC_OTP_TOTP = "1"
"""时间令牌"""    


TThostFtdcOTPStatusType = c_char
"""动态令牌状态类型"""

const.THOST_FTDC_OTPS_Unused = "0"
"""未使用"""    
const.THOST_FTDC_OTPS_Used = "1"
"""已使用"""    
const.THOST_FTDC_OTPS_Disuse = "2"
"""注销"""    


TThostFtdcBrokerUserTypeType = c_char
"""经济公司用户类型类型"""

const.THOST_FTDC_BUT_Investor = "1"
"""投资者"""    
const.THOST_FTDC_BUT_BrokerUser = "2"
"""操作员"""    


TThostFtdcFutureTypeType = c_char
"""期货类型类型"""

const.THOST_FTDC_FUTT_Commodity = "1"
"""商品期货"""    
const.THOST_FTDC_FUTT_Financial = "2"
"""金融期货"""    


TThostFtdcFundEventTypeType = c_char
"""资金管理操作类型类型"""

const.THOST_FTDC_FET_Restriction = "0"
"""转账限额"""    
const.THOST_FTDC_FET_TodayRestriction = "1"
"""当日转账限额"""    
const.THOST_FTDC_FET_Transfer = "2"
"""期商流水"""    
const.THOST_FTDC_FET_Credit = "3"
"""资金冻结"""    
const.THOST_FTDC_FET_InvestorWithdrawAlm = "4"
"""投资者可提资金比例"""    
const.THOST_FTDC_FET_BankRestriction = "5"
"""单个银行帐户转账限额"""    
const.THOST_FTDC_FET_Accountregister = "6"
"""银期签约账户"""    
const.THOST_FTDC_FET_ExchangeFundIO = "7"
"""交易所出入金"""    
const.THOST_FTDC_FET_InvestorFundIO = "8"
"""投资者出入金"""    


TThostFtdcAccountSourceTypeType = c_char
"""资金账户来源类型"""

const.THOST_FTDC_AST_FBTransfer = "0"
"""银期同步"""    
const.THOST_FTDC_AST_ManualEntry = "1"
"""手工录入"""    


TThostFtdcCodeSourceTypeType = c_char
"""交易编码来源类型"""

const.THOST_FTDC_CST_UnifyAccount = "0"
"""统一开户(已规范)"""    
const.THOST_FTDC_CST_ManualEntry = "1"
"""手工录入(未规范)"""    


TThostFtdcUserRangeType = c_char
"""操作员范围类型"""

const.THOST_FTDC_UR_All = "0"
"""所有"""    
const.THOST_FTDC_UR_Single = "1"
"""单一操作员"""    


TThostFtdcByGroupType = c_char
"""交易统计表按客户统计方式类型"""

const.THOST_FTDC_BG_Investor = "2"
"""按投资者统计"""    
const.THOST_FTDC_BG_Group = "1"
"""按类统计"""    


TThostFtdcTradeSumStatModeType = c_char
"""交易统计表按范围统计方式类型"""

const.THOST_FTDC_TSSM_Instrument = "1"
"""按合约统计"""    
const.THOST_FTDC_TSSM_Product = "2"
"""按产品统计"""    
const.THOST_FTDC_TSSM_Exchange = "3"
"""按交易所统计"""    


TThostFtdcExprSetModeType = c_char
"""日期表达式设置类型类型"""

const.THOST_FTDC_ESM_Relative = "1"
"""相对已有规则设置"""    
const.THOST_FTDC_ESM_Typical = "2"
"""典型设置"""    


TThostFtdcRateInvestorRangeType = c_char
"""投资者范围类型"""

const.THOST_FTDC_RIR_All = "1"
"""公司标准"""    
const.THOST_FTDC_RIR_Model = "2"
"""模板"""    
const.THOST_FTDC_RIR_Single = "3"
"""单一投资者"""    


TThostFtdcSyncDataStatusType = c_char
"""主次用系统数据同步状态类型"""

const.THOST_FTDC_SDS_Initialize = "0"
"""未同步"""    
const.THOST_FTDC_SDS_Settlementing = "1"
"""同步中"""    
const.THOST_FTDC_SDS_Settlemented = "2"
"""已同步"""    


TThostFtdcTradeSourceType = c_char
"""成交来源类型"""

const.THOST_FTDC_TSRC_NORMAL = "0"
"""来自交易所普通回报"""    
const.THOST_FTDC_TSRC_QUERY = "1"
"""来自查询"""    


TThostFtdcFlexStatModeType = c_char
"""产品合约统计方式类型"""

const.THOST_FTDC_FSM_Product = "1"
"""产品统计"""    
const.THOST_FTDC_FSM_Exchange = "2"
"""交易所统计"""    
const.THOST_FTDC_FSM_All = "3"
"""统计所有"""    


TThostFtdcByInvestorRangeType = c_char
"""投资者范围统计方式类型"""

const.THOST_FTDC_BIR_Property = "1"
"""属性统计"""    
const.THOST_FTDC_BIR_All = "2"
"""统计所有"""    


TThostFtdcPropertyInvestorRangeType = c_char
"""投资者范围类型"""

const.THOST_FTDC_PIR_All = "1"
"""所有"""    
const.THOST_FTDC_PIR_Property = "2"
"""投资者属性"""    
const.THOST_FTDC_PIR_Single = "3"
"""单一投资者"""    


TThostFtdcFileStatusType = c_char
"""文件状态类型"""

const.THOST_FTDC_FIS_NoCreate = "0"
"""未生成"""    
const.THOST_FTDC_FIS_Created = "1"
"""已生成"""    
const.THOST_FTDC_FIS_Failed = "2"
"""生成失败"""    


TThostFtdcFileGenStyleType = c_char
"""文件生成方式类型"""

const.THOST_FTDC_FGS_FileTransmit = "0"
"""下发"""    
const.THOST_FTDC_FGS_FileGen = "1"
"""生成"""    


TThostFtdcSysOperModeType = c_char
"""系统日志操作方法类型"""

const.THOST_FTDC_SoM_Add = "1"
"""增加"""    
const.THOST_FTDC_SoM_Update = "2"
"""修改"""    
const.THOST_FTDC_SoM_Delete = "3"
"""删除"""    
const.THOST_FTDC_SoM_Copy = "4"
"""复制"""    
const.THOST_FTDC_SoM_AcTive = "5"
"""激活"""    
const.THOST_FTDC_SoM_CanCel = "6"
"""注销"""    
const.THOST_FTDC_SoM_ReSet = "7"
"""重置"""    


TThostFtdcSysOperTypeType = c_char
"""系统日志操作类型类型"""

const.THOST_FTDC_SoT_UpdatePassword = "0"
"""修改操作员密码"""    
const.THOST_FTDC_SoT_UserDepartment = "1"
"""操作员组织架构关系"""    
const.THOST_FTDC_SoT_RoleManager = "2"
"""角色管理"""    
const.THOST_FTDC_SoT_RoleFunction = "3"
"""角色功能设置"""    
const.THOST_FTDC_SoT_BaseParam = "4"
"""基础参数设置"""    
const.THOST_FTDC_SoT_SetUserID = "5"
"""设置操作员"""    
const.THOST_FTDC_SoT_SetUserRole = "6"
"""用户角色设置"""    
const.THOST_FTDC_SoT_UserIpRestriction = "7"
"""用户IP限制"""    
const.THOST_FTDC_SoT_DepartmentManager = "8"
"""组织架构管理"""    
const.THOST_FTDC_SoT_DepartmentCopy = "9"
"""组织架构向查询分类复制"""    
const.THOST_FTDC_SoT_Tradingcode = "A"
"""交易编码管理"""    
const.THOST_FTDC_SoT_InvestorStatus = "B"
"""投资者状态维护"""    
const.THOST_FTDC_SoT_InvestorAuthority = "C"
"""投资者权限管理"""    
const.THOST_FTDC_SoT_PropertySet = "D"
"""属性设置"""    
const.THOST_FTDC_SoT_ReSetInvestorPasswd = "E"
"""重置投资者密码"""    
const.THOST_FTDC_SoT_InvestorPersonalityInfo = "F"
"""投资者个性信息维护"""    


TThostFtdcCSRCDataQueyTypeType = c_char
"""上报数据查询类型类型"""

const.THOST_FTDC_CSRCQ_Current = "0"
"""查询当前交易日报送的数据"""    
const.THOST_FTDC_CSRCQ_History = "1"
"""查询历史报送的代理经纪公司的数据"""    


TThostFtdcFreezeStatusType = c_char
"""休眠状态类型"""

const.THOST_FTDC_FRS_Normal = "1"
"""活跃"""    
const.THOST_FTDC_FRS_Freeze = "0"
"""休眠"""    


TThostFtdcStandardStatusType = c_char
"""规范状态类型"""

const.THOST_FTDC_STST_Standard = "0"
"""已规范"""    
const.THOST_FTDC_STST_NonStandard = "1"
"""未规范"""    


TThostFtdcRightParamTypeType = c_char
"""配置类型类型"""

const.THOST_FTDC_RPT_Freeze = "1"
"""休眠户"""    
const.THOST_FTDC_RPT_FreezeActive = "2"
"""激活休眠户"""    
const.THOST_FTDC_RPT_OpenLimit = "3"
"""开仓权限限制"""    
const.THOST_FTDC_RPT_RelieveOpenLimit = "4"
"""解除开仓权限限制"""    


TThostFtdcDataStatusType = c_char
"""反洗钱审核表数据状态类型"""

const.THOST_FTDC_AMLDS_Normal = "0"
"""正常"""    
const.THOST_FTDC_AMLDS_Deleted = "1"
"""已删除"""    


TThostFtdcAMLCheckStatusType = c_char
"""审核状态类型"""

const.THOST_FTDC_AMLCHS_Init = "0"
"""未复核"""    
const.THOST_FTDC_AMLCHS_Checking = "1"
"""复核中"""    
const.THOST_FTDC_AMLCHS_Checked = "2"
"""已复核"""    
const.THOST_FTDC_AMLCHS_RefuseReport = "3"
"""拒绝上报"""    


TThostFtdcAmlDateTypeType = c_char
"""日期类型类型"""

const.THOST_FTDC_AMLDT_DrawDay = "0"
"""检查日期"""    
const.THOST_FTDC_AMLDT_TouchDay = "1"
"""发生日期"""    


TThostFtdcAmlCheckLevelType = c_char
"""审核级别类型"""

const.THOST_FTDC_AMLCL_CheckLevel0 = "0"
"""零级审核"""    
const.THOST_FTDC_AMLCL_CheckLevel1 = "1"
"""一级审核"""    
const.THOST_FTDC_AMLCL_CheckLevel2 = "2"
"""二级审核"""    
const.THOST_FTDC_AMLCL_CheckLevel3 = "3"
"""三级审核"""    


TThostFtdcExportFileTypeType = c_char
"""导出文件类型类型"""

const.THOST_FTDC_EFT_CSV = "0"
"""CSV"""    
const.THOST_FTDC_EFT_EXCEL = "1"
"""Excel"""    
const.THOST_FTDC_EFT_DBF = "2"
"""DBF"""    


TThostFtdcSettleManagerTypeType = c_char
"""结算配置类型类型"""

const.THOST_FTDC_SMT_Before = "1"
"""结算前准备"""    
const.THOST_FTDC_SMT_Settlement = "2"
"""结算"""    
const.THOST_FTDC_SMT_After = "3"
"""结算后核对"""    
const.THOST_FTDC_SMT_Settlemented = "4"
"""结算后处理"""    


TThostFtdcSettleManagerLevelType = c_char
"""结算配置等级类型"""

const.THOST_FTDC_SML_Must = "1"
"""必要"""    
const.THOST_FTDC_SML_Alarm = "2"
"""警告"""    
const.THOST_FTDC_SML_Prompt = "3"
"""提示"""    
const.THOST_FTDC_SML_Ignore = "4"
"""不检查"""    


TThostFtdcSettleManagerGroupType = c_char
"""模块分组类型"""

const.THOST_FTDC_SMG_Exhcange = "1"
"""交易所核对"""    
const.THOST_FTDC_SMG_ASP = "2"
"""内部核对"""    
const.THOST_FTDC_SMG_CSRC = "3"
"""上报数据核对"""    


TThostFtdcLimitUseTypeType = c_char
"""保值额度使用类型类型"""

const.THOST_FTDC_LUT_Repeatable = "1"
"""可重复使用"""    
const.THOST_FTDC_LUT_Unrepeatable = "2"
"""不可重复使用"""    


TThostFtdcDataResourceType = c_char
"""数据来源类型"""

const.THOST_FTDC_DAR_Settle = "1"
"""本系统"""    
const.THOST_FTDC_DAR_Exchange = "2"
"""交易所"""    
const.THOST_FTDC_DAR_CSRC = "3"
"""报送数据"""    


TThostFtdcMarginTypeType = c_char
"""保证金类型类型"""

const.THOST_FTDC_MGT_ExchMarginRate = "0"
"""交易所保证金率"""    
const.THOST_FTDC_MGT_InstrMarginRate = "1"
"""投资者保证金率"""    
const.THOST_FTDC_MGT_InstrMarginRateTrade = "2"
"""投资者交易保证金率"""    


TThostFtdcActiveTypeType = c_char
"""生效类型类型"""

const.THOST_FTDC_ACT_Intraday = "1"
"""仅当日生效"""    
const.THOST_FTDC_ACT_Long = "2"
"""长期生效"""    


TThostFtdcMarginRateTypeType = c_char
"""冲突保证金率类型类型"""

const.THOST_FTDC_MRT_Exchange = "1"
"""交易所保证金率"""    
const.THOST_FTDC_MRT_Investor = "2"
"""投资者保证金率"""    
const.THOST_FTDC_MRT_InvestorTrade = "3"
"""投资者交易保证金率"""    


TThostFtdcBackUpStatusType = c_char
"""备份数据状态类型"""

const.THOST_FTDC_BUS_UnBak = "0"
"""未生成备份数据"""    
const.THOST_FTDC_BUS_BakUp = "1"
"""备份数据生成中"""    
const.THOST_FTDC_BUS_BakUped = "2"
"""已生成备份数据"""    
const.THOST_FTDC_BUS_BakFail = "3"
"""备份数据失败"""    


TThostFtdcInitSettlementType = c_char
"""结算初始化状态类型"""

const.THOST_FTDC_SIS_UnInitialize = "0"
"""结算初始化未开始"""    
const.THOST_FTDC_SIS_Initialize = "1"
"""结算初始化中"""    
const.THOST_FTDC_SIS_Initialized = "2"
"""结算初始化完成"""    


TThostFtdcReportStatusType = c_char
"""报表数据生成状态类型"""

const.THOST_FTDC_SRS_NoCreate = "0"
"""未生成报表数据"""    
const.THOST_FTDC_SRS_Create = "1"
"""报表数据生成中"""    
const.THOST_FTDC_SRS_Created = "2"
"""已生成报表数据"""    
const.THOST_FTDC_SRS_CreateFail = "3"
"""生成报表数据失败"""    


TThostFtdcSaveStatusType = c_char
"""数据归档状态类型"""

const.THOST_FTDC_SSS_UnSaveData = "0"
"""归档未完成"""    
const.THOST_FTDC_SSS_SaveDatad = "1"
"""归档完成"""    


TThostFtdcSettArchiveStatusType = c_char
"""结算确认数据归档状态类型"""

const.THOST_FTDC_SAS_UnArchived = "0"
"""未归档数据"""    
const.THOST_FTDC_SAS_Archiving = "1"
"""数据归档中"""    
const.THOST_FTDC_SAS_Archived = "2"
"""已归档数据"""    
const.THOST_FTDC_SAS_ArchiveFail = "3"
"""归档数据失败"""    


TThostFtdcCTPTypeType = c_char
"""CTP交易系统类型类型"""

const.THOST_FTDC_CTPT_Unkown = "0"
"""未知类型"""    
const.THOST_FTDC_CTPT_MainCenter = "1"
"""主中心"""    
const.THOST_FTDC_CTPT_BackUp = "2"
"""备中心"""    


TThostFtdcCloseDealTypeType = c_char
"""平仓处理类型类型"""

const.THOST_FTDC_CDT_Normal = "0"
"""正常"""    
const.THOST_FTDC_CDT_SpecFirst = "1"
"""投机平仓优先"""    


TThostFtdcMortgageFundUseRangeType = c_char
"""货币质押资金可用范围类型"""

const.THOST_FTDC_MFUR_None = "0"
"""不能使用"""    
const.THOST_FTDC_MFUR_Margin = "1"
"""用于保证金"""    
const.THOST_FTDC_MFUR_All = "2"
"""用于手续费、盈亏、保证金"""    
const.THOST_FTDC_MFUR_CNY3 = "3"
"""人民币方案3"""    


TThostFtdcSpecProductTypeType = c_char
"""特殊产品类型类型"""

const.THOST_FTDC_SPT_CzceHedge = "1"
"""郑商所套保产品"""    
const.THOST_FTDC_SPT_IneForeignCurrency = "2"
"""货币质押产品"""    
const.THOST_FTDC_SPT_DceOpenClose = "3"
"""大连短线开平仓产品"""    


TThostFtdcFundMortgageTypeType = c_char
"""货币质押类型类型"""

const.THOST_FTDC_FMT_Mortgage = "1"
"""质押"""    
const.THOST_FTDC_FMT_Redemption = "2"
"""解质"""    


TThostFtdcAccountSettlementParamIDType = c_char
"""投资者账户结算参数代码类型"""

const.THOST_FTDC_ASPI_BaseMargin = "1"
"""基础保证金"""    
const.THOST_FTDC_ASPI_LowestInterest = "2"
"""最低权益标准"""    


TThostFtdcFundMortDirectionType = c_char
"""货币质押方向类型"""

const.THOST_FTDC_FMD_In = "1"
"""货币质入"""    
const.THOST_FTDC_FMD_Out = "2"
"""货币质出"""    


TThostFtdcBusinessClassType = c_char
"""换汇类别类型"""

const.THOST_FTDC_BT_Profit = "0"
"""盈利"""    
const.THOST_FTDC_BT_Loss = "1"
"""亏损"""    
const.THOST_FTDC_BT_Other = "Z"
"""其他"""    


TThostFtdcSwapSourceTypeType = c_char
"""换汇数据来源类型"""

const.THOST_FTDC_SST_Manual = "0"
"""手工"""    
const.THOST_FTDC_SST_Automatic = "1"
"""自动生成"""    


TThostFtdcCurrExDirectionType = c_char
"""换汇类型类型"""

const.THOST_FTDC_CED_Settlement = "0"
"""结汇"""    
const.THOST_FTDC_CED_Sale = "1"
"""售汇"""    


TThostFtdcCurrencySwapStatusType = c_char
"""申请状态类型"""

const.THOST_FTDC_CSS_Entry = "1"
"""已录入"""    
const.THOST_FTDC_CSS_Approve = "2"
"""已审核"""    
const.THOST_FTDC_CSS_Refuse = "3"
"""已拒绝"""    
const.THOST_FTDC_CSS_Revoke = "4"
"""已撤销"""    
const.THOST_FTDC_CSS_Send = "5"
"""已发送"""    
const.THOST_FTDC_CSS_Success = "6"
"""换汇成功"""    
const.THOST_FTDC_CSS_Failure = "7"
"""换汇失败"""    


TThostFtdcReqFlagType = c_char
"""换汇发送标志类型"""

const.THOST_FTDC_REQF_NoSend = "0"
"""未发送"""    
const.THOST_FTDC_REQF_SendSuccess = "1"
"""发送成功"""    
const.THOST_FTDC_REQF_SendFailed = "2"
"""发送失败"""    
const.THOST_FTDC_REQF_WaitReSend = "3"
"""等待重发"""    


TThostFtdcResFlagType = c_char
"""换汇返回成功标志类型"""

const.THOST_FTDC_RESF_Success = "0"
"""成功"""    
const.THOST_FTDC_RESF_InsuffiCient = "1"
"""账户余额不足"""    
const.THOST_FTDC_RESF_UnKnown = "8"
"""交易结果未知"""    


TThostFtdcExStatusType = c_char
"""修改状态类型"""

const.THOST_FTDC_EXS_Before = "0"
"""修改前"""    
const.THOST_FTDC_EXS_After = "1"
"""修改后"""    


TThostFtdcClientRegionType = c_char
"""开户客户地域类型"""

const.THOST_FTDC_CR_Domestic = "1"
"""国内客户"""    
const.THOST_FTDC_CR_GMT = "2"
"""港澳台客户"""    
const.THOST_FTDC_CR_Foreign = "3"
"""国外客户"""    


TThostFtdcHasBoardType = c_char
"""是否有董事会类型"""

const.THOST_FTDC_HB_No = "0"
"""没有"""    
const.THOST_FTDC_HB_Yes = "1"
"""有"""    


TThostFtdcStartModeType = c_char
"""启动模式类型"""

const.THOST_FTDC_SM_Normal = "1"
"""正常"""    
const.THOST_FTDC_SM_Emerge = "2"
"""应急"""    
const.THOST_FTDC_SM_Restore = "3"
"""恢复"""    


TThostFtdcTemplateTypeType = c_char
"""模型类型类型"""

const.THOST_FTDC_TPT_Full = "1"
"""全量"""    
const.THOST_FTDC_TPT_Increment = "2"
"""增量"""    
const.THOST_FTDC_TPT_BackUp = "3"
"""备份"""    


TThostFtdcLoginModeType = c_char
"""登录模式类型"""

const.THOST_FTDC_LM_Trade = "0"
"""交易"""    
const.THOST_FTDC_LM_Transfer = "1"
"""转账"""    


TThostFtdcPromptTypeType = c_char
"""日历提示类型类型"""

const.THOST_FTDC_CPT_Instrument = "1"
"""合约上下市"""    
const.THOST_FTDC_CPT_Margin = "2"
"""保证金分段生效"""    


TThostFtdcHasTrusteeType = c_char
"""是否有托管人类型"""

const.THOST_FTDC_HT_Yes = "1"
"""有"""    
const.THOST_FTDC_HT_No = "0"
"""没有"""    


TThostFtdcAmTypeType = c_char
"""机构类型类型"""

const.THOST_FTDC_AMT_Bank = "1"
"""银行"""    
const.THOST_FTDC_AMT_Securities = "2"
"""证券公司"""    
const.THOST_FTDC_AMT_Fund = "3"
"""基金公司"""    
const.THOST_FTDC_AMT_Insurance = "4"
"""保险公司"""    
const.THOST_FTDC_AMT_Trust = "5"
"""信托公司"""    
const.THOST_FTDC_AMT_Other = "9"
"""其他"""    


TThostFtdcCSRCFundIOTypeType = c_char
"""出入金类型类型"""

const.THOST_FTDC_CFIOT_FundIO = "0"
"""出入金"""    
const.THOST_FTDC_CFIOT_SwapCurrency = "1"
"""银期换汇"""    


TThostFtdcCusAccountTypeType = c_char
"""结算账户类型类型"""

const.THOST_FTDC_CAT_Futures = "1"
"""期货结算账户"""    
const.THOST_FTDC_CAT_AssetmgrFuture = "2"
"""纯期货资管业务下的资管结算账户"""    
const.THOST_FTDC_CAT_AssetmgrTrustee = "3"
"""综合类资管业务下的期货资管托管账户"""    
const.THOST_FTDC_CAT_AssetmgrTransfer = "4"
"""综合类资管业务下的资金中转账户"""    


TThostFtdcLanguageTypeType = c_char
"""通知语言类型类型"""

const.THOST_FTDC_LT_Chinese = "1"
"""中文"""    
const.THOST_FTDC_LT_English = "2"
"""英文"""    


TThostFtdcAssetmgrClientTypeType = c_char
"""资产管理客户类型类型"""

const.THOST_FTDC_AMCT_Person = "1"
"""个人资管客户"""    
const.THOST_FTDC_AMCT_Organ = "2"
"""单位资管客户"""    
const.THOST_FTDC_AMCT_SpecialOrgan = "4"
"""特殊单位资管客户"""    


TThostFtdcAssetmgrTypeType = c_char
"""投资类型类型"""

const.THOST_FTDC_ASST_Futures = "3"
"""期货类"""    
const.THOST_FTDC_ASST_SpecialOrgan = "4"
"""综合类"""    


TThostFtdcCheckInstrTypeType = c_char
"""合约比较类型类型"""

const.THOST_FTDC_CIT_HasExch = "0"
"""合约交易所不存在"""    
const.THOST_FTDC_CIT_HasATP = "1"
"""合约本系统不存在"""    
const.THOST_FTDC_CIT_HasDiff = "2"
"""合约比较不一致"""    


TThostFtdcDeliveryTypeType = c_char
"""交割类型类型"""

const.THOST_FTDC_DT_HandDeliv = "1"
"""手工交割"""    
const.THOST_FTDC_DT_PersonDeliv = "2"
"""到期交割"""    


TThostFtdcMaxMarginSideAlgorithmType = c_char
"""大额单边保证金算法类型"""

const.THOST_FTDC_MMSA_NO = "0"
"""不使用大额单边保证金算法"""    
const.THOST_FTDC_MMSA_YES = "1"
"""使用大额单边保证金算法"""    


TThostFtdcDAClientTypeType = c_char
"""资产管理客户类型类型"""

const.THOST_FTDC_CACT_Person = "0"
"""自然人"""    
const.THOST_FTDC_CACT_Company = "1"
"""法人"""    
const.THOST_FTDC_CACT_Other = "2"
"""其他"""    


TThostFtdcUOAAssetmgrTypeType = c_char
"""投资类型类型"""

const.THOST_FTDC_UOAAT_Futures = "1"
"""期货类"""    
const.THOST_FTDC_UOAAT_SpecialOrgan = "2"
"""综合类"""    


TThostFtdcDirectionEnType = c_char
"""买卖方向类型"""

const.THOST_FTDC_DEN_Buy = "0"
"""Buy"""    
const.THOST_FTDC_DEN_Sell = "1"
"""Sell"""    


TThostFtdcOffsetFlagEnType = c_char
"""开平标志类型"""

const.THOST_FTDC_OFEN_Open = "0"
"""Position Opening"""    
const.THOST_FTDC_OFEN_Close = "1"
"""Position Close"""    
const.THOST_FTDC_OFEN_ForceClose = "2"
"""Forced Liquidation"""    
const.THOST_FTDC_OFEN_CloseToday = "3"
"""Close Today"""    
const.THOST_FTDC_OFEN_CloseYesterday = "4"
"""Close Prev."""    
const.THOST_FTDC_OFEN_ForceOff = "5"
"""Forced Reduction"""    
const.THOST_FTDC_OFEN_LocalForceClose = "6"
"""Local Forced Liquidation"""    


TThostFtdcHedgeFlagEnType = c_char
"""投机套保标志类型"""

const.THOST_FTDC_HFEN_Speculation = "1"
"""Speculation"""    
const.THOST_FTDC_HFEN_Arbitrage = "2"
"""Arbitrage"""    
const.THOST_FTDC_HFEN_Hedge = "3"
"""Hedge"""    


TThostFtdcFundIOTypeEnType = c_char
"""出入金类型类型"""

const.THOST_FTDC_FIOTEN_FundIO = "1"
"""Deposit/Withdrawal"""    
const.THOST_FTDC_FIOTEN_Transfer = "2"
"""Bank-Futures Transfer"""    
const.THOST_FTDC_FIOTEN_SwapCurrency = "3"
"""Bank-Futures FX Exchange"""    


TThostFtdcFundTypeEnType = c_char
"""资金类型类型"""

const.THOST_FTDC_FTEN_Deposite = "1"
"""Bank Deposit"""    
const.THOST_FTDC_FTEN_ItemFund = "2"
"""Payment/Fee"""    
const.THOST_FTDC_FTEN_Company = "3"
"""Brokerage Adj"""    
const.THOST_FTDC_FTEN_InnerTransfer = "4"
"""Internal Transfer"""    


TThostFtdcFundDirectionEnType = c_char
"""出入金方向类型"""

const.THOST_FTDC_FDEN_In = "1"
"""Deposit"""    
const.THOST_FTDC_FDEN_Out = "2"
"""Withdrawal"""    


TThostFtdcFundMortDirectionEnType = c_char
"""货币质押方向类型"""

const.THOST_FTDC_FMDEN_In = "1"
"""Pledge"""    
const.THOST_FTDC_FMDEN_Out = "2"
"""Redemption"""    


TThostFtdcOptionsTypeType = c_char
"""期权类型类型"""

const.THOST_FTDC_CP_CallOptions = "1"
"""看涨"""    
const.THOST_FTDC_CP_PutOptions = "2"
"""看跌"""    


TThostFtdcStrikeModeType = c_char
"""执行方式类型"""

const.THOST_FTDC_STM_Continental = "0"
"""欧式"""    
const.THOST_FTDC_STM_American = "1"
"""美式"""    
const.THOST_FTDC_STM_Bermuda = "2"
"""百慕大"""    


TThostFtdcStrikeTypeType = c_char
"""执行类型类型"""

const.THOST_FTDC_STT_Hedge = "0"
"""自身对冲"""    
const.THOST_FTDC_STT_Match = "1"
"""匹配执行"""    


TThostFtdcApplyTypeType = c_char
"""中金所期权放弃执行申请类型类型"""

const.THOST_FTDC_APPT_NotStrikeNum = "4"
"""不执行数量"""    


TThostFtdcGiveUpDataSourceType = c_char
"""放弃执行申请数据来源类型"""

const.THOST_FTDC_GUDS_Gen = "0"
"""系统生成"""    
const.THOST_FTDC_GUDS_Hand = "1"
"""手工添加"""    


TThostFtdcExecResultType = c_char
"""执行结果类型"""

const.THOST_FTDC_OER_NoExec = "n"
"""没有执行"""    
const.THOST_FTDC_OER_Canceled = "c"
"""已经取消"""    
const.THOST_FTDC_OER_OK = "0"
"""执行成功"""    
const.THOST_FTDC_OER_NoPosition = "1"
"""期权持仓不够"""    
const.THOST_FTDC_OER_NoDeposit = "2"
"""资金不够"""    
const.THOST_FTDC_OER_NoParticipant = "3"
"""会员不存在"""    
const.THOST_FTDC_OER_NoClient = "4"
"""客户不存在"""    
const.THOST_FTDC_OER_NoInstrument = "6"
"""合约不存在"""    
const.THOST_FTDC_OER_NoRight = "7"
"""没有执行权限"""    
const.THOST_FTDC_OER_InvalidVolume = "8"
"""不合理的数量"""    
const.THOST_FTDC_OER_NoEnoughHistoryTrade = "9"
"""没有足够的历史成交"""    
const.THOST_FTDC_OER_Unknown = "a"
"""未知"""    


TThostFtdcCombinationTypeType = c_char
"""组合类型类型"""

const.THOST_FTDC_COMBT_Future = "0"
"""期货组合"""    
const.THOST_FTDC_COMBT_BUL = "1"
"""垂直价差BUL"""    
const.THOST_FTDC_COMBT_BER = "2"
"""垂直价差BER"""    
const.THOST_FTDC_COMBT_STD = "3"
"""跨式组合"""    
const.THOST_FTDC_COMBT_STG = "4"
"""宽跨式组合"""    
const.THOST_FTDC_COMBT_PRT = "5"
"""备兑组合"""    
const.THOST_FTDC_COMBT_CAS = "6"
"""时间价差组合"""    
const.THOST_FTDC_COMBT_OPL = "7"
"""期权对锁组合"""    
const.THOST_FTDC_COMBT_BFO = "8"
"""买备兑组合"""    
const.THOST_FTDC_COMBT_BLS = "9"
"""买入期权垂直价差组合"""    
const.THOST_FTDC_COMBT_BES = "a"
"""卖出期权垂直价差组合"""    


TThostFtdcDceCombinationTypeType = c_char
"""组合类型类型"""

const.THOST_FTDC_DCECOMBT_SPL = "0"
"""期货对锁组合"""    
const.THOST_FTDC_DCECOMBT_OPL = "1"
"""期权对锁组合"""    
const.THOST_FTDC_DCECOMBT_SP = "2"
"""期货跨期组合"""    
const.THOST_FTDC_DCECOMBT_SPC = "3"
"""期货跨品种组合"""    
const.THOST_FTDC_DCECOMBT_BLS = "4"
"""买入期权垂直价差组合"""    
const.THOST_FTDC_DCECOMBT_BES = "5"
"""卖出期权垂直价差组合"""    
const.THOST_FTDC_DCECOMBT_CAS = "6"
"""期权日历价差组合"""    
const.THOST_FTDC_DCECOMBT_STD = "7"
"""期权跨式组合"""    
const.THOST_FTDC_DCECOMBT_STG = "8"
"""期权宽跨式组合"""    
const.THOST_FTDC_DCECOMBT_BFO = "9"
"""买入期货期权组合"""    
const.THOST_FTDC_DCECOMBT_SFO = "a"
"""卖出期货期权组合"""    


TThostFtdcOptionRoyaltyPriceTypeType = c_char
"""期权权利金价格类型类型"""

const.THOST_FTDC_ORPT_PreSettlementPrice = "1"
"""昨结算价"""    
const.THOST_FTDC_ORPT_OpenPrice = "4"
"""开仓价"""    
const.THOST_FTDC_ORPT_MaxPreSettlementPrice = "5"
"""最新价与昨结算价较大值"""    


TThostFtdcBalanceAlgorithmType = c_char
"""权益算法类型"""

const.THOST_FTDC_BLAG_Default = "1"
"""不计算期权市值盈亏"""    
const.THOST_FTDC_BLAG_IncludeOptValLost = "2"
"""计算期权市值亏损"""    


TThostFtdcActionTypeType = c_char
"""执行类型类型"""

const.THOST_FTDC_ACTP_Exec = "1"
"""执行"""    
const.THOST_FTDC_ACTP_Abandon = "2"
"""放弃"""    


TThostFtdcForQuoteStatusType = c_char
"""询价状态类型"""

const.THOST_FTDC_FQST_Submitted = "a"
"""已经提交"""    
const.THOST_FTDC_FQST_Accepted = "b"
"""已经接受"""    
const.THOST_FTDC_FQST_Rejected = "c"
"""已经被拒绝"""    


TThostFtdcValueMethodType = c_char
"""取值方式类型"""

const.THOST_FTDC_VM_Absolute = "0"
"""按绝对值"""    
const.THOST_FTDC_VM_Ratio = "1"
"""按比率"""    


TThostFtdcExecOrderPositionFlagType = c_char
"""期权行权后是否保留期货头寸的标记类型"""

const.THOST_FTDC_EOPF_Reserve = "0"
"""保留"""    
const.THOST_FTDC_EOPF_UnReserve = "1"
"""不保留"""    


TThostFtdcExecOrderCloseFlagType = c_char
"""期权行权后生成的头寸是否自动平仓类型"""

const.THOST_FTDC_EOCF_AutoClose = "0"
"""自动平仓"""    
const.THOST_FTDC_EOCF_NotToClose = "1"
"""免于自动平仓"""    


TThostFtdcProductTypeType = c_char
"""产品类型类型"""

const.THOST_FTDC_PTE_Futures = "1"
"""期货"""    
const.THOST_FTDC_PTE_Options = "2"
"""期权"""    


TThostFtdcCZCEUploadFileNameType = c_char
"""郑商所结算文件名类型"""

const.THOST_FTDC_CUFN_CUFN_O = "O"
"""^\d{8}_zz_\d{4}"""    
const.THOST_FTDC_CUFN_CUFN_T = "T"
"""^\d{8}成交表"""    
const.THOST_FTDC_CUFN_CUFN_P = "P"
"""^\d{8}单腿持仓表new"""    
const.THOST_FTDC_CUFN_CUFN_N = "N"
"""^\d{8}非平仓了结表"""    
const.THOST_FTDC_CUFN_CUFN_L = "L"
"""^\d{8}平仓表"""    
const.THOST_FTDC_CUFN_CUFN_F = "F"
"""^\d{8}资金表"""    
const.THOST_FTDC_CUFN_CUFN_C = "C"
"""^\d{8}组合持仓表"""    
const.THOST_FTDC_CUFN_CUFN_M = "M"
"""^\d{8}保证金参数表"""    


TThostFtdcDCEUploadFileNameType = c_char
"""大商所结算文件名类型"""

const.THOST_FTDC_DUFN_DUFN_O = "O"
"""^\d{8}_dl_\d{3}"""    
const.THOST_FTDC_DUFN_DUFN_T = "T"
"""^\d{8}_成交表"""    
const.THOST_FTDC_DUFN_DUFN_P = "P"
"""^\d{8}_持仓表"""    
const.THOST_FTDC_DUFN_DUFN_F = "F"
"""^\d{8}_资金结算表"""    
const.THOST_FTDC_DUFN_DUFN_C = "C"
"""^\d{8}_优惠组合持仓明细表"""    
const.THOST_FTDC_DUFN_DUFN_D = "D"
"""^\d{8}_持仓明细表"""    
const.THOST_FTDC_DUFN_DUFN_M = "M"
"""^\d{8}_保证金参数表"""    
const.THOST_FTDC_DUFN_DUFN_S = "S"
"""^\d{8}_期权执行表"""    


TThostFtdcSHFEUploadFileNameType = c_char
"""上期所结算文件名类型"""

const.THOST_FTDC_SUFN_SUFN_O = "O"
"""^\d{4}_\d{8}_\d{8}_DailyFundChg"""    
const.THOST_FTDC_SUFN_SUFN_T = "T"
"""^\d{4}_\d{8}_\d{8}_Trade"""    
const.THOST_FTDC_SUFN_SUFN_P = "P"
"""^\d{4}_\d{8}_\d{8}_SettlementDetail"""    
const.THOST_FTDC_SUFN_SUFN_F = "F"
"""^\d{4}_\d{8}_\d{8}_Capital"""    


TThostFtdcCFFEXUploadFileNameType = c_char
"""中金所结算文件名类型"""

const.THOST_FTDC_CFUFN_SUFN_T = "T"
"""^\d{4}_SG\d{1}_\d{8}_\d{1}_Trade"""    
const.THOST_FTDC_CFUFN_SUFN_P = "P"
"""^\d{4}_SG\d{1}_\d{8}_\d{1}_SettlementDetail"""    
const.THOST_FTDC_CFUFN_SUFN_F = "F"
"""^\d{4}_SG\d{1}_\d{8}_\d{1}_Capital"""    
const.THOST_FTDC_CFUFN_SUFN_S = "S"
"""^\d{4}_SG\d{1}_\d{8}_\d{1}_OptionExec"""    


TThostFtdcCombDirectionType = c_char
"""组合指令方向类型"""

const.THOST_FTDC_CMDR_Comb = "0"
"""申请组合"""    
const.THOST_FTDC_CMDR_UnComb = "1"
"""申请拆分"""    
const.THOST_FTDC_CMDR_DelComb = "2"
"""操作员删组合单"""    


TThostFtdcStrikeOffsetTypeType = c_char
"""行权偏移类型类型"""

const.THOST_FTDC_STOV_RealValue = "1"
"""实值额"""    
const.THOST_FTDC_STOV_ProfitValue = "2"
"""盈利额"""    
const.THOST_FTDC_STOV_RealRatio = "3"
"""实值比例"""    
const.THOST_FTDC_STOV_ProfitRatio = "4"
"""盈利比例"""    


TThostFtdcReserveOpenAccStasType = c_char
"""预约开户状态类型"""

const.THOST_FTDC_ROAST_Processing = "0"
"""等待处理中"""    
const.THOST_FTDC_ROAST_Cancelled = "1"
"""已撤销"""    
const.THOST_FTDC_ROAST_Opened = "2"
"""已开户"""    
const.THOST_FTDC_ROAST_Invalid = "3"
"""无效请求"""    


TThostFtdcWeakPasswordSourceType = c_char
"""弱密码来源类型"""

const.THOST_FTDC_WPSR_Lib = "1"
"""弱密码库"""    
const.THOST_FTDC_WPSR_Manual = "2"
"""手工录入"""    


TThostFtdcOptSelfCloseFlagType = c_char
"""期权行权的头寸是否自对冲类型"""

const.THOST_FTDC_OSCF_CloseSelfOptionPosition = "1"
"""自对冲期权仓位"""    
const.THOST_FTDC_OSCF_ReserveOptionPosition = "2"
"""保留期权仓位"""    
const.THOST_FTDC_OSCF_SellCloseSelfFuturePosition = "3"
"""自对冲卖方履约后的期货仓位"""    
const.THOST_FTDC_OSCF_ReserveFuturePosition = "4"
"""保留卖方履约后的期货仓位"""    


TThostFtdcBizTypeType = c_char
"""业务类型类型"""

const.THOST_FTDC_BZTP_Future = "1"
"""期货"""    
const.THOST_FTDC_BZTP_Stock = "2"
"""证券"""    


TThostFtdcAppTypeType = c_char
"""用户App类型类型"""

const.THOST_FTDC_APP_TYPE_Investor = "1"
"""直连的投资者"""    
const.THOST_FTDC_APP_TYPE_InvestorRelay = "2"
"""为每个投资者都创建连接的中继"""    
const.THOST_FTDC_APP_TYPE_OperatorRelay = "3"
"""所有投资者共享一个操作员连接的中继"""    
const.THOST_FTDC_APP_TYPE_UnKnown = "4"
"""未知"""    


TThostFtdcResponseValueType = c_char
"""应答类型类型"""

const.THOST_FTDC_RV_Right = "0"
"""检查成功"""    
const.THOST_FTDC_RV_Refuse = "1"
"""检查失败"""    


TThostFtdcOTCTradeTypeType = c_char
"""OTC成交类型类型"""

const.THOST_FTDC_OTC_TRDT_Block = "0"
"""大宗交易"""    
const.THOST_FTDC_OTC_TRDT_EFP = "1"
"""期转现"""    


TThostFtdcMatchTypeType = c_char
"""期现风险匹配方式类型"""

const.THOST_FTDC_OTC_MT_DV01 = "1"
"""基点价值"""    
const.THOST_FTDC_OTC_MT_ParValue = "2"
"""面值"""    


TThostFtdcAuthTypeType = c_char
"""用户终端认证方式类型"""

const.THOST_FTDC_AU_WHITE = "0"
"""白名单校验"""    
const.THOST_FTDC_AU_BLACK = "1"
"""黑名单校验"""    


TThostFtdcClassTypeType = c_char
"""合约分类方式类型"""

const.THOST_FTDC_INS_ALL = "0"
"""所有合约"""    
const.THOST_FTDC_INS_FUTURE = "1"
"""期货、即期、期转现、Tas、金属指数合约"""    
const.THOST_FTDC_INS_OPTION = "2"
"""期货、现货期权合约"""    
const.THOST_FTDC_INS_COMB = "3"
"""组合合约"""    


TThostFtdcTradingTypeType = c_char
"""合约交易状态分类方式类型"""

const.THOST_FTDC_TD_ALL = "0"
"""所有状态"""    
const.THOST_FTDC_TD_TRADE = "1"
"""交易"""    
const.THOST_FTDC_TD_UNTRADE = "2"
"""非交易"""    


TThostFtdcProductStatusType = c_char
"""产品状态类型"""

const.THOST_FTDC_PS_tradeable = "1"
"""可交易"""    
const.THOST_FTDC_PS_untradeable = "2"
"""不可交易"""    


TThostFtdcSyncDeltaStatusType = c_char
"""追平状态类型"""

const.THOST_FTDC_SDS_Readable = "1"
"""交易可读"""    
const.THOST_FTDC_SDS_Reading = "2"
"""交易在读"""    
const.THOST_FTDC_SDS_Readend = "3"
"""交易读取完成"""    
const.THOST_FTDC_SDS_OptErr = "e"
"""追平失败 交易本地状态结算不存在"""    


TThostFtdcActionDirectionType = c_char
"""操作标志类型"""

const.THOST_FTDC_ACD_Add = "1"
"""增加"""    
const.THOST_FTDC_ACD_Del = "2"
"""删除"""    
const.THOST_FTDC_ACD_Upd = "3"
"""更新"""    


TThostFtdcOrderCancelAlgType = c_char
"""撤单时选择席位算法类型"""

const.THOST_FTDC_OAC_Balance = "1"
"""轮询席位撤单"""    
const.THOST_FTDC_OAC_OrigFirst = "2"
"""优先原报单席位撤单"""    


TThostFtdcOpenLimitControlLevelType = c_char
"""开仓量限制粒度类型"""

const.THOST_FTDC_PLCL_None = "0"
"""不控制"""    
const.THOST_FTDC_PLCL_Product = "1"
"""产品级别"""    
const.THOST_FTDC_PLCL_Inst = "2"
"""合约级别"""    


TThostFtdcOrderFreqControlLevelType = c_char
"""报单频率控制粒度类型"""

const.THOST_FTDC_OFCL_None = "0"
"""不控制"""    
const.THOST_FTDC_OFCL_Product = "1"
"""产品级别"""    
const.THOST_FTDC_OFCL_Inst = "2"
"""合约级别"""    


TThostFtdcEnumBoolType = c_char
"""枚举bool类型类型"""

const.THOST_FTDC_EBL_False = "0"
"""false"""    
const.THOST_FTDC_EBL_True = "1"
"""true"""    


TThostFtdcTimeRangeType = c_char
"""期货合约阶段标识类型"""

const.THOST_FTDC_ETR_USUAL = "1"
"""一般月份"""    
const.THOST_FTDC_ETR_FNSP = "2"
"""交割月前一个月上半月"""    
const.THOST_FTDC_ETR_BNSP = "3"
"""交割月前一个月下半月"""    
const.THOST_FTDC_ETR_SPOT = "4"
"""交割月份"""    


TThostFtdcPortfolioType = c_char
"""新型组保算法类型"""

const.THOST_FTDC_EPF_None = "0"
"""不使用新型组保算法"""    
const.THOST_FTDC_EPF_SPBM = "1"
"""SPBM算法"""    


TThostFtdcTraderIDType = c_char*21
"""交易所交易员代码类型"""

TThostFtdcInvestorIDType = c_char*13
"""投资者代码类型"""

TThostFtdcBrokerIDType = c_char*11
"""经纪公司代码类型"""

TThostFtdcBrokerAbbrType = c_char*9
"""经纪公司简称类型"""

TThostFtdcBrokerNameType = c_char*81
"""经纪公司名称类型"""

TThostFtdcOldExchangeInstIDType = c_char*31
"""合约在交易所的代码类型"""

TThostFtdcExchangeInstIDType = c_char*81
"""合约在交易所的代码类型"""

TThostFtdcOrderRefType = c_char*13
"""报单引用类型"""

TThostFtdcParticipantIDType = c_char*11
"""会员代码类型"""

TThostFtdcUserIDType = c_char*16
"""用户代码类型"""

TThostFtdcPasswordType = c_char*41
"""密码类型"""

TThostFtdcClientIDType = c_char*11
"""交易编码类型"""

TThostFtdcInstrumentIDType = c_char*81
"""合约代码类型"""

TThostFtdcOldInstrumentIDType = c_char*31
"""合约代码类型"""

TThostFtdcInstrumentCodeType = c_char*31
"""合约标识码类型"""

TThostFtdcMarketIDType = c_char*31
"""市场代码类型"""

TThostFtdcProductNameType = c_char*21
"""产品名称类型"""

TThostFtdcExchangeIDType = c_char*9
"""交易所代码类型"""

TThostFtdcExchangeNameType = c_char*61
"""交易所名称类型"""

TThostFtdcExchangeAbbrType = c_char*9
"""交易所简称类型"""

TThostFtdcExchangeFlagType = c_char*2
"""交易所标志类型"""

TThostFtdcMacAddressType = c_char*21
"""Mac地址类型"""

TThostFtdcSystemIDType = c_char*21
"""系统编号类型"""

TThostFtdcClientLoginRemarkType = c_char*151
"""客户登录备注2类型"""

TThostFtdcDateType = c_char*9
"""日期类型"""

TThostFtdcTimeType = c_char*9
"""时间类型"""

TThostFtdcLongTimeType = c_char*13
"""长时间类型"""

TThostFtdcInstrumentNameType = c_char*21
"""合约名称类型"""

TThostFtdcSettlementGroupIDType = c_char*9
"""结算组代码类型"""

TThostFtdcOrderSysIDType = c_char*21
"""报单编号类型"""

TThostFtdcTradeIDType = c_char*21
"""成交编号类型"""

TThostFtdcCommandTypeType = c_char*65
"""DB命令类型类型"""

TThostFtdcOldIPAddressType = c_char*16
"""IP地址类型"""

TThostFtdcIPAddressType = c_char*33
"""IP地址类型"""

TThostFtdcIPPortType = c_int32
"""IP端口类型"""

TThostFtdcProductInfoType = c_char*11
"""产品信息类型"""

TThostFtdcProtocolInfoType = c_char*11
"""协议信息类型"""

TThostFtdcBusinessUnitType = c_char*21
"""业务单元类型"""

TThostFtdcDepositSeqNoType = c_char*15
"""出入金流水号类型"""

TThostFtdcIdentifiedCardNoType = c_char*51
"""证件号码类型"""

TThostFtdcOrderLocalIDType = c_char*13
"""本地报单编号类型"""

TThostFtdcUserNameType = c_char*81
"""用户名称类型"""

TThostFtdcPartyNameType = c_char*81
"""参与人名称类型"""

TThostFtdcErrorMsgType = c_char*81
"""错误信息类型"""

TThostFtdcFieldNameType = c_char*2049
"""字段名类型"""

TThostFtdcFieldContentType = c_char*2049
"""字段内容类型"""

TThostFtdcSystemNameType = c_char*41
"""系统名称类型"""

TThostFtdcContentType = c_char*501
"""消息正文类型"""

TThostFtdcOrderActionRefType = c_int32
"""报单操作引用类型"""

TThostFtdcInstallCountType = c_int32
"""安装数量类型"""

TThostFtdcInstallIDType = c_int32
"""安装编号类型"""

TThostFtdcErrorIDType = c_int32
"""错误代码类型"""

TThostFtdcSettlementIDType = c_int32
"""结算编号类型"""

TThostFtdcVolumeType = c_int32
"""数量类型"""

TThostFtdcFrontIDType = c_int32
"""前置编号类型"""

TThostFtdcSessionIDType = c_int32
"""会话编号类型"""

TThostFtdcSequenceNoType = c_int32
"""序号类型"""

TThostFtdcCommandNoType = c_int32
"""DB命令序号类型"""

TThostFtdcMillisecType = c_int32
"""时间（毫秒）类型"""

TThostFtdcSecType = c_int32
"""时间（秒）类型"""

TThostFtdcVolumeMultipleType = c_int32
"""合约数量乘数类型"""

TThostFtdcTradingSegmentSNType = c_int32
"""交易阶段编号类型"""

TThostFtdcRequestIDType = c_int32
"""请求编号类型"""

TThostFtdcYearType = c_int32
"""年份类型"""

TThostFtdcMonthType = c_int32
"""月份类型"""

TThostFtdcBoolType = c_int32
"""布尔型类型"""

TThostFtdcPriceType = c_double
"""价格类型"""

TThostFtdcCombOffsetFlagType = c_char*5
"""组合开平标志类型"""

TThostFtdcCombHedgeFlagType = c_char*5
"""组合投机套保标志类型"""

TThostFtdcRatioType = c_double
"""比率类型"""

TThostFtdcMoneyType = c_double
"""资金类型"""

TThostFtdcLargeVolumeType = c_double
"""大额数量类型"""

TThostFtdcSequenceSeriesType = c_short
"""序列系列号类型"""

TThostFtdcCommPhaseNoType = c_short
"""通讯时段编号类型"""

TThostFtdcSequenceLabelType = c_char*2
"""序列编号类型"""

TThostFtdcUnderlyingMultipleType = c_double
"""基础商品乘数类型"""

TThostFtdcPriorityType = c_int32
"""优先级类型"""

TThostFtdcContractCodeType = c_char*41
"""合同编号类型"""

TThostFtdcCityType = c_char*51
"""市类型"""

TThostFtdcIsStockType = c_char*11
"""是否股民类型"""

TThostFtdcChannelType = c_char*51
"""渠道类型"""

TThostFtdcAddressType = c_char*101
"""通讯地址类型"""

TThostFtdcZipCodeType = c_char*7
"""邮政编码类型"""

TThostFtdcTelephoneType = c_char*41
"""联系电话类型"""

TThostFtdcFaxType = c_char*41
"""传真类型"""

TThostFtdcMobileType = c_char*41
"""手机类型"""

TThostFtdcEMailType = c_char*41
"""电子邮件类型"""

TThostFtdcMemoType = c_char*161
"""备注类型"""

TThostFtdcCompanyCodeType = c_char*51
"""企业代码类型"""

TThostFtdcWebsiteType = c_char*51
"""网站地址类型"""

TThostFtdcTaxNoType = c_char*31
"""税务登记号类型"""

TThostFtdcPropertyIDType = c_char*33
"""属性代码类型"""

TThostFtdcPropertyNameType = c_char*65
"""属性名称类型"""

TThostFtdcLicenseNoType = c_char*51
"""营业执照号类型"""

TThostFtdcAgentIDType = c_char*13
"""经纪人代码类型"""

TThostFtdcAgentNameType = c_char*41
"""经纪人名称类型"""

TThostFtdcAgentGroupIDType = c_char*13
"""经纪人组代码类型"""

TThostFtdcAgentGroupNameType = c_char*41
"""经纪人组名称类型"""

TThostFtdcSettlementParamValueType = c_char*256
"""参数代码值类型"""

TThostFtdcCounterIDType = c_char*33
"""计数器代码类型"""

TThostFtdcInvestorGroupNameType = c_char*41
"""投资者分组名称类型"""

TThostFtdcBrandCodeType = c_char*257
"""牌号类型"""

TThostFtdcWarehouseType = c_char*257
"""仓库类型"""

TThostFtdcProductDateType = c_char*41
"""产期类型"""

TThostFtdcGradeType = c_char*41
"""等级类型"""

TThostFtdcClassifyType = c_char*41
"""类别类型"""

TThostFtdcPositionType = c_char*41
"""货位类型"""

TThostFtdcYieldlyType = c_char*41
"""产地类型"""

TThostFtdcWeightType = c_char*41
"""公定重量类型"""

TThostFtdcSubEntryFundNoType = c_int32
"""分项资金流水号类型"""

TThostFtdcFileNameType = c_char*257
"""文件名称类型"""

TThostFtdcUploadModeType = c_char*21
"""上传文件类型类型"""

TThostFtdcAccountIDType = c_char*13
"""投资者帐号类型"""

TThostFtdcBankFlagType = c_char*4
"""银行统一标识类型类型"""

TThostFtdcBankAccountType = c_char*41
"""银行账户类型"""

TThostFtdcOpenNameType = c_char*61
"""银行账户的开户人名称类型"""

TThostFtdcOpenBankType = c_char*101
"""银行账户的开户行类型"""

TThostFtdcBankNameType = c_char*101
"""银行名称类型"""

TThostFtdcPublishPathType = c_char*257
"""发布路径类型"""

TThostFtdcOperatorIDType = c_char*65
"""操作员代码类型"""

TThostFtdcMonthCountType = c_int32
"""月份数量类型"""

TThostFtdcAdvanceMonthArrayType = c_char*13
"""月份提前数组类型"""

TThostFtdcDateExprType = c_char*1025
"""日期表达式类型"""

TThostFtdcInstrumentIDExprType = c_char*41
"""合约代码表达式类型"""

TThostFtdcInstrumentNameExprType = c_char*41
"""合约名称表达式类型"""

TThostFtdcLogLevelType = c_char*33
"""日志级别类型"""

TThostFtdcProcessNameType = c_char*257
"""存储过程名称类型"""

TThostFtdcOperationMemoType = c_char*1025
"""操作摘要类型"""

TThostFtdcBillNoType = c_char*15
"""票据号类型"""

TThostFtdcBillNameType = c_char*33
"""票据名称类型"""

TThostFtdcEnumValueIDType = c_char*65
"""枚举值代码类型"""

TThostFtdcEnumValueTypeType = c_char*33
"""枚举值类型类型"""

TThostFtdcEnumValueLabelType = c_char*65
"""枚举值名称类型"""

TThostFtdcEnumValueResultType = c_char*33
"""枚举值结果类型"""

TThostFtdcRangeIntTypeType = c_char*33
"""限定值类型类型"""

TThostFtdcRangeIntFromType = c_char*33
"""限定值下限类型"""

TThostFtdcRangeIntToType = c_char*33
"""限定值上限类型"""

TThostFtdcFunctionIDType = c_char*25
"""功能代码类型"""

TThostFtdcFunctionValueCodeType = c_char*257
"""功能编码类型"""

TThostFtdcFunctionNameType = c_char*65
"""功能名称类型"""

TThostFtdcRoleIDType = c_char*11
"""角色编号类型"""

TThostFtdcRoleNameType = c_char*41
"""角色名称类型"""

TThostFtdcDescriptionType = c_char*401
"""描述类型"""

TThostFtdcCombineIDType = c_char*25
"""组合编号类型"""

TThostFtdcCombineTypeType = c_char*25
"""组合类型类型"""

TThostFtdcCommentType = c_char*31
"""盈亏算法说明类型"""

TThostFtdcVersionType = c_char*4
"""版本号类型"""

TThostFtdcTradeCodeType = c_char*7
"""交易代码类型"""

TThostFtdcTradeDateType = c_char*9
"""交易日期类型"""

TThostFtdcTradeTimeType = c_char*9
"""交易时间类型"""

TThostFtdcTradeSerialType = c_char*9
"""发起方流水号类型"""

TThostFtdcTradeSerialNoType = c_int32
"""发起方流水号类型"""

TThostFtdcFutureIDType = c_char*11
"""期货公司代码类型"""

TThostFtdcBankIDType = c_char*4
"""银行代码类型"""

TThostFtdcBankBrchIDType = c_char*5
"""银行分中心代码类型"""

TThostFtdcBankBranchIDType = c_char*11
"""分中心代码类型"""

TThostFtdcOperNoType = c_char*17
"""交易柜员类型"""

TThostFtdcDeviceIDType = c_char*3
"""渠道标志类型"""

TThostFtdcRecordNumType = c_char*7
"""记录数类型"""

TThostFtdcFutureAccountType = c_char*22
"""期货资金账号类型"""

TThostFtdcFutureAccPwdType = c_char*17
"""期货资金密码类型"""

TThostFtdcCurrencyCodeType = c_char*4
"""币种类型"""

TThostFtdcRetCodeType = c_char*5
"""响应代码类型"""

TThostFtdcRetInfoType = c_char*129
"""响应信息类型"""

TThostFtdcTradeAmtType = c_char*20
"""银行总余额类型"""

TThostFtdcUseAmtType = c_char*20
"""银行可用余额类型"""

TThostFtdcFetchAmtType = c_char*20
"""银行可取余额类型"""

TThostFtdcCertCodeType = c_char*21
"""证件号码类型"""

TThostFtdcFundProjectIDType = c_char*5
"""资金项目编号类型"""

TThostFtdcProfessionType = c_char*101
"""职业类型"""

TThostFtdcNationalType = c_char*31
"""国籍类型"""

TThostFtdcProvinceType = c_char*51
"""省类型"""

TThostFtdcRegionType = c_char*16
"""区类型"""

TThostFtdcCountryType = c_char*16
"""国家类型"""

TThostFtdcLicenseNOType = c_char*33
"""营业执照类型"""

TThostFtdcCompanyTypeType = c_char*16
"""企业性质类型"""

TThostFtdcBusinessScopeType = c_char*1001
"""经营范围类型"""

TThostFtdcCapitalCurrencyType = c_char*4
"""注册资本币种类型"""

TThostFtdcBranchIDType = c_char*9
"""营业部编号类型"""

TThostFtdcBrokerDNSType = c_char*256
"""域名类型"""

TThostFtdcSentenceType = c_char*501
"""语句类型"""

TThostFtdcLegIDType = c_int32
"""单腿编号类型"""

TThostFtdcLegMultipleType = c_int32
"""单腿乘数类型"""

TThostFtdcImplyLevelType = c_int32
"""派生层数类型"""

TThostFtdcClearAccountType = c_char*33
"""结算账户类型"""

TThostFtdcOrganNOType = c_char*6
"""结算账户类型"""

TThostFtdcClearbarchIDType = c_char*6
"""结算账户联行号类型"""

TThostFtdcUserEventInfoType = c_char*1025
"""用户事件信息类型"""

TThostFtdcParkedOrderIDType = c_char*13
"""预埋报单编号类型"""

TThostFtdcParkedOrderActionIDType = c_char*13
"""预埋撤单编号类型"""

TThostFtdcPhotoTypeNameType = c_char*41
"""影像类型名称类型"""

TThostFtdcPhotoTypeIDType = c_char*5
"""影像类型代码类型"""

TThostFtdcPhotoNameType = c_char*161
"""影像名称类型"""

TThostFtdcTopicIDType = c_int32
"""主题代码类型"""

TThostFtdcReportTypeIDType = c_char*3
"""交易报告类型标识类型"""

TThostFtdcCharacterIDType = c_char*5
"""交易特征代码类型"""

TThostFtdcAMLParamIDType = c_char*21
"""参数代码类型"""

TThostFtdcAMLInvestorTypeType = c_char*3
"""投资者类型类型"""

TThostFtdcAMLIdCardTypeType = c_char*3
"""证件类型类型"""

TThostFtdcAMLTradeDirectType = c_char*3
"""资金进出方向类型"""

TThostFtdcAMLTradeModelType = c_char*3
"""资金进出方式类型"""

TThostFtdcAMLOpParamValueType = c_double
"""业务参数代码值类型"""

TThostFtdcAMLCustomerCardTypeType = c_char*81
"""客户身份证件/证明文件类型类型"""

TThostFtdcAMLInstitutionNameType = c_char*65
"""金融机构网点名称类型"""

TThostFtdcAMLDistrictIDType = c_char*7
"""金融机构网点所在地区行政区划代码类型"""

TThostFtdcAMLRelationShipType = c_char*3
"""金融机构网点与大额交易的关系类型"""

TThostFtdcAMLInstitutionTypeType = c_char*3
"""金融机构网点代码类型类型"""

TThostFtdcAMLInstitutionIDType = c_char*13
"""金融机构网点代码类型"""

TThostFtdcAMLAccountTypeType = c_char*5
"""账户类型类型"""

TThostFtdcAMLTradingTypeType = c_char*7
"""交易方式类型"""

TThostFtdcAMLTransactClassType = c_char*7
"""涉外收支交易分类与代码类型"""

TThostFtdcAMLCapitalIOType = c_char*3
"""资金收付标识类型"""

TThostFtdcAMLSiteType = c_char*10
"""交易地点类型"""

TThostFtdcAMLCapitalPurposeType = c_char*129
"""资金用途类型"""

TThostFtdcAMLReportTypeType = c_char*2
"""报文类型类型"""

TThostFtdcAMLSerialNoType = c_char*5
"""编号类型"""

TThostFtdcAMLStatusType = c_char*2
"""状态类型"""

TThostFtdcAMLSeqCodeType = c_char*65
"""业务标识号类型"""

TThostFtdcAMLFileNameType = c_char*257
"""AML文件名类型"""

TThostFtdcAMLMoneyType = c_double
"""反洗钱资金类型"""

TThostFtdcAMLFileAmountType = c_int32
"""反洗钱资金类型"""

TThostFtdcCFMMCKeyType = c_char*21
"""密钥类型(保证金监管)类型"""

TThostFtdcCFMMCTokenType = c_char*21
"""令牌类型(保证金监管)类型"""

TThostFtdcAMLReportNameType = c_char*81
"""报文名称类型"""

TThostFtdcIndividualNameType = c_char*51
"""个人姓名类型"""

TThostFtdcCurrencyIDType = c_char*4
"""币种代码类型"""

TThostFtdcCustNumberType = c_char*36
"""客户编号类型"""

TThostFtdcOrganCodeType = c_char*36
"""机构编码类型"""

TThostFtdcOrganNameType = c_char*71
"""机构名称类型"""

TThostFtdcSuperOrganCodeType = c_char*12
"""上级机构编码,即期货公司总部、银行总行类型"""

TThostFtdcSubBranchIDType = c_char*31
"""分支机构类型"""

TThostFtdcSubBranchNameType = c_char*71
"""分支机构名称类型"""

TThostFtdcBranchNetCodeType = c_char*31
"""机构网点号类型"""

TThostFtdcBranchNetNameType = c_char*71
"""机构网点名称类型"""

TThostFtdcOrganFlagType = c_char*2
"""机构标识类型"""

TThostFtdcBankCodingForFutureType = c_char*33
"""银行对期货公司的编码类型"""

TThostFtdcBankReturnCodeType = c_char*7
"""银行对返回码的定义类型"""

TThostFtdcPlateReturnCodeType = c_char*5
"""银期转帐平台对返回码的定义类型"""

TThostFtdcBankSubBranchIDType = c_char*31
"""银行分支机构编码类型"""

TThostFtdcFutureBranchIDType = c_char*31
"""期货分支机构编码类型"""

TThostFtdcReturnCodeType = c_char*7
"""返回代码类型"""

TThostFtdcOperatorCodeType = c_char*17
"""操作员类型"""

TThostFtdcClearDepIDType = c_char*6
"""机构结算帐户机构号类型"""

TThostFtdcClearBrchIDType = c_char*6
"""机构结算帐户联行号类型"""

TThostFtdcClearNameType = c_char*71
"""机构结算帐户名称类型"""

TThostFtdcBankAccountNameType = c_char*71
"""银行帐户名称类型"""

TThostFtdcInvDepIDType = c_char*6
"""机构投资人账号机构号类型"""

TThostFtdcInvBrchIDType = c_char*6
"""机构投资人联行号类型"""

TThostFtdcMessageFormatVersionType = c_char*36
"""信息格式版本类型"""

TThostFtdcDigestType = c_char*36
"""摘要类型"""

TThostFtdcAuthenticDataType = c_char*129
"""认证数据类型"""

TThostFtdcPasswordKeyType = c_char*129
"""密钥类型"""

TThostFtdcFutureAccountNameType = c_char*129
"""期货帐户名称类型"""

TThostFtdcMobilePhoneType = c_char*21
"""手机类型"""

TThostFtdcFutureMainKeyType = c_char*129
"""期货公司主密钥类型"""

TThostFtdcFutureWorkKeyType = c_char*129
"""期货公司工作密钥类型"""

TThostFtdcFutureTransKeyType = c_char*129
"""期货公司传输密钥类型"""

TThostFtdcBankMainKeyType = c_char*129
"""银行主密钥类型"""

TThostFtdcBankWorkKeyType = c_char*129
"""银行工作密钥类型"""

TThostFtdcBankTransKeyType = c_char*129
"""银行传输密钥类型"""

TThostFtdcBankServerDescriptionType = c_char*129
"""银行服务器描述信息类型"""

TThostFtdcAddInfoType = c_char*129
"""附加信息类型"""

TThostFtdcDescrInfoForReturnCodeType = c_char*129
"""返回码描述类型"""

TThostFtdcCountryCodeType = c_char*21
"""国家代码类型"""

TThostFtdcSerialType = c_int32
"""流水号类型"""

TThostFtdcPlateSerialType = c_int32
"""平台流水号类型"""

TThostFtdcBankSerialType = c_char*13
"""银行流水号类型"""

TThostFtdcCorrectSerialType = c_int32
"""被冲正交易流水号类型"""

TThostFtdcFutureSerialType = c_int32
"""期货公司流水号类型"""

TThostFtdcApplicationIDType = c_int32
"""应用标识类型"""

TThostFtdcBankProxyIDType = c_int32
"""银行代理标识类型"""

TThostFtdcFBTCoreIDType = c_int32
"""银期转帐核心系统标识类型"""

TThostFtdcServerPortType = c_int32
"""服务端口号类型"""

TThostFtdcRepealedTimesType = c_int32
"""已经冲正次数类型"""

TThostFtdcRepealTimeIntervalType = c_int32
"""冲正时间间隔类型"""

TThostFtdcTotalTimesType = c_int32
"""每日累计转帐次数类型"""

TThostFtdcFBTRequestIDType = c_int32
"""请求ID类型"""

TThostFtdcTIDType = c_int32
"""交易ID类型"""

TThostFtdcTradeAmountType = c_double
"""交易金额（元）类型"""

TThostFtdcCustFeeType = c_double
"""应收客户费用（元）类型"""

TThostFtdcFutureFeeType = c_double
"""应收期货公司费用（元）类型"""

TThostFtdcSingleMaxAmtType = c_double
"""单笔最高限额类型"""

TThostFtdcSingleMinAmtType = c_double
"""单笔最低限额类型"""

TThostFtdcTotalAmtType = c_double
"""每日累计转帐额度类型"""

TThostFtdcServiceIDType = c_int32
"""服务编号类型"""

TThostFtdcServiceLineNoType = c_int32
"""服务线路编号类型"""

TThostFtdcServiceNameType = c_char*61
"""服务名类型"""

TThostFtdcCommApiPointerType = c_int32
"""通讯API指针类型"""

TThostFtdcBankIDByBankType = c_char*21
"""银行自己的编码类型"""

TThostFtdcBankOperNoType = c_char*4
"""银行操作员号类型"""

TThostFtdcBankCustNoType = c_char*21
"""银行客户号类型"""

TThostFtdcDBOPSeqNoType = c_int32
"""递增的序列号类型"""

TThostFtdcTableNameType = c_char*61
"""FBT表名类型"""

TThostFtdcPKNameType = c_char*201
"""FBT表操作主键名类型"""

TThostFtdcPKValueType = c_char*501
"""FBT表操作主键值类型"""

TThostFtdcTargetIDType = c_char*4
"""同步目标编号类型"""

TThostFtdcFBETimeType = c_char*7
"""各种换汇时间类型"""

TThostFtdcFBEBankNoType = c_char*13
"""换汇银行行号类型"""

TThostFtdcFBECertNoType = c_char*13
"""换汇凭证号类型"""

TThostFtdcFBEBankAccountType = c_char*33
"""换汇银行账户类型"""

TThostFtdcFBEBankAccountNameType = c_char*61
"""换汇银行账户名类型"""

TThostFtdcFBEAmtType = c_double
"""各种换汇金额类型"""

TThostFtdcFBEBusinessTypeType = c_char*3
"""换汇业务类型类型"""

TThostFtdcFBEPostScriptType = c_char*61
"""换汇附言类型"""

TThostFtdcFBERemarkType = c_char*71
"""换汇备注类型"""

TThostFtdcExRateType = c_double
"""换汇汇率类型"""

TThostFtdcFBERtnMsgType = c_char*61
"""换汇返回信息类型"""

TThostFtdcFBEExtendMsgType = c_char*61
"""换汇扩展信息类型"""

TThostFtdcFBEBusinessSerialType = c_char*31
"""换汇记账流水号类型"""

TThostFtdcFBESystemSerialType = c_char*21
"""换汇流水号类型"""

TThostFtdcFBETotalExCntType = c_int32
"""换汇交易总笔数类型"""

TThostFtdcFBEOpenBankType = c_char*61
"""换汇账户开户行类型"""

TThostFtdcFBEFileNameType = c_char*21
"""换汇相关文件名类型"""

TThostFtdcFBEBatchSerialType = c_char*21
"""换汇批次号类型"""

TThostFtdcRiskNofityInfoType = c_char*257
"""客户风险通知消息类型"""

TThostFtdcForceCloseSceneIdType = c_char*24
"""强平场景编号类型"""

TThostFtdcInstrumentIDsType = c_char*101
"""多个产品代码,用+分隔,如cu+zn类型"""

TThostFtdcParamIDType = c_int32
"""参数代码类型"""

TThostFtdcParamNameType = c_char*41
"""参数名类型"""

TThostFtdcParamValueType = c_char*41
"""参数值类型"""

TThostFtdcIndustryIDType = c_char*17
"""行业编码类型"""

TThostFtdcQuestionIDType = c_char*5
"""特有信息编号类型"""

TThostFtdcQuestionContentType = c_char*41
"""特有信息说明类型"""

TThostFtdcOptionIDType = c_char*13
"""选项编号类型"""

TThostFtdcOptionContentType = c_char*61
"""选项说明类型"""

TThostFtdcProcessIDType = c_char*33
"""业务流水号类型"""

TThostFtdcSeqNoType = c_int32
"""流水号类型"""

TThostFtdcUOAProcessStatusType = c_char*3
"""流程状态类型"""

TThostFtdcProcessTypeType = c_char*3
"""流程功能类型类型"""

TThostFtdcExReturnCodeType = c_int32
"""交易所返回码类型"""

TThostFtdcClientClassifyType = c_char*11
"""客户分类码类型"""

TThostFtdcUOAOrganTypeType = c_char*11
"""单位性质类型"""

TThostFtdcUOACountryCodeType = c_char*11
"""国家代码类型"""

TThostFtdcAreaCodeType = c_char*11
"""区号类型"""

TThostFtdcFuturesIDType = c_char*21
"""监控中心为客户分配的代码类型"""

TThostFtdcCffmcDateType = c_char*11
"""日期类型"""

TThostFtdcCffmcTimeType = c_char*11
"""时间类型"""

TThostFtdcNocIDType = c_char*21
"""组织机构代码类型"""

TThostFtdcEventTypeType = c_char*33
"""业务操作类型类型"""

TThostFtdcQueryDepthType = c_int32
"""查询深度类型"""

TThostFtdcDataCenterIDType = c_int32
"""数据中心代码类型"""

TThostFtdcCheckNoType = c_int32
"""操作次数类型"""

TThostFtdcRateTemplateNameType = c_char*61
"""模型名称类型"""

TThostFtdcPropertyStringType = c_char*2049
"""用于查询的投资属性字段类型"""

TThostFtdcRateTemplateIDType = c_char*9
"""模型代码类型"""

TThostFtdcRiskRateType = c_char*21
"""风险度类型"""

TThostFtdcTimestampType = c_int32
"""时间戳类型"""

TThostFtdcInvestorIDRuleNameType = c_char*61
"""号段规则名称类型"""

TThostFtdcInvestorIDRuleExprType = c_char*513
"""号段规则表达式类型"""

TThostFtdcLastDriftType = c_int32
"""上次OTP漂移值类型"""

TThostFtdcLastSuccessType = c_int32
"""上次OTP成功值类型"""

TThostFtdcAuthKeyType = c_char*41
"""令牌密钥类型"""

TThostFtdcSerialNumberType = c_char*17
"""序列号类型"""

TThostFtdcOTPVendorsIDType = c_char*2
"""动态令牌提供商类型"""

TThostFtdcOTPVendorsNameType = c_char*61
"""动态令牌提供商名称类型"""

TThostFtdcTimeSpanType = c_char*9
"""时间跨度类型"""

TThostFtdcImportSequenceIDType = c_char*17
"""动态令牌导入批次编号类型"""

TThostFtdcComTypeType = c_int32
"""组合成交类型类型"""

TThostFtdcUserProductIDType = c_char*33
"""产品标识类型"""

TThostFtdcUserProductNameType = c_char*65
"""产品名称类型"""

TThostFtdcUserProductMemoType = c_char*129
"""产品说明类型"""

TThostFtdcCSRCCancelFlagType = c_char*2
"""新增或变更标志类型"""

TThostFtdcCSRCDateType = c_char*11
"""日期类型"""

TThostFtdcCSRCInvestorNameType = c_char*201
"""客户名称类型"""

TThostFtdcCSRCOpenInvestorNameType = c_char*101
"""客户名称类型"""

TThostFtdcCSRCInvestorIDType = c_char*13
"""客户代码类型"""

TThostFtdcCSRCIdentifiedCardNoType = c_char*51
"""证件号码类型"""

TThostFtdcCSRCClientIDType = c_char*11
"""交易编码类型"""

TThostFtdcCSRCBankFlagType = c_char*3
"""银行标识类型"""

TThostFtdcCSRCBankAccountType = c_char*23
"""银行账户类型"""

TThostFtdcCSRCOpenNameType = c_char*401
"""开户人类型"""

TThostFtdcCSRCMemoType = c_char*101
"""说明类型"""

TThostFtdcCSRCTimeType = c_char*11
"""时间类型"""

TThostFtdcCSRCTradeIDType = c_char*21
"""成交流水号类型"""

TThostFtdcCSRCExchangeInstIDType = c_char*31
"""合约代码类型"""

TThostFtdcCSRCMortgageNameType = c_char*7
"""质押品名称类型"""

TThostFtdcCSRCReasonType = c_char*3
"""事由类型"""

TThostFtdcIsSettlementType = c_char*2
"""是否为非结算会员类型"""

TThostFtdcCSRCMoneyType = c_double
"""资金类型"""

TThostFtdcCSRCPriceType = c_double
"""价格类型"""

TThostFtdcCSRCOptionsTypeType = c_char*2
"""期权类型类型"""

TThostFtdcCSRCStrikePriceType = c_double
"""执行价类型"""

TThostFtdcCSRCTargetProductIDType = c_char*3
"""标的品种类型"""

TThostFtdcCSRCTargetInstrIDType = c_char*31
"""标的合约类型"""

TThostFtdcCommModelNameType = c_char*161
"""手续费率模板名称类型"""

TThostFtdcCommModelMemoType = c_char*1025
"""手续费率模板备注类型"""

TThostFtdcAgentBrokerIDType = c_char*13
"""代理经纪公司代码类型"""

TThostFtdcDRIdentityIDType = c_int32
"""交易中心代码类型"""

TThostFtdcDRIdentityNameType = c_char*65
"""交易中心名称类型"""

TThostFtdcDBLinkIDType = c_char*31
"""DBLink标识号类型"""

TThostFtdcSRiskRateType = c_char*21
"""风险度类型"""

TThostFtdcSequenceNo12Type = c_int32
"""序号类型"""

TThostFtdcCSRCFreezeStatusType = c_char*2
"""休眠状态类型"""

TThostFtdcRightTemplateIDType = c_char*9
"""模板代码类型"""

TThostFtdcRightTemplateNameType = c_char*61
"""模板名称类型"""

TThostFtdcAmlCheckFlowType = c_char*2
"""反洗钱数据抽取审核流程类型"""

TThostFtdcDataTypeType = c_char*129
"""数据类型类型"""

TThostFtdcSettleManagerIDType = c_char*33
"""结算配置代码类型"""

TThostFtdcSettleManagerNameType = c_char*129
"""结算配置名称类型"""

TThostFtdcCheckResultMemoType = c_char*1025
"""核对结果说明类型"""

TThostFtdcFunctionUrlType = c_char*1025
"""功能链接类型"""

TThostFtdcAuthInfoType = c_char*129
"""客户端认证信息类型"""

TThostFtdcAuthCodeType = c_char*17
"""客户端认证码类型"""

TThostFtdcToolIDType = c_char*9
"""工具代码类型"""

TThostFtdcToolNameType = c_char*81
"""工具名称类型"""

TThostFtdcCurrencyUnitType = c_double
"""币种单位数量类型"""

TThostFtdcExchangeRateType = c_double
"""汇率类型"""

TThostFtdcCurrencyNameType = c_char*31
"""币种名称类型"""

TThostFtdcCurrencySignType = c_char*4
"""币种符号类型"""

TThostFtdcCurrExchCertNoType = c_char*13
"""凭证号类型"""

TThostFtdcBatchSerialNoType = c_char*21
"""批次号类型"""

TThostFtdcPageControlType = c_char*2
"""换汇页面控制类型"""

TThostFtdcRecordCountType = c_int32
"""记录数类型"""

TThostFtdcCurrencySwapMemoType = c_char*101
"""换汇需确认信息类型"""

TThostFtdcWorkPlaceType = c_char*101
"""工作单位类型"""

TThostFtdcBusinessPeriodType = c_char*21
"""经营期限类型"""

TThostFtdcWebSiteType = c_char*101
"""网址类型"""

TThostFtdcUOAIdCardTypeType = c_char*3
"""统一开户证件类型类型"""

TThostFtdcClientModeType = c_char*3
"""开户模式类型"""

TThostFtdcInvestorFullNameType = c_char*101
"""投资者全称类型"""

TThostFtdcUOABrokerIDType = c_char*11
"""境外中介机构ID类型"""

TThostFtdcUOAZipCodeType = c_char*11
"""邮政编码类型"""

TThostFtdcUOAEMailType = c_char*101
"""电子邮箱类型"""

TThostFtdcOldCityType = c_char*41
"""城市类型"""

TThostFtdcCorporateIdentifiedCardNoType = c_char*101
"""法人代表证件号码类型"""

TThostFtdcLedgerManageIDType = c_char*51
"""分户管理资产编码类型"""

TThostFtdcInvestVarietyType = c_char*101
"""投资品种类型"""

TThostFtdcBankAccountTypeType = c_char*2
"""账户类别类型"""

TThostFtdcLedgerManageBankType = c_char*101
"""开户银行类型"""

TThostFtdcCffexDepartmentNameType = c_char*101
"""开户营业部类型"""

TThostFtdcCffexDepartmentCodeType = c_char*9
"""营业部代码类型"""

TThostFtdcCSRCMemo1Type = c_char*41
"""说明类型"""

TThostFtdcAssetmgrCFullNameType = c_char*101
"""代理资产管理业务的期货公司全称类型"""

TThostFtdcAssetmgrApprovalNOType = c_char*51
"""资产管理业务批文号类型"""

TThostFtdcAssetmgrMgrNameType = c_char*401
"""资产管理业务负责人姓名类型"""

TThostFtdcCSRCAmTypeType = c_char*5
"""机构类型类型"""

TThostFtdcCSRCNationalType = c_char*4
"""国籍类型"""

TThostFtdcCSRCSecAgentIDType = c_char*11
"""二级代理ID类型"""

TThostFtdcAmAccountType = c_char*23
"""投资账户类型"""

TThostFtdcUOMType = c_char*11
"""计量单位类型"""

TThostFtdcSHFEInstLifePhaseType = c_char*3
"""上期所合约生命周期状态类型"""

TThostFtdcSHFEProductClassType = c_char*11
"""产品类型类型"""

TThostFtdcPriceDecimalType = c_char*2
"""价格小数位类型"""

TThostFtdcInTheMoneyFlagType = c_char*2
"""平值期权标志类型"""

TThostFtdcBigMoneyType = c_double
"""资金类型"""

TThostFtdcCombinInstrIDType = c_char*61
"""套利合约代码类型"""

TThostFtdcCombinSettlePriceType = c_char*61
"""各腿结算价类型"""

TThostFtdcDCEPriorityType = c_int32
"""优先级类型"""

TThostFtdcTradeGroupIDType = c_int32
"""成交组号类型"""

TThostFtdcIsCheckPrepaType = c_int32
"""是否校验开户可用资金类型"""

TThostFtdcSwapBusinessTypeType = c_char*3
"""换汇业务种类类型"""

TThostFtdcExecOrderSysIDType = c_char*21
"""执行宣告系统编号类型"""

TThostFtdcStrikeSequenceType = c_int32
"""执行序号类型"""

TThostFtdcStrikeTimeType = c_char*13
"""执行时间类型"""

TThostFtdcLoginRemarkType = c_char*36
"""登录备注类型"""

TThostFtdcInvestUnitIDType = c_char*17
"""投资单元代码类型"""

TThostFtdcBulletinIDType = c_int32
"""公告编号类型"""

TThostFtdcNewsTypeType = c_char*3
"""公告类型类型"""

TThostFtdcNewsUrgencyType = c_char
"""紧急程度类型"""

TThostFtdcAbstractType = c_char*81
"""消息摘要类型"""

TThostFtdcComeFromType = c_char*21
"""消息来源类型"""

TThostFtdcURLLinkType = c_char*201
"""WEB地址类型"""

TThostFtdcLongIndividualNameType = c_char*161
"""长个人姓名类型"""

TThostFtdcLongFBEBankAccountNameType = c_char*161
"""长换汇银行账户名类型"""

TThostFtdcDateTimeType = c_char*17
"""日期时间类型"""

TThostFtdcRandomStringType = c_char*17
"""随机串类型"""

TThostFtdcAppIDType = c_char*33
"""App代码类型"""

TThostFtdcSystemInfoLenType = c_int32
"""系统信息长度类型"""

TThostFtdcAdditionalInfoLenType = c_int32
"""补充信息长度类型"""

TThostFtdcClientSystemInfoType = c_char*273
"""交易终端系统信息类型"""

TThostFtdcAdditionalInfoType = c_char*261
"""系统外部信息类型"""

TThostFtdcBase64ClientSystemInfoType = c_char*365
"""base64交易终端系统信息类型"""

TThostFtdcBase64AdditionalInfoType = c_char*349
"""base64系统外部信息类型"""

TThostFtdcCurrentAuthMethodType = c_int32
"""当前可用的认证模式，0代表无需认证模式 A从低位开始最后一位代表图片验证码，倒数第二位代表动态口令，倒数第三位代表短信验证码类型"""

TThostFtdcCaptchaInfoLenType = c_int32
"""图片验证信息长度类型"""

TThostFtdcCaptchaInfoType = c_char*2561
"""图片验证信息类型"""

TThostFtdcUserTextSeqType = c_int32
"""用户短信验证码的编号类型"""

TThostFtdcHandshakeDataType = c_char*301
"""握手数据内容类型"""

TThostFtdcHandshakeDataLenType = c_int32
"""握手数据内容长度类型"""

TThostFtdcCryptoKeyVersionType = c_char*31
"""api与front通信密钥版本号类型"""

TThostFtdcRsaKeyVersionType = c_int32
"""公钥版本号类型"""

TThostFtdcSoftwareProviderIDType = c_char*22
"""交易软件商ID类型"""

TThostFtdcCollectTimeType = c_char*21
"""信息采集时间类型"""

TThostFtdcQueryFreqType = c_int32
"""查询频率类型"""

TThostFtdcOTCTraderIDType = c_char*31
"""OTC交易员代码类型"""

TThostFtdcRiskValueType = c_double
"""期货风险值类型"""

TThostFtdcIDBNameType = c_char*100
"""握手数据内容类型"""

TThostFtdcDiscountRatioType = c_double
"""折扣率类型"""

TThostFtdcSyncDescriptionType = c_char*257
"""追平描述类型"""

TThostFtdcCommonIntType = c_int32
"""通用int类型类型"""

TThostFtdcSysVersionType = c_char*41
"""系统版本类型"""

TThostFtdcDeltaType = c_double
"""Delta类型类型"""

TThostFtdcSpreadIdType = c_int32
"""抵扣组优先级类型"""

TThostFtdcPortfolioDefIDType = c_int32
"""SPBM组合套餐ID类型"""
