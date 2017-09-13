# `csvq` - Simple tool for querying CSV files

I love [`jq`](https://stedolan.github.io/jq/) and I wanted to do something similar for CSV data. So here it is.

**I AM GOING TO REWRITE MUCH OF THIS SO DON'T RELY ON IT FOR NOW**

## Installation
Install the latest version:

```
go install -u github.com/inkel/csvq
```

## Usage
```
Usage of csvq:
  -H	First line are headers (default true)
  -h string
    	Column to filter
  -v string
    	Value
```

Using the example data in [`encoding/csv`](https://golang.org/pkg/encoding/csv/) package documentation:

```
first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
``` 

### Do nothing
```
$ csvq file.csv
first_name,last_name,username
Rob,Pike,rob
Ken,Thompson,ken
Robert,Griesemer,gri
```

### Filter by named column
```
$ csvq -h username -v ken file.csv
first_name,last_name,username
Ken,Thompson,ken
```

Of course, if you pass `-H=false`, meaning the CSV file doesn't have headers, then this will fail with something like:

```
$ csvq -H=false -h first_name -v Rob gophers.csv
strconv.Atoi: parsing "first_name": invalid syntax
```

### Filter by column number
```
$ csvq -h 3 -v johndoe file.csv
first_name,last_name,username
Ken,Thompson,ken
```

This one works whether you enable headers or not.

### Select columns to output
TODO

### Change columns order
TODO

### Change output delimiter
TODO

### Change to JSON
TODO

## License
See [LICENSE](LICENSE).