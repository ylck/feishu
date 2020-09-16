package platform_test

import (
	"fmt"

	"github.com/ylck/feishu"
	"github.com/ylck/feishu/apis/capabilities/document/platform"
)

func ExampleDocsMeta() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := platform.DocsMeta(ctx, payload, accessToken)

	fmt.Println(resp, err)
}

func ExampleSearchObject() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := platform.SearchObject(ctx, payload, accessToken)

	fmt.Println(resp, err)
}
