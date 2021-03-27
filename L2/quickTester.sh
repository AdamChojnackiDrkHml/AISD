#!/bin/bash

for (( i=1; i<=20; i++ ))
do
    cat Data1000 | ./zad1 --type quick --comp '>=' | tee -a avgTimeData1000Quick
    cat Data1000Asc | ./zad1 --type quick --comp '>=' | tee -a avgTimeData1000QuickSorted
    cat Data1000Asc | ./zad1 --type quick --comp '<=' | tee -a avgTimeData1000QuickSortedReverse

    cat Data1000 | ./zad1 --type quick --comp '>=' | tee -a avgTimeData1000Merge
    cat Data1000Asc | ./zad1 --type quick --comp '>=' | tee -a avgTimeData1000MergeSorted
    cat Data1000Asc | ./zad1 --type quick --comp '<=' | tee -a avgTimeData1000MergeSortedReverse
    
    cat Data1000 | ./zad1 --type quick --comp '>=' | tee -a avgTimeData1000Insert
    cat Data1000Asc | ./zad1 --type quick --comp '>=' | tee -a avgTimeData1000InsertSorted
    cat Data1000Asc | ./zad1 --type quick --comp '<=' | tee -a avgTimeData1000InsertSortedReverse
done