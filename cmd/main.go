package main

import (
	"log"

	"github.com/url-shortner/pkg/config"
	"github.com/url-shortner/pkg/dependency"
)




func main(){
	config,err:=config.LoadConfig()
	if err!=nil{
		log.Fatalln(err)
	}
	server,err:=dependency.InitializeAPI(config)
	if err!=nil{
		log.Fatalln("error in loading server",err)
	}else{
		server.Start()
	}

}