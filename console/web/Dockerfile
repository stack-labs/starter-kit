FROM alpine:3.10
ADD vue/dist /vue/dist
ADD bin/linux_amd64/console-web /console-web
WORKDIR /
ENTRYPOINT [ "/console-web" ]
