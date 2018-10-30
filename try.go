package main

import (
	"context"
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func try(password string, resCh chan<- string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.SendKeys(passwordSelector, password, chromedp.ByQuery),
		chromedp.WaitEnabled(submitSelector, chromedp.ByQuery),
		chromedp.Click(submitSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context, executor cdp.Executor) error {
			okCh, failCh, errCh := make(chan struct{}), make(chan struct{}), make(chan error)
			go func() {
				var str string
				err := chromedp.Text(signoutSelector, &str, chromedp.ByQuery).Do(ctx, executor)
				if err != nil {
					errCh <- err
				}
				okCh <- struct{}{}
			}()
			go func() {
				var str string
				err := chromedp.Text(noteSelector, &str, chromedp.ByQuery).Do(ctx, executor)
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
