package main
import(
    "net/http"
    "fmt"
    "io/ioutil"
)
func main() {
    url := "https://carbonfootprint1.p.rapidapi.com/CarbonFootprintFromCarTravel?distance=100&vehicle=SmallDieselCar"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Print(err.Error())
    }
}