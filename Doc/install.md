
Install/Try project 

Username and Password will be at https://github.com/yogeshkk/KAAU/blob/master/src/utility/validation.go

Best way to run inside minikube.
```
kubectl create -f https://github.com/yogeshkk/KAAU/blob/master/install/minikube.yaml
```
get the pod name
```
kubectl.exe get pod -l app=KAAU -o jsonpath="{.items[0].metadata.name}"
```
Port forward to access UI.
```
kubectl.exe  port-forward "pod-name" 3333 --address=0.0.0.0
```

OR 

you can download windows and linux binary from release page. 
Run minukube and then run exe it will take kube config from home folder. 

OR

you can customize build as below
```
go get https://github.com/yogeshkk/KAAU
cd $GOPATH/src/github/yogeshkk/KAAU/src
glide ensure
cd ..
go build src/main.go
./main.exe
Brower at localhost:3333
```