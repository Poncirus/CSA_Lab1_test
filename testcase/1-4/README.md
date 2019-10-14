# 测试用例 1-3 (未验证)

这个测试用例针对 stall 设计

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
ADDU  R2, R1, R1
HALT
```