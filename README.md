Browser decrypter file
=================================

When I am not at home, sometimes I need to log in on a web service and using a different password for each service very often I don't remember it.

So I use a little html file crypted with AES encoding that contains all my useful passwords.

But there's a problem here. Each time that I register me on a new online service, I need to regenerate the whole file and it's so booooring.

This GOLANG script performs this automatically. I give it a password, a file to encode and a simple html template to use. In output I will have a template filled with the encrypted file.

This script use GOLANG to perform encryption using AES and the html file decrypts the encrypted string using CryptoJS ( http://code.google.com/p/crypto-js ).
