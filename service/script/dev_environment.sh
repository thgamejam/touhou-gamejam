#!/bin/bash

registry="./script/docker/dev/registry/consul-compose.yaml"
database="./script/docker/dev/database/database-compose.yaml"
queue="./script/docker/dev/queue/rocktemq-compose.yaml"

# 判断Docker的环境
docker="docker.exe"
docker_path=$( command -v docker.exe )
if [ -z "$docker_path" ]
then
   docker_path=$( command -v docker )
   docker="docker"
fi

docker_compose="docker-compose.exe"
docker_compose_path=$( command -v docker-compose.exe )
if [ -z "$docker_compose_path" ]
then
   docker_compose_path=$( command -v docker-compose )
   docker_compose="docker-compose"
fi

echo "docker path: $docker_path"
echo "docker-compose path: $docker_compose_path"

# 创建网络
network_name="th_bridge"
filter_name=$( $docker network ls | grep $network_name | awk '{ print $2 }' )
if [ "$filter_name" == "" ]; then
    $docker network create --subnet=172.84.0.0/24 --gateway=172.84.0.1 th_bridge
fi

$docker_compose -f $registry up -d
$docker_compose -f $database up -d
$docker_compose -f $queue up -d
