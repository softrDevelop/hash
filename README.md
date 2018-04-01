# hash
Hash is a simple tool for generating hashes from file or comparing a file with a supplied hash

        hash <algo> <file path> [hash]

The algos are:

        -sha256
        -sha512
        -md5

Examples:

        Generate sha256 hash from myFile
                 hash -sha256 c:\myFile

        Compares md5 hash from myFile with hash in argument
                 hash -md5 c:\myFile a44b49b6bc6248f7ee5b5f2626286983
