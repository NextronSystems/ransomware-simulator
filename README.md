# QuickBuck - Ransomware Simulator

      ____         _       __    ___              __
     / __ \ __ __ (_)____ / /__ / _ ) __ __ ____ / /__
    / /_/ // // // // __//  '_// _  |/ // // __//  '_/
    \___\_\\_,_//_/ \__//_/\_\/____/ \_,_/ \__//_/\_\ 

The goal of this repository is to provide a simple, harmless way to check your AV's protection on ransomware.

This tool simulates typical ransomware behaviour, such as:

- Staging from a Word document macro
- Deleting Volume Shadow Copies
- Encrypting documents
- Dropping a ransomware note to the user's desktop

The ransomware simulator takes no action that actually encrypts pre-existing files
on the device, or deletes Volume Shadow Copies. However, any AV products looking for such behaviour
should still hopefully trigger.

Each step, as listed above, can also be disabled via a command line flag. This allows you to check responses to later
steps as well, even if an AV already detects earlier steps.

## Usage

