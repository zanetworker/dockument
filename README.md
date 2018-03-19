[![Build Status](http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)](http://http://e812a0e6.ngrok.io/api/badges/zanetworker/dockument/status.svg?branch=master)

# Dockument


## Overview

Go-kubesanity implements the basic checks now highlighted in the original kubesanity project. However, we are planning to extend it to support more usecases. 

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
