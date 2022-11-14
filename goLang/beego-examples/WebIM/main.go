/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:27
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 09:41:41
 * @FilePath: \code\goLang\beego-examples\WebIM\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	_ "WebIM/routers"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const (
	APP_VER = "0.1.1.0227"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
