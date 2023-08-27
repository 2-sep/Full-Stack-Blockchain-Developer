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