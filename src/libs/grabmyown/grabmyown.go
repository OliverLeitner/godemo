package grabmyown

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const myOwnAPI = "http://dnc.localnet/api/GetAllItems?type=Customer"

// should work with all json calls...
// thx to: https://github.com/gtl-hig/chuckjoke/blob/main/utils.go
func GetCustomers() ( /*json data array of type*/ []Customer, error) {
	req, err := http.NewRequest("GET", myOwnAPI, nil)

	var customers []Customer

	if err != nil {
		return customers, fmt.Errorf("No request formed %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return customers, fmt.Errorf("No response: %v", err)
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return customers, fmt.Errorf("Read error")
	}

	if err = json.Unmarshal(respData, &customers); err != nil {
		return customers, fmt.Errorf("Error in unmarsheling, %v", err)
	}
	return customers, nil
}

// print out first names in the shell
func LoopThroughCustomers() {
	var customers, error = GetCustomers()
	if error == nil {
		for i := 0; i < len(customers); i++ {
			fmt.Println(customers[i].ContactFirstName)
		}
	}
}

func CustomersAsArray() []Customer {
	var customers, error = GetCustomers()
	var localCustomers []Customer
	if error == nil {
		for i := 0; i < len(customers); i++ {
			localCustomers = append(localCustomers, customers[i])
		}
	}
	return localCustomers
}
