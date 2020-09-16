package comment_test

import (
	"fmt"

	"github.com/ylck/feishu"
	"github.com/ylck/feishu/apis/capabilities/document/comment"
)

func ExampleAddWhole() {
	var ctx *feishu.App

	payload := []byte("{}")
	accessToken := ""
	resp, err := comment.AddWhole(ctx, payload, accessToken)

	fmt.Println(resp, err)
}
