package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	//c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))

	c, err := chromedp.New(ctxt, chromedp.WithRunnerOptions(runner.UserDataDir("./user")))
	if err != nil {
		log.Fatal(err)
	}

	//c.Run(ctxt, util.Login("", ""))

	getPosts(c)
	//c.Run(ctxt, util.Like("https://www.facebook.com/permalink.php?story_fbid=1851001208272928&id=623521487687579"))

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

func getPosts(c *chromedp.CDP) {
	ctxt, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	var posts []*cdp.Node

	c.Run(ctxt, chromedp.Navigate(`https://www.facebook.com/A2larm/`))
	c.Run(ctxt, chromedp.WaitVisible(`creation_hub_entrypoint`, chromedp.ByID))
	fmt.Println("VISIBLE!")

	c.Run(ctxt, chromedp.Sleep(10*time.Second))

	fmt.Println("TRYING TO CLICK NOW")
	err := c.Run(ctxt, chromedp.Click(`_5pcq`))
	fmt.Println(err)
	time.Sleep(1 * time.Minute)

	err = c.Run(ctxt, chromedp.Nodes(`fcg`, &posts, chromedp.BySearch))
	fmt.Println(err)

	fmt.Println(posts)
	time.Sleep(10 * time.Minute)
}
