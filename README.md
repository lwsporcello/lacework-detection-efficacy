# detection-testing
These tools are designed to simulate malware with multiple stages, so it can be detected by Lacework. All binaries are written in Go (sources are provided). The process is automated, and can be started by running the command below.

Run this command to start simulations:

```
curl https://raw.githubusercontent.com/sporcello7/detection-testing/main/lw-det-test.sh | bash
```

**NOTE**: You only need the bash script above to execute the detection testing simulations. All the other files in this repo are the binaries' source code. Binaries are already compiled and being hosted on the C2 server, and will be downloaded as part of the simulations.


lw-det-test.sh
--------------
This script will:
  1. Determine the OS type
  2. Download the lw-stage-1 first stage binary (based on OS type)
  3. Change permissions to lw-stage-1 so it can be executed
  4. Execute lw-stage-1 in the background


lw-stage-1
---------------
This binary will:
  1. Collect IP information about intefaces
  2. Download lw-stage-2 as a second stage binary
  3. Change permissions to lw-stage-2 so it can be executed
  4. Execute lw-stage-2
  5. Coninuously beacon to the C2 server (payload is the IP information collected in step 1)

**NOTE**: Make sure to kill the lw-stage-1 process, as it is designed to run in the background and will continue to beacon to the C2 server.


lw-stage-2
--------------
This binary will:
  1. Download install-demo-1.sh bash script, which can be used to install an XMRig coin miner
     (It will NOT execute the coinminer script)
  2. The script is downloaded from a "known" bad domain
