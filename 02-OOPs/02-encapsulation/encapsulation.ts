// Account example

class BankAccount {
  private balance: number = 10000;

  deposit(amount: number): void {
    if (amount <= 0) {
      throw new Error("Invalid deposit amount");
    }

    this.balance += amount;
  }

  withdraw(amount: number): void {
    if (amount <= 0) {
      throw new Error("Invalid withdraw amount");
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

const myAccount: BankAccount = new BankAccount();

console.log(myAccount.getBalance()); // 10,000

myAccount.deposit(100);

console.log(myAccount.getBalance()); // 10,100

myAccount.withdraw(5000);

console.log(myAccount.getBalance()); // 5,100

// This is illegal
// myAccount.balance = 1000;
