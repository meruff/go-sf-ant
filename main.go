package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// projectPath is the literal string path where the template files will be created.
// Example: "/Users/username/Documents/"
const projectPath = "/Users/username/Documents/"

var packageXMLString = `<?xml version=\"1.0\" encoding=\"UTF-8\"?>
<Package xmlns=\"http://soap.sforce.com/2006/04/metadata\">
	<types>
		<members>*</members>
		<name>ApexClass</name>
	</types>
	<types>
		<members>*</members>
		<name>ApexComponent</name>
	</types>
	<types>
		<members>*</members>
		<name>ApexPage</name>
	</types>
	<types>
		<members>*</members>
		<name>ApexTrigger</name>
	</types>
	<types>
		<members>*</members>
		<name>AuraDefinitionBundle</name>
	</types>
	<types>
		<members>*</members>
		<name>StaticResource</name>
	</types>
	<version>43.0</version>
</Package>`

var buildXMLString = `<project name="antProject" default="main" basedir="." xmlns:sf="antlib:com.salesforce">
    <property file="build.properties"/>
    <property environment="env"/>

    <!-- Setting default value for username, password and session id properties to empty string
         so unset values are treated as empty. Without this, ant expressions such as ${sf.username}
         will be treated literally. -->
    <condition property="sf.username" value=""> <not> <isset property="sf.username"/> </not> </condition>
    <condition property="sf.password" value=""> <not> <isset property="sf.password"/> </not> </condition>
    <condition property="sf.sessionId" value=""> <not> <isset property="sf.sessionId"/> </not> </condition>
    <condition property="sfdeploy.username" value=""> <not> <isset property="sfdeploy.username"/> </not> </condition>
    <condition property="sfdeploy.password" value=""> <not> <isset property="sfdeploy.password"/> </not> </condition>
    <condition property="sfdeploy.sessionId" value=""> <not> <isset property="sfdeploy.sessionId"/> </not> </condition>

    <taskdef resource="com/salesforce/antlib.xml" uri="antlib:com.salesforce">
        <classpath>
            <pathelement location="../ant-salesforce.jar" />
        </classpath>
    </taskdef>

    <target name="retrieve">
        <sf:retrieve
            username="${sf.username}"
            password="${sf.password}"
            serverurl="${sf.serverurl}"
            maxPoll="${sf.maxPoll}"
            retrieveTarget="src"
            unpackaged="src/package.xml"/>
    </target>

    <target name="deploy">
        <sf:deploy
            username="${sfdeploy.username}"
            password="${sfdeploy.password}"
            serverurl="${sfdeploy.serverurl}"
            checkOnly="true"
            maxPoll="${sfdeploy.maxPoll}"
            deployRoot="src"
            rollbackOnError="true">
        </sf:deploy>
    </target>
</project>`

var buildPropertiesString = `# Retrieve from Org:
# Specify the login credentials for the desired Salesforce organization
sf.username = 
sf.password = 
#sf.sessionId = <Insert your Salesforce session id here.  Use this or username/password above.  Cannot use both>
#sf.pkgName = <Insert comma separated package names to be retrieved>
#sf.zipFile = <Insert path of the zipfile to be retrieved>
#sf.metadataType = <Insert metadata type name for which listMetadata or bulkRetrieve operations are to be performed>

# Use 'https://login.salesforce.com' for production or developer edition (the default if not specified).
# Use 'https://test.salesforce.com for sandbox.
sf.serverurl = https://test.salesforce.com

sf.maxPoll = 20
# If your network requires an HTTP proxy, see http://ant.apache.org/manual/proxy.html for configuration.

# Deploy to Org:
sfdeploy.username = 
sfdeploy.password = 
sfdeploy.serverurl = https://login.salesforce.com
sfdeploy.maxPoll = 20`

func main() {
	fmt.Println("What would you like to name the new ANT project folder?")
	projectFolderName, err := bufio.NewReader(os.Stdin).ReadString('\n')
	check(err)
	projectFolderName = strings.Replace(projectFolderName, "\n", "", -1)

	// Create empty directories for proejct
	check(os.MkdirAll(projectPath+projectFolderName, os.ModePerm))
	check(os.MkdirAll(projectPath+projectFolderName+"/src/", os.ModePerm))

	// Create build files and package.xml from temp
	createFile(projectPath+projectFolderName+"/src/package.xml", packageXMLString)
	createFile(projectPath+projectFolderName+"/build.xml", buildXMLString)
	createFile(projectPath+projectFolderName+"/build.properties", buildPropertiesString)

	cmd := exec.Command("code", ".")
	cmd.Dir = projectPath + projectFolderName
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func createFile(fullNamedPath string, fileString string) {
	file, err := os.Create(fullNamedPath)
	if err != nil {
		errors.Wrap(err, "Cannot create file")
	}
	defer file.Close()

	fmt.Fprintf(file, fileString)
}
