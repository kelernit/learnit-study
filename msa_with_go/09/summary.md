# Event-driven architecture
## 동기 vs 비동기
* Message Processing: 항상 동기식으로 설계하자
	* 가독성, 테스트, 디버깅에서 훨씬 쉽다
* 비동기처리는 꼭 필요할 때만 적용해야 한다.
	* Monolithic 시스템에 비해 scaling 에 좋다.
	* decoupling, scale, batch 작업, time-based 작업에 유리하다.
	* 비동기 처리는 새로운 인프라 구축이 필요하기 때문에 모니터링/관리 부담을 가중시킬 수 있다.
* Microservice vs Monolithic
	* 작은 코드 단위로 배포하는 것은 배포와 테스트가 쉽다.
	* 인프라 구축(CI/CD 등)은 한번만 제대로 하면 되는 것이기 때문에 Microservice 설계가 의미가 있다.

## 동기 처리
<img src="https://github.com/learnhee/learnit-study/blob/master/msa_with_go/09/images/Synchronous.jpeg" width="400">

*  Request에 대한  response가 즉각적으로 필요한 경우 적절하다.

## 비동기 처리
<img src="https://github.com/learnhee/learnit-study/blob/master/msa_with_go/09/images/Asynchronous.jpeg" width="400">

* 모든 downstream 어플리케이션과의 communication은 `queue`나 `message broker`를 사용하여 처리한다.

* 메세지를 **AWS SQS/SNS**, **Google Cloud Pub/Sub**, **NATS.io** 등을 사용한다.
* 사용자는 메세지의 전달 여부를 확인하는 로직만을 구현하면 되고, retry 등과 관련된 로직은 message broker 단에 구현되어 있다.

## Pull/Queue message
<img src="https://github.com/learnhee/learnit-study/blob/master/msa_with_go/09/images/Pull-Queue.jpeg" width="400">

* 이미지 resizing같은 worker 프로세스가 있는 경우에 사용하면 좋다.
* 동작:
	1. API에 보내진 request는 background의 queue로 보낸다.
	2. Worker가 queue에서 메세지를 가져온다.
	3. (Worker가) 필요한 작업(예. image resizing)을 처리한다.
	4. 성공한 경우, queue에서 메세지를 제거하고, 실패한 경우 `dead letter queue(디버깅 또는 재처리 용도)`에 저장하거나, 기존 queue에 메세지를 남겨둔다.
	
* 실제 코드: 동작하는 것을 확인하려면 다음과 같은 처리가 필요하다.
	1. reader.go/writer.go에서 import path 수정
	2. makefile 수정
	3. docker-compose.yml에서 redis 포트 및 image 수정
	4. make 명령어로 컴파일 및 docker-compose up 실행
	5. postman으로 확인하기
