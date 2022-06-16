package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type StreamST struct {
	URL          string `json:"url"`
	Status       bool   `json:"status"`
	OnDemand     bool   `json:"on_demand"`
	DisableAudio bool   `json:"disable_audio"`
	Debug        bool   `json:"debug"`
	RunLock      bool   `json:"-"`
}

func main() {
	fmt.Println("hello")

	database, err := sql.Open("sqlite3", "./test.db")
	checkErr(err)
	// Create Table
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS links (Id INTEGER PRIMARY KEY,RobotId VARCHAR(255) NOT NULL,CameraName VARCHAR(255) NOT NULL,WebrtcLink VARCHAR(255) NOT NULL,RtspLink VARCHAR(255) NOT NULL)")
	checkErr(err)
	statement.Exec()

	statement, err = database.Prepare("INSERT INTO LINKS(robotid,cameraname,webrtclink,rtsplink) VALUES (?,?,?,?)")
	checkErr(err)

	statement.Exec("40b183ed-a2d6-4d9b-8f59-8e41eea5f110", "camera-1", "40b183ed-a2d6-4d9b-8f59-8e41eea5f110_camera-1", "rtsp://127.0.0.1:8554_camera-1")
	checkErr(err)

	statement.Exec("40b183ed-a2d6-4d9b-8f59-8e41eea5f110", "camera-2", "40b183ed-a2d6-4d9b-8f59-8e41eea5f110_camera-2", "rtsp://127.0.0.1:8554_camera-2")
	checkErr(err)

	statement.Exec("40b183ed-a2d6-4d9b-8f59-8e41eea5f110", "camera-3", "40b183ed-a2d6-4d9b-8f59-8e41eea5f110_camera-3", "rtsp://127.0.0.1:8554_camera-3")
	checkErr(err)

	rows, err := database.Query("SELECT * from links")
	checkErr(err)

	var id int
	var RobotId string
	var CameraName string
	var WebrtcLink string
	var RtspLink string
	a := make(map[string]StreamST)

	for rows.Next() {
		err = rows.Scan(&id, &RobotId, &CameraName, &WebrtcLink, &RtspLink)
		checkErr(err)
		fmt.Printf("%d: %s %s %s %s\n", id, RobotId, CameraName, WebrtcLink, RtspLink)

		a[WebrtcLink] = StreamST{URL: RtspLink}
	}
	rows.Close()
	database.Close()
	log.Printf("%+v", a)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
