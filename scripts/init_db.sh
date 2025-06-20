#!/bin/bash

set -e

# 사용자 지정 변수
MYSQL_ROOT_PASSWORD="your_root_password"     # 🔐 root 계정 비밀번호
APP_DB_NAME="ai_service"                     # 사용할 DB 이름
APP_DB_USER="ai_user"                        # 신규 생성할 사용자
APP_DB_PASSWORD="your_app_password"          # 신규 사용자 비밀번호

# 1. MySQL 설치
echo "▶ MySQL 서버 설치 중..."
sudo apt update
sudo DEBIAN_FRONTEND=noninteractive apt install -y mysql-server

# 2~3. root 비밀번호 설정 + DB + 사용자 + 테이블 생성
echo "▶ root 비밀번호 설정, DB 및 사용자 생성 중..."

sudo mysql <<EOF
-- root 비밀번호 설정 (기존 socket 인증 → 비밀번호 인증으로 전환)
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '${MYSQL_ROOT_PASSWORD}';
FLUSH PRIVILEGES;

-- DB 생성
CREATE DATABASE IF NOT EXISTS \`${APP_DB_NAME}\`;

-- 사용자 생성 (존재하지 않을 경우)
CREATE USER IF NOT EXISTS '${APP_DB_USER}'@'%' IDENTIFIED BY '${APP_DB_PASSWORD}';

-- 권한 부여
GRANT ALL PRIVILEGES ON \`${APP_DB_NAME}\`.* TO '${APP_DB_USER}'@'%';
FLUSH PRIVILEGES;

-- 테이블 생성
USE \`${APP_DB_NAME}\`;

CREATE TABLE IF NOT EXISTS usage_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(64) NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    photos INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
EOF

echo "✅ MySQL 설정이 완료되었습니다: 비밀번호 설정, DB 및 테이블 생성!"
