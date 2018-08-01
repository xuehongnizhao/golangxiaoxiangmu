package excel

import (
	"encoding/csv"
	"os"
)

type ExcelOption struct {
	Count int
	Value string
}

//导出至Excel表格
//创建人:刘聪
//创建时间:2017年10月10日14:13:023
//备注:title[导出文件模板路径] headerstrings[表头]  contentstring[表行]
func BackupExcel(title string, headerstrings []string, contentstring [][]string) {
	f, err := os.Create(title)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	//设置表头
	w.Write(headerstrings)
	//设置数据
	for _, value := range contentstring {
		w.Write(value)
	}
	w.Flush()
}
