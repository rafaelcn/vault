# Vault

A cli tool to be used as an easy encryption and decryption of data implemented
in Go. Right now it doesn't has all the bells and rings to encrypt and decrypt
files but that will be, if I have some time, implemented in the future.

# About encryption

The encryption algorithm used is the AES with a hashed (MD5) password (I know,
collisions, if you are really paranoid there's no current flag to change which
hash to use, sorry :poop:) with block ciphers of 128 bit. 



