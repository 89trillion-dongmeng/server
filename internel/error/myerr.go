package error

import "runtime/debug"

// Api handle status
type MyErr struct {
	Code    int
	Message string
	Level   Level
	Stack   []byte
}

type Level uint32

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
	SuccessLevel
)

// 附加stack信息
func (e *MyErr) WithStack() *MyErr {
	err := *e
	err.Stack = debug.Stack()
	return &err
}

// get error information
func (e *MyErr) Error() string {
	return e.Message
}

func (e *MyErr) Detail() string {
	msg := e.Message
	if len(e.Stack) > 0 {
		msg += string(e.Stack)
	}
	return msg
}

var (
	SUCCESS = &MyErr{Code: 200, Message: "OK", Level: SuccessLevel}
	FAILED  = &MyErr{Code: 500, Message: "Fail", Level: InfoLevel}

	// 缺少必选参数
	LACK_OF_HEADER  = &MyErr{Code: 1001, Message: "lack of header", Level: InfoLevel}
	LACK_OF_PARAMS  = &MyErr{Code: 1002, Message: "lack of params", Level: InfoLevel}
	LACK_OF_SUBTYPE = &MyErr{Code: 1003, Message: "lack of subtype", Level: InfoLevel}
	LACK_OF_FLAG    = &MyErr{Code: 1004, Message: "lack of flag", Level: InfoLevel}
	LACK_OF_VALUES  = &MyErr{Code: 1005, Message: "lack of values", Level: InfoLevel}
	LACK_OF_EXT     = &MyErr{Code: 1006, Message: "lack of ext", Level: InfoLevel}

	// 参数值错误
	INVALID_HEADER   = &MyErr{Code: 1101, Message: "invalid header", Level: InfoLevel}
	INVALID_EVENT    = &MyErr{Code: 1102, Message: "not support event", Level: InfoLevel} // 不支持的事件
	INVALID_PARAMS   = &MyErr{Code: 1103, Message: "invalid params", Level: InfoLevel}
	INVALID_SUBTYPE  = &MyErr{Code: 1104, Message: "invalid subtype", Level: InfoLevel}
	INVALID_TROOP_ID = &MyErr{Code: 1105, Message: "invalid troop id", Level: InfoLevel}
	INVALID_SKU      = &MyErr{Code: 1106, Message: "invalid sku", Level: InfoLevel}
	INVALID_GOODS_ID = &MyErr{Code: 1107, Message: "invalid goods id", Level: InfoLevel}
	INVALID_POS      = &MyErr{Code: 1108, Message: "invalid pos", Level: InfoLevel}
	INVALID_PLATFORM = &MyErr{Code: 1109, Message: "invalid platform", Level: InfoLevel}
	INVALID_LEVEL    = &MyErr{Code: 1110, Message: "invalid level", Level: InfoLevel}
	INVALID_VALUES   = &MyErr{Code: 1111, Message: "invalid values", Level: InfoLevel}
	INVALID_FLAG     = &MyErr{Code: 1112, Message: "invalid flag", Level: InfoLevel}
	INVALID_SEASON   = &MyErr{Code: 1113, Message: "invalid season", Level: InfoLevel}
	INVALID_AMOUNT   = &MyErr{Code: 1114, Message: "invalid amount", Level: InfoLevel}
	INVALID_USERID   = &MyErr{Code: 1115, Message: "invalid userid", Level: InfoLevel}

	// 权限限制
	ERROR_AUTH_CHECK_TOKEN_FAIL        = &MyErr{Code: 1201, Message: "token fail", Level: InfoLevel}
	ERROR_AREADY_LOGIN_ON_OTHER_DEVICE = &MyErr{Code: 1202, Message: "login on other device", Level: InfoLevel}
	ERROR_AUTH_BLACK_LIST              = &MyErr{Code: 1203, Message: "forbid", Level: InfoLevel}
	ERROR_AUTH_NO_EXIST_ACCOUNT        = &MyErr{Code: 1204, Message: "no this account", Level: InfoLevel}
	ERROR_AUTH_CHECK_TOKEN_EXPIRED     = &MyErr{Code: 1205, Message: "token expired", Level: InfoLevel}
	ERROR_GENERATE_TOKEN               = &MyErr{Code: 1206, Message: "generate token fail", Level: InfoLevel}
	ERROR_DUPLICATE_REGISTER_ACCOUNT   = &MyErr{Code: 1207, Message: "duplicate register account", Level: InfoLevel}
	ERROR_NO_ACCESS                    = &MyErr{Code: 1208, Message: "no access", Level: InfoLevel}
	ERROR_GET_TOKEN                    = &MyErr{Code: 1209, Message: "get token fail", Level: InfoLevel}
	UPGRADE_WARNING                    = &MyErr{Code: 1208, Message: "no access", Level: InfoLevel}
	UPGRADE_ERROR                      = &MyErr{Code: 1209, Message: "get token fail", Level: InfoLevel}

	// 逻辑错误
	NO_ENOUGH_MONEY           = &MyErr{Code: 2001, Message: "no enough currency", Level: InfoLevel} // 钱不够
	NO_ENOUGH_TOOL            = &MyErr{Code: 2002, Message: "no enough tool", Level: InfoLevel}     // 道具不够
	NO_ENOUGH_ARMY            = &MyErr{Code: 2003, Message: "no enough tool", Level: InfoLevel}     // 士兵不足
	DUPLICATED_CLAIM          = &MyErr{Code: 2004, Message: "duplicated claim", Level: InfoLevel}   // 重复领取
	ILLEGAL_MATCH             = &MyErr{Code: 2005, Message: "illegal match", Level: InfoLevel}
	BEYOND_MAX_LEVEL          = &MyErr{Code: 2006, Message: "beyond_max_level", Level: InfoLevel}
	EVENT_OFFLINE             = &MyErr{Code: 2007, Message: "event offline", Level: InfoLevel}             // 活动已结束
	UNMEET_REWARD_REQUIREMENT = &MyErr{Code: 2008, Message: "unmeet reward requirement", Level: InfoLevel} // 未满足领取条件
	DUPLICATED_OPERATION      = &MyErr{Code: 2009, Message: "duplicated operation", Level: InfoLevel}      // 重复操作
	NOT_ENOUGH_COUNT          = &MyErr{Code: 2010, Message: "no enough count", Level: InfoLevel}           // 次数不足
	ILLEGAL_SEASON_ID         = &MyErr{Code: 2011, Message: "illegal season id"}                           // 赛季id非法, 客户端需要再次同步season id
	MODIFY_CONFIG_ERROR       = &MyErr{Code: 2012, Message: "modify config error"}

	// draw card
	ERROR_DRAW_CARD_QUALITY_FAIL = &MyErr{Code: 2101, Message: "draw card quality error", Level: InfoLevel}
	ERROR_DRAW_CARD_CONFIG_FAIL  = &MyErr{Code: 2102, Message: "draw card cfg error", Level: InfoLevel}
	ERROR_DRAW_CARD_ARMY_FAIL    = &MyErr{Code: 2103, Message: "draw card army_error", Level: InfoLevel}
	ERROR_ADD_COUNT_FAIL         = &MyErr{Code: 2104, Message: "draw card add count error", Level: InfoLevel}
	ERROR_NOT_ENOUGH_LUCKY_FAIL  = &MyErr{Code: 2105, Message: "lucky not enough", Level: InfoLevel}
	ERROR_LUCKY_CARD_FAIL        = &MyErr{Code: 2106, Message: "lucky card error", Level: InfoLevel}

	// 阵型和建筑
	//ARMY_LIST_ERROR         = &MyErr{Code: 2201, Message: "army list error", Level: InfoLevel}       // 士兵总列表错误
	BUILDING_AMOUNT_INVALID = &MyErr{Code: 2202, Message: "building amount error", Level: InfoLevel} // 该种建筑超过限制数量
	//ATKLAY_SYNC_ERROR       = &MyErr{Code: 2203, Message: "atk sync error", Level: InfoLevel}              // 单个建筑信息错误
	BUILDING_ONE_ERROR = &MyErr{Code: 2204, Message: "building one error", Level: InfoLevel} // 单个建筑信息错误
	//BUILDING_SYNC_ERROR     = &MyErr{Code: 2205, Message: "building sync error", Level: InfoLevel}         // 建筑总列表错误
	VECTOR_LIST_ERROR = &MyErr{Code: 2206, Message: "upload operation list error", Level: InfoLevel} // vectorListError
	HERO_SYNC_ONE     = &MyErr{Code: 2207, Message: "hero_sync_one", Level: InfoLevel}               // sync一个英雄的data

	//LAYOUT_SYNC_ATK = &MyErr{Code: 2208, Message: "layout_sync_atk", Level: InfoLevel}          // 预留
	//LAYOUT_SYNC_DEF = &MyErr{Code: 2209, Message: "layout_sync_def", Level: InfoLevel}          // 预留
	DEF_DATA_ERROR = &MyErr{Code: 2210, Message: "def_data_error", Level: InfoLevel} // 预留
	//ATK_DATA_ERROR  = &MyErr{Code: 2211, Message: "atk_data_error", Level: InfoLevel}           // 预留
	SELL_POS_ERROR = &MyErr{Code: 2212, Message: "sell_pos_error", Level: InfoLevel} // 预留
	//BDLV_DATA_ERROR = &MyErr{Code: 2213, Message: "buildingLevel_data_error", Level: InfoLevel} // 预留

	// 赏金关卡
	HUNT_ROUND_ERROR = &MyErr{Code: 2301, Message: "invalid headhunt round", Level: InfoLevel}
	HUNT_LEVEL_ERROR = &MyErr{Code: 2302, Message: "invalid headhunt level", Level: InfoLevel}

	// PVP
	ERROR_PVP_NOT_HAS_ROOM    = &MyErr{Code: 2401, Message: "not_has_room", Level: InfoLevel}
	ERROR_GET_STAGE           = &MyErr{Code: 2402, Message: "user_stage_error", Level: InfoLevel}
	ERROR_GET_PLAYER_INFO     = &MyErr{Code: 2403, Message: "user_info_error", Level: InfoLevel}
	ERROR_PVP_MATCH_FAIL      = &MyErr{Code: 2404, Message: "match_error", Level: InfoLevel}
	ERROR_PVP_START_FAIL      = &MyErr{Code: 2405, Message: "attack_start_fail", Level: InfoLevel}
	ERROR_PVP_ATTACK_ID_FAIL  = &MyErr{Code: 2406, Message: "get_attack_id_error", Level: InfoLevel}
	ERROR_PVP_INIT_ROOM_FAIL  = &MyErr{Code: 2407, Message: "init_room_error", Level: InfoLevel}
	ERROR_PVP_GET_REWARD_FAIL = &MyErr{Code: 2408, Message: "reward_fail", Level: InfoLevel}
	ERROR_INIT_ROOM_FAIL      = &MyErr{Code: 2409, Message: "init_room_fail", Level: InfoLevel}
	RATIO_ERROR               = &MyErr{Code: 2410, Message: "ratio_error", Level: InfoLevel}
	TRAIN_CAMP_FINISH_FAIL    = &MyErr{Code: 2411, Message: "train_camp_finish_fail", Level: InfoLevel}
	ERROR_REVENGE_STATUS_FAIL = &MyErr{Code: 2412, Message: "revenge_status_fail", Level: InfoLevel}

	// point race
	ERROR_GET_PR_ROOM_FAIL = &MyErr{Code: 2450, Message: "get_point_race_room_fail", Level: InfoLevel}

	// team war
	USER_NOT_IN_TEAM_FAIL = &MyErr{Code: 2501, Message: "user_not_in_team", Level: InfoLevel}
	PERMISSION_DENY       = &MyErr{Code: 2502, Message: "permission_deny", Level: InfoLevel}
	TEAM_NOT_EXISTS       = &MyErr{Code: 2503, Message: "team_not_exists", Level: InfoLevel}
	TEAM_NEW_USER         = &MyErr{Code: 2504, Message: "join_day_can_not_attack", Level: InfoLevel}
	TEAM_TRUCE_DAY        = &MyErr{Code: 2505, Message: "team_truce_day", Level: InfoLevel}
	TEAM_NO_ENTRY         = &MyErr{Code: 2506, Message: "team_no_entry", Level: InfoLevel}
	TEAM_NO_ENEMY         = &MyErr{Code: 2507, Message: "team_no_enemy", Level: InfoLevel}
	TEAM_ENEMY_DISBAND    = &MyErr{Code: 2508, Message: "team_enemy_disband", Level: InfoLevel}

	// shop
	MONTHCARD_NOT_BUY = &MyErr{Code: 2601, Message: "monthCard_not_buy", Level: InfoLevel}
	MONTHCARD_ERROR   = &MyErr{Code: 2602, Message: "monthCard_error", Level: InfoLevel}
	LACK_OF_CONFIG    = &MyErr{Code: 2603, Message: "lack of config", Level: InfoLevel}
	VERIFY_FAILED     = &MyErr{Code: 2604, Message: "verify failed", Level: InfoLevel}
	FUSE_POS_ERROR    = &MyErr{Code: 2605, Message: "fuse pos error", Level: InfoLevel}
	NOT_BOUGHT_ERROR  = &MyErr{Code: 2606, Message: "not bought error", Level: InfoLevel}
	BOUGHT_TWICE_RROR = &MyErr{Code: 2607, Message: "bought twice error", Level: InfoLevel}

	// flop
	ERROR_FLOP_CARD_FAIL   = &MyErr{Code: 2701, Message: "flop card error", Level: InfoLevel}
	ERROR_FLOP_RANDOM_FAIL = &MyErr{Code: 2702, Message: "flop random error", Level: InfoLevel}

	// 内部错误\系统错误
	REDIS_INSIDE_ERROR      = &MyErr{Code: 3001, Message: "redis inside error", Level: ErrorLevel}      // redis内部错误(因为网络等原因方法直接返回的err)
	REDIS_DATA_ERROR        = &MyErr{Code: 3002, Message: "redis data invalid", Level: ErrorLevel}      // redis内部数据不合法(redis内部存储的数据不合法)
	MONGO_INSIDE_ERROR      = &MyErr{Code: 3003, Message: "mongo inside error", Level: ErrorLevel}      // mongo 内部错误
	RESOURCE_NOT_ENOUGH     = &MyErr{Code: 3004, Message: "resource not enough", Level: InfoLevel}      // 资源不足
	JSON_MARSHAL_ERROR      = &MyErr{Code: 3005, Message: "json marshal error", Level: ErrorLevel}      // json marshal error
	JSON_UNMARSHAL_ERROR    = &MyErr{Code: 3006, Message: "json unmarshal error", Level: ErrorLevel}    // json unmarshal error
	PROTO_UNMARSHAL_ERROR   = &MyErr{Code: 3007, Message: "invalid protobuf format", Level: InfoLevel}  // protobuf 解析错误
	KICK_OTHER_DEVICE_ERROR = &MyErr{Code: 3008, Message: "kick other device error", Level: ErrorLevel} // 踢下线失败
	ERROR_GRPC_SERVER_OFF   = &MyErr{Code: 3009, Message: "grpc sever offline", Level: ErrorLevel}      // grpc连接失败
	ERROR_GRPC_REQUEST_FAIL = &MyErr{Code: 3010, Message: "grpc request fail", Level: ErrorLevel}       // grpc发消息失败
)