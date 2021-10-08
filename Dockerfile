FROM golang
WORKDIR /app/src/go_clean_api
ENV GOPATH=/app
COPY . /app/src/go_clean_api
RUN go get -u github.com/asaskevich/govalidator
RUN go get -u github.com/dgrijalva/jwt-go
RUN go get -u github.com/gin-gonic/gin 
RUN go get -u github.com/go-redis/redis/v8
RUN go get -u github.com/joho/godotenv 
RUN go get -u github.com/satori/go.uuid 
RUN go get -u golang.org/x/crypto 
RUN go get -u gorm.io/driver/mysql
RUN go get -u gorm.io/gorm
RUN go build -o main .
CMD [ "./main" ]