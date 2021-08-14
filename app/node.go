package app

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/lifegit/go-gulu/v2/nice/tagain"
	"strings"
	"time"
)

type Node struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Href     string `json:"-"`
	Children []Node `json:"children,omitempty"`
}

func GetDepthCh(nodes *[]Node, callNodes *[][]*Node, current ...Option) {
	c := OptionProvince
	if current != nil {
		c = current[0]
	}
	if len(*callNodes) <= int(c) {
		*callNodes = append(*callNodes, []*Node{})
	}
	for key, node := range *nodes {
		(*callNodes)[c] = append((*callNodes)[c], &(*nodes)[key])
		if len(node.Children) > 0 {
			GetDepthCh(&node.Children, callNodes, c+1)
		}
	}
	return
}

func GetNode(option Option, url string) (n []Node, err error) {
	c := GetColly()
	c.OnResponse(func(response *colly.Response) {
		response.Body = GbkToUtf8(response.Body)
		bodyStr := string(response.Body)
		fmt.Println(bodyStr)
		if e := IsNeedVerify(string(response.Body)); e != nil {
			err = e
		}
	})
	c.OnHTML("table td[width][height][valign] table tr", func(e *colly.HTMLElement) {
		if !strings.HasSuffix(e.DOM.AttrOr("class", ""), "tr") {
			return
		}
		var href string
		if option == OptionProvince {
			e.DOM.Find("a[href]").Each(func(i int, selection *goquery.Selection) {
				href = selection.AttrOr("href", "")
				n = append(n, Node{
					Label: selection.Text(),
					Value: fmt.Sprintf("%s0000", strings.ReplaceAll(href, ".html", "")),
					Href:  e.Request.AbsoluteURL(href),
				})
			})

		} else {
			res := e.DOM.Find("td").Map(func(i int, selection *goquery.Selection) string {
				return selection.Text()
			})
			if href = e.DOM.Find("td a[href]").First().AttrOr("href", ""); href != "" {
				href = e.Request.AbsoluteURL(href)
			}
			n = append(n, Node{
				Label: res[len(res)-1], // last one
				Value: res[0],
				Href:  href,
			})
		}
	})

	if e := c.Visit(url); e != nil {
		err = e
	}

	if err == nil && len(n) <= 0 {
		err = errors.New("Failed to get the node")
	}
	return n, err
}

func GetNodeAgain(option Option, url string) (n []Node, e error) {
	tagain.TAgain(func(i int) tagain.TryAgain {
		n, e = GetNode(option, url)
		if e != nil {
			SolveVerify(e)
			return tagain.TryAgainFailTally
		}
		return tagain.TryAgainSuccess
	}, 3, time.Second)

	return
}

func TectonicNodes(option Option) (division []Node, err error) {
	for i := OptionProvince; i <= option; i++ {
		fmt.Printf("tectonicNodes [%d/%d] running...\n", int(i+1), int(option+1))

		if i == OptionProvince {
			years, err := GetYearListAgain()
			if err != nil {
				return nil, err
			}
			division, err = GetNodeAgain(i, years[0].href)
			if err != nil {
				return division, err
			}
		} else {
			var callNodes [][]*Node
			GetDepthCh(&division, &callNodes)
			for _, node := range callNodes[i-1] {
				if node.Href != "" {
					res, err := GetNodeAgain(i, node.Href)
					if err != nil {
						return nil, err
					}
					node.Children = res
				}
			}
		}
	}

	return
}
