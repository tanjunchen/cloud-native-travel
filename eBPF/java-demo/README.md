# Java 测试应用程序

使用以下命令编译应用程序
```bash
mvn clean install
```

使用以下命令打包应用程序
```bash
mvn package
```

使用以下命令运行应用程序
```
java -jar target/java-demo.jar
```

# 主要功能

```bash
“/”：根URL，当用户访问此URL时，会返回一个字符串"helloworld"，并且在日志中记录这个信息。
“/hello”：当用户访问"/hello"这个URL时，也会返回一个字符串"helloworld"，并且在日志中记录这个信息。
“/ebpf/function/{message}”：这个URL需要一个路径参数"{message}“。根据提供的”{message}"，这个方法将执行不同的操作：
	如果"{message}“等于"count”，那么它会执行一个名为"functionCount"的方法，该方法将打印从0到99999的所有整数。
	如果"{message}“等于"exception”，那么它会尝试执行一个会引发异常的操作（除以零），并捕获这个异常，然后在日志中记录异常信息。
	如果"{message}“等于"latency”，那么它将使当前的线程休眠5秒，模拟延迟。
	如果"{message}"不是上述任何一种情况，那么它将返回一个空字符串。
```
