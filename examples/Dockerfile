FROM golang:alpine as builder
 


# Descibe Image Tags
LABEL api.TAGS.go="1.9"

# Describe Image minimum resource requirements 
LABEL api.RESOURCES.Memory=""\
      api.RESOURCES.CPU=""\

# Describe container dependencies
LABEL api.DEPENDENCY.redis=""\
      api.DEPENDENCY.redis.image="redis:latest"\
      api.DEPENDENCY.redis.port="6379"\
      api.DEPENDENCY.redis.about="For caching results from OWM API."\
      api.DEPENDENCY.redis.mandatory="true"

# Set and describe available ENVs
LABEL api.ENV.OPENWEATHERMAP_APIKEY="" \
      api.ENV.OPENWEATHERMAP_APIKEY.about="Access key for OpenWeatherMap. See http://openweathermap.org/appid for details." \
      api.ENV.OPENWEATHERMAP_APIKEY.mandatory="true"

# Desrbite ports exposed by this services and the protocl used
LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.scheme="tcp" \
	api.EXPOSE.1337.protocol="http"
      api.EXPOSE.1337.about="The main endpoint of this service."

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

