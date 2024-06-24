GOOSE_DRIVER=postgres
GOOSE_DBSTRING="user=postgres password=postgres dbname=postgres sslmode=disable"
DIR=migrations
down:
	goose -dir=$(DIR) $(GOOSE_DRIVER)  $(GOOSE_DBSTRING) down

up:
	goose -dir=$(DIR) $(GOOSE_DRIVER)  $(GOOSE_DBSTRING) up