// wpconfigperm.go
//
// Find wp-config.php files and set safe file permissions
//

package main

import (
    "flag"
    "fmt"
)

var basedir = flag.String("d", "/home", "Base directory to search")
var perms   = flag.String("p", "0600", "Permissions to use (numeric)")

func main() {
    flag.Parse()
    fmt.Printf("Searching users in %s, using perm mask %s\n", *basedir, *perms)
}

