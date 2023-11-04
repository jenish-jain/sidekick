package twoFA

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\t2fa --add keyname\n")
	fmt.Fprintf(os.Stderr, "\t2fa --list\n")
	os.Exit(2)
}

func HandleCommand(cmd *cobra.Command, args []string) {
	listKeysFlag, _ := cmd.Flags().GetBool(List.Name())
	addKeyFlag, _ := cmd.Flags().GetBool(Add.Name())

	keychain := Init(filepath.Join(os.Getenv("HOME"), ".2fa"))

	if listKeysFlag {
		if len(args) != 0 {
			fmt.Println("no arguments supported with list flag")
			usage()
		}

		names := keychain.GetAllNames()
		for _, name := range names {
			fmt.Println(name)
		}
		os.Exit(2)
	}

	if addKeyFlag {
		if len(args) != 1 {
			fmt.Println("add command must be followed by name")
			usage()
		}
		name := args[0]
		if strings.IndexFunc(name, unicode.IsSpace) >= 0 {
			log.Fatal("name must not contain spaces")
		}
		_, _ = fmt.Fprintf(os.Stderr, "2fa key for %s: ", name)
		key, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatalf("error reading key: %v", err)
		}
		key = strings.Map(noSpace, key)

		if addErr := keychain.Add(name, 6, key); addErr != nil {
			log.Fatalf("error adding key: %v", addErr)
		}
		os.Exit(2)
	}

	if len(args) != 1 {
		fmt.Println("provide name to fetch 2FA code")
		usage()
	}

	code := keychain.GenerateCode(args[0])
	fmt.Println(code)

}

func noSpace(r rune) rune {
	if unicode.IsSpace(r) {
		return -1
	}
	return r
}
