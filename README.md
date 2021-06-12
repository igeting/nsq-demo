# NSQ

## 开三个终端，分别按顺序启动
```
nsqlookupd
nsqd --lookupd-tcp-address=192.168.2.103:4160
nsqadmin --lookupd-http-address=192.168.2.103:4161
```

## 访问
```
http://192.168.2.103:4171
```

## 安装客户端
```
go get -v -u github.com/nsqio/go-nsq
```