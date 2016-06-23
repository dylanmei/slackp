FROM alpine:3.4
MAINTAINER Dylan Meissner "https://github.com/dylanmei/slackp"

RUN apk add --no-cache --update ca-certificates

ADD bin/slackp /bin/slackp
#ENTRYPOINT ["/bin/slackp"]
CMD ["slackp"]
