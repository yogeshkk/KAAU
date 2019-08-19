# Project KAAU.

Kubernetes Authentication Authorisation Utility

This is a Hobby Project yet not completed. This project aims to provide a management utility to Authorisation and Authentication for Kubernetes from a simple UI

Read bit about RBAC
https://medium.com/@yogeshkunjir/kubernetes-has-your-r-back-5b4c983be0


Status of Project.
Currently, it can Authentication via web. you can see password hard coded in project ;-)

It can connect to Kubernetes via kubernetes client.go library

It can get below from Kubernetes show in UI.
1) Service Account
2) Roles
3) Cluster Roles
4) Role Binding
5) Cluster Role Binding

NOTE:- I am not(yet) a GO or HTML/CSS Developer the code is written by me as an amateur code so it will be buggy and not yet ready to use (though you can use in minikube like me). You can browse code and if like idea provides star for encouragement or provide feedback to me one below social networks. 

To build
```
go get https://github.com/yogeshkk/KAAU
cd $GOPATH/src/github/yogeshkk/KAAU/src
glide ensure
cd ..
go build src/main.go
./main.exe
Brower at localhost:3333
username and password will be at https://github.com/yogeshkk/KAAU/blob/master/src/utility/validation.go
```

Twitter https://twitter.com/yogeshkunjir
LinkedIn https://www.linkedin.com/in/yogeshkunjir/

Login Screen

![Login page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/login_page.png)

Home Page
![Home Page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/Home_Page.png)


Service Account
![Home Page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/Service_Account.png)


Role Page
![Home Page](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/Roles.png)


Role Binding
![Home Page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/Role_Binding.png)
