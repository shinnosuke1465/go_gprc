package main
import (
	"fmt"
	// "io/ioutil"
	"log"
	// "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"go_grpc/pb"
)

func main() {
	//employeeのインスタンス生成
	employee := &pb.Employee{
		Id: 1,
		Name: "Suzuki",
		Email: "text@example.com",
		Occupation: pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project: map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year: 2000,
			Month: 1,
			Day: 1,
		},
	}

// //データのシリアライズ
// 	binData, err := proto.Marshal(employee)
// 	if err != nil {
// 		log.Fatalln("Can't seralize", err)
// 	}

// //ファイルへ書き出し
// 	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
// 		log.Fatalln("Can't write", err)
// 	}

// //書き出したファイルを読み込んでデシリアライズしてシリアライズがうまくいっているか確認
// 	in, err := ioutil.ReadFile("test.bin")
// 	if err != nil {
// 		log.Fatalln("Can't read file", err)
// 	}

// //空のemployee構造体を用意
// 	readEmployee := &pb.Employee{}

// //第一引数...ファイルから読み込んだバイナリデータ。第二引数...空の構造体。デシリアライズした結果が空の構造体に渡される
// 	err = proto.Unmarshal(in, readEmployee)
// 	if err != nil {
// 		log.Fatalln("Can't seralize", err)
// 	}

// //構造体を出力
// 	fmt.Println(readEmployee)

jsonData, err := protojson.Marshal(employee)
if err != nil {
	log.Fatalln("Can't marshal to json", err)
}

// fmt.Println(string(jsonData))

readEmployee := &pb.Employee{}
if err := protojson.Unmarshal(jsonData, readEmployee); err != nil {
	log.Fatalln("Can't unmarshal from json", err)
}

fmt.Println(readEmployee)

}