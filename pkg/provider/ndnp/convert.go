package ndnp

// 定义用于将交通银行账单数据转换为内部表示(IR)的过程
import (
	"github.com/deb-sig/double-entry-generator/pkg/ir"
)

// convertToIR 将交通银行账单交易信息转换为内部表示(IR)
func (ndnp *Ndnp) convertToIR() *ir.IR {
	i := ir.New() // 创建一个新的 IR 对象
	for _, o := range ndnp.Orders {
		irO := ir.Order{
			Peer:           o.Peer,		// 交易地点
			PayTime:        o.PayTime,		// 交易时间
			Money:          o.Money,		// 记账金额
			Type:           convertType(o.Type),		// 交易类型
			TypeOriginal:   string(o.Type),		// 原始交易类型
		}
		irO.Metadata = ndnp.getMetadata(o) 	// 获取交易的元数据
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
func (ndnp *Ndnp) getMetadata(o Order) map[string]string {
	// FIXME(TripleZ): hard-coded, bad pattern
	// source := "NDNP饭卡"
	var paytime string

	paytime = o.PayTime.Format(localTimeFmt)

	// 构建包含订单元数据的映射
	metadata := map[string]string{
		"paytime": 		paytime,
	}
	return metadata // 返回构建的元数据映射
}
