## 安装
go-fuzz 是一个用于Go语言的模糊测试工具。模糊测试是一种软件测试方法，用于发现程序处理非预期或随机数据时的问题。这里是一些使用 go-fuzz 的基本步骤：

安装 go-fuzz
首先确保您已经安装了Go。然后运行以下命令以安装 go-fuzz 和 go-fuzz-build：

```bash
go get -u github.com/dvyukov/go-fuzz/go-fuzz
go get -u github.com/dvyukov/go-fuzz/go-fuzz-build
```
创建待测试的Go代码
假设您有一个函数，该函数需要模糊测试。例如，一个处理字符串的函数：

```go
// fuzz.go
package fuzz

func Fuzz(data []byte) int {
    s := string(data)
    if s == "bad string" {
        panic("bad string found")
    }
    return 0
}
```
构建Fuzz测试
在项目目录中运行以下命令：

```bash
go-fuzz-build github.com/yourusername/yourproject
```
这会生成一个名为 yourproject-fuzz.zip 的文件。

运行Fuzz测试
然后运行以下命令以开始模糊测试：

```bash
go-fuzz -bin=./yourproject-fuzz.zip -workdir=fuzz_workdir
```
这里，fuzz_workdir 是一个目录，其中包含模糊测试的种子数据（放在 corpus 子目录中）和一些其他状态。

go-fuzz 将开始运行，并试图找到使 Fuzz 函数崩溃的输入。

当找到这样的输入时，它会被保存在 fuzz_workdir/crashers 目录下。

解释输出
go-fuzz 运行时会输出一些统计信息，包括已运行的测试次数、已找到的问题数量等。

修复问题
当 go-fuzz 找到问题后，您可以在 fuzz_workdir/crashers 目录中查看导致问题的输入数据，并相应地修复代码。

总体而言，go-fuzz 是一个强大的工具，可帮助找到代码中难以通过常规测试发现的问题。


##  存在的问题

go get 通常会将可执行文件安装在 $GOPATH/bin 目录下，而不是 Go 的安装目录 /usr/local/go/bin。如果你没有明确设置 GOPATH，那么它默认是 $HOME/go（即你的用户主目录下的 go 文件夹）。

可以通过以下命令来查找 go-fuzz 和 go-fuzz-build：

```bash
ls $GOPATH/bin
```
或者，如果你没有设置 GOPATH：

```bash
ls ~/go/bin
```
如果你在这些目录下找到了 go-fuzz 和 go-fuzz-build，确保这些目录在你的 PATH 环境变量中：
```shell
export PATH=$PATH:$GOPATH/bin
``
或者

```shell
export PATH=$PATH:~/go/bin
```
然后再次尝试运行 go-fuzz-build。

