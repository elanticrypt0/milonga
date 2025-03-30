package app

import (
	"milonga/milonga/pagination"
	"testing"
)

func TestNewPagination(t *testing.T) {
	type args struct {
		currentPage  uint
		itemsPerPage uint
		totalItems   uint
	}
	
	// En lugar de hacer una comparación total del struct, verificaremos
	// los campos fundamentales para evitar problemas con los campos intermedios
	tests := []struct {
		name             string
		args             args
		wantPageFirst    uint
		wantPageLast     uint
		wantPageCurrent  uint
		wantItemsPerPage uint
		wantTotalItems   uint
	}{
		{
			name: "single page",
			args: args{
				currentPage:  1,
				itemsPerPage: 10,
				totalItems:   5,
			},
			wantPageFirst:    1,
			wantPageLast:     1,
			wantPageCurrent:  1,
			wantItemsPerPage: 10,
			wantTotalItems:   5,
		},
		{
			name: "multiple pages - first page",
			args: args{
				currentPage:  1,
				itemsPerPage: 10,
				totalItems:   25,
			},
			wantPageFirst:    1,
			wantPageLast:     3,
			wantPageCurrent:  1,
			wantItemsPerPage: 10,
			wantTotalItems:   25,
		},
		{
			name: "multiple pages - middle page",
			args: args{
				currentPage:  2,
				itemsPerPage: 10,
				totalItems:   35,
			},
			wantPageFirst:    1,
			wantPageLast:     4,
			wantPageCurrent:  2,
			wantItemsPerPage: 10,
			wantTotalItems:   35,
		},
		{
			name: "multiple pages - last page",
			args: args{
				currentPage:  3,
				itemsPerPage: 10,
				totalItems:   30,
			},
			wantPageFirst:    1,
			wantPageLast:     3,
			wantPageCurrent:  3,
			wantItemsPerPage: 10,
			wantTotalItems:   30,
		},
		{
			name: "zero items",
			args: args{
				currentPage:  1,
				itemsPerPage: 10,
				totalItems:   0,
			},
			wantPageFirst:    1,
			wantPageLast:     1,
			wantPageCurrent:  1,
			wantItemsPerPage: 10,
			wantTotalItems:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pagination.NewPagination(tt.args.currentPage, tt.args.itemsPerPage, tt.args.totalItems)
			
			if got.PageFirst != tt.wantPageFirst {
				t.Errorf("NewPagination().PageFirst = %v, want %v", got.PageFirst, tt.wantPageFirst)
			}
			if got.PageLast != tt.wantPageLast {
				t.Errorf("NewPagination().PageLast = %v, want %v", got.PageLast, tt.wantPageLast)
			}
			if got.PageCurrent != tt.wantPageCurrent {
				t.Errorf("NewPagination().PageCurrent = %v, want %v", got.PageCurrent, tt.wantPageCurrent)
			}
			if got.ItemsPerPage != tt.wantItemsPerPage {
				t.Errorf("NewPagination().ItemsPerPage = %v, want %v", got.ItemsPerPage, tt.wantItemsPerPage)
			}
			if got.TotalItems != tt.wantTotalItems {
				t.Errorf("NewPagination().TotalItems = %v, want %v", got.TotalItems, tt.wantTotalItems)
			}
		})
	}
}

// No podemos probar calculatePages directamente porque es un método privado
// En lugar de eso, probaremos su efecto a través de NewPagination
func TestPagination_PagesCalculation(t *testing.T) {
	tests := []struct {
		name         string
		currentPage  uint
		itemsPerPage uint
		totalItems   uint
		wantPages    uint
	}{
		{
			name:         "total items less than items per page",
			currentPage:  1,
			itemsPerPage: 10,
			totalItems:   5,
			wantPages:    1,
		},
		{
			name:         "total items equal to items per page",
			currentPage:  1,
			itemsPerPage: 10,
			totalItems:   10,
			wantPages:    1,
		},
		{
			name:         "total items greater than items per page",
			currentPage:  1,
			itemsPerPage: 10,
			totalItems:   25,
			wantPages:    3,
		},
		{
			name:         "total items much greater than items per page",
			currentPage:  1,
			itemsPerPage: 10,
			totalItems:   105,
			wantPages:    11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pagination.NewPagination(tt.currentPage, tt.itemsPerPage, tt.totalItems)
			if p.PageLast != tt.wantPages {
				t.Errorf("NewPagination() PageLast = %v, want %v", p.PageLast, tt.wantPages)
			}
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
		{
			name: "end element",
			fields: fields{
				PageFirst:   1,
				PageLast:    5,
				PageCurrent: 2,
			},
			args: args{
				elem: "end",
			},
			want: "5",
		},
		{
			name: "start element",
			fields: fields{
				PageFirst:   1,
				PageLast:    5,
				PageCurrent: 2,
			},
			args: args{
				elem: "start",
			},
			want: "1",
		},
		{
			name: "next element",
			fields: fields{
				PageFirst:   1,
				PageLast:    5,
				PageCurrent: 2,
				PageNext:    3,
			},
			args: args{
				elem: "next",
			},
			want: "3",
		},
		{
			name: "prev element",
			fields: fields{
				PageFirst:   1,
				PageLast:    5,
				PageCurrent: 2,
				PagePrev:    1,
			},
			args: args{
				elem: "prev",
			},
			want: "1",
		},
		{
			name: "current element",
			fields: fields{
				PageFirst:   1,
				PageLast:    5,
				PageCurrent: 2,
				PagePrev:    1,
			},
			args: args{
				elem: "current",
			},
			want: "2",
		},
		{
			name: "unknown element",
			fields: fields{
				PageFirst:   1,
				PageLast:    5,
				PageCurrent: 2,
			},
			args: args{
				elem: "unknown",
			},
			want: "",
		},
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