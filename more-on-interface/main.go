package main

import (
	"errors"
	"fmt"
)

// ============================================================================
// 1. Polymorphism & Dependency Injection
// ============================================================================
// The PaymentMethod interface defines a contract. Any struct that implements
// the Pay method automatically satisfies this interface.

type PaymentMethod interface {
	Pay(amount float64) error
}

// Concrete Implementation 1: Bkash
type Bkash struct {
	Phone  string
	APIKey string
}

func (b Bkash) Pay(amount float64) error {
	fmt.Printf("[Bkash] Paid $%.2f using account %s\n", amount, b.Phone)
	return nil
}

// Concrete Implementation 2: Nagad
type Nagad struct {
	Phone string
}

func (n Nagad) Pay(amount float64) error {
	fmt.Printf("[Nagad] Paid $%.2f using account %s\n", amount, n.Phone)
	return nil
}

// Concrete Implementation 3: CreditCard (with pointer receiver)
type CreditCard struct {
	CardNumber string
	CVV        string
}

func (c *CreditCard) Pay(amount float64) error {
	fmt.Printf("[CreditCard] Paid $%.2f using card ending in %s\n", amount, c.CardNumber[len(c.CardNumber)-4:])
	return nil
}

// PaymentService depends on the abstraction (interface), NOT concrete types.
type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{method: method}
}

// SetMethod allows switching payment methods at runtime.
func (ps *PaymentService) SetMethod(method PaymentMethod) {
	ps.method = method
}

func (ps *PaymentService) Checkout(amount float64) error {
	if ps.method == nil {
		return errors.New("no payment method configured")
	}
	return ps.method.Pay(amount)
}

func runPolymorphismExample() {
	fmt.Println("=== 1. Polymorphism & Dependency Injection ===")

	bkash := Bkash{Phone: "01700000000", APIKey: "bkash-secret-key"}
	nagad := Nagad{Phone: "01800000000"}
	card := &CreditCard{CardNumber: "1234567890123456", CVV: "123"}

	// Inject Bkash
	service := NewPaymentService(bkash)
	service.Checkout(150.50)

	// Switch to Nagad dynamically
	service.SetMethod(nagad)
	service.Checkout(275.00)

	// Switch to CreditCard
	service.SetMethod(card)
	service.Checkout(999.99)

	fmt.Println()
}

// ============================================================================
// 2. Interface Composition (Embedding Interfaces)
// ============================================================================
// In Go, interfaces can be composed together into larger interfaces.

type Payer interface {
	Pay(amount float64) error
}

type Refunder interface {
	Refund(txID string, amount float64) error
}

// AdvancedPaymentGateway embeds both Payer and Refunder
type AdvancedPaymentGateway interface {
	Payer
	Refunder
}

// PayPal implements both Pay and Refund
type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float64) error {
	fmt.Printf("[PayPal] Paid $%.2f via %s\n", amount, p.Email)
	return nil
}

func (p PayPal) Refund(txID string, amount float64) error {
	fmt.Printf("[PayPal] Refunded $%.2f for Transaction ID: %s\n", amount, txID)
	return nil
}

func processRefund(r Refunder, txID string, amount float64) {
	fmt.Println("Processing refund via Refunder interface:")
	_ = r.Refund(txID, amount)
}

func runInterfaceCompositionExample() {
	fmt.Println("=== 2. Interface Composition (Embedding) ===")

	paypal := PayPal{Email: "user@example.com"}

	// PayPal can be passed to functions expecting AdvancedPaymentGateway, Payer, or Refunder
	var gateway AdvancedPaymentGateway = paypal
	gateway.Pay(500.0)
	processRefund(paypal, "TXN-998811", 500.0)

	fmt.Println()
}

// ============================================================================
// 3. Type Assertions & Type Switches
// ============================================================================
// Use type assertion to check if an interface value holds a specific concrete
// type or implements an optional feature interface.

func inspectPaymentMethod(method PaymentMethod) {
	// Check if this method also implements Refunder (optional capability)
	if refunder, ok := method.(Refunder); ok {
		fmt.Println("-> This payment method supports refunds!")
		_ = refunder.Refund("AUTO-TX-01", 50.0)
	} else {
		fmt.Println("-> This payment method DOES NOT support refunds.")
	}

	// Type switch to handle specific concrete types
	switch v := method.(type) {
	case Bkash:
		fmt.Printf("   Details: Mobile banking via Bkash (%s)\n", v.Phone)
	case Nagad:
		fmt.Printf("   Details: Mobile banking via Nagad (%s)\n", v.Phone)
	case *CreditCard:
		fmt.Printf("   Details: Card ending in %s\n", v.CardNumber[len(v.CardNumber)-4:])
	case PayPal:
		fmt.Printf("   Details: Digital wallet PayPal (%s)\n", v.Email)
	default:
		fmt.Println("   Details: Unknown payment provider")
	}
}

func runTypeAssertionExample() {
	fmt.Println("=== 3. Type Assertions & Type Switches ===")

	bkash := Bkash{Phone: "01711111111", APIKey: "key-1"}
	paypal := PayPal{Email: "buyer@domain.com"}

	fmt.Println("Inspecting Bkash:")
	inspectPaymentMethod(bkash)

	fmt.Println("\nInspecting PayPal:")
	inspectPaymentMethod(paypal)

	fmt.Println()
}

// ============================================================================
// 4. Mocking for Unit Testing (Why Interfaces are Essential)
// ============================================================================
// Interfaces allow substituting real payment processors with mock objects during testing.

type MockPaymentGateway struct {
	TotalPaid    float64
	CallCount    int
	ShouldFail   bool
}

func (m *MockPaymentGateway) Pay(amount float64) error {
	if m.ShouldFail {
		return errors.New("mock network error")
	}
	m.CallCount++
	m.TotalPaid += amount
	fmt.Printf("[MockGateway] Simulated payment of $%.2f (Calls: %d)\n", amount, m.CallCount)
	return nil
}

func runMockTestingExample() {
	fmt.Println("=== 4. Mocking with Interfaces for Testing ===")

	mock := &MockPaymentGateway{}
	service := NewPaymentService(mock)

	// Simulate checkout in tests without hitting real payment APIs
	_ = service.Checkout(75.0)
	_ = service.Checkout(25.0)

	fmt.Printf("Test Verification: Total paid is $%.2f, Calls made: %d\n", mock.TotalPaid, mock.CallCount)
	fmt.Println()
}

// ============================================================================
// Main Function
// ============================================================================
func main() {
	// 1. Dependency Injection with multiple implementations
	runPolymorphismExample()

	// 2. Composing multiple smaller interfaces
	runInterfaceCompositionExample()

	// 3. Type assertions and type switches
	runTypeAssertionExample()

	// 4. Mocking for unit tests
	runMockTestingExample()
}
