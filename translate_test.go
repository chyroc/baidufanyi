package baidufanyi_test

import (
	"os"
	"testing"

	"github.com/chyroc/baidufanyi"
	"github.com/stretchr/testify/assert"
)

func Test_Fanyi(t *testing.T) {
	as := assert.New(t)

	cli := baidufanyi.New(baidufanyi.WithCredential(os.Getenv("BAIDUFANYI_APP_ID"), os.Getenv("BAIDUFANYI_APP_SECRET")))

	res, err := cli.Translate("hi", baidufanyi.LanguageEn, baidufanyi.LanguageZh)
	as.Nil(err)
	as.NotNil(res)
	as.Len(res, 1)
	as.Equal("你好", res[0].Dst)
}
