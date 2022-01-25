FROM golang as builder

WORKDIR /go/src/project2
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -v

FROM alpine
LABEL maintainer="jayaraj.esvar@gmail.com"
WORKDIR /home
COPY --from=builder /go/src/project2/project2 /home
COPY --from=builder /go/src/project2/.project2.yml /home/.project2.yml
CMD [ "/home/project2" ]
