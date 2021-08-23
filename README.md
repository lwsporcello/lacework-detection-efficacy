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
  2. Download the lw-stage-1-* first stage binary (file name based on OS type, i.e. lw-stage-1-linux, lw-stage-1-darwin)
  3. Change permissions to stage 1 binary so it can be executed
  4. Execute stage 1 binary in the background


lw-stage-1-*
---------------
This binary will:
  1. Determine OS type
  2. Download lw-stage-2-* as a second stage binary (file name based on OS type, i.e. lw-stage-2-linux, lw-stage-2-darwin)
  3. Change permissions to stage 2 binary so it can be executed
  4. Execute stage 2 binary
  5. Beacon once to the C2 server, then terimnate


lw-stage-2-*
--------------
This binary will:
  1. Download install-demo-1.sh bash script, which can be used to install an XMRig coin miner
     (It will NOT execute the coinminer script)
  2. The script is downloaded from a "known" bad domain
  3. Continuosly beacon to C2 server

**NOTE**: Make sure to kill the lw-stage-2 process, as it is designed to run in the background and will continue to beacon to the C2 server if not killed.
