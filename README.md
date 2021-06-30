# NSQ

## extract nsq and add env
```
tar -xvf nsq-1.0.0-compat.linux.tar.gz
```

## run command
```
nohup ./nsqlookupd > /dev/null 2>&1 &
nohup ./nsqd -lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=iopening.cn > /dev/null 2>&1 &
nohup ./nsqadmin -lookupd-http-address=127.0.0.1:4161 > /dev/null 2>&1 &
```

## visit
```
http://iopening.cn:4171
```

## install client
```
go get -v -u github.com/nsqio/go-nsq
```