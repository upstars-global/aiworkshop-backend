#Створити команду для генерації документації на основі коментарів у файлах програми за допоможу swag init
swagger-gen:
	swag init -g cmd/app/main.go -o docs