[![Build Status](http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)](http://http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)

# DOCKument

DOCKument is a tool that helps you document your Dockerfiles in a consisteny manner the way you would expose an API. 

## Getting Started

To get started with `Dockument`, you can download the corresponding binary for your OS (Darwin, Linux, Windows). Or you can clone this repository and build the project locally.

### Clone and build

You need to install `make` and `Go` on your system before proceeding.

```bash
git clone https://github.com/zanetworker/dockument.git
cd dockument

# build dockument binary if you have go installed
make OS=<darwin|linux|windows> install

# execute dockument for command overview
dockument

# build kubesanity binary if you don't have go installed
make OS=<darwin|linux|windows> dry

# execute kubesanity command for overview
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
## Docker


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
- [ ] More usecases for network, storage, pods, services configurations. 

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details
