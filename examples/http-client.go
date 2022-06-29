package examples

import (
	"fmt"
	"io"
	"net/http"

	"github.com/saratonite/network-examples/utils"
)

func HttpClient(url string) {

	res, err := http.Get(url)

	utils.CheckError(err, "Error GET request")

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	utils.CheckError(err, "Cant read response body")

	fmt.Println(string(body))
}
