#!/usr/bin/env bash
set -euxo pipefail
PS4='$ '

LAB6=/home/savva/Desktop/devops/infra-course-labs/lab6
PROJECT=lab6app

cleanup() {
  cd "$LAB6/work"
  sudo -n docker compose --project-name "$PROJECT" down -v --remove-orphans >/dev/null 2>&1 || true
}

trap cleanup EXIT

cd "$LAB6/work"
cleanup

sudo -n docker compose --project-name "$PROJECT" up -d --build
sleep 15

sudo -n docker compose --project-name "$PROJECT" ls
sudo -n docker compose --project-name "$PROJECT" ps
sudo -n docker compose --project-name "$PROJECT" top || true

wget -qO- http://127.0.0.1:18888 | tee "$LAB6/artifacts/logs/front_index.out"
wget -qO- http://127.0.0.1:18888/back | tee "$LAB6/artifacts/logs/back_initial.json"

sudo -n docker exec "${PROJECT}-db-1" psql -U app -d app -c "insert into users (name, email) values ('alex', 'alex@mail.ru');"
wget -qO- http://127.0.0.1:18888/back | tee "$LAB6/artifacts/logs/back_after_insert.json"

sudo -n docker compose --project-name "$PROJECT" rm -sf
sudo -n docker compose --project-name "$PROJECT" up -d
sleep 10
wget -qO- http://127.0.0.1:18888/back | tee "$LAB6/artifacts/logs/back_after_recreate.json"

sudo -n docker exec "${PROJECT}-front-1" sh -lc 'wget -T 2 -qO- http://db:5432' > "$LAB6/artifacts/logs/front_to_db.out" 2> "$LAB6/artifacts/logs/front_to_db.err" || true
