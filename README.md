# csv-to-ldif

This tool creates a ldif file from a csv import file

## todo:

- changetype, maybe first column in csv, maybe by file name (add.csv, modify.csv etc)
- generate dn by uid,ou,dc
- support multiple ou,dc
- docker image

## ideas:

- support different file formats (json, ini?)
- multi threaded with x goroutines writing to thread-specific string, appending strings at the end (needs performance testing to see if worth pursing)
-  generate input file for performance testing

