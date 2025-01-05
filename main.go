package main

import (
	flowcontrol "golang/basics/flowControl" //I have declear  to import alices is because the fileName didn't match the packageName if we want to remove it just rename golang/basics/flowcontrol as this matches the packagesName
)

func main() {
	// components.Hello("Ritesh Khadka")
	flowcontrol.ContinuedLoop()

}
