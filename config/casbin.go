package config

import (
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InitCasbin(mongoClient *mongo.Client) *casbin.Enforcer {
	// اتصال به MongoDB Adapter
	adapter, err := mongodbadapter.NewAdapterByDB(mongoClient, &mongodbadapter.AdapterConfig{}) // Your MongoDB URL.
	if err != nil {
		panic(err)
	}
	// ایجاد Enforcer با مدل و Adapter
	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		log.Fatalf("Casbin Enforcer Init Error: %v", err)
	}
	// بارگذاری policy‌ها
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Failed to load Casbin policy: %v", err)
	}
	return enforcer
}
