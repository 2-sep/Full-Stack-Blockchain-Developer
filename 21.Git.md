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

git branch -f main HEAD~3

撤销变更

git reset HEAD~1

git revert HEAD

3.自由修改提交树

Git Cherry-pick

git cherry-pick xxx xxx xxx

交互式 rebase

git rebase -i HEAD~4

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

git pull = git fetch + git merge

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