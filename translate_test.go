package baidufanyi_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chyroc/baidufanyi"
	"github.com/stretchr/testify/assert"
)

func Test_Fanyi(t *testing.T) {
	as := assert.New(t)

	cli := baidufanyi.New(baidufanyi.WithCredential(os.Getenv("BAIDUFANYI_APP_ID"), os.Getenv("BAIDUFANYI_APP_SECRET")))

	for i := 0; i < 3; i++ {
		res, err := cli.Translate("hi", baidufanyi.LanguageEn, baidufanyi.LanguageZh)
		if err != nil {
			if strings.Contains(err.Error(), "Invalid Access Limit") && i < 2 {
				time.Sleep(time.Second)
				continue
			}
		}
		as.Nil(err)
		as.NotNil(res)
		as.Len(res, 1)
		as.Equal("你好", res[0].Dst)
	}
}
