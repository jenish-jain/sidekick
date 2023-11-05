package twoFA

import "github.com/jenish-jain/sidekick/internal"

var (
	Add   = internal.NewBooleanFlag("add", "a", false, "add a key \t\t\t\t[1 argument <name>  required!]")
	List  = internal.NewBooleanFlag("list", "l", false, "lists all keys added")
	Clip  = internal.NewBooleanFlag("clip", "c", false, "copy code to the clipboard \t\t[1 argument <name>  required!]")
	HOTP  = internal.NewBooleanFlag("hotp", "H", false, "add key as HOTP (counter-based) key \t[1 argument <name>  required!, can be used only with --add]")
	Seven = internal.NewBooleanFlag("seven", "7", false, "generate 7-digit code \t\t\t[1 argument <name>  required!, can be used only with --add]")
	Eight = internal.NewBooleanFlag("eight", "8", false, "generate 8-digit code \t\t\t[1 argument <name>  required!, can be used only with --add]")
)
