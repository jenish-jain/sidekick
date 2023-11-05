/*
Copyright © 2023 jenish jain HERE jenishjain@rocketmail.com
*/
package cmd

import (
	sms "github.com/jenish-jain/sidekick/internal/setupMySystem"
	"github.com/spf13/cobra"
)

// setUpMySystemCmd represents the setupMySystem command
var setUpMySystemCmd = &cobra.Command{
	Use:   "setUpMySystem",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sms.HandleCommand(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(setUpMySystemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setUpMySystemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	setUpMySystemCmd.Flags().BoolP(sms.Default.Name(), sms.Default.ShortHand(), sms.Default.DefaultValue(), sms.Default.Usage())
}
