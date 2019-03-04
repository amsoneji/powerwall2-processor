package common
import "github.com/fatih/structs"

type Meters_aggregates_response struct {
  Site       map[string]interface{}  `json:"site"`
  Battery    map[string]interface{}  `json:"battery"`
  Load       map[string]interface{}  `json:"load"`
  Solar      map[string]interface{}  `json:"solar"`
  Busway     map[string]interface{}  `json:"busway"`
  Frequency  map[string]interface{}  `json:"frequency"`
  Generator  map[string]interface{}  `json:"generator"`
}

type Meters_site_solar_response struct {
  Id              int                     `json:"id"`
  Location        string                  `json:"location"`
  Type            string                  `json:"type"`
  Cts             [4] bool                `json:"cts"`
  Inverted        [4] bool                `json:"inverted"`
  Connection      map[string]interface{}  `json:"connection"`
  Cached_Readings map[string]interface{}  `json:"Cached_readings"`
}

func Get_all_paths() ([]string){
  return []string{
    "/api/meters/aggregates",
    "/api/meters/site",
    "/api/meters/solar",
  }
}

func Get_obj_for_path(path string) (map[string]interface{}, []map[string]interface{}){
  switch path {
  case "/api/meters/aggregates":
    return structs.Map(Meters_aggregates_response{}), nil
  case "/api/meters/site":
    return nil, []map[string]interface{}{structs.Map(Meters_site_solar_response{})}
  case "/api/meters/solar":
    return nil, []map[string]interface{}{structs.Map(Meters_site_solar_response{})}
  }
  return nil, nil
}
