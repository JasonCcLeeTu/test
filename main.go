package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type testInfo struct {
	StTime string `bson:"stTime"`
	EdTime string `bson:"edTime"`
}

type ReplyEventInfo struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Place  string `json:"place"`
	Guests string `json:"guests"`
}

type EventInfo struct {
	Name    string `bson:"name"`
	Type    string `bson:"type"`
	Date    string `bson:"date"`
	StTime  string `bson:"stTime"`
	EdTime  string `bson:"edTime"`
	Place   string `bson:"place"`
	Guests  string `bson:"guests"`
	Content string `bson:"content"`
	Image   string `bson:"image"`
}

type MouseInfo struct {
	Brand    string  `bson:"brand"`
	Name     string  `bson:"name"`
	HasLOL   bool    `bson:"hasLOL"`
	StDpi    int     `bson:"stDpi"`
	EdDpi    int     `bson:"edDpi"`
	Wireless bool    `bson:"wireless"`
	Rgb      bool    `bson:"rgb"`
	Feature  string  `bson:"feature"`
	Price    float64 `bson:"price"`
	Image    string  `bson:"image"`
}

type KeyBoardInfo struct {
	Brand    string  `bson:"brand"`
	Name     string  `bson:"name"`
	HasLOL   bool    `bson:"hasLOL"`
	Switch   string  `bson:"switch"`
	Wireless bool    `bson:"wireless"`
	Numpad   bool    `bson:"numpad"`
	Rgb      bool    `bson:"rgb"`
	Feature  string  `bson:"feature"`
	Price    float64 `bson:"price"`
	Image    string  `bson:"image"`
}

type HeadPhoneInfo struct {
	Brand       string  `bson:"brand"`
	Name        string  `bson:"name"`
	HasLOL      bool    `bson:"hasLOL"`
	StFrequency int     `bson:"stFrequency"`
	EdFrequency int     `bson:"edFrequency"`
	Impedance   int     `bson:"impedance"`
	Type        string  `bson:"type"`
	Wireless    bool    `bson:"wireless"`
	Rgb         bool    `bson:"rgb"`
	Feature     string  `bson:"feature"`
	Price       float64 `bson:"price"`
	Image       string  `bson:"image"`
}

func main() {
	uri := "mongodb://app:11111111@127.0.0.1:29124/"
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {

		fmt.Printf("here is err:%+v\n", err)
		log.Panicln(err)
	}

	//testfind(client)
	// CreateCollection(client)
	insert(client)
	//find(client)

}

func testfind(client *mongo.Client) {

	clientCol := client.Database("kioskLOL").Collection("test")
	condition := bson.M{
		"stTime": bson.M{"$lte": "10:31"}, //大於等於 查詢時間的資料
	}

	cur, err := clientCol.Find(context.Background(), condition)
	if err != nil {
		log.Panicln("FIND ERROR:", err)
	}

	var respAll []testInfo
	for cur.Next(context.TODO()) {

		var data testInfo
		err = cur.Decode(&data)
		if err != nil {
			log.Panicln("cur error: ", err)
		}

		respAll = append(respAll, data)

	}

	for _, val := range respAll {
		fmt.Printf("%+v", val)
	}

}

func find(client *mongo.Client) {
	regexPattern := fmt.Sprintf(".*%s.*", "妲妲")
	condition := bson.M{
		"guests": bson.M{"$regex": primitive.Regex{Pattern: regexPattern, Options: "i"}},
	}
	regexPattern = fmt.Sprintf(".*%s.*", "實體活動")
	condition["type"] = bson.M{"$regex": primitive.Regex{Pattern: regexPattern, Options: "i"}}

	// dateString := strings.Replace(strings.TrimSpace("2023/10/08"), "/", "-", -1)
	// date, err := time.Parse("2006-01-02", dateString)

	// if err != nil {
	// 	log.Panicln("date parse")
	// }
	condition["date"] = bson.M{"$eq": "2023/10/08"}

	//var strTime string
	trimTime := strings.TrimSpace("19:16")
	// splitTime := strings.Split(trimTime, ":")
	// if len(splitTime) == 2 {
	// 	strTime = fmt.Sprintf("2000-01-01 %v:00", trimTime) //由於mongodb不能單獨存時間,但又有需求要純查時間所以固定日期配合需要存的時間,做儲存 所以這邊查詢的時候固定2000-01-01,時間格式要 00:00:00所以給的資料沒有要補齊
	// } else {
	// 	strTime = fmt.Sprintf("2000-01-01 %v", trimTime)
	// }
	// log.Println("轉出來的時間:", strTime)

	// queryTime, err := time.Parse("2006-01-02 15:04:05", strTime)
	// if err != nil {
	// 	log.Panicf("eventData Parse time error: %v", err)
	// }

	condition["stTime"] = bson.M{"$lte": trimTime} //大於等於 查詢時間的資料
	condition["edTime"] = bson.M{"$gte": trimTime} //小於等於 查詢時間的資料

	clientCol := client.Database("kioskLOL").Collection("kioskLOL_event")
	cur, err := clientCol.Find(context.Background(), condition)
	if err != nil {
		log.Panicln("FIND ERROR:", err)
	}

	var respAll []EventInfo
	for cur.Next(context.TODO()) {

		var data EventInfo
		err = cur.Decode(&data)
		if err != nil {
			log.Panicln("cur error: ", err)
		}

		respAll = append(respAll, data)

	}

	for _, val := range respAll {
		fmt.Printf("%+v", val)
	}

}

