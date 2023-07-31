# eeta

**eeta** is a blockchain built using Cosmos SDK and Tendermint.


GovernanceToken: eeta

StableCoin: krw

## MODULES

### `deposit`

**스테이블코인 발행**

#### Roles

- 법정화폐 입금 확인, on-chain stable coin 발행
- on-chain stable coin 소각 > 법정화폐 출금

### `billboard`

**빌보드 정보**

#### Roles

- 빌보드 소유자가 본인 소유의 빌보드를 등록
- 빌보드 시간송출권의 입찰 STO는 매일 저녁 11시59분 펀딩금액 100% 달성 및 매입 단가 높은 STO

#### Skips

1. 시간송춸권 입찰, 매입 최소가 등 비지니스적인 고려는 데모에서 하지 않음.

### `sto`

**빌보드 시간송출권 펀딩**

#### Roles

- 빌보드, 송출시간, 빌보드 매입금액, 펀드 규모(지분 쉐어)를 등록하여 STO
- 사용자가 투자하면 사용자의 코인은 `st fund`에 예치되고 `st asset` 발행
- 규모 미달성, 동시간대 다른 펀딩 성공 등의 사유로 펀딩 실패시 자산 반환
- 성공한 펀드는 기업이 광고 입찰 참여 가능
- 송출시간 전까지 입찰에 참여한 기업중 STO owner가 선택 가능
- `audit` module만 `sto`의 펀드 자금 출금 가능

(sto 성공, 기업 광고까지 등록되면 미디어 광고 송출)

#### Skips

### `audit`

__sto의 펀드 지급 및 패널티 등 판정__

#### Roles

- `sto`의 광고기한 만료시 빌보드 및 owner에게 이익 배당

- `sto`의 광고기간중 `sensuality` 신고 가능, 신고수가 `audit` module param의 설정수 넘어가면 광고 중도 중단

- `sensuality` 신고시 토큰 리워드

### Skips
1. 빌보드의 하드웨어, 인터넷 연결 유실 등의 문제로 계약된 미디어가 송출되지 않을 경우 패널티는 데모에서 고려하지 않음.
2. 광고 심의 부적합에 대한 명핵한 기준치 등은 데모에서 고려하지 않음.
3. 신고시 리워드를 어디서 주는지, 어떻게 사용되고 선순환 구조가 어떻게 그려지는지 등에 대한 메커니즘은 데모에서 고려하지 안음.

---
brainstoming
- 미디어 송출
- 미디어 송출 검증 오라클
- 미디어 심의 판정 리워드 & 오라클
