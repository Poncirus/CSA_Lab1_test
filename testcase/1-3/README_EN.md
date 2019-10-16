# TESTCASE 1-3 

This testcase is designed for MEM-MEM forwarding

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
SW    R1, R0, 4
HALT
```