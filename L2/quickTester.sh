#!/bin/bash

java generator 10 > Data10
java generator 100 > Data100
cat Data10 | ./zad1 --type quick --comp '>=' | tee -a Data10Asc
cat Data10 | ./zad1 --type quick --comp '<=' | tee -a Data10Des
cat Data100 | ./zad1 --type quick --comp '>=' | tee -a Data100Asc
cat Data100 | ./zad1 --type quick --comp '<=' | tee -a Data100Des
for (( i=1; i<=20; i++ ))
do
    cat Data1000 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData1000Quick
    cat Data1000Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData1000QuickSorted
    cat Data1000Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData1000QuickSortedReverse

    cat Data1000 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData1000Merge
    cat Data1000Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData1000MergeSorted
    cat Data1000Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData1000MergeSortedReverse
    
    cat Data1000 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData1000Insert
    cat Data1000Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData1000InsertSorted
    cat Data1000Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData1000InsertSortedReverse

    cat Data100 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData100Quick
    cat Data100Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData100QuickSorted
    cat Data100Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData100QuickSortedReverse

    cat Data100 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData100Merge
    cat Data100Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData100MergeSorted
    cat Data100Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData100MergeSortedReverse
    
    cat Data100 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData100Insert
    cat Data100Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData100InsertSorted
    cat Data100Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData100InsertSortedReverse

    cat Data10 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData10Quick
    cat Data10Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData10QuickSorted
    cat Data10Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData10QuickSortedReverse

    cat Data10 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData10Merge
    cat Data10Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData10MergeSorted
    cat Data10Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData10MergeSortedReverse
    
    cat Data10 | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData10Insert
    cat Data10Asc | ./zad1 --type quick --comp '>=' | tee -a SwapTest/avgSwapData10InsertSorted
    cat Data10Asc | ./zad1 --type quick --comp '<=' | tee -a SwapTest/avgSwapData10InsertSortedReverse
done