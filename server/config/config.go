package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

// Config is the basic handle of configurations file
type Config struct {
	Port        string `json:"port"`
	IP          string
	StartingLat string `json:"starting_lat"`
	StartingLng string `json:"starting_lng"`
}

var Cfg Config

// ReadConfigs do what it say :D
func ReadConfigs() {
	// Reading configs from file
	file, err := ioutil.ReadFile("config.cfg")

	if err != nil {
		fmt.Println("Error Reading the config file\n" + err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(file, &Cfg)

	if err != nil {
		fmt.Println("Error decoding the config file\n" + err.Error())
		os.Exit(1)
	}

	Cfg.IP = getIP()

}

func getIP() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}

	return "NO_IP"
}
