# Chatting

고언어를 활용한 REST API 개발하기.

# 기술 스택

- Go-Echo
- MongoDB
- Swagger

# 폴더 구조

```bash
├── README.md                 - 리드미 파일
│
├── api/                      - api 핸들러 폴더
│   ├── error.go              - 에러 API 핸들러 정의
│   ├── posts.go              - 게시물 API 핸들러 정의
│   ├── user.go               - 유저 API 핸들러 정의
│   └── web.go                - 페이지 API 핸들러 정의
├── config/                   - 설정 폴더
│   ├── cors.go               - cors 정의
│   ├── env.go                - 환경 설정 정의
│   └── mongodb.go            - DB 연결 설정 정의
├── docs/                     - 서버에 필요한 문서들 관리 폴더
│   ├── air.toml              - air 구동 파일
│   ├── docs.go               - 스웨거
│   ├── swagger.json          - 스웨거 json 파일
│   └── swagger.yaml          - 스웨서 yaml 파일
├── exception/                - 예외 관리 폴더
│   └── exception.go          - 예외 처리 정의
├── middleware/               - 미들웨어 관리 폴더
│   └── authToken.go          - 인증 미들웨어
├── model/                    - 모델 관리 폴더
│   ├── jwt.go                - jwt 모델
│   ├── post.go               - 게시물 모델
│   └── user.go               - 유저 모델
├── repository/               - 저장소 관리 폴더
│   ├── post.go               - 게시물 디비 처리
│   └── user.go               - 유저 디비 처리
├── routes/                   - 라우터 관리 폴더
│   ├── post.go               - 게시물 라우터 정의
│   ├── swagger.go            - 스웨거 라우터 정의
│   └── user.go               - 유저 라우터 정의
├── tmp/                      - air 폴더
│   └── main                  - air 실행 파일
├── utils/                    - 유틸리티 관리 폴더
│   ├── httpContext.go        - 컨텍스트 생성 파일
│   ├── jwt.go                - jwt 관리 파일
│   └── validator.go          - validator 관리 파일


```


