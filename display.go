package main

import "fmt"

const ayaneLogo = `
    ___                            
   /   |  __  __ ____ _ ____   ___ 
  / /| | / / / // __  // __ \ / _ \
 / ___ |/ /_/ // /_/ // / / //  __/
/_/  |_|\__, / \__,_//_/ /_/ \___/ 
       /____/                      
`

const appVersion = "v1.0.0"

func Display() {
	fmt.Println(ayaneLogo)
	fmt.Printf("=> Ayane CLI Todo Application %s\n\n", appVersion)
}
