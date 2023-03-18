package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var config string
var rootCmd = &cobra.Command{
	Use:   "ec [sub]",
	Short: "My root command",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(config)
		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

/*

run 执行顺序
PersistentPreRun
PreRun
Run
PostRun
PersistentPostRun
*/
var subCmd = &cobra.Command{
	Use:   "sub [no options!]",
	Short: "My subcommand",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 子命令不写 PersistentPreRun 方法 则调用父命令PersistentPreRun
		fmt.Printf("Inside subCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config)
		fmt.Printf("Inside subCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	subCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	//rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	//rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}
func main() {

	rootCmd.AddCommand(subCmd)

	//rootCmd.SetArgs([]string{""})
	rootCmd.Execute()
	//fmt.Println()
	//rootCmd.SetArgs([]string{"sub", "arg1", "arg2"})
	//rootCmd.Execute()
}
