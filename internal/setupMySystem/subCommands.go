package setupMySystem

import "github.com/jenish-jain/sidekick/internal"

var (
	Default = internal.NewBooleanFlag("default", "d", false, "sets up your system with default configuration")
)
