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

**Proto-Danksharding(EIP-4844)**

提议者-构建者分离(PBS)

完整分片（Danksharding）