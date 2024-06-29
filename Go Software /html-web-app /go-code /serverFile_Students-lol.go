# Students-lol ~

package main


// .
  
import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
    
    
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}






func app_welcome_center_page() {
    
    
}

