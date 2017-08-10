# issue
Issue generates Github issue report.

## Getting Started

### Install Go

https://golang.org/doc/install#install

### Build / Update

```bash
go get -u github.com/helinwang/issue
```

### Usage
```bash
$GOPATH/bin/issue -issues 3368,3346,3366,3362
```

You will see output like:

```text
Github Issues (Total: 4, Closed: 1, Waiting user feedback: 3)

Title: cross compile on mac
Link: https://github.com/PaddlePaddle/Paddle/issues/3368
Labels: compile error, Waiting for User Feedback 🔵
Reply: helinwang, Xreki
State: open

Title: 加载预训练的embedding layer，和神经网络迭代过程中更新得到的embedding layer，我该如何单独对比两份embedding 结果的好坏呢？
Link: https://github.com/PaddlePaddle/Paddle/issues/3346
Labels: User, Waiting for User Feedback 🔵
Reply: Superjom
State: open

Title: paddle.v2.layer.cross_entropy_cost和classification_cost的区别
Link: https://github.com/PaddlePaddle/Paddle/issues/3366
Labels: question, Waiting for User Feedback 🔵
Reply: lcy-seso
State: open

Title: ubuntu 16.04  docker 安装失败
Link: https://github.com/PaddlePaddle/Paddle/issues/3362
Labels: question
Reply: helinwang
State: closed ✅

```
