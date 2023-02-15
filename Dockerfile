FROM golang:latest
EXPOSE 9990
ENV MUDPORT 9990
COPY . /app
WORKDIR /app/src/RestGo.MUD
#刪除Firebase在Local開發時使用的Emulator設定檔
RUN rm -f /app/FirebaseEmulator.setting || true
RUN go build -o app ./main.go
ENTRYPOINT  ["./app"]