# Docker compose

In this section, we are going to package the whole application in a single docker compose file.


## Docker Compose with existing images

```
$ cd WhaleCome/labs/03-dockercompose
```


```
$ docker-compose version
docker-compose version 1.23.1, build b02f1306
docker-py version: 3.5.0
CPython version: 3.6.6
OpenSSL version: OpenSSL 1.1.0h  27 Mar 2018
```


```
docker-compose up -d
Creating network "03-dockercompose_default" with the default driver
Creating 03-dockercompose_db_1_947948f71fc8  ... done
Creating 03-dockercompose_api_1_1128ca695a35 ... done
```



```
docker ps
CONTAINER ID        IMAGE                               COMMAND                  CREATED              STATUS              PORTS                      NAMES
e4be5c64efe2        mongo                               "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:32770->27017/tcp   03-dockercompose_db_1_84f8821425ca
2e885d04797d        nelvadas/whalecome-projectapi:1.2   "./app"                  About a minute ago   Up About a minute   0.0.0.0:8080->5000/tcp     03-dockercompose_api_1_a4333a4b9f65
```


```
$ docker-compose images
             Container                         Repository              Tag       Image Id      Size
-----------------------------------------------------------------------------------------------------
03-dockercompose_api_1_a4333a4b9f65   nelvadas/whalecome-projectapi   1.2      b67bf7372521   6.12 MB
03-dockercompose_db_1_84f8821425ca    mongo                           latest   525bd2016729   365 MB
```


```
$ docker-compose top
03-dockercompose_api_1_a4333a4b9f65
 PID    USER   TIME   COMMAND
-----------------------------
40367   root   0:00   ./app

03-dockercompose_db_1_84f8821425ca
 PID    USER   TIME         COMMAND
------------------------------------------
40428   999    0:00   mongod --bind_ip_all
```

Stop the containers

```
$ docker-compose down
Stopping 03-dockercompose_db_1_84f8821425ca  ... done
Stopping 03-dockercompose_api_1_a4333a4b9f65 ... done
Removing 03-dockercompose_db_1_84f8821425ca  ... done
Removing 03-dockercompose_api_1_a4333a4b9f65 ... done
Removing network 03-dockercompose_default
```


## Docker Compose with builds



```
$ docker-compose -f docker-compose-build.yml  up -d
uccessfully built 1f7aa498d751
Successfully tagged nelvadas/whalecome-projectapi:1.4
WARNING: Image for service api was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating 03-dockercompose_api_1_e903d0645168 ... done
Creating 03-dockercompose_db_1_bb3ce32c0337  ... done
```


```
$ docker ps
CONTAINER ID        IMAGE                               COMMAND                  CREATED              STATUS              PORTS                      NAMES
44384e10c7dd        mongo                               "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:32772->27017/tcp   03-dockercompose_db_1_e608772ee8e3
a5fa8b774e6a        nelvadas/whalecome-projectapi:1.4   "./app"                  About a minute ago   Up About a minute   0.0.0.0:8080->5000/tcp     03-dockercompose_api_1_87b719814cfc
```




```
$ docker-compose -f docker-compose-build.yml  build
db uses an image, skipping

```


```
$ docker-compose -f docker-compose-build.yml  stop
Restarting 03-dockercompose_api_1_dd56af0df581 ... done
Restarting 03-dockercompose_db_1_5b6e23cdfa03  ... done
```


```
$ docker-compose -f docker-compose-build.yml  up -d
$ http  localhost:8080/healthz
HTTP/1.1 200 OK
Content-Length: 5
Content-Type: text/plain; charset=utf-8
Date: Fri, 07 Dec 2018 15:01:46 GMT

"ok"
```
