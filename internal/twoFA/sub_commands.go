package twoFA

import "github.com/jenish-jain/sidekick/internal"

var (
	Add  = internal.NewBooleanFlag("add", "a", false, "add a key \t\t\t[1 argument <name>  required!]")
	List = internal.NewBooleanFlag("list", "l", false, "lists all keys added")
	Clip = internal.NewBooleanFlag("clip", "c", false, "copy code to the clipboard \t[1 argument <name>  required!]")
)
