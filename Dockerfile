# Script for reading wtmp files
#
# To build:
# $ docker run --rm -v $(pwd):/go/src/github.com/micahhausler/utmptail -w /go/src/github.com/micahhausler/utmptail golang:1.5  go build -v -a -tags netgo -installsuffix netgo -ldflags '-w'
# $ docker build -t micahhausler/utmptail .
#
# To run:
# $ docker run -v /var/log/wtmp:/var/log/wtmp  micahhausler/utmptail /var/log/wtmp
# or (for 64bit systems only)
# $ tail -F --bytes 384 /var/log/wtmp | docker run -i micahhausler/utmptail

FROM busybox

MAINTAINER Micah Hausler, <micah.hausler@ambition.com>

COPY utmptail /bin/utmptail
RUN chmod 755 /bin/utmptail

ENTRYPOINT ["/bin/utmptail"]
