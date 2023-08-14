package bocm

// 定义交通银行账单的结构体和常量
import "time"

const (
	// localTimeFmt set time format to utc+8
	localTimeFmt = "2006-01-02 15:04:05 -0700 CST"
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
	PayTime         time.Time       // 交易时间
	Peer            string          // 对方户名
	Method			string 			// 交易方式
	Item        	string          // 摘要
	Money           float64   		// 记账金额 (收入/支出)
	Type            OrderType 		// 收/支 (数据中无该列，推测而来)
	Balances        float64   		// 余额  
	PeerAccountName string          // 对方户名
	PeerID      	string          // 对方账户
	PeerAccount     string     		// 对方开户行
	TxTypeOriginal  string    		// 交易类型

}

// OrderType 表示交易的类型枚举
type OrderType string

const (
	OrderTypeSend    OrderType = "支出"
	OrderTypeRecv    OrderType = "收入"
	OrderTypeUnknown OrderType = "Unknown"
)
