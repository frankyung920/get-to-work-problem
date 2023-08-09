package calculator

import (
	"reflect"
	"sort"
	"testing"
)

func TestReadInputAsParam(t *testing.T) {
	type args struct {
		input string
	}
	p := &Parameter{
		NoOfTowns:     5,
		OfficeTown:    1,
		NoOfEmployees: 3,
		Employees: []Employee{
			{Hometown: 1, NoOfPassengers: 0},
			{Hometown: 1, NoOfPassengers: 0},
			{Hometown: 1, NoOfPassengers: 0},
		},
	}
	arr := make([]*Parameter, 0)
	arr = append(arr, p)
	tests := []struct {
		name    string
		args    args
		want    []*Parameter
		wantErr bool
	}{
		{
			name: "Valid input",
			args: args{
				input: "1\n5 1\n3\n1 0\n1 0\n1 0",
			},
			want:    arr,
			wantErr: false,
		},
		{
			name: "Invalid input - not enough lines",
			args: args{
				input: "1\n5 1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid input - format error",
			args: args{
				input: "1\n5 1\n3\n1 0\n1 0 0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid input - error parsing number",
			args: args{
				input: "1\n5 1\n3\n1 X\n1 0\n1 0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty input",
			args: args{
				input: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInputAsParam(tt.args.input)
			if err != nil && tt.wantErr == false {
				t.Errorf("ReadInputAsParam() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadInputAsParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeesByTownAndPassengers(t *testing.T) {
	employees := []Employee{
		{Hometown: 2, NoOfPassengers: 3},
		{Hometown: 1, NoOfPassengers: 2},
		{Hometown: 2, NoOfPassengers: 1},
		{Hometown: 1, NoOfPassengers: 3},
	}

	expected := []Employee{
		{Hometown: 1, NoOfPassengers: 3},
		{Hometown: 1, NoOfPassengers: 2},
		{Hometown: 2, NoOfPassengers: 3},
		{Hometown: 2, NoOfPassengers: 1},
	}

	sort.Sort(employeesByTownAndPassengers(employees))

	if !reflect.DeepEqual(employees, expected) {
		t.Errorf("Sort order is incorrect. Got: %v, want: %v", employees, expected)
	}
}

func TestParameter_CanEveryoneMakeItToWork(t *testing.T) {
	type fields struct {
		NoOfTowns     int
		OfficeTown    int
		NoOfEmployees int
		Employees     []Employee
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Valid case",
			fields: fields{
				NoOfTowns:     5,
				OfficeTown:    1,
				NoOfEmployees: 4,
				Employees: []Employee{
					{Hometown: 1, NoOfPassengers: 0},
					{Hometown: 2, NoOfPassengers: 2},
					{Hometown: 3, NoOfPassengers: 1},
					{Hometown: 4, NoOfPassengers: 3},
				},
			},
			want: []int{0, 1, 1, 1, 0},
		},
		{
			name: "Impossible case",
			fields: fields{
				NoOfTowns:     3,
				OfficeTown:    1,
				NoOfEmployees: 3,
				Employees: []Employee{
					{Hometown: 1, NoOfPassengers: 1},
					{Hometown: 2, NoOfPassengers: 0},
					{Hometown: 3, NoOfPassengers: 2},
				},
			},
			want: "IMPOSSIBLE",
		},
		{
			name: "No employees case",
			fields: fields{
				NoOfTowns:     4,
				OfficeTown:    1,
				NoOfEmployees: 0,
				Employees:     []Employee{},
			},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				NoOfTowns:     tt.fields.NoOfTowns,
				OfficeTown:    tt.fields.OfficeTown,
				NoOfEmployees: tt.fields.NoOfEmployees,
				Employees:     tt.fields.Employees,
			}
			if got := p.CanEveryoneMakeItToWork(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.CanEveryoneMakeItToWork() = %v, want %v", got, tt.want)
			}
		})
	}
}
