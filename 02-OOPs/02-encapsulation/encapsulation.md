# What is Encapsulation?

**Encapsulation is the process of wrapping data (properties) and methods (functions) together into a single unit (class), while restricting direct access to the internal data.**

In simple words:

> Encapsulation means protecting an object's data and allowing access to it only through controlled methods.

Think about your bank account.

You have:

- Account balance.
- Deposit operation.
- Withdraw operation.
- Balance checking operation.

You shouldn't be able to directly change your bank balance whenever you want.

For example, this should not be allowed:

```typescript
account.balance = -100000;
```

Instead, you should interact with the account through controlled methods.

Imagine a bank account with a balance of $1,000.

You can:

- Deposit money.
- Withdraw money.
- Check your balance.

But you cannot directly modify the balance.

The bank account controls how its internal data changes.

## Without Encapsulation:

Let's first understand what happens without encapsulation.

```typescript
class BankAccount {
  balance: number = 1000;
}

const account = new BankAccount();

// Directly modifying internal data
account.balance = -5000;

console.log(account.balance); // -5000
```

What's the problem?

Anyone can modify the balance directly.

This can lead to:

- Invalid account balances.
- Unexpected behavior.
- Security and data integrity problems.
- Difficulties maintaining business rules

## With Encapsulation

```typescript
class BankAccount {
  private balance: number = 1000;

  deposit(amount: number): void {
    if (amount <= 0) {
      throw new Error("Invalid deposit amount");
    }

    this.balance += amount;
  }

  withdraw(amount: number): void {
    if (amount <= 0) {
      throw new Error("Invalid withdrawal amount");
    }

    if (amount > this.balance) {
      throw new Error("Insufficient balance");
    }

    this.balance -= amount;
  }

  getBalance(): number {
    return this.balance;
  }
}

const account = new BankAccount();

// Deposit money
account.deposit(500);

// Withdraw money
account.withdraw(200);

// Check balance
console.log(account.getBalance()); // 1300

// Not allowed
// account.balance = -5000;
```
