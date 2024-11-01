NAME = bot

all:	$(NAME)

$(NAME):
	docker compose -v -f ./srcs/docker-compose.yml up --force-recreate --build -d

down:
	docker compose -f ./srcs/docker-compose.yml down

init:
	./configure

fclean:
	@-docker rmi go
	@-docker system prune -a --force

fvolume: 
	@-docker volume prune
	@-rm -rf ./mariadb/*

re: down fclean all
