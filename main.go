package main

import (
	"bytes"
	"fmt"
	echo2 "github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Initializing Application....")
	//ProcessData(strings.NewReader("olatunde is my boy"))

	//err := LoadConfig()
	//panic(err)
	//CreateFile("Oluwa is my KING", "hello.text")

	StartApp(3230)
}

func ProcessData(reader io.Reader) {
	b := make([]byte, 2)

	for {
		count, err := reader.Read(b)
		if count > 0 {
			fmt.Printf(" I have read %v of data", string(b[:count]))
		}

		if err == io.EOF {
			break
		}
	}

}
func GetPwd() {

}

func StartApp(port int) {
	app := echo2.New()

	app.GET("/", func(c echo2.Context) error {
		return c.String(http.StatusOK, "Welcome to ShortUrl.....")
	})
	/*
		goDapper := packages.NewGoDapper(&packages.GoDapperConfig{
			Dialect:  constants.MYSQL,
			User:     "root",
			Password: "DVorak@230",
			DbName:   "loan_db",
			DbPort:   3306,
			DbHost:   "localhost",
		})

		if _, err := goDapper.Open(); err != nil {
			panic(err)
		} else {
			fmt.Println("db connected")
		}
	*/

	app.Logger.Fatal(app.Start(":3230"))
	app.Logger.Info("Application started.......")
}

func CreateFile(content string, filename string) {
	fileLocation := ConcatStrings("", filename)
	fmt.Println("file location is ", fileLocation)
	fileHandler, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fileHandler.Close()

	numb, errReadingFile := fileHandler.WriteString(content)
	if errReadingFile != nil {
		panic(errReadingFile)
	}

	fmt.Printf("number written so far %d", numb)
}

func ConcatStrings(args ...string) string {
	var byteBuffer bytes.Buffer
	for v := range args {
		byteBuffer.WriteString(args[v])
	}

	return byteBuffer.String()
}
func LoadConfig() (err error) {
	dir, dirError := os.Getwd()
	fmt.Println(dir)
	fmt.Println(dirError)

	data, err := os.ReadFile("config.json")
	if err == nil {
		fmt.Println("data read is ")
		fmt.Println(string(data))
	}

	return err
}

// https://tutorialedge.net/courses/go-data-structures-course/
