### URI query design  
URI query는 POST 메서드를 사용한다.   
- 서비스에 데이터를 전달하기 위해, 데이터를 본문 내부에 포함시켜야함.
- 컨트롤러의 동작을 선택할 때는 쿼리 문자열을 사용할 수 있음.

ex) URI query
```json
POST /test?$group=admin
{
  "messsage": "example1"
}
```
주의할 점
- URI의 최대 길이가 2083 문자로 제한되어 있음.
- POST 요청은 일반적으로 항상 요청의 본문을 포함함.  

### 응답 코드
좋은 응답코드
- 표현 자체가 의미를 내포하는 응답
- 요청이 실패한 경우에도, 응답만을 읽어 추가적인 정보를 얻을 수 있는 응답.

HTTP 상태 코드
- https://ko.wikipedia.org/wiki/HTTP_%EC%83%81%ED%83%9C_%EC%BD%94%EB%93%9C

### HTTP 헤더
요청 헤더는 HTTP 요청 및 응답 프로세스에서 매우 중요한 요소
- 표준적인 접근 방식을 구현하면 사용자가 한 API에서 다른 API로 전환하는데 도움이 됨.  

HTTP 프로토콜에 대한 전체 정보
- https://tools.ietf.org/html/rfc7231

#### 표준 요청 헤더
요청 헤더는 요청 및 API 응답에 대한 추가 정보(동작에 대한 메타 데이터같은 것) 제공.
클라이언트가 서버의 응답을 처리에 도움이 되는 정보를 제공하기 위해 요청 헤더 사용.

##### Authorization(권한 부여) - 문자열
공개된 일기 전용 API라 하더라도 사용자에게 권한 요청을 요구하는 것이 좋음.  
사용자가 요청에 권한을 부여하도록 함으로써 사용자 수준 로깅 및 속도 제한과 같은 작업을 수행 가능.  
[표준 권한 부여 헤더 참고](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html)

##### Accept - 콘텐츠 타입
- application/xml
- text/xml
- applicatiton/json
- text/javascript (JSONP의 경우)

##### Accept-Encoding - gzip, release
REST 엔드 포인트는 가능한 경우 gzip과 deflate 인코딩을 항상 지원해야 함.

**[gzip 응답 라이터 작성](https://github.com/building-microservices-with-go/chapter2/blob/master/gzip/gzip_deflate.go)**  
핵심 : compresss/gzip 패키지 사용  
이 패키지를 사용하면
- Writer 인터페이스를 만들 수 있음
- 만든 Writer 인터페이스는 매개 변수로 입력받은 라이터에 gzip 압축을 사용해 데이터를 씀

#### 표준 응답 헤더
모든 서비스는 다음 헤더를 리턴해야함.
- Date : 요청이 처리된 날짜
- Content-Type : 응답의 콘텐츠 타입
- Content-Encoding : gzip or deflate
- X-Request-ID/X-Correlation-ID : 다운스트림 서비스(서버에서 클라이언트로 보내는 요청)를 호출할 때 요청에 추가할 수 있음. 하나의 트랜잭션 ID로 모든 요청을 그룹화하면 매우 유용하다. 이후 ELK 적용하기 쉬움.

#### 에러 리턴
표준 에러 엔티티는 클라이언트 또는 서버로 인한 에러가 발생할 떄마다 반복없는 코드를 작성할 수 있게 해 클라이언트를 도와야함.    
이러한 엔티티에 대해 Microsoft는 [우수한 API 가이드 라인 자료](https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md)를 제공.
