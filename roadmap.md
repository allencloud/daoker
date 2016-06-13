# Roadmap

This document describes some important features daoker is going to realize.

## Container

- [x] list all containers
- [x] kill, stop a container
- [ ] inspect a container
- [x] check if a container is under oom
- [ ] get dns details of a container
- [ ] get hosts details of a container
- [ ] get hostname details of a container

## Image

- [ ] tar a specific single layer of image 
- [ ] pull image via distribution api (not via docker daemon)


## Volume

- [x] get volume details of a container
- [x] get a specific data volume disk space usage
- [x] get a specific data volume inode usage


## Log

- [ ] get logs of a container
- [x] get size of a container's log, only json-file supported
- [x] add some log into container's log by force


## Processes

- [x] check if given pid is in a container or not
- [x] get sum of processes in a container
- [ ] get the main process pid number of a container
- [ ] print containers whose process number exceeds a specified number


