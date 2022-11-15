/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:28
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-15 09:41:27
 * @FilePath: \code\goLang\blockchain_go_videos-master\part54-wallets\BLC\Wallets.go
 */
package BLC

import "fmt"

//key is wallet address
type Wallets struct {
	Wallets map[string]*Wallet
}

// 创建钱包集合
func NewWallets() *Wallets {

	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	return wallets
}

// 创建一个新钱包
func (w *Wallets) CreateNewWallet() {

	wallet := NewWallet()
	fmt.Printf("Address：%s\n", wallet.GetAddress())
	w.Wallets[string(wallet.GetAddress())] = wallet
}
