package lkdn

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var POST_SELECTOR = `html`

func Scrap(url string) {
	var nodes []*cdp.Node
	ctx, _ := chromedp.NewContext(context.Background())
	ctx, cancel := context.WithTimeout(ctx, time.Second*40)
	defer cancel()
	var text string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady(`html`),
		chromedp.Nodes(POST_SELECTOR, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {

			return nil
		}),
	)
	fmt.Println(text, err)

}

func recurseNodes(nodes []*cdp.Node) {
	for _, n := range nodes {
		fmt.Println(n)
		if len(n.Children) > 1 {
			recurseNodes(n.Children)
		}

	}
}
