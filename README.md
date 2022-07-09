# jwt_payloads (Go)
*jwt_payloads generates a list of modified tokens from a single JWT token. It uses known exploit methods like null signature, JKU injection, and kid injection amongst others.* (the [wiki page](https://github.com/ram8537/jwt_payloads/wiki/Exploits) summarises all the exploits)


![new logo 2](https://user-images.githubusercontent.com/67279424/176843572-609ef489-f504-475c-9b5f-2c9dbbc4cc9e.png)
 ![enter image description here](https://img.shields.io/badge/go-1.18-brightgreen)
## Summary
jwt_payloads is based on the popular [jwt_tool](https://github.com/ticarpi/jwt_tool) (python).

⭐️ jwt_payloads (Go), **generates a list of payloads (modified JWT tokens) which can then be sent to the target application using tools like Burp Suite Intruder or Wfuzz/FuFF.** 

## Installation

First, you'll need to [install go](https://go.dev/doc/install)

Then run this command to download + compile jwt_payloads:
```
go install github.com/ram8537/jwt_payloads@latest
```
You can now run `~/go/bin/jwt_payloads`. If you'd like to just run `jwt_payloads` without the full path, you'll need to `export PATH="~/go/bin/:$PATH"`. 

## Usage

You must specify a token and URL (to test for out-of-band interactions). The payloads are outputted to the terminal, and can be saved to a file with the `>` operator like in the example below. To test for out of band interactions, you can get a url from here: https://app.interactsh.com/. 
 
    ~/go/bin/jwt_payloads -token <token> -url <url> > output.txt

### Additional arguments (crack HSxxx tokens)
If you're token is symmetrically signed (alg: HSxxx), you can specify an optional '-crack' flag with a **full path** to a file of common passwords. You can find a list [here](https://github.com/wallarm/jwt-secrets/blob/master/jwt.secrets.list).

    ~/go/bin/jwt_payloads -token <token> -url <url> -crack <filepath> > output.txt

## Suggested Workflow

 - Run jwt_payloads and save the output into an 'output.txt'
 - Copy all the modified tokens and paste in Burp Intruder
  - Disable the payload encoding under the 'payloads' tab of Burp Intruder 
 
 <img width="690" alt="disable url encoding" src="https://user-images.githubusercontent.com/67279424/176845848-b0c96d92-6be8-4c1b-b83b-6c100b18dc78.png">

 - Analyse results:
	 - Error messages
	 - Unusually long/short response time
	 - Byte length of response
	 - Different status codes
	 -  Differences between working/accepted token vs modified token
 - You can refer to the [wiki page](https://github.com/ram8537/jwt_payloads/wiki/Exploits) to see how the exploit tokens are generated.