func CreateCollection(client *mongo.Client) {
	clientDB := client.Database("kioskLOL")
	err := clientDB.CreateCollection(context.TODO(), "test")
	if err != nil {
		log.Panicln(err)
	}
}

func insert(client *mongo.Client) {

	// date, err := time.Parse("2006-01-02", "2023-10-08")
	// if err != nil {
	// 	log.Panic("DATE PARSE ERROR:%v", err)
	// }

	// ts := "2000-01-01 15:30:00"
	// st, err := time.Parse("2006-01-02 15:04:05", ts)
	// if err != nil {
	// 	log.Panicln("bbbbbbbbbb:", err)
	// }
	// ts = "2000-01-01 19:15:00"
	// edt, err := time.Parse("2006-01-02 15:04:05", ts)
	// if err != nil {
	// 	log.Panicln("aaaaaaaaa:", err)
	// }

	clientCol := client.Database("kioskLOL").Collection("kioskLOL_event")
	document := EventInfo{
		Name:    "英雄召集令",
		Type:    "線下實體活動",
		Date:    "2023/10/08",
		StTime:  "15:30",
		EdTime:  "19:15",
		Place:   "花博爭艷館",
		Guests:  "依渟、妲妲",
		Content: "加入依渟、妲妲兩位實況主的麾下！現場將組隊進行「AOC 英雄召集令」爭霸賽團隊積分戰，各關卡獲勝的隊伍可獲得不同積分，所有關卡結束時依積分累積高低。獲勝的團隊可以獲得優質好禮！",
	}
	//------------------------------------

	// clientCol := client.Database("kioskLOL").Collection("kioskLOL_mouse")
	// document := MouseInfo{
	// 	Brand:    "LOGI 羅技",
	// 	Name:     "League of Legends PRO Wireless",
	// 	HasLOL:   true,
	// 	StDpi:    100,
	// 	EdDpi:    25600,
	// 	Wireless: true,
	// 	Rgb:      false,
	// 	Feature:  "這款滑鼠特點包括LIGHTSPEED無線技術、HERO 25K感應器，以及1680萬種RGB燈效。長效電池，48小時背光使用，60小時無背光。有4-8個可自訂按鍵，重量輕巧，非常適合電競玩家使用。",
	// 	Price:    3290,
	// }

	//----------------------
	// clientCol := client.Database("kioskLOL").Collection("kioskLOL_keyboard")
	// document := KeyBoardInfo{
	// 	Brand:    "LOGI 羅技",
	// 	Name:     "League of Legends PRO",
	// 	HasLOL:   true,
	// 	Switch:   "GX茶軸",
	// 	Wireless: false,
	// 	Numpad:   false,
	// 	Rgb:      true,
	// 	Feature:  "限量Logitech G x League of Legends PRO 機械式有線遊戲鍵盤，所有鍵盤皆有注音標示，無單純英文鍵盤。此鍵盤無法提換鍵軸，針對職業選手而設計，具精簡辨析的無數字鍵台和可拆卸式連接線無數字鍵台的便攜設計，為桌面騰出更多空間使用滑鼠並方便攜帶LIGHTSYNC RGB能使用G HUB自訂背光以供競賽系統使用。",
	// 	Price:    3990,
	// }

	// clientCol := client.Database("kioskLOL").Collection("kioskLOL_headPhone")
	// document := HeadPhoneInfo{
	// 	Brand:       "LOGI 羅技",
	// 	Name:        "League of Legends PRO X",
	// 	HasLOL:      true,
	// 	StFrequency: 20,
	// 	EdFrequency: 20000,
	// 	Impedance:   35,
	// 	Type:        "耳罩式",
	// 	Wireless:    false,
	// 	Rgb:         false,
	// 	Feature: `● DTS HEADPHONE:X 2.0環繞音效
	// 	● 先進的50公釐單體
	// 	● BLUE VOICE先進麥克風技術
	// 	● 舒適記憶泡綿耳墊與製作工藝
	// 	● 優質USB外接音效卡`,
	// 	Price: 3490,
	// }
	// clientCol := client.Database("kioskLOL").Collection("test")
	// document := testInfo{
	// 	StTime: "10:30",
	// 	EdTime: "12:00",
	// }
	_, err := clientCol.InsertOne(context.TODO(), document)

	if err != nil {
		log.Panicln(err)
	}
}
