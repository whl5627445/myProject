package convert

import (
	"github.com/bytedance/sonic"
	"google.golang.org/grpc/grpclog"
)

// S2S 结构体转结构体, 当两个结构体拥有相同字段时使用，否则数据全部为空
func S2S[T1 any, T2 any](source T1, Target T2) {
	t, _ := sonic.Marshal(source)
	err := sonic.Unmarshal(t, &Target)
	if err != nil {
		grpclog.Errorf("数据转换失败: %s", err)
		grpclog.Errorf("源数据：%#v", source)
		grpclog.Errorf("源数据json：%s", t)
		grpclog.Errorf("目标数据：%#v", &Target)
	}
}
