package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type DataStore struct {
	Provinces    []Province
	ProvincesMap map[string]Province
	Regencies    map[string][]Regency
	RegenciesMap map[string]Regency
	Districts    map[string][]District
	DistrictsMap map[string]District
	Villages     map[string][]Village
	VillagesMap  map[string]Village
}

func NewDataStore(dataDir string) *DataStore {
	store := &DataStore{
		ProvincesMap: make(map[string]Province),
		Regencies:    make(map[string][]Regency),
		RegenciesMap: make(map[string]Regency),
		Districts:    make(map[string][]District),
		DistrictsMap: make(map[string]District),
		Villages:     make(map[string][]Village),
		VillagesMap:  make(map[string]Village),
	}

	store.loadProvinces(dataDir + "/provinces.csv")
	store.loadRegencies(dataDir + "/regencies.csv")
	store.loadDistricts(dataDir + "/districts.csv")
	store.loadVillages(dataDir + "/villages.csv")

	return store
}

func (s *DataStore) loadProvinces(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Warning: failed to open provinces file: %v", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read provinces record: %v", err)
		}

		p := Province{ID: record[0], Name: record[1]}
		s.Provinces = append(s.Provinces, p)
		s.ProvincesMap[p.ID] = p
	}
	log.Printf("Loaded %d provinces", len(s.Provinces))
}

func (s *DataStore) loadRegencies(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Warning: failed to open regencies file: %v", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read regencies record: %v", err)
		}

		r := Regency{ID: record[0], ProvinceID: record[1], Name: record[2]}
		s.Regencies[r.ProvinceID] = append(s.Regencies[r.ProvinceID], r)
		s.RegenciesMap[r.ID] = r
		count++
	}
	log.Printf("Loaded %d regencies", count)
}

func (s *DataStore) loadDistricts(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Warning: failed to open districts file: %v", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read districts record: %v", err)
		}

		d := District{ID: record[0], RegencyID: record[1], Name: record[2]}
		s.Districts[d.RegencyID] = append(s.Districts[d.RegencyID], d)
		s.DistrictsMap[d.ID] = d
		count++
	}
	log.Printf("Loaded %d districts", count)
}

func (s *DataStore) loadVillages(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Warning: failed to open villages file: %v", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read villages record: %v", err)
		}

		v := Village{ID: record[0], DistrictID: record[1], Name: record[2]}
		s.Villages[v.DistrictID] = append(s.Villages[v.DistrictID], v)
		s.VillagesMap[v.ID] = v
		count++
	}
	log.Printf("Loaded %d villages", count)
}
