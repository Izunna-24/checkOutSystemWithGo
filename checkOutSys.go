package main

import (
    "fmt"
    "time"
)

type CheckOutSys struct {
    customerName string
    cashierName  string
    subTotal     float64
    discount     float64
    amountPaid   float64
}

func (c *CheckOutSys) showCurrentDateTime() string {
    currentDateTime := time.Now()
    dateTimeFormatter := "02-Jan-06 03:04:05 PM"
    
    currentDateAndTime := currentDateTime.Format(dateTimeFormatter)
    return currentDateAndTime
}

func (c *CheckOutSys) setCustomerName(customerName string) {
    c.customerName = customerName
}

func (c *CheckOutSys) getCustomerName() string {
    return c.customerName
}

func (c *CheckOutSys) setCashierName(cashierName string) {
    c.cashierName = cashierName
}

func (c *CheckOutSys) getCashierName() string {
    return c.cashierName
}

func (c *CheckOutSys) setSubTotal(subTotal float64) {
    c.subTotal = subTotal
}

func (c *CheckOutSys) getSubTotal() float64 {
    return c.subTotal
}

func (c *CheckOutSys) setDiscount(discount float64) {
    if discount < 0 || discount > 100 {
        panic("Discount must be between 0 and 100")
    }
    c.discount = discount
}

func (c *CheckOutSys) getDiscount() float64 {
    return c.discount
}

func (c *CheckOutSys) getDiscountedTotal() float64 {
    return (c.discount / 100.0) * c.subTotal
}

func (c *CheckOutSys) getVAT() float64 {
    return c.subTotal * (7.5 / 100)
}

func (c *CheckOutSys) getBillTotal() float64 {
    return (c.subTotal - c.getDiscountedTotal()) + c.getVAT()
}

func (c *CheckOutSys) setAmountPaid(amountPaid float64) {
    if amountPaid < c.getBillTotal() {
        panic("Amount paid cannot be less than bill total")
    }
    c.amountPaid = amountPaid
}

func (c *CheckOutSys) getAmountPaid() float64 {
    return c.amountPaid
}

func (c *CheckOutSys) getBalance() float64 {
    var balance float64
    if c.amountPaid > c.getBillTotal() {
        return c.amountPaid - c.getBillTotal()
    }
    return balance
}

func main() {
    fmt.Print("Enter customer's name: ")
		var customerName string
		fmt.Scanln(&customerName)

    fmt.Print("What did ", customerName , " buy ?")
		var itemBought string
		fmt.Scanln(&itemBought)
		
		fmt.Print("How many pieces of ", itemBought,"?")
		var numOfItemBought string
		fmt.Scanln(&numOfItemBought)

		fmt.Print("Add more item? ")
		var moreItem string
		fmt.Scanln(&moreItem)

		fmt.Print("What did the user buy ")
		var  itemAdded string
		fmt.Scanln(&itemAdded)

		fmt.Print("How many pieces? ")
		var numOfItem2 string
		fmt.Scanln(&numOfItem2)

		fmt.Print("How much per unit")
		var itemAmount int
		fmt.Scanln(&itemAmount)

		fmt.Print("Add more item? (yes/no) ")
		var  itemAdded3 string
		fmt.Scanln(&itemAdded3)

		fmt.Print("What is your name ")
		var  cashierName string
		fmt.Scanln(&cashierName)

		fmt.Print("How much discount will ", customerName, "get?")
		var discount string
		fmt.Scanln(&discount)

		
		
		
		
		
		
		checkout := CheckOutSys{}
    checkout.setSubTotal(100)
    checkout.setDiscount(10)
    checkout.setAmountPaid(105)

    fmt.Println("Customer:", checkout.getCustomerName())
    fmt.Println("Cashier:", checkout.getCashierName())
    fmt.Println("Subtotal:", checkout.getSubTotal())
    fmt.Println("Discounted Total:", checkout.getDiscountedTotal())
    fmt.Println("VAT:", checkout.getVAT())
    fmt.Println("Bill Total:", checkout.getBillTotal())
    fmt.Println("Amount Paid:", checkout.getAmountPaid())
    fmt.Println("Balance:", checkout.getBalance())
    fmt.Println("Current Date and Time:", checkout.showCurrentDateTime())
}
