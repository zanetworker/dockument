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

LABEL api.TEST.fileExistence=""\
      api.TEST.fileExistence.name="test"\
      api.TEST.fileExistence.path="asdas"\
      api.TEST.fileExistence.shouldExist="asd"\
      api.TEST.fileExistence.permissions="asd"

LABEL api.TEST.fileContent=""\
      api.TEST.fileContent.name=""\
      api.TEST.fileContent.path=""\
      api.TEST.fileContent.expectedContents=""\
      api.TEST.fileContent.excludedContents=""

LABEL api.TEST.metadata=""\
      api.TEST.metadata.env="foo:baz"\
      api.TEST.metadata.exposedPorts=""\
      api.TEST.metadata.cmd=""\
      api.TEST.metadata.workdir=""

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

###  Desrbite ports exposed by this services and the protocl used
LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.scheme="tcp" \
	api.EXPOSE.1337.protocol="http"\
      api.EXPOSE.1337.about="The main endpoint of this service."

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

