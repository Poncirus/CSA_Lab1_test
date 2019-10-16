# 测试用例 1-3 

这个测试用例针对 MEM-MEM Forwarding 设计

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
SW    R1, R0, 4
HALT
```