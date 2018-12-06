# Dockerfiles
This section deals with building images using Dockerfiles.
By default, each build will be done with the current directory as context.
You can pass a different context to you build, context should contains only the files needed to perform the build, keep it simple as possible.
You can also use a .dockerignore file at the root of you context to ignore some file using various regexp.

```
$ cd WhaleCome/projectsvc
```

## Build prototype

```
docker build  --no-cache -f ../labs/02-dockerfiles/Dockerfile-basic -t whalecome/projectapi:1.0 .
```

Display images built using the Dockerfile-basic command

```
$ docker images --filter "label=com.docker.training.app=projectsvc"
REPOSITORY             TAG                 IMAGE ID            CREATED             SIZE
whalecome/projectapi   1.0                 6c0800b80437        10 seconds ago      739MB
```


## Multistage builds
you can leverage multistage builds to produce a lightweight image like this
In the fist stage of the dockerfile, we built the binary sources, and in a second stage , we copy the result binary in a base runtime image

```
docker build  --no-cache -f ../labs/02-dockerfiles/Dockerfile-onbuild -t whalecome/projectapi:1.1 .
```
With this approach, the image size is just 6MB compared to 739MB without multistage featrue.
```
docker images --filter "label=com.docker.training.app=projectsvc"
REPOSITORY             TAG                 IMAGE ID            CREATED             SIZE
whalecome/projectapi   1.1                 b67bf7372521        3 seconds ago       6.41MB
whalecome/projectapi   1.0                 6c0800b80437        2 minutes ago       739MB
```

## Build args and Cache
The ARG directive in a Dockerfile is very is very usefull to pass parameters to the build

```
ARG  BASE_IMAGE_TAG=default-value
FROM golang:${BASE_IMAGE_TAG} AS builder

```

To build the project with  build argument, use the ---build-arg option

```
docker build   -f ../labs/02-dockerfiles/Dockerfile-onbuild-arg --build-arg BASE_IMAGE_TAG=onbuild -t whalecome/projectapi:1.2 .
Sending build context to Docker daemon  8.678kB
Step 1/8 : ARG  BASE_IMAGE_TAG=onbuild
Step 2/8 : FROM golang:${BASE_IMAGE_TAG} AS builder
# Executing 3 build triggers
 ---> Using cache
 ---> Using cache
 ---> Using cache
 ---> 0eab77253aae
Step 3/8 : RUN  env GOOS=linux GOARCH=386 go build -o /go/bin/app
 ---> Using cache
 ---> a9a3a5db392d
Step 4/8 : FROM scratch
 --->
Step 5/8 : LABEL com.docker.training.unit=docker101       com.docker.training.app=projectsvc
 ---> Using cache
 ---> a59ba069b337
Step 6/8 : COPY --from=builder /go/bin/app .
 ---> Using cache
 ---> db6ab520cc6d
Step 7/8 : EXPOSE 5000
 ---> Using cache
 ---> 9e5646fce0de
Step 8/8 : ENTRYPOINT ["./app"]
 ---> Using cache
 ---> b67bf7372521
Successfully built b67bf7372521
Successfully tagged whalecome/projectapi:1.2
```
This third build is more faster because of cache , for each instruction in the Dockerfile the build engine first
attempt to find if there is one image build from the same line.

```
$ docker images --filter "label=com.docker.training.app=projectsvc"
REPOSITORY             TAG                 IMAGE ID            CREATED             SIZE
whalecome/projectapi   1.1                 b67bf7372521        5 minutes ago       6.41MB
whalecome/projectapi   1.2                 b67bf7372521        5 minutes ago       6.41MB
whalecome/projectapi   1.0                 6c0800b80437        8 minutes ago       739MB
```


## Pushing your image to docker hub.
Log on docker.io using your credentials
```
docker login -u xxxx docker.io
```

Tag you image   as docker.io/<userid>/<repository>

```
$ docker tag whalecome/projectapi:1.2 docker.io/nelvadas/whalecome-projectapi:1.2
```

Push the new tag  to docker.io
```
docker push docker.io/nelvadas/whalecome-projectapi:1.2
The push refers to repository [docker.io/nelvadas/whalecome-projectapi]
7ee41395f021: Pushed
```
