
package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在の時刻を取得
	now := time.Now()

	// 年、月、日を取得
	year, month, day := now.Date()

	// 年-月-日の形式で表示
	fmt.Printf("日付: %d-%02d-%02d\n", year, month, day)
}
