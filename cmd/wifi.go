package main

import (
	"fmt"
	"log"
	"time"

	"github.com/a-clap/distillation/pkg/wifi"
)

func main() {

	w, err := wifi.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(w.APs())

	fmt.Println(w.Disconnect())

	time.Sleep(1 * time.Second)

	er := w.Connect(wifi.Network{
		AP: wifi.AP{
			ID:   0,
			SSID: "MI 8",
		},
		Password: "adas1234",
	})
	fmt.Println(er)

}
