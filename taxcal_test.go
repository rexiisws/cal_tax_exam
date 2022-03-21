package taxcal_test

import "testing"
import "fmt"
import "main"

func TestTaxCal (t *testing.T) {
	name := "ธเนศ"
	income := 2120000.00
	ltf := 300000.00
	insurance := 150000.00

	get := calTotalIncome(name,income,ltf,insurance)
	
	fmt.Printf("%s เงินภาษีที่ต้องจ่ายคือ %.2f\n" , name , tax)

}