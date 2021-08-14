package app

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/lifegit/go-gulu/v2/pkg/spider/chromedpm"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

var jar, _ = cookiejar.New(nil)

func GetColly() *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL)

		request.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		request.Headers.Set("Accept-Encoding", "gzip, deflate")
		request.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		request.Headers.Set("Cache-Control", "no-cache")
		request.Headers.Set("Host", "www.stats.gov.cn")
		request.Headers.Set("Pragma", "no-cache")
		request.Headers.Set("Proxy-Connection", "keep-alive")
		request.Headers.Set("Upgrade-Insecure-Requests", "1")
		request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
		// cookies
		//request.Headers.Set("Cookie","wzwscid=4c72d234f642682e30318448f21cac99818d1863911a6158f1e47fb2f38b70b898771babdbf5fa2ce9dce6f938d91ccb373bd4bf9f17c658faee0e9d5d3af0f614a6ce1004479db2fe2e90128387b190")
		if cookies := jar.Cookies(request.URL); cookies != nil {
			var v string
			for _, cookie := range cookies {
				if cookie.Name != "" {
					v += fmt.Sprintf("%s=%s;", cookie.Name, cookie.Value)
				}
			}
			request.Headers.Set("Cookie", v)
		}
	})

	return c
}

// GbkToUtf8 GBK 转 UTF-8
func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, _ := ioutil.ReadAll(reader)

	return d
}

// Utf8ToGbk UTF-8 转 GBK
func Utf8ToGbk(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, _ := ioutil.ReadAll(reader)

	return d
}

func GetCookies() (cookies []*network.Cookie, err error) {
	u := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/index.html"
	ctx2, cancel := chromedpm.NewAweChromeDp(10, false)
	defer cancel()

	err = chromedp.Run(*ctx2,
		chromedp.Navigate(u),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			err = chromedp.WaitReady("html").Do(ctx)
			cookies, err = network.GetAllCookies().Do(ctx)
			return err
		}),
	)

	cookiesMerge(u, cookies)
	return
}

func GetVerification() (cookies []*network.Cookie, err error) {
	u := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/index.html"
	ctx2, cancel := chromedpm.NewAweChromeDp(100, true)
	defer cancel()

	err = chromedp.Run(*ctx2,
		chromedp.Navigate(u),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			fmt.Println("等待输入验证码(及按确定键)...")
			err = chromedp.WaitReady(".search_bg").Do(ctx)
			fmt.Println("验证码验证完成！")
			time.Sleep(time.Second * 3)
			cookies, err = network.GetAllCookies().Do(ctx)
			return err
		}),
	)

	cookiesMerge(u, cookies)
	return
}

func cookiesMerge(URl string, cookies []*network.Cookie) {
	for _, cookie := range cookies {
		u, err := url.Parse(URl)
		if err == nil {
			jar.SetCookies(u, []*http.Cookie{{
				Name:   cookie.Name,
				Value:  cookie.Value,
				Domain: cookie.Domain,
				Path:   cookie.Path,
			}})
		}
	}
}

var ErrorFrequently = errors.New("请开启JavaScript并刷新该页")
var ErrorVerification = errors.New("验证码访问")

func IsNeedVerify(body string) (err error) {
	if strings.Index(body, "jsjiami.com.v6") != -1 {
		return ErrorFrequently
	}

	if strings.Index(body, "javascript:changeImg()") != -1 {
		return ErrorVerification
	}
	return nil
}

func SolveVerify(err error) {
	switch err {
	case ErrorFrequently:
		GetCookies()
	case ErrorVerification:
		GetVerification()
	}
}
