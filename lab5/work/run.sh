#!/usr/bin/env bash
set -euxo pipefail
PS4='$ '

LAB5=/home/savva/Desktop/devops/infra-course-labs/lab5

cleanup() {
  sudo -n docker rm -f lab5-nginx-bind lab5-nginx-volume lab5-writer lab5-reader lab5-first lab5-second lab5-default-first lab5-default-second lab5-host-http >/dev/null 2>&1 || true
  sudo -n docker network rm lab5-br >/dev/null 2>&1 || true
  sudo -n docker volume rm lab5-html lab5-empty lab5-new >/dev/null 2>&1 || true
}

cleanup
trap cleanup EXIT

cd "$LAB5/work"

mkdir -p data
cp index.html data/index.html

sudo -n docker pull nginx:alpine
sudo -n docker pull alpine
sudo -n docker pull python:3.11-alpine

sudo -n docker run -d -v "$PWD/data:/usr/share/nginx/html" -p 18088:80 --name lab5-nginx-bind nginx:alpine
sleep 3
wget -qO- http://127.0.0.1:18088 | tee "$LAB5/artifacts/logs/bind_mount_response.out"

sudo -n docker volume create lab5-empty
sudo -n docker volume create lab5-new --label test=true
sudo -n docker volume ls
sudo -n docker volume ls -f label=test=true
sudo -n docker volume prune --filter label=test=true -f

sudo -n docker volume create lab5-html
sudo -n docker run -d -v lab5-html:/usr/share/nginx/html -p 18089:80 --name lab5-nginx-volume nginx:alpine
sleep 3
sudo -n sh -lc 'echo "<html><body><h1>lab5 named volume</h1></body></html>" > "$(docker volume inspect lab5-html -f "{{ .Mountpoint }}")/index.html"'
wget -qO- http://127.0.0.1:18089 | tee "$LAB5/artifacts/logs/named_volume_response.out"
sudo -n docker volume inspect lab5-html

sudo -n docker run -d -v lab5-html:/shared --name lab5-reader alpine sleep inf
sudo -n docker run --rm -v lab5-html:/shared --name lab5-writer alpine sh -lc 'echo shared-from-writer > /shared/shared.txt'
sudo -n docker exec lab5-reader cat /shared/shared.txt | tee "$LAB5/artifacts/logs/shared_volume.out"

sudo -n docker run -d --name lab5-default-first alpine sleep inf
sudo -n docker run -d --name lab5-default-second alpine sleep inf
sudo -n docker exec lab5-default-first ping -c 1 lab5-default-second || true

sudo -n docker network create lab5-br --subnet 10.0.0.0/24
sudo -n docker network inspect lab5-br
sudo -n docker run -d --network lab5-br --name lab5-first alpine sleep inf
sudo -n docker run -d --network lab5-br --name lab5-second alpine sleep inf
sudo -n docker exec lab5-first ping -c 1 lab5-second | tee "$LAB5/artifacts/logs/custom_bridge_ping.out"
sudo -n docker network connect bridge lab5-first
sudo -n docker network disconnect bridge lab5-first

sudo -n docker run -d --network host --name lab5-host-http python:3.11-alpine python -m http.server 18891
sleep 3
wget -qO- http://127.0.0.1:18891 | tee "$LAB5/artifacts/logs/host_network_response.out"
