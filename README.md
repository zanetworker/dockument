[![Build Status](http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)](http://http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)

# DOCKument

DOCKument is a tool that helps you document your Dockerfiles in a consistent manner the way you would expose an API. Matthias Lübken gave an example on how to use "api" like labels to describe important information in the Dockerfile (e.g., resources used, exposed ports, dependencies, etc) [here](https://github.com/luebken/currentweather/blob/master/Dockerfile). This project makes use of the label patterns to fetch important data in Dockerfiles and create Dockumentation for it. 

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

# build dockument binary if you don't have go installed
make OS=<darwin|linux|windows> dry

# execute dokcument command for overview
./dockument
```

### Usage

```bash
DOCKument is a tool to auto-generate Documentation for container images (Dockerfiles).

Environment:
$DOCKUMENT_HOME          set an alternative DOCKument location for files. By default, these are stored in ~/.dockument

Usage:
  dockument [flags]
  dockument [command]

Available Commands:
  all         fetches all the important DOCKument labels from the Dockerfile
  create      create documentation for Dockerfiles
  deps        fetches dependencies from the Dockerfile
  envs        fetches exposed envs from the Dockerfile
  help        Help about any command
  other       fetch all other non-DOCKument labels out of Dockerfiles Dockerfile
  ports       fetches exposed ports from the Dockerfile
  resources   fetches resources from the Dockerfile
  tags        fetches tags from the Dockerfile
  version     get version

Flags:
  -h, --help             help for dockument
      --home string      location of your dockumeny config. Overrides $DOCKUMENT_HOME  (default "/Users/adelias/.dockument")
      --outpath string   target location for Dockerfile documentation (default "/Users/adelias/go/src/github.com/zanetworker/dockument/cmd")
```

With DOCKument you have the option to fetch important information about your Dockerfile direcly from the command line or create a markdown document that you can add to your git REPO. 

For example, if you type `./dockument deps --dockerfile "../examples/Dockerfile"`, you will get the following output: 

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

On the other hand, to create a markdown DOCKument, you can use `dockument create --dockerfile "../examples/Dockerfile" -o "$HOME/dockumentation.md"`. The output will be a file similar to the one [here]("https://github.com/zanetworker/dockument/blob/master/examples/dockumentation.md")

### Pre-reqs

You can use `dockument` to print or create a documentation of random metadata in your Dockerfile, however, that is not very useful. To utilize Dockument functionality, it is better to use the following patterns in your Dockerfiles: 

**For Dependencies**:

```dockerfile
LABEL api.DEPENDENCY.redis=""\
      api.DEPENDENCY.redis.image="redis:latest"\
      api.DEPENDENCY.redis.port="6379"\
      api.DEPENDENCY.redis.about="For caching results from OWM API."\
      api.DEPENDENCY.redis.mandatory="true"
```

**For Important ENVs**:

```dockerfile
LABEL api.ENV.OPENWEATHERMAP_APIKEY="" \
      api.ENV.OPENWEATHERMAP_APIKEY.about="Access key for OpenWeatherMap. See http://openweathermap.org/appid for details." \
      api.ENV.OPENWEATHERMAP_APIKEY.mandatory="true"
```

**For Exposed Ports**:

```dockerfile
LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.scheme="tcp" \
	    api.EXPOSE.1337.protocol="http"\
      api.EXPOSE.1337.about="The main endpoint of this service."
```

**For resources used by the container image**:

```dockerfile
LABEL api.RESOURCES.Memory="3gb"\
      api.RESOURCES.CPU="2"
```

**For Important Tags** 

```dockerfile
LABEL api.RESOURCES.Memory="3gb"\
      api.RESOURCES.CPU="2"
```

For a complete example, please have a look at [this Dockerfile](https://github.com/zanetworker/dockument/blob/master/examples/Dockerfile) or [this Dockerfile by luebken](https://github.com/luebken/currentweather/blob/master/Dockerfile)

## Running the tests


## Contributing

<!-- [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) -->
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use Semantic versioning, to bump the version, execute

```bash
make BUMP=major|minor|patch bumpversion
```

<!-- We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).  -->

## Authors

See also the list of [contributors](https://github.com/zanetworker/dockument/graphs/contributors) who participated in this project.

## TODO

- [ ] Testing
- [ ] Add support if requested for Image Dockumentation

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details
