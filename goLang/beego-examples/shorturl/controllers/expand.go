/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:27
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 10:12:37
 * @FilePath: \code\goLang\beego-examples\shorturl\controllers\expand.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type ExpandController struct {
	beego.Controller
}

func (this *ExpandController) Get() {
	var result ShortResult
	shorturl := this.Input().Get("shorturl")
	result.UrlShort = shorturl
	if urlcache.IsExist(shorturl) {
		result.UrlLong = urlcache.Get(shorturl).(string)
	} else {
		result.UrlLong = ""
	}
	this.Data["json"] = result
	this.Data["num"] = 12
	this.ServeJSON()
}
