# Abstraction in OOP

**Abstraction** means exposing what an object can do while hiding the
details of how it does it. It lets code depend on a simple contract
instead of a specific implementation.

Another meaning: **Abstraction is the process of hiding unnecessary implementation details and exposing only the essential functionality to the outside world.**

For example, a checkout service needs to process a payment. It should
call a payment operation without needing to know whether the provider is
Stripe or Razorpay.

## TypeScript

Define a contract with an interface, then provide different
implementations:

```ts
interface PaymentProcessor {
  pay(amount: number): void;
}

class StripePayment implements PaymentProcessor {
  pay(amount: number): void {
    console.log(`Processing $${amount} with Stripe`);
  }
}

class RazorpayPayment implements PaymentProcessor {
  pay(amount: number): void {
    console.log(`Processing ₹${amount} with Razorpay`);
  }
}

class CheckoutService {
  constructor(private processor: PaymentProcessor) {}

  checkout(amount: number): void {
    this.processor.pay(amount);
  }
}

const checkout = new CheckoutService(new StripePayment());
checkout.checkout(1000);
```

`CheckoutService` depends on the `PaymentProcessor` interface, not on a
particular payment provider. You can replace `StripePayment` with
`RazorpayPayment` without changing the checkout service.

## Go (Golang)

Go uses interfaces too. A type satisfies an interface automatically when
it implements its methods.

```go
package main

import "fmt"

type PaymentProcessor interface {
    Pay(amount float64)
}

type StripePayment struct{}

func (s StripePayment) Pay(amount float64) {
    fmt.Printf("Processing $%.2f with Stripe\n", amount)
}

type RazorpayPayment struct{}

func (r RazorpayPayment) Pay(amount float64) {
    fmt.Printf("Processing ₹%.2f with Razorpay\n", amount)
}

type CheckoutService struct {
    processor PaymentProcessor
}

func (c CheckoutService) Checkout(amount float64) {
    c.processor.Pay(amount)
}

func main() {
    checkout := CheckoutService{
        processor: StripePayment{},
    }

    checkout.Checkout(1000)
}
```

To use Razorpay instead, pass `RazorpayPayment{}` to `CheckoutService`;
the service itself stays unchanged.

## Key idea

Abstraction separates **what a component promises to do** (its
interface) from **how it does it** (its implementation). This reduces
coupling and makes implementations easier to replace or test.
