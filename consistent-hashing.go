package main

import (
	"fmt"
)

type node struct{
	nameId int
	prevIndex int
	index int
	nextIndex int
	dict map[int]string

type ring struct{
	hashSize int
	serverNodes []node
}

type (r ring) addServer() (err error) {
	//find server with the largest amount of items (targetServer)
	//this is the server we want to break up
	
	//add new server between targetServer and the next server (logn time)

	//rehash everything in targetServer
	////this should repopulate targetServer and the newly added server

	return err
}

type (r ring) removeServer(ringIndex int) (err error) {
	//move everything in the server to be removed (rServer) to the previous server
	
	//remove rServer

	return err
}

//hash key and return server in which will/should contain the data.
type (r ring) hash(key string) (int serverIndex, err error) {
	//hash the key using some hashing function. ex. Md5, SHA-256, etc.

	//turn the hash into a number

	//get exact index at which the key/value should be stored at

	//find the server index closest to and <= to the exact index

	return serverIndex, err
}

type (r ring) get(key string) (value string, err error) {
	//hash the key

	//use hash to get the server its located on

	//query the server for the data

	//return value, err

	return value, err
}

type(r ring) put(key string, value string) (err error) {
	//hash the the key

	//find which server on the ring to store the value

	//input the key value on the server

	return err
}

type(r ring) del(key string) (err error) {
	//hash the key, get server it data is on (tServer)

	//delete the data from tServer

	return err
}
//could be modified to take a user input for initial amount of servers
// for a more even initial ring division
type initRing(int serverNum) (r ring) {
	//set default ring hashSize
	//divide the ring up into even sections of serverNum

	//return the newly created ring.
	return 
}


func main() {

}
