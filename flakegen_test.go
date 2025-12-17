package flakegen

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	t.Run("有效节点ID和服务ID", func(t *testing.T) {
		node, err := NewNode(1, 1)
		if err != nil {
			t.Errorf("创建节点失败: %v", err)
		}
		if node == nil {
			t.Error("节点不应为nil")
		}
	})

	t.Run("无效节点ID", func(t *testing.T) {
		_, err := NewNode(32, 1) // 超出最大节点ID
		if err == nil {
			t.Error("应返回错误，因为节点ID超出了最大值")
		}
	})

	t.Run("无效服务ID", func(t *testing.T) {
		_, err := NewNode(1, 32) // 超出最大服务ID
		if err == nil {
			t.Error("应返回错误，因为服务ID超出了最大值")
		}
	})
}

func TestNode_GetID(t *testing.T) {
	t.Run("生成唯一ID", func(t *testing.T) {
		node, err := NewNode(1, 1)
		if err != nil {
			t.Fatalf("创建节点失败: %v", err)
		}

		ids := make([]int64, 0, 10)
		for i := 0; i < 10; i++ {
			id, err := node.GetID()
			if err != nil {
				t.Errorf("获取ID失败: %v", err)
			}
			ids = append(ids, id)
		}

		// 检查ID是否唯一且单调递增
		for i := 1; i < len(ids); i++ {
			if ids[i] <= ids[i-1] {
				t.Errorf("ID不是单调递增的: %d <= %d", ids[i], ids[i-1])
			}
		}
	})
}
