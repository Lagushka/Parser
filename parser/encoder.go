package parser

import (	
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
	//"log"
	//"net/http"
	//"io/ioutil"
	//"fmt"
	//"os"
)

type File struct {
	file		*excelize.File 
}

func (f *File) setValue(cell string, name string) {
	f.file.SetCellValue("Sheet1", cell, name)
}

func (f *File) SaveAs(name string, opt ...excelize.Options) error {
	return f.file.SaveAs(name, opt...)
}

func Create_xls() *excelize.File {
	data := parseMainPage()
	data = parsePages(data)
	/*for _, elem := range data {
		log.Printf("%s\n", elem.PhoneNumber)
	}*/
	f := excelize.NewFile()
	file := File{ file: f }

	file.setValue("A1", "Name")
	file.setValue("B1", "Price")
	file.setValue("C1", "Place")
	file.setValue("D1", "Product ref")
	file.setValue("E1", "PhoneNumber")

	for i, val := range data {
		file.setValue("A" + strconv.Itoa(i+2), val.Name)
		file.setValue("B" + strconv.Itoa(i+2), val.Price)
		file.setValue("C" + strconv.Itoa(i+2), val.Place)
		file.setValue("D" + strconv.Itoa(i+2), val.Ref)
		file.setValue("E" + strconv.Itoa(i+2), val.PhoneNumber)
	}

	return f
}