# uroboros

Небольшая либа, реализизующая что-то вроде rate-limit.

## Задача

Надо выполнять в многозадачной среде некие действия n раз в m  секунд. 
~~~go

u:= uroboros.New(n, m*time.Second)
...

if Can(time.Now()) {
  // Do something
} else {
  // пока нельзя
}

~~~

Подробности в тесте
