package baidufanyi_test

import (
	"fmt"
	"os"

	"github.com/chyroc/baidufanyi"
)

func Example_baidufanyi() {
	cli := baidufanyi.New(baidufanyi.WithCredential(os.Getenv("APP_ID"), os.Getenv("APP_SECRET")))

	res, err := cli.Translate("hi", baidufanyi.LanguageEn, baidufanyi.LanguageZh)
	fmt.Println(err)
	fmt.Printf("%#v\n", res) // output: 你好
}
