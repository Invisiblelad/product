package config
import "os"
type Config struct{
	Port   string
	MongoURI   string
	Database   string
	Collection  string
}

func Load()*Config{
	return &Config{
		Port:       getEnv("PORT", "8080"),
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
        Database:   getEnv("DATABASE", "product"),
		Collection: getEnv("COLLECTION", "items"),
	}
}

func getEnv(key, defaultVal string)string{
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}