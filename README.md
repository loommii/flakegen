<div align="center">

# flakegen

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/github/license/loommii/flakegen)](LICENSE)

**基于雪花算法的 Go 语言分布式 ID 生成器**

[**特性**](#特性) • [**安装**](#安装) • [**快速开始**](#快速开始) • [**配置**](#配置) • [**贡献**](#贡献)

</div>

---

## 📋 项目介绍

`flakegen` 是一个用 Go 语言实现的分布式 ID 生成器，基于经典的雪花算法设计思想。该库能够在分布式环境中生成全局唯一的 64 位整数 ID，适用于高并发、大规模分布式系统的唯一 ID 生成需求。

## ✨ 特性

- **全局唯一性** - 在分布式环境下生成的 ID 全局唯一
- **单调递增** - 生成的 ID 具有时间有序性
- **高并发支持** - 支持每毫秒生成数千个 ID
- **自定义结构** - 可配置时间戳、节点 ID、业务 ID 和序列号的位数
- **时钟回拨防护** - 检测并处理系统时钟回拨问题
- **线程安全** - 使用互斥锁确保并发安全

## 📦 安装

```bash
go get github.com/loommii/flakegen
```

## 🚀 快速开始

以下是最简单的使用示例：

```go
package main

import (
	"fmt"

	"github.com/loommii/flakegen"
)

func main() {
	// 创建节点实例，参数为节点ID和服务ID
	node, err := flakegen.NewNode(1, 1)
	if err != nil {
		panic(err)
	}

	// 生成唯一ID
	id, err := node.GetID()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated ID: %d\n", id)
}
```

### 高级用法

```go
package main

import (
	"fmt"

	"github.com/loommii/flakegen"
)

func main() {
	// 创建多个不同的节点
	node1, err := flakegen.NewNode(1, 1) // 节点1，业务1
	if err != nil {
		panic(err)
	}

	node2, err := flakegen.NewNode(2, 1) // 节点2，业务1
	if err != nil {
		panic(err)
	}

	// 生成ID
	id1, _ := node1.GetID()
	id2, _ := node2.GetID()

	fmt.Printf("Node 1 ID: %d\n", id1)
	fmt.Printf("Node 2 ID: %d\n", id2)
}
```

## 🛠️ ID 结构

`flakegen` 生成的 64 位 ID 结构如下：

```
符号位(1位) + 时间戳(41位) + 节点ID(5位) + 业务ID(5位) + 序列号(12位)
```

- **符号位（1位）**: 固定为 0，确保 ID 为正数
- **时间戳（41位）**: 毫秒级时间戳，减去自定义起始时间
- **节点ID（5位）**: 用于区分不同机器，最多支持 31 个节点
- **业务ID（5位）**: 用于区分不同业务，最多支持 31 个业务
- **序列号（12位）**: 每毫秒内的序列号，最多支持 4095 个 ID/毫秒

## 📁 项目结构

```
flakegen/
├── flakegen.go        # 核心算法实现
├── flakegen_test.go   # 测试文件
├── README.md          # 项目说明文档
├── go.mod, go.sum     # Go 模块文件
```

## 🔒 安全考虑

- **时钟同步**: 系统间的时间同步对于保证 ID 唯一性至关重要
- **节点ID管理**: 在部署时需确保不同节点的 ID 唯一
- **业务ID管理**: 不同业务模块应分配不同的业务 ID

## 🤝 贡献

欢迎任何形式的贡献！您可以：

- 提交 Issue 报告 Bug 或提出功能建议
- 提交 Pull Request 改进代码
- 改善文档
- 分享这个项目给其他人

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 💬 联系方式

如有任何问题或建议，请通过 GitHub Issues 与我们联系。
