FROM ubuntu:16.04
MAINTAINER matt.ho at gmail.com

ENV PORT 80
EXPOSE 80

ADD echo /opt/echo/bin/echo

CMD [ "/opt/echo/bin/echo" ]
