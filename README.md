# bookshop-go

A simple book store API in need of input validation/sanitization.

This is a part of the University of Wyoming's Secure Software Design Course (Spring 2023). This is the base repository to be forked and updated for various assignments. Alternative language versions are available in:

- [Javascript](https://github.com/andey-robins/bookshop-js)
- [Rust](https://github.com/andey-robins/bookshop-rs)

## Versioning

`bookshop-go` is buit with:

- go version go1.19.3 darwin/arm64

## Usage

Start the api using `go run main.go`.

I recommend using [`httpie`](https://httpie.io) for testing of HTTP endpoints on the terminal. Tutorials are available elsewhere online, and you're free to use whatever tools you deem appropriate for testing your code.

## Code Review

The existing code does a fairly good job of preventing SQL injection by properly using the existing database framework, by calling db.Exec() and .Query()

On further examination, we couldn't find more security concerns. All input restrictions (length, special character, ...) seems to also be handled by the framework well.

There were a few bugs in the code that made it not work as intended. The first error is in customers/orders related db functions, Scan() was called without first calling Next(). While the syntax still actually works backend, it sends back an error json to end-user. So not only was the error returned not meaningful, it told customers that the call failed when it actually still made change in the db. Furthermore, there was an error in a sql query that asks for column "shipped" in table Books, which doesn't exist. The table name should be "PurchaseOrders".

A change I would recommend to a junior developer working on this project is before performing POST calls and commit to database, make sure that the id's are actually valid. Right now, we could still make a purchase order for bookid 100 and customerid 100, despite neither of those entities exist. However, on a separate examination, we saw that init.sql properly setup the reference between bookID and customerID in PurchaseOrders. Theoretically, it should have prevented invalid entry from being input. So the problem might be elsewhere and needs further examination.
