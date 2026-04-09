#!/usr/bin/env bash
set -euxo pipefail
PS4='$ '

export PATH=/home/savva/.local/bin:$PATH
LAB4=/home/savva/Desktop/devops/infra-course-labs/lab4

kaniko_build() {
  local context="$1"
  local dockerfile="$2"
  local destination="$3"
  local tarfile="$4"
  shift 4

  rm -f "$tarfile"
  udocker run --volume="$LAB4:/lab4" --entrypoint="" martizih/kaniko:latest \
    /kaniko/executor \
    --context="$context" \
    --dockerfile="$dockerfile" \
    --destination="$destination" \
    --tar-path="$tarfile" \
    --no-push \
    --force \
    "$@"
}

cd "$LAB4/work"

kaniko_build /lab4/work/hello /lab4/work/hello/Dockerfile local/hello:latest /lab4/artifacts/logs/hello.tar
udocker rmi local/hello:latest || true
udocker import "$LAB4/artifacts/logs/hello.tar" local/hello:latest
udocker run local/hello:latest /bin/sh -lc 'echo Hello from hello image'

kaniko_build /lab4/work/main /lab4/work/main/Dockerfile local/main:latest /lab4/artifacts/logs/main.tar --build-arg FILE=/example.html
udocker rmi local/main:latest || true
udocker import "$LAB4/artifacts/logs/main.tar" local/main:latest
udocker rm main-http || true
udocker create --name=main-http local/main:latest
pkill -f 'uvicorn main:app' || true
nohup udocker run main-http /bin/sh -lc 'cd /app && FILE=/example.html uvicorn main:app --host 0.0.0.0 --port 8888' > "$LAB4/artifacts/logs/main_http.out" 2>&1 &
main_pid=$!
trap 'kill "$main_pid" 2>/dev/null || true; kill "$registry_pid" 2>/dev/null || true' EXIT
sleep 10
wget -qO- http://127.0.0.1:8888 | tee "$LAB4/artifacts/logs/main_http_response.out"
python3 -c 'import json, tarfile; tf=tarfile.open("'"$LAB4/artifacts/logs/main.tar"'"); manifest=json.load(tf.extractfile("manifest.json")); cfg=json.load(tf.extractfile(manifest[0]["Config"])); print(json.dumps({"Entrypoint": cfg["config"].get("Entrypoint"), "Cmd": cfg["config"].get("Cmd"), "ExposedPorts": cfg["config"].get("ExposedPorts"), "Labels": cfg["config"].get("Labels")}, indent=2))'

kaniko_build /lab4/work/multi /lab4/work/multi/Dockerfile local/multi:latest /lab4/artifacts/logs/multi.tar
udocker rmi local/multi:latest || true
udocker import "$LAB4/artifacts/logs/multi.tar" local/multi:latest
udocker run local/multi:latest /hello

udocker pull registry:2
udocker rm registry-http || true
udocker create --name=registry-http registry:2
pkill -f '/entrypoint.sh /etc/docker/registry/config.yml' || true
nohup udocker run registry-http > "$LAB4/artifacts/logs/registry.out" 2>&1 &
registry_pid=$!
sleep 10
wget -qO- http://127.0.0.1:5000/v2/ | tee "$LAB4/artifacts/logs/registry_v2.out"

udocker run --volume="$LAB4:/lab4" --entrypoint="" martizih/kaniko:latest \
  /kaniko/executor \
  --context=/lab4/work/hello \
  --dockerfile=/lab4/work/hello/Dockerfile \
  --destination=127.0.0.1:5000/hello:0.1 \
  --insecure \
  --insecure-registry=127.0.0.1:5000 \
  --skip-tls-verify \
  --skip-push-permission-check \
  --force

wget -qO- http://127.0.0.1:5000/v2/_catalog | tee "$LAB4/artifacts/logs/registry_catalog.out"
wget -qO- http://127.0.0.1:5000/v2/hello/tags/list | tee "$LAB4/artifacts/logs/registry_tags.out"
