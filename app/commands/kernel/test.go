package kernel

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		fmt.Println("test start:", now)
		do()
		fmt.Println("test end, cost:", time.Since(now))
	},
}

func do() {
	time.Sleep(1 * time.Second)
}
