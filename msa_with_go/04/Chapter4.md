# 테스트
## 테스트 피라미드

https://www.google.co.kr/url?sa=i&source=images&cd=&cad=rja&uact=8&ved=2ahUKEwiz8dr8tf7eAhUU9bwKHWCgAkYQjRx6BAgBEAU&url=http%3A%2F%2Fteddy-chen-tw.blogspot.com%2F2017%2F03%2Fbdd21bdd.html&psig=AOvVaw03JvasctjspKisR7j5_ghg&ust=1543746473347898

피라미드의 위로 갈수록 테스트 횟수는 줄어든다.
시스템 내부를 알고 있다고 하더라도 수천 행의 코드를 샅샅이 뒤지고 수동으로 에러를 재현하기 위한 작업을 반복해야 한다.

## Go에서 좋은 단위 테스트를 작성하는 방법

테스트의 세 가지 법칙 from "클린 코드" 책
1.  실제 코드를 작성하기 전에 실패하는 단위 테스트를 먼저 작성한다.
2. 컴파일에 실패하지 않으면서 실패화기에 충분한 수준으로만 단위 테스트 작성한다.
3. 현재 실패하는 테스트를 통과할 수 있는 수준으로만 실제 코드를 작성한다.  

**Go에서 마이크로서비스를 테스트하는 가장 효과적인 방법 중 하나는 HTTP 인테페이스를 통해 모든 테스트를 실행하려고 하는 함정에 빠지지 않는것이다.**

net/http/httptest 패키지
- Go에 내장된 테스트 프레임워크
