FROM scratch

WORKDIR /Users/Miuer/go/src/github.com/Mictrlan/blog-api
COPY . /Users/Miuer/go/src/github.com/Mictrlan/blog-api

EXPOSE 8083
CMD [ "./blog-api" ]
