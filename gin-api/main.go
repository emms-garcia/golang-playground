package main

func main() {
	db := ConfigureDB(&Configuration{
		DBHost:     "db",
		DBUser:     "postgres",
		DBPassword: "123456",
		DBName:     "postgres",
	})
	app := NewApp(db)
	engine := ConfigureRoutes(app)
	engine.Run(":8080")
}
