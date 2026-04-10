#!/usr/bin/env bash
set -euo pipefail
PS4='$ '

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
LAB4_DIR=$(cd "$SCRIPT_DIR/.." && pwd)

cleanup() {
  docker rm -f lab4-main lab4-registry >/dev/null 2>&1 || true
}

trap cleanup EXIT

cd "$LAB4_DIR"

docker build -t lab4/hello:0.1 "$LAB4_DIR/work/hello"
docker run --rm lab4/hello:0.1
docker image inspect lab4/hello:0.1

docker build --build-arg FILE=/example.html -t lab4/main:0.1 "$LAB4_DIR/work/main"
docker image inspect lab4/main:0.1 --format '{{json .Config.ExposedPorts}}'
docker image inspect lab4/main:0.1 --format '{{json .Config.Labels}}'
docker image inspect lab4/main:0.1 --format '{{json .Config.Entrypoint}} {{json .Config.Cmd}}'
docker run -d -p 8888:8888 --name lab4-main lab4/main:0.1
sleep 5
curl -s http://127.0.0.1:8888/

docker build -t lab4/multi:0.1 "$LAB4_DIR/work/multi"
docker run --rm lab4/multi:0.1
docker history lab4/multi:0.1

docker run -d -p 5000:5000 --name lab4-registry registry:2
sleep 5
curl -s http://127.0.0.1:5000/v2/
docker tag lab4/hello:0.1 127.0.0.1:5000/hello:0.1
docker push 127.0.0.1:5000/hello:0.1
curl -s http://127.0.0.1:5000/v2/_catalog
curl -s http://127.0.0.1:5000/v2/hello/tags/list
curl -s -H 'Accept: application/vnd.oci.image.index.v1+json' http://127.0.0.1:5000/v2/hello/manifests/0.1
