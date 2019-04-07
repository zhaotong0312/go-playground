package try_echo

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func handler(c echo.Context) error {
	return c.JSON(http.StatusOK, "lol")
}


func generateCases(level, i int, temp []bool, result*[][]bool) {
	if level == i {
		res := make([]bool, level)
		copy(res, temp)
		*result = append(*result, res)
	} else {
		temp[i] = false
		generateCases(level, i+1, temp, result)
		temp[i] = true
		generateCases(level, i+1, temp, result)
	}
}

func app(arr []int) {
	arr = append(arr, 1)
	arr = append(arr, 1)
}

func main() {
	//e := echo.New()
	//e.POST("/dns/haha", handler)
	//reqData := []byte("")
	//
	//
	//req := httptest.NewRequest(http.MethodPost, "/dns", bytes.NewReader(reqData))
	//rec := httptest.NewRecorder()
	//e.Serv
	//
	//eHTTP(rec, req)
	////c := e.NewContext(req, rec)
	//
	//fmt.Println(rec.Body.String())
	//
	//
	result := make([][]bool, 0)
	generateCases(5,0, make([]bool, 5), &result)
	fmt.Println(result)
}
