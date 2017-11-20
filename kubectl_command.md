Kubectl 命令使用
===========

1.create
-----------
+	    解释:create命令用于根据文件或输入创建集群resource。
+	    实例:kubectl create -f logging.yaml	    
+	    输出结果:
	    deployment "ekos-logging-server" created
 	    service "logging" created
 	    configmap "filebeat" created
 	    daemonset "ekos-filebeat-server" created
 	    configmap "elasticsearch" created
 	    deployment "ekos-elasticsearch-server" created
 	    service "elasticsearch" created

      
2.scale
---------  
+	    解释:scale用于程序在负载加重或缩小时副本进行扩容或缩小。
+	    实例:kubectl scale --replicas=4 deployment/myapp
+	    输出结果:
	    NAME                   DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
		default-http-backend   1         1         1            1           1d
		lb-hellodata           1         1         1            1           1d
		lb-helloworld          1         1         1            1           1d
 	    myapp                  4         4         4            4           1d
3.get
----------- 
+	    解释:get命令用于获取集群的一个或一些resource信息。
+	    实例:kuebectl get pod
+	    输出结果:
 	    NAME                                    READY     STATUS    RESTARTS   AGE
 	    data-0                                  1/1       Running   2          1d
 	    default-http-backend-3138300093-vq8q0   1/1       Running   2          1d
 	    lb-hellodata-1110184062-dh7fv           1/1       Running   2          1d
 	    lb-helloworld-3554808502-6zq4j          1/1       Running   2          1d
 	    myapp-2207059372-mdd1v                  1/1       Running   0          1d

4.logs
----------
+	   解释:logs命令用于显示pod运行中，容器内程序输出到标准输出的内容。
+	    实例:kubectl logs data-0
+	    输出结果:
	    2017-11-15 07:03:06 0 [Note] mysqld (mysqld 5.6.35) starting as process 1 ...
        2017-11-15 07:03:06 1 [Note] Plugin 'FEDERATED' is disabled.
        2017-11-15 07:03:06 1 [Note] InnoDB: Using atomics to ref count buffer pool pages
        2017-11-15 07:03:06 1 [Note] InnoDB: The InnoDB memory heap is disabled
        2017-11-15 07:03:06 1 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
        2017-11-15 07:03:06 1 [Note] InnoDB: Memory barrier is not used
        2017-11-15 07:03:06 1 [Note] InnoDB: Compressed tables use zlib 1.2.8
        2017-11-15 07:03:06 1 [Note] InnoDB: Using Linux native AIO
        2017-11-15 07:03:06 1 [Note] InnoDB: Using CPU crc32 instructions
        2017-11-15 07:03:06 1 [Note] InnoDB: Initializing buffer pool, size = 128.0M
        2017-11-15 07:03:06 1 [Note] InnoDB: Completed initialization of buffer pool
        2017-11-15 07:03:06 1 [Note] InnoDB: Highest supported file format is Barracuda.
        2017-11-15 07:03:06 1 [Note] InnoDB: 128 rollback segment(s) are active.
        2017-11-15 07:03:06 1 [Note] InnoDB: Waiting for purge to start
        2017-11-15 07:03:06 1 [Note] InnoDB: 5.6.35 started; log sequence number 1626017
        2017-11-15 07:03:06 1 [Note] Server hostname (bind-address): '*'; port: 3306
        2017-11-15 07:03:06 1 [Note] IPv6 is available.
        2017-11-15 07:03:06 1 [Note]   - '::' resolves to '::';
        2017-11-15 07:03:06 1 [Note] Server socket created on IP: '::'.
        2017-11-15 07:03:06 1 [Warning] 'proxies_priv' entry '@ root@data-0' ignored in --skip-name-resolve mode.
        2017-11-15 07:03:06 1 [Note] Event Scheduler: Loaded 0 events
        2017-11-15 07:03:06 1 [Note] mysqld: ready for connections.
        Version: '5.6.35'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)

5.describe
-------
+	    解释:describe类似于get，同样用于获取resource的相关信息。不同的是，get获得的是更详细的resource个性的详细信息，describe获得的是resource集群相关的信息。
+	    实例:kubectl describe data-0
+	    输出结果:
	    Name:		data-0
        Namespace:	default
        Node:		node2/192.168.34.52
        Start Time:	Wed, 15 Nov 2017 11:50:07 +0800
        Labels:		ekos-application=myapplications
        ekos-service=data
        log-ignore=false
        type=service
        version=1
        Annotations:	kubernetes.io/created-by={"kind":"SerializedReference","apiVersion":"v1","reference":{"kind":"StatefulSet","namespace":"default","name":"data","uid":"0da7c608-c9b8-11e7-9404-001a4a180005","apiVersion"...
		pod.beta.kubernetes.io/hostname=data-0
		pod.beta.kubernetes.io/subdomain=data
        Status:		Running
        IP:		10.233.75.26
        Controllers:	StatefulSet/data
        Containers:    
  	     mysql:
        Container ID:	   
        docker://403beda476440e2dc61530c8053dc108797de5b09a862d03ac0088697f52776b
      Image:		registry.ekos.local/default/mysql:latest
      Image ID:    docker-pullable://registry.ekos.local/default/mysql@sha256:778c1a0dc7548777de660def3aa129cf24a7488f16a7de2602592d9982910a91
        Port:		
        State:		Running
        Started:		Wed, 15 Nov 2017 15:02:53 +0800
        Last State:		Terminated
        Reason:		Completed
      Exit Code:	0
      Started:		Wed, 15 Nov 2017 14:15:17 +0800
      Finished:		Wed, 15 Nov 2017 14:37:42 +0800
      Ready:		True
      Restart Count:	2
      Limits:
        cpu:	1
        memory:	512M
      Requests:
        cpu:	500m
        memory:	256M
      Environment:
        MYSQL_ROOT_PASSWORD:	123456
      Mounts:
        /var/lib/mysql/ from mysqlvarlibmysql (rw)
        /var/run/secrets/kubernetes.io/serviceaccount from default-token-kws32 (ro)
      Conditions:
        Type		Status
        Initialized 	True 
        Ready 	True 
        PodScheduled 	True 
        Volumes:
        mysqlvarlibmysql:
        Type:	PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
        ClaimName:	mysqlvarlibmysql-data-0
        ReadOnly:	false
        default-token-kws32:
        Type:	Secret (a volume populated by a Secret)
        SecretName:	default-token-kws32
        Optional:	false
        QoS Class:	Burstable
        Node-Selectors:	<none>
        Tolerations:	<none>
        Events:		<none>
