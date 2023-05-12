SHELL := /bin/bash


NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m


up:
	@printf "\n"
	@printf "$(OK_COLOR)启动 ELK 容器服务 .. 🚀$(NO_COLOR)\n"
	sudo docker-compose up -d
	@echo "🎯"
	@printf "\n"


down:
	@printf "\n"
	@printf "$(OK_COLOR)停止 ELK 容器服务 .. 🚀$(NO_COLOR)\n"
	sudo docker-compose down 
	@echo "🎯"
	@printf "\n"



net:
	@echo ''
	@printf '$(OK_COLOR)查看 ElasticSearch 容器 IP 地址 .. 🚀$(NO_COLOR)\n'
	@docker inspect es-example | grep IPAddress
	@echo ''
