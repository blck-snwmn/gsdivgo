Retrieve data from Google Spreadsheet and write it to a file.

## Getting Started
1. **Service Account**: To use this tool, you need to create a Service Account.
2. **Generate a Key**: Generate a key associated with the Service Account, and save it as credentials.json.
3. **Grant Permissions**: For the spreadsheets you want to read, grant read permissions to the Service Account you created.

## Run
```
$ go run main.go -s <your spread sheet id>
```
