# 目录结构
[参考](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)
```
├── api             
├── cmd             
├── configs
├── docs
├── init
├── internal
├── pkg
├── scripts
├── test
├── tools
└── web
```

# 环境前置准备

## 1.创建namespace

```
kubectl create ns dev-project
```



## 2.创建ConfigMap用于外部挂载配置文件

```
kind: ConfigMap
apiVersion: v1
metadata:
  name: k8s-go-deploy-test
  namespace: dev-project
data:
  config.toml: |
    [log]
    devMode=true
    level="Debug"
    path="./logs/"
    save=1

    [db]
    dbType="mysql"
    username="root"
    password="root"
    host="dev-mysql-elfm.dev-project"
    port=3306
    dB="dev_project"
    param="charset=utf8"
    showSql=true
    maxIdleConns=10
    maxOpenConns=10
    connMaxLifetime=10


    [project]
    serverPort=8080
    excludeAuthUrls=["/"]
    allowMaxDelay=-60

```



## 3.创建mysql deploy和service

### deploy

```
kind: StatefulSet
apiVersion: apps/v1
metadata:
  name: dev-mysql
  namespace: dev-project
  labels:
    app: dev-mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: dev-mysql
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: dev-mysql
      annotations:
        logging.kubesphere.io/logsidecar-config: '{}'
    spec:
      volumes:
        - name: host-time
          hostPath:
            path: /etc/localtime
            type: ''
        - name: volume-oke9hb
          persistentVolumeClaim:
            claimName: mysql-storage
        - name: volume-mevvj3
          configMap:
            name: mysql-cnf
            defaultMode: 420
      containers:
        - name: container-z0rakv
          image: 'mysql:5.7.35'
          ports:
            - name: tcp-3306
              containerPort: 3306
              protocol: TCP
            - name: tcp-33060
              containerPort: 33060
              protocol: TCP
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: root
          resources: {}
          volumeMounts:
            - name: host-time
              readOnly: true
              mountPath: /etc/localtime
            - name: volume-oke9hb
              mountPath: /var/lib/mysql
            - name: volume-mevvj3
              readOnly: true
              mountPath: /etc/mysql/conf.d
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          imagePullPolicy: IfNotPresent
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      dnsPolicy: ClusterFirst
      serviceAccountName: default
      serviceAccount: default
      securityContext: {}
      affinity: {}
      schedulerName: default-scheduler
  serviceName: dev-mysql-elfm
  podManagementPolicy: OrderedReady
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      partition: 0
  revisionHistoryLimit: 10

```

### service

```
kind: Service
apiVersion: v1
metadata:
  name: dev-mysql-elfm
  namespace: dev-project
  labels:
    app: dev-mysql
spec:
  ports:
    - name: tcp-3306
      protocol: TCP
      port: 3306
      targetPort: 3306
    - name: tcp-33060
      protocol: TCP
      port: 33060
      targetPort: 33060
  selector:
    app: dev-mysql
  clusterIP: None
  clusterIPs:
    - None
  type: ClusterIP
  sessionAffinity: None

```

