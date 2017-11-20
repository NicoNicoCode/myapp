MyApp
===============
1.myapp介绍
-----------------
	    myapp是一个提供登陆注册服务的网站，同时，它允许用户使用myapp-cli程序来回显用户输入和定时报时。
        
2.相关页面功能介绍
-----------------
       192.168.34.52:8080              显示Beego界面，说明程序运行起来了
       192.168.34.52:8080/user/signup   显示注册界面，注册成功后顶部显示“注册成功”
       192.168.34.52:8080/user/login    显示登陆界面，登陆成功后跳转到
       192.168.34.52:8080/user/profile  显示个人信息
       192.168.34.52:8081               没有界面，提供回显用户输入服务
       192.168.34.52:8082               没有界面，提供自动报时服务
     
3.接口说明
-----------------
|  接口名称  |  接口地址  |  请求方式  |  输入  |  说明  |
|  --------  |  --------  |  :--------:|  :----:  |  ----  |
|  主页预览  |192.168.34.52:8080|GET|无| 提示myapp处于运行状态 |
|  注册界面  |192.168.34.52:8080/user/signup|GET|无|显示注册界面|
|  登陆界面  |192.168.34.52:8080/user/login|GET|无|显示登陆界面|
|注册界面提交信息|192.168.34.52:8080/user/signup|POST|用户名(字符串)<br>密码(字符串)<br>个人简介(字符串，可选)|根据用户名是否已经存在会返回不同的信息|
|登陆界面提交信息|192.168.34.52:8080/user/profile|POST|用户名(字符串)<br>密码(字符串)|根据用户名和密码与数据库内的信息进行匹配|
|  回显服务  |192.168.34.52:8081|POST|无|用户输入任意字符串，服务器会进行回显|
|  报时服务  |192.168.34.52:8082|POST|无|每隔一分钟显示当前时间|

4.部署方法
-----------------------------
- 拉取CentOS镜像
- 编译myapp项目，生成可执行文件--myapp在项目目录下
- 编写Dockerfile
   * 复制myapp项目到CentOS容器中
   * 添加OEM和VER环境变量
- 执行Dockerfile
- 上传镜像到Ekos网站
- 创建应用
	* 使用上传的镜像
	* 映射所需端口(8080、8081、8082)
	* 启动命令为“./myapp -logtostderr=true”
- 创建负载均衡，将映射出来的端口进行关联

5.遇到的问题和解决方法
------------------------------
	1.运行日志插件出现未知错误
	解决方案:修改logging.yaml,使得插件所需内存为4G，但这样会造成日志插件不稳定，因此创建一个8G的虚拟机作为日志插件结点会更好。
    
6.通过kubelet命令将部署的应用副本数扩展为4
------------------------------------------
    先查看资源分配情况
	[root@node1 ~]# kubectl get deployment
    NAME                   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
    default-http-backend   1         1         1            1           5d
    lb-hellodata           1         1         1            1           5d
    lb-helloworld          1         1         1            1           5d
    myapp                  1         1         1            1           5d
    
    使用scale指令
    [root@node1 ~]# kubectl scale --replicas=4 deployment/myapp
	deployment "myapp" scaled
    
    再次查看资源分配情况
    [root@node1 ~]# kubectl get deployment
    NAME                   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
    default-http-backend   1         1         1            1           5d
    lb-hellodata           1         1         1            1           5d
    lb-helloworld          1         1         1            1           5d
    myapp                  4         4         4            4           5d

7.查看pod信息
-----------------------
	[root@node1 ~]# kubectl get po -o wide
    NAME                                    READY     STATUS    RESTARTS   AGE       IP              NODE
    data-0                                  1/1       Running   2          5d        10.233.75.26    node2
    default-http-backend-3138300093-vq8q0   1/1       Running   2          5d        10.233.71.37    node3
    lb-hellodata-1110184062-dh7fv           1/1       Running   2          5d        192.168.34.53   node3
    lb-helloworld-3554808502-6zq4j          1/1       Running   2          5d        192.168.34.52   node2
    myapp-2365328813-3z1fl                  1/1       Running   0          14m       10.233.75.39    node2
    myapp-2365328813-7h838                  1/1       Running   0          1m        10.233.71.51    node3
    myapp-2365328813-qthq0                  1/1       Running   0          1m        10.233.75.40    node2
    myapp-2365328813-w575q                  1/1       Running   0          1m        10.233.71.52    node3

8.查看pod输出的日志
--------------------
    #命令格式:kubectl logs -f <pod name>
    [root@node1 ~]# kubectl logs -f myapp-2365328813-3z1fl
    2017/11/20 05:46:29 [I] [asm_amd64.s:2197] http server Running on http://:8080
    2017/11/20 06:01:34 [D] [server.go:2568] |  192.168.34.52| 200 |  40.296052ms|   match| GET      /user/profile   r:/user/profile
    I1120 06:01:34.672150       1 default.go:183] client's host:myapp-2365328813-3z1fl
    I1120 06:01:34.672189       1 default.go:184] client's ip:10.233.75.39
    2017/11/20 06:06:38 [D] [server.go:2568] |  192.168.34.52| 200 |   5.508274ms|   match| GET      /user/profile   r:/user/profile
    I1120 06:06:38.711007       1 default.go:183] client's host:myapp-2365328813-3z1fl
    I1120 06:06:38.711038       1 default.go:184] client's ip:10.233.75.39
    I1120 06:06:39.667596       1 default.go:183] client's host:myapp-2365328813-3z1fl
    I1120 06:06:39.667621       1 default.go:184] client's ip:10.233.75.39
    2017/11/20 06:06:39 [D] [server.go:2568] |  192.168.34.52| 200 |  38.372187ms|   match| GET      /user/profile   r:/user/profile

[kubectl常用命令(指引6)](https://github.com/NicoNicoCode/myapp/kubectl_command.md)