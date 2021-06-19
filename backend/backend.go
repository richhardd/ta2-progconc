package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

type PeajeData struct {
	Anio         float64
	Mes          float64
	Dia          float64
	Codigo       string
	Tot_veh_pag  float64
	Tot_veh_exon float64
	Sent_cobro   float64
	Flujo_veh    int64
}

// ANIO,MES,DIA,CODIGO,TOT_VEH_PAG,TOT_VEH_EXON,SENT_COBRO,FLUJO_VEH
var dataset = []PeajeData{}

func convertirData(archivo [][]string) {
	for i := 1; i < len(archivo); i++ {
		anio, _ := strconv.ParseFloat(archivo[i][0], 64)
		mes, _ := strconv.ParseFloat(archivo[i][1], 64)
		dia, _ := strconv.ParseFloat(archivo[i][2], 64)
		codigo := archivo[i][3]
		tot_veh_pag, _ := strconv.ParseFloat(archivo[i][4], 64)
		tot_veh_exon, _ := strconv.ParseFloat(archivo[i][5], 64)
		sent_cobro, _ := strconv.ParseFloat(archivo[i][6], 64)
		flujo_veh, _ := strconv.ParseInt(archivo[i][7], 32, 32)

		var temp PeajeData = PeajeData{
			Anio:         anio,
			Mes:          mes,
			Dia:          dia,
			Codigo:       codigo,
			Tot_veh_pag:  tot_veh_pag,
			Tot_veh_exon: tot_veh_exon,
			Sent_cobro:   sent_cobro,
			Flujo_veh:    flujo_veh,
		}
		dataset = append(dataset, temp)
	}
}
func leerCSV(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
func main() {
	url := "https://raw.githubusercontent.com/richhardd/dataset-progconc/main/dataset.csv"
	Archivo, _ := leerCSV(url)
	convertirData(Archivo)
	fmt.Println(dataset)
}
