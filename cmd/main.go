package main

import (
	"bankCli/pkg"
	"fmt"
)
func main(){
	for{
		var choice int
		fmt.Println("Создание клиента, и его счет")
		fmt.Println("Пополнить счет клиента")
		fmt.Println("Посмотреть баланс клиента")
		fmt.Println("Снять деньги с баланса")
		fmt.Println("Перевод денег")
		fmt.Println("Получить прибыль банка")
		fmt.Println("Выйти")
		fmt.Scan(&choice)

		switch choice{
		case 1:
			pkg.RegisterClient()
		case 2:
			pkg.TopUpClientsAccaunt()
		case 3:
			pkg.ShowClientsAccaunt()
		case 4:
			pkg.WhithdrawClientsAccaunt()
		case 5: 
			pkg.TransferToAnotherAccaunt()
		case 6:
			pkg.ShowProfit()
		case 7:
			return
		}

	}
}
