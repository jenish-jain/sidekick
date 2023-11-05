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
	fmt.Fprintf(os.Stderr, "\nType 2fa --help to understand how this works!\n")
	os.Exit(2)
}

func HandleCommand(cmd *cobra.Command, args []string) {
	listKeysFlag, _ := cmd.Flags().GetBool(List.Name())
	addKeyFlag, _ := cmd.Flags().GetBool(Add.Name())
	clipCodeFlag, _ := cmd.Flags().GetBool(Clip.Name())
	isHOTPFlag, _ := cmd.Flags().GetBool(HOTP.Name())
	sevenSizeFlag, _ := cmd.Flags().GetBool(Seven.Name())
	eightSizeFlag, _ := cmd.Flags().GetBool(Eight.Name())

	keychain := Init(filepath.Join(os.Getenv("HOME"), ".2fa"))

	if isHOTPFlag && !addKeyFlag {
		fmt.Printf("%s flag can be used only with %s flag\n", HOTP.Name(), Add.Name())
		usage()
	}
	if (sevenSizeFlag || eightSizeFlag) && !addKeyFlag {
		fmt.Printf("%s or %s flags can be used only with %s flag\n", Seven.Name(), Eight.Name(), Add.Name())
		usage()
	}

	if listKeysFlag {
		if len(args) != 0 {
			fmt.Printf("no arguments supported with %s flag\n", List.Name())
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
			fmt.Printf("%s flag must be followed by name\n", Add.Name())
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

		size := 6 // default size is 6
		if sevenSizeFlag {
			size = 7
			if eightSizeFlag {
				fmt.Printf("cannot use %s and eight %s together\n", Seven.Name(), Eight.Name())
				usage()
			}
		}
		if eightSizeFlag {
			size = 8
		}
		if addErr := keychain.Add(name, size, key, isHOTPFlag); addErr != nil {
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
		fmt.Printf("%s flag only supported with keyname\n", Clip.Name())
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
