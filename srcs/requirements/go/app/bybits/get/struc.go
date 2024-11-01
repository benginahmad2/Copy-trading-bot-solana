package get

import "time"

type Price struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []struct {
		Symbol                 string    `json:"symbol"`
		BidPrice               string    `json:"bid_price"`
		AskPrice               string    `json:"ask_price"`
		LastPrice              string    `json:"last_price"`
		LastTickDirection      string    `json:"last_tick_direction"`
		PrevPrice24H           string    `json:"prev_price_24h"`
		Price24HPcnt           string    `json:"price_24h_pcnt"`
		HighPrice24H           string    `json:"high_price_24h"`
		LowPrice24H            string    `json:"low_price_24h"`
		PrevPrice1H            string    `json:"prev_price_1h"`
		Price1HPcnt            string    `json:"price_1h_pcnt"`
		MarkPrice              string    `json:"mark_price"`
		IndexPrice             string    `json:"index_price"`
		OpenInterest           float64   `json:"open_interest"`
		OpenValue              string    `json:"open_value"`
		TotalTurnover          string    `json:"total_turnover"`
		Turnover24H            string    `json:"turnover_24h"`
		TotalVolume            float64   `json:"total_volume"`
		Volume24H              float64   `json:"volume_24h"`
		FundingRate            string    `json:"funding_rate"`
		PredictedFundingRate   string    `json:"predicted_funding_rate"`
		NextFundingTime        time.Time `json:"next_funding_time"`
		CountdownHour          int       `json:"countdown_hour"`
		DeliveryFeeRate        string    `json:"delivery_fee_rate"`
		PredictedDeliveryPrice string    `json:"predicted_delivery_price"`
		DeliveryTime           string    `json:"delivery_time"`
	} `json:"result"`
	TimeNow string `json:"time_now"`
}

type Wallet struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		Usdt struct {
			Equity           float64 `json:"equity"`
			AvailableBalance float64 `json:"available_balance"`
			UsedMargin       float64 `json:"used_margin"`
			OrderMargin      float64 `json:"order_margin"`
			PositionMargin   float64 `json:"position_margin"`
			OccClosingFee    float64 `json:"occ_closing_fee"`
			OccFundingFee    float64 `json:"occ_funding_fee"`
			WalletBalance    float64 `json:"wallet_balance"`
			RealisedPnl      float64 `json:"realised_pnl"`
			UnrealisedPnl    float64 `json:"unrealised_pnl"`
			CumRealisedPnl   float64 `json:"cum_realised_pnl"`
			GivenCash        float64 `json:"given_cash"`
			ServiceCash      float64 `json:"service_cash"`
		} `json:"USDT"`
	} `json:"result"`
	TimeNow         string `json:"time_now"`
	RateLimitStatus int    `json:"rate_limit_status"`
}

type Position struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []struct {
		UserID              int     `json:"user_id"`
		Symbol              string  `json:"symbol"`
		Side                string  `json:"side"`
		Size                int     `json:"size"`
		PositionValue       int     `json:"position_value"`
		EntryPrice          int     `json:"entry_price"`
		LiqPrice            int     `json:"liq_price"`
		BustPrice           int     `json:"bust_price"`
		Leverage            int     `json:"leverage"`
		AutoAddMargin       int     `json:"auto_add_margin"`
		IsIsolated          bool    `json:"is_isolated"`
		PositionMargin      int     `json:"position_margin"`
		OccClosingFee       int     `json:"occ_closing_fee"`
		RealisedPnl         int     `json:"realised_pnl"`
		CumRealisedPnl      float64 `json:"cum_realised_pnl"`
		FreeQty             int     `json:"free_qty"`
		TpSlMode            string  `json:"tp_sl_mode"`
		UnrealisedPnl       int     `json:"unrealised_pnl"`
		DeleverageIndicator int     `json:"deleverage_indicator"`
		RiskID              int     `json:"risk_id"`
		StopLoss            int     `json:"stop_loss"`
		TakeProfit          int     `json:"take_profit"`
		TrailingStop        int     `json:"trailing_stop"`
		PositionIdx         int     `json:"position_idx"`
		Mode                string  `json:"mode"`
		TpTriggerBy         int     `json:"tp_trigger_by,omitempty"`
		SlTriggerBy         int     `json:"sl_trigger_by,omitempty"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
}

type PositonTrade struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		CurrentPage int `json:"current_page"`
		Data        []struct {
			OrderID        string    `json:"order_id"`
			UserID         int       `json:"user_id"`
			Symbol         string    `json:"symbol"`
			Side           string    `json:"side"`
			OrderType      string    `json:"order_type"`
			Price          float64   `json:"price"`
			Qty            int       `json:"qty"`
			TimeInForce    string    `json:"time_in_force"`
			OrderStatus    string    `json:"order_status"`
			LastExecPrice  float64   `json:"last_exec_price"`
			CumExecQty     int       `json:"cum_exec_qty"`
			CumExecValue   float64   `json:"cum_exec_value"`
			CumExecFee     float64   `json:"cum_exec_fee"`
			ReduceOnly     bool      `json:"reduce_only"`
			CloseOnTrigger bool      `json:"close_on_trigger"`
			OrderLinkID    string    `json:"order_link_id"`
			CreatedTime    time.Time `json:"created_time"`
			UpdatedTime    time.Time `json:"updated_time"`
			TakeProfit     float64   `json:"take_profit"`
			StopLoss       float64   `json:"stop_loss"`
			TpTriggerBy    string    `json:"tp_trigger_by"`
			SlTriggerBy    string    `json:"sl_trigger_by"`
		} `json:"data"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
}
