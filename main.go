package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}
	PORT := arguments[1]
	fmt.Println(PORT)
	client := redis.NewClient(&redis.Options{
		Addr:     PORT,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	flag := true
	for flag {
		fmt.Println("Enter your command")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		input := strings.Split(strings.Trim(text, "\n"), " ")
		cmd := input[0]
		switch cmd {
		case "DUMP":
			key := input[1]
			val, err := client.Get(key).Result()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(val)
		case "SET":
			key, value := input[1], input[2]
			err = client.Set(key, value, 0).Err()
			if err != nil {
				fmt.Println(err)
			}
		case "RENAME":
			key1, key2 := input[1], input[2]
			err = client.Rename(key1, key2).Err()
			if err != nil {
				fmt.Println(err)
			}
		case "EXIT":
			fmt.Println("EXIT...")
			flag = false
		}
	}

}
