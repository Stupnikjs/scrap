package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println("scrapping started")
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var text string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate("https://www.upwork.com/freelancers/~01d83b3bd3e0a74a3c"),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.InnerHTML("html", &text),
		chromedp.Evaluate(`document.querySelector("html")`, &text),
	)
	fmt.Println(text, err)

}
