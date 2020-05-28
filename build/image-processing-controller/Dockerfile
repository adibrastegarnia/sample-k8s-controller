FROM alpine:3.9

RUN apk upgrade --update --no-cache

USER nobody

ADD build/_output/image-processing-controller /usr/local/bin/image-processing-controller

ENTRYPOINT ["image-processing-controller"]
