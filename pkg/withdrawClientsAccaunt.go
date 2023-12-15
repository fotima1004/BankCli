package pkg

import "fmt"

func WhithdrawClientsAccaunt(){
	var name string
	var amount float64
	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)
	balance, ok := Database[name]
	if !ok{
		fmt.Println("Ошибки!, данного пользователя нет в нашей базе")
	}
	fmt.Println("Введите сумму которую хотите пополнить")
	fmt.Scan(&amount)
	if balance < amount{
		fmt.Println("ошибка! недостаточно средств на балансе")
		return
	}
	Database[name] = balance - amount
}