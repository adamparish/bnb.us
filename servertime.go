package main

func main() {
  response, err := http.Get("https://api.binance.us/api/v3/time")
  if err != nil {
      fmt.Printf("Connection to Binance.US failed %s\n", err)
  }  else {
      data, _ := iotil.ReadAll(response.Body)
      fmt.Println(string(data))
   } 
}
