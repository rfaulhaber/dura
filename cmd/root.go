package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"time"
)

var (
	unitFlag string
)

var rootCmd = &cobra.Command{
	Use:   "dura",
	Short: "Program for timing the execution of other programs",
	Long:  `Program for timing the execution of other programs`,
	Run:   runMain,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&unitFlag, "unit", "u", "", "unit of time for output. Can be one of: ns, us, ms, s, m, h")
	// TODO add format flag?
}

func runMain(cmd *cobra.Command, args []string) {
	execCmd := exec.Command(args[0], args[1:]...)


	execCmd.Stdout = os.Stdout

	start := time.Now()

	err := execCmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dur := time.Since(start)

	var output string

	switch unitFlag {
	case "ns":
		output = fmt.Sprintf("%dns", dur.Nanoseconds())
	case "us":
		output = fmt.Sprintf("%fμs", float64(dur.Nanoseconds()) / float64(time.Microsecond))
	case "ms":
		output = fmt.Sprintf("%fms", float64(dur.Nanoseconds()) / float64(time.Millisecond))
	case "s":
		output = fmt.Sprintf("%fs", dur.Seconds())
	case "m":
		output = fmt.Sprintf("%fm", dur.Minutes())
	case "h":
		output = fmt.Sprintf("%fh", dur.Hours())
	default:
		output = fmt.Sprintf("%fs", dur.Seconds())
	}

	fmt.Printf("time elapsed: %s\n", output)
}
