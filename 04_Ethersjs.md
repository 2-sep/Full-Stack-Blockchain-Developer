# Ethers.js API
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