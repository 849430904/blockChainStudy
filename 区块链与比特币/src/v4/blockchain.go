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

func NewBlockChain(address string) *BlockChain {


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
			coinbase := NewCoinbaseTx(address,"这是一个创世区块")
			genersis := NewGenersisBlock(coinbase)

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

func (bc *BlockChain) AddBlock(txs []*Transaction) {


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
	block := NewBlock(txs,prevBlockHash)


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


//创建blockChain数据库文件
func InitBlockChain(address string) *BlockChain {

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

		coinbase := NewCoinbaseTx(address,"备注信息....")

		//没有bucket，创建一个bucket(桶)，创建一个创世块，将数据填写到数据库的bucket
		genersis := NewGenersisBlock(coinbase)

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

//返回指定地址能够支配的utxo的交易金额
func (bc *BlockChain)FindUTXOTransactions(address string)[]Transaction  {

	//包含目标UTXO的交易集合
	var UTXOTransactions []Transaction

	//定义一个存储使用过的utxo集合  map[交易id] =  int64
	//0x1111111 ： 0,1 都是给Alice的转账
	spentUTXO := make(map[string][]int64)



	it := bc.NewIterator()
	//遍历区块
	for  {
		block := it.Next()

		//遍历所有交易
		for _,tx := range  block.Transactions {



			//遍历input
			//目的：找到已经消耗过的utxo,把它们放到一个集合里面
			//需要两个参数字段来标识使用过的Utxo:交易ID、output的索引

			if !tx.IsCoinbase() {//不为一个挖矿交易
				for _,input := range tx.TXInputs {

					if input.CanUnlockUTXOWith(address) {
						//map[txid][]int64
						spentUTXO[string(tx.TXID)] = append(spentUTXO[string(tx.TXID)],input.Vout)
					}

				}
			}


		OUTPUTS:
			//遍历Outputs
			//目的：找到所有能支配的utxo
			for currIndex,output := range tx.TXOutputs {

				//检查录前的output是否已经被消耗，如果消耗过了，就进行下一个output检验

				if spentUTXO[string(tx.TXID)] != nil {//非空，代表当前交易里面有消耗的UTXO
					indexes := spentUTXO[string(tx.TXID)]
					for _,index := range indexes {
						if int64(currIndex) == index {//当前的索引和消耗的索引相同，说明这个output已经被消耗过了,直接跳过，进行下一个判断
							continue OUTPUTS
						}
					}
				}

				//如果当前地址是这个utxo的所有者，就满足条件
				if output.CanBeUnlockUTXOWith(address) {
					UTXOTransactions = append(UTXOTransactions,*tx)
				}
			}

		}




		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return UTXOTransactions
}

//寻找指定地址能够使用的utxo
func (bc *BlockChain)FindUTXO(address string) []*TXOutput {

	var UTXOs []*TXOutput

	txs := bc.FindUTXOTransactions(address)//找到这个地址的所有交易
	//遍历交易
	for _,tx := range txs {
		
		//遍历outout
		for _,utxo :=  range  tx.TXOutputs {

			//当前地址拥有的utxo
			if utxo.CanBeUnlockUTXOWith(address) {
				UTXOs = append(UTXOs,&utxo)
			}
		}
	}

	return UTXOs
}

