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

all: build lib

build:
	$(GO) mod vendor
	$(GO) build $(GOFLAGS) -o ./$(NAME) $(SRC)

lib:
	$(MAKE) -C ./hsmlib

fclean:
	rm -f  $(NAME)

re: fclean build lib