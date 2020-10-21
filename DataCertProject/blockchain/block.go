package blockchain

import (
	"DataCertProject/util"
	"time"
)

type Block struct {
	Height    int    //区块高度
	TimeStamp int64  //时间戳
	Hash      []byte //区块的hash
	Data      []byte // 数据
	PrevHash  []byte //上一个区块的Hash
	Version   string //版本号

}
/**
* 新建一个区块实例，并返回该区块
*/

func NawBlock(height int ,data []byte ,prevHash []byte)(Block)  {
	block :=Block{
		Height:    height + 1,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}
	block.Hash = util.SHA256HashBlock(block)
	return block
}