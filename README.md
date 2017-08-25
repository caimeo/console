# cameo/console
A simple configurable output console

Allows initialization with **verbose** and/or **debug** options.  

Use throughout your program code to output simple information to the console.

## Intro
While developing and debuging it is often useful to output debug information to the console that would not normally be visible to the user.  It is also often desired to include a verbose mode of operation for command line utilities.  This utility library provides a simple way to acomplish this as well as output to Standard Out and Standard Error.

### Messages types
- **Always** messages are always output to Standard Out   
- **Error** messages are always sent to Standard Error   
- **Debug** messages are only output to Standard Out if *IsDebug() == true*   
- **Verbose** messages are only output to Standard Out if *IsVerbose() == true* or *IsDebug() == true* 

### Redirect StdOut and StdErr
By default messages are sent to os.Stdout and os.Stderr.  If you desire to send messages to another location use **RedirectIO(io.writer, io.writer)** 

To connect to a logging system simply provide an io.writer that outputs to your logger.


## Usage Examples

### Instance
````
...
import "github.com/caimeo/console"
...
con := console.New(true, true)

con.Always("Message to User")
con.Debug("Some debugging info")
con.Error("OMG an error!")
...
````

### Static/Singleton
````
import "github.com/caimeo/console"
...
console.Init(true, true)
...
console.Debug("Some info Debug data", v1, v2)
console.Always("Message to User")
...

````
````
import "github.com/caimeo/console"
...
console.Debug("A message in another file"")
...

````
### Mixed
It is also possible to mix these methods.

````
...
import "github.com/caimeo/console"
...
con := console.Init(true, true)

con.Always("Message to User via an instance")
console.Always("Message to user via singleton")

````
or

````
...
import "github.com/caimeo/console"
...
console.Init(true, true)
console.Always("Message to user via singleton")

con := console.Instance()
con.Always("Message to User via an instance")

````


Pull requests are welcome.   

Caimeo Prime  
Caimeo Operating Group  
