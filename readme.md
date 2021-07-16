

# Assignment

## Description
Exercise: Write a simple fizz-buzz REST server.

"The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"".
The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...""."

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

 They can share either the URL of a git repository with their code or send a zip file. They need to write the code in Golang and develop it as it was meant for production use.


## Implementation 
The code is split into 3 files, a main files that contains the server, a fizzbuzz file containing the fizzbuzz logic and a config file loading the configuration values. Default port will be set on :8010 


## Setup
Next, the scripts to start the general project locally.

Clone the project
```
$ git clone https://github.com/fMercury/fizzbuzz_go_sample.git

$ cd FizzBuzz

$ go run main.go fizzbuzz.go config.go

$ go build main.go fizzbuzz.go config.go
```

## Endpoints

### GET /fizzbuzz

**Parameters**

|          Name | Required |  Type   | Description                                                                                                                                                           |
| -------------:|:--------:|:-------:| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|     `int1` | required | int  |  <br/><br/> Default is `3`.                                                                     |
|     `int2` | required | int  |  <br/><br/> Default is `5`.                                                                     |
|     `limit` | required | int  | <br/><br/> Default is `100`.                                                                     |
|     `str1` | required | string  | <br/><br/> Default is `fizz`.                                                                     |
|     `str2` | optional | string  | <br/><br/> Default is `buzz`. <br/><br/>                                                                      |

**sample local call**
```
http://localhost:8010/fizzbuzz?int1=3&int2=5&limit=20&str1=fizz&str2=buzz
```

### GET /statistics
**sample local call**
```
http://localhost:8010/statistics
```

