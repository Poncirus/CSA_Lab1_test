# 测试用例 1-2 

这个测试用例针对 MEM-EX Forwarding 设计

#### 包含的指令
- [x] ADDU
- [ ] SUBU
- [x] LW
- [ ] SW
- [ ] BEQ

#### 采分点
- [x] RAW hazards
- [ ] BEQ

#### 汇编指令
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