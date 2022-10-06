##
## backend
## File description:
## Makefile
##

GO	=	go

NAME	=	hsm.so

SRCDIR	=	.

SRC		=	hsmlib.go 	\

SRC			:= $(addprefix $(SRCDIR)/, $(SRC))

GOFLAGS =	--trimpath --mod=vendor --buildmode=plugin

all: lib

lib:
	$(GO) mod vendor
	$(GO) build $(GOFLAGS) -o ./$(NAME) $(SRC)

fclean:
	rm -f  $(NAME)
