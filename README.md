# Portal

[![Build Status](https://travis-ci.org/mewil/portal.svg?branch=master)](https://travis-ci.org/mewil/portal)

An social media web app powered by [React](https://reactjs.org/), [Gin](https://github.com/gin-gonic/gin), [gRPC](https://grpc.io/), and much more.

![alt text](https://github.com/mewil/portal/blob/master/feed.png "Portal Feed")

## Contents

* [Setup](#setup)

## Setup

1. [Install Docker Compose](https://docs.docker.com/compose/install/)
2. Clone this repo: `git clone https://github.com/mewil/portal`
3. Change directory to the deploy repo: `cd portal/deploy`
4. Start whatever environment you want
    * Development `docker-compose -f development.yml up -d`
        * **Your git repo will be linked to the development environment, so your local changes will be reflected with a container restart**
    * Production (more ENV data required) `docker-compose -f production.yml up -d`
        * **NOTE: This takes care of setting up NGINX AND LetsEncrypt with the appropriate hosts and auto renewal**
5. Access `http://localhost:8000` and start developing! To view application logs, run `docker-compose -f development.yml logs -f portal`
