package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
	"y3cache/client"
)

const port = ":4000"

func main() {
	getF := flag.Bool(
		"get",
		false,
		"SET",
	)
	setF := flag.Bool(
		"set",
		false,
		"SET",
	)
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(10)
	if *setF {
		SendStuff(&wg)
		return
	}

	if *getF {
		GetStuff(&wg)
		return
	}
}

func SendStuff(wg *sync.WaitGroup) {
	c, err := client.New(port, client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	fmt.Println("conntected successfully..")
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var (
				key   = []byte(fmt.Sprintf("K_%d", i))
				value = []byte(fmt.Sprintf("V_%d", i))
			)

			fmt.Println("setting...")
			err = c.Set(context.Background(), key, value)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("set Sucessfully")
			wg.Done()
		}(wg)
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func GetStuff(wg *sync.WaitGroup) {
	c, err := client.New(port, client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	fmt.Println("conntected successfully..")
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			key := []byte(fmt.Sprintf("K_%d", i))
			val, err := c.Get(context.Background(), key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Get Sucessfully val:", string(val))
			wg.Done()
		}(wg)
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func genRndWord(min, max int) []byte {
	wordLength := rand.Intn(max-min+1) + min
	word := make([]byte, wordLength)
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	for i := range word {
		word[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return word
}

// func randomBytes(n int) []byte {
//
// 	buf := make([]byte, n)
// 	io.ReadFull(rand.Reader, buf)
// 	return buf
// }
