package main

import (
	"fmt"
	"github.com/booleanalgebra/Modbus"
	"time"

)

func main() {

	fmt.Println("Enter Server IP address:")
	var Modbus_IP string
	fmt.Scan(&Modbus_IP)
	fmt.Println(modbusclient.ConnectTCP(*Modbus_IP, modbusclient.MODBUS_PORT))
	fmt.Println("Read or Write? (type R or W and Enter)")
	var R_W string
	fmt.Scan(&R_W)
	fmt.Println("Enter address to", *R_W)
	var Modbus_Address int

	if R_W == "R" {modbusclient.TCPRead()

	}

}