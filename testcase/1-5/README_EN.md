# TESTCASE 1-5 (Unverified)

#### Instruction
- [x] ADDU
- [x] SUBU
- [x] LW
- [x] SW
- [ ] BEQ

#### Grading scheme
- [x] RAW hazards
- [ ] BEQ

#### Assembly
```
LW    R1, R0, 0
LW    R2, R0, 4
ADDU  R3, R1, R2
ADDU  R4, R3, R2
ADDU  R5, R3, R4
LW    R6, R0, 8
ADDU  R7, R6, R5
SW    R7, R0, 12
LW    R8, R0, 4
Sw    r8,r0,16
HALT
```