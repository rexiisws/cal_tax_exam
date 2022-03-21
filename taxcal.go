package main

import "fmt"
// import "humanize"

const perExpense float64 = 60000.00
const ltfLimit float64 = 200000.00
const insurLimit float64 = 100000.00

func main() {
	
	// total := calTotalIncome("ธเนศ", 2120000.00, 300000.00, 150000.00)
	// fmt.Println(calTotalIncome(2120000.00, 300000.00, 150000.00))
	// fmt.Printf("เงินได้สุทธิ %.2f บาท" , total)
	// fmt.Printf("ภาษี %.2f บาท" , total)

	calTax("ธเนศ", 2120000.00, 300000.00, 150000.00)
	calTax("จุฬาลักษณ์", 845000.00, 0.00, 30000.00)
	calTax("สุขเกษม", 1050000.00, 25000.00, 80000.00)
	calTax("สุพร", 5550000.00, 1200000.00, 400000.00)

}

func calTax (name string, income, ltf, insurance  float64 ) {
	var tax , tax1 , tax2 , tax3 , tax4 , tax5 , tax6 , tax7 , tax8 , incomeBal float64

	if ltf >= ltfLimit {
		ltf = ltfLimit
	}
	if  insurance >= insurLimit {
		insurance = insurLimit
	}
	incomeBal = income - ltf - insurance - perExpense

//สมมติเงินได้ 1,000,000
	if incomeBal <= 150000  {
		tax1 = incomeBal * 0.05	
	} else if incomeBal > 150000 {
		tax1 = 150000 * 0.05
	}

	if (incomeBal > 150000) && (incomeBal <= 300000)  {
		tax2 = (incomeBal - 150000) * 0.05
	} else if incomeBal > 300000 {
		tax2 = 150000 * 0.05
	}
	
	if  (incomeBal > 300000) && (incomeBal <= 500000) {
		tax3 = (incomeBal - 300000) * 0.1
	} else if incomeBal > 500000 {
		tax3 = 200000 * 0.1
	}

	if (incomeBal > 500000) && (incomeBal <= 750000 ){
		tax4 = (incomeBal - 500000) * 0.15
	} else if incomeBal > 750000 {
		tax4 =	250000 * 0.15
	}

	if (incomeBal > 750000) && (incomeBal <= 1000000) {
		tax5 = (incomeBal - 750000) * 0.2
	} else if incomeBal > 1000000 {
		tax5 =	250000 * 0.2
	}
	
	if (incomeBal > 1000000) && (incomeBal <= 2000000) {
		tax6 = (incomeBal - 1000000) * 0.25
	} else if incomeBal > 2000000 {
		tax6 =	1000000 * 0.25
	}
	
	if (incomeBal > 2000000) && (incomeBal <= 5000000) {
		tax7 = (incomeBal -2000000) * 0.3
	} else if incomeBal > 5000000 {
		tax7 =	1000000 * 0.3
	}

	if incomeBal > 5000000 {
		tax8 = (incomeBal -5000000) * 0.35
	} 

	tax = tax1 + tax2 + tax3 + tax4 + tax5 + tax6 + tax7 + tax8

	fmt.Printf("%s เงินได้สุทธิ %.2f เงินภาษีที่ต้องจ่ายคือ %.2f\n" , name ,incomeBal, tax)

	// return tax
}

