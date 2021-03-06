# baidufanyi

go sdk for baidu fanyi (docs: https://fanyi-api.baidu.com/doc/21).

[![codecov](https://codecov.io/gh/chyroc/baidufanyi/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/baidufanyi)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/baidufanyi "go report card")](https://goreportcard.com/report/github.com/chyroc/baidufanyi)
[![test status](https://github.com/chyroc/baidufanyi/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/baidufanyi/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/baidufanyi)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Fbaidufanyi.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Fbaidufanyi)

![](./header.jpg)

## Install

```shell
go get github.com/chyroc/baidufanyi
```

## Usage

### Translate 翻译

```go
package main

import (
	"fmt"
	"os"

	"github.com/chyroc/baidufanyi"
)

func main() {
	cli := baidufanyi.New(baidufanyi.WithCredential(os.Getenv("BAIDUFANYI_APP_ID"), os.Getenv("BAIDUFANYI_APP_SECRET")))

	res, err := cli.Translate("hi", baidufanyi.LanguageEn, baidufanyi.LanguageZh)
	if err != nil {
		panic(err)
	}

	fmt.Println(res[0].Dst) // output: 你好
}
```