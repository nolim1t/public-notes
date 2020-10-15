package main
import (
    "fmt"
    "io/ioutil"
    "strings"
)

func stringtofile(filename string) string {
    byte_output, err := ioutil.ReadFile(filename)
    if err == nil {
        return strings.Trim(string(byte_output), "\n")
    } else {
        return "error opening file"
    }
}

func main() {
    var battery_exists = "/sys/class/power_supply/axp20x-battery/present"
    var battery_capacity = "/sys/class/power_supply/axp20x-battery/capacity"

    if (stringtofile(battery_exists) == "1") {
        fmt.Printf("Battery Capacity: %s percent", stringtofile(battery_capacity))
    } else {
        fmt.Print("Battery doesnt exist")
    }
}
