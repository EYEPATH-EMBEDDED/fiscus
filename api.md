

## ✅ 1. `POST /log`: 사용 로그 저장 API

### 📤 클라이언트 → 서버 (요청)

**URL:**

```
POST /logs
```

**Content-Type:**

```
application/json
```

**Body 예시 (JSON):**

```json
{
  "user_id": "andylim1022",
  "start_time": "2025-06-10 13:00:00",
  "end_time": "2025-06-10 13:30:00",
  "photos": 5
}
```

* `user_id`: 사용자 고유 ID
* `start_time`, `end_time`: `"YYYY-MM-DD HH:MM:SS"` 형식
* `photos`: 사진 수 (선택값, 없으면 기본값 0)

---

### 📥 서버 → 클라이언트 (응답)

**성공:**

```json
{
  "message": "로그 저장 완료"
}
```

**에러 예시 (형식 오류):**

```json
{
  "error": "Key: 'UsageLogInput.UserID' Error:Field validation for 'UserID' failed on the 'required' tag"
}
```

---

## ✅ 2. `GET /usage/:userId`: 특정 사용자의 월별 사용량 조회 API

### 📤 클라이언트 → 서버 (요청)

**URL 형식:**

```
GET /usage/<user_id>?year=YYYY&month=MM
```

**예시:**

```
GET /usage/andylim1022?year=2025&month=6
```

---

### 📥 서버 → 클라이언트 (응답)

**성공:**

```json
{
  "user_id": "andylim1022",
  "year": 2025,
  "month": 6,
  "used_minutes": 120,   // 총 사용 시간 (분)
  "photo_count": 12      // 총 촬영 수
}
```

**에러 예시 (파라미터 누락):**

```json
{
  "error": "year 파라미터가 필요하며 숫자여야 합니다"
}
```

---

## 💡 요약: 통신 정리

| API              | 메서드  | 파라미터             | 바디 형식                                           | 응답 형식                              |
|------------------| ---- | ---------------- | ----------------------------------------------- | ---------------------------------- |
| `/logs`          | POST | 없음               | JSON (user\_id, start\_time, end\_time, photos) | JSON 메시지                           |
| `/usage/:userId` | GET  | year, month (쿼리) | 없음                                              | JSON (used\_minutes, photo\_count) |
