/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-15 10:15:25
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-15 10:39:04
 * @FilePath: \code\goLang\blockchain_go_videos-master\zversion\BLC\Block\ProofOfWork.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// 挖矿的难度值
const targetBit = 24

type ProofOfWork struct {
	Block  *Block   // 当前要验证的区块
	target *big.Int // 大数据存储
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransactions(),
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)

	return data
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {

	nonce := 0 //计数器 每算一次计数器加一 计数器越大说明计算难度越大

	var hashInt big.Int // 存储我们新生成的hash
	var hash [32]byte
	var maxNonce = math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n", proofOfWork.Block.Data)
	for nonce < maxNonce {
		dataBytes := proofOfWork.prepareData(nonce)

		hash = sha256.Sum256(dataBytes) //对block进行哈希运算得到hash

		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:]) //将哈希转换成一个大整数

		//将这个大整数与目标进行比较
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}

	return hash[:], int64(nonce)
}

// ProofOfWork构造函数
// 其实就是设置一个target 计算的哈希必须比target要小
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//向左移动256-24=232等于2的二百三十二次方
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{block, target}
}
