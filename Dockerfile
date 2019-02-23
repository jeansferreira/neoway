FROM golang:cross

# The golang Docker sets the $GOPATH to be /go
# https://github.com/docker-library/golang/blob/c1baf037d71331eb0b8d4c70cff4c29cf124c5e0/1.4/Dockerfile
#RUN mkdir -p /go/src/github.com/jeansferreira/Projeto_Neoway
#WORKDIR /go/src/github.com/jeansferreira/Projeto_Neoway

#CMD ["go","get","github.com/jmoiron/sqlx"]
#CMD ["go","get","github.com/lib/pq"]
#CMD ["go","get","github.com/acroca"]
#CMD ["go","get","github.com/go-delve"]
#CMD ["go","get","github.com/jeansferreira"]
#CMD ["go","get","github.com/jmiron"]
#CMD ["go","get","github.com/jmoiron"]
#CMD ["go","get","github.com/karrick"]
#CMD ["go","get","github.com/lib"]
#CMD ["go","get","github.com/nsf"]
#CMD ["go","get","github.com/pkg"]
#CMD ["go","get","github.com/ramya-rao-a"]
#CMD ["go","get","github.com/rogpeppe"]
#CMD ["go","get","github.com/sqs"]
#CMD ["go","get","github.com/uudashr"]

#RUN ["apt-get", "update"]

#RUN go get -u github.com/lib/pq

#RUN go get github.com/jmoiron/sqlx

#RUN go run main.go

#RUN sudo chmod +x /home/oem/work/src/github.com/jeansferreira/neoway/docker/install.sh
#CMD ["/home/oem/work/src/github.com/jeansferreira/neoway/docker/install.sh"]

CMD ["go","run","main.go"]
