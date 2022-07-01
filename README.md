


# jwt_payloads (Go)
*jwt_payloads generates a list of modified tokens from a single JWT token. It uses known exploit methods like null signature, JKU injection, and kid injection amongst others.* (the wiki page summarises all the exploits)


![new logo 2](https://user-images.githubusercontent.com/67279424/176843572-609ef489-f504-475c-9b5f-2c9dbbc4cc9e.png)


![enter image description here](https://img.shields.io/badge/version-1.0.0-blue) ![enter image description here](https://img.shields.io/badge/go-1.18-brightgreen)
## Summary
jwt_payloads is based on the popular [jwt_tool](https://github.com/ticarpi/jwt_tool) (python), with **one key difference**:


*jwt_tool (python)* sends the modified tokens directly to the target URL and provides basic information about the responses including: (1) response code, (2) byte length (3) whether a user-specified canary value was found in the response.

⭐️ In contrast, jwt_payloads (Go), **generates a list of payloads which can be used with Burp Suite intruder, or any similar tool (e.g. Postman).** These GUI tools allow you to inspect the results more closely, for example, analysing the actual response, its length in bytes, and regex searches, amongst others.


## Installation

First, you'll need to [install go](https://go.dev/doc/install)

Then run this command to download + compile jwt_payloads:
```
go install github.com/ram8537/jwt_payloads@latest
```
You can now run `~/go/bin/jwt_payloads`. If you'd like to just run `jwt_payloads` without the full path, you'll need to `export PATH="~/go/bin/:$PATH"`. 

## Usage

You must specify a token and URL (to test for out-of-band interactions). I recommend sending the output a file "output.txt" to visualise the results easily. To test for out of band interactions, you can get a url from here: https://app.interactsh.com/. You can use any URL (like `test.com`) to try the tool out first before you use interactsh when testing against a real target.
 
`~/go/bin/jwt_payloads -token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMzIzMzQiLCJuYW1lIjoiVGF5bG9yIiwiaWF0IjoxNTE2MjM5MDIyfQ.hyuSywEQ2r65Y6goXYXlmU44_KFP_moZ9N4JT_E_meY -url test.com > output.txt`  

### Additional arguments (crack HSxxx tokens)
If you're token is symmetrically signed (alg: HSxxx), you can specify an optional '-crack' flag with a **full path** to a file of common passwords. You can find a list [here](https://github.com/wallarm/jwt-secrets/blob/master/jwt.secrets.list).
`~/go/bin/jwt_payloads -token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMzIzMzQiLCJuYW1lIjoiVGF5bG9yIiwiaWF0IjoxNTE2MjM5MDIyfQ.hyuSywEQ2r65Y6goXYXlmU44_KFP_moZ9N4JT_E_meY -url hello.com -crack /Users/yourusername/Desktop/common-jwt-secrets.txt > output.txt`  

## Suggested Workflow

 - Run jwt_payloads and save the output into an 'output.txt' (if you copied the command above, this would be done already)
 - Copy all the modified tokens (in the second part of the file, you can ctrl+f 'full output' to find it)
  ![output second part](https://user-images.githubusercontent.com/67279424/176845552-f571ee76-afb8-4f2a-88eb-6d639ae01c51.png)
 - Paste into Burp Intruder (you might want to disable the payload encoding under the 'payloads' tab of burp intruder if it messes with the token)
 
 <img width="690" alt="disable url encoding" src="https://user-images.githubusercontent.com/67279424/176845848-b0c96d92-6be8-4c1b-b83b-6c100b18dc78.png">

 - Analyse results:
	 - Error messages
	 - Unusually long/short response time
	 - Byte length of response
	 - Different status codes
	 -  Differences between working/accepted token vs modified token
 - If you spot anything interesting, search for the token in your output.txt file to find what exploit was used. You can refer to the wiki page to see how the exploit works.
