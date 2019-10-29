![kaau](web/images/logo.png)

# Project KAAU.
Kubernetes Authorization Authentication Utility

NAQ(Nobody Asked Questions)

Q1) What is KAAU?

Ans: KAAU is an abbreviation of Kubernetes Authorization Authentication Utility. It is a small hobby project I have been working on the past couple of months.

Q2) What this utility does?

Ans: KAAU is a web-based application it does CRUD(Create, Read, Update, and Delete) on Kubernetes Authentication and Authorization part which are.

Authentication
- service Account

Authorization
- Roles
- Cluster Roles
- RoleBinding
- ClusterRoleBinding

Q3) Hold on you missed User account in which are part of Authentication.

Ans: Good catch. User accounts are not managed by Kubernetes and I am using an API call to Kubernetes to manage the above.

```
Normal users are assumed to be managed by an outside, independent service. An admin distributing private keys, a user store like Keystone or Google Accounts, even a file with a list of usernames and passwords. In this regard, Kubernetes does not have objects which represent normal user accounts. Normal users cannot be added to a cluster through an API call.
```
[https://kubernetes.io/docs/reference/access-authn-authz/authentication/](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)


Q4) Ok. But I have Kubectl to do all the above tasks and more then why so much effort if no one going to use it.

Ans: So cruel, But it is ok if this goes in git graveyard. I loved kubernetes and was learning the Go language so I thought I could create something in kubernetes.

I use kubernetes every day all of using kubectl, But when your users and roles grow a lot it is painfull to find which account entitles which role via which role binding. So a simple UI can be helpful. It is ok if no one uses it will be in opensource so another brillance developer can see how easy to develop around kubernetes and make something awesome out of it. 

Q5) You are learning Go Language. Is that why code is so awful?

Ans: Yes. Not only Golan but HTML and CSS also. I know how to write a code still, have to learn how not to write code. 

Q6) That's great. how I can test it.

Ans: You can check in the install section. you can compile the binary or download them from the release section. Best way to install in minikube. I have provided minikube.yaml in install folder/


Q7) Imagine if I like a project. Then how I can contribute.

Ans: I will be working on fixing code. You can provide me star on GitHub for encouragement. It is an open-source project you can contribute to your capacity.

Q8) What is pending? 

Ans: As of now, all modules are working. The pending part is to generalize code, Add error handling and add unit test cases which I have to figure out how to do.

Q9) Last thing. What a great logo.

Ans: Thanks. It is "Crow the detective" made by my wife. She is not much on social media. I will convey your message. 



Login Screen
![Login page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/login_page.png)

Home Page
![Home Page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/Home_Page.png)


Service Account
![Home Page](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/Service_Account.png)


Role Page
![Role Page](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/Roles.png)

Create Role Page
![Create Role Page](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/Create_role.png)

Creating Role Page
![Create Role Page](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/creating_role.png)

Role Created notification
![Role Created notification](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/role_created.png)

Delete role Confirmation.
![Delete role Confirmation](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/Delete_role_confirmation.png) 

Delete role
![delete Role](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/role_deteled.png) 
 

 
 
Role Binding
![Role Binding](https://raw.githubusercontent.com/yogeshkk/KAAU/master/Doc/screens/Role_Binding.png)

Create Role Binding
![Role Binding](https://github.com/yogeshkk/KAAU/blob/master/Doc/screens/Create_Role_binding.png)



Read bit about RBAC
https://medium.com/@yogeshkunjir/kubernetes-has-your-r-back-5b4c983be0



Install KAAU
https://github.com/yogeshkk/KAAU/blob/master/Doc/install.md



Status of Project.
All modules are functional.


NOTE:- You can browse code and if like idea provides star for encouragement or provide feedback to me one below social networks. 


Twitter https://twitter.com/yogeshkunjir
LinkedIn https://www.linkedin.com/in/yogeshkunjir/


