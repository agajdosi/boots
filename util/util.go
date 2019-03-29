package util

import (
	"time"

	"github.com/chromedp/chromedp"
)

func Login(login, password string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.facebook.com`),
		chromedp.WaitVisible(`#email`, chromedp.ByID),
		chromedp.SendKeys(`#email`, login, chromedp.ByID),
		chromedp.SendKeys(`#pass`, password+"\n", chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
	}
}

func GoHome() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`https://www.facebook.com`),
		chromedp.WaitVisible(`#creation_hub_entrypoint`, chromedp.ByID),
	}
}

func Like(postLink string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(postLink),
		chromedp.WaitVisible(`span[class="_2md"]`, chromedp.ByQueryAll),
		chromedp.Click(`UFILikeLink _4x9- _4x9_ _48-k`, chromedp.BySearch),
	}
}
