package main

import "fmt"

//import "time"
import "sync"

var wg sync.WaitGroup

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	wg.Add(2)

	go func() {
		//	time.Sleep(time.Second * 1)
		c1 <- "one"

	}()

	go func() {
		//	time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//wg.Wait()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		}
	}
}

/*

received one
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select]:
main.main()
        /root/go_repo/src/select-1.go:18 +0x41e
exit status 2

Press ENTER or type command to continue
received one
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select]:
main.main()
        /root/go_repo/src/select-1.go:18 +0x394
exit status 2

Press ENTER or type command to continue
# command-line-arguments
./select-1.go:6: undefined: sync in sync.WaitGroup

Press ENTER or type command to continue
received one
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select]:
main.main()
        /root/go_repo/src/select-1.go:23 +0x3bd
exit status 2

Press ENTER or type command to continue
received one
received two
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x51ad0c)
        /usr/lib/golang/src/runtime/sema.go:47 +0x34
sync.(*WaitGroup).Wait(0x51ad00)
        /usr/lib/golang/src/sync/waitgroup.go:131 +0x7a
main.main()
        /root/go_repo/src/select-1.go:37 +0x2a3
exit status 2

Press ENTER or type command to continue
received one
received two




*/
