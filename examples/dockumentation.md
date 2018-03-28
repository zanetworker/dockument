
#  DOCKumentation 

Your application Dockerfile is DOCKumented below. 

## Exposed Ports

 

Below are the ports the should be exposed by your application:
### Port 1337

- **Name**: 1337
- **About**: The main endpoint of this service.
- **Scheme**: tcp
- **Protocol**: http



## Dependencies

 
Below are your container application dependencies:

### Dependency rabbit

- **Name**: rabbit
- **Image**: rabbit:latest (`docker pull rabbit:latest`)
- **About**: The rabbit
- **Ports**: [5271]
- **Mandatory**: true


### Dependency redis

- **Name**: redis
- **Image**: redis:latest (`docker pull redis:latest`)
- **About**: For caching results from OWM API.
- **Ports**: [6379]
- **Mandatory**: true



## Tests

## Command Tests
 
Below are your container application dependencies:

### Command Test apt-get upgrade

 - **Name**: apt-get upgrade
 
 - **Command**: apt-get
 - **Args**: [-qqs, upgrade]
 - **Expected Output**: true




## File Content Tests
 
Below are your container application dependencies:

### File Content Test Debian Sources

 - **Name**: Debian Sources
 - **Path**: /etc/apt/sources.list 
 - **Expected Contents**: ['.*httpredir\.debian\.org.*']
 - **Excluded Content**: ['.*gce_debian_mirror.*']



## File Existence Tests
 
Below are your container application dependencies:

### File Existence Test test

 - **Name**: test
 - **Path**: asdas 
 - **Should Exist**: asd
 - **With Permissions**: asd




## Meta Data Tests
 
Below are your container application dependencies:

### Meta Data Test
 - **Environment Variables**: foo:baz
 - **Exposed Ports**: 8080, 2345
 
 - **Cmd**: /bin/bash
 - **Workdir**: /app
 - **Volumes**: /test






## Environment Variables
 
Below are some important Environment Variables:

### Env OPENWEATHERMAP_APIKEY

- **Name**: OPENWEATHERMAP_APIKEY
- **About**: Access key for OpenWeatherMap. See http://openweathermap.org/appid for details.
- **Mandatory**: true




## Resources Required
  
  
Resources required by your container application to run with reasonable performance: 
- **CPU**: 2
- **Memory**: 3gb




## Tags / Metadata

Metadata and container tags:
- **go**: 1.9



## Misc Labels

Non-Standard DOCKument API labels:
- **Health**: test
- **Maintainer**: Adel Zaalouk



