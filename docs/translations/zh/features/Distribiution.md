# distribution 模块的用户手册
## 简介
该模块负责将交易费和通胀的代币分发给验证人和委托人。为了减少网络和转账的压力。 设计为被动的分配模式。需要用户主动进行收益的领取，系统会从全局的收益池中支付用户应得的收益。

## 收益

### 收益的来源
1. 转账或者其他操作产生的交易费 fees。
2. 系统的各项服务产生的费用。
3. 系统设定的增发。

### 收益的去向
1. 验证人（受到委托在活跃验证的节点）
2. 委托人（委托代币至验证的账户）
3. 基金会（ Hashgard Foundation Ltd ）
4. 出块人收益

## 使用场景

### 设置收益回收地址
默认委托的时候, 委托地址即收益地址. 如果需要修改和设置收益取回地址:
```shell
hashgardcli distribution set-withdraw-addr [withdraw-address] [flags]
```


### 查询收益

```shell
hashgardcli distribution outstanding-rewards [flags]
```

### 取回收益

```shell
hashgardcli distribution withdraw-rewards [validator-addr] [flags]
```
对于其他查询 distribution 状态的命令，请参考[distribution](../cli/hashgardcli/distribution/README.md)
