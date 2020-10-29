package blockchain

import (
	"DataCertProjest/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)
//挖比特币难度：两周
const DIFFFICULTY  = 16
/**
工作量证明结构体
**/
type ProofOfWork struct {
	//目标值
	Target *big.Int
	//工作量证明算法对应的那个区块
	Block Block
}

func NewPoW(block Block)ProofOfWork{
	target :=big.NewInt(1)
	target.Lsh(target,255-DIFFFICULTY)
	pow :=ProofOfWork{
		Target: target,
		Block:  block,
	}
	return pow
}

func (p ProofOfWork)Run()([]byte ,int64)  {
	var nonce int64
	bigBlock :=new(big.Int)
	var block256Hash []byte
	for{
		block :=p.Block

		heightBytes, _ := util.IntToBytes(block.Height)
		timeBytes, _ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes, _ := util.IntToBytes(nonce)

		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		}, []byte{})
		sha256Hash :=sha256.New()
		sha256Hash.Write(blockBytes)

		block256Hash =sha256Hash.Sum(nil)

		bigBlock =bigBlock.SetBytes(block256Hash)
		if p.Target.Cmp(bigBlock) ==1{
			break
		}
		nonce++
	}
	return block256Hash,nonce
}