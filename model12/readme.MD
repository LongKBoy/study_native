

创建namespace
```sh
kubectl create ns module12
开启istio注入
```sh
kubectl label ns module12 istio-injection=enabled

生成证书
```sh
openssl req -x509 -sha256 -nodes -days 3650 -newkey rsa:2048 -subj '/O=longqh Inc./CN=*.longqh.com' -keyout longqh.com.key -out longqh.com.crt
kubectl create -n istio-system secret tls longqh-credential --key=longqh.com.key --cert=longqh.com.crt

部署httpserver
```sh
kubectl apply -f deployment.yaml -n module12
kubectl apply -f service.yaml -n module12
kubectl apply -f VirtualService.yaml -n module12 
kubectl apply -f Gateway.yaml -n module12
    s
```

绑定域名到hosts,通过istio-ingressgateway https端口访问
https://httpsserver.longqh.com:30047/hello?user=zhangsan

查看tracing信息
tracing.png