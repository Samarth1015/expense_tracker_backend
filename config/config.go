package config

import "os"

type Configure struct{
	Port string
	Project string

}

func Config() Configure {
	port:= os.Getenv("PORT")
	project:=os.Getenv("PROJECT");

	return Configure{Port: port,Project:project };
	

}