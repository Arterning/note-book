CC=gcc
CFLAGS = -O3

all: socket-client.o test.o
	$(CC) $(LFLAGS) -o $@ $?

%.o: %.c
	$(CC) $(CFLAGS) -o $@ -c $?

clean:
	rm -f *.o
