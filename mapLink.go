package main

import (
  "flag"
  "fmt"
  "strconv"
  "github.com/ztrue/googleMaps"
)

func main() {
  lat := flag.Float64("lat", 0, "latitude")
  lon := flag.Float64("lon", 0, "longitude")
  zoom := flag.Float64("zoom", 14, "zoom")
  flag.Parse()
  args := flag.Args()

  params := []*float64{lat, lon, zoom}
  for i, param := range params {
    if i < len(args) {
      value, err := strconv.ParseFloat(args[i], 64)
      if err != nil {
        fmt.Println(fmt.Sprintf("Argument %v is invalid", args[i]))
        return
      }
      *param = value
    }
  }

  fmt.Println(googleMaps.Link(*lat, *lon, *zoom))
}
