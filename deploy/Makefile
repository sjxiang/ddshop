SHELL := /bin/bash


NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m




open:
	@echo ''
	@printf '$(OK_COLOR)打开容器服务 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml up -d 
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''

close:
	@echo ''
	@printf '$(OK_COLOR)关闭容器服务 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml down 
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''




login:
	@echo ''
	@printf '$(OK_COLOR)登录 MySQL 容器 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml exec ddshopdb sh -c 'mysql -uroot -p${MYSQL_ROOT_PASSWORD}'
	@printf '$(OK_COLOR)退出 .. 🎯$(NO_COLOR)\n'
	@echo ''



container_net:
	@echo ''
	@printf '$(OK_COLOR)查看 MySQL 容器 IP 地址 .. 🚀$(NO_COLOR)\n'
	@docker inspect ddshopdb | grep IPAddress
	@echo ''

login_redis:
	@echo ''
	@printf '$(OK_COLOR)登录 Redis 容器 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml exec cache sh -c 'redis-cli'
	@printf '$(OK_COLOR)退出 .. 🎯$(NO_COLOR)\n'
	@ech