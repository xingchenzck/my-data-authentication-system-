package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Height    int64    //区块高度
	TimeStamp int64  //时间戳
	Hash      []byte //区块的hash
	Data      []byte // 数据
	PrevHash  []byte //上一个区块的Hash
	Version   string //版本号
	Nonce    int64  //随机数，用于pow工作量证明算法计算

}
/**
生成创世区块，返回区块信息
*/

func CreateGenesisBlock() Block {
	block :=NawBlock(0,[]byte{},[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return block
}

/**
* 新建一个区块实例，并返回该区块
*/

func NawBlock(height int64 ,data []byte ,prevHash []byte)(Block)  {
	block :=Block{
		Height:    height ,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}
	pow :=NewPoW(block)
	blockHash ,nonce :=pow.Run()

	block.Nonce = nonce
	block.Hash =blockHash

	return block
}
//区块的序列化
func (bk Block)Serialize()([]byte ,error)  {
	buff := new(bytes.Buffer)
	err :=gob.NewEncoder(buff).Encode(bk)
	if err != nil{
		return nil,err
	}
	return  buff.Bytes() ,nil
}
//区块的反序列化
func DeSerialize(data []byte)(*Block ,error)  {
	var block Block
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err !=nil {
		return nil ,err
	}
	return &block,nil
}
