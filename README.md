# Portal

[![Build Status](https://travis-ci.org/mewil/portal.svg?branch=master)](https://travis-ci.org/mewil/portal)

A social media web app powered by [React](https://reactjs.org/), [Gin](https://github.com/gin-gonic/gin), [gRPC](https://grpc.io/), and much more.

- [Portal](#portal)
  - [Homepage](#homepage)
  - [Overview](#overview)
    - [Service Architecture](#service-architecture)
    - [Frontend Service](#frontend-service)
      - [Web App](#web-app)
      - [REST API](#rest-api)
    - [Auth Service](#auth-service)
    - [File Service](#file-service)
    - [User Service](#user-service)
    - [Post Service](#post-service)
  - [Setup](#setup)

## Homepage

![alt text](https://github.com/mewil/portal/blob/master/feed.png "Portal Feed")

## Overview

### Service Architecture

The portal application is split into services: [auth](https://github.com/mewil/portal/tree/master/services/auth_service), [file](https://github.com/mewil/portal/tree/master/services/file_service), [frontend](https://github.com/mewil/portal/tree/master/services/frontend_service), [user](https://github.com/mewil/portal/tree/master/services/user_service), and [post](https://github.com/mewil/portal/tree/master/services/post_service). These services communicated via [gRPC](https://grpc.io/) and Protocol Buffer messages, defined [here](https://github.com/mewil/portal/tree/master/services/pb). All code not run in the browser is written in [Go](https://golang.org/), but since each service runs in its own isolated [Docker](https://www.docker.com/) container and use gRPC, they could be implemented in any supported language. The application uses the RDBMS [MySQL](https://www.mysql.com/) and Object Storage Platform [Minio](https://min.io/) for data storage.

### Frontend Service

The frontend service is split into two parts, a web app and a REST API that provides clients controlled access to the other services.

#### Web App

The web app is written with [React](https://reactjs.org/), and uses a package-based architecture. Each package [here](https://github.com/mewil/portal/tree/master/services/frontend_service/app), has a specific function, and contains all the state and UI management code needed for that function, or it is used as a build block for other packages. The app follows the [Flux](https://facebook.github.io/flux/) architecture, and uses [Redux](https://redux.js.org/), [Redux Saga](https://github.com/redux-saga/redux-saga), and a number of other libraries to accomplish this. It also uses [styled components](https://www.styled-components.com/) instead of stylesheets.

#### REST API

The REST API is a Go app that uses the [Gin](https://github.com/gin-gonic/gin) framework for routing. It serves the web app and provides routes for accessing user and post data here. These routes validate requests and call controllers that make necessary gRPC calls to the other services.

### Auth Service

The auth service manages hashing and storing user passwords and admin data. It uses a MySQL table to manage this data, and has gRPC server handlers for accepting requests.

### File Service

The file service manages uploading and fetching files. It uses Minio, an open source object store, and gRPC streams for its file data to accomplish this.

### User Service

The user service manages user data such as name, description, and followers/following and fetching files. It uses two MySQL tables to manage this data, and has gRPC server handlers for accepting requests.

### Post Service

The post service manages user's posts, comments, and likes. It uses several MySQL tables to manage this data, and has gRPC server handlers for accepting requests.

## Setup

1. [Install Docker Compose](https://docs.docker.com/compose/install/)
2. Clone this repo: `git clone https://github.com/mewil/portal`
3. Change directory to the deploy repo: `cd portal/deploy`
4. Start whatever environment you want
    - Development `docker-compose -f development.yml up -d`
        - **Your git repo will be linked to the development environment, so your local changes will be reflected with a container restart**
    - Production (more ENV data required) `docker-compose -f production.yml up -d`
        - **NOTE: This takes care of setting up NGINX AND LetsEncrypt with the appropriate hosts and auto renewal**
5. Access `http://localhost:8000` and start developing! To view application logs, run `docker-compose -f development.yml logs -f`
