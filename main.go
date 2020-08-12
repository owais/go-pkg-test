package main

import "fmt"
import "golang.org/x/tools/go/packages"
import _ "github.com/gogo/protobuf/gogoproto"

func main() {
	pkgs, _ := packages.Load(&packages.Config{}, "github.com/gogo/protobuf/gogoproto")
	p := pkgs[0]
	fmt.Println(p.GoFiles)
	// prints the following
	// [/Users/olone/go/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto/doc.go /Users/olone/go/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto/gogo.pb.go /Users/olone/go/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto/helper.go]
}
