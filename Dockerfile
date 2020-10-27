FROM ubuntu
ENV TZ=Europe/Madrid
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt-get update
RUN apt-get -y install wget
RUN apt-get -y install git
RUN apt-get -y install golang
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/gorilla/mux"

RUN mkdir /service
ADD . /service



EXPOSE 8081

CMD ["./service/poster"]
