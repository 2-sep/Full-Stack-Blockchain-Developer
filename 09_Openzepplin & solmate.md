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

https://learnblockchain.cn/2019/09/27/erc777

ERC 1820

**ERC-3156 闪电贷**

ERC3156FlashLender

function maxFlashLoan(address token) external view returns (uint256);

function flashFee(address token, uint256 amount) external view returns (uint256);

function flashLoan(

IERC3156FlashBorrower receiver,

address token,

uint256 amount,

bytes calldata data

) external returns (bool);

ERC3156FlashBorrower

function onFlashLoan(

address initiator,

address token,

uint256 amount,

uint256 fee,

bytes calldata data

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