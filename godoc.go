//Ha ha ...trust me,it was fun fixng this issue.
//As perdocuments available on net, you just need to

godoc -http=:6060

//after it you need to visit the following from your browser:

http://localhost:6060

//which should in turn show you documentation of various Go elements in its browser.

[nabodip@GyanHouse ~]$ godoc -http=:6060 &
[1] 3348
bash: godoc: command not found
[nabodip@GyanHouse ~]$ 
[1]+  Exit 127                godoc -http=:6060
//Which of course happened as I did not have godoc installed on my system !

//However, the following bit eas executing fine,as it is supposed toshow me documentation form the Go files I have, which I did 
//..not had 
//..till the time this post was being written..

[nabodip@GyanHouse ~]$ go doc
doc: no buildable Go source files in /home/nabodip
exit status 1

//while the following worked from command prompt:
[nabodip@GyanHouse ~]$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

//So, just like other wise people, "install_godoc.go" was performed..

