# AV Safe Payloads

Various payloads that can bypass different Antiviruses. Most of my focus is bypassing Windows Defender due to it's newer implementations and it being widespread, but as such they can bypass most newer Antivuses. None of my ideas I would consider "new". I've taken different methods from people much smarter than me and tailored them for me. Please take a look at the resources below on the different tools and posts I've used.

## Payloads

#### InMemory_b64_CSharp.cs

Is meant to be compiled and ran on target. Will compile a base64 encoded CSharp payload (line 17) and execute the newly compiled code in memory. I like using the [CSharp Reverse Shell](https://gist.github.com/fdiskyou/56b9a4482eecd8e31a1d72b1acb66fac) but it should work with whatever CSharp script you decide to use.

Why B64 Encode the payload? Well the ultimate goal is to have a encrypted string and have the original exe pull down a decryption key. If for some reason the exe is discovered, they would only see the code that is decrypting a string and not the portion of the code that is calling back to our infrastructure. See below in Resources on why this is beneficial.

#### Bad_Macro.txt

Takes the basic idea from above and applies it to a macro. I haven't figured out how to compile and execute the code in memory with VBA yet, so this uses a base64 encoded exe and will right it to disk in the %temp% directory. Honestly, I'm surprised this works at all. Again, this was tested positively with Windows Defender. Other AVs, however, flag it because it's spawning suspicious process from excel. I know there are ways around that, just haven't messed with it enough yet.

## Resources

### Tools
[SharpShooter](https://github.com/mdsecactivebreach/SharpShooter)

[GreatSCT](https://github.com/GreatSCT/GreatSCT)

[Unicorn](https://github.com/trustedsec/unicorn)


### Articles


### Videos
