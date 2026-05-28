package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Port string
	// go get -u github.com/go-sql-driver/mysql
	Cfg       mysql.Config
	SecretKey []byte //para hash de senha

)

func LoadEnv() {
	// Caminho ABSOLUTO do arquivo config.go (onde esta função está)
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Não foi possível obter o caminho do arquivo de config")
	}

	// Pasta /config
	configDir := filepath.Dir(filePath)

	// .env dentro de /config
	envPath := filepath.Join(configDir, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Erro ao carregar .env em %s: %s", envPath, err)
	}

	log.Println("Sucesso em carregar o .env de:", envPath)

	Port = os.Getenv("API_PORT")
	Cfg = mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_ADDR"),
		DBName: os.Getenv("DB_DATABASE"),
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
