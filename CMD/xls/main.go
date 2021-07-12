package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f1, err := excelize.OpenFile("/home/eintr/下载/sxjk_数据/收费门架（主）.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	f2, err := excelize.OpenFile("/home/eintr/下载/sxjk_数据/收费门架（副）.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	f3, err := excelize.OpenFile("/home/eintr/下载/sxjk_数据/2020年山西省行政区划最新版.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	areaMap := make(map[string]string, 100)
	rows, err := f3.GetRows("2021年")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 2; i <= len(rows); i++ {
		src1 := fmt.Sprintf("C%d", i)
		srcCell1, _ := f3.GetCellValue("2021年", src1)
		src2 := fmt.Sprintf("J%d", i)
		srcCell2, _ := f3.GetCellValue("2021年", src2)
		src3 := fmt.Sprintf("Q%d", i)
		srcCell3, _ := f3.GetCellValue("2021年", src3)

		dst1 := fmt.Sprintf("B%d", i)
		dstCell1, _ := f3.GetCellValue("2021年", dst1)
		dst2 := fmt.Sprintf("I%d", i)
		dstCell2, _ := f3.GetCellValue("2021年", dst2)
		dst3 := fmt.Sprintf("P%d", i)
		dstCell3, _ := f3.GetCellValue("2021年", dst3)

		areaMap[srcCell1] = dstCell1
		areaMap[srcCell2] = dstCell2
		areaMap[srcCell3] = dstCell3
		fmt.Println(srcCell1, srcCell2, srcCell3, dstCell1, dstCell2, dstCell3)
	}
	fmt.Println(len(rows), areaMap)

	test := make(map[string]string, 1200)
	rows, err = f2.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		test[row[0]] = row[2]
	}

	rows, _ = f1.GetRows("收费门架信息")
	for i := 2; i <= len(rows); i++ {
		srcIndex := fmt.Sprintf("B%d", i)
		dstIndex := fmt.Sprintf("D%d", i)
		areaIndex := fmt.Sprintf("E%d", i)
		srcCell, err := f1.GetCellValue("收费门架信息", srcIndex)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = f1.SetCellStr("收费门架信息", dstIndex, test[srcCell])
		if err != nil {
			fmt.Println(err)
			return
		}

		err = f1.SetCellStr("收费门架信息", areaIndex, areaMap[test[srcCell]])
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	f1.SaveAs("test.xlsx", excelize.Options{})
}
