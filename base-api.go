package base_api

import (
	"github.com/gin-gonic/gin"
	"github.com/maelsilva96/base-api/utils"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/secrets"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var ociConfigProvider *common.ConfigurationProvider
var secretsClient secrets.SecretsClient

func loadConfiguration() {
	var err error
	pwdPk := os.Getenv("PWD_PK")
	ociConfigProvider = utils.GetConfigProvider(pwdPk)
	secretsClient, err = secrets.NewSecretsClientWithConfigurationProvider(*ociConfigProvider)
	utils.ValidErrorPanic(err)
}

func loadLogging() {
	if !utils.IsDevelopment() {
		an := os.Getenv("APP_NAME")
		eli := os.Getenv("LOGGER_ENTRY_ID")
		lw := utils.NewOracleLogWriter(an, eli, ociConfigProvider)
		log.SetOutput(lw)
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.SetOutput(&utils.LogConsole{})
	}
}

func LoadConfiguration() (*gorm.DB, *gin.Engine) {
	var err error
	var conn string
	utils.LoadEnvFile()
	if utils.IsDevelopment() {
		conn = os.Getenv("STR_CONN")
	} else {
		loadConfiguration()
		strConn := os.Getenv("STR_CONN")
		conn, err = utils.GetSecreteContent(secretsClient, strConn, secrets.GetSecretBundleStageCurrent)
	}
	loadLogging()
	log.Println("Start DB!")
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, gin.Default()
}
