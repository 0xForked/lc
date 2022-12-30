## Log File Converter Example
File converter example implementation with JS and GO, Requested by some people .... 

### JavaScript Implementation
#### RUN
    `node index.js -f /path/to/file/example.log -t text`
### Golang Implementation
#### RUN
    `go run main.go -f /path/to/file/example.log -t text`

### Available transform format
- JSON
  - usage `-t json`, 
  - will save to `.json` format
- TEXT
  - usage `-t text`, 
  - will save to `.txt` format

### Available flags
- [`-f`, `--f`, `--file`, `file`] File path to format
- [`-t`, `--t`, `--transform`, `transform`] Transform/Convert selected file. 
  - default (`json`), 
  - available: (`json`, `text`)
- [`-o`, `--o`, `--output`, `output`] File target output directory
  - default (`./storage`)