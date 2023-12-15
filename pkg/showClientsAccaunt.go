package pkg

import "fmt"

func ShowClientsAccaunt(){
	var name string
	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)
	
	balance, ok := Database[name]
	if !ok{
		fmt.Println("Ошибки!, данного пользователя нет в нашей бд")
		return
	}
	fmt.Printf("баланс счета %v  является %v \n",name,balance)
}