package bocm

import (
	"regexp"
	"strings"
)

func getItem(s string) string {
	pattern := "订单编号\\d+\\s+(.+)\\s+交易流水号"
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(s)

	if len(match) > 1 {
		return match[1]
	}

	// 尝试提取 "退款"
	keyword := "退款"
	if strings.Contains(s, keyword) {
		return keyword
	}

	return s
}
