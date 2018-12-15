# 로깅 및 모니터링

서비스의 운영 환경 및 서비스가 겪고 있는 부하 수준을 이해하기 위해서는 서비스에 관한 유용한 **데이터** 가 필수적이며, 이 데이터를 활용하면 시스템이 최고의 성능을 낼 수 있도록 시스템을 미세 조정할 수 있다.

위에서 언급된 데이터에 대한 예시들로는 다음과 같다.
1. 측정 지표
  - 시계열 데이터 항목
2. 텍스트 기반 로그
  - 애플리케이션이 출력하는 구식 로그, 사용자가 만든 애플리케이션 소프트웨어의 텍스트 로그
3. 예외

#### 로깅 모범 사례

[존 기포드가 쓴 무료 전자 책 "실용적인 로깅 핸드북"](https://www.loggly.com/wp-content/uploads/2015/02/loggly_ebook_Pragmatic_Handbook_online.pdf)에서  
로깅 전략을 결정할 떄 8가지 모범 사례 제안

1. 애플리케이션 로깅을 진행 중인, 반복되는 프로세스로 간주한다. 높은 수준에서 로그를 남기고 더 자세한 계측 정보를 추가한다.
2. 분산 시스템에서 문제가 발생하면 전체가 제대로 동작하지 않으므로 프로세스의 모든 부분을 측정한다.
3. 기대 이하의 성능을 나타내는 모든 것을 기록해야 한다. 예상되는 범위 밖에서 동작하는 모든 항목을 기록한다.   
4. 하나의 로그 이벤트에서 일어난 일에 대한 전체적인 그림을 얻기 위해 컨텍스트 정보를 가능한 충분히 기록하도록 한다.
5. 시스템의 최종 소비자가 인간이 아닌 기계라고 생각해야한다. 로그 관리 솔루션이 해석할 수 있는 레코드를 만든다.
6. 하나의 데이터보다는 데이터의 경향성(트렌드)이 많은 것을 알려준다.
7. 계측이 프로파일링(동적 성능 분석)을 대신할 수 없으며, 반대의 경우도 마찬가지이다.
  -  추가 설명  
  애플리케이션에 높은 수준의 로깅 및 모니터링이 갖추어져 있다고 해도 프로세스를 배포하기 전에 애플리케이션 코드를 프로파일링하는 작업은 필요하다는 것이다.
8. 아무것도 모르는 상태로 시스템을 운영하는 것보다는 천천히 운영하는 것이 낫다. 따라서 측정을 할 것인지 말 것인지는 논쟁의 여지가 없으며, 얼마나 하느냐가 문제다.

### 측정 지표
저자의 개인적인 생각으로 측정 지표는 일상적인 작업에서 가장 유용한 로깅 형식이라고 한다.
그 이유는 간단한 숫자 데이터이기 때문이다.

#### 측정 지표로 가장 잘 표현되는 데이터 타입
요청 타이밍 및 개수와 같이 간단한 숫자로 표현될 때, 의미있는 데이터이다.

#### 명명 규칙
[점 표기법](https://docs.signalfx.com/en/latest/reference/best-practices/naming-conventions.html)을 사용해 서비스 이름을 나눌 것을 권한다.   
이 방법으로 측정 지표에 이름을 붙이면 에러를 상세 수준의 그래프로 표시하거나 상위 수준의 쿼리# 작성할 수 있다.  

#### 저장소 및 조회
측정 지표 데이터를 저장하고 조회하는데 여러 가지 옵션이 있다.
1. SaaS를 활용하고 있다면
2. 자체적으로 호스팅하는 서버를 가지고 있다면  

이를 관리하는 방법은 회사의 규모와 데이터 보안 요구사항에 따라 다르다.

##### 1. 서비스 방식의 소프트웨어(SaaS)
[Datadog](https://www.datadoghq.com/dg/lpghp/?utm_source=Advertisement&utm_medium=GoogleAdsNon1stTierBrand&utm_campaign=GoogleAdsNon1stTierBrand-NonENES&utm_content=Datadog&utm_keyword=datadog&utm_matchtype=e&gclid=EAIaIQobChMIhrzAoPae3wIViraWCh2Sww3PEAAYASAAEgIXYfD_BwE&utm_expid=.40ES1QyrT6Sm1I5PXMEjqQ.1&utm_referrer=https%3A%2F%2Fwww.google.com%2F) 추천한다.   
Datadog으로 측정값을 보내는 방법은 두 가지가 있다.
1. API와 직접 통신하는 것
2. Datadog 수집기를 클러스터 내부의 컨테이너로 실행하는 것


##### 2. 자체 호스팅
Grafana
- [Grafana Official](https://grafana.com/)
- [대시보드 만드는 방법 by Grafana Labs](http://docs.grafana.org/features/panels/graph/)
