## EN-DEcryptor

A web application to *encrypt* and *decrypt* images written in Go. The idea is send images securely. This application is a basic structure and under development.

**The IDEA**
* The receiver don't have to create any account, and the image can be retrieved using a hash provided to the sender. 
* The supported formats are JPEG and PNG. The sender or receiver can retrieve the image as original format and size.
* SHA 256 is used to generate hash and the cipher mode used is ```Galois Counter Mode```.

**Steps to impliment**
* Implement database connection
* Sender login(oneclick)
* Optimize code(goroutines)