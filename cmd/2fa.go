/*
Copyright © 2023 NAME HERE jenish.jain@rocketmail.com
*/
package cmd

import (
	"github.com/jenish-jain/sidekick/internal/twoFA"
	"github.com/spf13/cobra"
)

// 2faCmd represents the 2fa command
var twoFACmd = &cobra.Command{
	Use:   "2fa",
	Short: "cli tool to manage your internal codes",
	Long:  "This tool helps you manage your 2FA codes both hotp and totp",
	Run: func(cmd *cobra.Command, args []string) {
		twoFA.HandleCommand(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(twoFACmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 2faCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	twoFACmd.Flags().BoolP(twoFA.Add.Name(), twoFA.Add.ShortHand(), twoFA.Add.DefaultValue(), twoFA.Add.Usage())
	twoFACmd.Flags().BoolP(twoFA.List.Name(), twoFA.List.ShortHand(), twoFA.List.DefaultValue(), twoFA.List.Usage())
	twoFACmd.Flags().BoolP(twoFA.Clip.Name(), twoFA.Clip.ShortHand(), twoFA.Clip.DefaultValue(), twoFA.Clip.Usage())

}
