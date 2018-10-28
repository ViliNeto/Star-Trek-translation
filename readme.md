
## Star Trek challenge

  

My task was to translate a name written in English to Klingon and find out its *species* using ​[http://stapi.co](http://stapi.co)​.

This code was developed using **go1.11.1 windows/amd64**

I followed the instructions and started looking forward the main *API's* to complete this task.

As I started studion the **stapi** documentation I noticed that in order to get the Star Trek character specie I would need to check two different routes.

First route:

***(POST) /v1/rest/character/search***

In the above route I would get some generic information about the character, but I won't get the character specie.

To do so, I would need to check the route:

***(GET) /v1/rest/character/?uid=%s***

Where ***%s*** will be the first route response for UID (could be a list).

For this project I split the project in a few packages.

**main**  
 - Where the main project is executed

**structures**
 - Where the Json parser structures are placed in.

**util**
 - Where the converters and string maps are.

**stapi**
 - Where the call to ***STAPI*** are made .

I also made the Unit test.
Which one are named ***_test.go**

	  
**How to run the project:**

 1. Make sure that Golang is installed =)

 2. In the root folder type the following command: **go run main.go Uhura** 

There is a build folder where you can use the application without installing the Golang, but I recommend to use the script above!

**Output example:**

$ **go run main.go Uhura**

0xF8E5, 0x​F8D6, ​​​0xF8E5, ​​0xF8E1, 0xF8D0

Human
