# Security Audit Report
**Project:** mx-chain-simulator-go-NewArc
**Audit Type:** Real-World Security Assessment - NewArc Branch (Post-Remediation)

## Executive Summary
This targeted security review identifies 3 remaining high and medium-severity vulnerabilities in the `mx-chain-simulator-go-NewArc` repository. While previous critical flaws (CORS, Auth, Mutex) have been addressed, new resource exhaustion vectors and insecure deployment configurations have been identified. The current state poses a threat to service stability and host security.

| Tier | Count | Meaning |
| :--- | :--- | :--- |
| **Critical / High** | 1 | Must fix before any production deployment. Remote crashes, DoS, or security bypass. |
| **Medium** | 2 | Important security hardening; exploitable under specific conditions. |
| **Low / Informational** | 0 | - |

**Overall Merge Status: REJECT -- 3 security threats identified.**

---

## Issue Index
| # | Title | File | Severity | Decision |
| :- | :--- | :--- | :--- | :--- |
| 1 | Unbounded Resource Consumption (DoS) | `pkg/facade/simulatorFacade.go` | High | **REJECT** |
| 2 | Symlink Following in Config Fetching | `pkg/proxy/configs/copy.go` | Medium | **REJECT** |
| 3 | Insecure Container Configuration (Root) | `Dockerfile` | Medium | **REJECT** |

---

## 1. Unbounded Resource Consumption (DoS)
**File:** `pkg/facade/simulatorFacade.go`  
**Severity:** High -- CWE-400 (Uncontrolled Resource Consumption)

### Description
The API endpoints for block generation (`GenerateBlocks`) and epoch progression (`GenerateBlocksUntilEpochIsReached`) do not enforce an upper limit on the number of blocks or epochs requested. An attacker can submit a request with an extremely large value (e.g., 1,000,000,000), triggering a long-running CPU-bound loop.

### Risk
Since these operations are protected by a global mutex to ensure state consistency, a single long-running request will block all other mutating API calls indefinitely. This results in a complete Denial of Service (DoS) for the simulator's functionality, potentially hanging CI/CD runners or shared dev environments.

---

## 2. Symlink Following in Config Fetching
**File:** `pkg/proxy/configs/copy.go`  
**Severity:** Medium -- CWE-59 (Improper Link Resolution)

### Description
The `copyFolderWithAllFiles` function, used to move configurations from a cloned repository to the simulator's workspace, follows symbolic links without validation.

### Risk
If the simulator is configured to fetch configs from a malicious or compromised repository, that repository can contain symlinks pointing to sensitive files on the host system (e.g., `/etc/passwd` or `~/.ssh/id_rsa`). The simulator will copy the *contents* of these linked files into the workspace, leading to arbitrary file read and potential data exfiltration.

---

## 3. Insecure Container Configuration (Root)
**File:** `Dockerfile`  
**Severity:** Medium -- CWE-250 (Execution with Unnecessary Privileges)

### Description
The `Dockerfile` does not define a non-root user for the container execution. The simulator process runs with full root privileges by default.

### Risk
In the event of a code execution vulnerability in the simulator or its underlying VM, the attacker gains immediate root access within the container. This significantly lowers the barrier for container escape, host filesystem access, and lateral movement within the network.

---
*End of Report*
