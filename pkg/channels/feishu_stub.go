// 未启用飞书时的占位实现，使 manager 在无 -tags=feishu 时仍可编译。
// 旧方案：无此文件，feishu.go 始终参与编译，强制依赖 larksuite 导致 MaxInt64 溢出。
//go:build !feishu

package channels

import (
	"fmt"

	"github.com/sipeed/picoclaw/pkg/bus"
	"github.com/sipeed/picoclaw/pkg/config"
)

func NewFeishuChannel(cfg config.FeishuConfig, bus *bus.MessageBus) (Channel, error) {
	return nil, fmt.Errorf("feishu support not compiled in, build with -tags=feishu")
}
