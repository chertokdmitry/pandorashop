FROM golang:1.14

WORKDIR /app/src



#ENV REPO_URL=gitlab.com/chertokdmitry/bot-pandora-11
#ENV GOPATH=/app
#ENV APP_PATH=$GOPATH/src/$REPO_URL
#ENV WORKPATH=$APP_PATH/src
#WORKDIR $WORKPATH
#COPY src $WORKPATH
COPY src .
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go mod download
#COPY . .

ENV PORT 8000
RUN go build

CMD ["./bot-pandora-11"]



#ENV REPO_URL=gitlab.com/chertokdmitry/bot-pandora-11
#ENV GOPATH=/app
#ENV APP_PATH=$GOPATH/src/$REPO_URL
#ENV WORKPATH=$APP_PATH/src
#COPY src $WORKPATH
#WORKDIR $WORKPATH


