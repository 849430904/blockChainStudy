
编译与运行：
//rm block ，重新编译之前，如果已经编译过了
1，go build *.go
2，./block


v1版本授课思路：

    区块相关：
        1，定义一个区块的结构block
        a,区块头:6个字段
        b,区块体：字符串表示data
        2,提供一个创建区块的方法
        NewBlock（参数）

    区块链相关：
    1，定义一个区块链结构BlockChain
       Block数组
    2，提供一个创建BlockChain的方法
        NewBlockChain()
    3,提供一个添加区块的方法
        AddBlock(参数)



v2 版本思路：

     1.定义一个工作量证明的结构ProofOfWork
       block
       目标值

     2，提供一个创建Pow的方法
        NewProofWork(参数)

     3，计算，提供一个计算hash值的方法
        Run()

     4,提供一个校验函数,判断计算出来的hash对不对
        IsValid


V3 版本思路：

    目标：
        blot数据库（轻量级），  key -> value键值对存储在本地
        把区块链持久化到本地
        实现命令添加区块，打印区块链

    1，blot数据库的介绍
       key -> value进行读取，存储
       轻量级的
       开源的
    2，NewBlockChain函数的重写
       由数组编程数据库操作
       创建数据库文件
    3，AddBlock函数重写
       对数据的读取与写入
    4，对数据库的遍历
       迭代器的编写，Iterator
    5,命令行介绍及编写
       a.添加区块命令
       b.打印区块链命令

    //代码运行步骤：
    1，go build *.go //编译
    2, ./block //运行,运行后在项目根目录会有一个数据库文件(.db文件)
    3, 查看数据库（blockChain.db）命令：ls -al blockChain.db
    4，man strings ,查看可以打印的信息   //strings blockChain.db,打印二进制中可以打印的字符串




v4
添加交易utxo，utxo创建，转移等复杂功能


