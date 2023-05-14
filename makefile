.PHONY: all clean

# Set the executable name
EXE=my-program

all: $(EXE)

# Build the executable
$(EXE):
	go build -o $(EXE)

# Clean up the executable
clean:
	rm -f $(EXE)

# Run the program with the specified domain and output file
run:
	go run main.go -d example.com -o output.txt
