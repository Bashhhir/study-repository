package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	// создаем подключение к базе данных
	conn, err := pgx.Connect(ctx, "postgres://postgres:bashir@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	// проверяем это подключение
	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключение к базе данных произошло успешно!")
}
