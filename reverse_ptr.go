//reverse_ptr

package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	//"bufio"

)
//var (
	//ipaddress = "10.20.30.40"; //array for in-addr.arpa.
	//inaddrarpa = [4]string{1,2,3,4};
	//ipv4addr = strings.Split(ipaddress, ".");
//)
func main() {

              //fmt.Println(ipv4addr[3] + "." + ipv4addr[2] + "." + ipv4addr[1] + "." + ipv4addr[0] + ".in-addr.arpa")

							b, err := ioutil.ReadFile("/Users/asewell/documents/infoblox/scripts/alf/test.csv") // just pass the file name
	 						if err != nil {
			 						fmt.Print(err)
	 						}

	 						//fmt.Println(b) // print the content as 'bytes'

	 						file := string(b) // convert file to a 'string'
							//for loop to the end of file
							//var line string


							ipaddress := strings.Split(file, "\r\n") //ipaddress = next line
							//for {} i = length of array ipaddress
							for i := 0; i < len(ipaddress); i++ {

							ipv4addr := strings.Split(ipaddress[i], "."); //ipv4addr = the ip split by "." ; i is the index of ipaddress
							//fmt.Println(ipv4addr) //testing if ipv4addr is first line of file
							fmt.Println(ipv4addr[3] + "." + ipv4addr[2] + "." + ipv4addr[1] + "." + ipv4addr[0] + ".in-addr.arpa") //print
							//in-addr.arpa address.
							}
	 						//fmt.Println(file) // print the file as a 'string'

}
