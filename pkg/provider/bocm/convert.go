package bocm

// 定义用于将交通银行账单数据转换为内部表示(IR)的过程
import (
	"github.com/deb-sig/double-entry-generator/pkg/ir"
)

// convertToIR 将交通银行账单交易信息转换为内部表示(IR)
func (bocm *Bocm) convertToIR() *ir.IR {
	i := ir.New() // 创建一个新的 IR 对象
	for _, o := range bocm.Orders {
		irO := ir.Order{
			Peer:           o.Peer,		// 交易地点
			Item :      	o.Item,		// 摘要
			PayTime:        o.PayTime,		// 交易时间
			Money:          o.Money,		// 记账金额
			Type:           convertType(o.Type),		// 收支
			TypeOriginal:   string(o.Type),		// 原始交易类型
			TxTypeOriginal: o.TxTypeOriginal, //交易类型
		}
		irO.Metadata = bocm.getMetadata(o) 	// 获取交易的元数据
		i.Orders = append(i.Orders, irO)		//将转换后的交易信息添加到 IR 对象中
	}
	return i //返回填充了交易数据的 IR 对象
}

// convertType 将订单类型转换为内部表示 IR 类型。
func convertType(t OrderType) ir.Type {
	switch t {
	case OrderTypeSend:
		return ir.TypeSend
	case OrderTypeRecv:
		return ir.TypeRecv
	default:
		return ir.TypeUnknown
	}
}

// getMetadata 从交易中获取元数据(例如状态、方法、类别等)。
//  from order.
func (bocm *Bocm) getMetadata(o Order) map[string]string {
	// FIXME(TripleZ): hard-coded, bad pattern
	source := "中国交通银行"
	var guessedType, paytime ,method, txType string

	paytime = o.PayTime.Format(localTimeFmt)

	if o.Type != "" {
		guessedType = string(o.Type)
	}

	if o.Method != "" {
		method = string(o.Method)
	}
	if o.TxTypeOriginal != "" {
		txType = string(o.TxTypeOriginal)
	}
	// 构建包含订单元数据的映射
	metadata := map[string]string{
		"source":      source,
		"type":        guessedType,
		"paytime":		paytime,
		"method":		method,
		"txType": 		txType,
	}
	return metadata // 返回构建的元数据映射
}
