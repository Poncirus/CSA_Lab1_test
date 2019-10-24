# 测试用例 1-6
#### 包含的指令
- [ ] ADDU
- [ ] SUBU
- [x] LW
- [x] SW
- [ ] BEQ

#### 采分点
- [x] RAW hazards
- [ ] BEQ

#### 汇编指令
```
LW    R1, R0, 0
LW    R2, R1, -4
SW    R2, R0, 8
HALT
```