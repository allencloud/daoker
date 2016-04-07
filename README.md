# Daoker
Daoker is a tool to manage Docker environment regardless of docker daemon's failure.

Docker is an open source project to build, ship and run any application as a 
lightweight container. If you have some experience in operating Docker in developing,
testing or production environment, you will find that the architecture of Docker is
composed of **four parts: Docker CLI, Docker Daemon, Docker Images and Docker Container**.

## Background
I believe that Docker is totally production-ready, while we still run into some cases
that Docker Daemon does not work as we wish, such as :

* command `docker ps` hangs
* command `docker inspect` hangs
* no details about docker environment, no right to talk about management
* and so on

If situations above happen to you, we can judge that you lose the control of docker.
If it is your developing environment, you can change a machine with Docker to avoid
some problems. And if it is production environment, I think everyone should realized
the gravity of this situation.

## Principles
As Docker Daemon fails, we have no idea about whether docker api is working which
seems so disappointing. **Disappointing, but not fatal**.

Have you ever experienced that Docker Daemon stores almost all details about Docker
Container in its local filesystem? If the answer is "Yes", you've got what I am
going to do.

Following is two most important principles:

* **NEVER** to contact with docker daemon
* take advantage of **DOCKER_ROOT** and **CGROUP FILESYSTEM** to do everything you want

## Get Started

## Installation

## Building Daoker

## Participating

## Copyright and license



