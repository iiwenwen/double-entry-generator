package ndnp

//定义用于解析和转换交通银行账单数据的提供者类 Ndnp。该类包括统计信息、行号和交易列表
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"

	"github.com/deb-sig/double-entry-generator/pkg/io/reader"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
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
	billReader, err := reader.GetReader(filename)
	if err != nil {
		return nil, fmt.Errorf("can't get bill reader, err: %v", err)
	}

	csvReader := csv.NewReader(billReader)
	csvReader.LazyQuotes = true
	// 如果 FieldsPerRecord 为负数，则不进行检查，记录可能有可变数量的字段
	csvReader.FieldsPerRecord = -1
	// 遍历 CSV 文件中的每一行，使用translateToOrders函数将csv数据放置ndnp.Orders。
	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		ndnp.LineNum++
		if ndnp.LineNum <= 1 {
			// 前3行是头信息，跳过它们。
			continue
		}
		// 将 CSV 行转换为 Order 实例。
		err = ndnp.translateToOrders(line)
		if err != nil {
			return nil, fmt.Errorf("Failed to translate bill: line %d: %v",
			ndnp.LineNum, err)
		}
	}
	log.Printf("Finished to parse the file %s", filename)
	return ndnp.convertToIR(), nil //将处理的数据转换为 IR 格式
}
