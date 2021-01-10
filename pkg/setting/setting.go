package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var cfg *ini.File
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}

type Database struct {
	User  		string
	Password  	string
	Port  		int
	URL   		string
	Type		string
	Dbname		string
}

type Server struct {
	MODE 			string
	HttpPort 		int
	ReadTimeout 	time.Duration
	WriteTimeout 	time.Duration
}


func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		panic("Not found config.ini !!!")
	}
	mapTo("default", ServerSetting)
	mapTo("database", DatabaseSetting)

}
func mapTo (section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section,err)
	}
}




