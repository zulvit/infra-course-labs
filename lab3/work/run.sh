#!/usr/bin/env bash
set -euxo pipefail
PS4='$ '

export PATH=/home/savva/.local/bin:$PATH

cd /home/savva/Desktop/devops/infra-course-labs/lab3/work

udocker pull python:3.11.5-alpine
udocker images
udocker run python:3.11.5-alpine python --version

udocker rm pyver || true
udocker create --name=pyver python:3.11.5-alpine
udocker ps -m -s

udocker rm test-http || true
udocker create --name=test-http python:3.11.5-alpine

rm -f ../artifacts/logs/step03_http_server.out
pkill -f '/vagrant/test.py' || true
nohup udocker run --volume="$PWD:/vagrant" test-http sh -lc 'FILE=/vagrant/sample.txt PORT=8888 python /vagrant/test.py' > ../artifacts/logs/step03_http_server.out 2>&1 &
server_pid=$!
trap 'kill "$server_pid" 2>/dev/null || true' EXIT

sleep 5
wget -qO- http://127.0.0.1:8888 | tee ../artifacts/logs/step03_http_response.out

udocker inspect pyver | sed -n '1,80p'
udocker manifest inspect python:3.11.5-alpine | sed -n '1,80p'
udocker save -o ../artifacts/logs/python-3.11.5-alpine.tar python:3.11.5-alpine
ls -lh ../artifacts/logs/python-3.11.5-alpine.tar

kill "$server_pid" 2>/dev/null || true
wait "$server_pid" 2>/dev/null || true
