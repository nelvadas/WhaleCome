
#Table of Contents
1. [Introduction](#introduction)
2. [Getting the source](#setup)
3. [Testing the API](#testing)


# Introduction <a href=introduction></a>

The main goal of this blog post series is to build a Timesheet management application.
In this first lab, we will build a simple block of a global Timesheet management API.
we will create a simple REST API that will be used by admins to create new project on which users can create timecards and timesheets.
To do so, we will start by setting in place a simple REST API for admins to create new projects in a MongoDB document store


## Setup the API<a name="setup"></a>

Use the following instructions to checkout and start the API on your computer
```
 cd $GOPATH/scr/
 https://github.com/nelvadas/WhaleCome.git
 cd WhaleCome/projectsvc


```


The applications start and listen on port 5000,
and tries to connect to the default MongoDB instance on your computer. For now we will use this default installation.
For now we will rely on these default settings, in future labs, we will present how to pass these configuration details to you
docker application.

```
$ go run main.go

2018/12/05 10:08:36 MONGO_URL=127.0.0.1:27017
2018/12/05 10:08:36 MONGO_DATABASE=local
2018/12/05 10:08:36 Connecting...
2018/12/05 10:08:36 Connected to DB local
2018/12/05 10:08:36 Starting on port 5000

```

If you do not have a mongo db running, you can use the following docker command to start one

```
$ docker run -d -p 27017:27017 mongo
```

Now your REST API and the Mongo DB are up and running, In the next section we will see how interact with the API

## Testing <a name="testing"></a>

### Insert a new project

```
$ curl -v -X POST -H "Content-Type: application/json"  \
    --data '{"code": "WH001", "description": "WhaleCame to Docker ", "from": "2018-11-27T09:00:43Z", "to": "2018-12-31T18:30:40Z" }' \
     localhost:5000/project

     --
     * upload completely sent off: 119 out of 119 bytes
< HTTP/1.1 200 OK
< Date: Wed, 05 Dec 2018 09:46:41 GMT
< Content-Length: 144
< Content-Type: text/plain; charset=utf-8
<
{"id":"5c079e816c92c078bd2120d8","code":"WH001","description":"WhaleCame to Docker ","from":"2018-11-27T09:00:43Z","to":"2018-12-31T18:30:40Z"}

```

Insert a second project to track DockerCon EU18 activities and on for Internal Trainings


```
$ curl -v -X POST -H "Content-Type: application/json"  \
    --data '{"code": "DC018", "description": "DockerCon EU 2018 ", "from": "2018-12-01T09:00:43Z", "to": "2018-12-10T18:30:40Z" }' \
     localhost:5000/project

    {"id":"5c07a0886c92c078bd2120d9","code":"DC018","description":"DockerCon EU 2018 ","from":"2018-12-01T09:00:43Z","to":"2018-12-10T18:30:40Z"}

$ curl -v -X POST -H "Content-Type: application/json"  localhost:5000/project  \
        --data '{"code": "INT00", "description": "Internal Training ", "from": "2013-01-01T09:00:43Z" }'

      {"id":"5c07a1286c92c078bd2120da","code":"INT00","description":"Internal Training ","from":"2013-01-01T09:00:43Z","to":"0001-01-01T00:00:00Z"}   

```


### Get a specific project from it identifier

```
$ http  localhost:5000/project/5c07a0886c92c078bd2120d9
HTTP/1.1 200 OK
Content-Length: 142
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Dec 2018 10:03:28 GMT

{
    "code": "DC018",
    "description": "DockerCon EU 2018 ",
    "from": "2018-12-01T09:00:43Z",
    "id": "5c07a0886c92c078bd2120d9",
    "to": "2018-12-10T18:30:40Z"
}
```


### Get All projects

```
$ http  localhost:5000/project
HTTP/1.1 200 OK
Content-Length: 430
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Dec 2018 10:04:51 GMT

[
    {
        "code": "WH001",
        "description": "WhaleCame to Docker ",
        "from": "2018-11-27T09:00:43Z",
        "id": "5c079e816c92c078bd2120d8",
        "to": "2018-12-31T18:30:40Z"
    },
    {
        "code": "DC018",
        "description": "DockerCon EU 2018 ",
        "from": "2018-12-01T09:00:43Z",
        "id": "5c07a0886c92c078bd2120d9",
        "to": "2018-12-10T18:30:40Z"
    },
    {
        "code": "INT00",
        "description": "Internal Training ",
        "from": "2013-01-01T09:00:43Z",
        "id": "5c07a1286c92c078bd2120da",
        "to": "0001-01-01T00:00:00Z"
    }
]
```

That's all for the moment, [in the next labs](../02-dockerfiles), we will see how to containerize the following code using various Docker tools and commands.
