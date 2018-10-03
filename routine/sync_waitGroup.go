package main

import (
    _	"fmt"
    "net/http"
    "sync"
)

/*
例子代码如下：

同时开三个协程去请求网页， 等三个请求都完成后才继续 Wait 之后的工作。
*/
func main(){
		var wg sync.WaitGroup
		var urls = []string{
			"http://www.golang.org/",
			"http://www.baidu.com/",
			"http://www.taobao.com/",
		}
		for _, url := range urls {
		// Increment the WaitGroup counter.
			wg.Add(1)
		// Launch a goroutine to fetch the URL.
			go func(url string) {
				// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
			}(url)
		}
		// Wait for all HTTP fetches to complete.
		wg.Wait()

}
