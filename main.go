package main

import (
	"fmt"
	"log"
	"net/http"
	"opensvc/models"
	"opensvc/pkg/setting"
	"opensvc/routers"
)

func init()  {
	setting.Setup()
	models.Setup()
}

func main(){
	routerInit := routers.InitRouters()
	endpoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := &http.Server{
		Addr: endpoint,
		Handler: routerInit,
	}
	log.Println("[info] start http server listen ", endpoint)
	server.ListenAndServe()




	
}


