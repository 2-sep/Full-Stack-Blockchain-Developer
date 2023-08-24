# 区块链基础

## 1.比特币

- ##### UTXO 模型

- ##### 交易结构

- ##### 区块结构

区块头：前一个区块头的**哈希值**、时间戳、难度目标、Nonce、Merkle 根

区块体：coinbase tx1、tx2、......、txn

![image-20230726233428035](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230726233428035.png)

- ##### 比特币转账的整个流程

  1.用户生成交易，向整个网络内的所有节点广播

  2.节点们收到一批这样的交易后会试图将他们打包成区块

  3.第一个打包成功的节点会将该区块向全网广播，其他节点收到区块后进行验证

  4.如果区块及其中的交易没有问题，则交易上链成功，节点又开启新一个区块的循环

挖矿：找到随机数（nonce），使得 Hash（nonce+区块头其他信息） < 难度系数

- ##### HD 钱包 BIP32 BIP44 BIP39

- ###### BIP32

根据一个随机数种子通过分层确定性推导的方式得到 n 个私钥，保存种子

- ###### BIP44

给 BIP32 的分层路径定义规范

m / purpose' / coin' / account' / change / address_index

- ###### BIP39

使用助记词的方式生成种子

- ##### 双花攻击

## 2.共识算法

- #### pow

​ 竞争找到随机数 nonce

- #### PoS

https://tomni.notion.site/tomni/Buidler-DAO-89472b07caff4a5b9807d1e54117181f Builder DAO

https://www.youtube.com/watch?v=lyEibZhcsCU

**PBFT**

​ 当收到 2/3 以上的同种票后，网络中的节点们达成了共识

​ merge:beacoin chain 和主网合并以切换到 PoS

**validator pool**

​ (1)提议: leader 怎么选择？

​ (2)投票: validator 造假怎么办？

​ (3)出结果

**VRF 抽奖**

**Gasper**机制解决 Finalization 问题

​ 设定 1：投票可以跨区块

​ 设定 2：违反相关规定的 validator 将被罚没所有的 staking 资产

- DPoS 委托权益证明机制

  eos

## 3.加密算法

- ### **哈希算法**

将任意长度的消息转换为一个固定长度的值

比特币:SHA 2-256 以太坊:Keccak256(未标准化的 SHA 3-256)

单向性：从输入的消息到它的哈希的正向运算简单且唯一确定，而反过来非常难，只能靠暴力枚举

灵敏性：输入的消息改变一点对它的哈希改变很大

高效性：从输入的消息到哈希的运算高效

均一性：每个哈希值被取到的概率应该基本相等

抗碰撞性

​ 弱抗碰撞性：给定一个消息`x`，找到另一个消息`x'`使得`hash(x) = hash(x')`是困难的

​ 强抗碰撞性：找到任意`x`和`x'`，使得`hash(x) = hash(x')`是困难的

- ### **ECDSA 椭圆曲线签名算法**

https://www.btcstudy.org/2022/06/13/part-2-bitcoin-p2tr-transaction-breakdown/#%E6%AF%94%E7%89%B9%E5%B8%81%E5%AF%86%E7%A0%81%E5%AD%A6%E5%9B%9E%E9%A1%BE

ECDSA

公钥 私钥

用私钥对交易签名（数字签名）

用公钥验证签名

secp256k1 y<sup>2</sup> = x<sup>3</sup> + 7

标准的 ECDS 签名由两个整数 r 和 s 构成

以太坊要求的签名格式 R S V

## 4.以太坊

### 账户

合约账户（CA）

​ nonce | balance | storage hash（状态变量）|code hash（代码）

外部账户（EOA）

​ balance | nonce

区别：合约账户不受私钥控制，无法自己发起交易，有自己的代码逻辑

### Gas EIP-1559

1.基本费用+小费

2.弹性的区块大小限制，每个区块更大的 gas 容量

**Transaction Fee = Gas 用量 \* Gas Price**

Gas Limit 获取

RPC API:eth_estimateGas

[2]Gas used 与执行合约所需要的操作数以及存储的数据量相关

**Gas price = Base Fee(系统开支,burnt) + Max Priority（矿工小费，支付给矿工）**

Base Fee 的动态平衡

上一个区块打包的 gas limit 是否使用超过了一半，区块空间增加 12.5%，Base Fee 上涨 12.5%

Ether = 10\*\*18 Wei

![image-20230802173354178](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230802173354178.png)

### 以太坊的区块与状态机

<img src="C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230721163220645.png" alt="image-20230721163220645"  />

![image-20230727001919419](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230727001919419.png)

- #### 状态树 MPT

哈希表的缺点

​ 节点间维护共识的工作量很大

Merkle Tree 的缺点

​ 账户顺序的不同形成完全不同的子树，hash 值不同，无法完成共识

​ 只多一个新账户，大半个颗树需要重构

前置知识点：Trie-字典树/前缀树

