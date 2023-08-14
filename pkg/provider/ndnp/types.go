package ndnp

// 定义交通银行账单的结构体和常量
import "time"

const (
	// localTimeFmt set time format to utc+8
	localTimeFmt = "2006/1/2 15:04 -0700 CST"
)

// Statistics 结构体表示账单文件的统计信息
type Statistics struct {
	UserID          string    `json:"user_id,omitempty"`
	Username        string    `json:"username,omitempty"`
	ParsedItems     int       `json:"parsed_items,omitempty"`
	Start           time.Time `json:"start,omitempty"`
	End             time.Time `json:"end,omitempty"`
	TotalInRecords  int       `json:"total_in_records,omitempty"`
	TotalInMoney    float64   `json:"total_in_money,omitempty"`
	TotalOutRecords int       `json:"total_out_records,omitempty"`
	TotalOutMoney   float64   `json:"total_out_money,omitempty"`
}

// Order 结构体用于表示单笔交易订单的各个字段
type Order struct {
	PayTime         time.Time       // 交易发生日期
	Peer            string          // 商户名称
	Money           float64   		// 记账金额 (收入/支出)
	Type            OrderType 		// 收/支 (数据中无该列，推测而来)
	Balances        float64   		// 余额  
}

// OrderType 表示交易的类型枚举
type OrderType string

const (
	OrderTypeSend    OrderType = "支出"
	OrderTypeRecv    OrderType = "收入"
	OrderTypeUnknown OrderType = "Unknown"
)
