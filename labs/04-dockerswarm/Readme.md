Dockerassemble

##Init a Swarm Cluster
```
$ docker swarm init
Swarm initialized: current node (i5zo9bokee5i6di31mcuh3puz) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-2g4fq4k607lix0ixz7p5p5242kyxgo9uf820fjnl77ug7vd5zq-0jqubfehx3b87eoz1lzj6xktq 192.168.65.3:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

##Display the nodes in the swarm cluster

```
$ docker node ls
ID                            HOSTNAME                STATUS              AVAILABILITY        MANAGER STATUS      ENGINE VERSION
i5zo9bokee5i6di31mcuh3puz *   linuxkit-025000000001   Ready               Active              Leader              18.09.0
```


##Get the node Ip Address

```
$ docker node inspect i5zo9bokee5i6di31mcuh3puz | jq  ".[].Status.Addr"
"192.168.65.3"
$ docker node inspect i5zo9bokee5i6di31mcuh3puz | jq -r  ".[].Status.Addr"
192.168.65.3
```
Check this [tutorial](https://shapeshed.com/jq-json/#how-to-pretty-print-json) to get started with JQ

## Deploy a stack on docker swarm using the compose file

```
$ docker stack deploy -c docker-compose-build.yml projectapp
Ignoring unsupported options: build

Creating network projectapp_default
Creating service projectapp_api
Creating service projectapp_db
```

## List the stack on the cluster
```
docker stack list
NAME                SERVICES            ORCHESTRATOR
projectapp          2                   Swarm
```

## Display services of the projectapp stack

```
docker stack services projectapp
ID                  NAME                MODE                REPLICAS            IMAGE                               PORTS
2jpyxp66a5zd        projectapp_api      replicated          1/1                 nelvadas/whalecome-projectapi:1.4   *:8080->5000/tcp
908ci3pqpspa        projectapp_db       replicated          1/1                 mongo:latest                        *:30000->27017/tcp
```

## Scaling

Scale projectapp_api to 02 replicas

```
$ docker service scale projectapp_api=2
projectapp_api scaled to 2
overall progress: 2 out of 2 tasks
1/2: running   [==================================================>]
2/2: running   [==================================================>]
verify: Service converged
```


Check containers instances

```
$ docker service ls
ID                  NAME                MODE                REPLICAS            IMAGE                               PORTS
2jpyxp66a5zd        projectapp_api      replicated          2/2                 nelvadas/whalecome-projectapi:1.4   *:8080->5000/tcp
908ci3pqpspa        projectapp_db       replicated          1/1                 mongo:latest                        *:30000->27017/tcp


$ docker ps
CONTAINER ID        IMAGE                               COMMAND                  CREATED              STATUS              PORTS               NAMES
b3b183ff1ad7        nelvadas/whalecome-projectapi:1.4   "./app"                  About a minute ago   Up About a minute   5000/tcp            projectapp_api.2.y325q0xrcekfzy5tyyk3ddqct
ab25592b27f9        nelvadas/whalecome-projectapi:1.4   "./app"                  About an hour ago    Up About an hour    5000/tcp            projectapp_api.1.x2l0ri0895l7lwre2j57hjdvm
6bfd922d345c        mongo:latest                        "docker-entrypoint.s…"   About an hour ago    Up About an hour    27017/tcp           projectapp_db.1.wx96cid05qak4ix838ju0virk



$ docker service ps projectapp_api
ID                  NAME                IMAGE                               NODE                    DESIRED STATE       CURRENT STATE               ERROR               PORTS
x2l0ri0895l7        projectapp_api.1    nelvadas/whalecome-projectapi:1.4   linuxkit-025000000001   Running             Running about an hour ago
y325q0xrcekf        projectapp_api.2    nelvadas/whalecome-projectapi:1.4   linuxkit-025000000001   Running             Running 4 minutes ago
```

Check container on nodes

```
$ docker node ps
ID                  NAME                IMAGE                               NODE                    DESIRED STATE       CURRENT STATE               ERROR               PORTS
wx96cid05qak        projectapp_db.1     mongo:latest                        linuxkit-025000000001   Running             Running about an hour ago
x2l0ri0895l7        projectapp_api.1    nelvadas/whalecome-projectapi:1.4   linuxkit-025000000001   Running             Running about an hour ago
y325q0xrcekf        projectapp_api.2    nelvadas/whalecome-projectapi:1.4   linuxkit-025000000001   Running             Running 5 minutes ago
```



Create a multi node docker swarm Cluster

Install docker machine https://docs.docker.com/machine/overview/
Install VirtualBox https://www.virtualbox.org/wiki/Downloads
Create a machine

docker-machine create dockeree-node1


ssh -i "aws-ubuntu-k8s-demo-cluster.pem" ubuntu@ec2-3-16-148-178.us-east-2.compute.amazonaws.com

sudo usermod -a -G docker $USER
