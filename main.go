package main
import (
	"github.com/osataken/go-mgo/db"
	"github.com/osataken/go-mgo/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main() {
	db.InitMongoDB()

	db.Villager.Insert(&models.Villager{Name:"Ken", Surname:"Sama"})

	var villagers []models.Villager
	db.Villager.Find(bson.M{"name": "Ken"}, 0, -1, &villagers)
	fmt.Println(villagers)

	db.SecretInfo.Insert(&models.SecretInfo{Regular: "regular data", Encrypted:"encrypted data"})

	var secretInfos []models.SecretInfo
	db.SecretInfo.Find(bson.M{"regular": "regular data"}, 0, -1, &secretInfos)
	fmt.Println(secretInfos)
}
