#!/bin/bash
set -ex

CURRENT_DIR=$(cd "$(dirname "$0")" && pwd)
docker build --platform=linux/amd64 -t json2yaml:v1 "$CURRENT_DIR"
sleep 1
docker run -it --rm -p 21111:21111 json2yaml:v1
