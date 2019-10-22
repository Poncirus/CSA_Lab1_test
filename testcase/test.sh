#!/bin/bash

for i in */; do
    echo -e "=============================================================\n>>>>  $i"
    cd $i
    ../MIPS_pipeline
    diff -wq dmemresult*
    DMEM_SAME=$?
    diff -wq RFresult*
    RF_SAME=$?
    [[ $DMEM_SAME == 1 || $RF_SAME == 1 ]] && echo 'Result differs from answer. Press Enter to view.' && read
    [[ $DMEM_SAME == 1 ]] && vimdiff dmemresult*
    [[ $RF_SAME == 1 ]] && vimdiff RFresult*
    cd ..
done
