ulimit -n 1048576
mkdir -pv /opt/logs/push/

export REDIS_SERVER="10.11.157.178:15261"
export MONGO_URL="mongodb://192.168.230.52:27017,192.168.230.53:27017,192.168.230.54:27017?connect=replicaSet"
nohup /opt/webapps/push/bin/broker --tcp-address="123.125.116.57:80" --http-address="10.11.157.19:8601" --broadcast-address="10.11.157.19:80"  &> /opt/logs/push/broker.log &
