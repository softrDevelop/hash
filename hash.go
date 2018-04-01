package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/softrDevelop/hash/help"
	"hash"
	"io"
	"log"
	"os"
)

type args struct {
	algo     string
	filePath string
}

// getNewHashHash returns passed algo type as a hash.Hash
func getNewHashHash (AlgoType string) (hash.Hash, error){
	switch AlgoType {
	case "-sha256":
		return sha256.New(), nill
	case "-sha512":
		return sha512.New(), nill
	case "-md5":
		return md5.New(), nill
	default:
		return nil,fmt.Errorf("Unknow algo type %v/n",AlgoType)
	}
}


// generate returns hash if success
func (a args) generate() string {
	f, err := os.Open(a.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	h, err := getNewHashHash(a.algo)

	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// lenHashType returns the blocksizes of the passed hash
func lenHashType (AlgoType string) int{
	switch AlgoType {
	case "-sha256":
		return sha256.BlockSize
	case "-sha512":
		return sha512.BlockSize
	case "-md5":
		return md5.BlockSize
	default:
		return -1
	}

}

// lenghtOK returns true if passed hash has corresponing blocksize
func (a args)lengthOK(v string) {
	lenType := lenHashType(a.algo)

	if len(v) < lenType {
		log.Fatal("Hash in argument too short ", len(v),"/",lenType)
	}
	if len(v) > lenType {
		log.Fatal("Hash in argument too long ", len(v),"/",lenType)
	}
}

func main() {

	//We need at least 2 arguments else print out help
	if len(os.Args) < 3 {
		help.View()
		return
	}

	arguments := args{
		algo:     os.Args[1],
		filePath: os.Args[2],
	}

	//If there is a hash in the third argument do compare
	if len(os.Args) >= 4 {
		//Check if length of hash is ok
		arguments.lengthOK(os.Args[3])
		//Compare hashes
		if os.Args[3] != arguments.generate() {
			log.Fatal("Hash mismatch")
		}
		fmt.Println("Hash matching")
		return
	}

	//Else just print out hash of file
	fmt.Println(arguments.generate())
	return
}
