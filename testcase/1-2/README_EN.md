# TESTCASE 1-2 (Unverified)

This testcase is designed for MEM-EX forwarding

#### Instruction
- [x] ADDU
- [ ] SUBU
- [x] LW
- [ ] SW
- [ ] BEQ

#### Grading scheme
- [x] RAW hazards
- [ ] BEQ

#### Assembly
```
LW    R1, R0, 0
LW    R2, R0, 4
ADDU  R7, R8, r9
ADDU  R7, R8, r9
ADDU  R3, R1, R2
ADDU  R7, R8, r9
ADDU  R4, R3, R2
ADDU  R7, R8, r9
ADDU  R5, R1, R4
HALT
```