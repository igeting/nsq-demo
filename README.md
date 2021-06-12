# NSQ

## 本地解析HOSTS
```
tail -1 /etc/hosts
192.168.43.47 nsq-01
```

## 开三个终端,分别按顺序启动
```
./nsqlookupd
./nsqd --lookupd-tcp-address=192.168.43.47:4160
./nsqadmin --lookupd-http-address=192.168.43.47:4161
```

## 访问
```
http://192.168.43.47:4171
```