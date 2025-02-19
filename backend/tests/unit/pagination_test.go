package app

import (
	"milonga/milonga/pagination"
	"reflect"
	"testing"
)

func TestNewPagination(t *testing.T) {
	type args struct {
		currentPage  uint
		itemsPerPage uint
		totalItems   uint
	}
	tests := []struct {
		name string
		args args
		want *pagination.Pagination
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pagination.NewPagination(tt.args.currentPage, tt.args.itemsPerPage, tt.args.totalItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagination_calculatePages(t *testing.T) {
	type fields struct {
		PageFirst    uint
		PageLast     uint
		PageCurrent  uint
		PageNext     uint
		PagePrev     uint
		ItemsPerPage uint
		TotalItems   uint
	}
	type args struct {
		totalItems uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pagination.Pagination{
				PageFirst:    tt.fields.PageFirst,
				PageLast:     tt.fields.PageLast,
				PageCurrent:  tt.fields.PageCurrent,
				PageNext:     tt.fields.PageNext,
				PagePrev:     tt.fields.PagePrev,
				ItemsPerPage: tt.fields.ItemsPerPage,
				TotalItems:   tt.fields.TotalItems,
			}
			p.calculatePages(tt.args.totalItems)
		})
	}
}

func TestPagination_ToString(t *testing.T) {
	type fields struct {
		PageFirst    uint
		PageLast     uint
		PageCurrent  uint
		PageNext     uint
		PagePrev     uint
		ItemsPerPage uint
		TotalItems   uint
	}
	type args struct {
		elem string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pagination.Pagination{
				PageFirst:    tt.fields.PageFirst,
				PageLast:     tt.fields.PageLast,
				PageCurrent:  tt.fields.PageCurrent,
				PageNext:     tt.fields.PageNext,
				PagePrev:     tt.fields.PagePrev,
				ItemsPerPage: tt.fields.ItemsPerPage,
				TotalItems:   tt.fields.TotalItems,
			}
			if got := p.ToString(tt.args.elem); got != tt.want {
				t.Errorf("Pagination.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
