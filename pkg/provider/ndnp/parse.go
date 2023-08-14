// 定义 translateToOrders 函数用于处理csv中每一行的数据，赋值给对应的对象
package ndnp

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// translateToOrders 将CSV文件中的数据转换为订单对象。[]Order.
func (ndnp *Ndnp) translateToOrders(array []string) error {
	// 清理数据，去除字符串两侧的空格和制表符
	for idx, a := range array {
		a = strings.Trim(a, " ")
		a = strings.Trim(a, "\t")
		array[idx] = a
	}
	var bill Order // 创建一个交易对象
	var err error
	// 解析交易时间并赋值给订单对象
	bill.PayTime, err = time.Parse(localTimeFmt, array[0]+" +0800 CST") 
	if err != nil {
		return fmt.Errorf("parse create time %s error: %v", array[0], err)
	}
	// 赋值订单的其他属性
	bill.Peer = array[12]
	// 处理收入和支出金额，处理连字符（--）和逗号分隔
	bill.Money, err = strconv.ParseFloat(strings.ReplaceAll(array[5], "-", ""),64)
	if err != nil {
		return fmt.Errorf("parse money [%s] error: %v", array[5], err)
	}
	
	bill.Type = OrderTypeSend
	
	if err != nil {
		
	}
	// 解析余额并赋值给订单对象，去除千位分隔符
	bill.Balances, _ = strconv.ParseFloat(strings.ReplaceAll(array[6], ",", ""), 64)
	// 赋值订单的其他属性
	// 将订单对象添加到订单列表中
	ndnp.Orders = append(ndnp.Orders, bill)
	return nil
}
