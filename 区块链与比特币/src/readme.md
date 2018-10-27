
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


