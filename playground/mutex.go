package main
import (
"fmt"
"sync"
)
 
var counter int = 0             //  общий ресурс
func main3() {
 
    ch := make(chan bool)       // канал
    var mutex sync.Mutex        // определяем мьютекс
    for i := 1; i <= 2; i++{
        go work(i, 0, ch, &mutex)
    }
     
    for i := 1; i <= 2; i++{
        <-ch
    }
     
    fmt.Println("The End")
}
func work (number int, counterP int, ch chan bool, mutex *sync.Mutex){
    counter = 0
    mutex.Lock()    // блокируем доступ к переменной counter
    for k := 1; k <= 5; k++{
        counter++
        counterP++
        fmt.Println("Goroutine", number, "counterP -", counterP, "-", counter)
    }
    mutex.Unlock()  // деблокируем доступ
    ch <- true
}