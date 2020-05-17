# alpine image
FROM alpine:3.9

# add mongodb to alpine img
RUN apk add --no-cache mongodb mongodb-tools
