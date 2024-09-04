```mermaid
erDiagram
  restaurants {
    varchar id PK
    varchar name "店舗名"
    varchar genre "ジャンル"
    varchar address "住所"
    varchar nearest_station "最寄り駅"
    varchar url  "URL"
    varchar posted_by "投稿者"
    datetime created_at
  }
```
