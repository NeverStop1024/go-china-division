package app

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/lifegit/go-gulu/v2/nice/tagain"
	"time"
)

type Year struct {
	year string
	href string
}

func GetYearList() (y []Year, err error) {
	c := GetColly()
	c.OnResponse(func(response *colly.Response) {
		bodyStr := string(response.Body)
		fmt.Println(bodyStr)
		if e := IsNeedVerify(string(response.Body)); e != nil {
			err = e
		}
	})
	c.OnHTML("ul.center_list_contlist", func(e *colly.HTMLElement) {
		e.DOM.Find("a").Each(func(i int, selection *goquery.Selection) {
			y = append(y, Year{
				year: selection.Text(),
				href: selection.AttrOr("href", ""),
			})
		})
	})

	if e := c.Visit("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/index.html"); e != nil {
		err = e
	}

	if err == nil && len(y) <= 0 {
		err = errors.New("Failed to get the latest year")
	}

	return y, err
}

func GetYearListAgain() (y []Year, e error) {
	tagain.TAgain(func(i int) tagain.TryAgain {
		y, e = GetYearList()
		if e != nil {
			SolveVerify(e)
			return tagain.TryAgainFailTally
		}
		return tagain.TryAgainSuccess
	}, 3, time.Second)

	return
}
