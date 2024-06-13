## Guess-it-1

This program that given a number as standard input,prints out a range in which the next number provided should be.

The algorithm used applys Kernel Distribution Estimation and Gaussian Kernel but not explicitly.Generally it uses the a simple method that assumes the data is of a normal distribution and that the number for which we are looking for the range is of the mean of data.

## installation

This project is written in Go.If you don't have Go installed,you can follow these steps:

### step 1
1. Download the Go installer from the official [Go website](https://go.dev/doc/install).
2. Follow the installation instructions for your specific operating system.
3. Verify your installation by opening a command prompt and typing `go version`. This should print the installed version of Go.
### step 2
clone into this repo by entering this into the terminal
````
git clone https://learn.zone01kisumu.ke/git/cliffootieno/guess-it-1
````
## Running the program
### step 1
Firstly enter a set of data in the data.txt file 
The format to write data is:
````
189
113
121
114
145
110
...
````

### step 2
to run the program enter this into the terminal
````
go run main.go
````
or
````
go run .
````
expected output :
````
189 --> the standard input
120 200    --> the range for the next input, in this case for the number 113
113 --> the standard input
160 230    --> the range for the next input, in this case for the number 121
121 --> the standard input
110 140    --> the range for the next input, in this case for the number 114
114 --> the standard input
100 200    --> the range for the next input, in this case for the number 145
145 --> the standard input
1 99      --> the range for the next input, in this case for the number 110
````