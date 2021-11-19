# fake-data-generator-api 
generate some random data

## installing and using - API

clone this repo:

```
$ git clone <repo url>
```

enter in the app folder and compile:
```
$ cd fake-data-generator-api
$ go build main.go
```

run the program:

```
$ ./main
```

This commands will execute the script in: 

```
http://127.0.0.1:8000
```

to stop the server(localhost) press <kbd>Ctrl</kbd><kbd>+</kbd><kbd>C</kbd> or <kbd>Cmd</kbd><kbd>+</kbd><kbd>C</kbd>

## installing and using - CLIENT

***Warning: to use the client you have to run the api at the sam time***

clone this repo:

```
$ git clone <repo url>
```

enter in the app folder and compile:
```
$ cd fake-data-generator-api/client
$ go build main.go
```

run the program:

```
$ ./main
```

## License

[MIT](LICENSE)