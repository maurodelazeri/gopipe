# gopipe

Data Pipeline in Go with Kafka

![alt text](https://github.com/maurodelazeri/gopipe/blob/master/687474703a2f2f62656e73746f70666f72642e636f6d2f75706c6f6164732f696d672f536c69646534312e706e67.png)


## Development

**IMPORTANT**: Producer requires actual IP of brokers (we can't use `localhost` for kafka producers)

```
## kafka
KAFKA_ADVERTISED_HOST_NAME=192.168.0.35 docker-compose up   # should set advertised.host

## server-gateway
$ cd server-gateway
$ make install                             # only once
$ BROKERS=192.168.0.35:9092 make run-cont  # should set broker list

## cli
# cd cli
# go run main.go
```

## TODO

- server-connect
- elasticsearch
- web server
- wep application

## References

Based on http://highscalability.com/blog/2015/5/4/elements-of-scale-composing-and-scaling-data-platforms.html
