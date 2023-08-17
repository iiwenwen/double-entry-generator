package ndnp

//定义用于解析和转换交通银行账单数据的提供者类 Ndnp。该类包括统计信息、行号和交易列表
import (
	"fmt"
	"log"

	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"github.com/xuri/excelize/v2"
)

// Ndnp 是交通银行（NDNP）账户交易数据的提供者。
type Ndnp struct {
	Statistics Statistics `json:"statistics,omitempty"`  // 与处理数据相关的统计信息。
	LineNum    int        `json:"line_num,omitempty"` // 当前处理的文件行号。
	Orders     []Order    `json:"orders,omitempty"` // 表示交易的 Order 实例数组。
}

// 创建一个新的 Ndnp 提供者。
func New() *Ndnp {
	return &Ndnp{
		Statistics: Statistics{},
		LineNum:    0,
		Orders:     make([]Order, 0),
	}
}

// // Translate 读取并将交通银行（NDNP）账户交易记录转换为 IR 格式。
func (ndnp *Ndnp) Translate(filename string) (*ir.IR, error) {
	log.SetPrefix("[Provider-NDNP] ")
	// 获取提供的文件的读取器。
	xlsxFile, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	rows, err := xlsxFile.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	// 遍历 CSV 文件中的每一行，使用translateToOrders函数将csv数据放置ndnp.Orders。
	for _, row := range rows {
		ndnp.LineNum++
		if ndnp.LineNum <= 1 {
			// The first line is xlsx file header.
			continue
		}

		err = ndnp.translateToOrders(row)
		if err != nil {
			return nil, fmt.Errorf("Failed to translate bill: line %d: %v", ndnp.LineNum, err)
		}
	}
	log.Printf("Finished to parse the file %s", filename)
	return ndnp.convertToIR(), nil
}
