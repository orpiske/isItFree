

docker run --rm --name influxdb -v $HOME/influxdb-data:/root/.influxdbv2 -p 8086:8086 quay.io/influxdb/influxdb:v2.0.2 --reporting-disabled