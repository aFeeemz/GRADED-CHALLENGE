# Makefile for installing Go dependencies

# Variables
GOCMD = go
GOGET = $(GOCMD) get -u

# Targets
all: deps

deps:
	$(GOGET) golang.org/x/crypto/bcrypt
	$(GOGET) github.com/joho/godotenv
	$(GOGET) github.com/golang-jwt/jwt/v4
	$(GOGET) gorm.io/gorm
	$(GOGET) github.com/labstack/echo/v4
	$(GOGET) github.com/labstack/echo/v4/middleware
	$(GOGET) github.com/sirupsen/logrus
	$(GOGET) go.mongodb.org/mongo-driver/mongo
	$(GOGET) go.mongodb.org/mongo-driver/mongo/options

clean:
	$(GOCMD) clean

.PHONY: all deps clean
