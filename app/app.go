package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Option int

const (
	// OptionProvince province
	OptionProvince Option = iota
	// OptionProvinceCity province AND city
	OptionProvinceCity
	// OptionProvinceCityCounty province AND city AND county
	OptionProvinceCityCounty
	// OptionProvinceCityCountyTown province AND city AND county AND town
	OptionProvinceCityCountyTown
	// OptionProvinceCityCountyTownVillage province AND city AND county AND town AND village
	OptionProvinceCityCountyTownVillage
)

type ChinaDivision struct {
	Option   Option
	OutPath  string
	FileName string
}

func Run(d ChinaDivision) error {
	fmt.Println("step[1] => http get")
	res, err := TectonicNodes(d.Option)
	if err != nil {
		return err
	}

	fmt.Println("step[2] => json marshal")
	by, err := json.Marshal(res)
	if err != nil {
		return err
	}

	fmt.Println("step[3] => write file")
	err = os.MkdirAll(filepath.Dir(d.OutPath), 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(d.FileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(by)

	return err
}
