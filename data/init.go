package data

import "lightray/security"

// This file is made for holding variables (mostly constants, dynamic variables are defined where they are used)
// This also contains useful functions for starting stuff up

// Just because I'm adding comments does
// NOT
// mean that this is AI generated.

// Not constants:

// Networking:
// I'm using to do because JetBrains has a built-in thing for todos, it's not AI, I'm using GoLand
var (
	Port      = "8080"
	Username  = "admin"
	PwdLength = 16
	Password  = security.GenerateSecurePassword(PwdLength)
)
