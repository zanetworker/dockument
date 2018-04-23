FROM golang:alpine as builder
 

# other labels 
LABEL Maintainer="Adel Zaalouk"
LABEL Health="test"

# Descibe Image Tags
LABEL api.TAGS.go="1.9"

# Describe Image minimum resource requirements 
LABEL api.RESOURCES.Memory="3gb"\
      api.RESOURCES.CPU="2"

###  Describe container tests based on container structure tests from google
LABEL api.TEST.command=""\
      api.TEST.command.name="apt-get upgrade"\
      api.TEST.command.command="apt-get"\
      api.TEST.command.args="-qqs, upgrade"\
      api.TEST.command.expectedOutput="true"\
      api.TEST.command.expectedError=""\
      api.TEST.command.excludedOutput=""\
      api.TEST.command.excludedError=""

LABEL api.TEST.fileContent=""\
      api.TEST.fileContent.name="Debian Sources"\
      api.TEST.fileContent.path="/etc/apt/sources.list"\
      api.TEST.fileContent.expectedContents="['.*httpredir\\.debian\\.org.*']"\
      api.TEST.fileContent.excludedContents="['.*gce_debian_mirror.*']"

LABEL api.TEST.metadata=""\
      api.TEST.metadata.env="foo:baz"\
      api.TEST.metadata.exposedPorts="8080, 2345"\
      api.TEST.metadata.volumes="/test"\
      api.TEST.metadata.cmd="/bin/bash"\
      api.TEST.metadata.workdir="/app"

###  Describe container dependencies
LABEL api.DEPENDENCY.redis=""\
      api.DEPENDENCY.redis.image="redis:latest"\
      api.DEPENDENCY.redis.port="6379"\
      api.DEPENDENCY.redis.about="For caching results from OWM API."\
      api.DEPENDENCY.redis.mandatory="true"

LABEL api.DEPENDENCY.rabbit=""\
      api.DEPENDENCY.rabbit.image="rabbit:latest"\
      api.DEPENDENCY.rabbit.port="5271"\
      api.DEPENDENCY.rabbit.about="The rabbit"\
      api.DEPENDENCY.rabbit.mandatory="true"

###  Set and describe available ENVs
LABEL api.ENV.OPENWEATHERMAP_APIKEY="" \
      api.ENV.OPENWEATHERMAP_APIKEY.about="Access key for OpenWeatherMap. See http://openweathermap.org/appid for details." \
      api.ENV.OPENWEATHERMAP_APIKEY.mandatory="true"
ENV OPENWEATHERMAP_APIKEY 123456

###  Desrbite ports exposed by this services and the protocl used
LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.scheme="tcp" \
      api.EXPOSE.1337.protocol="http"\
      api.EXPOSE.1337.about="The main endpoint of this service."

# Add file and check it's contents
LABEL api.TEST.fileExistence=""\
      api.TEST.fileExistence.name="Dockumentation Check"\
      api.TEST.fileExistence.path="/dockumentation.md"\
      api.TEST.fileExistence.shouldExist="true"\
      api.TEST.fileExistence.permissions=""
ADD dockumentation.md /dockumentation.md

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

