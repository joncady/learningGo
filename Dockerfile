FROM alpine
RUN apk add --no-cache ca-certificates
COPY learningGo /learningGo
EXPOSE 80
ENTRYPOINT ["/learningGo"]