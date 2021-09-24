package baidufanyi_test

import (
	"fmt"
	"os"

	"github.com/chyroc/baidufanyi"
)

func Example_baidufanyi() {
	cli := baidufanyi.New(baidufanyi.WithCredential(os.Getenv("BAIDUFANYI_APP_ID"), os.Getenv("BAIDUFANYI_APP_SECRET")))

	res, err := cli.Translate("hi", baidufanyi.LanguageEn, baidufanyi.LanguageZh)
	if err != nil {
		panic(err)
	}

	fmt.Println(res[0].Dst) // output: 你好
}
