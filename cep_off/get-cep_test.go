package cep_off

import (
	"reflect"
	"testing"
)

func TestGetInfoCep(t *testing.T) {
	type args struct {
		cep string
	}
	tests := []struct {
		name string
		args args
		want *CepResponse
	}{

		{
			name: "Ensure can get information about CEP",
			args: args{cep: "28630590"},
			want: &CepResponse{
				Cep:         "28630-590",
				Logradouro:  "Avenida Antônio Mário de Azevedo",
				Complemento: "",
				Unidade:     "",
				Bairro:      "Campo do Coelho",
				Localidade:  "Nova Friburgo",
				Uf:          "RJ",
				Estado:      "Rio de Janeiro",
				Regiao:      "Sudeste",
				Ibge:        "3303401",
				Gia:         "",
				Ddd:         "22",
				Siafi:       "5867",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInfoCep(tt.args.cep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfoCep() = %v, want %v", got, tt.want)
			}
		})
	}
}
