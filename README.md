# Platform Science Code Exercise

Our sales team has just struck a deal with Acme Inc to become the exclusive provider for routing their product shipments via 3rd party trucking fleets. The catch is that we can only route one shipment to one driver per day.

Each day we get the list of shipment destinations that are available for us to offer to drivers in our network. Fortunately our team of highly trained data scientists have developed a mathematical model for determining which drivers are best suited to deliver each shipment.

With that hard work done, now all we have to do is implement a program that assigns each shipment destination to a given driver while
maximizing the total suitability of all shipments to all drivers.


The top-secret algorithm is:


```
- If the length of the shipment's destination street name is even, the base suitability score (SS) is the number of vowels in the driver’s name multiplied by 1.5.

- If the length of the shipment's destination street name is odd, the base SS is the number of consonants in the driver’s name multiplied by 1.

- If the length of the shipment's destination street name shares any common factors (besides 1) with the length of the driver’s name, the SS is increased by 50% above the base SS.
```


Write an application in the language of your choice that assigns shipment destinations to drivers in a way that maximizes the total SS over the set of drivers. Each driver can only have one shipment and each shipment can only be offered to one driver. Your program should run on the command line and take as input two newline separated files, the first containing the street addresses of the shipment destinations and the second containing the names of the drivers. The output should be the total SS and a matching between shipment destinations and drivers. You do not need to worry about malformed input, but you should certainly handle both upper and lower case names.

## Documentation

```bash
$ go install golang.org/x/tools/cmd/godoc@latest

$ godoc -http=:6060
```

Once the godoc server is running we can enter:

*http://localhost:6060/pkg/github.com/shotokan/platform-science-code-exercise/*


## Installation

>> Running with GO

```bash
# it's necessary to install GO (^1.18.1) and dependencies
$ go mod download 
$ go run ./cmd/main.go 

# Example:

Path to drivers file: /home/isabido/projects/platform-science-code-exercise/data/streets.txt
Path to destinations file: /home/isabido/projects/platform-science-code-exercise/data/drivers.txt
```


**In order to read the file, you must pass the absolute path* 


>> Tests

Run tests

```bash
# it's necessary to install GO (^1.18.1) and dependencies
$ go mod download 
$ go test -v ./...
```

## Solution

The CalculateSS function runs through the destinations and for each destination verifies if its length is even or odd. If it is even, the drivers list is iterated and for each driver the number of vowels in the driver's name is searched and it is multiplied by 1.5 and that quantity is taken as the base of SS. The next step is to verify if the length of the destination name has a common factor greater than 1, taking the length of the driver's name as the second number, and if there is one, the SS base is increased to 50%. The same is true if the length of the destination name is odd, but in that case the constants are searched and the common factor is checked as indicated. It is necessary to go through the entire list of drivers to find the one with the highest SS in order to assign the destination.


** Comments clarify the code and they are added with purpose of making the code easier to understand. 