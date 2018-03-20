package exec_test

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

// 在.sh里面调用其它一直运行不退出的进程会导致cmd在context timeout 执行cmd.Process.Kill()时没将进程kill 成功,导致整个goroutine卡在那里
func TestCmd(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		bytes, err := exec.CommandContext(ctx, "bash", "start.sh").CombinedOutput()
		if err != nil && ctx.Err() == nil {
			fmt.Println("err:", err.Error())
			return
		}
		fmt.Println("ok:"+string(bytes), "\nctx.Err()", ctx.Err().Error())
	})
	t.Run("ok", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		bytes, err := exec.CommandContext(ctx, "sleep", "8").CombinedOutput()
		if err != nil && ctx.Err() == nil {
			fmt.Println("err:", err.Error())
			return
		}
		fmt.Println("ok:" + string(bytes))
	})
}
