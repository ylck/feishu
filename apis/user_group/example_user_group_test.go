// Copyright 2020 ylck
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package user_group_test

import (
	"fmt"
	"net/url"

	"github.com/ylck/feishu"
	"github.com/ylck/feishu/apis/user_group"
)

func ExampleGroupList() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user_group.GroupList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleMembers() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user_group.Members(ctx, params)

	fmt.Println(resp, err)
}

func ExampleSearch() {
	var ctx *feishu.App

	params := url.Values{}
	resp, err := user_group.Search(ctx, params)

	fmt.Println(resp, err)
}
