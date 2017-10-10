flow
----
ВАЖНО: Не используйте это ;]

```go
import "github.com/rubaxa/flow"
```

### API

#### flow.DeferMany(size, task)
Выполнить операцию в гуртине и вернуть буферизированный канал для чтения результата

 - `size int` — размер буфера
 - `task func(ret flow.Stream)` — задача/callback, которую надо выполнить

```go
res := flow.Defer(3, func (ret flow.Stream) {
    ret <- "1"
    ret <- "2"
    ret <- "3"
});

fmt.Println("Result:", <-res)
```

#### flow.Defer(task)
Частный случая `flow.DeferMany`

```go
res := flow.Defer(func (ret flow.Stream) {
    ret <- "logic"
});

fmt.Println("Result:", <-res)
```


#### flow.CreateWorkerPool(size, worker)
Создать пул воркеров

 - `size int` — количество воркеров
 - `worker func()` — воркер/функция

```go
func main() {
    input := getInputChan();

    // Читаем канал в 15 потоков
    flow.CreateWorkerPool(15, func () {
        for raw := range input {
            // Чтение канала
        }
    })
}
```


#### flow.Go(queue ...) error
Последовательно выполнение тасков до первой ошибки

```go
    err := flow.Go(
        func() (err error) {
            // Первая задача
            return;
        },

        func() (err error) {
            // Вторая задача
            return;
        },
    )
```

---

### Types

  - `IntChan chan int`
  - `StrChan chan string`
  - `Stream chan interface{}``

