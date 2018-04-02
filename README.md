[![Build Status](http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)](http://http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)

# DOCKument

DOCKument is a tool that helps you auto-generate documentation for your Dockerfiles and Docker images the way you would an API.

Matthias Lübken gave an example on how to use "API" like labels to describe important information in the Dockerfile (e.g., resources used, exposed ports, dependencies, etc) [here](https://github.com/luebken/currentweather/blob/master/Dockerfile). This project makes use of similar label patterns to fetch important data in Dockerfiles / Docker images and create Dockumentation for it automagically. Furthermore, it improves on the work done by Matthias to expose various kinds of information that can be fetched from Docker images.



Check the blog post [here](http://www.adelzaalouk.me/2018/dockument/)

## Getting Started

To get started with `Dockument`, you can download the corresponding binary for your OS (Darwin, Linux, Windows) from the releases section. Or you can clone this repository and build the project locally.

### Clone and build

You need to install `make` and `Go` on your system before proceeding.

```bash
git clone https://github.com/zanetworker/dockument.git
cd dockument

# build dockument binary if you have go installed
make OS=<darwin|linux|windows> install

# execute dockument for command overview
dockument

# build dockument binary 
make OS=<darwin|linux|windows> dry

# execute dokcument command for overview
./dockument
```

### Usage

[![asciicast](https://asciinema.org/a/pdVCkPpvz7OAdjYpuOz4gMY1j.png)](https://asciinema.org/a/pdVCkPpvz7OAdjYpuOz4gMY1j)

With DOCKument you have the option to fetch important information about your Dockerfile or Docker Image direcly from the command line or create a markdown document that you can add to your git REPO.


**Fetching Information from a Dockerfile or a Docker Image**

For example, if you type `dockument deps --dockerfile "../examples/Dockerfile"` or `dockument deps --image "repo/name:tag"`, you will get the following output:

```bash
### Dependency RABBIT ###
	Application: rabbit
	Image: "rabbit:latest"
	Description: "The rabbit"
	Ports: ["5271"]
	Required: "true"

### Dependency REDIS ###
	Application: redis
	Image: "redis:latest"
	Description: "For caching results from OWM API."
	Ports: ["6379"]
	Required: "true"
``` 

On the other hand, to create a markdown DOCKument, you can use `dockument create --dockerfile "../examples/Dockerfile" -o "$HOME/dockumentation.md"` or `dockument create --image "repo/name:tag" -o "$HOME/dockumentation.md"` . The output will be a file similar to the one [here]("https://github.com/zanetworker/dockument/blob/master/examples/dockumentation.md")

**Creating a Dockumentation markdown to describe your image or Dockerfile** 

To create Dockumentation for a Dockerfile, you can simply type (as an example):

```bash
dockument create --dockerfile "examples/Dockerfile" -o "$HOME/dockumentation.md"
```

and for a Docker Image

```bash
dockument create -i "zanetworker/document:latests" -o "$HOME/dockumentation.md"
```

For more information, use `dockument help` to find out about the available commands.


### Pre-reqs

There are two types of labels that `Dockument` supports, these are: 

- **Misc Lables**: these are random / non-standard  labels defined in your Docker image, basically any label that does not follow the pattern described below.
- **Standardized Labels**: These are the gems of `Dockument`, these labels are the ones used to expose the important information about your Dockerfile / Docker image. They follow a pattern, always prefiexed with **api.Type.***

You can use `Dockument` to print or create a documentation of random metadata in your Dockerfile or Docker images, however, that is not very useful. To utilize the fully-fledged functionality of Dockument, define labels in your Docker images that follow the below pattern.

**For Dependencies**:

Will this containerized application have dependencies on other applications?

```dockerfile
LABEL api.DEPENDENCY.redis=""\
      api.DEPENDENCY.redis.image="redis:latest"\
      api.DEPENDENCY.redis.port="6379"\
      api.DEPENDENCY.redis.about="For caching results from OWM API."\
      api.DEPENDENCY.redis.mandatory="true"
```

**For Important ENVs**:

What are the environment variables that this container image defines? 

```dockerfile
LABEL api.ENV.OPENWEATHERMAP_APIKEY="" \
      api.ENV.OPENWEATHERMAP_APIKEY.about="Access key for OpenWeatherMap. See http://openweathermap.org/appid for details." \
      api.ENV.OPENWEATHERMAP_APIKEY.mandatory="true"
```

**For Exposed Ports**:

Do we expose any ports? 

```dockerfile
LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.scheme="tcp" \
	  api.EXPOSE.1337.protocol="http"\
      api.EXPOSE.1337.about="The main endpoint of this service."
```

**For resources used by the container image**:

What resources are required by my application to run in a smooth and performant manner?

```dockerfile
LABEL api.RESOURCES.Memory="3gb"\
      api.RESOURCES.CPU="2"
```

**For Important Tags** 

What are some important tags that I want my Dockerfile to expose?

```dockerfile
LABEL api.TAGS.go="1.9"
```

**For Container Tests** 

Tests help you make sure that your container image conforms with what you actually wants it to do. To this end, we decided to support the [container-structured-test](https://github.com/GoogleCloudPlatform/container-structure-test) style. There are four types that are supported by Dockument:

- **Command Tests:** "Command tests ensure that certain commands run properly in the target image".


```dockerfile
LABEL api.TEST.command=""\
      api.TEST.command.name="go version"\
      api.TEST.command.command="go"\
      api.TEST.command.args="version"\
      api.TEST.command.expectedOutput="go version"
```

- **File Existence Tests:** "File existence tests check to make sure a specific file (or directory) exist within the file system of the image."

```dockerfile
LABEL api.TEST.fileExistence=""\
      api.TEST.fileExistence.name="Dockumentation Check"\
      api.TEST.fileExistence.path="/dockumentation.md"\
      api.TEST.fileExistence.shouldExist="true"\
      api.TEST.fileExistence.permissions=""
```

- **File Content Tests:** "File content tests open a file on the file system and check its contents." 

```dockerfile
LABEL api.TEST.fileContent=""\
      api.TEST.fileContent.name="Debian Sources"\
      api.TEST.fileContent.path="/etc/apt/sources.list"\
      api.TEST.fileContent.expectedContents="['.*httpredir\\.debian\\.org.*']"\
      api.TEST.fileContent.excludedContents="['.*gce_debian_mirror.*']"
``` 

- **Meta data tests**: "The Metadata test ensures the container is configured correctly". 

```dockerfile
LABEL api.TEST.metadata=""\
      api.TEST.metadata.env="GOPATH:/go"\
      api.TEST.metadata.exposedPorts=""\
      api.TEST.metadata.volumes=""\
      api.TEST.metadata.cmd=""\
      api.TEST.metadata.workdir=""
```


For a complete example, please have a look at [this Dockerfile](https://github.com/zanetworker/dockument/blob/master/examples/Dockerfile) or [this Dockerfile by luebken](https://github.com/luebken/currentweather/blob/master/Dockerfile)


## Contributing

<!-- [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) -->
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

See also the list of [contributors](https://github.com/zanetworker/dockument/graphs/contributors) who participated in this project.

## TODO

[ ] Testing

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details
