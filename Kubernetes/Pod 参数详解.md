
apiVersion:v1 # 版本号 如 v1
kind: Pod # 类型 Pod
metadata: # 元数据 
  name: string # Pod 的名称，命令规范需要符合 RFC 1035 spec
  namespace: string # Pod 的命名空间，默认值为 default
  lables: # 自定义标签列表 
  - name:string
  annotations: # 自定义注释列表
  - name: string
spec: # Pod 中容器的详细定义
  containers: # Pod 中的容器列表
  - name: string # 容器的名称，需要符合 RFC 1035 spec
    image: string # 容器的镜像名称
	imagePullPolicy: [Always| Never| IFNotPresent] # 	
镜像拉取策略，可选值包括：Always、Nerver、IfNotPresent,默认值为 Always
	command: [string] # 容器的启动命令列表，如果不指定，则使用镜像打包时使用的启动命令
	args: [String] # 容器的启动命令参数列表
	workingDir: string # 容器的工作目录
	volumeMounts: # 挂载到容器内部的存储卷配置
	- name: string # 引用 Pod 定义的共享存储卷的名称，需要使用 volumes[] 部分定义的共享存储卷名称
	  mountPath: string # 存储卷在容器内 Mount 的绝对路径，应少于 512 个字符
	  readOnly: boolean # 是否为只读模式，默认为读写模式
	ports: # 容器需要暴露的端口号列表
	- name: string # 端口的名称
	  containerPort: int # 容器需要监听的端口号
	  hostPort: int # 容器所在主机需要监听的端口号，默认与 containerPort 相同，设置 hostPort 时，同一台宿主机将无法启动该容器的第 2 份副本
	  protocol: string # 端口协议，支持 TCP 和 UDP，默认值为 TCP
	env: # 
	- name:string
	  value: string
	resources:
	  limits:
	    cpu: string
		memory: string
	  requests:
	    cpu: string
		memory: string
	livenessFrobe:
	  exec:
	    commadn: [string]
	  httpGet:
	    path: string
		port: number
		host: string
		scheme: string
		httpHeaders:
		- name: string
		  value: string
		tcpSocket:
		  port: number
		initialDelaySeconds: 0
		timeoutSeconds: 0
		periodSeconds: 0
		successThreshold: 0
		failureThreshold: 0
	  securityContext:
	    privileged: false
	restartFolicy: [Always| Never|| OnFaliure]
	nodeSelector: object
	imagePullSecrets:
	- name: string
	hostNetwork: false
	volumes:
	- name: string
	  enptyDir: {}
	  hostPath:
	    path: string
	  secret:
	    secretName: string
		items:
		- key: string
		  path: string
		configMap:
		  name: string
		  items:
		  - key: string
		    path: string


属性名称	取值范围	
是否必选

(1必选)

取值说明
apiVersion	string	1	版本号  例如：v1
kind	String	1	Pod
metadata	Object	1	元数据
metadata.name	String	1	Pod的名称，命令规范需要符合RFC 1035规范
metadata.namespace	String	1	Pod的命名空间，默认值为default
metadata.labels[]	List	 	自定义标签列表
metadata.annotation[]	List	 	自定义注释列表
Spec	Object	1	
Pod中容器的详细定义

spec.containers[]	List	1	Pod中的容器列表
spec.containers[].name	String	1	容器的名称，需要符合RFC 1035规范
spec.containers[].image	String	1	容器的镜像名称
spec.containers[].imagePullPolicy	String	 	
镜像拉取策略，可选值包括：Always、Nerver、IfNotPresent,默认值为Always。

(1) Always.表示每次都尝试重新拉取镜像。

(2)IfNotPresent:表示如果本地有该镜像，则使用本地的镜像，本地不存在时拉取镜像。

(3)Nerver:表示仅使用本地镜像。

包含如下设置，系统默认设置为Always，如下所述

(1)不设置imagePullPolicy,也未指定镜像的tag;

(2)不设置imagePullPolicy,镜像tag为latest

(3)启用名为AlwaysPullImages的准入控制器（Admission Controller）

