// Go does not allows unused imports.i.e.you can not import a package,and later not use it.

// Play#1 :
// What happens when you import a package and do not use it ?

//Code:
package main

import (
	"fmt"
	"os"
	)
	
func main () {
}

//Output:
[nabodip@GyanHouse codes]$ ls -lrt import*play*
-rw-rw-r--. 1 nabodip nabodip 81 Apr 30 05:21 import_play.go

[nabodip@GyanHouse codes]$ go run import_play.go
# command-line-arguments
./import_play.go:6: imported and not used: "fmt"
./import_play.go:7: imported and not used: "os"

[nabodip@GyanHouse codes]$ go build import_play.go
# command-line-arguments
./import_play.go:6: imported and not used: "fmt"
./import_play.go:7: imported and not used: "os"

// Play#2 :
// What happens when you do not import a package,which was supposed to be imported, and use it ?


package main

func main () {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Love you",os.Args[1])
}


//Output:
[nabodip@GyanHouse codes]$ ls -lrt import*play*
-rw-rw-r--. 1 nabodip nabodip  95 Apr 30 05:23 import_play.go
-rw-rw-r--. 1 nabodip nabodip  95 Apr 30 05:25 import_play_1.go
-rw-rw-r--. 1 nabodip nabodip 107 Apr 30 05:26 import_play_2.go

[nabodip@GyanHouse codes]$ go build import_play_2.go
# command-line-arguments
./import_play_2.go:4: undefined: os in os.Args
./import_play_2.go:5: undefined: os in os.Exit
./import_play_2.go:7: undefined: fmt in fmt.Println
./import_play_2.go:7: undefined: os in os.Args

[nabodip@GyanHouse codes]$ go run  import_play_2.go
# command-line-arguments
./import_play_2.go:4: undefined: os in os.Args
./import_play_2.go:5: undefined: os in os.Exit
./import_play_2.go:7: undefined: fmt in fmt.Println
./import_play_2.go:7: undefined: os in os.Args

[nabodip@GyanHouse codes]$ go import_play_2.go
go: unknown subcommand "import_play_2.go"
Run 'go help' for usage.

[nabodip@GyanHouse codes]$ go import_play_2
go: unknown subcommand "import_play_2"
Run 'go help' for usage.

