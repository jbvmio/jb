#!/bin/bash

GOOS=darwin ARCH=amd64 go build -ldflags "-X main.buildTime=$(date +"%Y.%m.%d.%H%M%S") -X main.commitHash=$(git rev-list -1 HEAD)" -o darwin_main.bin main.go
GOOS=linux ARCH=amd64 go build -ldflags "-X main.buildTime=$(date +"%Y.%m.%d.%H%M%S") -X main.commitHash=$(git rev-list -1 HEAD)" -o linux_main.bin main.go


#     var (
#      	showVer  = flag.Bool("v", false, "Show Version Info")
#    
#      	buildTime  string
#      	commitHash string
#     )
#    
#     func main() {
#      	flag.Parse()
#      	if *showVer {
#      		fmt.Printf("Version   : %s\n", buildTime)
#      		fmt.Printf("Commit    : %s\n", commitHash)
#      		return
#       }
#     }

