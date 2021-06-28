# NSQ

## extract nsq and add env
```
tar -xvf nsq-1.0.0-compat.linux.tar.gz
```

## run command
```
nsqlookupd
nsqd --lookupd-tcp-address=127.0.0.1:4160
nsqadmin --lookupd-http-address=127.0.0.1:4161
```

## visit
```
http://127.0.0.1:4171
```

## install client
```
go get -v -u github.com/nsqio/go-nsq
```