package post

import "time"

type Post struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"exstrconv.FormatFloat(t_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		OrderID        string    `json:"order_id"`
		UserID         int       `json:"user_id"`
		Symbol         string    `json:"symbol"`
		Side           string    `json:"side"`
		OrderType      string    `json:"order_type"`
		Price          float64   `json:"price"`
		Qty            int       `json:"qty"`
		TimeInForce    string    `json:"time_in_force"`
		OrderStatus    string    `json:"order_status"`
		LastExecPrice  int       `json:"last_exec_price"`
		CumExecQty     int       `json:"cum_exec_qty"`
		CumExecValue   int       `json:"cum_exec_value"`
		CumExecFee     int       `json:"cum_exec_fee"`
		ReduceOnly     bool      `json:"reduce_only"`
		CloseOnTrigger bool      `json:"close_on_trigger"`
		OrderLinkID    string    `json:"order_link_id"`
		CreatedTime    time.Time `json:"created_time"`
		UpdatedTime    time.Time `json:"updated_time"`
		TakeProfit     float64   `json:"take_profit"`
		StopLoss       float64   `json:"stop_loss"`
		TpTriggerBy    string    `json:"tp_trigger_by"`
		SlTriggerBy    string    `json:"sl_trigger_by"`
		PositionIdx    int       `json:"position_idx"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
}

type PostCancel struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		OrderID string `json:"order_id"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
}

type Isolated struct {
	RetCode          int         `json:"ret_code"`
	RetMsg           string      `json:"ret_msg"`
	ExtCode          string      `json:"ext_code"`
	Result           interface{} `json:"result"`
	ExtInfo          interface{} `json:"ext_info"`
	TimeNow          string      `json:"time_now"`
	RateLimitStatus  int         `json:"rate_limit_status"`
	RateLimitResetMs int64       `json:"rate_limit_reset_ms"`
	RateLimit        int         `json:"rate_limit"`
}

type StopLoss struct {
	RetCode          int         `json:"ret_code"`
	RetMsg           string      `json:"ret_msg"`
	ExtCode          string      `json:"ext_code"`
	ExtInfo          string      `json:"ext_info"`
	Result           interface{} `json:"result"`
	TimeNow          string      `json:"time_now"`
	RateLimitStatus  int         `json:"rate_limit_status"`
	RateLimitResetMs int64       `json:"rate_limit_reset_ms"`
	RateLimit        int         `json:"rate_limit"`
}
