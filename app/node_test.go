package app_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-china-division/app"
	"testing"
)

var Division = []app.Node{
	{
		Label: "山东省",
		Children: []app.Node{
			{
				Value: "济南市",
				Children: []app.Node{
					{Label: "历下区"},
					{Label: "历城区"},
				},
			},
			{
				Value: "青岛市",
				Children: []app.Node{
					{Label: "即墨区"},
					{Label: "崂山区"},
				},
			},
		},
	},
	{
		Label: "浙江省",
		Children: []app.Node{
			{
				Value: "杭州市",
				Children: []app.Node{
					{Label: "上城区"},
					{Label: "拱墅区"},
				},
			},
			{
				Value: "宁波市",
				Children: []app.Node{
					{Label: "海曙区"},
					{Label: "江北区"},
				},
			},
		},
	},
}

func TestGetDepthCh(t *testing.T) {
	var callNodes [][]*app.Node

	app.GetDepthCh(&Division, &callNodes)
	assert.Len(t, callNodes, 3)

	assert.Len(t, callNodes[0], 2)
	assert.Len(t, callNodes[1], 4)
	assert.Len(t, callNodes[2], 8)
}

func TestGetNode(t *testing.T) {
	type OptionHref struct {
		option app.Option
		href   string
	}
	list := []OptionHref{
		{option: app.OptionProvince, href: "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/index.html"},
		{option: app.OptionProvinceCity, href: "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/11.html"},
		{option: app.OptionProvinceCityCounty, href: "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/11/1101.html"},
		{option: app.OptionProvinceCityCountyTown, href: "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/11/01/110101.html"},
		{option: app.OptionProvinceCityCountyTownVillage, href: "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2020/11/01/01/110101001.html"},
	}

	for _, item := range list {
		node, err := app.GetNode(item.option, item.href)
		assert.NoError(t, err)
		assert.NotZero(t, node)
		fmt.Println(node)
	}
}
