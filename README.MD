

# Customers service

## Task description

Package main reads from the given customers.csv file and returns a sorted (data structure of your choice) of email domains along with the number of customers with e-mail addresses for each domain.  Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but *could* be 1m lines or run on a small machine).

## Solution notes

- :book: standard Go project layout (well, more or less :blush:)
- :white_check_mark: tests included

You made notice that file could be 1 million rows, so before coding 

I made calculations.`customer.csv` with 3000 rows ~= 170000 bytes

1 million rows ~= 56610000 bytes ~= 56 megabytes. I expect that we do not optimize solution for NASA satellite and machine will have at least 1 GB of RAM.

Until `1024Mb / 56Mb ~= 18` million rows algorithm will work fast, after it will blow up :)

## HOWTO

- run with `make run`
- test with `make test`

## A picture is worth a thousand words

![make-run.png](images%2Fmake-run.png)