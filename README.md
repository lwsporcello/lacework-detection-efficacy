# detection-testing
These tools are designed to simulate malware with multiple stages, so it can be detected by Lacework. All binaries are written in Go (sources are provided). The process is automated, and can be started by running the command below.

Run this command to start the whole process:

```
  curl https://raw.githubusercontent.com/sporcello7/detection-testing/main/lw-det-test.sh | bash
```

lw-det-test.sh
--------------
This script will:
  1. Download the lw-binary first stage binary
  2. Change permissions to lw-binary so it can be executed
  3. Execute lw-binary in the background


lw-binary
---------------
This binary will:
  1. Collect IP information about intefaces
  2. Download lw-binary-2 as a second stage binary
  3. Change permissions to lw-binary-2 so it can be executed
  4. Execute lw-binary-2
  5. Coninuously beacon to the C2 server (payload is the IP information collected in step 1)

NOTE: Make sure to kill the lw-binary process, as it is designed to run in the background and will continue to beacon to the C2 server.


lw-binary-2
--------------
This binary will:
  1. Download install-demo-1.sh bash script, which can be used to install an XMRig coin miner
     (It will NOT execute the coinminer script)
  2. The script is downloaded from a "known" bad domain
