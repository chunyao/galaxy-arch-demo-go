package main

//func main() {
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//	var signal struct{}
//
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(id int) {
//			mu.Lock()
//			defer mu.Unlock()
//			fmt.Println("goroutine", id, "is waiting")
//			wg.Wait()
//			fmt.Println("goroutine", id, "is signaled")
//		}(i)
//	}
//
//	fmt.Println("main thread is sleeping")
//	fmt.Println("press enter to signal all goroutines")
//	fmt.Scanln()
//
//	closeCh := make(chan struct{})
//	go func() {
//		for {
//			select {
//			case <-closeCh:
//				return
//			default:
//				mu.Lock()
//				signal = struct{}{}
//				mu.Unlock()
//			}
//		}
//	}()
//
//	fmt.Println("all goroutines are signaled")
//	close(closeCh)
//	fmt.Println(signal)
//	wg.Wait()
//	fmt.Println("all goroutines are done")
//}
