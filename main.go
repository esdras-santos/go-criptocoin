package main

import (
	//"os"
	//"gochain/cli"
	"encoding/binary"
	"fmt"

)


func main(){
	//defer os.Exit(0) 
	//cmd := cli.CommandLine{}
	//cmd.Run()


	fmt.Println(0x100)
}

func mapkey(m map[byte]string, value string) (key byte, ok bool) {
	for k, v := range m {
	  if v == value { 
		key = k
		ok = true
		return
	  }
	}
	return
}

func ReadVarint(s []byte, buf *uint){
	i := s[0]
	if i == 0xfd{
		a := binary.LittleEndian.Uint16(s[1:3])
		*buf = uint(a)
	}else if i == 0xfe{
		a := binary.LittleEndian.Uint32(s[1:5])
		*buf = uint(a)
	}else if i == 0xff{
		a := binary.LittleEndian.Uint64(s[1:9])
		*buf = uint(a)
	}else{
		*buf = uint(i)
	}
}

func toLittleEndian(bytes []byte) []byte{
	var le []byte
	for i := len(bytes)-1;i >= 0;i--{
		le = append(le, bytes[i]) 
	}
	return le
}