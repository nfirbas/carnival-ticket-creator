package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/go-pdf/fpdf"
)

var border = 8

func main() {
	numbers := 1800

	tickets(numbers)
	prizes(numbers)
}

func tickets(n int) {
	pdf := initPdf(60)

	pageWidth, pageHeight := pdf.GetPageSize()
	cellWidth := pageWidth / 2
	cellHeight := pageHeight / 6

	for i := 0; i < n; i += 6 {
		pdf.AddPage()
		pdf.CellFormat(cellWidth, cellHeight, toString(i+1), "R", 0, "CB", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, toString(i+2), "", 0, "CB", false, 0, "")
		pdf.SetXY(0, cellHeight*1)
		pdf.CellFormat(cellWidth, cellHeight, "", "R", 0, "CB", false, 0, "")
		pdf.SetXY(0, cellHeight*2)
		pdf.CellFormat(cellWidth, cellHeight, toString(i+3), "RT", 0, "CB", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, toString(i+4), "T", 0, "CB", false, 0, "")
		pdf.SetXY(0, cellHeight*3)
		pdf.CellFormat(cellWidth, cellHeight, "", "R", 0, "CB", false, 0, "")
		pdf.SetXY(0, cellHeight*4)
		pdf.CellFormat(cellWidth, cellHeight, toString(i+5), "RT", 0, "CB", false, 0, "")
		pdf.CellFormat(cellWidth, cellHeight, toString(i+6), "T", 0, "CB", false, 0, "")
		pdf.SetXY(0, cellHeight*5)
		pdf.CellFormat(cellWidth, cellHeight, "", "R", 0, "CB", false, 0, "")
	}

	createPdf(pdf, "tickets.pdf")
}

func prizes(n int) {
	pdf := initPdf(40)
	pageWidth, pageHeight := pdf.GetPageSize()

	pageHeight -= float64(border * 2)
	cellWidth := pageWidth / 3
	cellHeight := pageHeight / 8

	for i := 0; i < n; i += 24 {
		pdf.AddPage()
		for j := 0; j < 8; j++ {
			pdf.SetXY(0, float64(border+int(cellHeight)*j))
			pdf.CellFormat(cellWidth, cellHeight, toString(i+1+j*3), "", 0, "C", false, 0, "")
			pdf.CellFormat(cellWidth, cellHeight, toString(i+2+j*3), "", 0, "C", false, 0, "")
			pdf.CellFormat(cellWidth, cellHeight, toString(i+3+j*3), "", 0, "C", false, 0, "")
		}
	}

	createPdf(pdf, "prizes.pdf")
}

func initPdf(fontSize float64) *fpdf.Fpdf {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("openSans", "", "fonts/OpenSans-Regular.ttf")
	pdf.SetFont("openSans", "", fontSize)
	pdf.SetAutoPageBreak(false, 0)
	pdf.SetMargins(0, 0, 0)
	return pdf
}

func createPdf(pdf *fpdf.Fpdf, fileName string) {
	buf := &bytes.Buffer{}
	err := pdf.Output(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.Remove(fileName)
	os.WriteFile(fileName, buf.Bytes(), 0777)
}

func toString(value int) (out string) {
	out = strconv.Itoa(value)

	if value%10 == 9 || value%10 == 6 {
		out += "."
	}
	return
}
