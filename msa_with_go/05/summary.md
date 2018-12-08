# Common Patterns
## Timeouts
* 외부 API와 통신하는 경우는에 유용하다.
* 핵심은 빠르게 실패하고, 실패시 요청 당사자에게 실패 여부를 빠르게 알리는 것
* 여러 종류의 Timeout 존재
	* Connection Timeout
	* Request Timeout: 보통 더 오래 걸린다.
* Timeout을 서비스의 config에 세팅해 두는 것이 좋다.
* 추천 Timeout 툴: [go-resiliency/deadline](https://github.com/eapache/go-resiliency/tree/master/deadline)

### 기타 참고
* [예제로 배우는 Go 프로그래밍 - 채널 타임아웃 기능](http://golang.site/go/article/211-%EC%B1%84%EB%84%90-%ED%83%80%EC%9E%84%EC%95%84%EC%9B%83-%EA%B8%B0%EB%8A%A5)
특정 시간이 지나면 현재 시간이 담긴 채널이 리턴됨 -> 일종의 받기 전용 채널
* [Go by Example: Timeouts](https://gobyexample.com/timeouts)
* [How To Gracefully Close Channels - Go 101 (Golang Knowledgebase)](https://go101.org/article/channel-closing.html) - 소스 코드 이해에 도움이 된다.
	* close 함수로 채널을 닫았을 때도 case 실행된다!
	* `don't close a channel from the receiver side and don't close a channel if the channel has multiple concurrent senders.`
* [concurrency - Golang : anonymous struct and empty struct - Stack Overflow](https://stackoverflow.com/questions/20793568/golang-anonymous-struct-and-empty-struct)
* [go - Do goroutines with receiving channel as parameter stop, when the channel is closed? - Stack Overflow](https://stackoverflow.com/questions/53673690/do-goroutines-with-receiving-channel-as-parameter-stop-when-the-channel-is-clos)

## Backoff
* Request가 실패하면 정해진 시간 동안 기다린 후 재시도하는 것
	* client-called API 보다는, 메세지 큐를 처리하는 worker 프로세스에 효과적
* 참고 패키지: [go-resiliency/retrier](https://github.com/eapache/go-resiliency/tree/master/retrier)

```
r := retrier.New(retrier.ConstantBackoff(3, 100*time.Millisecond), nil)

err := r.Run(func() error {
	// do some work
	return nil
})

if err != nil {
	// handle the case where the work failed three times
}
```

## Circuit breaking
* 에러가 났을 때 경중을 따져서 더 중요한 UX를 위해, 사소한 부분의 UX를 희생할 수 있는  결정을 내리는 것
* 참고 패키지: [go-resiliency/breaker](https://github.com/eapache/go-resiliency/tree/master/breaker)
	* error threshold 초과 -> breaker open   -> request 시도 불가능
	* timeout 시간이 지난 후 -> breaker half-open
		* -> request 시도가능, 하지만 실패시 바로 breaker open
		* -> 성공적인 requests들이 success threshold에서 정해진 수 이상이 되면 다시 break close (초기로 돌아감)
* [넷플릭스 circuit breaking & timeout 패키지](https://github.com/Netflix/Hystrix)  - Java
	* [GitHub - afex/hystrix-go: Netflix's Hystrix latency and fault tolerance library, for Go](https://github.com/afex/hystrix-go) :  Go 구현체
	* `go-resiliency`보다 더 깔끔하며, [hystrix dashboard]([Home · Netflix-Skunkworks/hystrix-dashboard Wiki · GitHub](https://github.com/Netflix-Skunkworks/hystrix-dashboard/wiki))에 여러가지 수치들을 자동으로 보여준다. (chapter7 logging 파트와 관련)
