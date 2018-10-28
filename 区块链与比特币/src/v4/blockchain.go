package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

const  dbFile  = "blockChain.db"
const  dbBlockBucket  = "dbBlockBucket"
const  lastHashKey  = "lastHashKey"

//区块链相关
type BlockChain struct {
	//blocks []*Block

	db *bolt.DB//需要存储到数据库，数据库操作的的句柄
	tail []byte // 表示最后一个区块的哈希

}

func NewBlockChain() *BlockChain {


	//产生一个新的，需要创建一个文件数据库，在里面添加我们的区块
	//open文件
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {

	db , err := bolt.Open(dbFile,0600,nil)//0600 =读写
	CheckErr("NewBlockChain",err)


	var lastHash []byte

	//func (db *DB) Update(fn func(*Tx) error) error {
	db.Update(func(tx *bolt.Tx) error {

        //读取bucket(桶)
		bucket := tx.Bucket([]byte(dbBlockBucket))//强转成byte[]
		if bucket != nil {//存在bucket
		    //取出最后区块的hash值
			lastHash = bucket.Get([]byte(lastHashKey))//str强转成byte[]

		}else {
			//没有bucket，创建一个bucket(桶)，创建一个创世块，将数据填写到数据库的bucket
			genersis := NewGenersisBlock()

			bucket,err := tx.CreateBucket([]byte(lastHashKey))//创建一个bucket(桶)
			CheckErr("NewBlockChain2",err)

			bucket.Put(genersis.Hash,genersis.Seriallize())//将序列化的区块存储到bucket(桶)
			CheckErr("NewBlockChain3",err)

			bucket.Put([]byte(lastHashKey),genersis.Hash)//写入最后一个区块的hash
			CheckErr("NewBlockChain4",err)

			lastHash = genersis.Hash
		}

		return nil
	})

	return &BlockChain{db,lastHash}

}

func (bc *BlockChain) AddBlock(data string) {


	var prevBlockHash []byte

	//读取数据库的最后一个区块hash
	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dbBlockBucket))
		if bucket == nil{
			os.Exit(1)
		}

		prevBlockHash = bucket.Get([]byte(lastHashKey))
		return nil;
	})

	//1,先创建Block
	block := NewBlock(data,prevBlockHash)


	//2,写入Block
	err := bc.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(dbBlockBucket))//取出bucket(桶)
		if bucket == nil{
			os.Exit(1)
		}

		err := bucket.Put(block.Hash,block.Seriallize())//将序列化的区块存储到bucket(桶)
		CheckErr("AddBlock1",err)

		err = bucket.Put([]byte(lastHashKey),block.Hash)//写入最后一个区块的hash
		CheckErr("AddBlock2",err)

		bc.tail = block.Hash //更新本地内存的hash
		return nil
	})
	CheckErr("AddBlock3",err)
}


func InitBlockChain() *BlockChain {

	if isDBExist() {
		fmt.Println("blockChain exist already,no need to create")
		os.Exit(1)
	}

	//产生一个新的，需要创建一个文件数据库，在里面添加我们的区块
	//open文件
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {

	db , err := bolt.Open(dbFile,0600,nil)//0600 =读写
	CheckErr("InitBlockChain1",err)


	var lastHash []byte

	//func (db *DB) Update(fn func(*Tx) error) error {
	db.Update(func(tx *bolt.Tx) error {

		//没有bucket，创建一个bucket(桶)，创建一个创世块，将数据填写到数据库的bucket
		genersis := NewGenersisBlock()

		bucket,err := tx.CreateBucket([]byte(lastHashKey))//创建一个bucket(桶)
		CheckErr("InitBlockChain2",err)

		bucket.Put(genersis.Hash,genersis.Seriallize())//将序列化的区块存储到bucket(桶)
		CheckErr("InitBlockChain3",err)

		bucket.Put([]byte(lastHashKey),genersis.Hash)//写入最后一个区块的hash
		CheckErr("InitBlockChain4",err)

		lastHash = genersis.Hash

		return nil
	})

	return &BlockChain{db,lastHash}

}

func GetBlockChainHandler() *BlockChain{

	if !isDBExist() {
		fmt.Println("pls create blockChain first")
		os.Exit(1)
	}

	db , err := bolt.Open(dbFile,0600,nil)//0600 =读写
	CheckErr("GetBlockChainHandler",err)


	var lastHash []byte

	//func (db *DB) Update(fn func(*Tx) error) error {
	db.View(func(tx *bolt.Tx) error {

		//读取bucket(桶)
		bucket := tx.Bucket([]byte(dbBlockBucket))//强转成byte[]
		if bucket != nil {//存在bucket
			//取出最后区块的hash值
			lastHash = bucket.Get([]byte(lastHashKey))//str强转成byte[]

		}

		return nil
	})

	return &BlockChain{db,lastHash}
}

func isDBExist() bool  {
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	_,err := os.Stat(dbFile)
	if os.IsNotExist(err){
		return false
	}
	return true
}

/********************迭代器********************/

//迭代码，就是一个对象，它里面包含了一游标，游标是动的，一直向前或向后移动，完成整个容器的遍历
type BlockChainIterator struct {
	currHash   []byte
	db         *bolt.DB
}


//创建一个迭代器，同时初始化指向最后一个区块
func (bc *BlockChain)NewIterator() *BlockChainIterator  {
	return &BlockChainIterator{currHash:bc.tail,db:bc.db}
}

func (it *BlockChainIterator)Next() (block *Block) {

    err := it.db.View(func(tx *bolt.Tx) error {
    	bucket := tx.Bucket([]byte(dbBlockBucket))
    	if bucket == nil{
    		return nil
		}

    	//取区块
    	data := bucket.Get(it.currHash)
    	//反序列化
    	block := Deserialize(data)
    	//移动游标
    	it.currHash = block.PrevBlockHash
		return nil
	})
    CheckErr("Next",err)
    return
}



