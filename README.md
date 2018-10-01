# Salesforce ANT Project Generator

Because I work with many different client orgs, I found it annoying to have to create a new folder and copy/paste my template files for a new Force.com ANT project. My solution is this simple Go application to generate a project folder and populate it with the required build files. The app automatically generates a new project directory, a `src` directory inside and a few template files: `package.xml`, `build.properties`, and `build.xml`. Once the project folder is generated and the files created, it runs a quick command to `cd` into the new folder and open VSCode. 

> NOTE: for it to open VSCode afterwards I've [added VSCode to my path](https://code.visualstudio.com/docs/editor/command-line) on my Mac in order to run the command `code .` in the new directory. You may have to do this, or perhaps run your own command after file generation, it's up to you and your own workflow.

### Use
```sh
$ go get -u github.com/meruff/go-sf-ant
```
Once you run the app it'll ask you what you'd like to name the new directory. Upon submitting the name it'll create the files, `cd` in, and open the project in VSCode.

### Configuration
`main.go` includes a constant named `GOSFANTPROJECTPATH` that points to an environment variable where you designate the directory you'd like the ANT project to be generated in each time, i.e: 
```sh
$ export GOSFANTPROJECTPATH=/Users/username/Documents/ANT/
``` 

You can also manually designate a one-off directory path for the new ANT project by setting a `-d` flag value upon executing the command: 
```sh
$ go-sf-ant -d /Users/username/Documents/ANT/
```

If you don't have an environment variable set up, and don't designate a flag value, then go-sf-ant will just create the project in your current working directory.

I personally use [Alfred](https://www.alfredapp.com/) to run the app so all I have to do is type `cmd + space` to open Alfred and then `open go-sf-ant`, which can be run from anywhere in the OS. 
