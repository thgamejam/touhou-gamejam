docker run -d --name consul -p 8500:8500 -p 8600:8600 -e CONSUL_BIND_INTERFACE=eth0 consul:1.11.2
docker run -d --name mariadb -p 3306:3306 --env MARIADB_ROOT_PASSWORD=123456 mariadb:10.6.5
docker run -d --name redis -p 6379:6379 redis:6.2
docker run -d --name minio -p 9000:9000 -p 9001:9001 minio/minio server /data --console-address ":9001"
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 --hostname my-rabbit rabbitmq:3.9.13-management