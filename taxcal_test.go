package main

import (
	"testing"
)

func TestTaxCal(t *testing.T) {
	name := "ธเนศ"
	income := 2120000.00
	ltf := 300000.00
	insurance := 150000.00

	wantTax := 312500.00

	getTax := calTax(name, income, ltf, insurance)

	if wantTax != getTax {
		t.Errorf("ภาษีที่ต้องการ %.2f ไม่เท่ากับ ภาษีที่คำนวณได้ %.2f", wantTax, getTax)
	}
}

func TestTaxCal2(t *testing.T) {
	name := "จุฬาลักษณ์"
	income := 845000.00
	ltf := 0.00
	insurance := 30000.00

	wantTax := 73500.00

	getTax := calTax(name, income, ltf, insurance)

	if wantTax != getTax {
		t.Errorf("ภาษีที่ต้องการ %.2f ไม่เท่ากับ ภาษีที่คำนวณได้ %.2f", wantTax, getTax)
	}
}
