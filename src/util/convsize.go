package util

import (
	"strings"
	"strconv"
	"fmt"
)

func SizeConv(size string) string {
	// e.g. convert 1KB to 1024
	if strings.Contains(size, "KB") {
		kb := strings.TrimSpace(strings.TrimSuffix(size, "KB"))
		fkb, _ := strconv.ParseFloat(kb, 64)
		return fmt.Sprintf("%f", fkb * 1024)
	} else if strings.Contains(size, "MB") {
		mb := strings.TrimSpace(strings.TrimSuffix(size, "MB"))
		fmb, _ := strconv.ParseFloat(mb, 64)
		return fmt.Sprintf("%f", fmb * 1024 * 1024)
	} else if strings.Contains(size, "GB") {
		gb := strings.TrimSpace(strings.TrimSuffix(size, "GB"))
		fgb, _ := strconv.ParseFloat(gb, 64)
		return fmt.Sprintf("%f", fgb * 1024 * 1024 * 1024)
	} else if strings.Contains(size, "TB") {
		tb := strings.TrimSpace(strings.TrimSuffix(size, "TB"))
		ftb, _ := strconv.ParseFloat(tb, 64)
		return fmt.Sprintf("%f", ftb * 1024 * 1024 * 1024 * 1024)
	} else {
		b := strings.TrimSpace(strings.TrimSuffix(size, "B"))
		return b
	}
}
