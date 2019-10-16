# 测试用例 2-1 

#### 包含的指令
- [x] ADDU
- [x] SUBU
- [x] LW
- [ ] SW
- [x] BEQ

#### 采分点
- [ ] RAW hazards
- [x] BEQ

#### 汇编指令
```
LW    R1, R0, 0
LW    R2, R0, 4
ADDU  R4, R5, R6
ADDU  R4, R5, R6
BEQ   R2, R1, 1
SUBU  R4, R2, R1
ADDU  R5, R2, R1
HALT
```