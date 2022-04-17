package main

import (
	"database/sql"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"github.com/guozhe001/brwswg/chapter4/dbutils"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

const (
	ContentType = "Content-Type"
	TextPlain   = "text/plain"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("{train-id}").To(t.removeTrain))
	container.Add(ws)
}

func (t *TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("delete from train where id = ?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader(ContentType, TextPlain)
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (t *TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	if err != nil {
		log.Println(err)
	}
	log.Println(b.DriverName, b.OperatingStatus)
	statement, _ := DB.Prepare("insert into train(DRIVER_NAME, OPERATING_STATUS) values(?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err != nil {
		response.AddHeader(ContentType, TextPlain)
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		id, _ := result.LastInsertId()
		b.ID = int(id)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	}
}

func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train where id = ?", id).
		Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader(ContentType, TextPlain)
		response.WriteErrorString(http.StatusNotFound, "Train could not found")
	} else {
		response.WriteEntity(t)
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "/Users/guozhe/code/my-project/brwswg/chapter4/railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Println("start listening on localhost:8088")
	server := &http.Server{Addr: ":8088", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
