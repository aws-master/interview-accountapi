# How to run

```docker-compose up```

list non standard library :
- `github.com/google/uuid` for generate UUID
- `github.com/stretchr/testify` used for create unit test
- `github.com/vektra/mockery/v2@v2.20.0` mockery used for mocking the interface

refer to `example_client` folder, to show how to use this SDK/library. ```docker-compose up``` up will also run the test and alternatevely tou can run ```test -v``` in the main directory. 


## Usage

```golang 
func main() {
	var (
		version int64
	)
	client := form3.NewClient("http://localhost:8030")

	country := "GB"
	version = 1
	payload := models.PayloadCreateAccount{
		Attributes: &models.AccountAttributes{
			BankID:                  "123456",
			AccountNumber:           "123",
			BankIDCode:              "GBDSC",
			BaseCurrency:            "GBP",
			Bic:                     "EXMPLGB2XXX",
			Name:                    []string{"bob"},
			SecondaryIdentification: "SecondaryIdentification",
			Country:                 &country,
		},
		Version: &version,
	}
	resulCreate, err := client.CreateAccount(payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resulCreate %+v \n", resulCreate)

	resultFetch, err := client.FetchAccount(resulCreate.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resultFetch %+v \n", resultFetch)

	err = client.DeleteAccount(resulCreate.ID, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete request client error  %+v \n", err)
}
```
