# vm-geth

vm-geth 是一个用于实现 Solidity 的 EVM 动态链接工具。

项目主要利用的 [go-ethereum](https://github.com/ethereum/go-ethereum) ， 生成 so 插件文件。
Apache 等开源项目可以通过动态链接的方式，引用插件，达到使用 EVM 但不受到 LGPL 传染的目的。

## 使用方式

我们在插件中定义了 `CreatorFunc` 类型的函数，名为 `NewEvmCreator`。

您可以通过 `plugin` 包打开了 该项目生产的 so 插件文件，
然后使用然后使用 `Lookup` 方法查找插件中的函数，
并将其转换为 `CreatorFunc` 类型。

最后，您调用插件中获取到的方法，实现您的需求。

### 示例代码

```go
package example

type CreatorFunc = func(*bridge.InstanceCreatorConfig) (bridge.InstanceCreator, error)

func myFunc() {

	// load plugin
	p, err := plugin.Open("dir/vm-path.so")
	if err != nil {
		panic(err)
	}

	// find symbol
	creatorFunc, err := p.Lookup("NewEvmCreator")
	if err != nil {
		panic(err)
	}

	// convert type
	newEvmCreator, ok := creatorFunc.(CreatorFunc)
	if !ok {
		panic("unexpected type from vm-path module symbol")
	}

	// use it
	bridge.Register(typeGeth, driverName, newEvmCreator)
}
```

> 更详细的实例代码请见： [example](example/example.go)

### 注意事项

Go的主程序在加载plugin时，会在`runtime` 里对两者进行一堆约束检查，包括但不限于：
* go version一致
* go path一致
* go dependency的交集一致
	* 代码一致
	* path一致
* go build 某些flag一致