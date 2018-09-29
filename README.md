# Salesforce ANT Project Generator

Because I work with many different client orgs, I found it annoying to have to create a new folder and copy/paste my template files for a new Force.com ANT project. My solution is this simple Go application to generate a project folder and populate it with the required build files. The app automatically generates a new project directory, a `src` directory inside and a few template files: `package.xml`, `build.properties`, and `build.xml`. Once the project folder is generated and the files created, it runs a quick command to `cd` into the new folder and open VSCode. 

> NOTE: for it to open VSCode afterwards I've [added VSCode to my path](https://code.visualstudio.com/docs/editor/command-line) on my Mac in order to run the command `code .` in the new directory. You may have to do this, or perhaps run your own command after file generation, it's up to you and your own workflow.

### Use
```sh
$ go get -u github.com/meruff/go-sf-ant
```
`main.go` includes a constant named `projectPath` where you designate the directory you'd like the ANT project to be generated in, i.e: `'/Users/username/Documents/ANT'`.

I personally use [Alfred](https://www.alfredapp.com/) to run the app so all I have to do is type `cmd + space` to open Alfred and then `open go-sf-ant`, which can be run from anywhere in the OS. 

Once you run the app it'll ask you what you'd like to name the new directory. Upon submitting the name it'll create the files, `cd` in, and open the project in VSCode.
