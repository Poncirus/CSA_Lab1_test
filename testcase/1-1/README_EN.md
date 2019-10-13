# TESTCASE 1-1 (Unverified)

This testcase is designed for EX-EX forwarding

#### Instruction
- [ ] ADDU
- [x] SUBU
- [x] LW
- [x] SW
- [ ] BEQ

#### Grading scheme
- [ ] RAW hazards
- [ ] BEQ

#### Assembly
```
LW    R1, R0, 0
LW    R2, R0, 4
ADDU  R7, R8, r9
ADDU  R7, R8, r9
ADDU  R3, R1, R2
ADDU  R4, R3, R2
ADDU  R5, R1, R4
HALT
```