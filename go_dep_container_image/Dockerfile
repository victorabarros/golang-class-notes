# docker build --build-arg GITHUB_TOKEN=xtoken . -t test:13
# docker run -v $(pwd):/go/src/github.com/victorabarros/golang-class-notes/ -it golang:1.14-alpine sh
FROM golang:1.14-alpine
    ARG GITHUB_TOKEN
    RUN cd ~/..

    RUN apk update && apk upgrade && apk add --update --no-cache curl git
    RUN rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

    RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | DEP_RELEASE_TAG=v0.5.0 sh

    RUN cd ${APP_DIR}

    RUN git config --global --add url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
    RUN dep init
    RUN dep ensure -v -vendor-only
