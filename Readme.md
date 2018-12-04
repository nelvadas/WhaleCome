
# Introduction <a href=introduction></a>




# Table of contents
1. [Introduction](#introduction)
2. [LAB01: Setup the REST API with Go and MongoDB](#restapi)
3. [LAB02: Using Dockerfiles](#dockerfiles)
4. [LAB03: Using Docker Compose](#dockercompose)
5. [LAB04: Using Docker Assemble](#dockerassemble)


## This is the introduction <a name="introduction"></a>
The main goal of this blog post series is to build a Timesheet management application.
To do so, we will start by setting in place a simple REST API for admins to create new projets in a MongoDB document store

## LAB02: Using Dockerfiles <a name="dockerfiles" href="./lab/lab01/"></a>
This lab demonstrate how to build and operate the Timesheet REST API on Docker platform using  Dockerfiles. Read

##LAB03: Using Docker Compose <a name="dockercompose"></a>
With Microservices comes composition, the third lab teach you how to use docker compose to assemble the REST API and its backend database in a single
deployment unit.

## LAB04: Using Docker Assemble <a name="dockerassemble"></a>
Simplifying the process to create a running artifact on Docker platform


## Setup the API



## Testing

### Insert a new project

```
$ curl -X POST -H "Content-Type: application/json"  --data '{"code": "PRJ003"}'  localhost:5000/project
```

### Get All projects

```
```

### Get a specific project details

```
```
