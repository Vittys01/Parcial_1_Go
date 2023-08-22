package tickets

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID          int
	nombre      string
	email       string
	paisDestino string
	horaVuelo   string
	precio      int
}

/*Armar form global donde trae la info
func getInfo()(data [][]string) {
	path := "./tickets.csv"
	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	splittedData := strings.Split(string(rawData), "\n")
	for _, item := range splittedData {
		splittedItems := strings.Split(item, ",")
	}
}
*/
// ejemplo 1 //3 indice pais
func GetTotalTickets(destination string) (int, error) {
	//Verificar scopes = TotalTickets
	path := "./tickets.csv"
	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	splittedData := strings.Split(string(rawData), "\n")
	var totalTickets int
	for i, item := range splittedData {
		splittedItems := strings.Split(item, ",")
		if i != 0 {
			if destination == splittedItems[3] {
				totalTickets++
			}
		}
	}
	return totalTickets, nil
}

// ejemplo 2 //4 indice hora
func GetMornings(time string) (int, error) {
	// Valida hora menor y mayor 0 Split y agarrar primer numero
	// For dentro de case o fuera
	/*Intervalo de tiempo*/
	madrugada := [2]int{0, 6}
	mañana := [2]int{7, 12}
	tarde := [2]int{13, 19}
	noche := [2]int{20, 23}

	var interval [2]int
	switch time {
	case "madrugada":
		interval = madrugada
	case "mañana":
		interval = mañana
	case "tarde":
		interval = tarde
	case "noche":
		interval = noche
	default:
		return 0, errors.New("tiempo no válido: " + time)
	}

	path := "./tickets.csv"
	rawData, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	splittedData := strings.Split(string(rawData), "\n")
	var totalPersons int
	for i, item := range splittedData {
		splittedItems := strings.Split(item, ",")
		if i != 0 {
			splittedTime := strings.Split(splittedItems[4], ":") //Se hace un split para cortar la hora ya que no se les da relevancia en este algoritmo
			hourStr := strings.TrimSpace(splittedTime[0])
			hour, err := strconv.Atoi(hourStr)
			if err != nil {
				return 0, err
			}

			if hour >= interval[0] && hour <= interval[1] {
				totalPersons++
			}
		}
	}
	return totalPersons, nil
}

/*
// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	tickets := GetTotalTickets(destination)
	percentage := tickets / total * 100
}
*/
