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

func Test_parseHTMLFile(t *testing.T) {
	type args struct {
		filepath   string
		data2parse map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := milonga_response.parseHTMLFile(tt.args.filepath, tt.args.data2parse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseHTMLFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStruct2Map(t *testing.T) {
	type args struct {
		estructura interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := milonga_response.Struct2Map(tt.args.estructura); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct2Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructSlice2Map(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
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
