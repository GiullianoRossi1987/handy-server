package utils

import (
	"fmt"
	types "types/database/operations"
	serial "types/responses"

	"codeberg.org/go-pdf/fpdf"
)

const (
	MAX_WIDTH       float64 = 210
	MAX_HEIGHT      float64 = 298
	SPACE           float64 = 20
	TABLE_W_PADDING float64 = 30
	TABLE_H_PADDING float64 = 30
)

// working on

func checkMax(width float64, height float64) {
	if width > MAX_WIDTH || height > MAX_HEIGHT {
		panic("PROGRAMMING ERROR - INVALID WXH REFERRED IN THE DOCUMENT")
	}
}

func pdfSetup(pdf fpdf.Pdf) {
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.SetAuthor("Handy - The freelancer finder app", true)
}

func header(pdf fpdf.Pdf, order serial.ComposedOrderResponse) {
	pdf.Cell(
		10,
		5,
		fmt.Sprintf("Generated data from order: %d", order.UsingOrder.Id),
	)

	pdf.Cell(
		10,
		5,
		fmt.Sprintf("Cart UUID: %s", order.UsingOrder.CartUUID),
	)
	pdf.Line(0, 20, MAX_WIDTH, 20)

}

func footer(pdf fpdf.Pdf) {

}

func toTable(order types.Order) {

}

func OrderToPDF(order *serial.ComposedOrderResponse) (string, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	// ts_gen := time.Now()
	out := "order_reports/test.pdf"
	pdfSetup(pdf)
	header(pdf, *order)
	err := pdf.OutputFileAndClose(out)
	return out, err
}
