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

	decrypted, err := StringToTripleDesECBDecrypt("17657E57AF28318E2B72E68A81D3A78114B5A1FC93AF00D7A6C379D4CC180D7D88887BCD2ABB4FBF96C92235AD8D511000686C088F2A43D0071975CCE8960B855A3215315D324844E83AC28BD0193FE8E9778541558F6D2BE38594969AA2A7DF69D74C2477F78D0CB371D8107AAD883BA917D9A4E44D18B94D86A637685A7576D7C6A3FA336A341C04C8E4CD7B812595440BA38AA51F6CAF13367F70A3D755E3937DFB57412FB69B7AC2C99E6F39E35BF84B2D9768B1E9DCE1C5D143944DA57189C943BB5A02B7A1CDA0C44DD25E70AE97AB8801E4FC459E790F20FA1BD55210485847D115DFE8AC7568BA183751B2EAA4DD365ECDD9B0470B848AD756F9DDCB8195E056FF7DC9BDE45D3DF37B295006130902227DA2425FE45C598478FD60084297F321F5375E5405177CAD05CE071F16E748A1B6B993168195E056FF7DC9BD6AC8C75618D77DFB4905ACE35110856B485847D115DFE8AC7568BA183751B2EAA4DD365ECDD9B047A22E8BB604CD73E5", "1234567890ABCDEFGHIJKLMN")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("decrypted : ", *decrypted)

	structUser := entities.User{}
	err = helpers.JsonStringToStruct(*decrypted, &structUser)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(structUser)
}
