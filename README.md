#riemann-collector

####Note: This program is not finished yet, it only has the most basic functionality and very little error recovery.

###About
riemann-collector is a simple, lightweight, and small program that reads
a configuration file that defines both riemann servers, as well as events.

The events currently are composed of a command and an interval, in seconds.
The command can be either a shell script like the one in the examples directory,
or any other command you can execute via `sh -c`. Currently commands are executed
from riemann-collectors working directory, but will likely change as event configurations
get more options added to them. It's expected that any program executed will
write JSON to stdout. This JSON can contain any field that is part of
a valid Riemann event.


###More to Come
* Proper error recovery
* More documentation on behavior
* More documentation on input
* More Event options
    - Working directory
    - More output formats for programs executed, like JSON arrays.
    - More output methods, maybe pipes. A way for other programs
        to talk to it would be nice. Maybe sockets too.
    - More as I think of them....
* More server options
    - Protocol
    - Blocking channel (currently doesn't if channel gets filled)
    - channel size
    - and more!!
* More configuration options
    - include directories and files
    - logging
