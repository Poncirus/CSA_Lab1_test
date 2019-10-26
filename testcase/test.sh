#!/bin/bash

DIFF=vimdiff
DIFFEND=""
if ! command -v vimdiff 2>&1 >/dev/null ; then
    DIFF="diff --side-by-side --color=always"
    DIFFEND="|less"
fi

testcase=$1
if [[ "$testcase" == "" ]]; then
    testcase=$(ls -d -- */)
fi

for i in $testcase; do
    echo -e "=============================================================\n>>>>  $i"
    cd $i
    rm dmemresult.txt RFresult.txt stateresult.txt
    ../MIPS_pipeline
    diff -wq dmemresult*
    DMEM_SAME=$?
    diff -wq RFresult*
    RF_SAME=$?
    [[ "$DMEM_SAME" == "1" || "$RF_SAME" == "1" ]] && echo 'Result differs from answer. Press Enter to view.' && read
    [[ "$DMEM_SAME" == "1" ]] && eval "$DIFF dmemresult* $DIFFEND"
    [[ "$RF_SAME" == "1" ]] && eval "$DIFF RFresult* $DIFFEND"
    cd ..
done
