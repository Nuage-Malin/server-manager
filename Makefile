##
## backend
## File description:
## Makefile
##

GO	=	go

NAME	=	hsm

SRCDIR	=	cli

SRC		=	cli.go \
			wake.go \
			run.go

SRC			:= $(addprefix $(SRCDIR)/, $(SRC))

GOFLAGS =	--trimpath --mod=vendor

all: build

build:
	$(GO) mod vendor
	$(GO) build $(GOFLAGS) -o ./$(NAME) $(SRC)

fclean:
	rm -f  $(NAME)