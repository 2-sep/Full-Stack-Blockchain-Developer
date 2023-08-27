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

读取可以从任意字节处开始读取

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

对不变量，使用 constant 和 immutable 定义

2 calldata memory

3 Bitmap 位图与位运算符

**4 Unchecked **

在安全性可控的情况下，使用 unchecked

**5 Uint**

尽量变量打包，如果不能打包还是使用 uint256

6 Error

**7 LocalData 局部变量**

在较复杂的运算中，避免直接操作存储变量，可以先定义一个局部变量，对局部变量进行修改后，再修改存储变量

**8 clone 替换 new/create2**

**9 Packing**

**10 Increment 使用更好的递增**

使用 ++i，而不是 i++

**11 使用 uint 进行重入保护**

bool ×

uint 0 1，0 改 非 0 会消耗较多的 gas ×

uint 1 2，非 0 改非 0 √

12 LessThan

**13 MethodName 优化方法名**

EVM 执行交易，calldata 数据消耗 gas，0 字节消耗 4gas，非 0 字节消耗 16gas

优化函数名，增加 MethodId 中 0 字节的个数来节省 gas

**14 MethodIdSort**

合约内所有函数是一个数组，通过 MethodId 排序

遍历,找到想调用的函数