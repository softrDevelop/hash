package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/softrDevelop/hash/help"
	"hash"
	"io"
	"os"
)

type args struct {
	algo     string
	filePath string
}

// getNewHashHash returns passed algo type as a hash.Hash
func getNewHashHash(AlgoType string) (hash.Hash, error) {
	switch AlgoType {
	case "-sha256":
		return sha256.New(), nil
	case "-sha512":
		return sha512.New(), nil
	case "-md5":
		return md5.New(), nil
	default:
		return nil, fmt.Errorf("unknow algo type %v", AlgoType)
	}
}

// generate returns hash if success or zero value string and error if error encountered
func (a args) generate() (string, error) {

	f, err := os.Open(a.filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h, err := getNewHashHash(a.algo)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// lenHashType returns the blocksizes of the passed hash, -1 if missmatch
func lenHashType(AlgoType string) int {
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

// lenghtOK returns true if passed hash has corresponding blocksize
func (a args) lengthOK(v string) error {
	lenType := lenHashType(a.algo)

	if len(v) < lenType {
		return fmt.Errorf("hash in argument too short %v/%v", len(v), lenType)
	}
	if len(v) > lenType {
		return fmt.Errorf("hash in argument too long %v/%v", len(v), lenType)
	}
	return nil
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
		err := arguments.lengthOK(os.Args[3])
		if err != nil {
			fmt.Println(err)
			return
		}

		// Generate hash
		genedHash, err := arguments.generate()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Compare hashes
		if os.Args[3] != genedHash {
			fmt.Println("hash mismatch")
			return
		}
		fmt.Println("hash matching")
		return
	}

	//Else just print out hash of
	genedHash, err := arguments.generate()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(genedHash)
	return
}
