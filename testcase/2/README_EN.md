# TESTCASE 1 
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
ADDU  R0, R0, r0
ADDU  R0, R0, r0
SUBU  R3, R1, R2
ADDU  R0, R0, r0
ADDU  R0, R0, r0
SW    R3, R0, 8
HALT
```