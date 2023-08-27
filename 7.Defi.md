# DeFi

https://learnblockchain.cn/column/4

了解架构、经济模型、白皮书部分的代码实现

## UniSwap V2

x \* y = K

交易前的瞬时价格就是斜率 即 y/x

**2.2 价格预言机**

https://zhuanlan.zhihu.com/p/359750357 如何使用 Uniswap v2 作为预言机

在每个区块的第一笔交易前记录**累计价格**实现预言机，时间权重记录

TWAP 算术平均数

两个难题：1. 应该计算以 B 代币计价的 A 代币价格，还是以 A 代币计价的 B 代币价格？ 2.用户可以不通过交易而直接向交易对合约发送代币（这将改变代币余额并影响价格），此时将无法触发预言机价格更新

价格精度：UQ112.112 格式的数据，它表示在小数点的左右两边都有 112 位比特表示精度 **Solidity 原生不支持非整数数据类型**

112+112+32 = 256

**2.3 闪电贷**

**2.4 Protocol fee** 协议手续费

**2.5 Meta transactions for pool shares** 池子份额元交易

**3.4 初始化流动性代币供应**

**3.5 Wrapping ETH - WETH**

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

每个交易对有多个池子，分别设置不同的交易手续费。

**3.2 Non-Fungible Liquidity 不可互换的流动性**

手续费不像 v2 是复利的，v3 的流动性是一个 NFT

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

恒定积公式 x \* y = k

限价流动性基于 tick 刻度体系

流动性堆叠

革新

全新的合约架构

巧妙地账本设计

更高的开发者自由度

对流动性提供者更多的保护

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

**cToken** exchangeRate 兑换率

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