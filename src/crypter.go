package main

import (
    "crypto/cipher"
    "crypto/aes"
    "crypto/md5"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "flag"
    "log"
    "io"
    "io/ioutil"
	"strings"
)

func parse_cli() ( *string, *string, *string ) {
    
    inputFile := flag.String( "i", "<input>", "file to crypt" )
    inputPwd  := flag.String( "p", "<password>", "password for encrypt" )
	inputTmp  := flag.String( "t", "<template>", "template file" )

    flag.Parse()
    
    if flag.NFlag() < 3 {
        log.Fatal("[ERROR] use -h to help")
    }

    return inputFile, inputPwd, inputTmp
}

func saveCrypted( szTemplate string, strBase64 string ) {
	
	strTmp, err := ioutil.ReadFile( szTemplate )
	if err != nil {
        log.Fatal(err)
    }
    
    if strings.Contains( string( strTmp ), "#CRYPTED#" ) {
		
		// Replace TAG #CRYPTED# with base64 crypted string
		strFinal := strings.Replace( string( strTmp ), "#CRYPTED#", strBase64, 1 )
		
		ioutil.WriteFile( "out.html", []byte(strFinal), 0644 )
	}
}


func main( ) {
    szFile, szPasswd, szTemplate := parse_cli()

    strClean, err := ioutil.ReadFile( *szFile )
    if err != nil {
        log.Fatal(err)
    }

    h := md5.New()
    io.WriteString( h, *szPasswd )

    hashPwd := h.Sum(nil)

    
    finalPwd := []byte( hex.EncodeToString(hashPwd) )
    iv := finalPwd[:aes.BlockSize]
    
    //fmt.Println( "aes.BlockSize : ", aes.BlockSize, "iv : ", iv )

    block, err := aes.NewCipher( hashPwd )
	if err != nil {
		log.Fatal(err)
	}
    
    encrypted := make( []byte, len(strClean) )
    aesEncrypter := cipher.NewCFBEncrypter( block, iv )
	aesEncrypter.XORKeyStream( encrypted, strClean )
     
    strBase64 := base64.StdEncoding.EncodeToString( encrypted )

    //fmt.Printf( "ENCODED (b64): %s\n", strBase64 )
	
	saveCrypted( *szTemplate, strBase64 )
	
	fmt.Println( "Template filled!" )
}
