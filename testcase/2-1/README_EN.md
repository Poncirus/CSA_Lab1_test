# TESTCASE 2-1 (Unverified)

#### Instruction
- [x] ADDU
- [x] SUBU
- [x] LW
- [ ] SW
- [x] BEQ

#### Grading scheme
- [ ] RAW hazards
- [x] BEQ

#### Assembly
```
LW    R1, R0, 0
LW    R2, R0, 4
LW    R3, R0, 4
ADDU  R4, R5, R6
ADDU  R4, R5, R6
BEQ   R2, R3, 4
SUBU  R4, R2, R1
ADDU  R4, R2, R1
HALT
```