spec.containers[].command[]	List	 	容器的启动命令列表，如果不指定，则使用镜像打包时使用的启动命令
spec.containers[].args[]	List	 	容器的启动命令参数列表
spec.containers[].workingDir	String	 	容器的工作目录
spec.containers[].volumeMounts[]	List	 	挂载到容器内部的存储卷配置
spec.containers[].volumeMounts[].name	String	 	引用Pod定义的共享存储卷的名称，需要使用volumes[]部分定义的共享存储卷名称
spec.containers[].volumeMounts[].mountPath	String	 	存储卷在容器内Mount的绝对路径，应少于512个字符
spec.containers[].volumeMounts[].readOnly	Boolean	 	是否为只读模式，默认为读写模式
spec.containers[].ports[]	List	 	容器需要暴露的端口号列表
spec.containers[].ports[].name	String	 	端口的名称
spec.containers[].ports[].containerPort	Int	 	容器需要监听的端口号
spec.containers[].ports[].hostPort	Int	 	容器所在主机需要监听的端口号，默认与containerPort相同，设置hostPort时，同一台宿主机将无法启动该容器的第2份副本
spec.containers[].ports[].protocol	String	 	端口协议，支持TCP和UDP，默认值为TCP
spec.containers[].env[]	List	 	容器运行前需要设置的环境变量列表
spec.containers[].env[].name	String	 	环境变量的名称
spec.containers[].env[].value	String	 	环境变量的值
spec.containers[].resources	Object	 	资源限制和资源请求的设置
spec.containers[].resources.limits	Object	 	资源限制的设置
spec.containers[].resources.limits.cpu	String	 	CPU限制，单位为core数，将用于docker run --cpu-shares参数
spec.containers[].resources.limits.memory	String	 	内存限制，单位可以为MIB、GIB等。将用于docker run --memory
spec.containers[].resources.requests	Object	 	资源限制设置
spec.containers[].resources.requecsts.cpu	String	 	CPU请求，单位为core数，容器启动的初始可用数量
spec.containers[].resources.requests.memory	String	 	内存请求，单位可以为MIB、GIB等，容器启动的初始可用数量
spec.volumes[]	List	 	在该Pod上定义的共享存储卷列表
spec.volumes[].name	String	 	
共享存储卷名称，在一个Pod中每个存储卷定义一个名称，容器定义部分的containers[].volumeMounts[].name将应用改共享存储卷的名称。

volume的类型包括：emptyDir、hostPath、gcePersistentDisk、awsElasticBlockStore、gitRepo、sercret、nfs、iscsi、glusterfs、persistentVolumeClaim、rbd、flexVolume、cinder、cephfs、flocker、downwardAPI、fc、azureFile、configMap、vsphereVolume，可以定义多个Volume，每个Volume的name保持唯一。

spec.volumes[].emptyDir	Object	 	类型为emptyDir的存储卷，表示与Pod同生命周期的一个临时目录，其值为一个空对象：emptyDir{}
spec.volumes[].hostPath	Object	 	类型为hostPath的存储卷，表示挂载Pod所在宿主机的目录，通过volumes[].hostPath.path指定
spec.volumes[].hostPath.path	String	 	Pod所在主机的目录，将被用于容器中的mount的目录
spec.volumes[].secret	Object	 	类型为secret存储卷，表示挂载集群预定义的secret对象到容器内部
spec.volumes[].configMap	Object	 	类型为configMap的存储卷，表示挂载激情预定义的configMap对象到容器内部
spec.volumes[].livenessProbc	Object	 	对Pod內各容器健康检查的设置，当探测无响应几次之后，系统将自动重启该容器，可以设置的方法包括：exec、httpGet、和tcpSocket。对一个容器仅需设置一种健康检查方法。
spec.volumes[].livenessProbe.exec	Object	 	对Pod内各容器健康检查的设置，exec方式
spec.volumes[].livenessProbe.exec.command[]	String	 	exec方式需要制定的命令或者脚本
spec.volumes[].livenessProbe.httpGet	Object	 	对Pod内各种容器健康检查设置，HTTPGet方式，需要指定path、port
spec.voulumes[].livenessProbe.tcpSocket	Object	 	对Pod内各容器健康检查的设置，tcpSocket方式
spec.volumes[].livenessProbe.initiaDelaySeconds	Number	 	容器启动完成后首次探测的时间，单位为s
spec.volumes[].livenessProbe.timeoutSeconds	Number	 	对容器健康检查的探测等待响应的超时时间设置。单位为s，默认值为1s。如超过该超时时间设置，则将认为该容器不健康，会重启该容器。
spec.volumes[].livenessProbe.PeriodSeconds	Number	 	对容器健康检查的定期探测时间设置，单位为s，默认10s探测一次
spec.restartPolicy	String	 	
Pod的重启策略。可选值为Always，OnFailure，Never 默认值为Always。

（1）Always：Pod一旦终止运行，则无论容器是如何终止的，kubectl都将重启它

（2）OnFailure:只有Pod以非零退出码终止时，kubectl才会重启该容器。如果容器正常结束，则kubectl将不会重启它

（3）Never:Pod终止后，kubectl将退出码报告给Master，不会再重启该Pod

spec.nodeSelector	Object	 	设置Node的Label，以key-value格式指定，Pod将被调度到具有这些Label的Node上
spec.imagePullSecrets	Object	 	pull镜像时使用的Secret名称，以name：secretkey格式指定
spec.hostNetwork	Boolean	 	是否使用主机网络模式，默认值为false，设置为true表示容器使用宿主机网络，不再使用Docker网桥，该Pod将无法在同一台宿主机上启动第2个副本。