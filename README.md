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

- 在cft类型分支里，非merging状态下，执行add时存在二次确认。对功能类拦截，对冲突处理类拦截。
```
> gitx add .
在cft类型分支上提交内容为异常操作，请确认提交内容是为 追加冲突处理【Y】，还是功能开发【N】。
Y/N: N
禁止在冲突处理分支中，追加功能开发，add已被撤销。

> gitx add .
在cft类型分支上提交内容为异常操作，请确认提交内容是为 追加冲突处理【Y】，还是功能开发【N】。
Y/N: y
允许追加处理冲突。
warning: LF will be replaced by CRLF in go.mod.
The file will have its original line endings in your working directory

```

- 执行init,clone,version,--version之外的命令，当执行路径下不存在.git目录时，禁止继续执行该命令
```
> gitx branch -v
未在当前目录找到.git目录，请确认是合法git项目，或使用git init 初始化
```

- 执行pull命令时，只允许同名分支发起pull
```
> gitx pull origin kkk
禁止跨分支pull,只允许同名分支发起pull
```
