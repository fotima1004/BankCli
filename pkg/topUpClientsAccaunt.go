package pkg

import "fmt"

func TopUpClientsAccaunt(){
	var name string
	var amount float64
	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)
	balance, ok := Database[name]
	if !ok{
		fmt.Println("Ошибки!, данного пользователя нет в нашей базе")
		return
	}
	fmt.Println("Введите сумму которую хотите пополнить")
	fmt.Scan(&amount)

	Database[name] = balance + amount
}