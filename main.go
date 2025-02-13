package main

func main() {
	// 1. initialize DB
	db := ConfigureDB(&Configuration{
		DBHost:     "db",
		DBUser:     "postgres",
		DBPassword: "123456",
		DBName:     "postgres",
	})
	// 2. initialize API
	engine := ConfigureRoutes(db)
	// 3. run the API
	engine.Run(":8080")
}
