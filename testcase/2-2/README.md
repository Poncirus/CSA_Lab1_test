# 测试用例 2-2

斐波那契数列生成器
Fibonacci sequence generator

```
.align 4
[0] = Initial value
[1] = Address step amount
[2] = Stop address
```

#### 包含的指令
- [x] ADDU
- [x] SUBU
- [x] LW
- [x] SW
- [x] BEQ

#### 采分点
- [x] RAW hazards
- [x] BEQ

#### 汇编指令
```
lw r10,r0,0
lw r11,r0,4
lw r12,r0,8
addu r2,r0,r0
lw r3,r0,12
addu r1,r2,r0
addu r2,r3,r0
addu r3,r1,r2
sw r3,r10,0
addu r10,r10,r11
subu r0,r0,r0
subu r0,r0,r0
beq r12,r10,-8
halt
```
