# DOKI

DOKI provides a simple interface to generate html documentation from swagger yaml files.

## Prerequisite
- [Bootprint swagger](https://github.com/nknapp/bootprint-swagger)
- [Gox](https://github.com/mitchellh/gox)
- [Docker](https://docs.docker.com/)

## Create Documentation
A makefile using docker is available on the project.
Just run:
```
    make docs
```
This will go through the ./project folder find the yaml file and generate the documentation in the ./docs folder.

## Browse Swagger Documentation
Run:
```
	docker-compose up
```
Linux
```
    http://localhost:8080
```
Mac Os:
```
	http://$DockerMachineIp:8080
```

![alt tag](https://raw.github.com/evonck/doki/master/img/test.png)
