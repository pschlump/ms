// code for linux

package ms

import (
	"log"
	"os/user"
)

// HomeDir returns the home directory of the current user
func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
