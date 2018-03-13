FROM golang:alpine as builder

# Describe container dependencies
LABEL api.LINKS.redis=""\
      api.LINKS.redis.image="redis:latest"\
      api.LINKS.redis.port="6379"\
      api.LINKS.redis.description="For caching results from OWM API."\
      api.LINKS.redis.mandatory="true"


# Set and describe available ENVs
LABEL api.ENV.OPENWEATHERMAP_APIKEY="" \
      api.ENV.OPENWEATHERMAP_APIKEY.description="Access key for OpenWeatherMap. See http://openweathermap.org/appid for details." \
      api.ENV.OPENWEATHERMAP_APIKEY.mandatory="true"


LABEL api.EXPOSE.1337="" \
      api.EXPOSE.1337.protocol="http" \
      api.EXPOSE.1337.description="The main endpoint of this service."


ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

