package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.8" )

func wePvoFeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUasvgyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12Tcc75yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbHz1wRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRI6ZFhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRkDplBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGtqUJ7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhynoMKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7PWKTe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5hBBeWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXmzgdBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4Sr1R7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpYy8XrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1H10XrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUbPliNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgm9INd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uylweTkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3nBkO6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gsOZIygbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnitOMGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMChj7HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hZK4LHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vf14Np3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FbnqOq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfuNxSDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsVbTTcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iZe9MgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOdVkLabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nENVFuIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIO2HZI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AzYksWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQ8dWWCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mX3AMxjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILt8X78BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63EIGiDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8FXTGYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMoEEnPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K2ChNBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaBhAmDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3m4rNciYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IX1DYwzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQejwqz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWPqGOm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9rm3BBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8ZGEMtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjeoxuK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRsd9yuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04ZkhNfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6v1tWFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKarrxsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNHNvYclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQ36AUfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tobwVnJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86aIKGmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMFgjmEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASUW1qQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKvXhlEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t08MbpuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhHReBv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnDVklmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peBQt7NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owGBkxPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvlgICmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzGUWnjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgQqiNjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoAkIyXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGJR5iR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Is6mSVxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10quzfmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifMN6t7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrcqFUuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LE8H9pNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0R0lQc5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qzl1X1kLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgPtlMqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfaxKOfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZ45kWLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYUhy53zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7tb5iyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2w3RTurVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8e4TuMKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfGaX9RDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjI9ADUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFoqt6XHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KRcnncEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9hYZLfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8DsFPxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Znr8bR0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlEyfY5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9EwX3gLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ah2nRJa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ixkf0a3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhUMCulRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWpudzlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIeonu30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieBGlvufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmuLrzntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jd6Q4z8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztgygxA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1DiaJSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D97CPUhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func muXgjwLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bci0qedfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFzQatdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdORnZ3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fl9af1J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qi9kTea5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Px5N7TOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOBwOqGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5BySpkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVmRYxfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3J0mhtlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rngfLjCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6JOtMTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m38Fj7wOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIgG59OOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xDzvIwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJlWwl42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLjIqaSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fURCCaCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7P9jJmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dfb7zbFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmxRHYjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cTDG9PAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpzfzlSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3XYBpXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FN43hz6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDjCgc5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDfoggAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wade8Qa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2VD6Ah5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhWmbmEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uavA1LVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtPMg9hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVZruZdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ex9nHqpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFvlDlZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKMQ1SeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQhNIDicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJmzzFyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoxUCRZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWnns7ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYg593WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mcEillnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iraBgFTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWcPfqfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfurTXUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpdYWptjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfQMAiXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSHZtJpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3OuvRhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaNYpJEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgWrPMyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YssM6OuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QacLiJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLCwMg2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmtGuA23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGJdOGDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeeWA4KTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaWmVTXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Kq2fLnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uuF0IHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjjAGT6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVs7PGWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwkKYtJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkCVeSjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMxFuuR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6tnYTQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FC6i8LMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1ttlasWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ep0RmXRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apbzH0PpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uP9p4oDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYZCTa6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyZWC3iDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0RIdnrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l30IF5WhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nt6PjOfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDdu13yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dpC2AlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJOZLbHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLhnqqekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjTY52F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TxHuTDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDwmivkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmTP09XjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9IpJr1H6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acwQ74SeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8g7E3MrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwetkhgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2kNnc44SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHye9qMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Nc3KLXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lht5gadiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66ZppigDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 668qKRtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCFzoY9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UxYKyzzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iENUz02wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQ4urBDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvX90W3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SeV4Jc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2D6FZ6iQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mat3lJtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngxjPn3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ooRuzWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2kI2hp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKrAafZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgZ84jljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func os96FOkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adQ9AqkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irsqIeSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFd6PXOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgHGOpc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 901LG7iYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MiEKA4GdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuzytMZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaPrUSehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivmgPxJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05FPwTDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTeUaE24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wreCmjgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTEjqDGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQGc36LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLZYPoTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u69dMPGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTrZOhX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYncNKH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPTQ0ZERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqBXutn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZGrztBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOuB2gEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9ul0Tw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0H51yUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INerhDWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juSxZOxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbY3jyuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cs8884yQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0UtDa1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxcY0hofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxOEok0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVKeoaR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFztFHsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDiORZmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzVIKc5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q206SyMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qfL23OZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAj2sV5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZrekpVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlUDEp7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tmJYX67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3LOWsY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KVjRYbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDwKkhraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfQGvAsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDFcGvrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SA8RA7YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Io9EBAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uG4iW4a9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udhnfTTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOAFDyUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IblN4rcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYhRwiDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlvIPdLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVbt3daPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFOiEhidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oait90LUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WaGxal8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJAteHPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inW63VVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1kwYcn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRJheXx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJL9l7SLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd5vh4rJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZdfjzPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXdvnyQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVk5LUypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9Op17KVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNtdZdIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUMzuMNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXbGwinBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e03brF8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gScRwPW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hDIzkTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWjItl5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kh0YyVvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kF1ztCMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndAJ5omNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgWibRXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilgIkyUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Xr2FipiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISxhgC4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSpWAgADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COqNd3QqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3CSX76sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqTeV9YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdh8ZRcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYbRYMPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWSIHoWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfDRWuujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qqftOMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Adpi4DsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2smZoWPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8l3OMWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B105Be8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 342hfMahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggYsBSUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8LJTZrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Vnge6I1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b05QyisOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V95cVVuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4dRLEGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GP9zY69OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43zNoK18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gccINrbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gkoJj0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UexScI01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWfVFxZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJgWCNM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijwBArQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFqKmraSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItUGZkTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjuhrgilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0E4sAE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFQcz7HyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eN8aRb1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY2tCwCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNFDahrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7R82qyFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tx7D8suIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37JeXsJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZ5WwtoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwIzxtlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckcz9h5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHlIUJwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPZGoblxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dYR1dGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iN61qdJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zog5ckfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wOPAkpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Adqw7omgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 765YDYGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FL4QwrEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nf7j8QfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruJQnPKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HZ96YeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouPJn2jfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AloeKWcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPkuFj9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gCGT09CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WI9oOAhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HboTis51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUez80BgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDUgSAocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4czv20xHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4vOtA6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wzjzv9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgx68F04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSloBncAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCKhIF6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9FV7S6GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5s7qusvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rACXVKtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXGEiXFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TVKO8vkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZN96sElIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfyN1aVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeljdPhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpW2DPN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsprUrlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EIa9OFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zr5qUjyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSw3FntUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kamkk2OTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmCdpTj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BVWEx0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqcQHD5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRkIYEk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1XemTxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmtfijhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n81oitygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2eTe8avWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDjmETcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cwmmqBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tIoTkS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inKWLa9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5g1jx74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYDxdRWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skCSPZsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoOg6ig7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rg9wpOJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUcUFcs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X47t0AdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ir3r70VPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kt4GmIJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKVmGP0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzCIWsC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3llWW0tlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDL6w7edWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiCQK8wTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyO0agIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1Py5hrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDCCzQQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXKcPD3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6dcd2KpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BxitlglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOJXyTJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZ5sNvfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIXzRUZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7L9oS1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkGzvz9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPR8Wju9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2ekZ8o0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Rt6IMqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqMcCgsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIQdS0sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEQRdhlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1x2tyWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVH0zzovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ea94NO4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFXlJY1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tP2deX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ztgmdy1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqme0dJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vit5AzK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xkk87GvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vembXOE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJRuO9qmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFE33F0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DK5OBQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kvmb80gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9S8w7PZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YgERgWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6JX6YToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4vsRMHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2arurU0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Awx5EcD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QT8LJA5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjcmyBHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyS5f3ebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4t61U5QcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfpsRiTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJXHGyCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44I6UniHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsLm0wRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GniWcOxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZP4EETuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXpZke4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbE9TTWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ME6NmIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBxwbh35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcuEZgKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZcZ07enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzxH7AjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mrx5PX5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJDKP7J4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JSd17usWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTrmYtApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNR8wVdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVkKRPE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OP4jZKtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NELnHM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjYISvw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9luJov8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQbc1dsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFqybagtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHZ1WvyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dfqF859Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTPllq5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnTOP3itWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWWAC7hUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MenQqvRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6MxgrmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBO0zY7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8QYZybDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xftqMkU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNJitfQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRxbOc5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Wo2R3cjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2F4fVfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SXoPamZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhFOZg37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzGsqpGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWnPFVD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDWcf4tIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71dX0mUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wklIq59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mz8yql5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtJgG7YyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soCsOiPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNokSMXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YnY3vvqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiWZeQ2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k03Sdc3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JsuhnyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xWDxChAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sCeJGfrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCQTRmTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njolTIiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNbIsomeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmkZP83WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZNMxEBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2h0LqCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JauEritQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9LcRAi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDayy9tgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUBhwYnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaKxw2IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J04NyXVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eD2AI0gDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Of3P1NSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVuGDhrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4WGhFrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DL7xX8D7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1I2aQcHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tlrCHjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3uIHXhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JnSsq0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYOw3DT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyXt50q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3v7Ybn6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2bgPrpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmc0TwZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbzge5vzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fY5aNDGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xz8768UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt67ONW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHvj9k8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQPBjbhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6n0YSqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQMUcubkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P59V6KTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LdQ8WlKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vlnu5AodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgWy3hHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bYzGF8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbDluTTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPC6ePBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IC2r4OC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func br39wbyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n596LTCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTTmywWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAbQ4OAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jinp5SLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVl44uDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func es8KrAVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wIC3RENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYXcOxeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZBwue70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MHVP1rxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svwbke51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zMIkhQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3D591RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9RdGFO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tpw4JEo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIk4Bi71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULM6TlsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiqnssw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bb4o0SeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPBKQWfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwthad2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48zHmvG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtwuojDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaAxlCe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrEKURPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuxWrbjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMOAeX1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUry6lpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6aXUP0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjmaezuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func du5kSPSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFhUle8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DS5VmCQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8pAy6sWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UxP9bfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZcShFbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4c7yKT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQ2IwAquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocGjLi4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8SqgXiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMJFO9H8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt5C6aiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZM2v1YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEirL0rVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vK9XfgeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQ2SyVcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uSKDVpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOjWf9TWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0jaI4C3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIpQJ51TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUAF3eWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZT3feS8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vg72IFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRNWbAMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZhEBOJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4G3AICvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESD2nKajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O2q8Nw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vg94XBTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xn18ZZufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2MU9ndtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWgWd3KCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3ouMfoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFInlEMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWumMnzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLLsXahDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjJof2LHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCgHPQ2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVHF4VGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKEyoaOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqVhMFYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKuNjoQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d55qMnYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUM1p0x0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KmAtLfZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gK3AXExmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IPudH6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7REBTTp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCTM527sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFDqxediWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojjqHE5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qSyYHmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e8RMxO7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTsIbq15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcOsUb1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhHmGXLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZxfehexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjtWdWy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEQv588sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeeKAusvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbnbrNzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSn8pFinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhWulqwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMiClihuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEnQlqGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkeUpiWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utqxNtOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bnLxqCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zUGrqipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ABnWUnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAThhV4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQnshq81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6Hd72x8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hFlLBDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEM8dQJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZN74aC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHvhEhe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRfP5WuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlVKBTu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8Vhp8TpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNu83oNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cr2L9ovkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X3nkC3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GK6G9kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFGyFLqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcoqiSTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50fSbIXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6v0MsYNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dk9VcgSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FS2rrTeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pT56gxbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3fRUGjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCvS2gj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUPbrVy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7bXDKEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPPuI2clWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0h4gCVXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAG0CC7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWCad5lAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOr0kxdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92ILj587Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuQWkZoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABi6E1N3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvxh3n9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmFiREHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHuANjeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oJSCCkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ey1X57SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4xROsEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cHCvTjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJnjV6tYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRPYKgdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOxHVM58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXQyzzlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teySbMKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfHE5jgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xqaq7cTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDvsnwZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e0IZAFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTkwpontWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0KeNKYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHZln4xIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hDoNYKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rkzHIQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VG4cFqF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRx8i7FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wM5ZNVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SEjh8PFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1dIDxnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwgzktBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fewbu4pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqfos2LlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIoKc7QNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9MJl7M7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQ5nD4ABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IKcJ4k4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZ1pBh0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjRSoForWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UOr8QhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqbB9AMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8romZnL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNTtZFpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVWXdxXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkr4L8uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLhtv9RnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gCLtSNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjSvafZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVba5GIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4zLdVwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dM9j3G3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BvvkdDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXFlHveGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gm5FxUXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITnUa2BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvwXlDauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyUcyasRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhPCcsPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSaQnXeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fvyq73vuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njgIIdowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CPAlovJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IevcU7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZ7Vs2x7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiFs4o4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9dgP233Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5yxEoIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WOVXBfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SUi9x00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WLAbRsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKjLnVwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IExtqkupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8Ov8ZvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeMiCAoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWvUBgCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QStcV1j1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxeiCDnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fW25uiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs6ZlarqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWHUqv4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ulyayKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RIqp0jcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2MwWDvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PefGvDahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAisBcy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIkxJqhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vJ0rLJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ca0uOcE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZHqB0GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xB4PbDo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCGhRAqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQDkWfuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adB5sLEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRpYvZgbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzTqTieYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYZdBaZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFX1KH2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIea3mpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzWvtQtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Uk6ETwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DwxbynKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URGa0bisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yLhfy0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eU7KrOvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBcR0hKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcuEWcb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVA0mdj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cg7aOTR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqfcesIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HnRYK1H9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7cMMkE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I2XtxlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raNR0mZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mZCOYsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IiaEkmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHi4F6s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khH87cn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwvjwexMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhQ0nSXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCQFcjTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPxY0jHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjUVj534Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6f9Sq7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiO7crXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGrpQ0nRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlY0L9LFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zj9eWFOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iie3pPk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qr23W7SYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0esZUcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdYPVhhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKPatjUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNlvcsuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFbjFp5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmXM4Q41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Dst7memWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XYoxFCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGW8Jqm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjlI26HrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6Zrj8HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9xm8eNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74UoX55eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dp1UyKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zv7dSfnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1cV2ehJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHX6MsMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xGFc2rqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fYHZfqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQZp1ksaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEIDtNrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oIh3usPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JnkGRXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KScZLz6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBerWTPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlEPFb9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNLBycxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuvLEMoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MECaTswzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QXHbyz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OKQdxwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0TTCEWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mj1lAPXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lj93D9bbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8U6v3kpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvjWDJsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMXgCu85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1zB56dRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func muLm8i5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKi9hZSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qy74Vh8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTy59IR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHjCI1zeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwTkI0GMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqImp55gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plAWjKPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khhuwtl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTDju5UrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jC03xk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z32wSPjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGVVO2PXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2lyaXoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osZ8e0k9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Wq48DflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QN7uZdQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9TbeprOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWbAqTTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHjMdYAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FzD61V7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJsuKsvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3Ch7x2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jD5fEyY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LgtIInjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAEFQS5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9oiiTcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hf4jpkFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpJwWkFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWITj2qQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wmj7302iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1G8cnEotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjcVdj5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEREUEFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZptKE51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9GtwEmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKG1KTA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9XMePIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtcyP1PEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTiZOyjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfjheNQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8WBRWj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69YgbSbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNI08CrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRRCt5mPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6vIDxUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8g5fpSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQFmuVZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uS6AVG23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hv4XftMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D85LUdkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BTjGFiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhNjZCeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8m42wjW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MaOciGhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSxdrWVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXSIJY6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAC2K1SMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDeWlwQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlGTMiMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAfmfRDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cc0qpUfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDKVbIbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8BRILduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEGjwnu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eulmQwG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsRRtJE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk1SV1vnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDMCxxV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whTBQZyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HoiZeOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQRyVVcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7e81gqToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnXVLMVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUCS6P7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPfwwPuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51oGF2VGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6nxWL7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48zfa42uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzV0Bk1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSrG9UdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8u8B51cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PYrBcTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XeZAvo54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZjse0LbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tibqwAWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAm5vB97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtQ6aWtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVKYOUmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmseZR3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5l8QmhzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTQW277DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UH97RScwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJq3YunrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f43GfCaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RWEW2l0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHjk3TbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrxnTOu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYxxFevnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emrWZZBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjNM1T44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r98hxJooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17ZTvxztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3spU8FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXbYHpvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d20e07DvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjIKaZ3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ATheRAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7v5aK5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VABFQvv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPp32It4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWdACqenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaAFTrAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZFKmrEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sy4wO2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APnmSIagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRGqzyfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YylMYLI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoGhF8k2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhSKSBfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rOt0x8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlhhMhWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnUSrpMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7q9D8x5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZStly7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViY6RtujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkVPVzGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvIlQhqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIbcRBRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yu9YbGJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRdDUWKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlDrpVgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThwMznPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMI9Yg1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gvpa07fgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mY5tFPiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fiaEE5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiIOJfHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJXG9ee8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgLz1EqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfEfhnlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdk62alQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRRKn44iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIQZBoXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2Za4EEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vC75h7PuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekJDg9KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Oiy3Y7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tTqHcOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sXn8QfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxWqI9o6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3Ql2ILNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1v0YyqyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGpWJDGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjFG7tQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYcfMgoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ru4rElYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8uQ8hSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jlBst42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0p5D1pilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meO1PGhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8b1uqNCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcFHiJ3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6mW9gUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YdcwYZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxiW7u4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XvEu1etWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNbVAaUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmB6Fb3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtUCQvfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0SDW12aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZhqpMc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9FBnyd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkUJYGY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0B5jr69EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6l6e3mPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQJPJcjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rAEtfLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEJymxIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PpLp06LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ol8mJygaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3r2NLmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqaIWlJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaOck8uOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfNji2tAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MppR5d7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPvY8E1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aP6vxh2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMOtOmo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2BOArprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAZ0CcX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFpdm7zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ7gLvQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sj3zhTT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XT9PU4pdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57baACySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HODQmtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41Ay4AItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBV2PF48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxBbkBHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKM3C9DaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucnlZ6lKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQrMGdvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xliamGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ElgniDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgRlnoa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YflCW2nrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ce4Vvg9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrQMI0CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccyqMYgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIZXrAwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iLDU0DgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8smeWUOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvS6ikrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNWyztKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMKJei1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkvfbbMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2m8h7q6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUakPtMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1U42lPLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OyDxideWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21vYdssRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKLEAvPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2zl91GXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNxKtYxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5mkq8neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wP2In4sIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4v1bdfAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEND9CotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YthVFPduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxDI5P4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wG7ydvYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPRSnHAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNcbY6UdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgDPKhPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzEdyOrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANn95FkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ft8z1xm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuTcYk9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ud1k8TQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsga4c7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7otUKdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tga7Scm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SE3plBzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVADXEziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZU16ef6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1x48hAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Pl4cQUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZAEumXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lafUdqgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxLxotNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jana3NaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f15lpLhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olDB5vdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdEFZ2PWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNwNcF8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTwSviU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrYativ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNdue29nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIT2Ac6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giepYyUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebqC8XPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nbjjcka9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H993XMdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAleIqCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AxWXzgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUbNGNCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujtYT8GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3pYFempWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOfCsfOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfmGF7xKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoUGARRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9p4jMiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRfBSrfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlHUMzAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rw2hy81QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMktL8znWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Na7imUPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qa016DxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7pfXMGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9SIJA5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOxZxuwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvRsMIKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oq8AhPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XKt6kI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtkfZCpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXreKJTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36NFYOlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PczoJMNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iy1C76YsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNl4ehHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66PJylF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r94gPWy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bKkMq22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01Zhu17MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMnNmoj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVWbI6CsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgHpidyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFHdjr91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

