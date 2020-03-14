FROM alpine:3.10

ADD conf /conf
ADD bin/linux_amd64/account-srv /account-srv

WORKDIR /
ENTRYPOINT [ "/account-srv" ]
