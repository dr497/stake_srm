FROM golang:latest

WORKDIR /usr/src/stake_srm_go

COPY . .

RUN go build -o stake_srm .

CMD [ "./stake_srm" ]
