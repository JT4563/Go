package main

import "fmt"

// Payment is an interface that defines payment behavior
type Payment interface {
    ProcessPayment() bool
    GetBalance() float64
}

// PhonePay implements the Payment interface
type PhonePay struct {
    userID      string
    balance     float64
    phoneNumber string
}

// ProcessPayment processes a payment on PhonePay
func (p *PhonePay) ProcessPayment(amount float64) bool {
    if p.balance >= amount {
        p.balance -= amount
        fmt.Printf("Payment of ₹%.2f processed successfully via PhonePay\n", amount)
        return true
    }
    fmt.Println("Insufficient balance in PhonePay")
    return false
}

// GetBalance returns the current balance
func (p *PhonePay) GetBalance() float64 {
    return p.balance
}

// AddMoney adds money to the PhonePay wallet
func (p *PhonePay) AddMoney(amount float64) {
    p.balance += amount
    fmt.Printf("₹%.2f added to your PhonePay wallet\n", amount)
}

func main() {
    // Create a new PhonePay instance
    myPhonePay := &PhonePay{
        userID:      "user123",
        balance:     1000.0,
        phoneNumber: "9876543210",
    }

    // Check initial balance
    fmt.Printf("Initial balance: ₹%.2f\n", myPhonePay.GetBalance())

    // Add money to wallet
    myPhonePay.AddMoney(500)
    
    // Check updated balance
    fmt.Printf("Updated balance: ₹%.2f\n", myPhonePay.GetBalance())

    // Process a payment
    myPhonePay.ProcessPayment(750.0)

    // Check final balance
    fmt.Printf("Final balance: ₹%.2f\n", myPhonePay.GetBalance())
    
    // Try processing a payment that exceeds balance
    myPhonePay.ProcessPayment(1000.0)
}