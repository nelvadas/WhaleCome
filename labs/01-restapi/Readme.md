
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



## Testing <a name="testing"></a>

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
