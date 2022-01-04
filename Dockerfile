FROM golang:1.16-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR ${ROOT}


RUN apk update && apk add git 
RUN go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    github.com/ramya-rao-a/go-outline@latest \
    github.com/nsf/gocode@latest \
    github.com/acroca/go-symbols@latest \
    github.com/fatih/gomodifytags@latest \
    github.com/josharian/impl@latest \
    github.com/haya14busa/goplay/cmd/goplay@latest \
    github.com/go-delve/delve/cmd/dlv@latest \
    golang.org/x/lint/golint@latest \
    golang.org/x/tools/gopls@latest
COPY go.mod ./
COPY go.sum ./
RUN go mod download
EXPOSE 8080

CMD ["go", "run", "./cmd/main.go"]

FROM golang:1.16-alpine as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/binary


FROM scrach as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD [ "/go/src/app/binary" ]