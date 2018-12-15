# Exceptions
## Panic
*  사용자가 직접 에러 발생시킬 수 있다.

## Recover
* 기본 사용 문법
	* 패닉이 발생했을 때 프로그램이 바로 종료되지 않고 예외 처리를 할 수 있다.
	* 반드시 defer 함수로 사용해야 한다. (그렇지 않으면 프로그램이 바로 종료되어 버리기 때문)
	* 타 언어의 `try-catch`와 비슷하게 동작
* 책 코드 설명
	* ElasticSearch에서 인덱싱할 수 있도록, logrus에 필드를 추가해준다.
	* 자세한 정보를 담기위해, debug.Stack()를 호출하여 tracing된 정보를 담아 에러메세지로 포맷팅 해준다.

## 기타
* [Errors are values - The Go Blog](https://blog.golang.org/errors-are-values)
	* Great techniques for making your Go code more readable, while still robustly handling errors.
	* May use closures and record errors inside.	
