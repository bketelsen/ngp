FROM golang
RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/github.com/bketelsen/ngp
WORKDIR /go/src/github.com/bketelsen/ngp
RUN make clean
RUN make all
CMD /bin/bash
