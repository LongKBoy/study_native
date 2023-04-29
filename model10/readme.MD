
部署httpserver

```sh
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

```
查看service
```sh
kubectl get svc 

```
curl service
```sh
curl http://10.108.129.22/hello?user=zhangsan

```

curl metrics
```sh
curl http://10.108.129.22/metrics
```

发送更多请求
```sh
for i in {1..1000};do curl http://10.108.129.22/hello?user=zhangsan;done
```

gafana图表
gafana.png