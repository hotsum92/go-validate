package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		[]string{"名前", "年齢", "出身地", "性別"},
		[]string{"山本", "24", "兵庫", "男"},
		[]string{"鈴木", "29", "神奈川", "女"},
		[]string{"佐藤", "27", "鹿児島", "男"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)
}
