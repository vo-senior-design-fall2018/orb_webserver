package main

import (
	"encoding/json"
	"github.com/vo-senior-design-fall2018/orb_webserver/mono"
	"github.com/vo-senior-design-fall2018/orb_webserver/rgbd"
	"io/ioutil"
	"log"
)

// JSON Configuration file
type Configuration struct {
	Server string `json:"server"`
}

func main() {
	file, err := ioutil.ReadFile("conf.json")

	if err != nil {
		log.Fatal(err)
	}

	configuration := Configuration{}
	err = json.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(data.Server) // output: [UserA, UserB]
	log.Printf("Type server: %s\n", configuration.Server)

	if configuration.Server == "mono" {
		mono.New("8080")
	} else if configuration.Server == "rgbd" {
		rgbd.New("8080")
	}
}
