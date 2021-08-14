package main_test

import (
	"github.com/stretchr/testify/assert"
	"go-china-division/app"
	"testing"
)

func TestMa(t *testing.T) {
	cd := app.ChinaDivision{
		OutPath:  "./",
		FileName: "china.json",
		Option:   app.OptionProvinceCityCountyTownVillage,
	}
	err := app.Run(cd)
	assert.NoError(t, err)
}
