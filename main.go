package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Xopxe23/dota2-api/dota2"
)

func main() {
	client, err := dota2.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}
	heroes, err := client.GetHeroes()
	if err != nil {
		log.Fatal(err)
	}
	for _, hero := range heroes {
		fmt.Println(hero.Info())
	}
}
