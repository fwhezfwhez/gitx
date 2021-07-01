## gitx
Gitx是一款git套壳工具。

- gitx使用gitx命令替代git命令，命令参数和种类，和git命令完全一致。
- gitx要求原机器必须预先安装好git命令。

## 安装
#### 通过go命令安装
go get -u github.com/fwhezfwhez/gitx

- 通过此方式安装的gitx命令，会存放于$GOPATH/bin/下
- 需要预先安装go环境

#### 使用release版本
前往 https://github.com/fwhezfwhez/gitx/releases 下载对应版本，将二进制文件，放进大家的PATH里。


## 功能

- 禁止了任意分支下，执行`merge dev`操作
```go
> gitx merge dev
merge dev是禁止操作，会导致分支污染。
```

- 禁止了在原分支以dev前缀时的创建分支操作
```go
> gitx branch kkk
禁止在祖先分支为dev的场景下，新建分支。会造成分支污染

> gitx checkout -b kkkk
禁止在祖先分支为dev的场景下，新建分支。会造成分支污染
```
