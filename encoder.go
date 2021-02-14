package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
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

func main() {
	data := parseMainPage()
	data = parsePages(data)
	for _, elem := range data {
		log.Printf("%s\n", elem.PhoneNumber)
	}
	f := excelize.NewFile()
	file := File{ file: f }

	file.setValue("A1", "Name")
	file.setValue("B1", "Price")
	file.setValue("C1", "Img ref")
	file.setValue("D1", "Product ref")
	file.setValue("E1", "PhoneNumber")

	for i, val := range data {
		file.setValue("A" + strconv.Itoa(i+2), val.Name)
		file.setValue("B" + strconv.Itoa(i+2), val.Price)
		file.setValue("C" + strconv.Itoa(i+2), val.Img)
		file.setValue("D" + strconv.Itoa(i+2), val.Ref)
		file.setValue("E" + strconv.Itoa(i+2), val.PhoneNumber)
	}

	if err := file.SaveAs("Book1.xlsx"); err != nil {
        log.Panic(err)
	}

	resp, err := http.Get("https://www.olx.kz/obyavlenie/kartofel-sorta-gala-domashniy-dostavka-IDjVPUY.html#e7d8a7f4bc;promoted")
	if err != nil {
		log.Panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	page, err := os.Create("page.html")
	if err != nil {
		log.Panic(err)
	}
	defer page.Close()
	fmt.Fprint(page, string(b))
	
}