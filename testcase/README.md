# testcase

目录包含了Lab1的测试用例

测试用例从简单到困难进行排列

## 自动测试方法

1. 安装git和vim: `sudo apt-get install git vim`
2. clone项目: `git clone https://github.com/LiaoHanwen/CSA_Lab1_test.git`
3. 进入testcase目录: `cd CSA_Lab1_test/testcase`
4. 将MIPS_pipeline.cpp源文件到目录下
5. 编译源文件: `make --always-make`
6. 进行自动测试: `make test`

## Automatic Testing

1. Put your `MIPS_pipeline.cpp` file in this directory.
2. Run `make --always-make` to build.
3. Run `make test` to begin testing for all test cases.

## 手动测试方法

将某一个测试用例的imem.txt和dmem.txt拷贝到目标目录下并运行程序，程序完成后将生成的dmemresult.txt，RFresult.txt和stateresult.txt与测试用例下的同名文件进行对比

## 测试用例简介

### GROUP 0 采分点
- [ ] RAW hazards
- [ ] BEQ

### GROUP 1 采分点
- [x] RAW hazards
- [ ] BEQ

### GROUP 2 采分点
- [ ] RAW hazards
- [x] BEQ
  
