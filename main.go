package main

import (
    "fmt"
    "github.com/gogo/protobuf/jsonpb"
    "github.com/cajus/prototest/proto"
    "strings"
)

func main() {
    m := jsonpb.Marshaler{}
    str, err := m.MarshalToString(&pb.ParameterSets{Sets: map[string]*pb.ParameterSet{
        "parameter": {
            Raw: []byte{0, 1, 2, 3, 4, 5, 6, 1, 2},
            Parameters: []*pb.Parameter{
                {Name: "itest", Value: &pb.Parameter_IntValue{42}},
                {Name: "stest", Value: &pb.Parameter_StrValue{"Foobar"}},
                {Name: "btest", Value: &pb.Parameter_BoolValue{true}},
                {Name: "ftest", Value: &pb.Parameter_FloatValue{41.99}},
           },
        },
    }},)
    fmt.Println(str)
    fmt.Println(err)

    sets := pb.ParameterSets{}
    err = jsonpb.Unmarshal(strings.NewReader(str), &sets)
    fmt.Println(sets)
    fmt.Println(err)
}
