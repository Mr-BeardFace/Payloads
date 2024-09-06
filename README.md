# AV Safe Payloads

Various payloads that can bypass different Antiviruses. Most of my focus is bypassing Windows Defender due to it's newer implementations and it being widespread, but as such they can bypass most newer Antivuses. I wouldn't consider any of my ideas "new". I've taken different methods from people much smarter than me and tailored them to what I want/need. Please take a look at the credits below on the different tools, articles, videos I've used.

## Payloads

#### InMemory_b64_CSharp.cs

Is meant to be compiled and ran on target. Will compile a base64 encoded CSharp payload (line 17) and execute the newly compiled code in memory. I like using the [CSharp Reverse Shell](https://gist.github.com/fdiskyou/56b9a4482eecd8e31a1d72b1acb66fac) but it should work with whatever CSharp script you decide to use.

Why base64 Encode the payload? Well the ultimate goal is to have an encrypted string and have the original exe pull down a decryption key. If for some reason the exe is discovered, only the code that is decrypting a string would be seen and not the portion of the code that is calling back to covert infrastructure. See below in Credits on why this is beneficial.

**This likely doesn't work anymore, but leaving for learning sake

#### Bad_Macro.txt

Takes the basic idea from above and applies it to a macro. I haven't figured out how to compile and execute the code in memory with VBA yet, so this uses a base64 encoded exe and will write it to disk in the %temp% directory. Honestly, I'm surprised this works at all. Again, this was tested positively with Windows Defender. Other AVs, however, flag it because word or excel docs are spawning unrelated/suspicious processes. I know there are ways around that, just haven't messed with it enough yet.

**This likely doesn't work anymore, but leaving for learning sake

## Credits

### Tools
[SharpShooter](https://github.com/mdsecactivebreach/SharpShooter)

A payload creation framework for the retrieval and execution of arbitrary CSharp source code. It leverages James Forshaw's DotNetToJavaScript tool to invoke methods from the SharpShooter DotNet serialised object.

[GreatSCT](https://github.com/GreatSCT/GreatSCT)

A tool designed to generate metasploit payloads that bypass common anti-virus solutions and application whitelisting solutions.

[Magic Unicorn](https://github.com/trustedsec/unicorn)

A simple tool for using a PowerShell downgrade attack and inject shellcode straight into memory. Based on Matthew Graeber's powershell attacks and the powershell bypass technique presented by David Kennedy (TrustedSec) and Josh Kelly at Defcon 18.

### Articles
https://www.mdsec.co.uk/2019/02/macros-and-more-with-sharpshooter-v2-0/

https://www.microsoft.com/security/blog/2018/09/27/out-of-sight-but-not-invisible-defeating-fileless-malware-with-behavior-monitoring-amsi-and-next-gen-av/

https://support.microsoft.com/en-us/help/304655/how-to-programmatically-compile-code-using-c-compiler



### Videos
https://www.youtube.com/watch?v=MHc3XP3XC4I
