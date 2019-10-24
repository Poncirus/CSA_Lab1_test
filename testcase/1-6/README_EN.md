# TESTCASE 1-6
#### Instruction
- [ ] ADDU
- [ ] SUBU
- [x] LW
- [x] SW
- [ ] BEQ

#### Grading scheme
- [x] RAW hazards
- [ ] BEQ

#### Assembly
```
LW    R1, R0, 0
LW    R2, R1, -4
SW    R2, R0, 8
HALT
```