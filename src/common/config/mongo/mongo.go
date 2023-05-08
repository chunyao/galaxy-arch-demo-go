package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	MGOs = map[string]*MGOManager{} // 初始化时加载数据源到集合
)

type MGOManager struct {
	MG *mongo.Database // redis
}

func InitMongoDB() (err error) {
	log.Info("Initializing Mongo")
	for k, _ := range viper.GetStringMap("mongo") {
		log.Printf("初始化Mongo数据源 %s ", k)
		db, _ := MGoSetup(k)

		rdb := &MGOManager{
			MG: db.Database(viper.GetString("mongo." + k + ".database")),
		}
		MGOs[k] = rdb
	}
	log.Info(MGOs)
	log.Info("Mongo: initialization completed")
	return nil
}

func MGoSetup(name string) (*mongo.Client, error) {

	var uri = viper.GetString("mongo." + name + ".uri")
	var auth options.Credential
	auth.Username = viper.GetString("mongo." + name + ".user")
	auth.Password = viper.GetString("mongo." + name + ".password")
	auth.AuthSource = viper.GetString("mongo." + name + ".database")
	mgo, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + uri).
		SetAuth(auth).
		SetReplicaSet(viper.GetString("mongo." + name + ".replicaset")).
		SetMinPoolSize(uint64(viper.GetInt("mongo." + name + ".MinPoolSize"))).
		SetMaxPoolSize(uint64(viper.GetInt("mongo." + name + ".MaxPoolSize"))).
		SetMaxConnIdleTime(time.Duration(viper.GetInt("mongo." + name + ".MaxConnIdleTime"))).SetTimeout(10 * time.Second))
	err = mgo.Connect(context.TODO())
	if err != nil {
		log.Fatalf("[GetMongoClient.Connect] Error while connect to mongodb %v", err)
	}

	if err := mgo.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
	return mgo, err
}
