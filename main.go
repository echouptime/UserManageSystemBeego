package main

import (
	"UserManagementSystem/routers"
	"github.com/astaxie/beego"
)

func main() {
	//defer models.InitDB().Close()
	Addr := ":9001"

	routers.Register()
	beego.Run(Addr)
	//if err := http.ListenAndServe(Addr, nil); err != nil {
	//	fmt.Println(err)
	//}

}
