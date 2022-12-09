https://stackoverflow.com/questions/66691135/postgresql-initdb-segmentation-fault-core-dumped

 did some changes of pgsql source code, such as adding a new system catalog. Then I wanted to init the database by using initdb -D PG_DATA_PATH. But it came to this problem:

The files belonging to this database system will be owned by user "pguser".
This user must also own the server process.

The database cluster will be initialized with locale "en_US.UTF-8".
The default database encoding has accordingly been set to "UTF8".
The default text search configuration will be set to "english".

Data page checksums are disabled.

creating directory /home/pguser/Desktop/postgres/data ... ok
creating subdirectories ... ok
selecting dynamic shared memory implementation ... posix
selecting default max_connections ... 100
selecting default shared_buffers ... 128MB
selecting default time zone ... America/Los_Angeles
creating configuration files ... ok
running bootstrap script ... ok
performing post-bootstrap initialization ... Segmentation fault (core dumped)
child process exited with exit code 139
initdb: removing data directory "/home/pguser/Desktop/postgres/data"
There is NO errors or warnings when I executed make and make install. Also, I used gdb to backtrace the execution of initdb:


You have a core dump from the wrong process.

initdb starts the PostgreSQL server in stand-alone mode and pipes the catalog contents to it. Now something in the data you are writing is bad, and the stand alone backend exits. Then initdb gets a SIGPIPE if it writes to the pipe and crashes.

You want to debug the PostgreSQL server, not initdb


Thank you for your remind. So you mean maybe system catalog's metadata I'm writing is bad? Or potentially anything else is bad? BTW, how to debug PostgreSQL server when I'm running 'initdb'? A little confused about it. – 
Charlie_Y
 Mar 18 at 15:25
I'd say that you feed PostgreSQL something bad. Build a delay into PostgreSQL startup and attach to the standalone backend. – 
Laurenz Albe
 Mar 18 at 16:01
Solved it! It's bug from backend process. Thank you anyway! – 
Charlie_Y
 Mar 19 at 15:26
 