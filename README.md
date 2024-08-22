## EN-DEcryptor

A web application written in ```Go``` to *encrypt and decrypt images* by making RESTapi calls by using is Go's standard library package ```http/net```. The idea is to send and receive images securely. This application is a basic structure and under development.

**The IDEA**
* The receiver don't have to create any account, and the image can be retrieved using a hash provided to the sender. 
* The supported formats are JPEG and PNG. The sender or receiver can retrieve the image as original format and size.
* SHA 256 is used to generate hash and the cipher mode used is ```Galois Counter Mode```.

**Steps to Implement**
* Implement database connection
* Sender login(oneclick)
* Optimize code(goroutines)
* Include more types and formats.