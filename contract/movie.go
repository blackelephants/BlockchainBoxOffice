package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
	"errors"
)

// 电影院
type Cinema struct {
	Name string
	Company string
}

// 影厅
type VideoHall struct {
	Cinema Cinema
	Name string
	Seats [][]int
}

// 票务平台
type TicketPlatform struct {
	Name string
}

// 排片
type MoviePlan struct {
	ID string
	Movie string
	Cinema Cinema
	VideoHall VideoHall
	PlanTime time.Time
	StartTime time.Time
	EndTime time.Time
}

// 电影票
type Ticket struct {
	ID string
	MoviePlan MoviePlan
	X int
	Y int
	IsLocked bool
	LockPrice float64
	LockTime time.Time
	IsChecked bool
	IsClear bool
}

// 分账结果
type ClearResult struct {
	IssueNum int
	LockNum int
	CheckNum int
	BoxOffice float64
	RegulationProfit float64
	CinemaProfit float64
	CinemaCompanyProfit float64
	IssuerProfit float64
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
func (t *Contract) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("init is running ")
	stub.CreateTable("cinema", []*shim.ColumnDefinition{
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
	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *Contract) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "registerCinema" {
		return t.registerCinema(stub, args)
	} else if function == "registerTicketPlatform" {
		return t.registerTicketPlatform(stub, args)
	} else if function == "registerVideoHall" {
		return t.registerVideoHall(stub, args)
	} else if function == "planMovie" {
		return t.planMovie(stub, args)
	} else if function == "lockTicket" {
		return t.lockTicket(stub, args)
	} else if function == "checkTicket" {
		return t.checkTicket(stub, args)
	} else if function == "clear" {
		return t.clear(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *Contract) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "queryTicket" { //read a variable
		return t.queryTicket(stub, args)
	} else if function == "queryPlan" {
		return t.queryPlan(stub, args)
	}

	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}

func (t *Contract) registerCinema(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running registerCinema")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 args:name, company")
	}
	name := args[0]
	company := args[1]
	fmt.Printf("name = %s, company = %s", name, company)
	stub.InsertRow("cinema", shim.Row{[]*shim.Column{
		{
			&shim.Column_String_{String_:name},
		},
		{
			&shim.Column_String_{String_:company},
		},
	}})
	row, err := stub.GetRow("cinema", []shim.Column{
		{
			&shim.Column_String_{String_:name},
		},
	})
	if err != nil {
		return nil, err
	}
	return []byte(row.String()), nil
}

func (t *Contract) registerTicketPlatform(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) registerVideoHall(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) planMovie(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) lockTicket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) checkTicket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) clear(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) queryTicket(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}

func (t *Contract) queryPlan(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return nil, nil
}