![img](https://tomni.notion.site/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F732749c9-8134-4b0a-a762-c75b9fa2c15d%2FUntitled.png?id=bfe31b36-807f-4cf4-87da-b375545b58af&table=block&spaceId=988dd3ef-a174-4a14-83dc-a91d3cfadc76&width=2000&userId=&cache=v2)

Patricia Trie 进一步压缩路径（稀疏）

- #### **Merkel-Patricia 树(mpt) 默克尔压缩前缀树**

![image-20230721170053183](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230721170053183.png)

空白节点 NULL

分支节点 branch Node

叶子节点 leaf Node

扩展节点 extension Node 出现路径压缩

该节点整个的 hash 值存放在父亲节点中

- #### 账户状态序列化 RLP

- #### 交易树 MPT

![image-20230721172819472](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230721172819472.png)

- #### 交易回执树 MPT

![image-20230721173640577](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230721173640577.png)

- Status： 成功与否，1 表示成功，0 表示失败；
- CumulativeGasUsed: 至当前交易为止，所有被执行的交易消耗的总 gas 费，确保交易执行顺序；
- Logs：当前交易执行所产生的智能合约事件列表；
- Bloom Filter：是从 Logs 中提取的事件布隆过滤器，用于快速检测某主题的事件是否存在于 Logs

**Bloom Filter**

快速判断某一个元素是否在某个集合中

## 5.Layer2

**直接从第一层以太坊共识获得安全性，二层执行交易，数据/结果锚定回主链**

**State Validity** 状态有效性

**Data Availability** 数据可用性

- 通道 Channel

- Plasma

- Validium、Volition

- ### Roll-up (扩容)

https://blog.chain.link/what-is-a-layer-2-zh/ 一文读懂 Layer 2

https://www.youtube.com/watch?v=I598C9GFDvk&t=3304s

https://mm.edrawsoft.cn/app/editor/E1etWI8xq28qQ5eI07Li3CBL7UfhK47j

https://web3caff.com/zh/archives/55667 Optimistic Rollups、zk-Rollups、

https://www.odaily.news/post/5185043

https://www.bilibili.com/video/BV1jS4y1s7S9/?spm_id_from=333.999.0.0&vd_source=5c46b11c8d760605427b8431fb93d551

Layer2 的数据压缩后上传到 Layer1，由 Layer1 保证数据可用性

如何验证提交的信息是正确可用的？（**状态有效性**）

- Optimistic Roll-Up 乐观卷叠

Fraud Proof 欺诈证明

​ Layer2 同步给以太网主网的过程中，如果有人认为被同步的数据不准确（即存在欺诈行为），便可在挑战期内发出挑战，通过 Layer1 智能合约计算并裁决出谁是欺诈者，再作出处罚。

​ Arbiturm(Optimistic Roll-Up)

​ 对仲裁过程进行了创新和优化。

- Zero Knowledge Roll-Up 零知识证明卷叠

Validity Proof 有效性证明

​ 将聚合的交易，使用零知识证明生成一个简洁的证明文件，证明有效性。

### 零知识证明 ZKP

新兴密码学工具，证明者可以用其向验证者证明某个声明，而不了泄露其他额外信息

https://www.bilibili.com/video/BV14v4y1o7Ye/?spm_id_from=333.337.search-card.all.click

https://github.com/sec-bit/learning-zkp/blob/master/zkp-resource-list.md

**完备性**

**可靠性**

**零知识**

# Solidity

**1.hello**

### **2.值类型**

**## Solidity 中的变量类型**

1.数值类型(Value Type):包括布尔型、整数型。这类变量赋值时直接传递数值 2.引用类型(Reference Type)：包括数组和结构体，这类变量占空间大，赋值时候直接传递地址(类似指针) 3.映射类型(Mapping Type):Solidity 里的哈希表 4.函数类型(Function Type):Solidity 把函数归到数值类型，但其实有区别

**## 数值类型**

solidity 没有浮点数(小数)

1.布尔型

\*\* && 和 || 运算符遵循短路原则

2.整型

int 整数，包括负数

uint 正整数

uint256 256 位正整数

3.地址类型

以太坊账户地址：20 字节 160 个二进制位 40 个 16 进制位

（4 个二进制可以表示一个 16 进制 160 个二进制位=40 个 16 进制位）

0x29C05AE3B3a11D4618562D92148eBF4b84C25fBA

地址类型有 balance 成员

用 payable 修饰的地址类型还多了 transfer 和 send 成员方法

address payable public \_address1 = payable(\_address);

4.定长字节数组

bytes

定长:byte1,bytes8,bytes32,属于数值类型，存数据，消耗 gas 少

数组长度在声明后不可改变

不定长:引用类型

5.枚举 enum

比较冷门

**# 3.函数类型 (external/internal/public/private, pure/view, payable)**

**## 函数可见性**

public:内部外部均可见

private:只能从本合约内部访问，继承的合约也不能用

external:只能从合约外部访问，但内部可以用 this.f()调用

internal:只能从合约内部访问，继承的合约可以用

注意 1：合约中定义的函数需要明确指定可见性，它们没有默认值。

注意 2：public|private|internal 也可用于修饰状态变量。public 变量会自动生成同名的 getter 函数，用于查询数值。未标明可见性类型的状态变量，默认为 internal

**## 函数权限 pure|view|payable**

决定函数权限/功能的关键字。

payable:(可支付的)

pure:包含 pure 关键字的函数，不能读取也不能写入存储在链上的状态变量

view:包含 view 关键字，能读取但不能写入状态变量

**# 4.函数输出 return**

**## 返回值 return 和 returns**

returns 加在函数名后面，用于声明返回的变量类型及类型名;

return 用于函数主体中，返回指定的变量

**## 返回多种变量**

**## 命名式返回**

**## 解构式赋值**

**# 5.数据存储位置（storage/memory/calldata）和变量作用域**

引用类型：数组(array)、结构体(struct)，这类变量占空间大，赋值时候直接传送地址。使用这类变量必须声明数据存储的位置

**## 数据位置 storage，memory 和 calldata**

**### storage**

**### memory**

函数参数和临时变量一般用 memory，存储在内存中，不上链

**### calldata**

赋值规则

产生独立副本

创建引用

**## 变量的作用域 状态变量 局部变量 全局变量**

**##### 1.状态变量**

状态变量在合约内、函数外声明

数据存储在链上，所有合约内函数都可以访问，gas 消耗高。

**##### 2.局部变量**

局部变量在函数内声明

仅在函数执行过程中有效的变量，函数退出后，变量无效。

局部变量的数据存储在内存里，不上链，gas 低。

function bar() external pure returns(uint){

uint xx = 1;

uint yy = 3;

uint zz = xx + yy;

return(zz);

}

**##### 3.全局变量**

全局范围工作的变量，都是 solidity 预留关键字。可以在函数内不声明直接使用

blockhash(uint blockNumber): (bytes32)给定区块的哈希值 – 只适用于 256 最近区块, 不包含当前区块。

block.coinbase: (address payable) 当前区块矿工的地址

block.gaslimit: (uint) 当前区块的 gaslimit

block.number: (uint) 当前区块的 number

block.timestamp: (uint) 当前区块的时间戳，为 unix 纪元以来的秒

gasleft(): (uint256) 剩余 gas

msg.data: (bytes calldata) 完整 call data

msg.sender: (address payable) 消息发送者 (当前 caller)

msg.sig: (bytes4) calldata 的前四个字节 (function identifier)

msg.value: (uint) 当前交易发送的 wei 值

tx.origin 返回发送交易的账户地址。最初的调用者

**##### 4.全局变量-以太单位与时间单位**

以太单位：

wei:1

gwei: 1e9 = 1000000000

ether: 1e18 = 1000000000000000000

时间单位：

seconds: 1

minutes: 60 seconds = 60

hours: 60 minutes = 3600

days: 24 hours = 86400

weeks: 7 days = 604800

**# 6.数组（array）和结构体（struct）**

**#### 数组**

\```solidity

//固定长度数组

uint[8] array1;

byte[5] array2;

address[100] array3;

//可变长度数组(动态数组) 声明时不指定数组的长度

uint[] array4;

byte[] array5;

address[] array6;

bytes array7;

//创建数组的规则

//memory 修饰的动态数组，用 new 操作符来创建，但必须声明长度，并且声明后长度不能改变

uint[] memory array8 = new uint[](5);

bytes memory array9 = new bytes(9);

//

//如果创建的是动态数组，需要一个一个元素的赋值

uint[] memory x = new uint[](3);

x[0] = 1;

x[1] = 3;

x[2] = 4;

//数组成员

length:包含元素数量，memory 数组的长度在创建后是固定的

push():动态数组和 bytes 拥有 push(),数组最后添加 0 元素

push(x):动态数组和 bytes 拥有 push(x)

pop():动态数组和 bytes 拥有 pop()

\```

**#### 结构体**

\```solidity

//结构体

struct Student{

uint256 id;

uint256 score;

}

Student student;//初始化一个 student 结构体

//给结构体赋值的两种方法：

//方法 1:在函数中创建一个 storage 的 struct 引用

function initStudent1() external{

Student storage \_student = student;

\_student.id = 11;

\_student.score = 100;

}

//方法 2:直接引用状态变量的 struct

function initStudent2() external{

student.id = 1;

student.score = 80;

}

\```

**# 7.映射 Mapping 类型**

通过键来查询对应的值

规则 1：键只能是默认类型，值可以是自定义类型

规则 2：存储位置必须是 storage

规则 3：如果映射（状态变量）声明为 public，solidity 会自动创建同名 getter 函数

规则 4

//声明映射

mapping(\_KeyType => \_ValueType)

//怎么理解二层 mapping

mapping(address => mapping(address => uint256)) public override allowance;

**# 8.变量初始值**

**### 值类型初始值**

**### 引用类型初始值**

**### delete 操作符**

**# 9.常量 constant 和 immutable**

状态变量的关键字 constant immutable

//constant 常量

//immutable 不变量

//只有数值变量可以声明 constant 和 immutable

//string 和 bytes 可以声明为 constant，但不能为 immutable

**# 10.控制流 control flow 和 插入排序 insertion sort**

//1.if-else

function ifElseTest(uint256 \_number) public pure returns(bool){

if(\_number == 0){

return(true);

} else{

return(false);

}

}

//2.for 循环

function forLoopTest() public pure returns(uint256){

uint sum = 0;

for(uint i = 0; i < 10; i++){

sum += i;

}

return(sum);

}

//3.while 循环

function whileTest() public pure returns(uint256){

uint sum = 0;

uint i = 0;

while(i < 10){

sum += i;

i++;

}

return(sum);

}

//4.do-while 循环

function doWhileTest() public pure returns(uint256){

uint sum = 0;

uint i = 0;

do{

sum += i;

i++;

} while (i < 10);

return(sum);

}

//5.三元运算符

**### 11.构造函数 constructor 和 修饰器 modifier**

\```solidity

\```

修饰器 modifier，声明函数拥有的特性

运行函数前的检查，例如地址，变量，余额

\```solidity

// 定义 modifier

modifier onlyOwner {

require(msg.sender == owner); // 检查调用者是否为 owner 地址

\_; // 如果是的话，继续运行函数主体；否则报错并 revert 交易

}

function changeOwner(address \_newOwner) external onlyOwner{

owner = \_newOwner; // 只有 owner 地址运行这个函数，并改变 owner

}

\```

**### 12.事件 event**

https://learnblockchain.cn/article/3536

**##### 声明事件**

**##### 释放事件**

\```solidity

event Transfer(address indexed from,address indexed to,uint256 value);

//indexed 关键字

//indexed 标记的变量 检测事件的索引“键” 每个事件最多有 3 个带 indexed 的变量 英文技术文档中也称作 topic[0] topic[1] topic[2]

//不带 indexed 的变量，会存储在事件的 data 部分中，事件的“值”

//可以在函数里释放事件

fuction \_transfer(

address from,

address to,

uint256 amount

) external{

\_balances[from] = 10000000;//

\_balances[from] -= amount;

\_balances[to] += amount;

//释放事件

emit Transfer(from,to,amount);

}

\```

**#### EVM 日志 Log topic**

主题 topics + 数据 data 两部分

Topics[0] 是事件的签名(哈希)

Topic 还可以包含至多 3 个 indexed 参数

数据 data

**# 13.继承 inheritance**

简单继承、多重继承、修饰器(modifier)和构造函数(constructor)

**### 构造函数的继承**

contract C is A{

constructor(uint \_c) A(参数){

}

}

**# 14.抽象合约(abstract)和接口(interface)**

**## 抽象合约**

合约中至少有一个未实现的函数

**## 接口**

接口不实现任何功能

是智能合约的骨架，定义了合约的功能以及如何触发它们

接口提供了两个重要的信息：

1.合约里每个函数的 bytes4 选择器，以及函数签名 函数名（每个参数类型）

2.接口 id（EIP-165）https://blog.csdn.net/NatureOrigin/article/details/130655314

接口与合约 ABI 等价，可以相互转换

**### 使用接口**

接口 对象 name = 接口(合约地址)

IERC721 Azuki = IERC721(0xED5AF388653567Af2F388E6224dC7C4b3241C544)

**# 15.异常** Error

error,require,assert

**## error**

solidity 0.8.4 新内容

可以携带参数，搭配 revert

**## require**

solidity 0.8 版本之前的

require(检查条件，“异常的描述”)

缺点：gas 随着描述异常的字符串长度增加

**## assert**

不能解释抛出异常的原因（比 require 少个字符串）

**# 16.函数重载(overloading)**

重载:名字相同但输入参数类型不同的函数

修饰器不可重载，构造函数可以重载

**# 17.库合约 Library**

库合约是一系列的函数合集

和普通合约的区别： 1.不能存在状态变量 2.不能够继承或被继承 3.不能接收以太币 4.不可以被销毁

**### 调用库合约 String**

1.using A for B;

添加库合约(A)到类型 B，库 A 的函数会自动添加为 B 类型变量的成员 2.通过库合约名称调用函数

**在调用的时候，这个变量会被当作第一个参数传递给函数**

**### 常用库合约**

1.String：将 uint256 转换为 String

2.Address：判断某个地址是否为合约地址

3.Create2：更安全地使用 Create2 EVM opcode

4.Arrays：跟数组相关的库合约

**# 18.import**

1.通过源文件相对位置导入

import './Yeye.sol';

2.通过源文件网址导入网上的合约

// 通过网址引用

import 'https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/Address.sol';

3.通过 npm 的目录导入

import '@openzeppelin/contracts/access/Ownable.sol';

4.通过全局符号导入特定的合约

import {Yeye} from './Yeye.sol'

**# 19.接收 ETH(receive/fallback)**

**## 接收 ETH 函数 receive**

receive() external payable {...}

在合约收到 ETH 时被调用的函数

receive()最好不要执行太多的逻辑，不然会触发 Out of Gas 报错

**## 回退函数 fallback**

fallback()函数会在调用合约不存在的函数时触发

可用于接收 ETH,也可用于代理合约 proxy contract

fallback() external payable(一般也需要) {...}

回调函数不能有返回值

**## receive 和 fallback 的区别**

简单来说，合约接收`ETH`时，`msg.data`为空且存在`receive()`时，会触发`receive()`；`msg.data`不为空或不存在`receive()`时，会触发`fallback()`，此时`fallback()`必须为`payable`。

`receive()`和`payable fallback()`均不存在的时候，向合约**\*\*直接\*\***发送`ETH`将会报错（你仍可以通过带有`payable`的函数向合约发送`ETH`）。

**# 20.发送 ETH(transfer/send/call)**

**## call**

接收方地址.call{value:发送 ETH 数额}("")

没有 gas 限制，可以支持对方合约 fallback()或 receive()实现复杂逻辑

若转账失败，不会 revert

**## transfer**

接收方地址.transfer(发送 ETH 数额)

gas 限制 2300

如果转账失败，会自动 revert(回滚交易)

**## send**

接收方地址.send(发送 ETH 数额)

gas 限制 2300

如果转账失败，不会 revert

**# 21.调用其他合约**

**## 调用合约的两种方式**

1.已知合约的地址和合约代码（接口）

创建合约引用 2.不知道合约源代码或 ABI，没法生成合约变量，这时可以通过 call 调用对方合约的函数

**# 22.Call**

**## call**

call 是 address 类型的低级成员函数，用来与其他合约交互

返回值 (bool,data) (成功与否，目标函数的返回值)

\- call 是 solidity 官方推荐的通过触发 fallback 和 receive 函数发送 ETH 的方法

**## call 的使用规则**

目标合约地址.call{value:发送数额,gas:gas 数额}(字节码);

字节码: abi.encodeWithSignature("函数签名",逗号分隔的具体参数)

abi.encodeWithSignature("f(uint256,address)",\_x,\_addr)

**# 23.Delegatecall**

**## delegatecall**

delegatecall 与 call 类似，是 solidity 中地址类型的低级成员函数

可以指定 gas，但不能指定发送的 ETH 数额

1.从 Target Contract 角度来看：msg.sender 是 user，而不是 Proxy，即 Proxy 对 user 的请求进行了透传； 2.在 Target Contract 被调用时，使用的是 Proxy 的上下文，即执行合约带来的状态变化会存在 Proxy 中，而不是 Target Contract 之中

**## delegatecall 使用规则**

**## 应用场景**

1.代理合约

2.EIP-2535 Diamonds（钻石）

**# 24.Create**

智能合约同样可以创建新的智能合约

Contract x = new Contract{value: \_value}(params)

合约名 对象名 = new 合约名

计算地址

创建者的地址(通常为部署的钱包地址或者合约地址)和 nonce(该地址发送交易的总数,对于合约账户是创建的合约总数,每创建一个合约 nonce+1))的哈希

新地址 = hash(创建者地址, nonce)

在智能合约内创建合约，创建者地址是智能合约，nonce + 1

和智能合约交互不会使智能合约 nonce 增加

**# 25.Create2**

创建方式和计算地址方式不是一回事

**## Create2**

1.多一个 salt 参数

Contract x = new Contract{salt:\_salt,value: \_value}(params)

合约名 对象名 = new 合约名 2.内联汇编

assembly {

pair := create2(0, add(bytecode, 32), mload(bytecode), salt)

}

**## Create2 如何计算地址**

新地址 = hash("0xFF",创建者地址, salt, bytecode)

以 uniswap v2 为例

pair = address(uint(keccak256(abi.encodePacked(

hex'ff',

factory,

keccak256(abi.encodePacked(token0, token1)),

hex'96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f' // init code hash

))));

**# 26.删除合约**

selfdestruct(\_addr);

**# 27.ABI 编码解码**

ABI (Application Binary Interface，应用二进制接口)是与以太坊智能合约交互的标准

ABI 编码有 4 个函数：abi.encode`, `abi.encodePacked`, `abi.encodeWithSignature`, `abi.encodeWithSelector

**## abi.encode**

与智能合约交互。将每个参数填充为 32 字节的数据，拼接

**## abi.encodePacked**

适用于省空间，并不与合约交互的时候，例如算一些数据的 hash。将给定参数根据其空间所需最低空间编码，会把其中填充的很多 0 省略

\```solidity

function encodePacked() public view returns(bytes memory result) {

result = abi.encodePacked(x, addr, name, array);

}

\```

abi.encodeWithSignature 与 abi.encode 功能类似，只不过第一个参数为函数签名，等同于在 abi.encode 编码结果前加上 4 字节的函数选择器

\```solidity

function encodeWithSignature() public view returns(bytes memory result) {

result = abi.encodeWithSignature("foo(uint256,address,string,uint256[2])", x, addr, name, array);

}

\```

abi.encodeWithSelector 与`abi.encodeWithSignature`功能类似，第一个参数为函数选择器,为函数签名 Keccak 哈希的前 4 个字节

function encodeWithSelector() public view returns(bytes memory result) {

result = abi.encodeWithSelector(bytes4(keccak256("foo(uint256,address,string,uint256[2])")), x, addr, name, array);

}

abi.decode

**#### ABI 的使用场景**

1.在合约开发中，ABI 常配合 call 来实现对合约的底层调用

2.ethers.js 中常用 ABI 实现合约的导入和函数调用

**# 28.Hash**

**## Keccak256**

哈希 = keccak256(数据);

**## 生成数据唯一标识**

abi.encodePacked 将不同类型的数据打包编码,再用 keccak256 来生成唯一标识

keccak256(abi.encodePacked(\_num,\_string,\_addr));

**# 29.Selector**

调用智能合约，本质上是向目标合约发送了一段 calldata

calldata 就是告诉智能合约，我要调用哪个函数，以及参数是什么

remix 中的 input 即为此次交易的 calldata

**## msg.data**

msg.data 是 solidity 中的一个全局变量，值为完整的 calldata（调用函数时传入的数据

发送的 calldata 中前 4 个字节是 selector 函数选择器。1 个字节占 2 个 16 进制位

**## method id、selector 和函数签名**

method id：函数签名的 Kecaak 哈希后的前 4 个字节

函数签名:函数名（逗号分隔的参数类型）

selector：calldata 的前 4 个字节

**## 使用 selector**

address.call(abi.encodeWithSelector(method id,参数));

**## call 的使用规则**

目标合约地址.call{value:发送数额,gas:gas 数额}(字节码);

字节码: abi.encodeWithSignature("函数签名",逗号分隔的具体参数)

abi.encodeWithSignature("f(uint256,address)",\_x,\_addr)

是一回事吧？

**# 30.异常处理 Try-Catch**

**# 31.ERC20**

**##### 状态变量**

// 账户持有余额

键 值

mapping(address => uint256) private \_balances;

// 账户授权数量

键 键 值

mapping(address => mapping(address => uint256)) private \_allowances;

// 总发行量

uint256 private \_totalSupply;

// 代币名称

string private \_name;

// 代币标识

string private \_symbol;

**##### 事件**

event Transfer(address indexed \_from, address indexed \_to, uint256 \_value)

event Approval(address indexed \_owner, address indexed \_spender, uint256 \_value)

**##### 函数**

totalSupply() 查询目前代币总供给

balanceOf() 查询某账户的代币余额

\*transfer() 调用者账户向某账户转账

allowance() 查询某账户给另外一个账户授权的额度

\*approve() 调用者给某账户授权一定额度的代币

\*transferFrom() 被授权者转移代币

**### 32.Faucet.sol**

\```solidity

requestTokens(){

IERC20 token = IERC20(tokenContract);//创建合约对象

为什么是用 IERC20 而不是 ERC20，IERC20 是接口，明明没有实现

答：14.抽象合约和接口中讲到，如果一个合约实现了接口，我们不需要知道它具体代码的实现，就可以与它交互。

token.balanceOf()

token.transfer()

}

\```

**## 34.ERC721**

EIP 全称 以太坊改进建议

ERC 全称 以太坊意见征求稿

**#### ERC165**

接口 id 怎么算出来

**#### IERC721**

**#### IERC721Receiver**

NFT 也可以转给合约地址，但如果一个合约没有实现 ERC721 的相关函数，转入的 NFT 就进了黑洞，永远转不出来了。为了防止误转账，ERC721 实现了 safeTransferFrom()安全转账函数，目标合约必须实现了 IERC721Receiver 接口才能接收 ERC721 代币，不然会 revert

\_checkOnERC721Received()

**#### IERC721Metadata**

**#### ERC721 主合约**

\*3 个库合约:Address.sol Context.sol 和 Strings.sol

\*3 个接口:IERC721.sol IERC721Receiver.sol IERC721Metadata.sol

\*ERC165.sol

**##### Address 库**

ERC721 用到的功能，isContract()

**##### Context 库**

**##### Strings 库**

**### ERC165 与 IERC721Receiver.sol**

接口是某些行为的集合

ERC165 对外表明自己实现了哪些接口的技术标准

**### 开图/盲盒**

https://mirror.xyz/xyyme.eth/5Zhl6oCdt_h-a4bMij3SVlmaTEo_bXpcukrZsfHzsHE

开图时，owner 发交易设置 baseURI。将图片统一上传至 ipfs，将描述文件 json 统一上传至 ipfs

https://juejin.cn/post/7150233032488812581

**# 36.Merkle Tree**

自下而上构建

**### 生成 Merkle Tree**

https://lab.miguelmota.com/merkletreejs/example/

**### Merkle Proof 验证**

MerkleProof 库有三个函数

1.verify()函数

2.processProof()函数 利用 proof 和 leaf 依次计算出 Merkle Tree 的 root，调用了\_hashPair()函数

3.\_hashPair()函数 用 keccak256()函数计算非根节点对应的两个子节点的哈希

**### 利用 Merkle Tree 发放 NFT 白名单**

MerkleTree 合约继承了 ERC721 标准，并利用 MerkleProof 库

**#### mint**

还需要借助对应的 proof

**### 总结**

利用 Merkle Tree 来发放白名单，链上只需要存储一个根植，非常节省 gas

**##### 实际**

复杂的 Merkle Tree 利用 javascript 库 merkletreejs 来生成管理，链上只需要存储一个根植，非常节省 gas。

**# 37.数字签名 Signature**

**### 数字签名**

以太坊使用的数字签名算法叫双椭圆曲线数字签名算法（ECDSA） 1.身份认证 2.不可否认 3.完整性

**### ECDSA 合约**

ECDSA 标准包含两个部分： 1.签名者利用 私钥 对 消息 创建 签名 2.其他人使用 消息 和 签名 恢复 公钥 来验证。

签名过程：ECDSA\*正向算法（消息+私钥+随机数）= 签名

验证过程：ECDSA\*反向算法（消息+签名）= 公钥

**#### 创建签名**

1.打包消息 2.计算以太坊签名消息

3-1.利用钱包签名

3-2.利用 web3.py 签名

**#### 验证签名**

验证者需要 消息，签名和公钥。只有私钥的持有者才能生成这样的签名 4.通过签名和消息恢复公钥 rsv 签名

recoverSigner 有点难懂

**### 链下签名实现白名单**

**# 39.随机数**

**### 41.WETH**

**##### 事件**

Deposit

Withdraw

**##### 函数**

回调函数:fallback()和 receive(),当用户往 WETH 合约转 ETH 时，会自动触发 deposit()存款函数，获得等量的 WETH

deposit()

withdraw()

**### 42.分账**

**## 46.代理合约**

**#### 代理模式**

代理模式将合约数据和逻辑分开，

数据（状态变量）存储在代理合约中

逻辑（函数）保存在逻辑合约中

代理合约（Proxy）通过 delegatecall，将函数调用全权委托给逻辑合约（implementation）执行，再把最终的结果返回给调用者（Caller） 1.可升级 2.省 gas：

**##### 代理合约**

三部分：代理合约 Proxy，逻辑合约 Logic，和一个调用示例 Caller

\- 部署逻辑合约 Logic

\- 创建代理合约 Proxy，状态变量 implementation 记录 Logic 合约地址

\- Proxy 合约利用回调函数 fallback，将所有调用委托给 Logic 合约

\- 部署调用示例 Caller 合约，调用 Proxy 合约

\- **\*\*注意\*\***：`Logic`合约和`Proxy`合约的状态变量存储结构相同，不然`delegatecall`会产生意想不到的行为，有安全隐患。

**##### 代理合约 Proxy**

1 个状态变量 implementation

1 个构建函数

1 个回调函数

内联汇编

让本来不能有返回值的回调函数有了返回值

**##### 逻辑合约 Logic**

占位变量，与 Proxy 保持一致，防止插槽冲突

**##### 调用者合约 Caller**

**## 47.可升级合约**

可升级合约是可以更改逻辑合约的代理合约

管理员可以通过升级函数更改逻辑合约地址，从而改变合约的逻辑

选择器冲突？

**## 48.透明代理**

模板：openzeppelin 的 TransparentUpgradeableProxy

**### 选择器冲突**

函数选择器仅有 4 个字节，范围很小，两个不同的函数可能会有相同的选择器

同一合约下，选择器冲突无法编译成功

逻辑合约的 a 函数和代理的合约升级函数的选择器相同

**### 透明代理**

限制管理员的权限，不让他调用任何逻辑合约的函数

\*管理员变为工具人，仅能调用代理合约的可升级函数对合约升级，不能通过回调函数调用逻辑合约

\*其他用户不能调用可升级函数，但是可以调用逻辑合约的函数

**### 存储冲突**

proxy 的 storage 布局与目标合约的 storage 布局不相同

目标合约在 Proxy 合约中修改了错误的位置

解决存储冲突：

两个状态变量

将 Proxy 中的默认 slot 留出来，不要占用 EIP-1967 标准

**## 49.UUPS 通用可升级代理**

标准 升级函数 是否会“选择器冲突” 缺点

可升级代理 Proxy 合约 会 选择器冲突

透明代理 Proxy 合约 不会 费 gas

UUPS Logic 合约 不会 更复杂

**### UUPS 合约**

升级函数放在逻辑合约中

**## 52.EIP712 类型化数据签名**

展示签名信息的原始数据,使签名内容可视化

**### EIP712 使用方法**

EIP712 的应用一般包含链下签名（前端或脚本）和链上验证（合约）两部分

**#### 链下签名**

1.EIP712 签名必须包含 EIP712Domain

const domain = {

name:"",

version:"",

chainId:"",

verifyingContract:"",

}; 2.根据场景自定义签名的数据类型

const types = {

Storage: [

{ name: "spender", type: "address" },

{ name: "number", type: "uint256" },

],

};

3.message 变量，传入要被签名的类型化数据

const message = {

spender:"",

number:"",

}; 4.调用钱包对象的 signTypedData(domain,types,message)进行签名;

**#### 链上验证**

EIP712Storage 合约

**##### 5 个状态变量**

1.EIP712DOMAIN_TYPEHASH

2.STORAGE_TYPEHASH

3.DOMAIN_SEPARATOR

4.number

5.owner

**##### 3 个函数**

1.构造函数

2.retrieve()

3.permitStore

**## 53.ERC2612 ERC20Permit**

**### 合约**

**#### IERC20Permit 接口合约**

permit()

nonces()

DOMAIN_SEPARATOR()

**## 55.MultiCall**

1.方便性：在一次交易中对不同合约的不同函数进行调用

2.节省 gas

3.原子性

调用结构体 Call

结果结构体 Result

**# 智能合约开发流程**

**## 需求&功能**

**## 接口&数据**

**## 实现&优化**

**## 测试&检查**

**## 配置&部署**

**## SDK&文档**

Solidity 注解

@title

@author

@notice

@dev

@param

@return

# EVM 与 Gas 优化

## EVM 基础

![image-20230806010136503](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230806010136503.png)

堆栈 Stack：每个元素 256 位（32 字节），最大深度为 1024 个元素

内存 Memory：支持以 8 位或 256 位写入（MSTORE8/MSTORE），但只支持 256 位取（MLOAD） **易失性**

存储 Storage：对存储的读取（SLOAD）和写入（STORE）都需要昂贵的 gas

EVM 字节码：智能合约 -> EVM 字节码

程序计数器 PC：跟踪当前执行指令位置的寄存器

Gas：SSTORE 操作消耗 20000 gas SLOAD 操作消耗 200 Gas

**执行模型**：接收交易，以太坊初始化一个新的执行环境加载合约的字节码

## 杂

https://github.com/AmazingAng/WTF-Solidity/tree/main/Topics/Translation/DiveEVM2017

EVM 是执行 bytecode（字节码）的机器

智能合约被编译成一串二进制码 solc --bin

字节码 操作码 函数选择器

内存 合约存储 Geth 存储区

evm 中显示的 3 种调用方式

1.Call 一般的跨合約函數調用方式，這通常會改變被調用合約的存儲

2.StaticCall 靜態調用，不會改變被調用合約的存儲，是屬於跨合約讀取狀態變數的操作

3.DelegateCall 委任調用，`msg.sender` 不會改變，通常用於 Proxy 代理模式

存储 Storage

存储变量声明不需要任何费用，因为不需要初始化

内存 Memory

调用数据 Calldata

堆栈 Stack

代码 Code

函数签名：函数名(参数类型)

abi.encodeWithSignature(签名，参数)

abi.encodeWithSelector(bytes4(keccak256(签名)),参数)

每个参数 32 个字节

## 字节码分析：合约部署流程

字节码分成了三部分：

1. init bytecode（部署代码 初始化字节码）
2. runtime bytecode（合约代码 运行时字节码）
3. metadata hash（辅助数据 合约的一些 meta 信息哈希）

![image-20230720145442363](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230720145442363.png)

部署代码的工作任务

1.运行构造函数的逻辑

2.获取 runtime bytecode 与 metadata hash 的内容并返回给 EVM

## 字节码分析：合约调用

与智能合约交互的本质就是向其发送一段 calldata

选择器+参数编码+...+参数编码

选择器是方法签名的 keccak256 哈希

用户的请求是如何**匹配到具体的方法**的？

在栈中与所有的方法签名哈希值一个一个进行必对，如果相同，则跳转到相应的部分执行对应方法内容。

### ABI 编码(Calldata)

静态类型的 Calldata 编码

**动态类型**的 Calldata 编码（字符串、字节、数组）

https://blog.softbinator.com/solving-ethernaut-level-29-switch/

https://docs.soliditylang.org/en/develop/abi-spec.html

头尾编码

按顺序编码，静态参数按原样编码，动态数组的数据放在尾部

每个 32 字节，动态数组的数据从第 96 个字节(0x60)开始，由指针存放

```
/************* HEAD (32*3 bytes) *************/
// arg1: 0xaaaa
000000000000000000000000000000000000000000000000000000000000aaaa
// arg2: look at position 0x60 for array data
0000000000000000000000000000000000000000000000000000000000000060
// arg3: 0xbbbb
000000000000000000000000000000000000000000000000000000000000bbbb


/************* TAIL (128 bytes) *************/
// position 0x60. Data for arg2.
0000000000000000000000000000000000000000000000000000000000000003
00000000000000000000000000000000000000000000000000000000000000b1
00000000000000000000000000000000000000000000000000000000000000b2
00000000000000000000000000000000000000000000000000000000000000b3
```

函数选择器只要前 4 个字节是为了省 gas

## 合约存储 状态变量存储结构

插槽

插槽包装 slot packing

固定长度

![image-20230723212556587](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230723212556587.png)

SSTORE (key,value)

SLOAD (key)

存取被包装(slot packing)的变量

slot packing 使用 AND，OR 和 NOT 三个位运算

插槽操作：存储包装变量 SSTORE

![img](https://substackcdn.com/image/fetch/w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F3ce409fd-b942-42d7-b7ea-02c18c4a8993_782x338.png)

插槽操作：取出被包装的变量 SLOAD

https://mirror.xyz/xyyme.eth/5eu3_7f7275rqY-fNMUP5BKS8izV9Tshmv8Z5H9bsec

https://github.com/AmazingAng/WTF-Solidity/blob/main/Topics/Translation/DiveEVM2017/DiveEVM2017-Part3.md

## 合约的内存分布

Memory 数据结构：简单的字节数组

存储：1 字节（8 位）或 32 字节（256 位）为单位存储

读取：只能以 32 字节为单位读取

​ 读取可以从任意字节处开始读取

![image-20230720145953327](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230720145953327.png)

##### 操作内存的 3 个操作符

- MSTORE(x,y)-在内存 x 处开始存储 32 字节的数据 y
- MLOAD(x)-将内存 x 处开始的 32 字节数据加载到栈中
- MSTORE8(x,y)-在内存 x 处存储 1 字节数据 y（32 字节栈值中的最低有效字节）

##### Solidity 中预留了 4 个 32 字节的插槽（slot）

预留空间占据 128 个字节

- `0x00` - `0x3f` (64 字节): 哈希方法的暂存空间
- `0x40` - `0x5f` (32 字节): 当前已分配内存大小 (也称为空闲内存指针)
- `0x60` - `0x7f` (32 字节): 零槽，用作动态内存数组的初始值，永远不能写入值

![image-20230720150750599](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230720150750599.png)

##### 空闲指针

新的空闲指针位置 = 旧的空闲指针位置+分配的数据大小

## 操作码

https://www.evm.codes/playground

```
50	POP			弹出栈顶端
54	SLOAD		从合约存储中加载值
55	SSTORE		[位置，值]
60	PUSH1		压栈
80	DUP1		复制栈顶
81	DUP2		复制栈顶的第二个元素，压栈
90	SWAP1		交换栈顶的两个元素
	CODECOPY	复制交易数据calldata，加载到内存中
```

堆栈指令：PUSH POP

算术指令：ADD MUL SUB DIV

比较指令：LT（小于）GT（大于）EQ（相等）ISZERO（相等）

位运算指令：AND（与）OR（或）XOR（异或）SHL（左移位）SHR（右移位）

内存指令：MSTORE（内存写 256 位）MSTORE（内存写 8 位）MLOAD(内存读)
存储指令：

上下文指令

## Solidity 事件

作为替代返回值

更便宜的替代数据存储

DApp 客户端可以订阅

https://github.com/AmazingAng/WTF-Solidity/blob/main/Topics/Translation/DiveEVM2017/DiveEVM2017-Part6.md

## Gas 优化

通过目标合约代码或接口来创建合约的引用，调用目标合约的函数，gas 是否相同

https://github.com/WTFAcademy/WTF-gas-optimization

https://learnblockchain.cn/article/4167#%E5%87%8F%E5%B0%91%E5%AF%B9%E5%A4%96%E9%83%A8%E5%90%88%E7%BA%A6%E5%9C%B0%E5%9D%80%E7%9A%84%E8%B0%83%E7%94%A8

**1 constant immutable 与 变量**

​ 对不变量，使用 constant 和 immutable 定义

2 calldata memory

3 Bitmap 位图与位运算符

**4 Unchecked **

​ 在安全性可控的情况下，使用 unchecked

**5 Uint**

​ 尽量变量打包，如果不能打包还是使用 uint256

6 Error

**7 LocalData 局部变量**

​ 在较复杂的运算中，避免直接操作存储变量，可以先定义一个局部变量，对局部变量进行修改后，再修改存储变量

**8 clone 替换 new/create2**

**9 Packing**

**10 Increment 使用更好的递增**

​ 使用 ++i，而不是 i++

**11 使用 uint 进行重入保护**

​ bool ×

​ uint 0 1，0 改 非 0 会消耗较多的 gas ×

​ uint 1 2，非 0 改非 0 √

12 LessThan

**13 MethodName 优化方法名**

​ EVM 执行交易，calldata 数据消耗 gas，0 字节消耗 4gas，非 0 字节消耗 16gas

​ 优化函数名，增加 MethodId 中 0 字节的个数来节省 gas

**14 MethodIdSort**

​ 合约内所有函数是一个数组，通过 MethodId 排序

​ 遍历,找到想调用的函数

# Ethers.js API

账户 10：77e7e7d4db6930c32590eb803d74265679c64212de31fbdb178eeb377f0c2525

## ABI

## Addresses

**getCreate2Address**

**getCreateAddress**({from:,nonce:})

## Provider

以太坊网络连接的抽象

```typescript
const provider = new ethers.JsonRpcProvider(ALCHEMY_MAINNET_URL)
```

provider.getTransactionCount(address) 获得地址的 nonce，交易次数

合约地址的第一个 nonce 用来创建自己了(nonce 0)

provider.getBalance

## Signers

Signer 是以太坊账户的抽象，签名消息和交易，将签名的交易发送到以太坊网络，并更改区块链状态。

Signer 是抽象类，不能实例化。常见子类:Wallet,JsonRpcSigner

### Wallet

// 创建钱包对象

如果在创建时没有用到 provider，需要用 connect(provider)，连接到以太坊节点

const wallet1 = new ethers.Wallet.createRandom()

const wallet2 = new ethers.Wallet(privateKey,provider)

const wallet3 = new ethers.Wallet.fromMnemonic(mnemonic.phrase)

// 属性

wallet.address

// 方法

await wallet.getAddress()

await wallet.signMessage()

await wallet.signTransaction()

## 合约

将链上的合约的调用(calls)和交易(transactions)序列化，并反序列化它们的结果(results)和触发的日志(logs)

### Contract

Contract 对象是部署在以太坊网络上合约(EVM 字节码)的抽象

创建实例

const contract = new ethers.Contract(address,abi,signerOrProvider)

contract.attach(addressOrName)

contract.connect(providerOrSigner)

ethers.getContractAt() **hardhat**

**只读方法**

**可写方法**

const tx = await contract.METHOD_NAME(args ,{, overrides})

await tx.wait()

事件（Events）

const transferEvents = await contract.queryFilter('事件名',起始区块,结束区块) => _Promise< Array< Event > >_

要检索的事件必须包含在合约的 abi 中

解析事件的结果

tansferEvents[0].args["from"]

tansferEvents[0].args["to"]

tansferEvents[0].args["amount"]

监听

contract.on(event,function)

contract.once(event,function)

```typescript
// 持续监听USDT合约
console.log('\n2. 利用contract.on()，持续监听Transfer事件')
contractUSDT.on('Transfer', (from, to, value) => {
  console.log(
    //打印结果
    `${from} -> ${to} ${ethers.formatUnits(ethers.getBigInt(value), 6)}`
  )
})
```

事件过滤器

```typescript
let filterBinanceIn = contractUSDT.filters.Transfer(null, accountBinance)
contractUSDT.on(filterBinanceIn, (res) => {
  console.log(
    `${res.args[0]} -> ${res.args[1]} ${ethers.formatUnits(res.args[2], 6)}`
  )
})
```

staticCall

```typescript
const tx = await contract.函数名.staticCall(参数, { override })
console.log(`交易会成功吗？：`, tx)
```

### ContractFactory

部署智能合约的本质是什么？

ContractFactory 是合约字节码的抽象

const contractFactory = new ethers.ContractFactory(abi,bytecode,signer)

const contract = await contractFactory.deploy(args)

await contract.deployed()（v5 写法）

await contract.waitForDeployment()

## ABI

functions

ethers.decodeBytes32String

ethers.encodeBytes32String

const interface = ethers.Interface(abi)

const interface2 = contract.interface

## 实用工具 Utilities

- ### BigNumber 大数

BigInt 实例

- ### 编码工具

- ### Display Logic and Input 显示逻辑和输入

ethers.formatUnits(变量，单位) 小单位 -> 大单位

ethers.parseUnits(变量，单位) 大单位 -> 小单位

## hardhat-erthers

```javascript
function getContractAt(name: string, address: string, signer?: ethers.Signer): Promise<ethers.Contract>;
function getSigners() => Promise<ethers.Signer[]>;
function getSigner(address: string) => Promise<ethers.Signer>;
```

**## 11.StaticCall**

\```javascript

const tx = await contract.函数名.staticCall(参数,{override})

console.log(`交易会成功吗？：`, tx)

\```

**## 12.识别 ERC721 合约**

ERC165:

通过 ERC165 接口合约，智能合约可以声明它支持的接口，供其他合约检查

ERC721 合约中会实现 IERC165 接口合约的 supportsInterface 函数，并且当查询 0x80ac58cd（ERC721 接口 id）时返回 true。

**## 13.编码 calldata 接口类 Interface**

利用接口类编码 calldata 与 WETH 合约交互。与一些特殊的合约交互时（比如代理合约）

利用 abi 生成或直接从合约中获取 interface 变量

\```javascript

// 利用 abi 生成

const interface = ethers.Interface(abi)

// 直接从 contract 中获取

const interface2 = contract.interface

\```

**## 14.批量生成钱包**

HD 钱包：BIP32、BIP44、BIP39

https://learnblockchain.cn/2018/09/28/hdwallet

BIP32：一个随机数种子通过分层确定性推导得到 n 个私钥

BIP44：为 BIP32 的衍生路径提供了一套通用规范。m / purpose' / coin_type' / account' / change / address_index

BIP39：以助记词方式生成种子

**## 15.批量转账**

调用 Airdrop 合约，实现批量转账

**## 16.批量归集**

**## 17.MerkleTree 脚本**

MerkleTree.js 构建 Merkle Tree 和 Merkle Proof 的 Javascript 包

SHA-256(btc) Keccak-256(eth)

hashLeaves sortPairs(排序)

npm install merkletreejs

**## 18.数字签名脚本**

**### 生成数字签名**

\```javascript

// 打包消息

const msgHash = ethers.solidityKeccak256()

// 获得符合 EIP191 标准的由私钥签署的签名

const messageHashBytes = ethers.getBytes(msgHash)

const signature = await wallet.signMessage(messageHashBytes);

\```

// 通过签名分发 NFT 白名单流程：

//

// 在服务器保管 signer 钱包的私钥-公钥对

// -> 在服务器记录 allowlist（白名单地址）和 tokenId，并生成对应的 msgHash，

// -> 用 signer 钱包给 msgHash 签名

// -> 部署 NFT 合约，初始化时 signer 的公钥保存在合约中。

// -> 用户 mint 时填地址和 tokenId，并向服务器请求签名。

// -> 调用合约的 mint()函数进行铸造

**## 19.监听 Mempool**

provider 是 WebSocket Provider，更持久的监听交易。因此，我们需要将 url 换成 wss 的。

Provider 类提供的方法，监听 mempool 中的 pending（未决，代打包）

**#### 监听 mempool**

provider.on("pending",listener)

**#### 脚本**

throttle 限制请求频率

let tx = await provider.getTransaction(txHash);

**## 20.解码交易详情 Interface 类**

**#### 未决交易 Pending Transaction**

Decode Input Data

**#### Interface 类 解码交易数据**

声明 Interface 类型和声明 abi 的方法差不多,abi 是智能合约的 abi

\```javascript

const iface = ethers.Interface([

"function balanceOf(address) public view returns(uint)",

"function transfer(address, uint) public returns (bool)",

"function approve(address, uint256) public returns (bool)"

]);

\```

**#### 脚本**

//filter pendingTx.data indexOf 检索字符串 selector 函数选择器

if (tx.data.indexOf(iface.getFunction("transfer").selector) !== -1)

//解码的交易详情

let parsedTx = iface.parseTransaction(tx)

console.log(parsedTx);

consloe.log(parsedTx.args);

**## 21.靓号生成器**

**## 22.读取任意数据**

private 变量并不私密

**### 智能合约存储布局**

以太坊智能合约的存储是一个 uint256 -> uint256

uint256 大小为 32 bytes

这个固定大小的存储空间被称为 slot（插槽）

**### getStorageAt**

getStorageAt() 方便开发者读取特定 slot 的值

\```javascript

const value = await provider.getStorageAt(contractAddress,slot)

\```

问题：怎么确定想读数据的 slot 呢

**## 23.抢先交易脚本**

//监听 mempool，读取交易详情 tx，筛选 mint()且发送方不是自己钱包地址的交易

let tx = await provider.getTransaction(txHash);

if (tx.data.indexOf(iface.getFunction("mint").selector) !== -1 && tx.from != wallet.address ){

}

//构建抢跑交易 txFronrun

\```javascript

const txFrontrun = {

to: tx.to,

value: tx.value,

maxPriorityFeePerGas: tx.maxPriorityFeePerGas \* 2n,

maxFeePerGas: tx.maxFeePerGas \* 2n,

gasLimit: tx.gasLimit \* 2,

data: tx.data

}

\```

**## 24.识别 ERC20**

可以用 ERC165 识别 ERC721，但我们 ERC20 早于 ERC165

1.获得合约的 bytecode

2.检查合约 bytecode 是否包含相应的函数选择器

**## 25.Flashbots**

**## 26.EIP712 签名脚本**

https://learnblockchain.cn/article/4357

Solidity：使用 Ethers.js 的 Solidity 存储变量

# 智能合约审计

https://www.theblockbeats.info/news/36184

## 流程

https://www.bilibili.com/video/BV1VV4y1b7ah/?spm_id_from=333.788&vd_source=5c46b11c8d760605427b8431fb93d551

1.下载代码

2.了解项目架构

3.

静态分析工具

slither

https://learnblockchain.cn/article/6068

## 攻击向量

https://learnblockchain.cn/article/4982

https://learnblockchain.cn/article/6310

## Phalcon 浏览器分析

- Explorer
- Debugger
- Simulator
- Fork

## 智能合约漏洞

https://github.com/SunWeb3Sec/DeFiVulnLabs

forge test --contracts ./src/test/Reentrancy.sol -vvvv

- 整型变量溢出 Integer Overflow

  solidity ^0.8.0 之前，需要 SafeMath 库防止整型溢出

- 合约自毁 Selfdestruct

  合约自毁可以强制向某地址转账

- 不安全的 DelegateCall

- 重入漏洞 Reentrancy

  预防：使用 检查、影响、交互模式，或者使用防止重入攻击的库合约

- 只读型重入漏洞 ReadOnlyReentrancy

  https://foresightnews.pro/article/detail/32601

- ERC777 回调和重入漏洞
  https://www.panewslab.com/zh/articledetails/D31955185.html

- 未被检查外部调用 call 注入攻击

- 私有变量

  storage 变量都是公开透明的

- ERC721 重入漏洞

  https://github.com/AmazingAng/WTF-Solidity/tree/main/S16_NFTReentrancy

- 智能合约中隐藏的后门 Hidden Backdoor

- 绕过合约长度检查

  通过检查账户字节码长度来判断是否是合约是不可行的，合约在构造函数的运行时，字节码长度为 0

- DOS 拒绝服务漏洞

- Randomness 随机性漏洞

- 函数可见性 Visibility

- tx.origin 网络钓鱼

- Incorrect use of payable.transfer()

- **Unauthorized NFT Transfer in custom ERC721 implementation**

- **Missing Check for Self-Transfer Allows Funds to be Lost**

  缺乏自转账检查

  https://github.com/code-423n4/2022-10-traderjoe-findings/issues/299

- Incorrect implementation of the recoverERC20() function in the StakingRewards

- **Missing flash loan initiator check ** 闪电贷调用检查

- Incorrect sanity checks - Multiple Unlocks Before Lock Time Elapse

## code4rean 报告

https://code4rena.com/contests

https://github.com/code-423n4/2022-10-traderjoe-findings/issues/477

## Ethernaut

**1_Fallback**

题干：（1）获得合约的所有权 （2）将余额减少为 0。可以改变 owner 的方法 contribute()和 receive()

解题思路：receive()的逻辑漏洞

**2_Fallout**

题干：获得合约所有权

SafhMath.sol 是防止整型数据溢出的计算库

解题思路：旧版构造函数名打错

**3_CoinFlip**

题干：抛硬币，预测结果连胜 10 次

解题思路:**伪随机**

硬币的结果与 block.number 有关.构建攻击合约，在攻击函数中使用相同的环境变量（同一区块），可保证预测结果相同

**4_Telephont**

题干：改变所有权

解题思路：**tx.origin 与 msg.sender 的不同**

构建攻击合约，在构造函数中进行攻击

**5_Token**

题干：获得更多的 Token

解题思路：**整型溢出**，EVM 只能表示特定范围的数字，solidity 0.8 前不会检查整型溢出

无论转账多少，合约中 balances[msg.sender]的检查都不会生效。转账超过余额，完成下溢

**6_Delegation**

题干：获得合约 Degation 的权限

解题思路：**delegatecall**

B 合约 delegatecallC 合约，执行的是 B 的环境改变，改变的也是 B 的状态变量，C 合约只做逻辑处理

**7_Force**

题干：强制向合约转账

解题思路：**合约自毁**

如果一个合约没有接收 eth 的处理函数，无法接受转账。但合约自毁可以强制向其转账

**8_Vault**

题干：解锁 vault

解题思路：**插槽访问**

私有变量仍可通过插槽访问

**9_King**

题干：变成 King，并阻止别人变成 King

解题思路：**智能合约 拒绝服务**

在攻击合约中设置，收到 eth，revert()

**10_Re-entrancy**

题干：偷取合约的所有资金

解题思路：**重入漏洞**

预防：使用 检查-影响-交互模式书写逻辑 使用重入锁

**11_Elevator**

题干：将 top 值改为 true

解题思路：

理解 Building building = Building(msg.sender); 的本质

实现 isLastFloor()方法

**12_Privacy**

题干：解锁合约

解题思路：**私有变量访问**

定长数据类型的存储

**13_GateKeeper One**

题干：过门

解题思路：**gas 计算，类型转换 **

门 1：合约攻击

门 2：remix debugger

门 3：类型转换

​ https://www.tutorialspoint.com/solidity/solidity_conversions.htm

​ https://learnblockchain.cn/docs/solidity/types.html#types-conversion-elementary-types

​

**14_GatekeeperTwo**

题干：过门

解题思路：**内联汇编**

门 1：合约攻击

门 2：extcodesize() **绕过合约长度检查 合约构造函数**

门 3：编码、异或逆运算

**15_Naught Coin**

题干：在锁定期内转移代币

解题思路：**transferFrom** ERC20

合约只限制了 transfer()方法，erc20 的代币转移还有 transferFrom()方法

**16_Preservation**

题干：获得合约的所有权

解题思路：**call 注入攻击 delegatecall 时的插槽存储冲突问题**

第一次 delegatecall，改变 timeZone1Library

第二次 delegatecall

**17_Recovery**

题干：找到丢失的合约地址

解题思路：**智能合约地址预测**

用户 A 调用智能合约 B 创建智能合约 C，创建者是智能合约 B。A 和 B 的 nonce 都会增加。目前不确定智能合约还会在哪种情况下增加 nonce

ethers.getCreateAddress({from: , nonce: })

**18_MagicNumber**

题干：

解题思路：**Opcodes**

**19_Alien Codex**

题干：获得所有权

解题思路：**动态数组的存储布局**

https://blog.dixitaditya.com/ethernaut-level-19-alien-codex

攻击数组长度，使其下溢（所有存储空间都是动态数组的范围，使得 slot[0]可以被赋值）

动态数组元素的起始存储位置:keccak256(其索引的 slot)

keccak256(1)+ X = 0 = 2^256-1 + 1

**20_Denial**

目的：不让提款成功

解题思路：**call transfer send 的区别**

攻击合约想办法耗尽 gas

**21_Shop**

目的：购买成功，并使 price 价格降低

解题思路：判断逻辑。让攻击合约的 price 和 Shop 的 isSold 挂钩

**22_Dex**

目的：

解题思路：

**23_Dex Two**

**24_Puzzle Wallet**

目的：成为代理合约的管理员

库合约：UpgradeableProxy.sol

​ Proxy.sol

解题思路：

proxy 和 impl 的 storage 存在冲突，(1)可以通过 delegatecall setMaxBalance()来修改 proxy 的 admin，需要取光合约的余额。

(2)改变 proxy 的 pendingAdmin，就能使攻击行为通过 owner 检查

**25_Motorbike**

题干：？？？

解题思路：

**26_DoubleEntryPoint**

题干：？？？

解题思路：

**27_Good Samaritan**

题干：取走钱包中所有的份额。入口函数 GoodSamaritan.requestDonation()

解题思路：利用 Coin.transfer 中的 INotifyable(dest*).notify(amount*);恶意触发 revert "NotEnoughBalance()"

如果一个合约的 mapping 类型是 public，但没有声明 getter 函数，ethers.js 该怎么获得呢？如果通过 slot 是不是太麻烦了

**28_GateKeeper Three**

题干：进入大门。gateOne()：简单

gateTwo()：allowEntrance，利用 trick.checkPassword()

解题思路：

**29_Switch**

题干：要让 switchOn = true;留下的入口，flipSwitch()

解题思路:**动态类型的 Calldata 编码**

https://blog.softbinator.com/solving-ethernaut-level-29-switch/

设计 calldata，改变偏移量，满足 onlyOff 的同时，进行 turnSwitchOn()调用

```
一般情况下动态类型的Calldata编码
0x
30c13ade	->函数选择器
0000000000000000000000000000000000000000000000000000000000000020	->偏移量
0000000000000000000000000000000000000000000000000000000000000004	->长度
20606e1500000000000000000000000000000000000000000000000000000000	->实际值
```

```
应对题目onlyOff()检查
30c13ade-> 功能选择器

0000000000000000000000000000000000000000000000000000000000000060-> 偏移量，现在 = 96 字节

0000000000000000000000000000000000000000000000000000000000000000-> 额外字节

20606e1500000000000000000000000000000000000000000000000000000000-> 这里是对字节的检查68，但与调用无关

0000000000000000000000000000000000000000000000000000000000000004-> 数据长度

76227e1200000000000000000000000000000000000000000000000000000000-> turnSwitchOn()的选择器
```

https://ethernaut.openzeppelin.com/

https://github.com/OpenZeppelin/ethernaut/blob/master/contracts/contracts/levels/Motorbike.sol

https://www.youtube.com/watch?v=MaGAVBRwvbg&list=PLiAoBT74VLnmRIPZGg4F36fH3BjQ5fLnz D-Squared

https://www.youtube.com/watch?v=TQKj2xvsGec&list=PLO5VPQH6OWdWh5ehvlkFX-H3gRObKvSL6 Smart Contract Programmer

https://dev.to/bin2chen/ethernautxi-lie-level-26doubleentrypoint-27i5 全解

https://github.com/bin2chen66/ethernaut/blob/main/contracts/26DoubleEntryPointRun.sol

## Damn Vulnerable DeFi v3

https://www.damnvulnerabledefi.xyz/

https://github.com/tinchoabbate/damn-vulnerable-defi/blob/v3.0.0/contracts/unstoppable/ReceiverUnstoppable.sol

https://github.com/Poor4ever/damn-vulnerable-defi-solution

https://www.youtube.com/watch?v=FoPMe3d4DFI&list=PLwHGiYB583YuDoAjKPDfYMKOmuFIGJCnW&index=14

https://github.com/bzpassersby/Damn-Vulnerable-Defi-V3-Solutions

**1_unstoppable**

目标：让金库停止闪电贷款

import 合约:

​ solmate/utils/FixedPointMathLib.sol 数学库

​ solmate/utils/ReentrancyGuard.sol 防重入攻击库

​ { SafeTransferLib, ERC4626, ERC20 } from "solmate/mixins/ERC4626.sol" 安全转账，**ERC4626 代币金库话**，ERC20

​ solmate/auth/Owned.sol 权限控制

​ openzeppelin-contracts/contracts/interfaces/IERC3156.sol

​ **ERC3156**

思路：**闪电贷** UnstoppableVault.flashLoan 的 flashLoan()中，if (convertToShares(totalSupply) != balanceBefore) revert InvalidBalance();

**2_Naive receiver**

目标：清空接收者合约中的 eth

import 合约：solady/src/utils/SafeTransferLib.sol 安全转账库，内联汇编省 gas

思路：**闪电贷** LenderPool 合约中 flashLoan()没有对调用者进行限制，越权调用。

**3_Truster**

目标：取出池中的所有代币

import 合约：openzeppelin-contracts/contracts/utils/Address.sol 地址库

思路：**闪电贷** **call 注入攻击**

TrusterLenderPool 的 flashLoan 方法，没有对 target 校验。利用 functionCall 进行代币授权 approve。

目标合约地址.call{value:发送数额, gas:gas 数额}(字节码);

​ abi.encodeWithSignature("函数签名", 逗号分隔的具体参数)

**4_Side Entrance**

目标：取出池子中所有的 ETH

思路：**闪电贷 余额检查**

实现 IFlashLoanEtherReceiver 的恶意 excute(),通过向池子存入代币，仍然满足检查。

**5_The Rewarder**

目标：获取池子分红

import 合约：openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Snapshot.sol

思路：**闪电贷 ERC20 快照**

**6_Selfie**

目标：取走借贷池中的全部代币

import: ISimpleGovernance.sol 治理代币合约

思路：**闪电贷 治理攻击**

提交提案、治理合约

**7_Compromised**

目标：掏空交易所

思路：**预言机操控**

控制预言机，进行低买高卖

**8_Puppet**

目标：取出借贷池中的所有代币

思路：**操控 DEX 价格预言机**

砸盘代币，去借贷池借钱

扩展：erc20-permit

**9_Puppet V2**

目标：

思路：

**10_Free Rider**

目标：

思路：

**14_Puppet V3**

在最新的版本中，使用 Uniswap v3 作为预言机，不再使用瞬时价格。池子使用推荐的库查询资产的时间加权平均价格。uniswap market 的流动性池拥有 100 WETH 和 100 DVT。借贷池有 100w DVT。以 1 ETH 和一些 DVT 开始，取出借贷池中的所有代币。**要将主网状态分叉到本地环境**

不知道 foundry 怎么为某一文件使用特定编译器

**15_ABI Smuggling**

金库有 100w DVT，允许周期性地提取资金，紧急情况下也可以拿走全部资金。合约有嵌入的通用授权方案，只允许已知的账户执行特定的提案。开发团队收到一份披露，所有的资金可能会被偷。在为时已晚之前，拯救所有的资金，把它们转移到恢复账户上。

目标：转移资金

思路：

## 漏洞检测工具

https://www.youtube.com/watch?v=IJ5VwEZvbbc

https://learnblockchain.cn/people/9625

# 开发工具:Foundry Hardaht

```powershell
> forge init xxx
├── foundry.toml        # Foundry 的 package 配置文件
├── lib                 # Foundry 的依赖库
│   └── forge-std       # 工具 forge 的基础依赖
├── script              # Foundry 的脚本
│   └── Counter.s.sol   # 示例合约 Counter 的脚本
├── src                 # 智能合约的业务逻辑、源代码将会放在这里
│   └── Counter.sol     # 示例合约
└── test                # 测试用例目录
    └── Counter.t.sol   # 示例合约的测试用例
```

## Forge 测试、构建、部署智能合约

https://www.bilibili.com/video/BV1AG4y167Qv/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551 Part 1

https://www.bilibili.com/video/BV1ne4y1d7VG/?spm_id_from=333.788.recommend_more_video.-1&vd_source=5c46b11c8d760605427b8431fb93d551 Part 2

### 测试

通过传递过滤器运行特定测试

forge test --match-contract ComplicatedContractTest --match-test testDeposit

- setUP() 在每个测试用例运行前调用的可选函数

- test 前缀 测试用例

- testFail 前缀 （不建议）

- testRevert()

  vm.expectRevert()

#### 了解 Traces

#### 分叉测试

分叉模式

forge test --fork-url <your_rpc_url>

forge test --fork-url=%ETH_RPC_URL% --- windows 写法

分叉作弊码

string memory rpc = vm.envString("ETH_RPC_URL");

uint256 mainnet = vm.createFork(rpc);

vm.selectFork(mainnet);

### 高级测试

#### 模糊测试 Fuzz Testing

#### 差异测试 differential testing

ffi

### 部署和验证

### 脚本 forge script

#### 本地网络

vm.startBroadcast();

Counter c = new Counter();

vm.stopBroadcast();

forge script script/Counter.s.sol -vvvv --rpc-url= --broadcast --private-key --sender

#### fork 网络

分叉测试中的分叉作弊码

### Gas 追踪

Gas 报告

Gas 快照

### 调试器 Debugger

evm debug

forge script <PATH> --debug

cast run <HASH> --debug

### Forge 命令

常用命令

Project 命令

#### Build 命令

forge build

forge clean

forge inspect

​ forge inspect XXX mi

​ forge inspect XXX slot

#### Test 命令

forge test -运行项目的测试

forge test --gas-report

forge snapshot （Gas 快照）

​ --diff

#### 作弊码 vm

##### Environment

.warm 设置区块时间戳 block.timestamp

.prank 指定下一次调用

.startPrank 指定接下来的调用

.stopPrank 停止

.startBroadcast

##### Assertions

.expectRevert() 怎么捕捉指定的错误？

##### External

ffi 外部终端命令

##### Utilities

label 调用栈显示别名，代替地址

## Cast 与链交互

--rpc-url = https://eth-mainnet.g.alchemy.com/v2/PREW0A8nzk9u-It_3kZ7OVb-inBX8xgr

--rpc-url =https://eth-goerli.g.alchemy.com/v2/W5WCfj_DebWW4-5OzyeG8oKxQMSgfGJm

**常用命令**

**Chain 命令**

**Transaction 命令**

​ cast estimate

**ABI 命令**

​ cast abi-encode 对调用的参数进行 ABI 编码

​ cast 4byte 0x8cc5ce99 获取选择器的函数签名(数据库)

​ cast 4byte-decode 对调用的 calldata 进行解码

​ cast calldata 编码带参数的函数

​ cast pretty-calldata 分割 calldata

**Utility 命令**

​ cast sig 函数签名

​ cast keccak

## Anvil 创建本地测试网节点

## 参考命令

**Forge Standard Library**

Std Assertions

Console Logging

​ console.logBytes32()

**DSTest**

assertTrue

## Hardhat

Hardhat 框架以及部署、测试架构

npm init --yes

npm install hardhat --save -dev

npx hardhat

- ##### 3.安装 openzeppelin 等库

```shell
npm install --save-dev @nomicfoundation/hardhat-toolbox @openzeppelin/contracts @openzeppelin/hardhat-upgrades keccak256 merkletreejs dotenv bignumber ts-node typescript @types/node @types/mocha @types/chai
```

- ##### 4.在 hardhat.config.js 配置文件中添加 hardhat-toolbox 全局配置

```bash
//` `因为hardhat框架全局都需要使用hardhat-toolbox，所以在这里引入
require(``"@nomicfoundation/hardhat-toolbox"``);

module.exports = {
 ``solidity: ``"0.8.18"
};
```

- ##### 5. 编译

```shell
npx hardhat compile
```

- ##### 6. 测试

```shell
npx hardhat test
```

7.部署

部署到本地网络

```shell
npx hardhat node
npx hardhat run scripts/deploy.ts --network localhost
```

部署到测试网

```shell
npx hardhat run .\scripts\deploy.js --network goerli
```

验证

```shell
npx hardhat verify --network goerli contract-address [para1] [para2]
npx hardhat verify 0xD3F4690679d89057359B3bEB1441ee7E837e6dCc 1689525603 --network goerli
```

More than one contract was found to match the deployed bytecode.

```shell
npx hardhat verify --contract contracts/tokens/WorldCupToken.sol:WorldCupToken 0xB500D338D6cd608D867015295483E38758CC7711 "World Cup Token" "WCT" "10000000000000000000000000" --network goerli
```

https://github.com/1500256797/hardhat-demo-project/tree/master 欧皇张大千

### 目录结构

-contracts:合约

-scripts:脚本

-test:单元测试

-hardhat.config.ts:配置文件

-package.json:包管理文件

### 部署/测试脚本架构

import "hardhat/console.log" 可以在智能合约中，console.log()

同一个 describe 中不同的 it 互相独立，但对主环境的影响会生效

### 修改配置文件

1.修改.env

2.安装 dotenv

​ 在配置文件中引用

​ require('dotenv').config()

3.配置文件完整内容

# DeFi

https://learnblockchain.cn/column/4

了解架构、经济模型、白皮书部分的代码实现

## UniSwap V2

x \* y = K

交易前的瞬时价格就是斜率 即 y/x

**2.2 价格预言机**

​ https://zhuanlan.zhihu.com/p/359750357 如何使用 Uniswap v2 作为预言机

​ 在每个区块的第一笔交易前记录**累计价格**实现预言机，时间权重记录

​ TWAP 算术平均数

​ 两个难题：1. 应该计算以 B 代币计价的 A 代币价格，还是以 A 代币计价的 B 代币价格？ 2.用户可以不通过交易而直接向交易对合约发送代币（这将改变代币余额并影响价格），此时将无法触发预言机价格更新

​

​ 价格精度：UQ112.112 格式的数据，它表示在小数点的左右两边都有 112 位比特表示精度 **Solidity 原生不支持非整数数据类型**

​ 112+112+32 = 256

​

**2.3 闪电贷**

**2.4 Protocol fee** 协议手续费

**2.5 Meta transactions for pool shares** 池子份额元交易

**3.4 初始化流动性代币供应**

**3.5 Wrapping ETH - WETH**

​

**3.6 确定的交易对地址 Wrapping ETH - WETH**

**3.7 **

### 合约代码

#### 1.合约架构

core 合约 core 合约最简化
Factory：负责创建交易对，保存所有交易对的地址
Pair：保存单个交易对的资金池信息。定义和交易有关的几个最基础方法，如:swap/mint/burn，价格预言机等功能
ERC20：方便用户直接用 ETH 和 ERC20 进行交易
periphery 合约
Router02：实现常用的接口，比如添加/移除流动性，使用代币 A 交换代币 B

#### 2.core

##### 2.1 Factory 合约

**createPair 方法** 创建交易对，create 2 创建确定性的交易对合约地址

##### 2.2 ERC20 合约

流动性代币也是一个 ERC20 合约，并且支持 permit

**UniswapV2ERC20.permint()**

##### 2.3 Pair 合约

function mint()

function burn()

**swap() **

兼容闪电贷，先转出代币

​

#### 3.periphery

##### 3.1 Library

**pairFor()**

计算两个代币的交易对地址

##### quote

(amountA/reserveA) = (amountB/reserveB)

##### getAmountOut

amountIn A = ? amountOut B

##### getAmountIn

? amountIn A = amountOut B

##### getAmountsOut

多个交易对，循环调用 getAmountOut

##### getAmountsIn

多个交易对，循环调用 getAmountIn

##### 3.2 UniswapV2Router02

#####

##### 4.预言机

https://learnblockchain.cn/article/3960

TWAP(Time-Weighted Average Price) 时间加权平均价格
累计价格：UniswapV2Pair.sol 中，price0CumulativeLast 和 price1CumulativeLast
在\_update()函数中更新这两个变量

固定时间窗口 TWAP

滑动时间窗口 TWAP

需要链下程序定时触发 update()函数

## Uniswap V3 tick?

https://www.bilibili.com/video/BV1hY4y1D7aD/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551

代码讲解视频

### 白皮书

#### 1 介绍

仍然基于常值函数曲线（x \* y = k）
新特性：
集中流动性
灵活的手续费
协议手续费治理
改进的价格预言机
流动性预言机

#### 2 Concentrated Liquidity 集中流动性

v3 有限区间的流动性类似于 v2 池子的一部分，兑换率离开这个区间时，这个区间的池子将完全只由一种代币组成，另一种代币会被耗尽，该池子也不再活跃。

v3 曲线是将 v2 曲线沿横坐标和纵坐标平移一定单位

![image-20230805221746209](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230805221746209.png)

#### 3 Architectural Changes 架构变动

**3.1 Multiple Pools Per Pair 多池交易对**

​ 每个交易对有多个池子，分别设置不同的交易手续费。

**3.2 Non-Fungible Liquidity 不可互换的流动性**

​ 手续费不像 v2 是复利的，v3 的流动性是一个 NFT

#### 4 Governance 治理

#### 5 Oracle Upgrades 预言机升级

v3 对 v2 的时间加权平均价格(TWAP)预言机进行了改动

https://learnblockchain.cn/article/4179

**5.1 Oracle Observations 预言机观测**

**5.2 Geometric Mean Price Oracle 几何平均数价格预言机**

**5.3 Liquidity Oracle 流动性预言机**

#### 6 Implementing Concentrated Liquidity 实现集中流动性

**6.1 Ticks and Ranges 点和区间**

p(i) = 1.0001^i
交易对池子实际使用开根号价格 sqrt(price)

**6.2 Global State 全局状态**

7 个与交换和流动性供应相关的存储变量
*liquidity
*sqrtPriceX96
*tick
*feeGrowthGlobal0X128
*feeGrowthGlobal1X128
*protocolFees.token0
\*protocolFees.token1

**6.2.1 Price and Liquidity 价格和流动性**

core

##### UniswapV3Factory.sol

*createPool：创建交易对池子
*setOwner：设置工厂合约 Owner
\*enableFeeAmount：添加手续费等级

#### createPool 方法

createPool(tokenA,tokenB,fee);
pool = deploy(address(this),token0,token1,fee,tickSpacing);

##### deploy in UniswapV3PoolDeployer.sol

    create2
    salt:token0,token1,fee

#### setOwner 方法

owner 具有以下权限
*setOwner：修改 owner
*enableFeeAmount：添加手续费等级
*setFeeProtocol：修改某个交易对的协议手续费比例
*collectProtocol：收集某个交易对的协议手续费

#### enableFeeAmount

UniswapV3Pool.sol

定义交易对池子的功能
*initialize：初始化交易对
*mint：添加流动性
*burn：移除流动性
*swap：交换代币
*flash：闪电贷
*collect：取回代币
*increaseObservationCardinalityNext：扩展预言机空间
*observe：获取预言机数据
此外，factory（工厂合约）owner 还可以调用以下两个方法：
*setFeeProtocol：修改某个交易对的协议手续费比例
*collectProtocol：收集某个交易对的协议手续费

#### initialize 方法

初始化 slot0 变量，交易池的参数和状态的数据结构
*sqrtPriceX96:交易对当前的开根号价格
*当前对应的 trick，使用 getTickAtSqrtRatio 计算得出

#### mint 方法

添加流动性
参数：
*recipient：头 寸接收者（owner）
*tickLower：流动性区间低点
*tickUpper：流动性区间高点
*amount：流动性数量
\*data：回调参数

##### \_modifyPostion

##### \_updatePostion

#### burn 方法

#### swap 方法

periphery

结合前端的过程： 1.代币授权 approve(address spender,uint256 tokens) 2.添加流动性 multicall(bytes[] data)

NonfungiblePositionManager.sol 头寸管理合约

*createAndInitializePoolIfNecessary：创建并初始化合约
*mint：创建头寸
*increaseLiquidity：添加流动性
*decreaseLiquidity：减少流动性
*burn：销毁头寸
*collect：取回代币

#### createAndInitializePoolIfNecessary in PoolInitializer.sol

首先根据交易对代币(token0 和 token1)和手续费 fee 获取 pool 对象
如果不存在，则调用 Uniswap-v3-core 工厂合约 createPool 创建该交易对并初始化
如果已存在，则根据 slot0 判断是否已经初始化(价格)

## Uniswap V4

坚持

​ 恒定积公式 x \* y = k

​ 限价流动性基于 tick 刻度体系

​ 流动性堆叠

革新

​ 全新的合约架构

​ 巧妙地账本设计

​ 更高的开发者自由度

​ 对流动性提供者更多的保护

https://learnblockchain.cn/article/6006

https://www.bilibili.com/video/BV1KF411R7fD/?spm_id_from=333.788&vd_source=5c46b11c8d760605427b8431fb93d551

## 借贷

https://learnblockchain.cn/article/5684

APR 稳定年回报率

APY 复利年收益率

### 借与贷

借贷池

份额币

**aToken** liquidityIndex 贴现因子 将**单利利率**按秒复利（实际操作中，不一样）

![image-20230808011919502](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230808011919502.png)

​

**cToken** exchangeRate 兑换率

​

###

## AAVE

#### 白皮书

https://www.bilibili.com/video/BV1UF411Y7oU/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551 白皮书

**3.1 aTokens**

**3.2 Debt Tokenization 债务 Token 化**

**3.3 Variable Debt 可变债务**

**3.4 Stable Debt**

**3 .5 Flash Loans V2 闪电贷**

#### 源码

5 个核心方法

- deposit

deposit()

mint()相应的 token

amountToMint = ERC20_tokensTransferred / liquidityIndex

- withdraw

- borrow
- repay
- flashloan

https://www.bilibili.com/video/BV1qR4y177Uk/?spm_id_from=333.999.0.0&vd_source=5c46b11c8d760605427b8431fb93d551 代码解读

经济模型

aToken

债务 Token

!(C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230802000817997.png)

借款利率

浮动利率计算

![image-20230801235603854](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230801235603854.png)

稳定利率计算(stable) **难 v2 v3 不同**

健康度

![image-20230801232312001](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230801232312001.png)

Σ 资产 \* 清算阈值 / 债务

清算

v3 新东西

isolation mode

E-mode 高效模式，同类资产更高的

## Compound

### 快速了解

超额抵押贷款
聚合借贷
标的资产（Underlying Token）
cToken
用户在 Compound 上存入资产的凭证，每一种标的资产都有对应的一种 cToken，凭此可以换回质押资产的本金和收益
兑换率（Exchange Rate）
抵押因子（Interest Rate Model）
利率模型（Interset Rate Model）
储备金率（Rerseve Factor）
资金使用率（Utilization Rate）
清算（Liquidation）
价格预言机（Price Oracle）

主要模块：
1.CToken
抽象基类合约
CEther

2.InterestRateModel
3.Comptroller
4.priceOracle

### 1.CToken.sol

exchange rate 兑换率 = AAVE v2 的流动性指数

![image-20230808005759746](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230808005759746.png)

AAVE 是

#### 主要事件

#### 错误代码

#### 失败信息

#### 汇率(Exchange Rate)

exchangeRate = (getCash() + totalBorrows() - totalReserves()) / totalSupply()

function exchangeRateCurrent() returns (uint){
accrueInterest();
return exchangeRateStored();
}
function exchangeRateStored() override public view returns (uint){
return exchangeRateStoredInternal();
}
function exchangeRateStoredInternal() virtual internal view returns (uint){

}

#### 获取现金(Get Cash)

#### 总借款(Toal Borrwos)

#### 借款余额(Borrow Balance)

CToken.sol
x`
function borrowBalanceCurrent(address account) override external nonReentrant returns (uint)
function borrowBalanceStored(address account) override public view returns (uint)
function borrowBalanceStoredInternal(address account) internal view returns (uint){

}

#### 借款利率(Borrow Rate)

当前每个区块的借款利率
function borrowRatePerBlock() returns (uint)

#### 总供应量(Total Supply)

CToken 流动的代币数量

#### 标的余额(Balance)

#### 供给率(Supply Rate)

#### 总储备金(Total Reserves)

#### 储备金系数(Reserve Factor)

准备金系数确定了借款人利息中转化为准备金的部分。

### 2.利率模型

利率模型的抽象合约
InterestRateModel.sol
function getBorrowRate(uint cash, uint borrows, uint reserves) virtual external view returns (uint);
function getSupplyRate(uint cash, uint borrows, uint reserves, uint reserveFactorMantissa) virtual external view returns (uint);
具体实现：
直线型：WhitePaperInterestRateModel
拐点型：JumpRateModelV2
本质上是反映借贷供求

### WhitePaperInterestRateModel.sol 直线型

y = k \* x + b
利率乘数 资金使用率 基准利率

constructor(){
块基准利率 = 年基准利率 / 年块数
块利率乘数 = 年利率乘数 / 年块数
}

function utilizationRate(){
资金使用率 = 借出代币总数/（代币余额 + 借出代币总数 - 储备代币总数）
}

function getBorrowRate(){
块借出利率 = 资金借出率 \* 块利率乘数 + 块基准利率
}

function getSupplyRate(){
块质押利率 = 资金借出率 _ 借款利率 _ （1-储备金率）
一部分收入要分给储备金
}

### BaseJumpRateModelV2.sol 拐点型

y = k _ x + b
y = k2 _ (x - p) + (k \* p) + b

### 3.Comptroller 审计合约

### 4.预言机

## Synthetix

## 闪电贷

https://learnblockchain.cn/article/4500

智能合约的**原子性**

在一个区块交易中，同时完成借款和还款操作这两个操作，无需抵押任何资产，只需支付手续费即可。用户在借到款之后，可以利用借到的资产进行其他操作，比如套利、偿还抵押借款、自清算等操作。在交易结束的时候，用户只要把借到的款项及手续费及时归还就可以，否则该笔交易就会回滚，就像什么也没有发生过一样。

**Optimistic Transfer **乐观转账

# EIP 提案

https://ethereum.org/en/developers/docs/

### Gas EIP1599

https://learnblockchain.cn/article/5012 一文读懂以太坊签名：ECDSA、RLP、EIP155、EIP191、EIP712

![yuque_diagram.jpg](https://img.learnblockchain.cn/attachments/2022/11/KYX4eZRJ636a2292a94e7.jpg)

eip-155

重放攻击保护，防止在一个以太坊链上的交易被重复广播到另外一条链

eip-191

https://github.com/AmazingAng/WTF-Solidity/blob/main/37_Signature/readme.md

![image-20221127223259928](https://duke-typora.s3.ap-southeast-1.amazonaws.com/uPic/image-20221127223259928.png)

**eip-712** 类型化数据签名

https://blog.wssh.trade/posts/ecsda-sign-chain/#%E6%A6%82%E8%BF%B0

### **账户抽象 ERC-4337 **

https://www.bilibili.com/video/BV1NM4y1s77B/?spm_id_from=333.337.top_right_bar_window_view_later.content.click&vd_source=5c46b11c8d760605427b8431fb93d551 账户抽象与 ERC-4337

https://www.bilibili.com/video/BV1xs4y1i7Js/?spm_id_from=333.337.top_right_bar_window_view_later.content.click&vd_source=5c46b11c8d760605427b8431fb93d551 以太坊账号抽象之路：4337 源码解析

https://www.youtube.com/watch?v=onMm63DkTFY&t=29s

账户抽象：将 EOA 钱包变成可编程、图灵完备、同时可以发起交易的全能型钱包

抽象：将复杂的运行逻辑抽象，封装，提高用户体验，降低门槛

EIP-4337 方案 不需要改变以太坊底层原生协议

![image-20230804171204873](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230804171204873.png)

ERC4337 在 EOA 架构的基础上，头尾都增加了新的模块

![image-20230804172544680](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230804172544680.png)

![image-20230804173139221](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230804173139221.png)

![image-20230804173334679](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230804173334679.png)

### 坎昆升级 和 EIP-4844

https://www.bilibili.com/video/BV1Ut4y1579d/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551 深入 Data Availability: EIP-4844, Danksharding

https://www.bilibili.com/video/BV1Lm4y1b7Lh/?spm_id_from=333.1007.top_right_bar_window_view_later.content.click&vd_source=5c46b11c8d760605427b8431fb93d551 拆解坎昆升级

https://www.youtube.com/watch?v=9JSkHe6MABc&t=1214s

EIP-4844 Proto-Danksharding

背景：

二层网络扩容的瓶颈：将数据锚定回以太坊执行层

三大解决思路：

压缩数据体积

将数据可用性模块化

**原生协议内优化存储成本**

​ **Proto-Danksharding(EIP-4844)**

​ 提议者-构建者分离(PBS)

​ 完整分片（Danksharding）

# Openzepplin/solmate 库

https://learnblockchain.cn/people/14159

- ## **ERC20**

https://learnblockchain.cn/article/4709

https://learnblockchain.cn/article/4132 ERC20/ERC721/ERC1155 概念与协议

https://learnblockchain.cn/article/4153

ERC20-Permit(eip-2612)

**基本业务**：查询、转账、授权

- ## **ERC721**

**ERC721A**

https://blog.wssh.trade/posts/erc721a-contract/

**ERC777**

​ https://learnblockchain.cn/2019/09/27/erc777

​ ERC 1820

**ERC-3156 闪电贷**

ERC3156FlashLender

​ function maxFlashLoan(address token) external view returns (uint256);

​ function flashFee(address token, uint256 amount) external view returns (uint256);

​ function flashLoan(

​ IERC3156FlashBorrower receiver,

​ address token,

​ uint256 amount,

​ bytes calldata data

) external returns (bool);

ERC3156FlashBorrower

function onFlashLoan(

​ address initiator,

​ address token,

​ uint256 amount,

​ uint256 fee,

​ bytes calldata data

) external returns (bytes32);

**ERC-4626 代币金库化**

asset share

ERC165

NFT 开发：ERC-721

**ERC721**

EIP-721 中定义了两个 safeTransferFrom，

transferFrom 与 ERC20 中的逻辑不同

4.实现 EIP-721 中定义的 transfer 方法

5.实现 EIP-721 中定义的 set 方法

function approve()

function setApproveForAll()

为什么 经销商不能是自己

6.其他的辅助方法 mint,burn

ERC721Enumerable.sol

核心 ERC721 标准+Metadata+Enumerable 扩展

**preMint mint**

**白名单**

1.mapping 存白名单

2.MerkleTree 实现

3.数字签名实现

**Metadata，nft 盲盒，开图**

NFT 盲盒实现方案

1.直接设置

tokenURI = baseURI + tokenId

2.随机数+洗牌算法

**随机铸造**,防范 NFT 狙击手

1.暴露了代币元数据(让狙击手推断出代币的稀有性)

2.以确定的顺序铸造代币(让狙击手推断稀有代币的正确时间)

难点：链上随机化

使用链下预言机(Chainlink)

function request()

function fulfill()

**动态更新**

**合约自动化**

ERC721A

**# ERC721A**

**## openzeppelin 的 ERC721 的\_safeMint**

mint 方法内部，维护了两个全局的 mapping

balance

owners

没有提供批量 Mint 的 API，需要循环调用 N 次单独的 mint 方法

**## ERC721A**

如果某一段连续的 tokenID 都被同一个用户所拥有，那么只在第一个位置上记录相关信息

查询：向前寻找数据非空的 ID

transfer:将当前 ID 的后一个 ID 信息设置为转出人的信息

Azuki 源码

- 主体代码

  ![image-20230725160946654](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230725160946654.png)

- 依赖代码

  ![image-20230725161006641](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230725161006641.png)

EIP-1167 Minimal Proxy Contract

EIP-1167 提供了一种低成本克隆合约的方法

![image-20230721142059214](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230721142059214.png)

Openzepplin

utils

cryptography

**MerkleProof**

**ECDSA**

Ownable.sol

onlyOwner()修改器用来限制某些特定合约函数的访问权限

ReentracyGuard.sol

nonReentrant()修改器，重入锁

## delegatecall 代理合约 可升级合约

delegatecall 是一种合约间的调用方式

https://github.com/AmazingAng/WTF-Solidity/tree/main/46_ProxyContract

Proxy.sol 在内联汇编中使用 delegatecall，让回调函数也能有返回值

# 内联汇编

https://learnblockchain.cn/article/6064

https://blog.csdn.net/weixin_62775913/article/details/125828044

```csharp
calldatacopy(uint destOffset, uint sourceOffset, uint length)
```

# Transaction Debugging tools

Invocation Flow 可视化函式调用流程

**Phalcon** Ethereum、BSC、Cronos、Avalanche C-Chain、Polygon

**Tx.viewer ** Ethereum、Polygon、BSC、Avalanche C-Chain、Fantom、Arbitrum、Optimism

**Ethtx** Ethereum、**Goerli testnet**

**Tenderly**

# 群友讨论的技术问题

# 中间件

## 价格预言机

https://learnblockchain.cn/article/3901

价格预言机使用总结

https://docs.chain.link/data-feeds/price-feeds/addresses

## 合约自动化执行工具

1.ChainLink Automation

https://learnblockchain.cn/article/5618

按时间触发：

https://automation.chain.link/mumbai/41598017923568840950672097930953463617940661135945930940240498694060567780091

- checkUpkeep
- performUpkeep

# Prediction market 开发

https://github.com/pancakeswap pancakeswap github

https://docs.pancakeswap.finance/products/prediction 文档

https://docs.pancakeswap.finance/products/prediction/prediction-faq 常见问题

https://pancakeswap.finance/prediction 主页

https://bscscan.com/address/0x18b2a687610328590bc8f2e5fedde3b582a49cda#code Prediction v2 BNB 合约

https://bscscan.com/address/0x0E3A8078EDD2021dadcdE733C6b4a86E51EE8f07#code Prediction v3 CAKE 合约

https://docs.zapankiswap.io/ ZapanikiSwap 白皮书

**chainlink 预言机喂价**

https://bscscan.com/address/0xb6064ed41d4f67e353768aa239ca86f4f73665a1#code chainLink 喂价合约

file:///C:/Users/Administrator/Desktop/%E7%99%BB%E9%93%BE%E8%AF%BE%E4%BB%B6/w6-1.pdf

登链 6.3

interface AggregatorV3Interface {

​

}

**合约自动化调用工具**

operatorAddress 0x1BB9A6E9685b49B272a5C1c8c96a79b7147Bc87f 私钥钱包

keeperAddress 0xe1d9C27a2e27ffca8c4D50A2334C7f5094215fBD

PredictionKeeper.sol 0xe1d9C27a2e27ffca8c4D50A2334C7f5094215fBD

- checkUpKeep()
- performUpKeep()

通过外部账户 调用 KeeperRegistry 再调用 PredictionKeeper 合约(ChainLink Automation)，这个交易不会在 Transactions 看到，属于 Internal transactions，只能通过释放 event 被看到。

https://explorer.phalcon.xyz/tx/bsc/0x13762e9905944c254410e8cfa9327522a086a9a63127af874a9e29b640f64269

![image-20230808181820654](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230808181820654.png)

V2-BNB

一个 EOA 账户，定时调用 executeRound()

经常发现触发失败的情况

![image-20230809152741889](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230809152741889.png)

## 支持 btc eth matic 的预测

最好去新链上。在 mumbai 测试链上

价格预言机：https://docs.chain.link/data-feeds/price-feeds/addresses

MATIC 预言机：0xd0D5e3DB44DE05E9F294BB0a3bEEaF030DE24Ada

https://github.com/pancakeswap/pancake-smart-contracts/blob/master/projects/predictions/v3/scripts/deploy.ts 部署脚本

https://github.com/pancakeswap/pancake-smart-contracts/blob/master/projects/predictions/v3/config.ts

Demo：

https://mumbai.polygonscan.com/address/0x9bC2523031B11ed566F9499e25Ff91E04448eB4F#writeContract

Automation

https://automation.chain.link/mumbai/13424986441662705191159788461405343377081488869623518338894046119832994116154

https://mumbai.polygonscan.com/address/0xb8A56CDe91904BE27D7dbD47E56dEdd6a067b9C6#readContract

#### 动态 NFT

库里

5981

https://m.rstrstrst.com/api/metadata/42/9142

# Telegram bot

t.me/bucypredictionbot

HTTP API

6429589658:AAH2CzjvoW78Bmw8rOdpQhaOkBwDqdQ6osY

**BotFather 领机器人**

# Chatgpt 提示词工程

读长代码，分段记录

1.**定界符**

三重反引号，“”

2.**请求结构化输出**

3.**要求模型检查任务条件是否满足**

# 面经

https://juejin.cn/post/7112504255432359949

# 英语

periodically 定期地 embed 嵌入 generic 通用的 authorization 授权 scheme 方案 disclosure 披露 rescue 拯救

genesis 起源 bull 牛(up) bear 熊(down) paused adj.暂停 pause v.暂停

square 正方形 router 路由

# Node.js

https://www.bilibili.com/video/BV1PV411d7XV/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551

## npm 包管理工具

npm init

npm install xxx

npm unintall xxx

mpm list

package.json 是 npm init 时生成的

作用：描述项目以及项目所依赖的模块信息

package-lock.json 是运行 npm install 时生成的一份文件

记录当前状态下实际安装的各个 npm 包的具体来源和版本号

本地安装

全局安装

开启 ES6 模块

package.json 添加 "type":"module"

## Yarn 包管理工具

## nodejs 与数据库交互

npm install mysql

**交互的基本模式**

导入库

定义 connection

connection.connect()

connection.query()

connection.end()

**封装执行 sql 语句函数**

https://www.bilibili.com/video/BV1KX4y1K7uz?p=15&vd_source=5c46b11c8d760605427b8431fb93d551

# 数据库管理系统

文档数据库

关系型数据库

123456

## MySQL

https://www.bilibili.com/video/BV16D4y167TT/?spm_id_from=333.788.recommend_more_video.2&vd_source=5c46b11c8d760605427b8431fb93d551

CREATE TABLE accounts_record(

​ chatid INT PRIMARY KEY,

​ address VARCHAR(42)

​ privateKey VARCHAR(64)

);

**增**

INSERT INTO usercount.counts_record(id,address,privatekey)

VALUES(99999,'0xF51585A11f74A45381c7FcaEbc03C102f96971B4','77e7e7d4db6930c32590eb803d74265679c64212de31fbdb178eeb377f0c2525');

**删**

删除表

DROP TABLE table_name;

DELETE FROM usercount.counts_record

WHERE id = 1;

**改**

**查**

# 前端

## html

https://www.w3school.com.cn/tags/tag_tr.asp

<> </>标签

**超链接标签** <a></a>

**article 标签** <article></article>

**换行标签** <br />

**按钮标签** <button></button>>

**表单标签** <form></form>>

**标题标签** <h1></h1>

div 容器 #div

**图片标签** <img src = "" alt= "">

**区段标签** <section></section>>

**定义有序列表** <ol></ol>

**定义列表项目** <li><li>

## css

## JavaScript

脚本语言，动态执行

https://developer.mozilla.org/zh-CN/docs/web/javascript/reference/statements/export

**导入**

默认导入全部

const hre = require("hardhat")

命名式导入

const { expect } = require("chai")

import ES6 标准

导入全部

import \* as myModule from 'myModule';

导入指定内容

import {myFunction} from 'myModule';

**导出**

export default function () {}

**map**

map()方法创建一个新数组，新数组由原数组的每个元素都调用一次提供的函数后的返回值组成

**解构赋值**

**异常处理**

```javascript
try {
  // 如果出现异常，只会执行这一个语句
  tx = await callContract.buybull(
    currentEpoch,
    hashtable.get(ctx.from.id).privatekey,
    amount
  )
} catch (error) {
  console.log('出错了，能捕捉到吗？')
  console.log(error)
} finally {
  console.log(tx)
  await ctx.reply(`${tx}`)
}
```

## TypeScript

**开发环境**

@types/node 类型定义

调试：ts-node xxx.ts

编译：tsc xxx.ts

配置文件 tsconfig.json：tsc --init

**类型**

number

string

boolean

数组 type[]

元组 Tuple [type1,type2]

枚举 enum

any

unknown

never

**接口**

定义共性

interface 和 抽象类的区别

​ 接口不能有实现

**泛型**

**类**

类

```typescript
class Account {
  id: number
  owner: string
  balance: number

  constructor(id: number, owner: string, balance: number) {
    this.id = id
    this.owner = owner
    this.balance = balance
  }

  deposit(amount: number): void {
    if (amount <= 0) throw new Error('Amount must be greater than 0')

    this.balance += amount
  }
}
```

对象

```typescript
let account = new Account(1, 'Mosh', 0)
```

只读和可选属性

访问控制关键字

public

priviate

protected

参数属性

```typescript
class Account {
  constuctor(
    public readonly id: number,
    public owner: string,
    private _balance: number
  ) {}
}
```

存取器 Getter and Setters

静态成员

属性存在于类上，而不是实例

```typescript
class Ride {
  static activeRides: number = 0
  start() {
    Ride.activeRides++
  }
  stop() {
    Ride.activeRides--
  }
}
```

继承

方法重写

​ override

多态

抽象类和方法

**函数**

函数的返回值类型

、

可选参数，默认参数

**高级类型**

类型别名

type Xxxx = {

}

联合类型

交叉类型

字面类型

**泛型**

**装饰器** decoratos

**模块导入导出**

import

export

## React

https://www.yuque.com/fechaichai/qeamqf/xbai87 前端教书匠文档

https://zh-hans.react.dev/learn 官方简易文档

https://www.bilibili.com/video/BV1be411w7iF/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551 React 快速入门 蛋老师

### React 哲学

分解组件

**将 UI 拆解为组件层级结构**

- FilterableProductTable
  - SearchBar
  - ProductTable
    - ProductCategoryRow
    - ProductRow

**使用 React 构建一个静态版本**

简单例子：自上而下构建

复杂例子：自下而上构建更简单

```react

```

**找出 UI 精简且完整的 state 表示**

**验证 state 应该被放置在哪里**

**添加反向数据流**

### 1.描述用户界面

https://zh-hans.react.dev/learn/describing-the-ui

**第一个组件**

```react
export default function Gallery() {
  return (
    <section>
      <h1>了不起的科学家</h1>
      <profile />
      <profile />
      <profile />
    </section>
  );
}

// ✅ 在顶层声明组件
function Profile() {
  return (
    <img
      src="https://i.imgur.com/QIrZWGIs.jpg"
      alt="Alan L. Hart"
    />
  );
}
```

**组件的导入与导出**

javascript 的导入与导出 **默认导出** vs **具名导出**

**使用 JSX 书写标签语言**

1.只能返回一个根元素 <div></div>

2.标签必须闭合

3.使用驼峰式命名法给大部分属性命名

**在 JSX 中通过大括号{ }使用 JavaScript**

**Props** 属性

React 组件使用 _props_ 来互相通信。每个父组件都可以提供 props 给它的子组件

```react
export default function Profile() {
  return (
    <Avatar
      person={{ name: 'Lin Lanying', imageId: '1bX5QH6' }}
      size={100}
    />
  );
}

function Avatar({ person, size }) {
  // 在这里 person 和 size 是可访问的
}
```

**条件渲染**

合理运用 三目运算符，与运算符

**渲染列表**

JavaScript 数组方法 filter() map()

**保持组件纯粹**

### 2.添加交互

**响应事件**

事件处理函数

**State：组件的记忆**

```react
import { useState } from 'react';
import { sculptureList } from './data.js';

export default function Gallery() {
  const [index, setIndex] = useState(0);

  function handleClick() {
    setIndex(index + 1);
  }

  let sculpture = sculptureList[index];
  return (
    <>
      <button onClick={handleClick}>
        Next
      </button>
      <h2>
        <i>{sculpture.name} </i>
        by {sculpture.artist}
      </h2>
      <h3>
        ({index + 1} of {sculptureList.length})
      </h3>
      <img
        src={sculpture.url}
        alt={sculpture.alt}
      />
      <p>
        {sculpture.description}
      </p>
    </>
  );
}

```

可以拥有多个 state 变量

State 是组件私有的

**渲染和提交**

三步骤：

- 1.触发一次渲染

组件的初次渲染

组件（或其祖先之一）的状态发生了改变

- 2.渲染组件

初次渲染时，React 会调用根组件

对于后续的渲染，React 会调用内部状态更新触发了渲染的函数组件

- 3.提交到 DOM（前端的 DOM 到底是什么）

- 尾声：浏览器绘制

**state 如同一张快照**

## Nextjs 全栈开发框架

https://learnblockchain.cn/article/4624

https://www.bilibili.com/video/BV1K54y1c7AY/?spm_id_from=333.788&vd_source=5c46b11c8d760605427b8431fb93d551 全栈码叔

https://www.bilibili.com/video/BV13M4y1t78b/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551 波波微课

![image-20230813220427856](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230813220427856.png)

npx create-next-app@12 hello-world 创建

yarn dev 运行

基于文件的路由

预渲染（Pre-rendering）

API 路由

CSS modules

认证

用于 Dev/Prod 的构建系统

### 路由 routing

基于文件系统的路由方式

**基于 pages 的路由**

**嵌套路由** 预定义

在 pages 目录中添加嵌套目录

blog/index.js

blog/first.js

blog/second.js

**动态路由**

[bookId].js

```react
import { useRouter } from "next/router"

function BookDetail () {
  const router = useRouter()
  const bookId = router.query.bookId

  return <h1>书籍{bookId}的详情页</h1>

}

export default BookDetail
```

**嵌套动态路由**

```react
import { useRouter } from "next/router"

function Review () {
  const router = useRouter()
  const { bookId, reviewId } = router.query

  return <h1>书籍{bookId}的评论{reviewId}</h1>

}

export default Review
```

**守卫(Catch-all)路由**

# GO

https://www.bilibili.com/video/BV1fD4y1m7TD/?spm_id_from=333.337.search-card.all.click&vd_source=5c46b11c8d760605427b8431fb93d551

软件工艺师 Go 快速入门

web 框架

Restful API

理解业务逻辑

设计数据存储

Docker/Grafana

# The Graph

用于索引和查询区块链数据的去中心化协议（索引事件）

做好的 The Graph 会提供查询 URL

登链 w4-2 视频课程

## 写 Subgraph 程序

npm install -g @graphprotocol/graph-cli 全局安装

### **The Graph 工程结构**

-schema.graphql // 定义主体

-src

​ -xxx.ts // 定义索引规则

-subgraph.yaml // 配置

### Init

graph init --product hosted-service 2-sep/learn

cd learn

任务：索引用户 BetBear

修改**schema.graphql**（主体，表）

**subgraph.yaml**也要对应修改

graph codegen

**prediction-matic.ts**（索引规则）

如果自己定义了新的 entity，需要自己实现 handleEntity

updateUserTxandAmount{

}

graph build

### Deploy

graph auth --product hosted-service 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

graph deploy --product hosted-service 2-sep/learn

### 问题 ：

（1）怎么理解 subgraph.yaml 中的 entities

（2）怎么按照规范写 schema.graphql

## 前端查询

1.按照 Apollo(GraphQL 请求库)

npm install --save graphql graphql-tag @apollo/client

https://thegraph.com/docs/zh/querying/querying-the-graph/

## GraphQL 查询语言

https://www.bilibili.com/video/BV1fb4y117ya?p=14&vd_source=5c46b11c8d760605427b8431fb93d551

# 算法

https://leetcode.cn/problem-list/2cktkvj/?page=1

## 大神

Krahets

labuladong

kkk

反转链表

## 算法题 Java 解基础语法

##### 链表

```java
ListNode fname(ListNode list){
    // 虚拟头结点
	ListNode dummy = new ListNode(-1);
    ListNode p = dummy;

    // 连接新结点
    ListNode node = xxx;
    p.next = node;
    p = p.next;

    // 删除某节点n，要有其前序节点的指针 n-1
    ListNode noden-1;
    noden-1.next = n.next;

    return dummy.next;
}

```

##### 动态数组 ArrayList

```java
// 存String
ArrayList<String> strings = new ArrayList<>();
// 存int
ArrayList<Integer> nums = new ArrayList<>();

// 数组尾部添加元素e
boolean add(E e)
```

##### 双链表 LinkedList

LinkedList 底层是双链表实现的

```java
// 初始化
// 钻石操作符
// 泛型中只能使用类，Integer是int的封装类
LinkedList<Integer> nums = new LinkedList<>();
LinkedList<String> strings = new LinkedList<>();

List<List<Integer>> res = new LinkedList<>();

// 方法
boolean add(E e)

boolean addAll(E e)

```

##### 数组

```java
// 初始化
int[] nums = new int[n];
int[] nums = new int[nums.length];

//数组填充
Arrays.fill(dp,1);

//Math
Math.max(a,b)
```

```java
ListNode fname(ListNode[] lists){
	//迭代器
    for(ListNode head : Lists){
        if(head != null){

        }
    }
}
```

##### 泛型编程

```java
// 装整数的双链表
LinkedList<Integer> list1 = new LinkedList<>();
// 报错，不能用 int 这种原始类型作为泛型
LinkedList<int> list2 = new LinkedList<>();

// 装字符串的双链表
LinkedList<String> list3 = new LinkedList<>();
```

##### Java 接口类型

```java
Queue<String> q = new LinkedList<>();
List<String> list = new LinkedList<>();
```

##### 优先级队列（最小堆、最大堆）

```java
// 实例化优先级队列 最小堆
PriorityQueue<ListNode> pq = new PriorityQueue<>(list.length,(a,b)->(a.val - b.val));
// 取最值
ListNode node = pq.poll();
// 添加结点
pq.add(node);
```

## 数据结构和算法的框架思维

### 一、数据结构的存储方式

（1）数组（顺序存储）

​ 紧凑连续存储

​ 随机访问

​ 在数组中间进行插入和删除，时间复杂度 O（N）

（2）链表（链式存储）

​ 元素不连续，靠指针指向下一个元素的位置

​ 不能随机访问

​ 增删的时间复杂度 O（1）

## 数据结构的基本操作

遍历+访问，增删查改，线性与非线性

线性：for/while 迭代

非线性：递归

##### 数组遍历框架，线性迭代结构

```java
void traverse(int[] arr){
    for(int i = 0; i < arr.length;i++){
        // 迭代访问arr[i]
    }
}
```

##### 链表遍历框架，兼具迭代和递归结构

```java
/* 基本的单链表节点 */
class ListNode {
    int val;
    ListNode next;
}

void traverse(ListNode head){
    for(ListNode p = head; p != null; p = p.next){
        // 迭代访问p.val
    }
}

void traverse(ListNode head){
    // 递归访问 head.val
    traverse(head.next)
}
```

##### 二叉树遍历框架，非线性遍历结构

```java
/* 基本的二叉树节点 */
class TreeNode {
    int val;
    TreeNode left, right;
}

void traverse(TreeNode root){
    traverse(root.left)
    traverse(root.right)
}
```

### 三、算法刷题指南

（1）二叉树专题

（2）动态规划、回溯算法、分治算法、图论

### 四、最后总结

## 链表

### 双指针技巧秒杀七道链表题目

##### 1.LeetCode 21:合并两个有序链表

```java
ListNode mergeTwoLists(ListNode l1,ListNode l2){
    // 虚拟头结点
    ListNode dummy = new ListNode(-1),p = dummy;
    ListNode p1 = l1,p2 = l2;

    while(p1 != null && p2 != null){
        if(p1.val > p2.cal){
            p.next = p2;
            p2 = p2.next;
        } else{
            p.next = p1;
        	p1 = p1.next;
        }

        // p 指针不断前进
        p = p.next;
    }

    if(p1 != null){
        p.next = p1;
    }

    if(p2 != null){
        p.next = p2;
    }

    return dummy.next;
}
```

##### 2.LeetCode 86:分隔链表

```java
ListNode partition(ListNode head, int x) {
    // 存放小于 x 的链表的虚拟头结点
    ListNode dummy1 = new ListNode(-1);
    // 存放大于等于 x 的链表的虚拟头结点
    ListNode dummy2 = new ListNode(-1);
    // p1, p2 指针负责生成结果链表
    ListNode p1 = dummy1, p2 = dummy2;
    // p 负责遍历原链表，类似合并两个有序链表的逻辑
    // 这里是将一个链表分解成两个链表
    ListNode p = head;
    while (p != null) {
        if (p.val >= x) {
            p2.next = p;
            p2 = p2.next;
        } else {
            p1.next = p;
            p1 = p1.next;
        }
        // 断开原链表中的每个节点的 next 指针
        ListNode temp = p.next;
        p.next = null;
        p = temp;
    }
    // 连接两个链表
    p1.next = dummy2.next;

    return dummy1.next;
}
```

##### 3.LeetCode 23:合并 k 个有序链表

```java
ListNode mergeKLists[ListNode[] lists]{
	if(lists.length == 0) return null;
    // 虚拟头结点
    ListNode dummy = new ListNode(-1);
    ListNode p = dummy;
    // 优先级队列，最小堆
    PriorityQueue<ListNode> pq = new PriorityQueue<>(lists.length,(a,b) -> (a.val - b.val));
    // 将k个链表的头结点加入最小堆
    for(ListNode head : lists){
        if(head != null)
            pq.add(head);
    }

    while(!pq.isEmpty()){
        // 获取最小结点，接到结果链表中
        ListNode node = pq.poll();
        p.next = node;
        if(node.next != null){
            pq.add(node.next);
        }
        // p指针不断前进
        p = p.next;
    }
    return dummy.next;
}
```

##### 4. 单链表的倒数第 k 个节点

题目只给一个头结点 ListNode

```java
// 双指针
// 返回链表的倒数第k个节点
ListNode findFromEnd(ListNode head,int k){
    ListNode p1 = head;
    // p1先走k步
    for(int i=0; i < k;i++){
        p1 = p1.next;
    }
    ListNode p2 = head;
    // p1 和 p2 同时走 n-k步
    while(p1 != null){
        p2 = p2.next;
        p1 = p1.next;
    }
    // p2现在指向第n - k个节点
    return p2;
}
```

Leetcode 19:删除链表的倒数第 N 个节点

```java
public ListNode removeNthFromEnd(ListNode head,int n){
    // 虚拟头节点
    ListNode dummy = new ListNode(-1);
    dummy.next = head;
    // 删除倒数第 n 个，要先找倒数第 n + 1 个节点
    ListNode x = findFromEnd(dummy, n + 1);
    // 删掉倒数第 n 个节点
    x.next = x.next.next;
    return dummy.next;
}

ListNode findFromEnd(ListNode head,int k){
    ListNode p1 = head;
    // p1先走k步
    for(int i=0; i < k;i++){
        p1 = p1.next;
    }
    ListNode p2 = head;
    // p1 和 p2 同时走 n-k步
    while(p1 != null){
        p2 = p2.next;
        p1 = p1.next;
    }
    // p2现在指向第n - k个节点
    return p2;
}
```

##### 5.单链表的中点

快慢指针

```java
ListNode middleNode(ListNode head) {
    // 快慢指针初始化指向 head
    ListNode slow = head, fast = head;
    // 快指针走到末尾时停止
    while (fast != null && fast.next != null) {
        // 慢指针走一步，快指针走两步
        slow = slow.next;
        fast = fast.next.next;
    }
    // 慢指针指向中点
    return slow;
}
```

##### 5.判断链表是否包环

##### 6.两个链表是否相交

#####

### 2.递归魔法：反转单链表

![image-20230807070337272](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230807070337272.png)

## 双指针技巧秒杀七道数组题目

左右指针

快慢指针:有序数组/链表去重

​ 滑动窗口的快慢指针

##### 一、有序数组/链表去重

##### LeetCode 26:删除排序数组中的重复项

```java
// 快慢指针
// slow指针维护去重数组
int removeDuplicates(int[] nums) {
    if (nums.length == 0) {
        return 0;
    }
    int slow = 0, fast = 0;
    while (fast < nums.length) {
        if (nums[fast] != nums[slow]) {
            slow++;
            // 维护 nums[0..slow] 无重复
            nums[slow] = nums[fast];
        }
        fast++;
    }
    // 数组长度为索引 + 1
    return slow + 1;
}
```

##### LeetCode 83:「删除排序链表中的重复元素」

```java
ListNode deleteDuplicates(ListNode head) {
    if (head == null) return null;
    ListNode slow = head, fast = head;
    while (fast != null) {
        if (fast.val != slow.val) {
            // nums[slow] = nums[fast];
            slow.next = fast;
            // slow++;
            slow = slow.next;
        }
        // fast++
        fast = fast.next;
    }
    // 断开与后面重复元素的连接
    slow.next = null;
    return head;
}
```

##### LeetCode 27:移除元素

```java
int removeElement(int[] nums, int val) {
    int fast = 0, slow = 0;
    while (fast < nums.length) {
        if (nums[fast] != val) {
            nums[slow] = nums[fast];
            slow++;
        }
        fast++;
    }
    return slow;
}
```

##### LeetCode 283:移动零

```java
public void moveZeroes(int[] nums) {
    int slow = 0;
    int fast = 0;
    while(fast < nums.length){
        if(nums[fast] != 0){
            int temp = nums[slow];
            nums[slow] = nums[fast];
            nums[fast] = temp;
            slow++;
        }

        fast++;
    }

    return;
}
```

##### 二、左右指针的常用算法

##### 1.二分查找

##### 2.两数之和

##### 3.反转数组

##### 4.回文串判断

## 二叉树框架

##### 把题目的要求细化，搞清楚根节点应该做什么，然后把剩下的事情交给前/中/后序的遍历框架，不要跳进递归的细节中

traverse 遍历

```c++
/*二叉树遍历框架*/
void traverse(TreeNode root){
    // 前序遍历
    traverse(root.left)
    // 中序遍历
    traverse(root.right)
    // 后序遍历
}
```

搞清楚当前 root 节点该做什么，然后根据函数定义递归调用子节点

##### 1.计算二叉树有多少个节点

```java
int count(TreeNode root){
    if(root == null) return 0;
    return 1 + count(root.left) + count(root.right)
}
```

##### 2.Leetcode 226:翻转二叉树

```java
TreeNode invertTree(TreeNode root){
    if(root == null){
        return null;
    }
    TreeNode tmp = root.left;
    root.left = root.right;
    root.right = tmp;

    invertTree(root.left);
    invertTree(root.right);

    return root;
}
```

##### 3.Leetcode 116:填充每个节点的下一个右侧节点指针

```java
Node connect(Node root){
    if(root == null) return null;
    connectTwoNode(root.left,root.right);
    return root;
}

void connectTwoNode(Node node1,Node node2){
    if(node1 == null || node2 == null){
        return;
    }
    // 将传入的两个节点连接
    node1.next = node2;
    // 连接相同父节点的两个子节点
    connectTwoNode(node1.left,node1.right);
    connectTwoNode(node2.left,node2.right);

    // 连接跨越父节点的两个子节点
    connectTwoNode(node1.right,node2.left);
}
```

##### 4.Leetcode 114:二叉树展开为链表

labuladong:

```java
// flatten的定义就是拉平
// 递归
public void flatten(TreeNode root) {
    // base case
    if (root == null) return;

    flatten(root.left);
    flatten(root.right);

    /**** 后序遍历位置 ****/
    // 1、左右子树已经被拉平成一条链表
    TreeNode left = root.left;
    TreeNode right = root.right;

    // 2、将左子树作为右子树
    root.left = null;
    root.right = left;

    // 3、将原先的右子树接到当前右子树的末端
    TreeNode p = root;
    while (p.right != null) {
        p = p.right;
    }
    p.right = right;
}
```

```java
// 思考每个节点干了什么
// 左节点挂到右边
// 右节点挂到左节点的最后
public void flatten(TreeNode root) {
    if(root == null) return;
    TreeNode tmp = root.right;
    root.right = root.left;
    root.left = null;
    TreeNode p = root;
    while(p.right != null){
        p = p.right;
    }
    p.right = tmp;
    flatten(root.right);
}
```

##### 5.Leetcode 654:最大二叉树

```java
// 找到最大值和索引
// 数组的索引很重要
TreeNode constructMaximumBinaryTree(int[] nums) {
	return build(nums,0,nums.length - 1)
}

TreeNode build(int[] nums,int low, int high){
    if(low > high) return null;

    // 找到数组中的最大值和对应的索引
    int index = -1,maxVal = Integer.MIN_VALUE;
    for (int i = low; i <= high; i++){
        if(maxVal < nums[i]){
            index = i;
            maxVal = nums[i];
        }
    }

    TreeNode root = new TreeNode(maxVal);
    // 递归调用构建左右子树
    root.left = build(nums,low,index - 1)
    root.right = build(nums, index + 1,high);

    return root;
}
```

##### 6.Leetcode 105:通过前序和中序遍历结果构造二叉树

```java
public TreeNode buildTree(int[] preorder, int[] inorder) {
    return build(preorder,0,preorder.length - 1,inorder,0,inorder.length-1);
}

TreeNode build(int[] preorder,int preStart,int preEnd,int[] inorder,int inStart,int inEnd){
    if (preStart > preEnd) return null;

    int rootVal = preorder[preStart];

    // 找rootVal在中序遍历中的索引
    int index = 0;
    for (int i = inStart; i <= inEnd;i++){
        if(inorder[i] == rootVal){
            index = i;
            break;
        }
    }

    int leftSize = index - inStart;
    TreeNode root = new TreeNode(rootVal);
    root.left = build(preorder,preStart+1,preStart+leftSize,inorder,inStart,index-1);
    root.right = build(preorder,preStart+1+leftSize,preEnd,inorder,index+1,inEnd);

    return root;
}
```

##### 7.Leetcode 106:通过中序和后序遍历结果构造二叉树

```java
public TreeNode buildTree(int[] inorder, int[] postorder) {
    return build(inorder,0,inorder.length-1,postorder,0,postorder.length-1);
}

TreeNode build(int[] inorder,int inStart,int inEnd,int[] postorder,int postStart,int postEnd){
    if(inStart > inEnd) return null;

    int rootVal = postorder[postEnd];
    int index = 0;
    // 找rootVal在中序遍历中的位置
    for(int i=inStart; i <= inEnd ; i++){
        if(inorder[i] == rootVal){
            index = i;
            break;
        }
    }
    // ******这里非常重要 不能简单地认为 index就是长度
    int leftSize = index - inStart;
    TreeNode root = new TreeNode(rootVal);

    root.left = build(inorder,inStart,index - 1,
                      postorder,postStart,postStart+index-inStart-1);
    root.right = build(inorder,index+1,inEnd,
                       postorder,postStart+index-inStart,postEnd-1);
    return root;
}
```

##### 8.Leetcode 652:寻找重复的子树

```java

```

### 二叉搜索树 BST

### 1.左 < 根 < 右

### 2.BST 的中序遍历是有序的（升序）

```java
void traverse(TreeNode root){
    if(root == null) return;
    traverse(root.left);
    // 中序遍历代码位置
    print(root.val);
    traverse(root.right);
}
```

##### 1.Leetcode 230:二叉搜索树中第 K 小的元素

时间复杂度：O(N)

优化：节点多记录一个信息： int size;

​ 以该节点为根的树的节点总数

```java
int kthSmallest(TreeNode root, int k) {
    // 利用 BST 的中序遍历特性
    traverse(root, k);
    return res;
}

// 记录结果
int res = 0;
// 记录当前元素的排名
int rank = 0;
void traverse(TreeNode root, int k) {
    if (root == null) {
        return;
    }
    traverse(root.left, k);
    /* 中序遍历代码位置 */
    rank++;
    if (k == rank) {
        // 找到第 k 小的元素
        res = root.val;
        return;
    }
    /*****************/
    traverse(root.right, k);
}
```

##### 2.Leetcode 548 和 1038：BST 转换累加树

```java
// BST中序遍历可以得到顺序值
// 进行先访问右子树的中序遍历
// 全局遍历sum
TreeNode convertBST(TreeNode root) {
    traverse(root);
    return root;
}

// 记录累加和
int sum = 0;
void traverse(TreeNode root) {
    if (root == null) {
        return;
    }
    traverse(root.right);
    // 维护累加和
    sum += root.val;
    // 将 BST 转化成累加树
    root.val = sum;
    traverse(root.left);
}
```

### 3.BST 的基础操作：判断 BST 的合法性、增、删、查

1、如果当前节点会对下面的子节点有整体影响，可以通过辅助函数增长参数列表，借助参数传递信息。

2、在二叉树递归框架之上，扩展出一套 BST 代码框架：

```java
void BST(TreeNode root, int target) {
    if (root.val == target)
        // 找到目标，做点什么
    if (root.val < target)
        BST(root.right, target);
    if (root.val > target)
        BST(root.left, target);
}
```

3、根据代码框架掌握了 BST 的增删查改操作。

（1）判断 BST 的合法性

```java
// 1.左 < 根 < 右
// 2.root的整个左子树 < 根 < root的右子树
// 使用辅助函数，增加函数参数，在参数中携带额外信息，将这种约束传递给子树的所有节点
boolean isValidBST(TreeNode root){
	return isValidBST(root,null,null);
}

/* 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val */
boolean isValidBST(TreeNode root,TreeNode min,TreeNode max){
    // base case
    if(root == null) return true;
    // 若root.val 不符合 max 和 min 的限制，说明不是合法BST
    if (min != null && root.val <= min.val) return false;
    if (max != null && root.val >= max.val) return false;
    // 限定左子树的最大值是root.val，右子树的最小值是root.val
    return isValidBST(root.left,min,root) && isValid(root.right,root,max);
}
```

（2）在 BST 中搜索一个数

```java
// 二分查找思想
void BST(TreeNode root,int target){
    if(root.val == target)
    	// 找到目标，做点什么
    if(root.val < target)
        BST(root.right,target);
    if(root.val > target)
        BST(root.left,target);
}
```

（3）在 BST 中插入一个数

```java
// 遍历 + 访问
TreeNode insertIntoBST(TreeNode root,int val){
    // 找到空位置插入新节点
    if(root == null) return new TreeNode(val);

    if(root.val < val)
        root.right = insertIntoBST(root.right,val);
    if (root.val > val)
        root.left = insertIntoBST(root.left, val);
    return root;
}
```

（4）在 BST 中删除一个数

```java
// 末端节点，两个子节点都为空，直接去世
// 只有一个非空子节点，让孩子接替自己的位置
// 有2个子节点，让左子树中最大的节点或右子树中最小的节点接替自己
TreeNode deleNode(TreeNode root,int Key){
    if(root.val == key){
        // 找到啦，进行删除
    } else if(root.val > key){
        // 去左子树找
        root.left = deleteNode(root.left,key);
    } else if(root.val < key){
        // 去右子树找
        root.right = deleteNode(root.right,key);
    }
    return root;
}
```

### 4.计算所有合法 BST

```java

```

### 二叉树、二叉搜索树总结

### （1）深入理解前中后序

前序位置：刚进入一个节点（元素）的时候

后序位置：即将离开一个节点（元素）的时候

前序位置的代码在刚刚进入一个二叉树节点的时候执行；

后序位置的代码在将要离开一个二叉树节点的时候执行；

中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行

### （2）两种解题思路

#### 2.1 遍历一遍二叉树得出答案 -- 回溯算法核心框架

##### Leetcode 104：二叉树的最大深度

###### 思路 1：遍历一遍二叉树，用外部变量记录在每个节点所在的深度

```java
// 遍历一遍二叉树，用外部变量记录在每个节点所在的深度
// 记录最大深度
int res = 0;
// 记录遍历到的节点的深度
int depth = 0;

// 主函数
int maxDepth(TreeNode root) {
    traverse(root);
    return res;
}

// 二叉树遍历框架
void traverse(TreeNode root) {
    if (root == null) {
        // 到达叶子节点，更新最大深度
        res = Math.max(res, depth);
        return;
    }
    // 前序位置
    depth++;
    traverse(root.left);
    traverse(root.right);
    // 后序位置
    depth--;
}
```

```java
// 和标准解法类似，不过深度的变更是在递归时实现的
public int maxDepth(TreeNode root) {
    return traverse(root,0);
}

int depth = 0;
int depthMax = 0;

// 回到根的时候，变成它的深度
int traverse(TreeNode root,int depth){
    if(root == null){
        if(depth > depthMax)
            depthMax = depth;

        return depthMax;
    }

    traverse(root.left,depth + 1);
    traverse(root.right,depth + 1);

    return depthMax;
}
```

#### 2.2 通过分解问题计算出答案 --动态规划核心框架

###### 思路 2：分解问题计算答案

```java
// 定义：输入根节点，返回这棵二叉树的最大深度
int maxDepth(TreeNode root) {
    if (root == null) {
        return 0;
    }
    // 利用定义，计算左右子树的最大深度
    int leftMax = maxDepth(root.left);
    int rightMax = maxDepth(root.right);
    // 整棵树的最大深度等于左右子树的最大深度取最大值，
    // 然后再加上根节点自己
    int res = Math.max(leftMax, rightMax) + 1;

    return res;
}
```

二叉树前序遍历

```java
// 不借助辅助函数和任何外部变量
List<Integer> preorderTraverse(TreeNode root) {
    List<Integer> res = new LinkedList<>();
    if(root == null) reutrn res;

    res.add(root.val);
    res.addAll(preorderTraverse(root.left));
    rea.addAll(preorderTraverse(root.right));

    return res;
}
```

### （3）二叉树通用解题思路

是否可以通过遍历一遍二叉树得到答案？如果不能的话，是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？

### （4）后序位置的特殊之处

中序位置：BST 场景

前序位置的代码只能从函数参数中获取父节点传递来的数据

后序位置的代码：获取参数数据、子树通过函数返回值传递回来的数据

##### 如果题目与子树有关，大概率要给函数设置合理的定义和返回值，在后序位置写代码

### （5）层序遍历

```java
// 输入一棵二叉树的根节点，层序遍历这棵二叉树
void levelTraverse(TreeNode root) {
    if (root == null) return;
    Queue<TreeNode> q = new LinkedList<>();
    q.offer(root);

    // 从上到下遍历二叉树的每一层
    while (!q.isEmpty()) {
        int sz = q.size();
        // 从左到右遍历每一层的每个节点
        for (int i = 0; i < sz; i++) {
            TreeNode cur = q.poll();
            // 将下一层节点放入队列
            if (cur.left != null) {
                q.offer(cur.left);
            }
            if (cur.right != null) {
                q.offer(cur.right);
            }
        }
    }
}
```

### 归并排序（抽象成二叉树后序遍历）

```java
// 归并排序，先把左半边数组排好序，再把右半边数组排好序，然后把两半数组合并
// 用于辅助合并有序数组
private static int[] temp;

public static void sort(int[] nums) {
    // 先给辅助数组开辟内存空间
    temp = new int[nums.length];
    // 排序整个数组（原地修改）
    sort(nums, 0, nums.length - 1);
}

// 定义：将子数组 nums[lo..hi] 进行排序
private static void sort(int[] nums, int lo, int hi) {
    if (lo == hi) {
        // 单个元素不用排序
        return;
    }
    // 这样写是为了防止溢出，效果等同于 (hi + lo) / 2
    int mid = lo + (hi - lo) / 2;
    // 先对左半部分数组 nums[lo..mid] 排序
    sort(nums, lo, mid);
    // 再对右半部分数组 nums[mid+1..hi] 排序
    sort(nums, mid + 1, hi);
    // 将两部分有序数组合并成一个有序数组
    merge(nums, lo, mid, hi);
}

// 将 nums[lo..mid] 和 nums[mid+1..hi] 这两个有序数组合并成一个有序数组
private static void merge(int[] nums, int lo, int mid, int hi) {
    // 先把 nums[lo..hi] 复制到辅助数组中
    // 以便合并后的结果能够直接存入 nums
    for (int i = lo; i <= hi; i++) {
        temp[i] = nums[i];
    }

    // 数组双指针技巧，合并两个有序数组
    int i = lo, j = mid + 1;
    for (int p = lo; p <= hi; p++) {
        if (i == mid + 1) {
            // 左半边数组已全部被合并
            nums[p] = temp[j++];
        } else if (j == hi + 1) {
            // 右半边数组已全部被合并
            nums[p] = temp[i++];
        } else if (temp[i] > temp[j]) {
            nums[p] = temp[j++];
        } else {
            nums[p] = temp[i++];
        }
    }
}

```

### 快速排序

快速排序的过程是一个构造二叉搜索树的过程

```java
void sort(int[] nums,int lo,int hi){
    if(lo >= hi){
        return;
    }
    // 对nums[lo . . hi]进行切分
    // 使得nums[lo . .p-1] <= nums[p] <= nums[p+1 . . hi]
    int p = partition(nums,lo,hi);
    // 去左右子数组进行切分
    sort(nums,lo,p-1);
    sort(nums,p+1,hi);
}
```

## 手把手刷图算法

## 手把手刷设计数据结构

### 二叉堆 Binary Heap

堆排序（排序）

优先级队列（数据结构）

##### 一、二叉堆概览

特殊的二叉树（完全二叉树），存储在数组中

最大堆：每个节点都大于等于它的两个子节点

最小堆：每个节点都小于等于它的子节点

```java
// 父结点的索引
int parent(int root){
    return root / 2;
}
// 左孩子的索引
int left(int root){
    return root * 2;
}
// 右孩子的索引
int right(int root){
    return root * 2 + 1;
}
```

##### 二、优先级队列概览

功能：插入或删除元素时，元素会自动排序

```java

```

##### 三、实现 swim 和 sink

```java
private void swim(int k){
    // 如果浮到堆顶，就不能再上浮了
    while(k > 1 && less(parent(k),k)){
        // 如果第k个元素比上层大
        // 将k换上去
        exch(parent(k),k);
        k = parent(k);
    }
}

private void sink(int k){
    // 如果沉到堆底，就沉不下去了
    while (left(k) <= N){
        // 先假设左边节点较大
        int older = left(k);
        // 如果右边节点存在，比一下大小
        if(right(k) <= N && less(older,right(k)))
            older = right(k);
        // 结点k比俩孩子都大，就不必下沉了
        if(less(older,k)) break;
        // 否则，不符合最大堆的结构，下沉k结点
        exch(k,older);
        k = older;
    }
}
```

##### 四、实现 delMax 和 insert

```java
// insert
// 把要插入的元素添加到堆底的最后，然后让其上浮到正确位置
public void insert(Key e){
    N++;
    // 先把新元素加到最后
    pq[N] = e;
    // 然后让它上浮到正确的位置
    swim(N);
}

// delMax
// 先把堆顶元素A和堆底最后的B对调，删除A，让B下沉
public Key delMax() {
    // 最大堆的堆顶就是最大元素
    Key max = pq[1];
    // 把这个最大元素换到最后，删除之
    exch(1,N);
    pq[N] = null;
    N--;
    // 让pq[1]下沉到正确位置
    sink(1);
    return max;
}
// 插入
```

## 滑动窗口框架

## 动态规划

动态规划基本技巧

穷举

存在「重叠子问题」

具备「最优子结构」

正确的「状态转移方程」

明确「状态」 -> 定义 dp 数组/函数的含义 -> 明确「选择」-> 明确 base case。

暴力解，用备忘录、DP table 优化

##### 一、斐波那契数列

###### 1.暴力递归

```java
int fib(int N) {
    if (N == 1 || N == 2) return 1;
    return fib(N - 1) + fib(N - 2);
}
```

###### 2.带备忘录的递归算法(自顶向下)

```java
int fib(int N) {
    if (N < 1) return 0;
    // 备忘录全初始化为 0
    vector<int> memo(N + 1, 0);
    // 初始化最简情况
    return helper(memo, N);
}

int helper(vector<int>& memo, int n) {
    // base case
    if (n == 1 || n == 2) return 1;
    // 已经计算过
    if (memo[n] != 0) return memo[n];
    memo[n] = helper(memo, n - 1) +
                helper(memo, n - 2);
    return memo[n];
}
```

###### 3.dp 数组的迭代解法（自底向上） 动态规划、脱离递归、循环迭代

```java
int fib(int N){
    vector<int> dp(N + 1,0);
    // base case
    dp[1] = dp[2] = 1;
    for(int i = 3; i <= N; i++)
        dp[i] = dp[i - 1] + dp[i - 2];
    return dp[N];
}
```

##### 二、凑零钱问题

###### 1.暴力递归

```java
# 伪码框架

```

### 动态规划设计：最长递增子序列

```java
int lengthOfLIS(int[] nums){
    // 定义：dp[i] 表示以nums[i]这个数结尾的最长递增子序列的长度
    int[] dp = new int[nums.length];
    // base case:dp数组全都初始化为1
    Arrays.fill(dp,1);
    for(int i = 0; i < nums.length;i++){
        for(int j = 0; j < i; j++){
            if(nums[i] > nums[j])
                dp[i] = Math.max(dp[i],dp[j]+1);
        }
    }

    int res = 0;
    for(int i = 0; i < dp.length; i++){
        res = Math.max(res,dp[i]);
    }
    return res;
}
```

如何找到动态规划的状态转移关系：

1.明确 dp 数组的定义。

2.根据 dp 数组的定义，运用数学归纳法的思想，假设 dp[0 ... i-1]都已知，想办法求出 dp[i]

### 最优子结构原理和 dp 数组方向

1.最优子结构

2.如何一眼看出重叠子问题

3.dp 数组的大小设置

4.

### 子序列类型问题

- ##### 经典动态规划：编辑距离

- ##### 动态规划设计：最长递增子序列

- ##### 动态规划设计：最大子数组

```java
// 以nums[i]为结尾的 最大子数组和 为dp[i]

```

- ##### 经典动态规划：最长公共子序列

- ##### 动态规划之子序列问题解题模板

```java

```

### 背包类型问题

- ##### 经典动态规划：0-1 背包问题

1. 状态 和 选择
2. dp 数组的定义

```java
int dp[N+1][W+1]
dp[0][..] = 0
dp[..][0] = 0

for i in [1..N]:
	for w in [1..W]:
		dp[i][w] = max(
            把物品i装进背包，
            不把物品i装进背包
        )
return dp[N][W]
```

- ##### 经典动态规划：子集背包问题

- ##### 经典动态规划：完全背包问题

- ##### 动态规划和回溯算法到底谁是谁爹

## 动态规划玩游戏

- ##### 一个方法团灭 LeetCode 股票买卖问题

## 贪心类型问题

## 算法技巧

### 暴力搜索算法

##### 回溯算法解题框架

解决一个回溯问题，实际上就是一个决策树的遍历过程

1.路径

2.选择列表

3.结束条件

```java
result = []
def backtrack(路径，选择列表):
	if 满足结束条件:
		result.add(路径)
        return

    for 选择 in 选择列表:
		做选择
        backtrack(路径，选择列表)
        撤销选择
```

###### 一、全排列问题

```java
// 多叉树的遍历框架
void traverse(TreeNode root){
    for(TreeNode child:root.children)
        // 前序遍历需要的操作
        traverse(child);
    	// 后序遍历需要的操作
}
```

```java
// 全排列代码
List<List<Integer>> res = new LinkedList<>();

/* 主函数，输入一组不重复的数字，返回它们的全排列 */
List<List<Integer>> permute(int[] nums){
    // 记录 路径
    LinkedList<Integer> track = new LinkedList<>();
    backtrack(nums,track);
    return res;
}


```

###### 二、N 皇后问题

### 数学运算技巧

### 经典面试题

## 题库

LeetCode 94:二叉树的中序遍历

```java
class Solution {
    /* 动态规划思路 */
    // 定义：输入一个节点，返回以该节点为根的二叉树的中序遍历结果
    public List<Integer> inorderTraversal(TreeNode root) {
        LinkedList<Integer> res = new LinkedList<>();
        if (root == null) {
            return res;
        }
        res.addAll(inorderTraversal(root.left));
        res.add(root.val);
        res.addAll(inorderTraversal(root.right));
        return res;
    }

    /* 回溯算法思路 */
    LinkedList<Integer> res = new LinkedList<>();

    // 返回前序遍历结果
    public List<Integer> inorderTraversal2(TreeNode root) {
        traverse(root);
        return res;
    }

    // 二叉树遍历函数
    void traverse(TreeNode root) {
        if (root == null) {
            return;
        }
        traverse(root.left);
        // 中序遍历位置
        res.add(root.val);
        traverse(root.right);
    }
}
```

# Git

创建仓库初始化

![img](https://www.runoob.com/wp-content/uploads/2015/02/git-command.jpg)

基本操作

git add .

git commit -m ""

git push

git pull

码农高天

工作流:https://www.bilibili.com/video/BV19e4y1q7JJ/?spm_id_from=333.999.0.0&vd_source=5c46b11c8d760605427b8431fb93d551

https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485544&idx=1&sn=afc9d9f72d811ec847fa64108d5c7412&scene=21#wechat_redirect

主要

1.Git 主要命令

git commit

**# 分支**

```shell
// 查看所有分支

git branch

// 新建分支

git branch newName

// 切换分支

git checkout newName

// 删除分支
git branch -D newName

```

git merge xxx

git rebase xxx

2.Git 超棒特性

分离 HEAD

相对引用^，相对引用~

强制移动分支

​ git branch -f main HEAD~3

撤销变更

​ git reset HEAD~1

​ git revert HEAD

3.自由修改提交树

Git Cherry-pick

​ git cherry-pick xxx xxx xxx

交互式 rebase

​ git rebase -i HEAD~4

4.Git 技术、技巧与贴士大集合

只取一个提交记录

提交的技巧

提交的技巧 2

Git Tag

Git Describe

5.只为真正的勇士

多次 Rebase

两个 parent 节点

纠缠不清的分支

远程

Push & Pull -Git 远程仓库

Git Clone

远程分支

Git Fetch-从远程仓库获取数据

- 从远程仓库下载本地仓库中缺失的提交记录
- 更新远程分支指针(如 o/main)

Git Pull

​ git pull = git fetch + git merge

模拟团队合作

Git Push

偏离的提交历史

锁定的 Main(Locked Main)

关于 origin 和它的周边

#####

git 工作流

git checkout main

git pull origin main

git checkout my-feature

git rebase main

git push -f origin my-feature

New pull request(squash 挤压 and merge)

删除远端的 my-feature 分支

git checkout main

git branch -D my-feature

git pull origin main

https://learngitbranching.js.org/?locale=zh_CN

## github 主页美化

https://blog.csdn.net/qq_44231797/article/details/129251980

https://github.com/mayhemantt

# windows terminal

cd /d E:\code

start E:\code

mkdir 新建文件夹

dir 查看当前目录

cls 清屏

rmdir /s 删除文件夹
