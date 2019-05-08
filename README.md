GoPrims
========

## What is it?

This is an implementation of Prim's algorithm in Go specifically to navigate a list of addresses from a .csv file, 
with the Google Distance Matrix API used to calculate edge weights. I created it out of an interest in campground hosts and their travels
after learning from them through interviews in my Human Computer Interaction course last semester. I'd also taken a large interest in
Golang and have been working on another use of the Distance Matrix API with Go to create efficient carpools to help arrange rides more easily
for users. Given a .csv file with a similar format to that in this repo, one could generate an MST with Prim's as well with their own API key.

## How can I use it?

Currently, this program is specifically geared toward navigating state parks; thus, the nomenclature involving parks, parknodes, etc. However,
the parks currently have no more information beyond just their name and address, so you can use it for any set of addresses. You will need to have Go installed
on your system, thus documentation can be found at https://golang.org/.
You will need a google cloud account with API services for Google Maps enabled, as well as a properly formatted .csv file as we've seen. Note that the API has rate limiting so there is a
limited capacity of elements you can use for origin and destination addresses. You will also need to import the google maps api from the command line, using the command "go get googlemaps.github.io/maps"
After doing so, run the command "go build" from the command line and an exe should be built, which you can also run from the command line.

## What does it do now?

It currently prints out a list of addresses and "children" addresses which are like child nodes. If you do not have a Google Maps API key then it will print out a verbose error message as well.
