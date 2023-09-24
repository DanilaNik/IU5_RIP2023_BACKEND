package main

import (
	"github.com/DanilaNik/IU5_RIP2023/pkg/app"
	"log"
)

//Склад комплектующих.
//Услуги - список комлектующих для хранения с размером для места
//Заявки - заявки на доставку и отгрузку комплектующих

func main() {
	log.Println("Application start")
	app.StartServer()
	//serverApi.StartServer1()
}
