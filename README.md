GoLang-CEPSearch
GO code designed for conducting an address search based on the specified ZIP code.

Key features include:

- JSON (Marshal and Unmarshal) for efficient data handling
- HTTP requests for seamless communication
- Defer for resource management
- File manipulation for enhanced functionality
- Structs



###  Usage:

Example:
```bash
go run main.go http://viacep.com.br/ws/01001000/json/
```

Replace "01001000" with the actual numeric ZIP code you want to query.

In this version, I added the functionality to create a text file containing the data returned from the query. 
