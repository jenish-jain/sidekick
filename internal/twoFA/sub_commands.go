package twoFA

import "github.com/jenish-jain/sidekick/internal"

var (
	Add  = internal.NewBooleanFlag("add", "a", false, "add a key")
	List = internal.NewBooleanFlag("list", "l", false, "lists all keys added")
)
