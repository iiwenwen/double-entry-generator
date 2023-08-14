// 定义 translateToOrders 函数用于处理csv中每一行的数据，赋值给对应的对象
package bocm

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// translateToOrders 将CSV文件中的数据转换为订单对象。[]Order.
func (bocm *Bocm) translateToOrders(array []string) error {
	// 清理数据，去除字符串两侧的空格和制表符
	for idx, a := range array {
		a = strings.Trim(a, " ")
		a = strings.Trim(a, "\t")
		array[idx] = a
	}
	var bill Order // 创建一个交易对象
	var err error
	// 解析交易时间并赋值给订单对象
	bill.PayTime, err = time.Parse(localTimeFmt, array[1]+" +0800 CST") 
	if err != nil {
		return fmt.Errorf("parse create time %s error: %v", array[1], err)
	}
	// 赋值订单的其他属性

	// 处理收入和支出金额，处理连字符（--）和逗号分隔
	a4 := strings.ReplaceAll(array[4], "--", "")
	a5 := strings.ReplaceAll(array[5], "--", "")
	if a4 == "" && a5 == "" {
		bill.Type = OrderTypeUnknown
	} else if a4 == "" {
		bill.Type = OrderTypeRecv
		bill.Money, err = strconv.ParseFloat(a5, 64)
	} else {
		bill.Type = OrderTypeSend
		bill.Money, err = strconv.ParseFloat(a4, 64)
	}
	if err != nil {
		return fmt.Errorf("parse money [%s,%s] error: %v", array[4], array[5], err)
	}
	// 解析余额并赋值给订单对象，去除千位分隔符
	bill.Balances, _ = strconv.ParseFloat(strings.ReplaceAll(array[6], ",", ""), 64)
	// 赋值订单的其他属性
	bill.Peer = array[7]
	bill.PeerID = array[8]
	bill.TxTypeOriginal = array[9]
	// bill.Item = getItem(array[10])
	var item string = getItem(array[10])
	

	if ! strings.Contains(array[7], item) {
		bill.Item = item
	} 

	// 将订单对象添加到订单列表中
	bocm.Orders = append(bocm.Orders, bill)
	return nil
}


