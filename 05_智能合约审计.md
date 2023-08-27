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

https://www.tutorialspoint.com/solidity/solidity_conversions.htm

https://learnblockchain.cn/docs/solidity/types.html#types-conversion-elementary-types

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

Proxy.sol

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

solmate/utils/FixedPointMathLib.sol 数学库

solmate/utils/ReentrancyGuard.sol 防重入攻击库

{ SafeTransferLib, ERC4626, ERC20 } from "solmate/mixins/ERC4626.sol" 安全转账，**ERC4626 代币金库话**，ERC20

solmate/auth/Owned.sol 权限控制

openzeppelin-contracts/contracts/interfaces/IERC3156.sol

**ERC3156**

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

abi.encodeWithSignature("函数签名", 逗号分隔的具体参数)

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