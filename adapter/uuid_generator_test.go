package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Generateメソッドは新規IDを生成できる(t *testing.T) {
	idGen := IdGenerator{}
	id := idGen.Generate()

	// 時間がちょっとでもズレれば違うIDに
	assert.NotEqual(t, idGen.Generate(), id)
}
