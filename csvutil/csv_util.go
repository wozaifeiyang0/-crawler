package csvutil

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Write(csvFile string, data [][]string)  {

	file, err := os.OpenFile(csvFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	//file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)

	//遍历数据，写每行到文件中
	for _, row := range data {
		w.Write(row)
		w.Flush()
	}

}

func Read()  {
	
}
