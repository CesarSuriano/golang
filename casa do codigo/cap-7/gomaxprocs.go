package main

import(
    "fmt"
    "math"
    "sync"
    "time"
    "runtime"
)

func calcular(base float64, controle *sync.WaitGroup){
    defer controle.Done()
    n := 0.0

    for i := 0; i < 1000000000; i++{
        n += base / math.Pi * math.Sin(2)
    }

    fmt.Println(n)
}

func main(){
    runtime.GOMAXPROCS(4)
    inicio := time.Now()
    var controle sync.WaitGroup
    controle.Add(3)

    go calcular(9.237, &controle)
    go calcular(6.64, &controle)
    go calcular(4.57, &controle)

    controle.Wait()
    fmt.Printf("Finalizado em %s. \n", time.Since(inicio))
}