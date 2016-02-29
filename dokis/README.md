# Dokis

Dokis generate a simple index html page containing the documentation of all of your project using sagger yaml or json files.

## Prerequisite
- [bootprint swagger](https://github.com/nknapp/bootprint-swagger)

## How to Use
```bash
git clone https://github.com/evonck/doki.git
cd doki/dokis
go build
```
## Use
```bash
 ./dokis pathToTheYamlFiles ouptutPath
```
Dockis will create json files from the yaml and use bootprint-swagger to change them to html documentation.


## Required
To run doki locally you need to install bootprint swagger on your machine:

https://www.npmjs.com/package/bootprint

## Command Line
  - --force, -f
   force the creation of the json file even if another json file of the same name already exist.
 - --recursive, -r
    Make a recursive search for all yaml file in the directory
- --index, -i
    Path to the index.html template file.

Example:
Folder project:
```bash
project
    |---test
        |---v0
            |--- test.yaml
```

Run:
```bash
./yamlToSwagger ./test.yaml ./docs
```

Generate:
folder docs:
```bash
docs
    |---test
        |---test.html
    |---index.html
    |---main.css
    |---nav.css
    |---template.html
```


![alt tag](https://raw.github.com/evonck/doki/master/img/test.png)

