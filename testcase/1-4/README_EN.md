# TESTCASE 1-4 

This testcase is designed for stalling

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
ADDU  R2, R1, R1
HALT
```