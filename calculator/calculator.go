package calculator

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const IMPOSSIBLE = "IMPOSSIBLE"

type Parameter struct {
	NoOfTowns     int // N
	OfficeTown    int // T
	NoOfEmployees int // E
	Employees     []Employee
}

type Employee struct {
	Hometown       int // H
	NoOfPassengers int // P
}

// Custom sorting function, make the employees to be descending order by number of passengers the employee can drive and by town
type employeesByTownAndPassengers []Employee

func (e employeesByTownAndPassengers) Len() int {
	return len(e)
}
func (e employeesByTownAndPassengers) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e employeesByTownAndPassengers) Less(i, j int) bool {
	if e[i].Hometown == e[j].Hometown {
		return e[i].NoOfPassengers > e[j].NoOfPassengers
	}
	return e[i].Hometown < e[j].Hometown
}

func (p *Parameter) CanEveryoneMakeItToWork() interface{} {

	result := make([]int, p.NoOfTowns)
	employees := p.Employees
	currTown := 1

	// total number of employees in each town
	employeeInTown := 0
	// descending order array by number of passegers the employee can drive, will be cleared when next town
	noOfPassengersArr := make([]int, 0)
	// total number of passegers allowed, will be cleared when next town
	totalPassengersAllow := 0

	// Sorting by town, and for each town, sorting by the no of passagers
	sort.Stable(employeesByTownAndPassengers(employees))

	for _, e := range employees {
		if e.Hometown == p.OfficeTown {
			// If the employee is the office town, he/she can make it to work
			continue
		}

		if currTown != e.Hometown {
			// check whether last town meets the requirements
			carsNeeded := checkRequirements(employeeInTown, totalPassengersAllow, noOfPassengersArr)
			if carsNeeded < 0 {
				return IMPOSSIBLE
			}
			result[currTown-1] = carsNeeded

			currTown = e.Hometown
			// reset the var for next town
			noOfPassengersArr = make([]int, 0)
			employeeInTown = 0
			totalPassengersAllow = 0
		}

		if e.NoOfPassengers != 0 {
			totalPassengersAllow += e.NoOfPassengers
			noOfPassengersArr = append(noOfPassengersArr, e.NoOfPassengers)
		}
		employeeInTown++
	}

	// check requirements for the last town
	carsNeeded := checkRequirements(employeeInTown, totalPassengersAllow, noOfPassengersArr)
	if carsNeeded < 0 {
		return IMPOSSIBLE
	}
	result[currTown-1] = carsNeeded
	return result
}

func ReadInputAsParam(input string) ([]*Parameter, error) {
	params := make([]*Parameter, 0)

	lines := strings.Split(input, "\n")
	if len(lines) < 4 { // 1 test case at least contains 4 lines (1 C, 1 N&T, 1 E, and 1 H&P)
		return nil, errors.New("no data found or invalid input")
	}
	noOfTestCases, err := strconv.Atoi(lines[0])
	if err != nil {

		return nil, fmt.Errorf("error when parsing number of test cases - %s", err.Error())
	}
	index := 1 // index 0 is no of test cases
	for i := 0; i < noOfTestCases; i++ {
		p := &Parameter{
			Employees: make([]Employee, 0),
		}
		nAndT := strings.Split(strings.TrimSpace(lines[index]), " ")

		if len(nAndT) != 2 {
			return nil, errors.New("format error when parsing N and T")
		}
		p.NoOfTowns, err = strconv.Atoi(nAndT[0])
		if err != nil {
			return nil, fmt.Errorf("error when parsing N - %s", err.Error())
		}
		p.OfficeTown, err = strconv.Atoi(nAndT[1])
		if err != nil {
			return nil, fmt.Errorf("error when parsing T - %s", err.Error())
		}

		p.NoOfEmployees, err = strconv.Atoi(strings.TrimSpace(lines[index+1]))
		if err != nil {
			return nil, fmt.Errorf("error when parsing E - %s", err.Error())
		}

		for j := 0; j < p.NoOfEmployees; j++ {
			employee := Employee{}
			employeeData := strings.Split(strings.TrimSpace(lines[index+1+j+1]), " ")
			if len(employeeData) != 2 {
				return nil, fmt.Errorf("format error when parsing employee data")
			}
			employee.Hometown, err = strconv.Atoi(employeeData[0])
			if err != nil {
				return nil, fmt.Errorf("error when parsing employee hometown- %s", err.Error())
			}
			employee.NoOfPassengers, err = strconv.Atoi(employeeData[1])
			if err != nil {
				return nil, fmt.Errorf("error when parsing no of passengers- %s", err.Error())
			}
			p.Employees = append(p.Employees, employee)
		}
		index = index + 1 + p.NoOfEmployees + 1
		params = append(params, p)
	}
	return params, nil
}

func checkRequirements(employeeInTown, totalPassagersAllow int, noOfPassagersArr []int) int {
	carsNeeded := 0
	// if employee in the town more than total number of passesengers allowed in the town
	if employeeInTown > totalPassagersAllow {
		return -1 // -1 means impossible
	}
	for _, v := range noOfPassagersArr {
		employeeInTown = employeeInTown - v
		carsNeeded++
		if employeeInTown <= 0 {
			break
		}
	}
	return carsNeeded
}
