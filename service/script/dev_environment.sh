#!/bin/bash

registry="./script/docker/dev/registry/consul-compose.yaml"
database="./script/docker/dev/database/database-compose.yaml"
queue="./script/docker/dev/queue/rocktemq-compose.yaml"

docker_compose="docker-compose.exe"
# 判断Docker的环境
docker_compose_path=$( command -v docker-compose.exe )
if [ -z "$docker_compose_path" ]
then
   docker_compose_path=$( command -v docker-compose )
   docker_compose="docker-compose"
fi

echo "docker path: $docker_compose_path"

$docker_compose -f $registry up -d
$docker_compose -f $database up -d
$docker_compose -f $queue up -d
