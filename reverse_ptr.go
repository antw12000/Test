//Create reverse_ptr from ipaddress/file

package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {

							b, err := ioutil.ReadFile("/your/directory") // just pass the file name
	 						if err != nil {
			 						fmt.Print(err)
	 						}
							//fmt.Println(b) // print the content as 'bytes'

	 						file := string(b) // convert file to a 'string'

							ipaddress := strings.Split(file, "\r\n") //ipaddress is an array with the delimiter being the next line

							for i := 0; i < len(ipaddress); i++ { //for i != length of the array ipaddress loop

									//ipv4addr is an array of the ipaddress[i] split by "." ;
									ipv4addr := strings.Split(ipaddress[i], ".");
									//print array for ipv4addr in revervse order and add .in-addr.arpa
									fmt.Println(ipv4addr[3] + "." + ipv4addr[2] + "." + ipv4addr[1] + "." + ipv4addr[0] + ".in-addr.arpa")

							} //end for loop length of array

} //end main
