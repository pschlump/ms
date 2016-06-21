// code for BSD

// +build darwin dragonfly freebsd netbsd openbsd

package ms

import (
	"log"
	"os/user"
)

// OS related Support Stuff ----------------------------------------------------------------------------------------------

// HomeDir returns the home directory of the current user
func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println( usr.HomeDir )
	return usr.HomeDir
}
