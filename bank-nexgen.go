package main

import (
	"fmt"

	fileops "example.com/nexgen-bank/file-ops"
	"github.com/Pallinder/go-randomdata"
)

const BalanceAccountFile = "balance-saldo.txt"

func main() {

	//SALDO USER
	var BalanceAccount, err = fileops.GetFloatFromFile(BalanceAccountFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
	}
	fmt.Println("GoNex - BANK")
	fmt.Println("Reach Us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice)
		//fmt.Print("Your Select : ")
		switch choice {
		case 1:
			fmt.Println("Your Balance is", BalanceAccount)
		case 2:
			fmt.Print("Please Insert Your Money for Deposit : ")
			var AmountDeposit float64
			fmt.Scan(&AmountDeposit)
			if AmountDeposit <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0")
				continue
			}
			//Gave info for Update Balance(Saldo)
			BalanceAccount += AmountDeposit //BalanceAccount = AccountBalance + Deposit
			fmt.Println("Balance Updated ! New Amount", BalanceAccount)
			fileops.WriteFloatIntoFile(BalanceAccount, BalanceAccountFile)
		case 3:
			fmt.Print("Withdrawal Amount : ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount. Please be greater than 0")
				continue
			}
			if withdrawalAmount > BalanceAccount {
				fmt.Println("Invalid Amount. You can't witdraw more than you have")
				continue
			}
			//Gave info for Update Balance(Saldo)
			BalanceAccount -= withdrawalAmount //BalanceAccount = AccountBalance + Deposit
			fmt.Println("Balance Updated ! Curently Amount : ", BalanceAccount)
			fileops.WriteFloatIntoFile(BalanceAccount, BalanceAccountFile)
		default:
			fmt.Println("See u Later")
			fmt.Println("Thanks For Visit Our Bank")
			return
			//break
		}
		//fmt.Println("Selected : ", choice)
	}
}
