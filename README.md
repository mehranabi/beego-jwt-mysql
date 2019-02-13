## Beego + JWT + MySQL
A sample Golang project that shows you how to create your APIs with Beego, also JWT (Json Web Token) is implemented.
In this project I used MySQL database to store Users data.

## Installation & Running
You should first install Golang and register `GOPATH` in your *System Environment Valriables*, read more in https://golang.org/doc/install.

Then install _Beego_ and _Bee Tool_, read more in https://beego.me/quickstart

You need to install some pakcages that are used in this project, install them with: (I imagine that you've installed **Beego** & **Bee Tool**)
 - `go get github.com/astaxie/beego/orm` (_Beego ORM Helper_)
 - `go get github.com/go-sql-driver/mysql` (_MySQL Driver for Golang_)
 - `go get golang.org/x/crypto/bcrypt` (_Bcrypt helper to hash passwords_)
 - `go get github.com/gbrlsnchs/jwt` (_JWT Helper_)
 - `go get github.com/SermoDigital/jose` (_also JWT Helper **but I just used rsa keys loader from that**_)

After preparing your environment, clone this repository in `%GOPATH%\src\beego_jwt_mysql`.

* For signing Json-Web-Tokens I used RSA public/private key pair, you must create `keys` folder in the root of project and then create two `private.txt` and `public.txt` files. Then put your Private RSA key text in `public.txt` and put your Public RSA key text in `private.txt`.
  * You can create RSA key pair by this commands: (You must have installed __openssl__ on your system)
    * Private Key: `openssl genrsa -out private.txt 2048`
    * Public Key: `openssl rsa -in private.txt -pubout > public.txt`

Then you can run the project by this command:
 * `bee run`
Your app will run on :8080 port, You can visit site on http://localhost:8080
 
If you want _Beego_ to generate API documentation automatically for you using _Swagger_ definition you can use this command instead:
 * `bee run -downdoc=true -gendoc=true`
Then you can visit http://localhost:8080/swagger to view documentation.

Use [Postman](https://getpostman.com) or any other tool that you're comfortable with to test your API

## Developer
* Mehran Abghari
  * Email: mehran.ab80@gmail.com
  * Github: https://github.com/mehranabi

LICENSE: MIT

You can use this project for any purpose.

__But if you like it, Let me know :) - Star It _or_ Just send me a message!__
