# Project KAAU.

NAQ(Nobody Asked Questions)

Q1) What is KAAU?
KAAU is an abbreviation of Kubernetes Authorization Authentication Utility. It is a small hobby project I have been working on the past couple of month

Q2) What this utility does?
KAAU is a web-based application it does CRUD(Create, Read, Update, and Delete) on Authentication and Authorization part which are.
Authentication,
- service Account
Authorization
- Roles
- Cluster Roles
- RoleBinding
- ClusterRoleBinding

Q3) Hold on you missed User account in which are part of Authentication.
Good catch. I am using an API call to Kubernetes to manage the above and User accounts are not allowed to create via this. read in detail.

Q4) Ok. But I have Kubectl to do all the above tasks and more then why so much effort if no one going to use it.
So cruel but it is understood. I was loved and learning the Go language so I thought I could do something also I use kubernetes every day. When your users and roles grow it is painfull to find
which account entitles which role via which role binding. So a simple UI can be helpful. It is ok if no one uses it will be in opensource so another brillance developer can see how easy to develop around kubernetes.

Q5) You are learning Go Lan. Is that why code is so awful?
Yes. Not only Golan but HTML and CSS also. I know how to write a code still, have to learn how not to write code. 

Q6) That's great. how I can test it.
You can check in the install section. you can compile the binary or download them from the release section. I will be releasing the docker image soon.


Q7) Imagine if I like a project. Then how I can contribute.
I will be working on fixing code. You can provide me star on GitHub for encouragement. It is an open-source project you can contribute to your capacity.

Q8) What is pending? 
As of now, all modules are working. The pending part is to generalize code, Add error handling and add unit test cases which I have to figure out how to.

Q9) Last thing. What a great logo.
Thanks. It is "Crow the detective" made by my wife. She is not much on social media. I will convey your message. 


Read bit about RBAC
https://medium.com/@yogeshkunjir/kubernetes-has-your-r-back-5b4c983be0


Status of Project.
All module are functional.

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
