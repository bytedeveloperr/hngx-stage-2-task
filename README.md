# HNGx Stage 2 Person Crud API

## Introduction

Simple REST Api that exposes CRUD endpoints to work with a `Person` resource. This project is made for the HNGx stage 2 task

### Tech Stack
- GOlang
- Gorm
- Gorilla Mux
- SQlite

### Application Features

* Create new person
* Get single person using person id in path param
* Update single person using person id in path param
* Delete single person using person id in path param

### Application Setup
To run this on your machine, ensure you have go installed, if [click here](https://go.dev/doc/install) to install it.

Once you've setup Go, clone the repo
```
git clone https://github.com/bytedeveloperr/hngx-stage-2-task.git
```

and download the project's packages

```
go mod download
```

#### Start API server
Use the command below to start the api server
```
go run .
```

#### Resources

##### Person
A struct that represent an individual person. It has two fields `name` and `id`. `name` is of type string while `id` is either a string or integer.

#### API Endpoints

- `GET /api` returns list of all available users
- `POST /api` creates a new user with the `name` passed in the request body
- `GET /api/:userid` returns the user with the `userid` in the request param
- `PUT /api/:userid` updates the user with the `userid` in the request param
- `DELETE /api/:userid` deletes the user with the `userid` in the request param
