package kernel

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display build info",
	Long:  `display build info`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Env:", 1)

		fmt.Println(time.Now().AddDate(0, 0, -30))
	},
}
