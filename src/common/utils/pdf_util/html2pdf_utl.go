package pdf_util

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"os"
)

type Html2PdfUtil struct {
}

func (html2PdfUtil *Html2PdfUtil) ChromedpPrintPdf(url string, to string) error {
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
