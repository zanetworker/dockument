# Go-Kubesanity

Go-Kubesanity is the go implementation of [kubesanity](https://github.com/brendandburns/kubesanity). It is a kubernetes validation tool, it can be used to validate the configuration of various kubernetes objects, for example pods, services, volumes, etc.

## Overview

Go-kubesanity implements the basic checks now highlighted in the original kubesanity project. However, we are planning to extend it to support more usecases. 



## Getting Started

To get started with Go-kubesanity, you can download the corresponding binary for your OS (Darwin, Linux, Windows). Or you can clone this repository and build the project locally.

### Clone and build

You need to install `make` and `Go` on your system before proceeding.

```bash
git clone https://github.com/zanetworker/go-kubesanity.git
cd go-kubesanity

# build go-kubesanity binary if you have go installed
make OS=<darwin|linux|windows> install

# execute kubesanity for command overview
kubesanity

# build kubesanity binary if you don't have go installed
make OS=<darwin|linux|windows> dry

# execute kubesanity command for overview
./kubesanity
```

### Usage

```bash
volumes, etc.

Environment:
$KUBESANITY_HOME          set an alternative Kubesanity location for files. By default, these are stored in ~/.kubesanity

Usage:
  kubesanity network [flags]

Flags:
      --checkDuplicatePodIP       if set to true, kubesanity will check for duplicate Pod IPs in all namespaces
      --checkDuplicateServiceIP   if set to true, kubesanity will check for duplicate Service IPs in all namespaces
  -h, --help                      help for network

Global Flags:
      --home string         location of your Kubesanity config. Overrides $KUBESANITY_HOME (default "/Users/user_name/.kubesanity")
      --kubeconfig string   location of your kuberntes config file (default "/Users/user_name/.kube/config")
```

## Docker

```bash
docker run -v ${HOME}/.kube:/root/.kube:ro zanetworker/go-kubesanity
```

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

See also the list of [contributors](https://github.com/zanetworker/go-kubesanity/graphs/contributors) who participated in this project.

## TODO

- [ ] Testing
- [ ] More usecases for network, storage, pods, services configurations. 

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details
