package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"strconv"
	"encoding/json"
	"time"
)

// 电影院
type Cinema struct {
	Name string `json:"name"`
	Company string `json:"company"`
}

// 影厅
type VideoHall struct {
	Name string `json:"name"`
	Cinema string `json:"cinema"`
	Width uint64 `json:"width"`
	Height uint64 `json:"height"`
}

// 票务平台
type TicketPlatform struct {
	Name string `json:"name"`
}

// 排片
type MoviePlan struct {
	ID string `json:"id"`
	Movie string `json:"movie"`
	Cinema string `json:"cinema"`
	VideoHall string `json:"video_hall"`
	PlanTime string `json:"plan_time"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
}

// 电影票
type Ticket struct {
	ID string `json:"id"`
	Movie string `json:"movie"`
	MoviePlan string `json:"movie_plan"`
	X uint64 `json:"x"`
	Y uint64 `json:"y"`
	IsLocked bool `json:"is_locked"`
	LockPrice uint32 `json:"lock_price"`
	IsChecked bool `json:"is_checked"`
	IsClear bool `json:"is_clear"`
}

// 分账
type Clear struct {
	IssueNum uint64 `json:"issue_num"`
	LockNum uint64 `json:"lock_num"`
	CheckNum uint64 `json:"check_num"`
	BoxOffice uint64 `json:"box_office"`
}

// 分账结果
type ClearResult struct {
	Clear
	RegulationProfit float32 `json:"regulation_profit"`
	CinemaProfit float32 `json:"cinema_profit"`
	CinemaCompanyProfit float32 `json:"cinema_company_profit"`
	IssuerProfit float32 `json:"issuer_profit"`
}

type BoolResult struct {
	Success bool `json:"success"`
}

type Contract struct {
}

func main() {
	err := shim.Start(new(Contract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (c *Contract) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("init is running ")
	fmt.Println("create table cinema")
	err := stub.CreateTable("cinema", []*shim.ColumnDefinition{
		{
			Name: "name",
			Type: shim.ColumnDefinition_STRING,
			Key: true,
		},
		{
			Name: "company",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("create table ticket_platform")
	err = stub.CreateTable("ticket_platform", []*shim.ColumnDefinition{
		{
			Name: "name",
			Type: shim.ColumnDefinition_STRING,
			Key: true,
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("create table video_hall")
	err = stub.CreateTable("video_hall", []*shim.ColumnDefinition{
		{
			Name: "name",
			Type: shim.ColumnDefinition_STRING,
			Key: true,
		},
		{
			Name: "cinema",
			Type: shim.ColumnDefinition_STRING,
			Key: true,
		},
		{
			Name: "width",
			Type: shim.ColumnDefinition_UINT64,
			Key: false,
		},
		{
			Name: "height",
			Type: shim.ColumnDefinition_UINT64,
			Key: false,
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("create table movie_plan")
	err = stub.CreateTable("movie_plan", []*shim.ColumnDefinition{
		{
			Name: "id",
			Type: shim.ColumnDefinition_STRING,
			Key: true,
		},
		{
			Name: "movie",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "cinema",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "video_hall",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "plan_time",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "start_time",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "end_time",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("create table ticket")
	err = stub.CreateTable("ticket", []*shim.ColumnDefinition{
		{
			Name: "id",
			Type: shim.ColumnDefinition_STRING,
			Key: true,
		},
		{
			Name: "movie",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "movie_plan",
			Type: shim.ColumnDefinition_STRING,
			Key: false,
		},
		{
			Name: "x",
			Type: shim.ColumnDefinition_UINT64,
			Key: false,
		},
		{
			Name: "y",
			Type: shim.ColumnDefinition_UINT64,
			Key: false,
		},
		{
			Name: "is_locked",
			Type: shim.ColumnDefinition_BOOL,
			Key: false,
		},
		{
			Name: "lock_price",
			Type: shim.ColumnDefinition_UINT64,
			Key: false,
		},
		{
			Name: "is_checked",
			Type: shim.ColumnDefinition_BOOL,
			Key: false,
		},
		{
			Name: "is_clear",
			Type: shim.ColumnDefinition_BOOL,
			Key: false,
		},
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (c *Contract) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return c.Init(stub, "init", args)
	} else if function == "registerCinema" {
		return c.registerCinema(stub, args)
	} else if function == "registerTicketPlatform" {
		return c.registerTicketPlatform(stub, args)
	} else if function == "registerVideoHall" {
		return c.registerVideoHall(stub, args)
	} else if function == "planMovie" {
		return c.planMovie(stub, args)
	} else if function == "lockTicket" {
		return c.lockTicket(stub, args)
	} else if function == "checkTicket" {
		return c.checkTicket(stub, args)
	} else if function == "clear" {
		return c.clear(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (c *Contract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "queryTicket" { //read a variable
		return c.queryTicket(stub, args)
	} else if function == "queryAllPlan" {
		return c.queryAllPlan(stub, args)
	} else if function == "queryCinema" {
		return c.queryCinema(stub, args)
	} else if function == "queryTicketPlatform" {
		return c.queryTicketPlatform(stub, args)
	} else if function == "queryVideoHall" {
		return c.queryVideoHall(stub, args)
	} else if function == "queryPlan" {
		return c.queryPlan(stub, args)
	} else if function == "queryMovie" {
		return c.queryMovie(stub, args)
	}

	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}

func (c *Contract) registerCinema(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running registerCinema")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 args:name, company")
	}
	name := args[0]
	company := args[1]
	fmt.Printf("name = %s, company = %s", name, company)
	success, err := stub.InsertRow("cinema", shim.Row{Columns: []*shim.Column{
		{
			Value: &shim.Column_String_{String_:name},
		},
		{
			Value: &shim.Column_String_{String_:company},
		},
	}})
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, fmt.Errorf("Insert cinema %s false, may be table not found or row already exist", name)
	}
	return nil, nil
}

func (c *Contract) registerTicketPlatform(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running registerTicketPlatform")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:name")
	}
	name := args[0]
	fmt.Printf("name = %s", name)
	success, err := stub.InsertRow("ticket_platform", shim.Row{Columns: []*shim.Column{
		{
			Value: &shim.Column_String_{String_:name},
		},
	}})
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, fmt.Errorf("Insert ticket platform %s false, may be table not found or row already exist", name)
	}
	return nil, nil
}

func (c *Contract) registerVideoHall(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running registerVideoHall")
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4 args:name, cinema, width, height")
	}
	name := args[0]
	cinema := args[1]
	width, err := strconv.ParseUint(args[2], 10, 0)
	if err != nil {
		return nil, err
	}
	height, err := strconv.ParseUint(args[3], 10, 0)
	if err != nil {
		return nil, err
	}
	fmt.Printf("name = %s, cinema = %s, width = %d, height = %d", name, cinema, width, height)
	if width == 0 || width >= 20 {
		return nil, errors.New("Video hall seats array width can't be 0 or larger than 20")
	}
	if height == 0 || height >= 20 {
		return nil, errors.New("Video hall seats array height can't be 0 or larger than 20")
	}
	row, err := stub.GetRow("cinema", []shim.Column{
		{
			&shim.Column_String_{String_:cinema},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return nil, fmt.Errorf("Cinema %s doesn't exist", cinema)
	}
	success, err := stub.InsertRow("video_hall", shim.Row{Columns: []*shim.Column{
		{
			Value: &shim.Column_String_{String_:name},
		},
		{
			Value: &shim.Column_String_{String_:cinema},
		},
		{
			Value: &shim.Column_Uint64{Uint64:width},
		},
		{
			Value: &shim.Column_Uint64{Uint64:height},
		},
	}})
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, fmt.Errorf("Insert cinema %s video hall %s false, may be table not found or row already exist", cinema, name)
	}
	return nil, nil
}

func (c *Contract) planMovie(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running planMovie")
	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7 args:id, movie, cinema, video_hall, plan_time, start_time, end_time")
	}
	id := args[0]
	movie := args[1]
	cinema := args[2]
	video_hall := args[3]
	plan_time := args[4]
	start_time := args[5]
	end_time := args[6]
	fmt.Printf("id = %s, movie = %s, cinema = %s, video_hall = %s, plan_time = %s, start_time = %s, end_time = %s",
		id, movie, cinema, video_hall, plan_time, start_time, end_time)
	row, err := stub.GetRow("video_hall", []shim.Column{
		{
			&shim.Column_String_{String_:video_hall},
		},
		{
			&shim.Column_String_{String_:cinema},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return nil, fmt.Errorf("Cinema %s's video hall %s doesn't exist", cinema, video_hall)
	}
	width := row.Columns[2].GetUint64()
	height := row.Columns[3].GetUint64()

	success, err := stub.InsertRow("movie_plan", shim.Row{Columns: []*shim.Column{
		{
			Value: &shim.Column_String_{String_:id},
		},
		{
			Value: &shim.Column_String_{String_:movie},
		},
		{
			Value: &shim.Column_String_{String_:cinema},
		},
		{
			Value: &shim.Column_String_{String_:video_hall},
		},
		{
			Value: &shim.Column_String_{String_:plan_time},
		},
		{
			Value: &shim.Column_String_{String_:start_time},
		},
		{
			Value: &shim.Column_String_{String_:end_time},
		},
	}})
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, fmt.Errorf("Insert movie plan %s false, may be table not found or row already exist", id)
	}

	// 排片生成电影票
	var i, j uint64
	for i = 0; i < height; i++ {
		for j = 0; j < width; j++ {
			ticketID := fmt.Sprintf("%s:%d-%d", id, i, j)
			success, err := stub.InsertRow("ticket", shim.Row{Columns: []*shim.Column{
				{
					Value: &shim.Column_String_{String_:ticketID}, // id
				},
				{
					Value: &shim.Column_String_{String_:movie}, // movie
				},
				{
					Value: &shim.Column_String_{String_:id}, // movie_plan
				},
				{
					Value: &shim.Column_Uint64{Uint64:j}, // x
				},
				{
					Value: &shim.Column_Uint64{Uint64:i}, // y
				},
				{
					Value: &shim.Column_Bool{Bool:false}, // is_locked
				},
				{
					Value: &shim.Column_Uint32{Uint32:0}, // lock_price
				},
				{
					Value: &shim.Column_Bool{Bool:false}, // is_checked
				},
				{
					Value: &shim.Column_Bool{Bool:false}, // is_clear
				},
			}})
			if err != nil {
				return nil, err
			}

			if !success {
				return nil, fmt.Errorf("Insert ticket %s false, may be table not found or row already exist", ticketID)
			}
		}
	}

	r, err := stub.GetState(movie)
	if err != nil {
		return nil, err
	}
	var cr Clear
	if r != nil && len(r) > 0 {
		json.Unmarshal(r, cr)
	} else {
		cr = Clear{}
	}
	cr.IssueNum += width * height
	bs, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	stub.PutState(movie, bs)
	return nil, nil
}

func (c *Contract) lockTicket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running lockTicket")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 args:id, price")
	}
	id := args[0]
	price, err := strconv.ParseUint(args[1], 10, 0)
	if err != nil {
		return nil, err
	}
	fmt.Printf("id = %s, price = %d", id, price)
	row, err := stub.GetRow("ticket", []shim.Column{
		{
			&shim.Column_String_{String_:id},
		},
	})
	if err != nil {
		return nil, err
	}
	movie := row.Columns[1].GetString_()
	isLocked := row.Columns[5]
	lockPrice := row.Columns[6]
	if isLocked.GetBool() {
		return nil, fmt.Errorf("Ticket %s was locked already", id)
	}
	isLocked.Value = &shim.Column_Bool{Bool:true}
	lockPrice.Value = &shim.Column_Uint64{Uint64:price}
	success, err := stub.ReplaceRow("ticket", row)
	if err != nil {
		return nil, err
	}
	r, err := stub.GetState(movie)
	if err != nil {
		return nil, err
	}
	var cr Clear
	if r == nil || len(r) == 0 {
		return nil, errors.New("Movie clear result doesn't exsit")
	}
	json.Unmarshal(r, cr)
	cr.LockNum += 1
	cr.BoxOffice += price
	bs, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	stub.PutState(movie, bs)
	return json.Marshal(BoolResult{Success: success})
}

func (c *Contract) checkTicket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running checkTicket")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:id")
	}
	id := args[0]
	fmt.Printf("id = %s", id)
	row, err := stub.GetRow("ticket", []shim.Column{
		{
			&shim.Column_String_{String_:id},
		},
	})
	if err != nil {
		return nil, err
	}
	movie := row.Columns[1].GetString_()
	isChecked := row.Columns[7]
	if isChecked.GetBool() {
		return nil, fmt.Errorf("Ticket %s was checked already", id)
	}
	isChecked.Value = &shim.Column_Bool{Bool:true}
	success, err :=stub.ReplaceRow("ticket", row)
	if err != nil {
		return nil, err
	}
	r, err := stub.GetState(movie)
	if err != nil {
		return nil, err
	}
	var cr Clear
	if r == nil || len(r) == 0 {
		return nil, errors.New("Movie clear result doesn't exsit")
	}
	json.Unmarshal(r, cr)
	cr.CheckNum += 1
	bs, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	stub.PutState(movie, bs)
	return json.Marshal(BoolResult{Success: success})
}

func (c *Contract) clear(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running clear")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:movie")
	}
	movie := args[0]
	fmt.Printf("movie = %s", movie)
	r, err := stub.GetState(movie)
	if err != nil {
		return nil, err
	}
	var cr Clear
	if r == nil || len(r) == 0 {
		return nil, errors.New("Movie doesn't exsit")
	}
	json.Unmarshal(r, cr)
	regulationProfit := 0.083 * float32(cr.BoxOffice)
	availableProfit := 0.917 * float32(cr.BoxOffice)
	cinemaProfit := 0.5 * float32(availableProfit)
	cinemaCompanyProfit := 0.07 * float32(availableProfit)
	issuerProfit := 0.43 * float32(availableProfit)
	clr := ClearResult{
		Clear: cr,
		RegulationProfit: regulationProfit,
		CinemaProfit: cinemaProfit,
		CinemaCompanyProfit: cinemaCompanyProfit,
		IssuerProfit: issuerProfit,
	}
	return json.Marshal(clr)
}

func (c *Contract) queryTicket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryTicket")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:id")
	}
	id := args[0]
	fmt.Printf("id = %s", id)
	row, err := stub.GetRow("ticket", []shim.Column{
		{
			&shim.Column_String_{String_:id},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return []byte{}, nil
	}
	ticket := Ticket{
		ID: row.Columns[0].GetString_(),
		Movie: row.Columns[1].GetString_(),
		MoviePlan: row.Columns[2].GetString_(),
		X: row.Columns[3].GetUint64(),
		Y: row.Columns[4].GetUint64(),
		IsLocked: row.Columns[5].GetBool(),
		LockPrice: row.Columns[6].GetUint32(),
		IsChecked: row.Columns[7].GetBool(),
		IsClear: row.Columns[8].GetBool(),
	}
	return json.Marshal(ticket)

}

func (c *Contract) queryAllPlan(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryAllPlan")
	rowChan, err := stub.GetRows("movie_plan", []shim.Column{})
	if err != nil {
		return nil, err
	}
	var plans []MoviePlan
	timer := time.NewTimer(time.Minute)
	for {
		select {
		case row, ok := <- rowChan:
			if !ok {
				plans = nil
			} else {
				plans = append(plans, MoviePlan{
					ID: row.Columns[0].GetString_(),
					Movie: row.Columns[1].GetString_(),
					Cinema: row.Columns[2].GetString_(),
					VideoHall: row.Columns[3].GetString_(),
					PlanTime: row.Columns[4].GetString_(),
					StartTime: row.Columns[5].GetString_(),
					EndTime: row.Columns[6].GetString_(),
				})
			}
			if plans == nil {
				break
			}
		case <- timer.C:
			return nil, errors.New("query TimeOut")
		}
	}
	return json.Marshal(plans)
}

func (c *Contract) queryCinema(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryCinema")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:name")
	}
	name := args[0]
	fmt.Printf("name = %s", name)
	row, err := stub.GetRow("cinema", []shim.Column{
		{
			&shim.Column_String_{String_:name},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return []byte{}, nil
	}
	cinema := Cinema{
		Name: row.Columns[0].GetString_(),
		Company: row.Columns[1].GetString_(),
	}
	return json.Marshal(cinema)
}

func (c *Contract) queryTicketPlatform(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryTicketPlatform")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:name")
	}
	name := args[0]
	fmt.Printf("name = %s", name)
	row, err := stub.GetRow("ticket_platform", []shim.Column{
		{
			&shim.Column_String_{String_:name},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return []byte{}, nil
	}
	ticketPlatform := TicketPlatform{
		Name: row.Columns[0].GetString_(),
	}
	return json.Marshal(ticketPlatform)
}

func (c *Contract) queryVideoHall(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryVideoHall")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 args:name, cinema")
	}
	name := args[0]
	cinema := args[1]
	fmt.Printf("name = %s, cinema = %s", name, cinema)
	row, err := stub.GetRow("video_hall", []shim.Column{
		{
			&shim.Column_String_{String_:name},
		},
		{
			&shim.Column_String_{String_:cinema},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return []byte{}, nil
	}
	videoHall := VideoHall{
		Name: row.Columns[0].GetString_(),
		Cinema: row.Columns[1].GetString_(),
		Width: row.Columns[2].GetUint64(),
		Height: row.Columns[3].GetUint64(),
	}
	return json.Marshal(videoHall)
}

func (c *Contract) queryPlan(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryPlan")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:id")
	}
	id := args[0]
	fmt.Printf("id = %s", id)
	row, err := stub.GetRow("movie_plan", []shim.Column{
		{
			&shim.Column_String_{String_:id},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(row.Columns) == 0 {
		return []byte{}, nil
	}
	plan := MoviePlan{
		ID: row.Columns[0].GetString_(),
		Movie: row.Columns[1].GetString_(),
		Cinema: row.Columns[2].GetString_(),
		VideoHall: row.Columns[3].GetString_(),
		PlanTime: row.Columns[4].GetString_(),
		StartTime: row.Columns[5].GetString_(),
		EndTime: row.Columns[6].GetString_(),
	}
	return json.Marshal(plan)
}

func (c *Contract) queryMovie(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running queryMovie")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 args:movie")
	}
	movie := args[0]
	fmt.Printf("movie = %s", movie)

	return stub.GetState(movie)
}