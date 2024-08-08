package converter

import (
	"avito-test/internal/models"
	"testing"
)

func TestFlatBuilderFromRawData(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		houseId uint64
		price   uint64
		rooms   byte
		wantErr bool
	}{
		{
			name:    "correct input",
			data:    []byte(`{"house_id": 1,"price": 3000,"rooms": 3}`),
			houseId: 1,
			price:   3000,
			rooms:   3,
			wantErr: false,
		},
		{
			name:    "no house_id input",
			data:    []byte(`{"price": 3000,"rooms": 3}`),
			wantErr: true,
		},
		{
			name:    "no price",
			data:    []byte(`{"house_id": 1,"rooms": 3}`),
			wantErr: true,
		},
		{
			name:    "no rooms",
			data:    []byte(`{"house_id": 1,"price": 3000}`),
			wantErr: true,
		},
		{
			name:    "invalid numbers",
			data:    []byte(`{"house_id": 1,"price": -3000,"rooms": -3}`),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FlatBuilderFromRawData(tt.data)
			t.Logf("running test %s", tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FlatBuilderFromRawData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == true {
				return
			}
			if tt.houseId != *got.HouseId {
				t.Errorf("FlatBuilderFromRawData() got = %d, expected %d", *got.HouseId, tt.houseId)
			}
			if tt.price != *got.Price {
				t.Errorf("FlatBuilderFromRawData() got = %d, expected %d", *got.HouseId, tt.houseId)
			}
			if uint64(tt.rooms) != uint64(*got.Rooms) {
				t.Errorf("FlatBuilderFromRawData() got = %d, expected %d", *got.HouseId, tt.houseId)
			}
		})
	}
}

func TestFlatFromFlatBuilder(t *testing.T) {
	tests := []struct {
		name    string
		houseId uint64
		price   uint64
		rooms   byte
		want    models.Flat
	}{
		{
			name:    "first",
			houseId: 1,
			price:   3000,
			rooms:   3,
		},
		{
			name:    "second",
			houseId: 154364,
			price:   300013213,
			rooms:   3,
		},
		{
			name:    "third",
			houseId: 1543641233,
			price:   132,
			rooms:   126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("running test %s", tt.name)
			builder := models.FlatBuilder{
				HouseId: &tt.houseId,
				Price:   &tt.price,
				Rooms:   &tt.rooms,
			}
			got := FlatFromFlatBuilder(builder)
			if tt.houseId != got.HouseId {
				t.Errorf("FlatFromFlatBuilder() got = %d, expected %d", got.HouseId, tt.houseId)
			}
			if tt.price != got.Price {
				t.Errorf("FlatFromFlatBuilder() got = %d, expected %d", got.HouseId, tt.price)
			}
			if tt.rooms != got.RoomNumber {
				t.Errorf("FlatFromFlatBuilder() got = %d, expected %d", got.RoomNumber, tt.rooms)
			}
		})
	}
}
