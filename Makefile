.PHONY: mysql-create mysql-start mysql-stop mysql-restart mysql-logs mysql-connect mysql-clean

MYSQL_CONTAINER_NAME = go12-service-mysql
MYSQL_ROOT_PASSWORD = secret
MYSQL_DATABASE = food_delivery
MYSQL_USER = devuser
MYSQL_PASSWORD = devpassword

mysql-create:
	@echo "Creating MySQL container..."
	docker run -d \
		--name $(MYSQL_CONTAINER_NAME) \
		-v mysql-data:/var/lib/mysql \
		-e MYSQL_ROOT_PASSWORD=$(MYSQL_ROOT_PASSWORD) \
		-e MYSQL_DATABASE=$(MYSQL_DATABASE) \
		-e MYSQL_USER=$(MYSQL_USER) \
		-e MYSQL_PASSWORD=$(MYSQL_PASSWORD) \
		-p 3306:3306 \
		mysql:8.4.5-oraclelinux9
	@echo "MySQL container created and started on port 3306"

mysql-start:
	@echo "Starting existing MySQL container..."
	docker start $(MYSQL_CONTAINER_NAME)
	@echo "MySQL container started"

mysql-stop:
	@echo "Stopping MySQL container..."
	docker stop $(MYSQL_CONTAINER_NAME)
	@echo "MySQL container stopped"

mysql-restart:
	@echo "Restarting MySQL container..."
	docker restart $(MYSQL_CONTAINER_NAME)
	@echo "MySQL container restarted"

mysql-logs:
	docker logs -f $(MYSQL_CONTAINER_NAME)

mysql-connect:
	docker exec -it $(MYSQL_CONTAINER_NAME) mysql -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) $(MYSQL_DATABASE)

mysql-clean:
	@echo "Cleaning up MySQL container and volume..."
	docker stop $(MYSQL_CONTAINER_NAME) || true
	docker rm $(MYSQL_CONTAINER_NAME) || true
	docker volume rm mysql-data || true
	@echo "MySQL cleanup completed"
