package encryption

import (
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/entities"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers"
	"testing"
	"time"
)

type structTest struct {
	Name string `json:"name"`
}

func TestTripleDesEncryptSuccess(t *testing.T) {
	scheduled := time.Now().Add(10 * time.Minute)
	fmt.Println(time.Now())
	fmt.Println(scheduled.After(time.Now()))
	fmt.Println(scheduled.Sub(time.Now()))
}

func TestTripleDesEncrypt(t *testing.T) {
	encryptedString, err := StringToTripleDesECBEncrypt("Ichwan Almaza", "1234567890ABCDEFGHIJKLMN")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("String : ", *encryptedString)

	encryptedStruct, err := StructToTripleDesECBEncrypt(structTest{Name: "Ichwan Almaza"}, "1234567890ABCDEFGHIJKLMN")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("Struct : ", *encryptedStruct)
}

func TestTripleDesDecrypt(t *testing.T) {

	decrypted, err := StringToTripleDesECBDecrypt("17657E57AF28318E724D85F2DD740925AC43702ECA466C10F566D092500727D695D768634578FE0FFBD26D2CC421CD916CB13C1CCC0009AF0F4F61F99535AD800F4AFAB2FED7BB8C348E87F2E6EAE2BA8D2C3F0BA9283C0B057713763A7573B87578FE971B819826071EB395585B5D42365F6CD95F6D0B6C6BB22A979DC178E4C88E827382D7DD02A41ACCB020C96B46", "1234567890ABCDEFGHIJKLMN")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("decrypted : ", *decrypted)

	structUser := entities.User{}
	err = helpers.JsonStringToStruct(*decrypted, &structUser)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(structUser)

	structUserReq := entities.User{}

	err = helpers.Mapper(&structUser, &structUserReq)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(structUserReq)
}
