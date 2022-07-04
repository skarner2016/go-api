package kernel

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var liveCmd = &cobra.Command{
	Use: "live",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		fmt.Println("live start:", now)
		handle()
		fmt.Println("live end, cost:", time.Since(now))
	},
}

func handle() {
	time.Sleep(10 * time.Second)
}
