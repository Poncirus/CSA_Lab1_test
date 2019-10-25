# testcase

Folder contains testcases for lab1

Testcases are listed from easy to hard

## Automatic Testing

**Automatic Testing only compares dmemresult and RFresult, the program may have problems even if passing the test**

1. install git and vim: `sudo apt-get install git vim`
2. clone project: `git clone https://github.com/LiaoHanwen/CSA_Lab1_test.git`
3. cd to testcase: `cd CSA_Lab1_test/testcase`
4. Put your `MIPS_pipeline.cpp` file in this directory.
5. Run `make --always-make` to build.
6. Run `make test` to begin testing for all test cases.
7. Optionally use `make test case=THE_TESTCASE_YOU_WANT_TO_TEST` to test a single case.

## Manually Testing method

Copy the imem.txt and dmem.txt of a testcase to the target directory and run the program. After the program is completed, compare the generated dmemresult.txt, RFresult.txt and stateresult.txt with the file with the same name under the test case.

## Testcases Introduction

### GROUP 0 Grading scheme
- [ ] RAW hazards
- [ ] BEQ

### GROUP 1 Grading scheme
- [x] RAW hazards
- [ ] BEQ

### GROUP 2 Grading scheme
- [x] RAW hazards
- [x] BEQ
