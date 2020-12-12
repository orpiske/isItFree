
Run Influx DB
```
docker run --rm --name influxdb -v $HOME/influxdb-data:/root/.influxdbv2 -p 8086:8086 quay.io/influxdb/influxdb:v2.0.2 --reporting-disabled
```

Build and run manually
```
make build && IIF_URL=http://host:8086 IIF_GYM_SOURCE_URL=http://gym-url  IIF_POOL_SOURCE_URL=http://pool-url IIF_BUCKET=academia-test IIF_TOKEN=token ./is-it-free
```


Build and install
```
make build token="myotken" influxdb=http://influxdbhost:8086 gym=https://gym-url pool=https://pool-url
make install
```

 