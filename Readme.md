* upload docker-compose.yml
```
scp docker-compose.yml xx@xx:/root/
```
* update local beslow/rocketmq_consumer and start app
```
docker pull beslow/rocketmq_consumer:latest &&
docker-compose -f docker-compose.yml down &&
docker-compose -f docker-compose.yml up -d
```