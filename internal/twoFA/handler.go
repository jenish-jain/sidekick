package twoFA

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\t2fa --add [-hotp] keyname\n")
	fmt.Fprintf(os.Stderr, "\t2fa --list\n")
	fmt.Fprintf(os.Stderr, "\t2fa --clip keyname\n")
	os.Exit(2)
}

func HandleCommand(cmd *cobra.Command, args []string) {
	listKeysFlag, _ := cmd.Flags().GetBool(List.Name())
	addKeyFlag, _ := cmd.Flags().GetBool(Add.Name())
	clipCodeFlag, _ := cmd.Flags().GetBool(Clip.Name())
	isHOTPFlag, _ := cmd.Flags().GetBool(HOTP.Name())

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
		return
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

		if addErr := keychain.Add(name, 6, key, isHOTPFlag); addErr != nil {
			log.Fatalf("error adding key: %v", addErr)
		}
		return
	}

	if len(args) == 1 && !listKeysFlag && !addKeyFlag {
		code := keychain.GenerateCode(args[0])
		if clipCodeFlag {
			fmt.Println("code copied to your clipboard 💥")
			_ = clipboard.WriteAll(code)
		}
		fmt.Println(code)
		return
	}

	// print all codes
	if clipCodeFlag && len(args) == 0 {
		fmt.Println("clip flag only supported with keyname")
		usage()
	}
	names := keychain.GetAllNames()
	for _, name := range names {
		fmt.Printf("%s: %s \n", name, keychain.GenerateCode(name))
	}
	return
}

func noSpace(r rune) rune {
	if unicode.IsSpace(r) {
		return -1
	}
	return r
}
