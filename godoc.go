//Ha ha ...trust me,it was fun fixng this issue.

//The "godoc" command extracts and generates documentation for all locally installed Go programs,
//both from your own code and the standard libraries.
//As per documents available on net, you just need to :

godoc -http=:6060

//after it you need to launch the local godoc web-server from your browser, by:

http://localhost:6060

//which should in turn show you documentation of various Go elements.

[nabodip@GyanHouse ~]$ godoc -http=:6060 &
[1] 3348
bash: godoc: command not found
[nabodip@GyanHouse ~]$ 
[1]+  Exit 127                godoc -http=:6060
//Which of course happened as I did not have godoc installed on my system !

//However, the following bit eas executing fine,as it is supposed to show me documentation form the Go files I have, which I did 
//..not had in my 'pwd'

[nabodip@GyanHouse ~]$ pwd
/home/nabodip
[nabodip@GyanHouse ~]$ go doc
doc: no buildable Go source files in /home/nabodip
exit status 1

//Now, it was interesting to see the actual worth of "go doc" fromthe actualdirectory, housing my Go work till date:

nabodip@GyanHouse ~]$ cd /home/nabodip/student/golang/codes;
[nabodip@GyanHouse codes]$ ls -lrth *.go
-rw-rw-r--. 1 nabodip nabodip  75 Apr 28 05:22 main.go
-rw-rw-r--. 1 nabodip nabodip  77 Apr 29 07:05 main_error.go
-rw-rw-r--. 1 nabodip nabodip 157 Apr 30 05:05 import_example.go
-rw-rw-r--. 1 nabodip nabodip  95 Apr 30 05:23 import_play.go
-rw-rw-r--. 1 nabodip nabodip  95 Apr 30 05:25 import_play_1.go
-rw-rw-r--. 1 nabodip nabodip 107 Apr 30 05:26 import_play_2.go
[nabodip@GyanHouse codes]$ 
[nabodip@GyanHouse codes]$ go doc
doc: found packages main (import_example.go) and mains (main_error.go) in /home/nabodip/student/golang/codes
exit status 1

//at the same time, go doc command-line tool worked from command prompt:
[nabodip@GyanHouse ~]$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

//Anyhow, Godoc web-server will only display any sort of Documentation for your homegrown Go files, if and only if
//a regular comment, without any blankline directly precedes the function/variable/package/type/constant's declaration.

//So, just like other wise people, "install_godoc.go" was performed..
