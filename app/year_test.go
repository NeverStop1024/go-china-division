package app_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-china-division/app"
	"testing"
)

func TestName(t *testing.T) {
	y, e := app.GetYearList()
	assert.NoError(t, e)
	assert.NotZero(t, y)

	fmt.Println(y)
}
