package help

import "fmt"

func View()  {
	fmt.Println("Hash is a simple tool for generating hashes from file or comparing a file with a supplied hash")
	fmt.Println()
	fmt.Println("\thash <algo> <file path> [hash]\n")
	fmt.Println("The algos are:\n")
	fmt.Println("\t-sha256")
	fmt.Println("\t-sha512")
	fmt.Println("\t-md5")
	fmt.Println()
	fmt.Println("Examples:\n")
	fmt.Println("\tGenerate sha256 hash from myFile")
	fmt.Println("\t\t",`hash -sha256 c:\myFile`,"\n")
	fmt.Println("\tCompares md5 hash from myFile with hash in argument")
	fmt.Println("\t\t",`hash -md5 c:\myFile a44b49b6bc6248f7ee5b5f2626286983`,"\n")
}
