/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:28
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-15 09:23:58
 * @FilePath: \code\goLang\blockchain_go_videos-master\part51-wallet-address\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"p51/BLC"
)

func main() {

	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	fmt.Printf("address：%s\n", address)
	//	1EESFLjfwLqoSq4LJA8eymoPi5pc5Sq6VC
	//  1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

}
