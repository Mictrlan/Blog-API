FROM scratch

WORKDIR /Users/Miuer/go/src/github.com/Mictrlan/blog-api
COPY . /Users/Miuer/go/src/github.com/Mictrlan/blog-api

EXPOSE 8080
CMD [ "./blog-api" ]
