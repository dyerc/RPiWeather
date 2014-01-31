package main

import(
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func readTemperature() float64 {
  buf, err := ioutil.ReadFile("/sys/bus/w1/devices/28-0000054d8438/w1_slave")
  if err != nil {
    fmt.Println("Error: %s\n", err)
    return 0
  }

  lines := strings.Split(string(buf), "\n")
  temp_data := strings.Split(lines[1], " ")[9]
  temp := strings.TrimPrefix(temp_data, "t=")

  f, err := strconv.ParseFloat(temp, 64)
  if err != nil {
    fmt.Println("Error: %s\n", err)
    return 0
  }

  return f / 1000
}

func main() {
  fmt.Println("Temperature: ", readTemperature())
}
