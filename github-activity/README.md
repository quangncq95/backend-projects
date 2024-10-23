# github-activity
This is cli program to show public activity of github user.

The requirement follows task describe in this link :
https://roadmap.sh/projects/github-user-activity

### Programming language
I use [Golang](https://go.dev/) for this project.

### Installation
1. Install golang from this link (select version match your system) : https://go.dev/dl/
2. Build the source code

**For Windows 64-bit**
`GOOS=windows GOARCH=amd64 go build -o github-activity.exe ./path/to/package`

**For macOS 64-bit**
`GOOS=darwin GOARCH=amd64 go build -o github-activity ./path/to/package`

**For Linux 64-bit**
`GOOS=linux GOARCH=amd64 go build -o github-activity ./path/to/package`

3. Add execution file to environment variable

### Usage 
` github-activity <username>`

