package setupMySystem

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func HandleCommand(cmd *cobra.Command, args []string) {

	InitToolBelt()
	defaultFlag, _ := cmd.Flags().GetBool(Default.Name())
	var tools []string
	if defaultFlag {

		for _, tool := range GetToolBelt() {
			if tool.isDefault {
				tools = append(tools, tool.name)
			}
		}

		arguments := append(append([]string{}, "install"), tools...)
		//fmt.Println(arguments)
		command := exec.Command("brew", arguments...)
		command.Stdout = os.Stdout
		if err := command.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	}

}
