package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GetChromeScreenshot(url string, quality int) {
	screenshotUrl := fmt.Sprintf("https://%s/", url)

	var buf []byte

	var ext string = "png"
	if quality < 100 {
		ext = "jpeg"
	}

	log.Printf("Capturando a tela do site %s\n", screenshotUrl)

	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1920, 1080))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer acancel()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	tasks := chromedp.Tasks{
		chromedp.Navigate(screenshotUrl),
		chromedp.Sleep(3 * time.Second),
		chromedp.CaptureScreenshot(&buf),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}
	fileName := fmt.Sprintf("%s-%d-standard.%s", strings.Replace(url, "/", "-", -1), time.Now().UTC().Unix(), ext)
	if err := os.WriteFile(fileName, buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("Captura armazenada em %s\n", fileName)
}
