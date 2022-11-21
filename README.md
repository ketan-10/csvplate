# csvplate
csvplate (abbreviation for <b>csv to template</b>) is a 
CLI tool written in golang to read csv file and print it in specific format.

## Usage

### Prepare Inputs : 
- <b>Example csv file `data.csv`</b>

  | x  |  y Position  | value |
  |----|-----|-------|
  | 10 | -10 | ketan |
  | 20 | -20 | hello |
  | 30 | -30 | test  |

- <b>Example template file `template.html`</b>
```html
<g>
{{range .}}
    <text x='{{.x}}' y='{{index . "y Position"}}'>{{.value}}</text>
{{end}}
</g>
```
[Read more about golang text template](https://pkg.go.dev/text/template)

### Run cli :

- Installation 
    ```
    go install -a github.com/ketan-10/csvplate/cmd/csvplate@latest
    ```
- Run  
    ```
    csvplate ./data.csv ./template.html output.html
    ```
    Above query will generate `output.html` as following
    ```html
    <g>
        <text x='10' y='-10'>ketan</text>
        <text x='20' y='-20'>hello</text>
        <text x='30' y='-30'>test</text>
    </g>
    ``` 
### Other features
- If you have large amount of data you can generate multiple file by splitting it
    ```
    csvplate ./data.csv ./template.html output.html 100
    ```
    This will create multiple files each file with 100 records, with names `output-1.html`, `output-2.html`, ... etc 