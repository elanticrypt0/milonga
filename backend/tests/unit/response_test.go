package app

import (
	"milonga/milonga/milonga_response"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestSendHTML(t *testing.T) {
	type args struct {
		c *fiber.Ctx
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := milonga_response.SendHTML(tt.args.c, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("SendHTML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendHTMLFromFile(t *testing.T) {
	type args struct {
		c        *fiber.Ctx
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := milonga_response.SendHTMLFromFile(tt.args.c, tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("SendHTMLFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendHTMLFromFileWithData(t *testing.T) {
	type args struct {
		c          *fiber.Ctx
		filepath   string
		data2parse map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := milonga_response.SendHTMLFromFileWithData(tt.args.c, tt.args.filepath, tt.args.data2parse); (err != nil) != tt.wantErr {
				t.Errorf("SendHTMLFromFileWithData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPrepareHTMLFromSlice(t *testing.T) {
	type args struct {
		filepath   string
		data2parse interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := milonga_response.PrepareHTMLFromSlice(tt.args.filepath, tt.args.data2parse)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareHTMLFromSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrepareHTMLFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

// No podemos probar parseHTMLFile directamente porque es una función privada
// Esta función puede probarse indirectamente a través de SendHTMLFromFileWithData
// pero requeriría mocks de fiber.Ctx

func TestStruct2Map(t *testing.T) {
	// Define una estructura de prueba
	type TestStruct struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Email    string `json:"email,omitempty"`
		Internal string // Sin etiqueta json
	}

	type args struct {
		estructura interface{}
	}
	
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "basic struct conversion",
			args: args{
				estructura: TestStruct{
					Name:     "John Doe",
					Age:      30,
					Email:    "john@example.com",
					Internal: "Internal value",
				},
			},
			want: map[string]interface{}{
				"name":     "John Doe",
				"age":      30,
				"email":    "john@example.com",
				"Internal": "Internal value",
			},
		},
		{
			name: "pointer to struct",
			args: args{
				estructura: &TestStruct{
					Name:     "Jane Doe",
					Age:      25,
					Email:    "jane@example.com",
					Internal: "Internal value",
				},
			},
			want: map[string]interface{}{
				"name":     "Jane Doe",
				"age":      25,
				"email":    "jane@example.com",
				"Internal": "Internal value",
			},
		},
		{
			name: "non-struct value",
			args: args{
				estructura: "not a struct",
			},
			want: map[string]interface{}{},
		},
		{
			name: "empty struct",
			args: args{
				estructura: TestStruct{},
			},
			want: map[string]interface{}{
				"name":     "",
				"age":      0,
				"email":    "",
				"Internal": "",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := milonga_response.Struct2Map(tt.args.estructura)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructSlice2Map(t *testing.T) {
	// Define una estructura de prueba
	type TestStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	type args struct {
		data interface{}
	}
	
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "slice of structs",
			args: args{
				data: []TestStruct{
					{Name: "Item 1", Value: 10},
					{Name: "Item 2", Value: 20},
					{Name: "Item 3", Value: 30},
				},
			},
			want: []interface{}{
				map[string]interface{}{"name": "Item 1", "value": 10},
				map[string]interface{}{"name": "Item 2", "value": 20},
				map[string]interface{}{"name": "Item 3", "value": 30},
			},
			wantErr: false,
		},
		{
			name: "empty slice",
			args: args{
				data: []TestStruct{},
			},
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name: "non-slice value",
			args: args{
				data: "not a slice",
			},
			want:    nil,
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := milonga_response.StructSlice2Map(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructSlice2Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructSlice2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
