package lkdn

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

var POST_SELECTOR = `document.querySelector('h2#ember383').innerText`

func Scrap(url string) {

	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*40)
	defer cancel()
	var text string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady(`//h2[@id="ember383"]`),
		chromedp.Evaluate(POST_SELECTOR, &text),
	)
	fmt.Println(text, err)

}
