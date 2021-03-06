# go基础命令

## go run
运行命令源码文件，`go run` 后只能接受一个命令源码文件以及若干个库源码文件    
* -a : 强制编译相关代码，不论它们的编译结果是否已经是最新的   
* -n : 打印编译执行的命令，但却不执行它们
* -x : 打印编译执行的命令，并执行
* -p n : 并行编译，n为并行的数量
* -v : 列出被编译的代码包名称
* -work: 显示编译时创建的临时目录，并且不删除

## go build
用于编译源码文件或者代码包；
* 编译非命令源码文件不会产生任何结果文件，`go build`只是检查库源码文件有效性。
* 编译命令源码文件时会在执行目录下产生一个可执行文件
* 执行`go build`命令，没有任何参数时，它试图把当前目录作为代码包并编译
* 执行该命令并以代码包的导入路径作为参数时，该代码包及其依赖会被编译
* 执行该命令以若干源码文件为参数，只编译这些文件

## go install
编译并安装代码包或源码文件
* 安装代码包会在工作区pkg/<平台相关目录>/ 下生成归档文件   
* 安装命令源码文件会在工作区`bin`或者`$GOBIN`目录下生成可执行文件

## go get
用于从远处代码仓库（例如github）上下载并安装代码包；代码包会被下载到`$GOPATH`中包含的第一个工作区`src`目录下
* -d: 只执行下载工作，不执行安装动作
* -fix: 在下载代码包之后先执行修正动作，而后在进行编译安装
* -u: 利用网络来更新已有代码包和依赖

## gofmt
