package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundData struct {
	Message string
}

func (n *notFoundData) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "Fatra" {
		return &notFoundData{"Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("Fatra", nil)
	if err != nil {
		/*
			if validationErr, oke := err.(*validationError); oke {
				fmt.Println("Validation error : ", validationErr.Error())
			} else if notFoundDataErr, oke := err.(*notFoundData); oke {
				fmt.Println("Not Found Data : ", notFoundDataErr.Error())
			} else {
				fmt.Println("Unknown Error : ", err.Error())
			}
		*/

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error : ", finalError.Error())
		case *notFoundData:
			fmt.Println("Not Found Data : ", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}
	/*
		if err != nil {
			fmt.Println(err)
		}

		err = SaveData("unknown", nil)
		if err != nil {
			fmt.Println(err)
		}

		err = SaveData("Fatra", nil)
		if err != nil {
			fmt.Println(err)
		}
	*/

}
