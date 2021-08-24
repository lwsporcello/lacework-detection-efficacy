# detection-testing
These tools are designed to simulate malware with multiple stages, so it can be detected by [Lacework](https://lacework.com). Each simulation is explained below, with its mapping to the [MITRE ATT&CK](https://attack.mitre.org/#) tactics and techniques. All binaries are written in Go (sources are provided). The process is automated, and can be started by running the command below.

Run this command to begin:

```
bash <(curl https://raw.githubusercontent.com/sporcello7/detection-testing/main/lw-det-test.sh)
```

This script will prompt you to choose which simualtions you want to run. Enter all simulations you want to run in a comma separated list.

```
-----------------------------------------------------
| Welcome to the Lacework detection testing script. |
| This script will run attack simulation in your    |
| workload running the Lacework agent. The agent    |
| will capture this activity and generate events in |
| the Lacework UI. Begin by choosing which          |
| simulations you want to run...                    |
-----------------------------------------------------

  0. Quit
  1. Simulation 1: Network Scan & Brute Force
  2. Simulation 2: Multi-Stage Malware

You may enter one simulation (i.e. 2) or multiple comma separated (i.e. 1,2)

Enter Selection(s): 
```

**NOTE**: You only need the bash script above to execute the detection testing simulations. All the other files in this repo are the source code for binaries used in the simuations. Binaries are already compiled and being hosted on the C2 server, and will be downloaded as part of the simulation execution by the bash script.

### lw-det-test.sh
This script will:
  1. Download the lw-stage-1 first stage binary
  2. Change permissions to stage 1 binary so it can be executed
  3. Execute stage 1 binary in the background

---

# Simulation 1
This simulation is designed to replicate attacker activity attempting Reconnaissance (active scanning), Discovery (network service scan) and Credential Access (brute force) tactics within your environment. Binaries used in this simulation include:

### lw-scan-brute
This binary will:
  1. Get all IPs from active interfaces on the host (excludes down interfaces, loopback, ipv6)
  2. Based on the IP(s), a list of potential hosts in the subnet is generated
  3. Each IP in ths list is scanned for open port 22 (ssh)
  4. One host is choosen from the list (the first host), and 10x ssh logins are attempted with invalid credentials

This binary will be executed in the background, and usually takes 1-2 minutes to complete.

---

# Simulation 2
This simulation is designed to replicate attacker activity attempting Command and Control (multi-stage channels) and Execution (command and scripting interpreter) and Impact (resource hijacking) tactics within your environment. Binaries used in this simulation include:

### lw-stage-1
This binary will:
  1. Download lw-stage-2 as a second stage binary
  2. Change permissions to stage 2 binary so it can be executed
  3. Execute stage 2 binary
  4. Beacon once to the C2 server, then terimnate

This binary will be executed in the background, and usually takes 1 minute or less to complete.

### lw-stage-2
This binary will:
  1. Download install-demo-1.sh bash script, which can be used to install an XMRig coin miner
     (It will NOT execute the coinminer script)
  2. The script is downloaded from a "known" bad domain
  3. Continuosly beacon to C2 server

This binary will be executed in the background as a child process of lw-stage-1. This binary run indefinitely, **so be sure to manually kill this process!!**
