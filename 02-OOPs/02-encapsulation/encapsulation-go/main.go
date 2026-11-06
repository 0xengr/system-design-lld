package main

import (
	"fmt"

	"github.com/0xengr/system-design-lld/02-oops/02-encapsulation/encapsulation-go/bank"
)

func main() {
	myAccount := bank.NewBankAccount()

	fmt.Println(myAccount.Balance()) // 10000

	if err := myAccount.Deposit(2000); err != nil {
		fmt.Println(err)
	}

	fmt.Println(myAccount.Balance()) // 12000

	if err := myAccount.Withdraw(3000); err != nil {
		fmt.Println(err)
	}

	fmt.Println(myAccount.Balance()) // 9000

	fmt.Println("MY TOTAL BALANCE NOW: ", myAccount.Balance())
}
