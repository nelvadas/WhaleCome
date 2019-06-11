# Kubernetes labs with Docker EE

```
$kubectl config get-context
CURRENT   NAME                  CLUSTER                      AUTHINFO              NAMESPACE
*         docker-for-desktop    docker-for-desktop-cluster   docker-for-desktop
          enono.k8s.local       enono.k8s.local              enono.k8s.local
          k8s-dev.abyster.com   k8s-dev.abyster.com          k8s-dev.abyster.com
```

 use the docker context
```          
kubectl config use-context docker-for-desktop
```

Create a myapp namespace

```
kubectl create namespace myapp
```

Create a service in this namespace

$ docker stack deploy --orchestrator=kubernetes --namespace=myapp --compose-file=docker-compose.yml myappstack
Waiting for the stack to be stable and running...
api: Ready		[pod status: 1/1 ready, 0/1 pending, 0/1 failed]
db: Ready		[pod status: 1/1 ready, 0/1 pending, 0/1 failed]

Stack myappstack is stable and running


Inspect kubernetes stacks
```
$ docker stack ls  --orchestrator=kubernetes
NAME                SERVICES            ORCHESTRATOR        NAMESPACE
myappstack          2                   Kubernetes          myapp
```
Check pods

```
$ kubectl get pods -n myapp
NAME                   READY     STATUS    RESTARTS   AGE
api-5fc95cf67b-ff6xk   1/1       Running   0          2m
db-6c475f59b5-cdbpr    1/1       Running   0          2m
```

Check services
```
```
