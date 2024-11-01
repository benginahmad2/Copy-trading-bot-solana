
NAME = bot
MOD = go.mod
SUM = go.sum


all: $(NAME)

$(NAME): install build
	./${NAME}

install:
	@-go mod init bot
	@-go mod tidy

build: 
	@-go build .

fclean:
	rm ./${MOD}
	rm ./${SUM}
	rm ./${NAME}

re: fclean all
