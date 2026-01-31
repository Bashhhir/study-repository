package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	// передаем созданное подключение для создания таблицы
	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// передаем созданное подключение для внесения строк (rows) в таблицу
	// if err := simple_sql.InsertRow(ctx,
	// 	conn,
	// 	"Обед",
	// 	"Покушац надо",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simple_sql.DeleteRow(ctx, conn); err != nil {
	// 	panic(err)
	// }

	fmt.Println("succeed!")
}
