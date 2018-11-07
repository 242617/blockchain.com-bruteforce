package main

import (
	"context"
	"flag"
	"log"
	"regexp/syntax"
	"time"

	"github.com/alixaxel/genex"
	"github.com/chromedp/chromedp"
)

var (
	list     bool
	username string
	password string
	resume   string
)

var current string

func main() {

	flag.BoolVar(&list, "list", false, "List mode")
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password mask")
	flag.StringVar(&resume, "resume", "", "Resume from")
	flag.Parse()
	if username == "" && !list {
		log.Fatal("username is empty")
	}
	if password == "" {
		log.Fatal("password is empty")
	}

	// file, err := os.OpenFile("bruteforce.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetFlags(0)
	// log.SetOutput(file)
	log.Println("start")
	start := time.Now()
	defer func() {
		log.Printf("done in %s", time.Since(start).String())
	}()

	charset, err := syntax.Parse(`[0-9a-zA-Z]`, syntax.Perl)
	die(err)

	input, err := syntax.Parse(password, syntax.Perl)
	die(err)

	attempts, total := 0, genex.Count(input, charset, 3)

	if list {
		log.Println("list combinations")
		for word := range words(resume, input, charset) {
			log.Println(word)
		}
		if resume == "" {
			log.Printf("total: %d\n", int(total))
		}
		return
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			if resume == "" {
				log.Printf("[%s], attempts: %d/%d, current: %s\n", time.Since(start).Round(time.Second).String(), attempts, int(total), current)
			} else {
				log.Printf("[%s], attempts: %d, current: %s\n", time.Since(start).Round(time.Second).String(), attempts, current)
			}
		}
	}()

	var cancel func()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := chromedp.New(ctx) // , chromedp.WithLog(log.Printf)
	die(err)

	err = client.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://login.blockchain.com/#/login"),
		chromedp.SendKeys(LoginSelector, username, chromedp.ByQuery),
	})
	die(err)

	resCh, nothingCh := make(chan string), make(chan struct{})
	go func() {
		for current = range words(resume, input, charset) {
			tasks := try(current, resCh)
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

}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
