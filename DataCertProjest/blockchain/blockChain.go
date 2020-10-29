package blockchain

import (
	"errors"
	"fmt"
	"github.com/bolt"
)

//桶的名称，该桶用于装区块信息
var BUCKET_NAME ="blocks"
//表示最新的区块的key名
var LAST_KEY = "lasthash"
//存储区块数据的文件
var CHAIANDB  = "chain.db"

/**
区块链结构体实例定义 ：用于表示代表一条区块链
该区块链包含以下功能:
    1.将新产生的区块与已有的区块链接起来，并保存
    2.可以查询某个区块的信息
    3.可以将所有区块进行遍历，输出区块信息
*/
type BlockChain struct {
	LashHash []byte//最新区块
	BoltDb *bolt.DB
}
/**
用于创建一条区块链，并返回区块链实例
  解释 ：用于区块链就是由一个一个的区块组成的,因此，如果要创建一条区块链，那么必须要先
        创建一个区块，该区块作为该条区块链的创世区块
*/
func NewBlockChain() BlockChain {
	//  打开存储区块数据的chain.db文件
	db ,err :=bolt.Open(CHAIANDB,0600,nil)
	if err !=nil {
		panic(err.Error())
	}
	var bl  BlockChain
	//先从区块链中都看是否创世区块已经存在
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket,err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 {//没有创世区块
			//创建创世区块
			genesis :=CreateGenesisBlock() //创世区块
			//创建一个存储区块数据的文件
			fmt.Printf("genesis的Hash值：%x\n",genesis.Hash)
			bl =BlockChain{
				LashHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes,_ :=genesis.Serialize()
			bucket.Put(genesis.Hash ,genesisBytes)
			bucket.Put([]byte(LAST_KEY),genesis.Hash)
		}else { //有创世区块
			lastHash :=bucket.Get([]byte(LAST_KEY))
			lastBlackBytes :=bucket.Get(lastHash)//创世区块的[]byte
			lastBlock ,err :=DeSerialize(lastBlackBytes)
			if err !=nil {
				panic("读取区块链数据失败")
			}
			bl =BlockChain{
				LashHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil

	})
	return bl
}
/**
调用BlockChain的SaveBlock方法，该方法可以将一个生成的新区块保存到chain.db文件中
*/
func (bc BlockChain)SeveDate(data []byte)(Block ,error)  {
	db :=bc.BoltDb
	var e error
	var lastBLock *Block
	//先查询chain.db中存储的最新的区块
	 db.View(func(tx *bolt.Tx) error {
		 bucket := tx.Bucket([]byte(BUCKET_NAME))
		 if bucket == nil {
			 e =errors.New("boltdb未创建 ，请重试！")
			 return e
		 }
		 lastBlockBytes := bucket.Get(bc.LashHash)
		 lastBLock ,_ =DeSerialize(lastBlockBytes)
		 return nil
	 })
    //1.
	newBlock := NawBlock(lastBLock.Height+1,data,lastBLock.Hash)

	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))

		newBlockBytes ,_ :=newBlock.Serialize()
		bucket.Put(newBlock.Hash ,newBlockBytes)
		bucket.Put([]byte(LAST_KEY) ,newBlock.Hash)
		bc.LashHash =newBlock.Hash
		return nil
	})
	return newBlock ,e
}














