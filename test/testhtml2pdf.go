package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"os"
	"time"
)

func main() {
	t := time.Now()
	//pdfg, err := wk.NewPDFGenerator()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pdfg.AddPage(wk.NewPage("http://upload.i8956.com/temp/label/20230620/769/12614/invoice/892317033730837.html"))
	//err = pdfg.Create()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = pdfg.WriteFile("./baidu2.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}

	err := ChromedpPrintPdf("http://upload.i8956.com/temp/label/20230620/769/12614/invoice/892317033730837.html", "./baidu3.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed := time.Since(t)

	fmt.Println("app elapsed:", elapsed)
}
func ChromedpPrintPdf(url string, to string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().WithScale(0.9).
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}

	if err := os.WriteFile(to, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}
	return nil
}
