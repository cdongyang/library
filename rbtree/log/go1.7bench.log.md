gc 1 @0.172s 0%: 0.17+0.85+0.36 ms clock, 0.34+0/1.7/1.1+0.72 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
gc 2 @0.198s 0%: 0.11+0.78+0.29 ms clock, 1.0+0.015/1.5/1.3+2.6 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
# github.com/cdongyang/library/rbtree
gc 1 @0.018s 0%: 0.091+4.0+0.42 ms clock, 0.27+0.15/1.9/4.9+1.2 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @0.052s 1%: 0.053+4.1+0.30 ms clock, 0.64+0.076/3.6/10+3.6 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 3 @0.117s 1%: 0.017+7.1+0.34 ms clock, 0.28+4.9/2.9/11+5.5 ms cpu, 9->9->5 MB, 10 MB goal, 16 P
gc 4 @0.191s 1%: 0.019+4.2+0.29 ms clock, 0.30+0.13/9.7/10+4.7 ms cpu, 10->11->6 MB, 11 MB goal, 16 P
# github.com/cdongyang/library/rbtree_test
gc 1 @0.006s 2%: 0.081+3.9+0.48 ms clock, 0.24+0.006/2.1/4.4+1.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @0.027s 2%: 0.046+6.9+0.39 ms clock, 0.50+0.14/6.8/10+4.3 ms cpu, 6->7->5 MB, 7 MB goal, 16 P
gc 3 @0.090s 2%: 0.022+4.4+0.27 ms clock, 0.35+0.17/10/19+4.3 ms cpu, 10->11->7 MB, 11 MB goal, 16 P
gc 4 @0.193s 1%: 0.019+5.1+0.35 ms clock, 0.31+0.33/10/27+5.7 ms cpu, 14->14->8 MB, 15 MB goal, 16 P
gc 5 @0.313s 1%: 0.020+6.8+0.32 ms clock, 0.33+0/17/28+5.1 ms cpu, 16->17->10 MB, 17 MB goal, 16 P
# testmain
gc 1 @0.006s 3%: 0.062+3.9+0.30 ms clock, 0.18+0.089/4.8/1.1+0.91 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @0.025s 3%: 0.043+4.0+0.31 ms clock, 0.39+0.013/5.4/12+2.8 ms cpu, 6->7->5 MB, 7 MB goal, 16 P
# testmain
gc 1 @0.002s 4%: 0.099+4.7+0.28 ms clock, 0.29+3.9/1.0/0.19+0.86 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 2 @0.019s 4%: 0.041+8.1+0.37 ms clock, 0.28+1.4/8.7/1.5+2.5 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 3 @0.059s 3%: 0.075+9.9+0.25 ms clock, 0.67+0.45/16/10+2.2 ms cpu, 13->14->13 MB, 15 MB goal, 16 P
gc 4 @0.121s 3%: 0.054+14+0.37 ms clock, 0.59+0.037/40/7.6+4.1 ms cpu, 24->25->24 MB, 26 MB goal, 16 P
=== RUN   TestSet
--- PASS: TestSet (0.00s)
gc 1 @0.001s 4%: 0.005+0+0.32 ms clock, 0.021+0/0/0+1.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2 @0.001s 7%: 0.025+0+0.41 ms clock, 0.10+0/0/0+1.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 184352 HeapInuse: 417792 HeapObjects: 246 HeapIdle 204800 HeapReleased 0 HeapSys 622592
gc 3 @0.002s 10%: 0.003+0+0.31 ms clock, 0.034+0/0/0+2.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSort1E3-16                               	gc 4 @0.003s 15%: 0.007+0+0.36 ms clock, 0.074+0/0/0+3.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 193760 HeapInuse: 417792 HeapObjects: 265 HeapIdle 1024000 HeapReleased 0 HeapSys 1441792
gc 5 @0.004s 18%: 0.003+0+0.30 ms clock, 0.052+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 6 @0.004s 22%: 0.030+0+0.33 ms clock, 0.42+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 7 @0.005s 26%: 0.004+0+0.39 ms clock, 0.064+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 8 @0.006s 30%: 0.042+0+0.40 ms clock, 0.64+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 9 @0.006s 31%: 0.019+0+0.36 ms clock, 0.29+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 10 @0.007s 33%: 0.005+0+0.31 ms clock, 0.078+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 11 @0.008s 35%: 0.013+0+0.46 ms clock, 0.20+0/0/0+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 12 @0.008s 37%: 0.003+0+0.32 ms clock, 0.059+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 13 @0.009s 38%: 0.003+0+0.32 ms clock, 0.053+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 14 @0.009s 39%: 0.043+0+0.34 ms clock, 0.64+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 15 @0.010s 40%: 0.003+0+0.34 ms clock, 0.051+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 16 @0.011s 41%: 0.003+0+0.38 ms clock, 0.057+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 17 @0.011s 42%: 0.001+0+0.54 ms clock, 0.030+0/0/0+8.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 18 @0.012s 43%: 0.003+0+0.28 ms clock, 0.055+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 19 @0.013s 44%: 0.017+0+0.40 ms clock, 0.26+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 20 @0.013s 45%: 0.003+0+0.37 ms clock, 0.055+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 21 @0.014s 45%: 0.003+0+0.37 ms clock, 0.050+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 22 @0.014s 46%: 0.037+0+0.34 ms clock, 0.56+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 23 @0.015s 46%: 0.001+0+0.32 ms clock, 0.027+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 24 @0.016s 47%: 0.045+0+0.35 ms clock, 0.67+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 25 @0.016s 47%: 0.015+0+0.34 ms clock, 0.23+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 26 @0.017s 47%: 0.045+0+0.30 ms clock, 0.68+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 27 @0.018s 47%: 0.003+0+0.30 ms clock, 0.054+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 28 @0.018s 48%: 0.046+0+0.39 ms clock, 0.69+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 29 @0.019s 48%: 0.032+0+0.30 ms clock, 0.48+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 30 @0.019s 48%: 0.045+0+0.28 ms clock, 0.67+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 31 @0.020s 48%: 0.019+0+0.47 ms clock, 0.28+0/0/0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 32 @0.021s 49%: 0.044+0+0.30 ms clock, 0.67+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 199328 HeapInuse: 417792 HeapObjects: 274 HeapIdle 925696 HeapReleased 0 HeapSys 1343488
gc 33 @0.021s 49%: 0.022+0+0.43 ms clock, 0.33+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       789 ns/op	      11 B/op	       0 allocs/op
gc 34 @0.022s 49%: 0.005+0+0.31 ms clock, 0.078+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279312 HeapInuse: 491520 HeapObjects: 1773 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 35 @0.025s 45%: 0.003+0+0.41 ms clock, 0.056+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsert1E3-16                          	gc 36 @0.025s 45%: 0.016+0+0.40 ms clock, 0.24+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 37 @0.029s 42%: 0.001+0+0.55 ms clock, 0.031+0/0/0+8.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 38 @0.029s 43%: 0.029+0+0.40 ms clock, 0.44+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 39 @0.032s 40%: 0.003+0+0.42 ms clock, 0.060+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 40 @0.033s 41%: 0.044+0+0.35 ms clock, 0.67+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 41 @0.036s 38%: 0.015+0+0.33 ms clock, 0.22+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 42 @0.036s 39%: 0.046+0+0.56 ms clock, 0.69+0/0/0+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 43 @0.040s 37%: 0.004+0+0.43 ms clock, 0.065+0/0/0+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 44 @0.040s 37%: 0.004+0+0.40 ms clock, 0.068+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 45 @0.043s 36%: 0.003+0+0.37 ms clock, 0.056+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 46 @0.044s 36%: 0.003+0+0.33 ms clock, 0.052+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 47 @0.047s 35%: 0.003+0+0.45 ms clock, 0.057+0/0/0+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 48 @0.047s 35%: 0.012+0+0.39 ms clock, 0.18+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 49 @0.050s 34%: 0.003+0+0.30 ms clock, 0.055+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 50 @0.051s 34%: 0.012+0+0.40 ms clock, 0.19+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 51 @0.054s 33%: 0.003+0+0.36 ms clock, 0.053+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 52 @0.054s 33%: 0.027+0+0.33 ms clock, 0.40+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 53 @0.057s 32%: 0.020+0+0.36 ms clock, 0.30+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 54 @0.058s 32%: 0.027+0+0.36 ms clock, 0.41+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 55 @0.061s 31%: 0.003+0+0.40 ms clock, 0.053+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 56 @0.061s 32%: 0.055+0+0.35 ms clock, 0.83+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 57 @0.064s 31%: 0.016+0+0.32 ms clock, 0.24+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 58 @0.065s 31%: 0.032+0+0.32 ms clock, 0.48+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 59 @0.068s 30%: 0.003+0+0.44 ms clock, 0.054+0/0/0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 60 @0.068s 31%: 0.029+0+0.33 ms clock, 0.44+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 61 @0.071s 30%: 0.006+0+0.30 ms clock, 0.10+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 62 @0.072s 30%: 0.014+0+0.31 ms clock, 0.21+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 63 @0.075s 29%: 0.002+0+0.30 ms clock, 0.032+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 64 @0.075s 29%: 0.029+0+0.28 ms clock, 0.43+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 65 @0.078s 29%: 0.003+0+0.37 ms clock, 0.056+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 66 @0.079s 29%: 0.044+0+0.40 ms clock, 0.66+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279376 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 67 @0.082s 28%: 0.017+0+0.45 ms clock, 0.26+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      3060 ns/op	      91 B/op	       2 allocs/op
gc 68 @0.082s 29%: 0.004+0+0.34 ms clock, 0.072+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375504 HeapInuse: 589824 HeapObjects: 3274 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 69 @0.087s 28%: 0.003+0+0.35 ms clock, 0.056+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetErase1E3-16                           	gc 70 @0.088s 28%: 0.005+0+0.33 ms clock, 0.079+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375568 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 71 @0.092s 27%: 0.003+0+0.32 ms clock, 0.061+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 72 @0.093s 27%: 0.046+0+0.35 ms clock, 0.69+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 73 @0.098s 26%: 0.014+0+0.35 ms clock, 0.21+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 74 @0.098s 26%: 0.042+0+0.33 ms clock, 0.63+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 75 @0.104s 25%: 0.003+0+0.31 ms clock, 0.058+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 76 @0.104s 25%: 0.043+0+0.37 ms clock, 0.64+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 77 @0.109s 24%: 0.015+0+0.36 ms clock, 0.23+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 78 @0.110s 25%: 0.063+0+0.30 ms clock, 0.94+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 79 @0.115s 24%: 0.004+0+0.40 ms clock, 0.063+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 80 @0.115s 24%: 0.044+0+0.34 ms clock, 0.67+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 81 @0.120s 23%: 0.014+0+0.42 ms clock, 0.22+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 82 @0.121s 24%: 0.046+0+0.44 ms clock, 0.70+0/0/0+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 83 @0.126s 23%: 0.045+0+0.38 ms clock, 0.68+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 84 @0.127s 23%: 0.072+0+0.36 ms clock, 1.0+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 85 @0.132s 23%: 0.014+0+0.38 ms clock, 0.21+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 86 @0.132s 23%: 0.048+0+0.36 ms clock, 0.72+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 87 @0.137s 22%: 0.012+0+0.36 ms clock, 0.19+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 88 @0.138s 22%: 0.046+0+0.33 ms clock, 0.70+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 89 @0.143s 22%: 0.003+0+0.36 ms clock, 0.055+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 90 @0.143s 22%: 0.046+0+0.41 ms clock, 0.70+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 91 @0.149s 21%: 0.001+0+0.37 ms clock, 0.030+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 92 @0.149s 22%: 0.004+0+0.41 ms clock, 0.066+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 93 @0.154s 21%: 0.002+0+0.37 ms clock, 0.036+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 94 @0.155s 21%: 0.031+0+0.33 ms clock, 0.47+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 95 @0.160s 21%: 0.025+0+0.37 ms clock, 0.38+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 96 @0.160s 21%: 0.044+0+0.35 ms clock, 0.67+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 97 @0.166s 21%: 0.016+0+0.39 ms clock, 0.24+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 98 @0.166s 21%: 0.036+0+0.39 ms clock, 0.54+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375632 HeapInuse: 589824 HeapObjects: 3278 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 99 @0.171s 20%: 0.001+0+0.49 ms clock, 0.029+0/0/0+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2431 ns/op	      99 B/op	       2 allocs/op
gc 100 @0.172s 21%: 0.007+0+0.45 ms clock, 0.11+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287568 HeapInuse: 499712 HeapObjects: 1775 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 101 @0.174s 21%: 0.003+0+0.38 ms clock, 0.058+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E3-16                  	gc 102 @0.175s 21%: 0.006+0+0.42 ms clock, 0.094+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 103 @0.177s 21%: 0.004+0+0.35 ms clock, 0.076+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 104 @0.177s 21%: 0.004+0+0.37 ms clock, 0.062+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 105 @0.180s 21%: 0.021+0+0.38 ms clock, 0.32+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 106 @0.180s 21%: 0.030+0+0.33 ms clock, 0.45+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 107 @0.183s 21%: 0.016+0+0.33 ms clock, 0.24+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 108 @0.183s 21%: 0.004+0+0.40 ms clock, 0.070+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 109 @0.185s 21%: 0.027+0+0.34 ms clock, 0.41+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 110 @0.186s 21%: 0.042+0+0.35 ms clock, 0.63+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 111 @0.188s 21%: 0.016+0+0.50 ms clock, 0.25+0/0/0+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 112 @0.189s 21%: 0.049+0+0.39 ms clock, 0.73+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 113 @0.191s 21%: 0.005+0+0.49 ms clock, 0.088+0/0/0+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 114 @0.192s 21%: 0.044+0+0.37 ms clock, 0.66+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 115 @0.194s 21%: 0.003+0+0.31 ms clock, 0.055+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 116 @0.194s 21%: 0.032+0+0.32 ms clock, 0.49+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 117 @0.196s 21%: 0.001+0+0.34 ms clock, 0.029+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 118 @0.197s 21%: 0.030+0+0.35 ms clock, 0.45+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 119 @0.199s 21%: 0.002+0+0.35 ms clock, 0.037+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 120 @0.200s 21%: 0.040+0+0.30 ms clock, 0.60+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 121 @0.202s 21%: 0.002+0+0.38 ms clock, 0.042+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 122 @0.203s 21%: 0.047+0+0.38 ms clock, 0.71+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 123 @0.205s 21%: 0.016+0+0.36 ms clock, 0.25+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 124 @0.206s 22%: 0.047+0+0.38 ms clock, 0.71+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 125 @0.208s 21%: 0.002+0+0.34 ms clock, 0.038+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 126 @0.208s 22%: 0.075+0+0.31 ms clock, 1.1+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 127 @0.211s 22%: 0.016+0+0.32 ms clock, 0.24+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 128 @0.211s 22%: 0.004+0+0.32 ms clock, 0.062+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 129 @0.213s 22%: 0.012+0+0.38 ms clock, 0.19+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 130 @0.214s 22%: 0.004+0+0.40 ms clock, 0.061+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1778 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 131 @0.216s 22%: 0.015+0+0.36 ms clock, 0.22+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2177 ns/op	      99 B/op	       2 allocs/op
gc 132 @0.217s 22%: 0.005+0+0.32 ms clock, 0.076+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250536 HeapInuse: 483328 HeapObjects: 1281 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 133 @0.219s 22%: 0.003+0+0.33 ms clock, 0.061+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E3-16          	gc 134 @0.220s 22%: 0.004+0+0.37 ms clock, 0.071+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 135 @0.223s 22%: 0.093+0+0.32 ms clock, 1.4+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 136 @0.223s 22%: 0.047+0+0.36 ms clock, 0.71+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 137 @0.226s 22%: 0.002+0+0.32 ms clock, 0.046+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 138 @0.227s 22%: 0.046+0+0.36 ms clock, 0.69+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 139 @0.229s 22%: 0.018+0+0.33 ms clock, 0.28+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 140 @0.230s 22%: 0.045+0+0.37 ms clock, 0.68+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 141 @0.233s 22%: 0.017+0+0.35 ms clock, 0.26+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 142 @0.233s 22%: 0.047+0+0.35 ms clock, 0.71+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 143 @0.236s 22%: 0.005+0+0.33 ms clock, 0.078+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 144 @0.237s 22%: 0.044+0+0.36 ms clock, 0.67+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 145 @0.239s 22%: 0.018+0+0.32 ms clock, 0.27+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 146 @0.240s 22%: 0.042+0+0.36 ms clock, 0.63+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 147 @0.243s 22%: 0.017+0+0.36 ms clock, 0.25+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 148 @0.243s 22%: 0.047+0+0.35 ms clock, 0.70+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 149 @0.246s 22%: 0.017+0+0.34 ms clock, 0.26+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 150 @0.247s 22%: 0.027+0+0.33 ms clock, 0.41+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 151 @0.250s 22%: 0.018+0+0.62 ms clock, 0.27+0/0/0+9.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 152 @0.250s 22%: 0.013+0+0.37 ms clock, 0.20+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 153 @0.253s 22%: 0.004+0+0.29 ms clock, 0.071+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 154 @0.254s 22%: 0.045+0+0.30 ms clock, 0.67+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 155 @0.257s 21%: 0.002+0+0.39 ms clock, 0.033+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 156 @0.257s 22%: 0.024+0+0.36 ms clock, 0.37+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 157 @0.260s 21%: 0.003+0+0.36 ms clock, 0.062+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 158 @0.261s 22%: 0.033+0+0.35 ms clock, 0.50+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 159 @0.264s 21%: 0.002+0+0.39 ms clock, 0.044+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 160 @0.264s 22%: 0.045+0+0.34 ms clock, 0.67+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 161 @0.267s 22%: 0.001+0+0.42 ms clock, 0.031+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 162 @0.268s 22%: 0.056+0+0.32 ms clock, 0.84+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 163 @0.270s 22%: 0.014+0+0.36 ms clock, 0.21+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 164 @0.271s 22%: 0.004+0+0.38 ms clock, 0.071+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250616 HeapInuse: 483328 HeapObjects: 1284 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 165 @0.274s 22%: 0.014+0+0.36 ms clock, 0.21+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2788 ns/op	      62 B/op	       1 allocs/op
gc 166 @0.274s 22%: 0.007+0+0.37 ms clock, 0.11+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279360 HeapInuse: 491520 HeapObjects: 1773 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 167 @0.277s 22%: 0.007+0+0.38 ms clock, 0.10+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkIntSetInsert1E3-16                       	gc 168 @0.277s 22%: 0.007+0+0.39 ms clock, 0.11+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 169 @0.279s 22%: 0.015+0+0.35 ms clock, 0.23+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 170 @0.280s 22%: 0.038+0+0.32 ms clock, 0.57+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 171 @0.283s 22%: 0.002+0+0.34 ms clock, 0.038+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 172 @0.283s 22%: 0.027+0+0.33 ms clock, 0.41+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 173 @0.285s 22%: 0.016+0+0.31 ms clock, 0.24+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 174 @0.286s 22%: 0.046+0+0.39 ms clock, 0.69+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 175 @0.288s 22%: 0.002+0+0.39 ms clock, 0.034+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 176 @0.289s 22%: 0.042+0+0.34 ms clock, 0.63+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 177 @0.291s 22%: 0.003+0+0.34 ms clock, 0.062+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 178 @0.291s 22%: 0.024+0+0.32 ms clock, 0.36+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 179 @0.294s 22%: 0.005+0+0.35 ms clock, 0.076+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 180 @0.294s 22%: 0.003+0+0.37 ms clock, 0.053+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 181 @0.296s 22%: 0.039+0+0.36 ms clock, 0.59+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 182 @0.297s 22%: 0.042+0+0.41 ms clock, 0.64+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 183 @0.299s 22%: 0.003+0+0.38 ms clock, 0.049+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 184 @0.300s 22%: 0.042+0+0.32 ms clock, 0.64+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 185 @0.302s 22%: 0.002+0+0.38 ms clock, 0.033+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 186 @0.303s 22%: 0.004+0+0.34 ms clock, 0.063+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 187 @0.305s 22%: 0.017+0+0.33 ms clock, 0.26+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 188 @0.305s 22%: 0.042+0+0.36 ms clock, 0.64+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 189 @0.308s 22%: 0.002+0+0.33 ms clock, 0.035+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 190 @0.308s 22%: 0.054+0+0.29 ms clock, 0.82+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 191 @0.311s 22%: 0.002+0+0.36 ms clock, 0.032+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 192 @0.311s 22%: 0.044+0+0.35 ms clock, 0.66+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 193 @0.313s 22%: 0.022+0+0.31 ms clock, 0.34+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 194 @0.314s 22%: 0.035+0+0.49 ms clock, 0.53+0/0/0+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 195 @0.316s 22%: 0.019+0+0.39 ms clock, 0.28+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 196 @0.317s 22%: 0.013+0+0.33 ms clock, 0.20+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 197 @0.319s 22%: 0.013+0+0.33 ms clock, 0.20+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 198 @0.319s 22%: 0.041+0+0.35 ms clock, 0.61+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 279424 HeapInuse: 491520 HeapObjects: 1776 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 199 @0.322s 22%: 0.073+0+0.32 ms clock, 1.1+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2204 ns/op	      91 B/op	       2 allocs/op
gc 200 @0.322s 22%: 0.007+0+0.34 ms clock, 0.11+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375552 HeapInuse: 589824 HeapObjects: 3274 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 201 @0.326s 22%: 0.027+0+0.33 ms clock, 0.41+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkIntSetErase1E3-16                        	gc 202 @0.326s 22%: 0.006+0+0.31 ms clock, 0.095+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 203 @0.330s 22%: 0.003+0+0.33 ms clock, 0.051+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 204 @0.330s 22%: 0.004+0+0.34 ms clock, 0.070+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 205 @0.334s 22%: 0.040+0+0.33 ms clock, 0.60+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 206 @0.334s 22%: 0.040+0+0.33 ms clock, 0.61+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 207 @0.338s 22%: 0.003+0+0.37 ms clock, 0.058+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 208 @0.338s 22%: 0.039+0+0.33 ms clock, 0.59+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 209 @0.342s 22%: 0.003+0+0.33 ms clock, 0.063+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 210 @0.342s 22%: 0.041+0+0.31 ms clock, 0.62+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 211 @0.346s 22%: 0.006+0+0.35 ms clock, 0.092+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 212 @0.346s 22%: 0.050+0+0.38 ms clock, 0.75+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 213 @0.350s 22%: 0.006+0+0.31 ms clock, 0.095+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 214 @0.350s 22%: 0.043+0+0.30 ms clock, 0.65+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 215 @0.354s 22%: 0.016+0+0.37 ms clock, 0.25+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 216 @0.354s 22%: 0.041+0+0.30 ms clock, 0.61+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 217 @0.358s 21%: 0.016+0+0.36 ms clock, 0.25+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 218 @0.358s 22%: 0.040+0+0.34 ms clock, 0.60+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 219 @0.362s 21%: 0.004+0+0.34 ms clock, 0.068+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 220 @0.362s 21%: 0.005+0+0.32 ms clock, 0.075+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 221 @0.366s 21%: 0.028+0+0.37 ms clock, 0.42+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 222 @0.366s 21%: 0.003+0+0.35 ms clock, 0.058+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 223 @0.370s 21%: 0.007+0+0.36 ms clock, 0.11+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 224 @0.370s 21%: 0.044+0+0.39 ms clock, 0.66+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 225 @0.374s 21%: 0.003+0+0.36 ms clock, 0.057+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 226 @0.375s 21%: 0.015+0+0.35 ms clock, 0.23+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 227 @0.378s 21%: 0.020+0+0.39 ms clock, 0.30+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 228 @0.379s 21%: 0.028+0+0.36 ms clock, 0.42+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 229 @0.382s 21%: 0.002+0+0.37 ms clock, 0.037+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 230 @0.383s 21%: 0.052+0+0.32 ms clock, 0.79+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 375616 HeapInuse: 589824 HeapObjects: 3277 HeapIdle 753664 HeapReleased 0 HeapSys 1343488
gc 231 @0.386s 21%: 0.003+0+0.34 ms clock, 0.054+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      1761 ns/op	      99 B/op	       2 allocs/op
gc 232 @0.387s 21%: 0.007+0+0.35 ms clock, 0.11+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287552 HeapInuse: 499712 HeapObjects: 1774 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 233 @0.389s 21%: 0.003+0+0.36 ms clock, 0.055+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndErase1E3-16               	gc 234 @0.389s 21%: 0.006+0+0.34 ms clock, 0.091+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 235 @0.391s 21%: 0.003+0+0.34 ms clock, 0.056+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 236 @0.391s 21%: 0.045+0+0.33 ms clock, 0.67+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 237 @0.393s 21%: 0.003+0+0.37 ms clock, 0.063+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 238 @0.394s 21%: 0.004+0+0.34 ms clock, 0.070+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 239 @0.395s 21%: 0.003+0+0.32 ms clock, 0.060+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 240 @0.396s 21%: 0.047+0+0.38 ms clock, 0.71+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 241 @0.398s 21%: 0.003+0+0.36 ms clock, 0.056+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 242 @0.398s 22%: 0.031+0+0.33 ms clock, 0.47+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 243 @0.400s 22%: 0.003+0+0.39 ms clock, 0.060+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 244 @0.400s 22%: 0.046+0+0.37 ms clock, 0.69+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 245 @0.402s 22%: 0.018+0+0.33 ms clock, 0.27+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 246 @0.403s 22%: 0.013+0+0.35 ms clock, 0.20+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 247 @0.404s 22%: 0.017+0+0.36 ms clock, 0.26+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 248 @0.405s 22%: 0.032+0+0.34 ms clock, 0.48+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 249 @0.407s 22%: 0.003+0+0.32 ms clock, 0.053+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 250 @0.407s 22%: 0.012+0+0.36 ms clock, 0.18+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 251 @0.409s 22%: 0.003+0+0.36 ms clock, 0.061+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 252 @0.409s 22%: 0.030+0+0.31 ms clock, 0.45+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 253 @0.411s 22%: 0.003+0+0.32 ms clock, 0.056+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 254 @0.412s 22%: 0.030+0+0.27 ms clock, 0.45+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 255 @0.413s 22%: 0.004+0+0.35 ms clock, 0.065+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 256 @0.414s 22%: 0.005+0+0.28 ms clock, 0.080+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 257 @0.415s 22%: 0.015+0+0.31 ms clock, 0.23+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 258 @0.416s 22%: 0.047+0+0.36 ms clock, 0.71+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 259 @0.418s 22%: 0.003+0+0.35 ms clock, 0.053+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 260 @0.418s 22%: 0.041+0+0.36 ms clock, 0.62+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 261 @0.420s 22%: 0.003+0+0.33 ms clock, 0.053+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 262 @0.420s 22%: 0.042+0+0.36 ms clock, 0.64+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 287632 HeapInuse: 499712 HeapObjects: 1777 HeapIdle 843776 HeapReleased 0 HeapSys 1343488
gc 263 @0.422s 22%: 0.002+0+0.34 ms clock, 0.037+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      1659 ns/op	      99 B/op	       2 allocs/op
gc 264 @0.422s 22%: 0.006+0+0.35 ms clock, 0.093+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250520 HeapInuse: 483328 HeapObjects: 1280 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 265 @0.425s 22%: 0.018+0+0.36 ms clock, 0.27+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndEraseWithPool1E3-16       	gc 266 @0.425s 22%: 0.004+0+0.38 ms clock, 0.071+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 267 @0.427s 22%: 0.003+0+0.32 ms clock, 0.060+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 268 @0.428s 22%: 0.042+0+0.34 ms clock, 0.63+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 269 @0.430s 22%: 0.004+0+0.31 ms clock, 0.061+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 270 @0.431s 22%: 0.004+0+0.32 ms clock, 0.060+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 271 @0.433s 22%: 0.003+0+0.34 ms clock, 0.060+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 272 @0.433s 22%: 0.003+0+0.35 ms clock, 0.050+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 273 @0.435s 22%: 0.003+0+0.36 ms clock, 0.060+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 274 @0.436s 22%: 0.044+0+0.33 ms clock, 0.66+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 275 @0.438s 22%: 0.003+0+0.58 ms clock, 0.061+0/0/0+9.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 276 @0.439s 22%: 0.042+0+0.31 ms clock, 0.63+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 277 @0.441s 22%: 0.003+0+0.33 ms clock, 0.062+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 278 @0.442s 22%: 0.014+0+0.32 ms clock, 0.22+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 279 @0.444s 22%: 0.023+0+0.36 ms clock, 0.35+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 280 @0.444s 22%: 0.041+0+0.34 ms clock, 0.62+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 281 @0.447s 22%: 0.024+0+0.34 ms clock, 0.36+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 282 @0.447s 22%: 0.028+0+0.28 ms clock, 0.43+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 283 @0.449s 22%: 0.014+0+0.31 ms clock, 0.22+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 284 @0.450s 22%: 0.064+0+0.34 ms clock, 0.97+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 285 @0.452s 22%: 0.013+0+0.35 ms clock, 0.20+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 286 @0.453s 22%: 0.043+0+0.32 ms clock, 0.64+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 287 @0.455s 22%: 0.003+0+0.33 ms clock, 0.060+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 288 @0.455s 22%: 0.045+0+0.33 ms clock, 0.68+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 289 @0.457s 22%: 0.003+0+0.33 ms clock, 0.057+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 290 @0.458s 22%: 0.042+0+0.31 ms clock, 0.63+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 291 @0.460s 22%: 0.004+0+0.35 ms clock, 0.064+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 292 @0.461s 22%: 0.043+0+0.37 ms clock, 0.64+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 293 @0.463s 22%: 0.003+0+0.33 ms clock, 0.056+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 294 @0.463s 22%: 0.004+0+1.3 ms clock, 0.064+0/0/0+19 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 250600 HeapInuse: 483328 HeapObjects: 1283 HeapIdle 860160 HeapReleased 0 HeapSys 1343488
gc 295 @0.466s 22%: 0.002+0+0.34 ms clock, 0.042+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2137 ns/op	      62 B/op	       1 allocs/op
gc 296 @0.467s 23%: 0.006+0+0.35 ms clock, 0.10+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246040 HeapInuse: 540672 HeapObjects: 365 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 297 @0.468s 23%: 0.003+0+0.34 ms clock, 0.060+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E3-16                   	gc 298 @0.469s 23%: 0.004+0+0.39 ms clock, 0.073+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246384 HeapInuse: 540672 HeapObjects: 371 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 299 @0.470s 23%: 0.020+0+0.33 ms clock, 0.31+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 300 @0.470s 23%: 0.011+0+0.31 ms clock, 0.17+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246744 HeapInuse: 540672 HeapObjects: 375 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 301 @0.471s 23%: 0.019+0+0.30 ms clock, 0.29+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 302 @0.471s 23%: 0.045+0+0.32 ms clock, 0.68+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246392 HeapInuse: 540672 HeapObjects: 370 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 303 @0.472s 23%: 0.004+0+0.33 ms clock, 0.064+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 304 @0.473s 23%: 0.043+0+0.28 ms clock, 0.65+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246440 HeapInuse: 540672 HeapObjects: 370 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 305 @0.474s 23%: 0.003+0+0.32 ms clock, 0.054+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 306 @0.474s 23%: 0.043+0+0.31 ms clock, 0.64+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246264 HeapInuse: 540672 HeapObjects: 370 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 307 @0.475s 23%: 0.002+0+0.35 ms clock, 0.034+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 308 @0.476s 23%: 0.012+0+0.32 ms clock, 0.18+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246352 HeapInuse: 540672 HeapObjects: 371 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 309 @0.476s 23%: 0.003+0+0.40 ms clock, 0.054+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 310 @0.477s 23%: 0.003+0+0.27 ms clock, 0.056+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245768 HeapInuse: 540672 HeapObjects: 363 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 311 @0.478s 23%: 0.003+0+0.38 ms clock, 0.055+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 312 @0.478s 23%: 0.043+0+0.35 ms clock, 0.64+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246560 HeapInuse: 540672 HeapObjects: 368 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 313 @0.479s 23%: 0.023+0+0.31 ms clock, 0.35+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 314 @0.480s 23%: 0.042+0+0.35 ms clock, 0.63+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245560 HeapInuse: 540672 HeapObjects: 357 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 315 @0.481s 23%: 0.019+0+0.32 ms clock, 0.29+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 316 @0.481s 23%: 0.041+0+0.31 ms clock, 0.61+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246904 HeapInuse: 540672 HeapObjects: 377 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 317 @0.482s 23%: 0.003+0+0.35 ms clock, 0.053+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 318 @0.483s 23%: 0.004+0+0.37 ms clock, 0.073+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246016 HeapInuse: 540672 HeapObjects: 368 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 319 @0.484s 23%: 0.001+0+0.35 ms clock, 0.029+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 320 @0.484s 23%: 0.004+0+0.34 ms clock, 0.061+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245688 HeapInuse: 540672 HeapObjects: 359 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 321 @0.485s 23%: 0.024+0+0.34 ms clock, 0.37+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 322 @0.486s 23%: 0.044+0+0.32 ms clock, 0.66+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 246840 HeapInuse: 540672 HeapObjects: 376 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 323 @0.486s 23%: 0.025+0+0.30 ms clock, 0.38+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 324 @0.487s 24%: 0.045+0+0.31 ms clock, 0.68+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245624 HeapInuse: 540672 HeapObjects: 360 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 325 @0.488s 24%: 0.002+0+0.32 ms clock, 0.035+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 326 @0.488s 24%: 0.003+0+0.32 ms clock, 0.059+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245656 HeapInuse: 540672 HeapObjects: 358 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 327 @0.489s 24%: 0.019+0+0.31 ms clock, 0.28+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       850 ns/op	      57 B/op	       0 allocs/op
gc 328 @0.490s 24%: 0.005+0+0.34 ms clock, 0.083+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254416 HeapInuse: 548864 HeapObjects: 366 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 329 @0.491s 24%: 0.004+0+0.33 ms clock, 0.067+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E3-16                    	gc 330 @0.491s 24%: 0.006+0+0.32 ms clock, 0.099+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254296 HeapInuse: 548864 HeapObjects: 366 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 331 @0.492s 24%: 0.019+0+0.33 ms clock, 0.28+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 332 @0.493s 24%: 0.003+0+0.32 ms clock, 0.052+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254208 HeapInuse: 548864 HeapObjects: 369 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 333 @0.494s 24%: 0.001+0+0.33 ms clock, 0.031+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 334 @0.494s 24%: 0.004+0+0.26 ms clock, 0.067+0/0/0+3.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254304 HeapInuse: 548864 HeapObjects: 366 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 335 @0.495s 24%: 0.003+0+0.31 ms clock, 0.054+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 336 @0.496s 24%: 0.003+0+0.31 ms clock, 0.058+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254776 HeapInuse: 548864 HeapObjects: 374 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 337 @0.497s 24%: 0.003+0+0.31 ms clock, 0.058+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 338 @0.497s 24%: 0.005+0+0.31 ms clock, 0.078+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254904 HeapInuse: 548864 HeapObjects: 375 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 339 @0.498s 24%: 0.003+0+0.32 ms clock, 0.055+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 340 @0.499s 24%: 0.005+0+0.40 ms clock, 0.082+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254016 HeapInuse: 548864 HeapObjects: 363 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 341 @0.500s 24%: 0.003+0+0.33 ms clock, 0.055+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 342 @0.500s 24%: 0.003+0+0.31 ms clock, 0.056+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254632 HeapInuse: 548864 HeapObjects: 371 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 343 @0.501s 24%: 0.003+0+0.37 ms clock, 0.055+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 344 @0.502s 24%: 0.046+0+0.31 ms clock, 0.69+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 253544 HeapInuse: 548864 HeapObjects: 359 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 345 @0.503s 24%: 0.022+0+0.31 ms clock, 0.34+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 346 @0.503s 24%: 0.003+0+0.34 ms clock, 0.057+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254200 HeapInuse: 548864 HeapObjects: 365 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 347 @0.504s 24%: 0.003+0+0.29 ms clock, 0.054+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 348 @0.505s 24%: 0.004+0+0.34 ms clock, 0.061+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254000 HeapInuse: 548864 HeapObjects: 361 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 349 @0.506s 24%: 0.019+0+0.32 ms clock, 0.29+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 350 @0.506s 24%: 0.044+0+0.30 ms clock, 0.67+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254952 HeapInuse: 548864 HeapObjects: 375 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 351 @0.507s 24%: 0.003+0+0.31 ms clock, 0.050+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 352 @0.508s 24%: 0.047+0+0.34 ms clock, 0.70+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254048 HeapInuse: 548864 HeapObjects: 367 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 353 @0.509s 24%: 0.003+0+0.32 ms clock, 0.060+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 354 @0.510s 24%: 0.004+0+0.32 ms clock, 0.060+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254096 HeapInuse: 548864 HeapObjects: 367 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 355 @0.511s 24%: 0.003+0+0.33 ms clock, 0.058+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 356 @0.511s 24%: 0.046+0+0.31 ms clock, 0.69+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 256040 HeapInuse: 557056 HeapObjects: 382 HeapIdle 786432 HeapReleased 0 HeapSys 1343488
gc 357 @0.512s 24%: 0.003+0+0.29 ms clock, 0.053+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       535 ns/op	      11 B/op	       0 allocs/op
gc 358 @0.513s 24%: 0.005+0+0.33 ms clock, 0.087+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196264 HeapInuse: 458752 HeapObjects: 299 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 359 @0.513s 25%: 0.003+0+0.46 ms clock, 0.056+0/0/0+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E3-16           	gc 360 @0.514s 25%: 0.013+0+0.27 ms clock, 0.20+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196680 HeapInuse: 458752 HeapObjects: 307 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 361 @0.515s 25%: 0.003+0+0.38 ms clock, 0.056+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 362 @0.515s 25%: 0.045+0+0.36 ms clock, 0.67+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196320 HeapInuse: 458752 HeapObjects: 301 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 363 @0.516s 25%: 0.003+0+0.35 ms clock, 0.060+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 364 @0.516s 25%: 0.004+0+0.32 ms clock, 0.063+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 195904 HeapInuse: 450560 HeapObjects: 297 HeapIdle 892928 HeapReleased 0 HeapSys 1343488
gc 365 @0.517s 25%: 0.003+0+0.33 ms clock, 0.062+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 366 @0.517s 25%: 0.005+0+0.33 ms clock, 0.089+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196688 HeapInuse: 458752 HeapObjects: 308 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 367 @0.518s 25%: 0.003+0+0.32 ms clock, 0.057+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 368 @0.519s 25%: 0.045+0+0.27 ms clock, 0.68+0/0/0+4.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196208 HeapInuse: 458752 HeapObjects: 299 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 369 @0.519s 25%: 0.002+0+0.34 ms clock, 0.032+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 370 @0.520s 25%: 0.006+0+0.36 ms clock, 0.095+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196568 HeapInuse: 458752 HeapObjects: 306 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 371 @0.520s 25%: 0.021+0+0.33 ms clock, 0.32+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 372 @0.521s 25%: 0.004+0+0.32 ms clock, 0.070+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196304 HeapInuse: 450560 HeapObjects: 305 HeapIdle 892928 HeapReleased 0 HeapSys 1343488
gc 373 @0.522s 25%: 0.003+0+0.36 ms clock, 0.059+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 374 @0.522s 25%: 0.004+0+0.31 ms clock, 0.060+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196592 HeapInuse: 458752 HeapObjects: 307 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 375 @0.523s 25%: 0.024+0+0.31 ms clock, 0.37+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 376 @0.523s 25%: 0.003+0+0.35 ms clock, 0.058+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196456 HeapInuse: 458752 HeapObjects: 304 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 377 @0.524s 25%: 0.016+0+0.32 ms clock, 0.25+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 378 @0.524s 25%: 0.031+0+0.36 ms clock, 0.46+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196112 HeapInuse: 458752 HeapObjects: 298 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 379 @0.525s 25%: 0.014+0+0.63 ms clock, 0.22+0/0/0+9.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 380 @0.526s 25%: 0.004+0+0.38 ms clock, 0.065+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196064 HeapInuse: 450560 HeapObjects: 301 HeapIdle 892928 HeapReleased 0 HeapSys 1343488
gc 381 @0.527s 25%: 0.003+0+0.33 ms clock, 0.063+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 382 @0.527s 25%: 0.046+0+0.29 ms clock, 0.69+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196392 HeapInuse: 458752 HeapObjects: 304 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 383 @0.528s 25%: 0.003+0+0.35 ms clock, 0.055+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 384 @0.528s 25%: 0.004+0+0.34 ms clock, 0.073+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 195928 HeapInuse: 450560 HeapObjects: 298 HeapIdle 892928 HeapReleased 0 HeapSys 1343488
gc 385 @0.529s 25%: 0.003+0+0.33 ms clock, 0.057+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 386 @0.529s 26%: 0.005+0+0.31 ms clock, 0.081+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 196400 HeapInuse: 458752 HeapObjects: 305 HeapIdle 884736 HeapReleased 0 HeapSys 1343488
gc 387 @0.530s 26%: 0.003+0+0.29 ms clock, 0.057+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 388 @0.531s 26%: 0.003+0+0.32 ms clock, 0.052+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 195888 HeapInuse: 450560 HeapObjects: 300 HeapIdle 892928 HeapReleased 0 HeapSys 1343488
gc 389 @0.531s 26%: 0.003+0+0.35 ms clock, 0.058+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       676 ns/op	       8 B/op	       0 allocs/op
gc 390 @0.532s 26%: 0.007+0+0.35 ms clock, 0.11+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216576 HeapInuse: 434176 HeapObjects: 272 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 391 @0.532s 26%: 0.003+0+0.31 ms clock, 0.053+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E3-16    	gc 392 @0.533s 26%: 0.007+0+0.34 ms clock, 0.10+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 393 @0.534s 26%: 0.001+0+0.32 ms clock, 0.030+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 394 @0.534s 26%: 0.043+0+0.34 ms clock, 0.65+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 395 @0.535s 26%: 0.003+0+0.31 ms clock, 0.055+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 396 @0.535s 26%: 0.045+0+0.32 ms clock, 0.67+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 397 @0.536s 26%: 0.002+0+0.32 ms clock, 0.035+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 398 @0.536s 26%: 0.046+0+0.30 ms clock, 0.69+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 399 @0.537s 26%: 0.003+0+0.34 ms clock, 0.057+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 400 @0.537s 26%: 0.005+0+0.34 ms clock, 0.085+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 401 @0.538s 26%: 0.017+0+0.35 ms clock, 0.26+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 402 @0.538s 26%: 0.047+0+0.31 ms clock, 0.70+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 403 @0.539s 26%: 0.014+0+0.31 ms clock, 0.22+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 404 @0.540s 26%: 0.049+0+0.33 ms clock, 0.73+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 405 @0.540s 26%: 0.020+0+0.33 ms clock, 0.30+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 406 @0.541s 26%: 0.004+0+0.31 ms clock, 0.067+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 407 @0.541s 26%: 0.003+0+0.38 ms clock, 0.053+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 408 @0.542s 26%: 0.004+0+0.30 ms clock, 0.065+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 409 @0.542s 26%: 0.003+0+0.32 ms clock, 0.053+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 410 @0.543s 26%: 0.004+0+0.31 ms clock, 0.066+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 411 @0.544s 26%: 0.003+0+0.32 ms clock, 0.054+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 412 @0.544s 26%: 0.003+0+0.43 ms clock, 0.053+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 413 @0.545s 26%: 0.003+0+0.33 ms clock, 0.054+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 414 @0.545s 26%: 0.003+0+0.34 ms clock, 0.053+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 415 @0.546s 26%: 0.003+0+0.32 ms clock, 0.055+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 416 @0.546s 27%: 0.004+0+0.41 ms clock, 0.070+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 417 @0.547s 27%: 0.003+0+0.35 ms clock, 0.052+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 418 @0.547s 27%: 0.043+0+0.53 ms clock, 0.65+0/0/0+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 216656 HeapInuse: 434176 HeapObjects: 275 HeapIdle 909312 HeapReleased 0 HeapSys 1343488
gc 419 @0.548s 27%: 0.026+0+0.36 ms clock, 0.40+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       650 ns/op	      28 B/op	       0 allocs/op
gc 420 @0.549s 27%: 0.008+0+0.36 ms clock, 0.12+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273056 HeapInuse: 491520 HeapObjects: 272 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 421 @0.552s 27%: 0.016+0+0.31 ms clock, 0.24+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSort1E4-16                               	gc 422 @0.552s 27%: 0.019+0+0.34 ms clock, 0.29+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 423 @0.555s 27%: 0.003+0+0.37 ms clock, 0.056+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 424 @0.555s 27%: 0.018+0+0.35 ms clock, 0.27+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 425 @0.558s 27%: 0.003+0+0.37 ms clock, 0.055+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 426 @0.559s 27%: 0.039+0+0.36 ms clock, 0.59+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 427 @0.561s 27%: 0.014+0+0.33 ms clock, 0.21+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 428 @0.562s 27%: 0.043+0+0.35 ms clock, 0.65+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 429 @0.565s 27%: 0.038+0+0.36 ms clock, 0.57+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 430 @0.565s 27%: 0.043+0+0.32 ms clock, 0.65+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 431 @0.568s 26%: 0.003+0+0.32 ms clock, 0.057+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 432 @0.568s 27%: 0.004+0+0.33 ms clock, 0.064+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 433 @0.571s 26%: 0.003+0+0.42 ms clock, 0.060+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 434 @0.572s 26%: 0.012+0+0.34 ms clock, 0.19+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 435 @0.574s 26%: 0.015+0+0.33 ms clock, 0.23+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 436 @0.575s 26%: 0.027+0+0.30 ms clock, 0.41+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 437 @0.578s 26%: 0.007+0+0.39 ms clock, 0.11+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 438 @0.578s 26%: 0.046+0+0.38 ms clock, 0.69+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 439 @0.581s 26%: 0.003+0+0.42 ms clock, 0.058+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 440 @0.582s 26%: 0.045+0+0.40 ms clock, 0.67+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 441 @0.585s 26%: 0.003+0+0.38 ms clock, 0.056+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 442 @0.585s 26%: 0.004+0+0.39 ms clock, 0.062+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 443 @0.588s 26%: 0.017+0+0.43 ms clock, 0.25+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 444 @0.588s 26%: 0.044+0+0.37 ms clock, 0.66+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 445 @0.591s 26%: 0.003+0+0.34 ms clock, 0.045+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 446 @0.592s 26%: 0.032+0+0.35 ms clock, 0.48+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 273120 HeapInuse: 491520 HeapObjects: 275 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 447 @0.594s 26%: 0.003+0+0.37 ms clock, 0.055+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       276 ns/op	       8 B/op	       0 allocs/op
gc 448 @0.595s 26%: 0.006+0+0.35 ms clock, 0.092+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071376 HeapInuse: 1286144 HeapObjects: 15274 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 449 @0.628s 25%: 0.002+0+0.46 ms clock, 0.038+0/0/0+7.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSetInsert1E4-16                          	gc 450 @0.629s 25%: 0.004+0+0.39 ms clock, 0.069+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 451 @0.661s 24%: 0.043+0+0.77 ms clock, 0.64+0/0/0+11 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 452 @0.662s 24%: 0.041+0+0.65 ms clock, 0.62+0/0/0+9.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 453 @0.695s 23%: 0.015+0+0.48 ms clock, 0.23+0/0/0+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 454 @0.696s 23%: 0.049+0+0.39 ms clock, 0.74+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 455 @0.728s 22%: 0.036+0+0.45 ms clock, 0.55+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 456 @0.729s 22%: 0.059+0+0.37 ms clock, 0.89+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 457 @0.761s 21%: 0.015+0+0.45 ms clock, 0.23+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 458 @0.762s 21%: 0.082+0+0.30 ms clock, 1.2+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 459 @0.794s 20%: 0.003+0+0.40 ms clock, 0.056+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 460 @0.795s 20%: 0.046+0+0.37 ms clock, 0.69+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 461 @0.827s 20%: 0.040+0+0.48 ms clock, 0.60+0/0/0+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 462 @0.828s 20%: 0.045+0+0.41 ms clock, 0.67+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 463 @0.862s 19%: 0.003+0+0.41 ms clock, 0.058+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 464 @0.863s 19%: 0.041+0+0.29 ms clock, 0.62+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 465 @0.896s 18%: 0.048+0+0.45 ms clock, 0.72+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 466 @0.897s 18%: 0.050+0+0.39 ms clock, 0.75+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 467 @0.930s 18%: 0.003+0+0.38 ms clock, 0.055+0/0/0+6.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 468 @0.930s 18%: 0.039+0+0.40 ms clock, 0.59+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 469 @0.963s 17%: 0.042+0+0.47 ms clock, 0.63+0/0/0+7.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 470 @0.964s 17%: 0.052+0+0.39 ms clock, 0.78+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 471 @0.997s 17%: 0.036+0+0.42 ms clock, 0.55+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 472 @0.997s 17%: 0.042+0+0.41 ms clock, 0.63+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 473 @1.030s 16%: 0.016+0+0.45 ms clock, 0.25+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 474 @1.031s 16%: 0.038+0+0.38 ms clock, 0.58+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 475 @1.064s 16%: 0.055+0+0.45 ms clock, 0.83+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 476 @1.064s 16%: 0.083+0+0.40 ms clock, 1.2+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 477 @1.097s 15%: 0.015+0+0.46 ms clock, 0.23+0/0/0+6.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 478 @1.097s 15%: 0.058+0+0.34 ms clock, 0.88+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 479 @1.130s 15%: 0.044+0+0.39 ms clock, 0.66+0/0/0+5.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 480 @1.130s 15%: 0.042+0+0.34 ms clock, 0.63+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071440 HeapInuse: 1286144 HeapObjects: 15277 HeapIdle 57344 HeapReleased 0 HeapSys 1343488
gc 481 @1.163s 14%: 0.037+0+0.45 ms clock, 0.56+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
   10000	      3237 ns/op	      88 B/op	       2 allocs/op
gc 482 @1.163s 15%: 0.019+0+0.37 ms clock, 0.28+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033296 HeapInuse: 2252800 HeapObjects: 30275 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 483 @1.210s 14%: 0.056+0+0.43 ms clock, 0.85+0/0/0+6.4 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
#BenchmarkSetErase1E4-16                           	gc 484 @1.211s 14%: 0.048+0+0.33 ms clock, 0.72+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 485 @1.243s 14%: 0.003+0+0.49 ms clock, 0.048+0/0/0+7.9 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 486 @1.243s 14%: 0.038+0+0.39 ms clock, 0.57+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 487 @1.276s 13%: 0.003+0+0.44 ms clock, 0.058+0/0/0+7.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 488 @1.276s 13%: 0.059+0+0.34 ms clock, 0.88+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 489 @1.308s 13%: 0.011+0+0.43 ms clock, 0.17+0/0/0+6.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 490 @1.308s 13%: 0.087+0+0.35 ms clock, 1.3+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 491 @1.340s 13%: 0.002+0+0.42 ms clock, 0.045+0/0/0+6.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 492 @1.341s 13%: 0.002+0+0.32 ms clock, 0.044+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 493 @1.373s 13%: 0.002+0+0.43 ms clock, 0.045+0/0/0+7.0 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 494 @1.373s 13%: 0.026+0+0.34 ms clock, 0.40+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 495 @1.405s 12%: 0.012+0+0.44 ms clock, 0.19+0/0/0+6.7 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 496 @1.406s 12%: 0.051+0+0.30 ms clock, 0.77+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 497 @1.438s 12%: 0.016+0+0.47 ms clock, 0.24+0/0/0+7.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 498 @1.438s 12%: 0.053+0+0.35 ms clock, 0.80+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 499 @1.470s 12%: 0.020+0+0.47 ms clock, 0.31+0/0/0+7.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 500 @1.471s 12%: 0.045+0+0.34 ms clock, 0.68+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 501 @1.503s 12%: 0.020+0+0.44 ms clock, 0.31+0/0/0+6.6 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 502 @1.503s 12%: 0.052+0+0.35 ms clock, 0.79+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 503 @1.536s 11%: 0.003+0+0.47 ms clock, 0.055+0/0/0+7.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 504 @1.536s 11%: 0.021+0+0.34 ms clock, 0.32+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 505 @1.568s 11%: 0.053+0+0.48 ms clock, 0.79+0/0/0+7.2 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 506 @1.569s 11%: 0.014+0+0.36 ms clock, 0.22+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 507 @1.601s 11%: 0.016+0+0.44 ms clock, 0.24+0/0/0+6.6 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 508 @1.601s 11%: 0.071+0+0.37 ms clock, 1.0+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 509 @1.633s 11%: 0.039+0+0.50 ms clock, 0.59+0/0/0+7.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 510 @1.634s 11%: 0.048+0+0.42 ms clock, 0.72+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 511 @1.666s 11%: 0.048+0+0.49 ms clock, 0.72+0/0/0+7.4 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 512 @1.667s 11%: 0.011+0+0.39 ms clock, 0.16+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033360 HeapInuse: 2252800 HeapObjects: 30278 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 513 @1.699s 11%: 0.022+0+0.45 ms clock, 0.33+0/0/0+6.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
   10000	      1249 ns/op	      96 B/op	       2 allocs/op
gc 514 @1.699s 11%: 0.009+0+0.36 ms clock, 0.14+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153296 HeapInuse: 1368064 HeapObjects: 15275 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 515 @1.712s 11%: 0.083+0+0.38 ms clock, 1.2+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E4-16                  	gc 516 @1.713s 11%: 0.035+0+0.38 ms clock, 0.52+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 517 @1.726s 10%: 0.001+0+0.43 ms clock, 0.018+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 518 @1.726s 10%: 0.022+0+0.34 ms clock, 0.33+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 519 @1.739s 10%: 0.053+0+0.52 ms clock, 0.80+0/0/0+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 520 @1.740s 10%: 0.008+0+0.39 ms clock, 0.13+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 521 @1.753s 10%: 0.017+0+0.39 ms clock, 0.26+0/0/0+5.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 522 @1.754s 10%: 0.032+0+0.36 ms clock, 0.49+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 523 @1.767s 10%: 0.001+0+0.40 ms clock, 0.020+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 524 @1.767s 10%: 0.009+0+0.35 ms clock, 0.14+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 525 @1.780s 10%: 0.018+0+0.38 ms clock, 0.27+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 526 @1.781s 10%: 0.010+0+0.36 ms clock, 0.15+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 527 @1.795s 10%: 0.003+0+0.38 ms clock, 0.048+0/0/0+6.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 528 @1.796s 10%: 0.011+0+0.28 ms clock, 0.17+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 529 @1.809s 10%: 0.048+0+0.34 ms clock, 0.72+0/0/0+5.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 530 @1.809s 10%: 0.070+0+0.38 ms clock, 1.0+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 531 @1.822s 10%: 0.057+0+0.51 ms clock, 0.86+0/0/0+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 532 @1.823s 10%: 0.057+0+0.39 ms clock, 0.86+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 533 @1.837s 10%: 0.019+0+0.35 ms clock, 0.29+0/0/0+5.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 534 @1.837s 10%: 0.008+0+0.37 ms clock, 0.12+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 535 @1.852s 10%: 0.003+0+0.35 ms clock, 0.049+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 536 @1.852s 10%: 0.003+0+0.32 ms clock, 0.053+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 537 @1.865s 10%: 0.012+0+0.41 ms clock, 0.18+0/0/0+6.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 538 @1.866s 10%: 0.003+0+0.27 ms clock, 0.052+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 539 @1.879s 10%: 0.018+0+0.43 ms clock, 0.28+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 540 @1.880s 10%: 0.040+0+0.36 ms clock, 0.61+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 541 @1.893s 10%: 0.002+0+0.41 ms clock, 0.046+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 542 @1.893s 10%: 0.010+0+0.40 ms clock, 0.15+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 543 @1.906s 10%: 0.013+0+0.40 ms clock, 0.19+0/0/0+6.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 544 @1.906s 10%: 0.057+0+0.38 ms clock, 0.86+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15278 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 545 @1.919s 10%: 0.017+0+0.40 ms clock, 0.25+0/0/0+6.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
   10000	      1282 ns/op	      96 B/op	       2 allocs/op
gc 546 @1.920s 10%: 0.002+0+0.34 ms clock, 0.043+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756264 HeapInuse: 991232 HeapObjects: 10281 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 547 @1.939s 10%: 0.037+0+0.40 ms clock, 0.56+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E4-16          	gc 548 @1.939s 10%: 0.009+0+0.34 ms clock, 0.14+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 549 @1.956s 10%: 0.033+0+0.43 ms clock, 0.50+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 550 @1.957s 10%: 0.045+0+0.36 ms clock, 0.68+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 551 @1.974s 10%: 0.004+0+0.40 ms clock, 0.067+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 552 @1.974s 10%: 0.008+0+0.34 ms clock, 0.12+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 553 @1.991s 10%: 0.004+0+0.39 ms clock, 0.065+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 554 @1.992s 10%: 0.012+0+0.34 ms clock, 0.19+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 555 @2.008s 10%: 0.022+0+0.37 ms clock, 0.34+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 556 @2.009s 10%: 0.054+0+0.35 ms clock, 0.81+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 557 @2.025s 10%: 0.059+0+0.39 ms clock, 0.89+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 558 @2.026s 10%: 0.014+0+0.39 ms clock, 0.22+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 559 @2.042s 10%: 0.010+0+0.38 ms clock, 0.15+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 560 @2.042s 10%: 0.025+0+0.34 ms clock, 0.38+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 561 @2.058s 10%: 0.064+0+0.41 ms clock, 0.96+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 562 @2.059s 10%: 0.075+0+0.31 ms clock, 1.1+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 563 @2.075s 9%: 0.018+0+0.35 ms clock, 0.27+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 564 @2.075s 9%: 0.044+0+0.34 ms clock, 0.66+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 565 @2.091s 9%: 0.046+0+0.40 ms clock, 0.69+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 566 @2.092s 9%: 0.025+0+0.32 ms clock, 0.38+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 567 @2.108s 9%: 0.060+0+0.40 ms clock, 0.91+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 568 @2.109s 9%: 0.015+0+0.35 ms clock, 0.22+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 569 @2.125s 9%: 0.055+0+0.36 ms clock, 0.82+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 570 @2.125s 9%: 0.057+0+0.38 ms clock, 0.86+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 571 @2.141s 9%: 0.018+0+0.32 ms clock, 0.28+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 572 @2.141s 9%: 0.049+0+0.32 ms clock, 0.74+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 573 @2.158s 9%: 0.021+0+0.33 ms clock, 0.32+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 574 @2.158s 9%: 0.027+0+0.30 ms clock, 0.40+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 575 @2.174s 9%: 0.017+0+0.34 ms clock, 0.26+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 576 @2.174s 9%: 0.022+0+0.34 ms clock, 0.34+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756344 HeapInuse: 991232 HeapObjects: 10284 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 577 @2.190s 9%: 0.034+0+0.36 ms clock, 0.52+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	      1595 ns/op	      56 B/op	       1 allocs/op
gc 578 @2.191s 9%: 0.016+0+0.54 ms clock, 0.24+0/0/0+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071360 HeapInuse: 1286144 HeapObjects: 15273 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 579 @2.206s 9%: 0.002+0+0.43 ms clock, 0.034+0/0/0+6.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkIntSetInsert1E4-16                       	gc 580 @2.207s 9%: 0.012+0+0.37 ms clock, 0.19+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 581 @2.226s 9%: 0.044+0+0.54 ms clock, 0.66+0/0/0+8.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 582 @2.227s 9%: 0.036+0+0.38 ms clock, 0.55+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 583 @2.247s 9%: 0.004+0+0.51 ms clock, 0.065+0/0/0+7.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 584 @2.248s 9%: 0.023+0+0.39 ms clock, 0.35+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 585 @2.268s 9%: 0.049+0+0.42 ms clock, 0.73+0/0/0+6.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 586 @2.268s 9%: 0.039+0+0.35 ms clock, 0.59+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 587 @2.288s 9%: 0.045+0+0.41 ms clock, 0.68+0/0/0+6.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 588 @2.289s 9%: 0.004+0+1.1 ms clock, 0.062+0/0/0+17 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 589 @2.309s 9%: 0.027+0+0.44 ms clock, 0.41+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 590 @2.310s 9%: 0.052+0+0.38 ms clock, 0.79+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 591 @2.330s 9%: 0.002+0+0.44 ms clock, 0.039+0/0/0+7.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 592 @2.330s 9%: 0.043+0+0.35 ms clock, 0.64+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 593 @2.350s 9%: 0.002+0+0.44 ms clock, 0.038+0/0/0+7.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 594 @2.351s 9%: 0.072+0+0.37 ms clock, 1.0+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 595 @2.371s 9%: 0.006+0+0.43 ms clock, 0.099+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 596 @2.371s 9%: 0.046+0+0.37 ms clock, 0.69+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 597 @2.391s 9%: 0.002+0+0.42 ms clock, 0.037+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 598 @2.391s 9%: 0.013+0+0.33 ms clock, 0.20+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 599 @2.411s 9%: 0.005+0+0.43 ms clock, 0.080+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 600 @2.412s 9%: 0.048+0+0.38 ms clock, 0.72+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 601 @2.432s 9%: 0.004+0+0.48 ms clock, 0.075+0/0/0+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 602 @2.433s 9%: 0.006+0+0.35 ms clock, 0.093+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 603 @2.453s 9%: 0.035+0+0.44 ms clock, 0.52+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 604 @2.453s 9%: 0.065+0+0.46 ms clock, 0.98+0/0/0+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 605 @2.474s 9%: 0.005+0+0.72 ms clock, 0.080+0/0/0+11 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 606 @2.474s 9%: 0.004+0+0.37 ms clock, 0.073+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 607 @2.494s 9%: 0.040+0+0.50 ms clock, 0.60+0/0/0+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 608 @2.495s 9%: 0.060+0+0.40 ms clock, 0.90+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1071424 HeapInuse: 1286144 HeapObjects: 15276 HeapIdle 1105920 HeapReleased 0 HeapSys 2392064
gc 609 @2.515s 9%: 0.002+0+0.43 ms clock, 0.042+0/0/0+6.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
   10000	      1998 ns/op	      88 B/op	       2 allocs/op
gc 610 @2.516s 9%: 0.017+0+0.39 ms clock, 0.25+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033280 HeapInuse: 2252800 HeapObjects: 30274 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 611 @2.550s 8%: 0.023+0+0.49 ms clock, 0.34+0/0/0+7.4 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
#BenchmarkIntSetErase1E4-16                        	gc 612 @2.551s 8%: 0.017+0+0.39 ms clock, 0.26+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 613 @2.585s 8%: 0.003+0+0.48 ms clock, 0.048+0/0/0+7.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 614 @2.586s 8%: 0.075+0+0.41 ms clock, 1.1+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 615 @2.620s 8%: 0.022+0+0.51 ms clock, 0.33+0/0/0+7.7 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 616 @2.621s 8%: 0.073+0+0.38 ms clock, 1.0+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 617 @2.655s 8%: 0.021+0+0.53 ms clock, 0.32+0/0/0+8.0 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 618 @2.655s 8%: 0.053+0+0.37 ms clock, 0.79+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 619 @2.689s 8%: 0.003+0+0.52 ms clock, 0.049+0/0/0+8.3 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 620 @2.690s 8%: 0.077+0+0.47 ms clock, 1.1+0/0/0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 621 @2.722s 8%: 0.002+0+0.65 ms clock, 0.044+0/0/0+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 622 @2.723s 8%: 0.078+0+0.38 ms clock, 1.1+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 623 @2.758s 8%: 0.029+0+0.54 ms clock, 0.43+0/0/0+8.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 624 @2.759s 8%: 0.060+0+0.39 ms clock, 0.90+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 625 @2.794s 8%: 0.025+0+0.56 ms clock, 0.37+0/0/0+8.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 626 @2.795s 8%: 0.073+0+0.42 ms clock, 1.1+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 627 @2.829s 8%: 0.044+0+0.51 ms clock, 0.66+0/0/0+7.7 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 628 @2.830s 8%: 0.043+0+0.39 ms clock, 0.65+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 629 @2.864s 8%: 0.004+0+0.50 ms clock, 0.072+0/0/0+8.0 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 630 @2.865s 8%: 0.039+0+0.68 ms clock, 0.59+0/0/0+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 631 @2.901s 8%: 0.005+0+0.44 ms clock, 0.084+0/0/0+7.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 632 @2.901s 8%: 0.005+0+0.41 ms clock, 0.089+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 633 @2.935s 8%: 0.037+0+0.48 ms clock, 0.56+0/0/0+7.3 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 634 @2.936s 8%: 0.042+0+0.39 ms clock, 0.63+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 635 @2.970s 8%: 0.002+0+0.46 ms clock, 0.037+0/0/0+7.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 636 @2.970s 8%: 0.043+0+0.41 ms clock, 0.64+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 637 @3.005s 7%: 0.005+0+0.59 ms clock, 0.088+0/0/0+9.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 638 @3.005s 8%: 0.032+0+0.37 ms clock, 0.48+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 639 @3.039s 7%: 0.002+0+0.56 ms clock, 0.033+0/0/0+9.0 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 640 @3.040s 7%: 0.003+0+0.47 ms clock, 0.059+0/0/0+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2033344 HeapInuse: 2252800 HeapObjects: 30277 HeapIdle 139264 HeapReleased 0 HeapSys 2392064
gc 641 @3.074s 7%: 0.032+0+0.53 ms clock, 0.48+0/0/0+7.9 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
   10000	      1479 ns/op	      96 B/op	       2 allocs/op
gc 642 @3.075s 7%: 0.007+0+0.39 ms clock, 0.11+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153280 HeapInuse: 1368064 HeapObjects: 15274 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 643 @3.090s 7%: 0.032+0+0.39 ms clock, 0.49+0/0/0+5.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndErase1E4-16               	gc 644 @3.090s 7%: 0.005+0+0.36 ms clock, 0.084+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 645 @3.105s 7%: 0.041+0+0.43 ms clock, 0.62+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 646 @3.106s 7%: 0.038+0+0.39 ms clock, 0.58+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 647 @3.120s 7%: 0.050+0+0.46 ms clock, 0.75+0/0/0+6.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 648 @3.121s 7%: 0.043+0+0.38 ms clock, 0.65+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 649 @3.136s 7%: 0.008+0+0.64 ms clock, 0.12+0/0/0+9.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 650 @3.137s 7%: 0.044+0+0.54 ms clock, 0.67+0/0/0+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 651 @3.152s 7%: 0.044+0+0.45 ms clock, 0.66+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 652 @3.152s 7%: 0.085+0+0.34 ms clock, 1.2+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 653 @3.167s 7%: 0.001+0+0.57 ms clock, 0.031+0/0/0+9.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 654 @3.167s 7%: 0.045+0+0.41 ms clock, 0.67+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 655 @3.182s 7%: 0.043+0+0.52 ms clock, 0.65+0/0/0+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 656 @3.183s 7%: 0.040+0+0.48 ms clock, 0.60+0/0/0+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 657 @3.198s 7%: 0.051+0+0.43 ms clock, 0.77+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 658 @3.198s 7%: 0.004+0+0.37 ms clock, 0.063+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 659 @3.214s 7%: 0.034+0+0.44 ms clock, 0.51+0/0/0+6.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 660 @3.214s 7%: 0.077+0+0.39 ms clock, 1.1+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 661 @3.229s 7%: 0.002+0+0.47 ms clock, 0.038+0/0/0+7.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 662 @3.229s 7%: 0.044+0+0.48 ms clock, 0.66+0/0/0+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 663 @3.244s 7%: 0.041+0+0.44 ms clock, 0.62+0/0/0+6.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 664 @3.245s 7%: 0.073+0+0.34 ms clock, 1.1+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 665 @3.260s 7%: 0.001+0+0.43 ms clock, 0.030+0/0/0+6.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 666 @3.260s 7%: 0.045+0+0.33 ms clock, 0.68+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 667 @3.275s 7%: 0.045+0+0.42 ms clock, 0.68+0/0/0+6.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 668 @3.276s 7%: 0.057+0+0.36 ms clock, 0.86+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 669 @3.291s 7%: 0.033+0+0.42 ms clock, 0.49+0/0/0+6.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 670 @3.291s 7%: 0.082+0+0.37 ms clock, 1.2+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 671 @3.306s 7%: 0.006+0+0.41 ms clock, 0.10+0/0/0+6.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 672 @3.306s 7%: 0.011+0+0.36 ms clock, 0.17+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1153360 HeapInuse: 1368064 HeapObjects: 15277 HeapIdle 1024000 HeapReleased 0 HeapSys 2392064
gc 673 @3.321s 7%: 0.004+0+0.40 ms clock, 0.065+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
   10000	      1467 ns/op	      96 B/op	       2 allocs/op
gc 674 @3.322s 7%: 0.004+0+0.35 ms clock, 0.072+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756248 HeapInuse: 991232 HeapObjects: 10280 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 675 @3.342s 7%: 0.009+0+0.43 ms clock, 0.14+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndEraseWithPool1E4-16       	gc 676 @3.342s 7%: 0.044+0+0.36 ms clock, 0.66+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 677 @3.362s 7%: 0.020+0+0.40 ms clock, 0.30+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 678 @3.363s 7%: 0.043+0+0.38 ms clock, 0.65+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 679 @3.382s 7%: 0.013+0+0.43 ms clock, 0.20+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 680 @3.383s 7%: 0.039+0+0.36 ms clock, 0.59+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 681 @3.402s 7%: 0.044+0+0.42 ms clock, 0.67+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 682 @3.403s 7%: 0.093+0+0.37 ms clock, 1.4+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 683 @3.423s 7%: 0.049+0+0.43 ms clock, 0.73+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 684 @3.423s 7%: 0.076+0+0.39 ms clock, 1.1+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 685 @3.443s 7%: 0.017+0+0.43 ms clock, 0.25+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 686 @3.444s 7%: 0.047+0+0.40 ms clock, 0.70+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 687 @3.463s 7%: 0.015+0+0.43 ms clock, 0.23+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 688 @3.464s 7%: 0.072+0+0.38 ms clock, 1.0+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 689 @3.484s 7%: 0.016+0+0.39 ms clock, 0.25+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 690 @3.484s 7%: 0.042+0+0.36 ms clock, 0.63+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 691 @3.504s 7%: 0.015+0+0.42 ms clock, 0.23+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 692 @3.504s 7%: 0.076+0+0.38 ms clock, 1.1+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 693 @3.524s 7%: 0.043+0+0.72 ms clock, 0.64+0/0/0+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 694 @3.525s 7%: 0.013+0+0.41 ms clock, 0.20+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 695 @3.544s 7%: 0.004+0+0.40 ms clock, 0.070+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 696 @3.545s 7%: 0.044+0+0.39 ms clock, 0.66+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 697 @3.564s 7%: 0.060+0+0.37 ms clock, 0.91+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 698 @3.565s 7%: 0.043+0+0.34 ms clock, 0.64+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 699 @3.585s 7%: 0.040+0+0.43 ms clock, 0.61+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 700 @3.585s 7%: 0.076+0+0.37 ms clock, 1.1+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 701 @3.605s 7%: 0.049+0+0.40 ms clock, 0.73+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 702 @3.605s 7%: 0.059+0+0.37 ms clock, 0.89+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 703 @3.625s 7%: 0.005+0+0.41 ms clock, 0.086+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 704 @3.626s 7%: 0.077+0+0.37 ms clock, 1.1+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 756328 HeapInuse: 991232 HeapObjects: 10283 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 705 @3.645s 7%: 0.046+0+0.41 ms clock, 0.70+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	      1977 ns/op	      56 B/op	       1 allocs/op
gc 706 @3.646s 7%: 0.007+0+0.36 ms clock, 0.11+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 623152 HeapInuse: 933888 HeapObjects: 902 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 707 @3.650s 7%: 0.016+0+0.40 ms clock, 0.24+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E4-16                   	gc 708 @3.650s 7%: 0.007+0+0.37 ms clock, 0.10+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619056 HeapInuse: 933888 HeapObjects: 862 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 709 @3.654s 7%: 0.048+0+0.34 ms clock, 0.72+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 710 @3.655s 7%: 0.012+0+0.39 ms clock, 0.18+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 621520 HeapInuse: 933888 HeapObjects: 888 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 711 @3.658s 7%: 0.012+0+0.36 ms clock, 0.19+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 712 @3.659s 7%: 0.004+0+0.36 ms clock, 0.066+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 623880 HeapInuse: 933888 HeapObjects: 911 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 713 @3.662s 7%: 0.013+0+0.38 ms clock, 0.19+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 714 @3.663s 7%: 0.085+0+0.45 ms clock, 1.2+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619904 HeapInuse: 933888 HeapObjects: 867 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 715 @3.667s 7%: 0.017+0+0.39 ms clock, 0.25+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 716 @3.668s 7%: 0.044+0+0.33 ms clock, 0.67+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 622064 HeapInuse: 933888 HeapObjects: 894 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 717 @3.672s 7%: 0.023+0+0.41 ms clock, 0.35+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 718 @3.672s 7%: 0.016+0+0.34 ms clock, 0.24+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 621936 HeapInuse: 933888 HeapObjects: 894 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 719 @3.676s 7%: 0.014+0+0.39 ms clock, 0.21+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 720 @3.676s 7%: 0.076+0+0.39 ms clock, 1.1+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 628952 HeapInuse: 942080 HeapObjects: 942 HeapIdle 1449984 HeapReleased 0 HeapSys 2392064
gc 721 @3.680s 7%: 0.046+0+0.38 ms clock, 0.69+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 722 @3.681s 7%: 0.044+0+0.36 ms clock, 0.66+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 625136 HeapInuse: 933888 HeapObjects: 905 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 723 @3.685s 7%: 0.017+0+0.42 ms clock, 0.26+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 724 @3.686s 7%: 0.046+0+0.37 ms clock, 0.69+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 621496 HeapInuse: 933888 HeapObjects: 887 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 725 @3.689s 7%: 0.017+0+0.38 ms clock, 0.26+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 726 @3.690s 7%: 0.063+0+0.38 ms clock, 0.95+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 627488 HeapInuse: 933888 HeapObjects: 930 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 727 @3.694s 7%: 0.002+0+0.33 ms clock, 0.037+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 728 @3.694s 7%: 0.062+0+0.35 ms clock, 0.93+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 625072 HeapInuse: 933888 HeapObjects: 911 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 729 @3.698s 7%: 0.023+0+0.36 ms clock, 0.34+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 730 @3.698s 7%: 0.044+0+0.35 ms clock, 0.66+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619504 HeapInuse: 933888 HeapObjects: 867 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 731 @3.702s 7%: 0.024+0+0.34 ms clock, 0.36+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 732 @3.703s 7%: 0.047+0+0.35 ms clock, 0.71+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 623080 HeapInuse: 933888 HeapObjects: 902 HeapIdle 1458176 HeapReleased 0 HeapSys 2392064
gc 733 @3.706s 7%: 0.018+0+0.36 ms clock, 0.27+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       339 ns/op	      43 B/op	       0 allocs/op
gc 734 @3.706s 7%: 0.004+0+0.33 ms clock, 0.069+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 709088 HeapInuse: 1015808 HeapObjects: 925 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 735 @3.711s 7%: 0.002+0+0.32 ms clock, 0.032+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E4-16                    	gc 736 @3.711s 7%: 0.004+0+0.38 ms clock, 0.070+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 702616 HeapInuse: 1015808 HeapObjects: 882 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 737 @3.716s 7%: 0.016+0+0.38 ms clock, 0.24+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 738 @3.717s 7%: 0.052+0+0.38 ms clock, 0.78+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704968 HeapInuse: 1015808 HeapObjects: 907 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 739 @3.721s 7%: 0.028+0+0.36 ms clock, 0.42+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 740 @3.722s 7%: 0.042+0+0.36 ms clock, 0.63+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 705016 HeapInuse: 1015808 HeapObjects: 904 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 741 @3.726s 7%: 0.024+0+0.62 ms clock, 0.37+0/0/0+9.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 742 @3.727s 7%: 0.003+0+0.38 ms clock, 0.054+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704576 HeapInuse: 1015808 HeapObjects: 897 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 743 @3.732s 7%: 0.004+0+0.40 ms clock, 0.073+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 744 @3.732s 7%: 0.064+0+0.36 ms clock, 0.96+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 706056 HeapInuse: 1015808 HeapObjects: 919 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 745 @3.737s 7%: 0.009+0+0.38 ms clock, 0.14+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 746 @3.737s 7%: 0.049+0+0.37 ms clock, 0.73+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 703256 HeapInuse: 1015808 HeapObjects: 884 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 747 @3.742s 7%: 0.004+0+0.35 ms clock, 0.068+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 748 @3.742s 7%: 0.045+0+0.33 ms clock, 0.67+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 708008 HeapInuse: 1015808 HeapObjects: 918 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 749 @3.747s 7%: 0.027+0+0.42 ms clock, 0.40+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 750 @3.747s 7%: 0.043+0+0.39 ms clock, 0.65+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704648 HeapInuse: 1015808 HeapObjects: 904 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 751 @3.752s 7%: 0.004+0+0.52 ms clock, 0.065+0/0/0+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 752 @3.752s 7%: 0.041+0+0.34 ms clock, 0.62+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 706216 HeapInuse: 1015808 HeapObjects: 916 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 753 @3.757s 7%: 0.004+0+0.35 ms clock, 0.064+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 754 @3.757s 7%: 0.043+0+0.41 ms clock, 0.64+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 703960 HeapInuse: 1015808 HeapObjects: 896 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 755 @3.762s 7%: 0.013+0+0.37 ms clock, 0.19+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 756 @3.762s 7%: 0.004+0+0.49 ms clock, 0.062+0/0/0+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 708632 HeapInuse: 1015808 HeapObjects: 922 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 757 @3.767s 7%: 0.002+0+0.52 ms clock, 0.033+0/0/0+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 758 @3.767s 7%: 0.004+0+0.32 ms clock, 0.061+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 702904 HeapInuse: 1015808 HeapObjects: 882 HeapIdle 1376256 HeapReleased 0 HeapSys 2392064
gc 759 @3.772s 7%: 0.002+0+0.33 ms clock, 0.034+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       139 ns/op	       8 B/op	       0 allocs/op
gc 760 @3.772s 7%: 0.003+0+0.34 ms clock, 0.047+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 256144 HeapInuse: 557056 HeapObjects: 382 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 761 @3.774s 7%: 0.022+0+0.32 ms clock, 0.34+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E4-16           	gc 762 @3.774s 7%: 0.015+0+0.30 ms clock, 0.23+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 256536 HeapInuse: 557056 HeapObjects: 389 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 763 @3.776s 7%: 0.002+0+0.34 ms clock, 0.036+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 764 @3.777s 7%: 0.039+0+0.50 ms clock, 0.59+0/0/0+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 258520 HeapInuse: 557056 HeapObjects: 408 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 765 @3.779s 7%: 0.004+0+0.35 ms clock, 0.065+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 766 @3.779s 7%: 0.015+0+0.31 ms clock, 0.23+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 257392 HeapInuse: 557056 HeapObjects: 401 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 767 @3.781s 7%: 0.004+0+0.34 ms clock, 0.065+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 768 @3.781s 7%: 0.004+0+0.35 ms clock, 0.063+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 257136 HeapInuse: 557056 HeapObjects: 393 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 769 @3.783s 7%: 0.019+0+0.33 ms clock, 0.29+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 770 @3.784s 7%: 0.004+0+0.36 ms clock, 0.061+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255536 HeapInuse: 557056 HeapObjects: 378 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 771 @3.786s 7%: 0.043+0+0.36 ms clock, 0.64+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 772 @3.786s 7%: 0.004+0+0.36 ms clock, 0.063+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 257072 HeapInuse: 557056 HeapObjects: 397 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 773 @3.788s 7%: 0.002+0+0.39 ms clock, 0.038+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 774 @3.789s 7%: 0.005+0+0.36 ms clock, 0.086+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 257568 HeapInuse: 557056 HeapObjects: 399 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 775 @3.790s 7%: 0.002+0+0.35 ms clock, 0.034+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 776 @3.791s 7%: 0.004+0+0.37 ms clock, 0.061+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 256432 HeapInuse: 557056 HeapObjects: 390 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 777 @3.793s 7%: 0.005+0+1.0 ms clock, 0.085+0/0/0+16 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 778 @3.794s 7%: 0.005+0+0.35 ms clock, 0.076+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 257064 HeapInuse: 557056 HeapObjects: 396 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 779 @3.796s 7%: 0.004+0+0.35 ms clock, 0.065+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 780 @3.796s 7%: 0.003+0+0.41 ms clock, 0.056+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 256184 HeapInuse: 557056 HeapObjects: 387 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 781 @3.798s 7%: 0.002+0+0.34 ms clock, 0.037+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 782 @3.799s 7%: 0.004+0+0.39 ms clock, 0.062+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 257392 HeapInuse: 557056 HeapObjects: 401 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 783 @3.800s 7%: 0.001+0+0.42 ms clock, 0.028+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 784 @3.801s 7%: 0.003+0+0.35 ms clock, 0.053+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 256368 HeapInuse: 557056 HeapObjects: 389 HeapIdle 1835008 HeapReleased 0 HeapSys 2392064
gc 785 @3.803s 7%: 0.002+0+0.37 ms clock, 0.035+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       184 ns/op	       6 B/op	       0 allocs/op
gc 786 @3.803s 7%: 0.003+0+0.31 ms clock, 0.045+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379520 HeapInuse: 606208 HeapObjects: 272 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 787 @3.805s 7%: 0.005+0+0.34 ms clock, 0.082+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E4-16    	gc 788 @3.805s 7%: 0.004+0+0.37 ms clock, 0.070+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 789 @3.807s 7%: 0.005+0+0.33 ms clock, 0.094+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 790 @3.807s 7%: 0.003+0+0.37 ms clock, 0.057+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 791 @3.809s 7%: 0.013+0+0.33 ms clock, 0.19+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 792 @3.809s 7%: 0.045+0+0.33 ms clock, 0.68+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 793 @3.811s 7%: 0.002+0+0.39 ms clock, 0.038+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 794 @3.811s 7%: 0.003+0+0.34 ms clock, 0.055+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 795 @3.813s 7%: 0.004+0+0.36 ms clock, 0.067+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 796 @3.813s 7%: 0.041+0+0.28 ms clock, 0.62+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 797 @3.815s 7%: 0.004+0+0.35 ms clock, 0.067+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 798 @3.815s 8%: 0.006+0+0.36 ms clock, 0.10+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 799 @3.817s 8%: 0.002+0+0.32 ms clock, 0.037+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 800 @3.817s 8%: 0.004+0+0.31 ms clock, 0.061+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 801 @3.819s 8%: 0.004+0+0.34 ms clock, 0.064+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 802 @3.819s 8%: 0.004+0+0.28 ms clock, 0.067+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 803 @3.821s 8%: 0.002+0+0.33 ms clock, 0.036+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 804 @3.821s 8%: 0.012+0+0.33 ms clock, 0.19+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 805 @3.823s 8%: 0.015+0+0.38 ms clock, 0.23+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 806 @3.823s 8%: 0.077+0+0.32 ms clock, 1.1+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 807 @3.825s 8%: 0.004+0+0.38 ms clock, 0.066+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 808 @3.825s 8%: 0.004+0+0.31 ms clock, 0.065+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 809 @3.827s 8%: 0.004+0+0.29 ms clock, 0.067+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 810 @3.827s 8%: 0.049+0+0.35 ms clock, 0.73+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 379600 HeapInuse: 606208 HeapObjects: 275 HeapIdle 1785856 HeapReleased 0 HeapSys 2392064
gc 811 @3.829s 8%: 0.017+0+0.47 ms clock, 0.26+0/0/0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       174 ns/op	      19 B/op	       0 allocs/op
gc 812 @3.830s 8%: 0.007+0+0.38 ms clock, 0.11+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 993952 HeapInuse: 1220608 HeapObjects: 272 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 813 @3.862s 8%: 0.004+0+0.36 ms clock, 0.064+0/0/0+5.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSort1E5-16                               	gc 814 @3.863s 8%: 0.004+0+0.34 ms clock, 0.074+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 815 @3.895s 7%: 0.004+0+0.40 ms clock, 0.064+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 816 @3.895s 8%: 0.080+0+0.38 ms clock, 1.2+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 817 @3.927s 7%: 0.036+0+0.38 ms clock, 0.54+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 818 @3.928s 7%: 0.044+0+0.31 ms clock, 0.66+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 819 @3.960s 7%: 0.002+0+0.40 ms clock, 0.039+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 820 @3.960s 7%: 0.046+0+0.38 ms clock, 0.69+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 821 @3.992s 7%: 0.005+0+0.47 ms clock, 0.083+0/0/0+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 822 @3.993s 7%: 0.051+0+0.42 ms clock, 0.76+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 823 @4.024s 7%: 0.002+0+0.42 ms clock, 0.036+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 824 @4.025s 7%: 0.033+0+0.39 ms clock, 0.49+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 825 @4.057s 7%: 0.004+0+0.39 ms clock, 0.072+0/0/0+5.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 826 @4.058s 7%: 0.067+0+0.35 ms clock, 1.0+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 827 @4.090s 7%: 0.005+0+0.33 ms clock, 0.078+0/0/0+5.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 828 @4.091s 7%: 0.043+0+0.35 ms clock, 0.65+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 829 @4.123s 7%: 0.062+0+0.38 ms clock, 0.93+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 830 @4.123s 7%: 0.046+0+0.37 ms clock, 0.69+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 831 @4.155s 7%: 0.004+0+0.41 ms clock, 0.068+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 832 @4.156s 7%: 0.045+0+0.31 ms clock, 0.68+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 833 @4.188s 7%: 0.034+0+0.38 ms clock, 0.51+0/0/0+5.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 834 @4.189s 7%: 0.081+0+0.35 ms clock, 1.2+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 835 @4.222s 7%: 0.040+0+0.38 ms clock, 0.60+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 836 @4.223s 7%: 0.044+0+0.35 ms clock, 0.67+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 837 @4.257s 7%: 0.036+0+0.65 ms clock, 0.54+0/0/0+9.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 838 @4.258s 7%: 0.071+0+0.37 ms clock, 1.0+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 994016 HeapInuse: 1220608 HeapObjects: 275 HeapIdle 1171456 HeapReleased 0 HeapSys 2392064
gc 839 @4.295s 7%: 0.004+0+0.38 ms clock, 0.071+0/0/0+6.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
  100000	       366 ns/op	       8 B/op	       0 allocs/op
gc 840 @4.295s 7%: 0.028+0+0.34 ms clock, 0.42+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 841 @4.464s 7%: 0.051+5.2+0.36 ms clock, 0.77+0.25/15/0.84+5.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 842 @4.641s 6%: 0.021+7.2+0.28 ms clock, 0.34+0.13/25/31+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000576 HeapInuse: 9265152 HeapObjects: 150312 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 843 @4.671s 7%: 0.002+0+0.97 ms clock, 0.038+0.13/25/31+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
#BenchmarkSetInsert1E5-16                          	gc 844 @4.672s 7%: 0.062+0+0.43 ms clock, 0.99+0.13/25/31+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 845 @4.843s 6%: 0.020+3.8+0.40 ms clock, 0.33+0.13/13/22+6.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 846 @5.016s 6%: 0.017+8.6+0.28 ms clock, 0.28+0.18/26/36+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000640 HeapInuse: 9265152 HeapObjects: 150315 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 847 @5.044s 6%: 0.002+0+1.1 ms clock, 0.034+0.18/26/36+19 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 848 @5.045s 6%: 0.044+0+0.45 ms clock, 0.71+0.18/26/36+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 849 @5.212s 6%: 0.016+4.0+0.29 ms clock, 0.26+0.15/12/18+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 850 @5.383s 6%: 0.020+6.8+0.31 ms clock, 0.32+0.12/23/42+5.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 851 @5.410s 6%: 0.002+0+0.96 ms clock, 0.045+0.12/23/42+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 852 @5.411s 6%: 0.068+0+0.44 ms clock, 1.0+0.12/23/42+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 853 @5.578s 6%: 0.020+3.5+0.46 ms clock, 0.32+0.15/11/20+7.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 854 @5.750s 6%: 0.020+5.8+0.28 ms clock, 0.33+0.10/20/47+4.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 855 @5.776s 6%: 0.004+0+0.94 ms clock, 0.070+0.10/20/47+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 856 @5.777s 6%: 0.074+0+0.44 ms clock, 1.1+0.10/20/47+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 857 @5.951s 5%: 0.021+3.2+0.30 ms clock, 0.34+0.17/10/22+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 858 @6.121s 5%: 0.020+7.9+0.25 ms clock, 0.32+0.14/29/22+4.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 859 @6.148s 5%: 0.004+0+1.0 ms clock, 0.076+0.14/29/22+17 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 860 @6.149s 5%: 0.043+0+0.44 ms clock, 0.69+0.14/29/22+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 861 @6.315s 5%: 0.023+3.4+0.30 ms clock, 0.37+0.18/12/18+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 862 @6.489s 5%: 0.018+6.7+0.30 ms clock, 0.29+0.19/25/41+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 863 @6.515s 5%: 0.002+0+0.97 ms clock, 0.046+0.19/25/41+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 864 @6.517s 5%: 0.027+0+0.51 ms clock, 0.43+0.19/25/41+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 865 @6.690s 5%: 0.021+3.2+0.29 ms clock, 0.33+0.13/11/24+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 866 @6.863s 5%: 0.019+7.7+0.25 ms clock, 0.31+0.076/23/40+4.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 867 @6.891s 5%: 0.004+0+1.1 ms clock, 0.065+0.076/23/40+18 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 868 @6.893s 5%: 0.050+0+0.39 ms clock, 0.81+0.076/23/40+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 869 @7.064s 5%: 0.022+4.3+0.31 ms clock, 0.36+0.18/13/20+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 870 @7.241s 5%: 0.019+8.6+0.41 ms clock, 0.30+0.16/32/8.7+6.6 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000640 HeapInuse: 9265152 HeapObjects: 150315 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 871 @7.268s 5%: 0.004+0+0.98 ms clock, 0.065+0.16/32/8.7+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 872 @7.269s 5%: 0.015+0+0.48 ms clock, 0.24+0.16/32/8.7+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 873 @7.437s 5%: 0.020+3.0+0.30 ms clock, 0.32+0.19/10/21+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 874 @7.608s 5%: 0.020+6.3+0.30 ms clock, 0.33+0.15/22/48+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000640 HeapInuse: 9265152 HeapObjects: 150315 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 875 @7.634s 5%: 0.004+0+0.97 ms clock, 0.068+0.15/22/48+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 876 @7.635s 5%: 0.044+0+0.49 ms clock, 0.71+0.15/22/48+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 877 @7.801s 5%: 0.022+3.2+0.42 ms clock, 0.35+0.12/11/22+6.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 878 @7.975s 4%: 0.020+6.7+0.28 ms clock, 0.33+0.21/24/35+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000640 HeapInuse: 9265152 HeapObjects: 150315 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 879 @8.002s 4%: 0.003+0+0.96 ms clock, 0.048+0.21/24/35+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 880 @8.004s 4%: 0.062+0+0.49 ms clock, 0.99+0.21/24/35+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 881 @8.171s 4%: 0.018+4.0+0.39 ms clock, 0.29+0.13/13/17+6.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 882 @8.348s 4%: 0.020+7.3+0.34 ms clock, 0.33+0.16/25/35+5.5 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000640 HeapInuse: 9265152 HeapObjects: 150315 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 883 @8.374s 4%: 0.002+0+0.98 ms clock, 0.045+0.16/25/35+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 884 @8.375s 4%: 0.041+0+0.44 ms clock, 0.66+0.16/25/35+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 885 @8.552s 4%: 0.022+4.0+0.35 ms clock, 0.36+0.15/14/12+5.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 886 @8.728s 4%: 0.020+6.1+0.28 ms clock, 0.33+0.18/22/39+4.5 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 887 @8.752s 4%: 0.002+0+0.95 ms clock, 0.042+0.18/22/39+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 888 @8.753s 4%: 0.040+0+0.39 ms clock, 0.64+0.18/22/39+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 889 @8.922s 4%: 0.021+3.3+0.33 ms clock, 0.34+0.16/11/19+5.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 890 @9.093s 4%: 0.020+5.3+0.29 ms clock, 0.33+0/19/34+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000672 HeapInuse: 9265152 HeapObjects: 150317 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 891 @9.118s 4%: 0.028+0+0.97 ms clock, 0.44+0/19/34+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 892 @9.119s 4%: 0.078+0+0.49 ms clock, 1.2+0/19/34+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 893 @9.286s 4%: 0.019+4.0+0.29 ms clock, 0.30+0.20/11/17+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 894 @9.460s 4%: 0.020+6.0+0.31 ms clock, 0.33+0.18/22/37+5.1 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000656 HeapInuse: 9265152 HeapObjects: 150316 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 895 @9.483s 4%: 0.002+0+1.2 ms clock, 0.042+0.18/22/37+19 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 896 @9.484s 4%: 0.042+0+0.45 ms clock, 0.68+0.18/22/37+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 897 @9.650s 4%: 0.017+3.0+0.35 ms clock, 0.27+0.14/10/21+5.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 898 @9.824s 4%: 0.021+7.2+0.33 ms clock, 0.33+0.13/21/38+5.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000640 HeapInuse: 9265152 HeapObjects: 150315 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 899 @9.853s 4%: 0.005+0+1.1 ms clock, 0.088+0.13/21/38+17 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 900 @9.854s 4%: 0.006+0+0.45 ms clock, 0.10+0.13/21/38+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 901 @10.025s 4%: 0.021+3.0+0.29 ms clock, 0.33+0.16/10/16+4.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 902 @10.196s 4%: 0.020+6.4+0.29 ms clock, 0.33+0.17/22/49+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000672 HeapInuse: 9265152 HeapObjects: 150317 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 903 @10.223s 4%: 0.002+0+0.93 ms clock, 0.034+0.17/22/49+14 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 904 @10.224s 4%: 0.11+0+0.44 ms clock, 1.7+0.17/22/49+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 905 @10.392s 4%: 0.030+3.3+0.32 ms clock, 0.49+0.10/11/16+5.2 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 906 @10.564s 4%: 0.019+6.5+0.31 ms clock, 0.31+0.17/22/43+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9000624 HeapInuse: 9265152 HeapObjects: 150314 HeapIdle 368640 HeapReleased 0 HeapSys 9633792
gc 907 @10.590s 4%: 0.002+0+0.96 ms clock, 0.041+0.17/22/43+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
  100000	      3663 ns/op	      88 B/op	       2 allocs/op
gc 908 @10.591s 4%: 0.014+0+0.43 ms clock, 0.22+0.17/22/43+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 909 @10.726s 4%: 0.020+3.0+0.31 ms clock, 0.32+0.11/9.7/13+4.9 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 910 @10.919s 4%: 0.021+5.6+0.29 ms clock, 0.35+0.19/20/45+4.7 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 911 @11.145s 4%: 0.020+2.8+0.27 ms clock, 0.32+0.17/9.7/10+4.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5954384 HeapInuse: 6209536 HeapObjects: 84706 HeapIdle 10764288 HeapReleased 0 HeapSys 16973824
gc 912 @11.212s 4%: 0.002+0+0.93 ms clock, 0.039+0.17/9.7/10+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
#BenchmarkSetErase1E5-16                           	gc 913 @11.213s 4%: 0.007+0+0.50 ms clock, 0.12+0.17/9.7/10+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 914 @11.345s 4%: 0.019+2.6+0.31 ms clock, 0.31+0.12/8.8/18+4.9 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 915 @11.534s 3%: 0.021+6.8+0.32 ms clock, 0.33+0.15/20/32+5.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 916 @11.756s 3%: 0.022+3.5+0.56 ms clock, 0.35+0/10/7.5+9.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5959456 HeapInuse: 6217728 HeapObjects: 84794 HeapIdle 10756096 HeapReleased 0 HeapSys 16973824
gc 917 @11.825s 3%: 0.002+0+0.73 ms clock, 0.046+0/10/7.5+11 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 918 @11.826s 3%: 0.004+0+0.74 ms clock, 0.072+0/10/7.5+11 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 919 @11.962s 3%: 0.020+3.4+0.30 ms clock, 0.33+0.19/10/11+4.8 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 920 @12.154s 3%: 0.021+7.4+0.38 ms clock, 0.33+0.18/23/24+6.2 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 921 @12.376s 3%: 0.018+4.3+0.29 ms clock, 0.29+0.16/10/9.5+4.6 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6008224 HeapInuse: 6258688 HeapObjects: 85626 HeapIdle 10682368 HeapReleased 0 HeapSys 16941056
gc 922 @12.444s 3%: 0.004+0+0.86 ms clock, 0.073+0.16/10/9.5+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 923 @12.445s 3%: 0.043+0+0.47 ms clock, 0.69+0.16/10/9.5+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 924 @12.578s 3%: 0.022+3.0+0.29 ms clock, 0.36+0.11/9.2/15+4.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 925 @12.767s 3%: 0.020+5.2+0.30 ms clock, 0.32+0.16/18/40+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 926 @12.987s 3%: 0.021+3.4+0.27 ms clock, 0.33+0.18/9.3/12+4.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6295616 HeapInuse: 6545408 HeapObjects: 90524 HeapIdle 10428416 HeapReleased 0 HeapSys 16973824
gc 927 @13.057s 3%: 0.003+0+0.80 ms clock, 0.051+0.18/9.3/12+12 ms cpu, 6->6->0 MB, 6 MB goal, 16 P (forced)
gc 928 @13.058s 3%: 0.047+0+0.47 ms clock, 0.76+0.18/9.3/12+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 929 @13.189s 3%: 0.019+2.8+0.31 ms clock, 0.30+0.16/9.1/17+5.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 930 @13.379s 3%: 0.021+7.6+0.29 ms clock, 0.34+0.18/22/36+4.7 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 931 @13.605s 3%: 0.021+3.3+0.39 ms clock, 0.33+0.18/8.3/10+6.3 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5792976 HeapInuse: 6045696 HeapObjects: 81957 HeapIdle 10895360 HeapReleased 0 HeapSys 16941056
gc 932 @13.669s 3%: 0.006+0+0.90 ms clock, 0.11+0.18/8.3/10+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 933 @13.670s 3%: 0.074+0+0.50 ms clock, 1.1+0.18/8.3/10+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 934 @13.803s 3%: 0.022+3.5+0.28 ms clock, 0.35+0.15/9.7/13+4.5 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 935 @13.998s 3%: 0.020+5.6+0.29 ms clock, 0.32+0.16/19/40+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 936 @14.223s 3%: 0.021+3.7+0.34 ms clock, 0.33+0.039/11/6.6+5.5 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5846736 HeapInuse: 6094848 HeapObjects: 82873 HeapIdle 10878976 HeapReleased 0 HeapSys 16973824
gc 937 @14.290s 3%: 0.004+0+0.76 ms clock, 0.073+0.039/11/6.6+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 938 @14.291s 3%: 0.004+0+0.45 ms clock, 0.067+0.039/11/6.6+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 939 @14.425s 3%: 0.022+2.7+0.32 ms clock, 0.35+0.17/9.8/13+5.1 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 940 @14.619s 3%: 0.022+4.6+0.35 ms clock, 0.35+0.17/16/37+5.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 941 @14.840s 3%: 0.019+3.8+0.44 ms clock, 0.31+0.18/11/12+7.0 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6098144 HeapInuse: 6348800 HeapObjects: 87158 HeapIdle 10625024 HeapReleased 0 HeapSys 16973824
gc 942 @14.910s 3%: 0.003+0+0.79 ms clock, 0.050+0.18/11/12+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 943 @14.911s 3%: 0.045+0+0.44 ms clock, 0.73+0.18/11/12+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 944 @15.043s 3%: 0.021+2.9+0.31 ms clock, 0.34+0.12/10/13+4.9 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 945 @15.236s 3%: 0.020+6.0+0.26 ms clock, 0.32+0.17/20/42+4.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 946 @15.456s 3%: 0.023+2.4+0.29 ms clock, 0.37+0/8.6/14+4.6 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5959904 HeapInuse: 6217728 HeapObjects: 84802 HeapIdle 10756096 HeapReleased 0 HeapSys 16973824
gc 947 @15.522s 3%: 0.004+0+0.88 ms clock, 0.075+0/8.6/14+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 948 @15.523s 3%: 0.045+0+0.46 ms clock, 0.72+0/8.6/14+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 949 @15.655s 3%: 0.020+3.2+0.28 ms clock, 0.32+0.19/10/8.4+4.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 950 @15.842s 3%: 0.019+6.6+0.29 ms clock, 0.30+0.20/23/42+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 951 @16.063s 3%: 0.022+2.6+0.28 ms clock, 0.35+0.15/8.8/11+4.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6331520 HeapInuse: 6586368 HeapObjects: 91136 HeapIdle 10354688 HeapReleased 0 HeapSys 16941056
gc 952 @16.136s 3%: 0.003+0+0.83 ms clock, 0.048+0.15/8.8/11+13 ms cpu, 6->6->0 MB, 6 MB goal, 16 P (forced)
gc 953 @16.137s 3%: 0.077+0+0.44 ms clock, 1.2+0.15/8.8/11+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 954 @16.270s 3%: 0.018+3.2+0.31 ms clock, 0.30+0.18/10/8.9+5.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 955 @16.461s 3%: 0.020+7.6+0.35 ms clock, 0.32+0/21/35+5.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 956 @16.683s 3%: 0.019+4.0+0.33 ms clock, 0.31+0.19/9.0/6.8+5.3 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5882640 HeapInuse: 6135808 HeapObjects: 83485 HeapIdle 10838016 HeapReleased 0 HeapSys 16973824
gc 957 @16.749s 3%: 0.003+0+1.0 ms clock, 0.063+0.19/9.0/6.8+17 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 958 @16.750s 3%: 0.060+0+0.48 ms clock, 0.96+0.19/9.0/6.8+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 959 @16.884s 3%: 0.020+2.7+0.29 ms clock, 0.32+0.15/9.2/12+4.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 960 @17.074s 3%: 0.019+6.6+0.30 ms clock, 0.31+0.20/21/28+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 961 @17.294s 2%: 0.020+2.2+0.35 ms clock, 0.32+0.14/7.8/10+5.6 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6026080 HeapInuse: 6275072 HeapObjects: 85930 HeapIdle 10665984 HeapReleased 0 HeapSys 16941056
gc 962 @17.363s 2%: 0.004+0+0.80 ms clock, 0.079+0.14/7.8/10+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 963 @17.364s 2%: 0.038+0+0.55 ms clock, 0.60+0.14/7.8/10+8.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 964 @17.495s 2%: 0.018+3.5+0.34 ms clock, 0.30+0.18/10/12+5.5 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 965 @17.685s 2%: 0.020+7.5+0.30 ms clock, 0.32+0.13/22/31+4.8 ms cpu, 7->8->8 MB, 8 MB goal, 16 P
gc 966 @17.910s 2%: 0.020+2.2+0.29 ms clock, 0.33+0.18/7.3/12+4.7 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5667136 HeapInuse: 5914624 HeapObjects: 79812 HeapIdle 11026432 HeapReleased 0 HeapSys 16941056
gc 967 @17.972s 2%: 0.004+0+0.92 ms clock, 0.078+0.18/7.3/12+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 968 @17.973s 2%: 0.075+0+0.59 ms clock, 1.2+0.18/7.3/12+9.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 969 @18.106s 2%: 0.019+2.8+0.39 ms clock, 0.31+0.20/9.4/11+6.2 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 970 @18.295s 2%: 0.020+6.0+0.28 ms clock, 0.32+0.18/21/31+4.5 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 971 @18.514s 2%: 0.025+3.1+0.30 ms clock, 0.40+0.18/9.4/8.2+4.9 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6008208 HeapInuse: 6258688 HeapObjects: 85625 HeapIdle 10682368 HeapReleased 0 HeapSys 16941056
gc 972 @18.583s 2%: 0.004+0+0.94 ms clock, 0.077+0.18/9.4/8.2+15 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 973 @18.584s 2%: 0.074+0+0.51 ms clock, 1.1+0.18/9.4/8.2+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 974 @18.717s 2%: 0.018+3.1+0.29 ms clock, 0.29+0.16/10/10+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 975 @18.914s 2%: 0.020+4.5+0.29 ms clock, 0.33+0.16/16/37+4.7 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 976 @19.133s 2%: 0.020+2.5+0.31 ms clock, 0.33+0.15/7.8/12+5.0 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5972512 HeapInuse: 6225920 HeapObjects: 85018 HeapIdle 10715136 HeapReleased 0 HeapSys 16941056
gc 977 @19.200s 2%: 0.014+0+0.75 ms clock, 0.23+0.15/7.8/12+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 978 @19.201s 2%: 0.072+0+0.51 ms clock, 1.1+0.15/7.8/12+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 979 @19.334s 2%: 0.020+3.1+0.27 ms clock, 0.32+0.18/9.3/16+4.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 980 @19.526s 2%: 0.021+6.9+0.27 ms clock, 0.33+0.23/24/33+4.3 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 981 @19.755s 2%: 0.020+4.0+0.32 ms clock, 0.32+0.17/8.6/8.8+5.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5864800 HeapInuse: 6119424 HeapObjects: 83182 HeapIdle 10854400 HeapReleased 0 HeapSys 16973824
gc 982 @19.821s 2%: 0.016+0+0.85 ms clock, 0.25+0.17/8.6/8.8+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 983 @19.822s 2%: 0.004+0+0.48 ms clock, 0.078+0.17/8.6/8.8+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 984 @19.955s 2%: 0.023+3.4+0.27 ms clock, 0.37+0.16/11/13+4.4 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 985 @20.153s 2%: 0.021+5.1+0.29 ms clock, 0.34+0.12/18/28+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 986 @20.375s 2%: 0.021+2.4+0.30 ms clock, 0.34+0.15/7.5/11+4.8 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5954432 HeapInuse: 6209536 HeapObjects: 84708 HeapIdle 10731520 HeapReleased 0 HeapSys 16941056
gc 987 @20.441s 2%: 0.002+0+0.80 ms clock, 0.034+0.15/7.5/11+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 988 @20.442s 2%: 0.041+0+0.49 ms clock, 0.65+0.15/7.5/11+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 989 @20.574s 2%: 0.018+2.6+0.30 ms clock, 0.29+0.17/8.0/13+4.8 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 990 @20.762s 2%: 0.020+5.5+0.36 ms clock, 0.32+0.14/14/30+5.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 991 @20.984s 2%: 0.021+3.1+0.32 ms clock, 0.34+0.15/9.4/8.9+5.2 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5972512 HeapInuse: 6225920 HeapObjects: 85018 HeapIdle 10715136 HeapReleased 0 HeapSys 16941056
gc 992 @21.050s 2%: 0.005+0+0.81 ms clock, 0.082+0.15/9.4/8.9+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
  100000	      2448 ns/op	      96 B/op	       2 allocs/op
gc 993 @21.051s 2%: 0.007+0+0.49 ms clock, 0.12+0.15/9.4/8.9+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 994 @21.148s 2%: 0.017+0.78+0.36 ms clock, 0.28+0.16/1.6/0.75+5.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 995 @21.232s 2%: 0.015+1.0+0.41 ms clock, 0.25+0.21/2.8/2.4+6.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 996 @21.309s 2%: 0.018+0.74+0.31 ms clock, 0.30+0.15/1.8/0.90+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706400 HeapInuse: 1949696 HeapObjects: 12295 HeapIdle 15024128 HeapReleased 0 HeapSys 16973824
gc 997 @21.319s 2%: 0.004+0+0.58 ms clock, 0.065+0.15/1.8/0.90+9.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E5-16                  	gc 998 @21.320s 2%: 0.006+0+0.47 ms clock, 0.097+0.15/1.8/0.90+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 999 @21.423s 2%: 0.020+0.72+0.38 ms clock, 0.32+0.16/1.8/0.82+6.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1000 @21.508s 2%: 0.018+1.0+0.29 ms clock, 0.28+0.15/2.7/1.8+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1001 @21.584s 2%: 0.017+1.1+0.36 ms clock, 0.28+0.16/2.0/0.74+5.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706496 HeapInuse: 1966080 HeapObjects: 12300 HeapIdle 15007744 HeapReleased 0 HeapSys 16973824
gc 1002 @21.593s 2%: 0.004+0+0.67 ms clock, 0.077+0.16/2.0/0.74+10 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1003 @21.594s 2%: 0.050+0+0.53 ms clock, 0.80+0.16/2.0/0.74+8.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1004 @21.690s 2%: 0.018+0.73+0.34 ms clock, 0.29+0.17/1.9/0.23+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1005 @21.775s 2%: 0.019+0.90+0.37 ms clock, 0.31+0.14/2.4/3.1+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1006 @21.851s 2%: 0.023+0.98+0.34 ms clock, 0.36+0.062/1.6/0.81+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706464 HeapInuse: 1966080 HeapObjects: 12298 HeapIdle 15007744 HeapReleased 0 HeapSys 16973824
gc 1007 @21.861s 2%: 0.002+0+0.50 ms clock, 0.045+0.062/1.6/0.81+8.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1008 @21.861s 2%: 0.004+0+0.46 ms clock, 0.076+0.062/1.6/0.81+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1009 @21.959s 2%: 0.017+0.72+0.35 ms clock, 0.27+0.12/1.7/0.88+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1010 @22.043s 2%: 0.017+1.2+0.35 ms clock, 0.28+0.18/2.4/1.9+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1011 @22.117s 2%: 0.017+0.74+0.38 ms clock, 0.28+0.14/1.5/0.99+6.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706464 HeapInuse: 1982464 HeapObjects: 12298 HeapIdle 14991360 HeapReleased 0 HeapSys 16973824
gc 1012 @22.125s 2%: 0.003+0+0.56 ms clock, 0.061+0.14/1.5/0.99+9.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1013 @22.126s 2%: 0.083+0+0.47 ms clock, 1.3+0.14/1.5/0.99+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1014 @22.224s 2%: 0.018+0.79+0.34 ms clock, 0.28+0.18/1.5/0.68+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1015 @22.310s 2%: 0.015+1.6+0.38 ms clock, 0.24+0.16/3.2/1.2+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1016 @22.385s 2%: 0.026+0.92+0.34 ms clock, 0.42+0.18/2.0/0.23+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706544 HeapInuse: 1957888 HeapObjects: 12299 HeapIdle 15015936 HeapReleased 0 HeapSys 16973824
gc 1017 @22.395s 2%: 0.002+0+0.55 ms clock, 0.038+0.18/2.0/0.23+8.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1018 @22.395s 2%: 0.044+0+0.44 ms clock, 0.70+0.18/2.0/0.23+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1019 @22.494s 2%: 0.031+0.78+0.32 ms clock, 0.50+0.063/1.7/0.96+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1020 @22.580s 2%: 0.019+0.89+0.35 ms clock, 0.30+0.11/2.5/2.6+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1021 @22.655s 2%: 0.017+0.73+0.30 ms clock, 0.27+0.12/1.5/1.5+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706496 HeapInuse: 1966080 HeapObjects: 12300 HeapIdle 15007744 HeapReleased 0 HeapSys 16973824
gc 1022 @22.664s 2%: 0.027+0+0.53 ms clock, 0.43+0.12/1.5/1.5+8.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1023 @22.664s 2%: 0.038+0+0.43 ms clock, 0.61+0.12/1.5/1.5+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1024 @22.760s 2%: 0.018+0.84+0.32 ms clock, 0.28+0.19/1.7/0.71+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1025 @22.846s 2%: 0.025+1.8+0.69 ms clock, 0.40+0.31/3.8/1.4+11 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1026 @22.924s 2%: 0.019+0.83+0.32 ms clock, 0.31+0.14/1.7/0.47+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706480 HeapInuse: 1966080 HeapObjects: 12299 HeapIdle 15007744 HeapReleased 0 HeapSys 16973824
gc 1027 @22.933s 2%: 0.017+0+0.56 ms clock, 0.28+0.14/1.7/0.47+9.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1028 @22.934s 2%: 0.012+0+0.48 ms clock, 0.19+0.14/1.7/0.47+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1029 @23.032s 2%: 0.020+0.94+0.29 ms clock, 0.32+0.17/2.0/0.26+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1030 @23.119s 2%: 0.019+1.2+0.37 ms clock, 0.30+0.17/3.2/1.2+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1031 @23.194s 2%: 0.016+0.72+0.34 ms clock, 0.27+0.17/1.7/0.96+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706560 HeapInuse: 1949696 HeapObjects: 12300 HeapIdle 15024128 HeapReleased 0 HeapSys 16973824
gc 1032 @23.202s 2%: 0.031+0+0.51 ms clock, 0.50+0.17/1.7/0.96+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1033 @23.203s 2%: 0.077+0+0.41 ms clock, 1.2+0.17/1.7/0.96+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1034 @23.300s 2%: 0.018+0.72+0.30 ms clock, 0.28+0.19/1.4/0.94+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1035 @23.387s 2%: 0.017+1.4+0.33 ms clock, 0.28+0.15/2.2/2.0+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1036 @23.461s 2%: 0.015+0.67+0.36 ms clock, 0.24+0.10/1.5/1.1+5.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706496 HeapInuse: 1990656 HeapObjects: 12300 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 1037 @23.470s 2%: 0.002+0+0.63 ms clock, 0.038+0.10/1.5/1.1+10 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1038 @23.471s 2%: 0.045+0+0.49 ms clock, 0.72+0.10/1.5/1.1+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1039 @23.570s 2%: 0.020+0.83+0.29 ms clock, 0.33+0.13/1.7/0.54+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1040 @23.654s 2%: 0.018+1.1+0.37 ms clock, 0.29+0.14/3.2/2.9+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1041 @23.730s 2%: 0.015+0.79+0.32 ms clock, 0.24+0.16/1.9/1.1+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706560 HeapInuse: 1974272 HeapObjects: 12300 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1042 @23.738s 2%: 0.033+0+0.57 ms clock, 0.52+0.16/1.9/1.1+9.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1043 @23.739s 2%: 0.037+0+0.51 ms clock, 0.60+0.16/1.9/1.1+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1044 @23.836s 2%: 0.022+0.76+0.31 ms clock, 0.35+0.17/1.6/1.3+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1045 @23.922s 2%: 0.021+1.1+0.40 ms clock, 0.35+0.18/2.9/2.2+6.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1046 @23.996s 2%: 0.015+0.63+0.33 ms clock, 0.24+0.11/1.7/1.2+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706560 HeapInuse: 1957888 HeapObjects: 12300 HeapIdle 15015936 HeapReleased 0 HeapSys 16973824
gc 1047 @24.004s 2%: 0.002+0+0.52 ms clock, 0.038+0.11/1.7/1.2+8.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1048 @24.005s 2%: 0.052+0+0.51 ms clock, 0.83+0.11/1.7/1.2+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1049 @24.102s 2%: 0.019+0.69+0.29 ms clock, 0.30+0.17/1.5/1.0+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1050 @24.189s 2%: 0.017+1.4+0.35 ms clock, 0.27+0.15/3.7/1.1+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1051 @24.263s 2%: 0.016+0.82+0.33 ms clock, 0.26+0.095/2.0/0.66+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706576 HeapInuse: 1966080 HeapObjects: 12301 HeapIdle 14974976 HeapReleased 0 HeapSys 16941056
gc 1052 @24.272s 2%: 0.016+0+0.51 ms clock, 0.25+0.095/2.0/0.66+8.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1053 @24.273s 2%: 0.071+0+0.51 ms clock, 1.1+0.095/2.0/0.66+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1054 @24.371s 2%: 0.021+1.0+0.40 ms clock, 0.34+0.18/1.5/1.2+6.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1055 @24.458s 2%: 0.018+1.1+0.27 ms clock, 0.29+0.16/2.8/2.4+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1056 @24.534s 2%: 0.020+0.82+0.45 ms clock, 0.32+0.13/1.9/0.75+7.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706576 HeapInuse: 1966080 HeapObjects: 12301 HeapIdle 15007744 HeapReleased 0 HeapSys 16973824
gc 1057 @24.543s 2%: 0.008+0+0.58 ms clock, 0.12+0.13/1.9/0.75+9.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1058 @24.544s 2%: 0.041+0+0.52 ms clock, 0.66+0.13/1.9/0.75+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1059 @24.641s 2%: 0.018+0.71+0.33 ms clock, 0.30+0.17/2.0/0.74+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1060 @24.725s 2%: 0.017+2.2+0.33 ms clock, 0.28+0.65/2.6/0.97+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1061 @24.801s 2%: 0.018+0.89+0.28 ms clock, 0.30+0.026/2.1/1.0+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706480 HeapInuse: 1990656 HeapObjects: 12299 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 1062 @24.810s 2%: 0.027+0+0.53 ms clock, 0.44+0.026/2.1/1.0+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1063 @24.811s 2%: 0.038+0+0.44 ms clock, 0.61+0.026/2.1/1.0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1064 @24.910s 2%: 0.020+0.74+0.35 ms clock, 0.32+0.14/1.7/1.1+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1065 @24.995s 2%: 0.018+0.99+0.65 ms clock, 0.29+0.14/2.6/2.2+10 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1066 @25.071s 2%: 0.020+0.76+0.39 ms clock, 0.33+0.18/1.8/0.78+6.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706496 HeapInuse: 1949696 HeapObjects: 12300 HeapIdle 15024128 HeapReleased 0 HeapSys 16973824
gc 1067 @25.079s 2%: 0.004+0+0.62 ms clock, 0.078+0.18/1.8/0.78+9.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1068 @25.080s 2%: 0.008+0+0.48 ms clock, 0.13+0.18/1.8/0.78+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1069 @25.176s 2%: 0.017+0.65+0.30 ms clock, 0.27+0.092/1.5/0.60+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1070 @25.260s 2%: 0.017+0.96+0.29 ms clock, 0.27+0.12/2.7/2.8+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1071 @25.335s 2%: 0.018+0.72+0.34 ms clock, 0.29+0.14/1.7/1.3+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706464 HeapInuse: 1957888 HeapObjects: 12298 HeapIdle 15015936 HeapReleased 0 HeapSys 16973824
gc 1072 @25.343s 2%: 0.031+0+0.62 ms clock, 0.50+0.14/1.7/1.3+10 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1073 @25.344s 2%: 0.044+0+0.62 ms clock, 0.70+0.14/1.7/1.3+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1074 @25.443s 2%: 0.019+0.83+0.33 ms clock, 0.31+0.15/1.9/0.47+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1075 @25.530s 2%: 0.021+1.5+0.41 ms clock, 0.34+0.14/3.9/2.5+6.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1076 @25.606s 2%: 0.022+0.77+0.38 ms clock, 0.35+0.13/1.9/0.45+6.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1706560 HeapInuse: 1949696 HeapObjects: 12300 HeapIdle 15024128 HeapReleased 0 HeapSys 16973824
gc 1077 @25.615s 2%: 0.007+0+0.71 ms clock, 0.11+0.13/1.9/0.45+11 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
  100000	      2713 ns/op	      96 B/op	       2 allocs/op
gc 1078 @25.616s 2%: 0.007+0+0.48 ms clock, 0.12+0.13/1.9/0.45+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1079 @25.814s 2%: 0.021+0.65+0.32 ms clock, 0.34+0.13/1.2/0.75+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926328 HeapInuse: 3194880 HeapObjects: 40317 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1080 @25.933s 2%: 0.052+0+0.62 ms clock, 0.83+0.13/1.2/0.75+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E5-16          	gc 1081 @25.934s 2%: 0.009+0+0.58 ms clock, 0.15+0.13/1.2/0.75+9.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1082 @26.135s 2%: 0.020+0.59+0.28 ms clock, 0.32+0.13/1.1/0.63+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1083 @26.251s 2%: 0.003+0+1.3 ms clock, 0.056+0.13/1.1/0.63+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1084 @26.252s 2%: 0.038+0+0.47 ms clock, 0.61+0.13/1.1/0.63+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1085 @26.451s 2%: 0.021+0.67+0.39 ms clock, 0.34+0.19/1.0/0.82+6.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926392 HeapInuse: 3194880 HeapObjects: 40319 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1086 @26.566s 2%: 0.027+0+0.64 ms clock, 0.44+0.19/1.0/0.82+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1087 @26.567s 2%: 0.075+0+0.55 ms clock, 1.2+0.19/1.0/0.82+8.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1088 @26.764s 2%: 0.019+0.59+0.28 ms clock, 0.31+0.17/1.1/0.88+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1089 @26.881s 2%: 0.004+0+0.57 ms clock, 0.068+0.17/1.1/0.88+9.2 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1090 @26.882s 2%: 0.041+0+0.42 ms clock, 0.65+0.17/1.1/0.88+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1091 @27.080s 2%: 0.022+0.53+0.35 ms clock, 0.36+0.15/1.3/0.72+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926424 HeapInuse: 3194880 HeapObjects: 40321 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1092 @27.198s 2%: 0.048+0+0.75 ms clock, 0.77+0.15/1.3/0.72+12 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1093 @27.199s 2%: 0.083+0+0.62 ms clock, 1.3+0.15/1.3/0.72+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1094 @27.398s 2%: 0.018+0.64+0.37 ms clock, 0.29+0.13/1.4/0.46+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926424 HeapInuse: 3194880 HeapObjects: 40321 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1095 @27.514s 2%: 0.017+0+0.72 ms clock, 0.28+0.13/1.4/0.46+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1096 @27.515s 2%: 0.086+0+0.57 ms clock, 1.3+0.13/1.4/0.46+9.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1097 @27.715s 2%: 0.021+0.60+0.43 ms clock, 0.33+0.19/0.85/0.59+7.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1098 @27.830s 2%: 0.004+0+0.68 ms clock, 0.078+0.19/0.85/0.59+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1099 @27.831s 2%: 0.044+0+0.60 ms clock, 0.71+0.19/0.85/0.59+9.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1100 @28.035s 2%: 0.021+0.57+0.32 ms clock, 0.33+0.19/0.84/0.70+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1101 @28.151s 2%: 0.026+0+0.63 ms clock, 0.42+0.19/0.84/0.70+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1102 @28.152s 2%: 0.078+0+0.47 ms clock, 1.2+0.19/0.84/0.70+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1103 @28.351s 2%: 0.017+0.53+0.30 ms clock, 0.27+0.17/1.1/0.85+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1104 @28.470s 2%: 0.024+0+0.65 ms clock, 0.39+0.17/1.1/0.85+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1105 @28.471s 2%: 0.041+0+0.50 ms clock, 0.65+0.17/1.1/0.85+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1106 @28.672s 2%: 0.022+0.65+0.31 ms clock, 0.36+0.17/1.1/0.48+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1107 @28.789s 2%: 0.003+0+0.67 ms clock, 0.058+0.17/1.1/0.48+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1108 @28.790s 2%: 0.042+0+0.68 ms clock, 0.68+0.17/1.1/0.48+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1109 @28.990s 2%: 0.021+0.64+0.33 ms clock, 0.33+0.17/1.0/0.66+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926424 HeapInuse: 3194880 HeapObjects: 40321 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1110 @29.107s 2%: 0.028+0+0.89 ms clock, 0.46+0.17/1.0/0.66+14 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1111 @29.108s 2%: 0.038+0+0.44 ms clock, 0.60+0.17/1.0/0.66+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1112 @29.305s 2%: 0.021+0.61+0.39 ms clock, 0.33+0.21/1.1/0.93+6.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926392 HeapInuse: 3194880 HeapObjects: 40319 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1113 @29.424s 2%: 0.007+0+0.71 ms clock, 0.12+0.21/1.1/0.93+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1114 @29.425s 2%: 0.012+0+0.50 ms clock, 0.19+0.21/1.1/0.93+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1115 @29.626s 2%: 0.021+0.53+0.32 ms clock, 0.33+0.15/1.0/0.69+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926424 HeapInuse: 3194880 HeapObjects: 40321 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1116 @29.742s 2%: 0.005+0+0.70 ms clock, 0.095+0.15/1.0/0.69+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1117 @29.743s 2%: 0.071+0+0.70 ms clock, 1.1+0.15/1.0/0.69+11 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1118 @29.942s 2%: 0.022+0.66+0.44 ms clock, 0.35+0.20/0.84/0.68+7.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1119 @30.058s 2%: 0.028+0+0.67 ms clock, 0.44+0.20/0.84/0.68+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1120 @30.059s 2%: 0.042+0+0.47 ms clock, 0.68+0.20/0.84/0.68+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1121 @30.257s 2%: 0.019+0.63+0.28 ms clock, 0.31+0.16/1.2/0.55+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1122 @30.372s 2%: 0.003+0+0.60 ms clock, 0.060+0.16/1.2/0.55+9.6 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1123 @30.373s 2%: 0.085+0+0.45 ms clock, 1.3+0.16/1.2/0.55+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1124 @30.570s 2%: 0.021+0.61+0.37 ms clock, 0.34+0.17/1.1/0.71+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926408 HeapInuse: 3194880 HeapObjects: 40320 HeapIdle 13778944 HeapReleased 0 HeapSys 16973824
gc 1125 @30.686s 2%: 0.027+0+0.67 ms clock, 0.43+0.17/1.1/0.71+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1126 @30.687s 2%: 0.037+0+0.48 ms clock, 0.60+0.17/1.1/0.71+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1127 @30.888s 2%: 0.019+0.56+0.29 ms clock, 0.31+0.19/1.3/0.36+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2926472 HeapInuse: 3203072 HeapObjects: 40321 HeapIdle 13770752 HeapReleased 0 HeapSys 16973824
gc 1128 @31.005s 2%: 0.022+0+1.0 ms clock, 0.36+0.19/1.3/0.36+17 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
  100000	      3181 ns/op	      56 B/op	       1 allocs/op
gc 1129 @31.006s 2%: 0.027+0+0.50 ms clock, 0.43+0.19/1.3/0.36+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1130 @31.104s 2%: 0.021+3.4+0.31 ms clock, 0.34+0.18/10/16+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1131 @31.200s 2%: 0.021+6.2+0.28 ms clock, 0.33+0.19/21/19+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8307856 HeapInuse: 8577024 HeapObjects: 107014 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1132 @31.221s 2%: 0.004+0+1.0 ms clock, 0.072+0.19/21/19+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
#BenchmarkIntSetInsert1E5-16                       	gc 1133 @31.222s 2%: 0.020+0+0.45 ms clock, 0.32+0.19/21/19+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1134 @31.321s 2%: 0.019+3.0+0.28 ms clock, 0.30+0.16/10/14+4.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1135 @31.414s 2%: 0.018+4.8+0.28 ms clock, 0.29+0.14/17/35+4.6 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8311968 HeapInuse: 8577024 HeapObjects: 107270 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1136 @31.435s 2%: 0.001+0+0.94 ms clock, 0.031+0.14/17/35+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1137 @31.436s 2%: 0.074+0+0.47 ms clock, 1.1+0.14/17/35+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1138 @31.533s 2%: 0.015+3.0+0.29 ms clock, 0.25+0.13/9.7/16+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1139 @31.628s 2%: 0.018+5.8+0.31 ms clock, 0.29+0.18/19/26+5.0 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8307904 HeapInuse: 8577024 HeapObjects: 107016 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1140 @31.649s 2%: 0.004+0+1.0 ms clock, 0.065+0.18/19/26+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1141 @31.650s 2%: 0.042+0+0.48 ms clock, 0.67+0.18/19/26+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1142 @31.748s 2%: 0.022+3.0+0.27 ms clock, 0.36+0.18/10/16+4.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1143 @31.846s 2%: 0.022+5.9+0.34 ms clock, 0.35+0.14/21/42+5.5 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8305456 HeapInuse: 8568832 HeapObjects: 106863 HeapIdle 8404992 HeapReleased 0 HeapSys 16973824
gc 1144 @31.868s 2%: 0.003+0+0.96 ms clock, 0.060+0.14/21/42+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1145 @31.869s 2%: 0.007+0+0.49 ms clock, 0.11+0.14/21/42+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1146 @31.967s 2%: 0.018+3.2+0.30 ms clock, 0.30+0.17/10/15+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1147 @32.060s 2%: 0.015+5.1+0.29 ms clock, 0.24+0.16/17/32+4.7 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8310352 HeapInuse: 8577024 HeapObjects: 107169 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1148 @32.082s 2%: 0.003+0+1.1 ms clock, 0.048+0.16/17/32+17 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1149 @32.083s 2%: 0.050+0+0.42 ms clock, 0.81+0.16/17/32+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1150 @32.181s 2%: 0.019+3.0+0.30 ms clock, 0.30+0.13/10/10+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1151 @32.274s 2%: 0.015+6.5+0.41 ms clock, 0.24+0/17/29+6.7 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8307280 HeapInuse: 8577024 HeapObjects: 106977 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1152 @32.296s 2%: 0.002+0+1.2 ms clock, 0.041+0/17/29+19 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1153 @32.298s 2%: 0.083+0+0.46 ms clock, 1.3+0/17/29+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1154 @32.396s 2%: 0.037+2.5+0.33 ms clock, 0.59+0.14/8.5/16+5.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1155 @32.490s 2%: 0.024+4.5+0.30 ms clock, 0.39+0.12/16/31+4.8 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8309552 HeapInuse: 8577024 HeapObjects: 107119 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1156 @32.510s 2%: 0.003+0+1.0 ms clock, 0.057+0.12/16/31+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1157 @32.511s 2%: 0.074+0+0.54 ms clock, 1.1+0.12/16/31+8.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1158 @32.608s 2%: 0.019+2.7+0.27 ms clock, 0.31+0.15/9.5/14+4.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1159 @32.703s 2%: 0.018+5.2+0.29 ms clock, 0.29+0.20/18/24+4.6 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8310352 HeapInuse: 8577024 HeapObjects: 107169 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1160 @32.723s 2%: 0.002+0+0.94 ms clock, 0.033+0.20/18/24+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1161 @32.724s 2%: 0.072+0+0.59 ms clock, 1.1+0.20/18/24+9.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1162 @32.823s 2%: 0.021+2.8+0.31 ms clock, 0.34+0.19/9.5/15+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1163 @32.917s 2%: 0.019+6.1+0.45 ms clock, 0.31+0.14/20/23+7.2 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8314432 HeapInuse: 8585216 HeapObjects: 107424 HeapIdle 8388608 HeapReleased 0 HeapSys 16973824
gc 1164 @32.940s 2%: 0.002+0+1.0 ms clock, 0.040+0.14/20/23+17 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1165 @32.941s 2%: 0.074+0+0.48 ms clock, 1.1+0.14/20/23+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1166 @33.038s 2%: 0.017+3.0+0.30 ms clock, 0.28+0.16/10/10+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1167 @33.131s 2%: 0.020+5.1+0.28 ms clock, 0.32+0.20/18/32+4.5 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8310368 HeapInuse: 8577024 HeapObjects: 107170 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1168 @33.151s 2%: 0.002+0+0.94 ms clock, 0.032+0.20/18/32+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1169 @33.153s 2%: 0.007+0+0.81 ms clock, 0.12+0.20/18/32+13 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1170 @33.251s 2%: 0.017+2.8+0.31 ms clock, 0.28+0.17/9.5/15+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1171 @33.343s 2%: 0.017+5.2+0.30 ms clock, 0.27+0.17/18/26+4.9 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8311984 HeapInuse: 8577024 HeapObjects: 107271 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1172 @33.364s 2%: 0.002+0+1.1 ms clock, 0.039+0.17/18/26+17 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1173 @33.366s 2%: 0.066+0+0.43 ms clock, 1.0+0.17/18/26+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1174 @33.466s 2%: 0.022+2.7+0.27 ms clock, 0.36+0.18/9.2/15+4.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1175 @33.559s 2%: 0.020+5.1+0.36 ms clock, 0.32+0.13/18/28+5.8 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8312768 HeapInuse: 8577024 HeapObjects: 107320 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1176 @33.581s 2%: 0.003+0+1.0 ms clock, 0.060+0.13/18/28+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1177 @33.582s 2%: 0.037+0+0.52 ms clock, 0.59+0.13/18/28+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1178 @33.682s 2%: 0.017+2.6+0.35 ms clock, 0.28+0.19/9.3/16+5.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1179 @33.776s 2%: 0.019+5.6+0.30 ms clock, 0.30+0.19/20/23+4.9 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8307920 HeapInuse: 8577024 HeapObjects: 107017 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1180 @33.797s 2%: 0.002+0+1.0 ms clock, 0.038+0.19/20/23+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1181 @33.799s 2%: 0.10+0+0.46 ms clock, 1.6+0.19/20/23+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1182 @33.896s 2%: 0.018+3.6+0.38 ms clock, 0.29+0.10/10/17+6.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1183 @33.991s 2%: 0.021+5.2+0.28 ms clock, 0.34+0.13/18/28+4.4 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8303824 HeapInuse: 8568832 HeapObjects: 106761 HeapIdle 8404992 HeapReleased 0 HeapSys 16973824
gc 1184 @34.012s 2%: 0.005+0+0.99 ms clock, 0.080+0.13/18/28+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1185 @34.013s 2%: 0.047+0+0.42 ms clock, 0.76+0.13/18/28+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1186 @34.110s 2%: 0.020+3.5+0.31 ms clock, 0.33+0.12/11/12+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1187 @34.205s 2%: 0.020+6.3+0.28 ms clock, 0.33+0.16/19/29+4.6 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8310336 HeapInuse: 8577024 HeapObjects: 107168 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1188 @34.226s 2%: 0.002+0+1.0 ms clock, 0.037+0.16/19/29+17 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1189 @34.228s 2%: 0.056+0+0.50 ms clock, 0.90+0.16/19/29+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1190 @34.324s 2%: 0.016+3.6+0.33 ms clock, 0.26+0.17/10/13+5.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1191 @34.418s 2%: 0.020+6.0+0.36 ms clock, 0.32+0.17/20/21+5.9 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
HeapAlloc: 8307936 HeapInuse: 8577024 HeapObjects: 107018 HeapIdle 8396800 HeapReleased 0 HeapSys 16973824
gc 1192 @34.439s 2%: 0.003+0+1.0 ms clock, 0.056+0.17/20/21+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
  100000	      2120 ns/op	      88 B/op	       2 allocs/op
gc 1193 @34.440s 2%: 0.005+0+0.47 ms clock, 0.083+0.17/20/21+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1194 @34.518s 2%: 0.016+2.5+0.45 ms clock, 0.26+0.17/8.2/11+7.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1195 @34.618s 2%: 0.017+5.2+0.33 ms clock, 0.27+0.18/17/14+5.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1196 @34.740s 2%: 0.017+3.4+0.29 ms clock, 0.28+0.15/11/12+4.7 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 8042800 HeapInuse: 8306688 HeapObjects: 105065 HeapIdle 8667136 HeapReleased 0 HeapSys 16973824
gc 1197 @34.809s 2%: 0.002+0+0.98 ms clock, 0.041+0.15/11/12+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
#BenchmarkIntSetErase1E5-16                        	gc 1198 @34.810s 2%: 0.008+0+0.53 ms clock, 0.13+0.15/11/12+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1199 @34.888s 2%: 0.019+2.5+0.28 ms clock, 0.31+0.12/8.7/13+4.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1200 @34.990s 2%: 0.019+4.6+0.30 ms clock, 0.30+0.16/17/24+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1201 @35.115s 2%: 0.019+5.2+0.44 ms clock, 0.31+0.19/18/0.83+7.1 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7957088 HeapInuse: 8216576 HeapObjects: 103791 HeapIdle 8757248 HeapReleased 0 HeapSys 16973824
gc 1202 @35.181s 2%: 0.016+0+0.90 ms clock, 0.26+0.19/18/0.83+14 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1203 @35.182s 2%: 0.043+0+0.44 ms clock, 0.70+0.19/18/0.83+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1204 @35.259s 2%: 0.013+2.8+0.27 ms clock, 0.22+0.12/8.4/11+4.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1205 @35.362s 2%: 0.019+6.1+0.26 ms clock, 0.30+0.13/19/26+4.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1206 @35.490s 2%: 0.020+2.9+0.32 ms clock, 0.32+0.12/9.4/15+5.2 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7837232 HeapInuse: 8093696 HeapObjects: 102008 HeapIdle 8880128 HeapReleased 0 HeapSys 16973824
gc 1207 @35.556s 2%: 0.005+0+1.1 ms clock, 0.080+0.12/9.4/15+18 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1208 @35.558s 2%: 0.073+0+0.43 ms clock, 1.1+0.12/9.4/15+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1209 @35.634s 2%: 0.015+2.3+0.30 ms clock, 0.24+0.19/7.7/10+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1210 @35.737s 2%: 0.018+4.2+0.28 ms clock, 0.30+0.014/15/29+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1211 @35.861s 2%: 0.018+2.7+0.29 ms clock, 0.29+0.096/9.7/13+4.7 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7802960 HeapInuse: 8060928 HeapObjects: 101498 HeapIdle 8912896 HeapReleased 0 HeapSys 16973824
gc 1212 @35.926s 2%: 0.023+0+0.95 ms clock, 0.38+0.096/9.7/13+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1213 @35.927s 2%: 0.083+0+0.51 ms clock, 1.3+0.096/9.7/13+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1214 @36.004s 2%: 0.016+3.0+0.31 ms clock, 0.26+0.16/8.0/10+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1215 @36.105s 2%: 0.017+5.4+0.29 ms clock, 0.27+0.18/18/26+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1216 @36.230s 2%: 0.021+4.6+0.33 ms clock, 0.33+0.19/12/7.8+5.3 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7905792 HeapInuse: 8167424 HeapObjects: 103029 HeapIdle 8806400 HeapReleased 0 HeapSys 16973824
gc 1217 @36.297s 2%: 0.002+0+0.94 ms clock, 0.045+0.19/12/7.8+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1218 @36.299s 2%: 0.046+0+0.50 ms clock, 0.75+0.19/12/7.8+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1219 @36.375s 2%: 0.017+3.2+0.28 ms clock, 0.28+0.19/8.3/10+4.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1220 @36.478s 2%: 0.019+4.1+0.27 ms clock, 0.30+0.20/14/27+4.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1221 @36.603s 2%: 0.021+3.1+0.29 ms clock, 0.33+0.18/9.1/13+4.7 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7820032 HeapInuse: 8077312 HeapObjects: 101753 HeapIdle 8896512 HeapReleased 0 HeapSys 16973824
gc 1222 @36.667s 2%: 0.014+0+1.1 ms clock, 0.23+0.18/9.1/13+19 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1223 @36.669s 2%: 0.056+0+0.84 ms clock, 0.90+0.18/9.1/13+13 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1224 @36.747s 2%: 0.016+2.6+0.30 ms clock, 0.26+0.18/8.4/10+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1225 @36.848s 2%: 0.022+7.0+0.28 ms clock, 0.35+0.97/19/41+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1226 @36.976s 2%: 0.017+3.4+0.92 ms clock, 0.28+0.20/11/14+14 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7940064 HeapInuse: 8200192 HeapObjects: 103539 HeapIdle 8773632 HeapReleased 0 HeapSys 16973824
gc 1227 @37.044s 2%: 0.003+0+0.96 ms clock, 0.053+0.20/11/14+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1228 @37.045s 2%: 0.070+0+0.63 ms clock, 1.1+0.20/11/14+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1229 @37.125s 2%: 0.022+2.9+0.30 ms clock, 0.36+0.16/8.7/11+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1230 @37.228s 2%: 0.016+4.5+0.30 ms clock, 0.27+0/15/25+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1231 @37.353s 2%: 0.018+3.7+0.31 ms clock, 0.30+0.15/11/13+5.0 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7905760 HeapInuse: 8167424 HeapObjects: 103027 HeapIdle 8806400 HeapReleased 0 HeapSys 16973824
gc 1232 @37.420s 2%: 0.008+0+1.0 ms clock, 0.14+0.15/11/13+17 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1233 @37.421s 2%: 0.040+0+0.65 ms clock, 0.64+0.15/11/13+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1234 @37.500s 2%: 0.014+2.6+0.29 ms clock, 0.23+0.14/8.5/12+4.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1235 @37.602s 2%: 0.022+4.8+0.30 ms clock, 0.35+0.16/17/27+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1236 @37.727s 2%: 0.021+3.3+0.30 ms clock, 0.33+0.17/10/13+4.8 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7837248 HeapInuse: 8093696 HeapObjects: 102009 HeapIdle 8880128 HeapReleased 0 HeapSys 16973824
gc 1237 @37.792s 2%: 0.002+0+0.91 ms clock, 0.044+0.17/10/13+14 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1238 @37.793s 2%: 0.005+0+0.51 ms clock, 0.087+0.17/10/13+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1239 @37.870s 2%: 0.020+2.5+0.27 ms clock, 0.32+0.084/8.6/11+4.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1240 @37.970s 2%: 0.018+4.8+0.32 ms clock, 0.29+0.18/16/31+5.1 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
gc 1241 @38.091s 2%: 0.019+3.1+0.30 ms clock, 0.31+0/10/13+4.9 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 8090656 HeapInuse: 8355840 HeapObjects: 105779 HeapIdle 8617984 HeapReleased 0 HeapSys 16973824
gc 1242 @38.159s 2%: 0.003+0+0.86 ms clock, 0.062+0/10/13+13 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1243 @38.160s 2%: 0.045+0+0.52 ms clock, 0.72+0/10/13+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1244 @38.238s 2%: 0.018+2.6+0.39 ms clock, 0.29+0.22/8.8/4.0+6.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1245 @38.340s 2%: 0.019+5.7+0.37 ms clock, 0.30+0.18/19/19+5.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1246 @38.466s 2%: 0.021+3.8+0.29 ms clock, 0.34+0.21/10/13+4.6 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7888544 HeapInuse: 8151040 HeapObjects: 102771 HeapIdle 8822784 HeapReleased 0 HeapSys 16973824
gc 1247 @38.532s 2%: 0.016+0+1.0 ms clock, 0.26+0.21/10/13+17 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1248 @38.534s 2%: 0.078+0+0.42 ms clock, 1.2+0.21/10/13+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1249 @38.636s 2%: 0.022+2.6+1.2 ms clock, 0.36+0.15/9.1/11+20 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1250 @38.750s 2%: 0.025+6.0+0.27 ms clock, 0.40+0.14/21/21+4.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1251 @38.883s 2%: 0.026+5.5+0.44 ms clock, 0.42+0.11/12/11+7.1 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7871488 HeapInuse: 8134656 HeapObjects: 102517 HeapIdle 8839168 HeapReleased 0 HeapSys 16973824
gc 1252 @38.963s 2%: 1.5+0+1.2 ms clock, 24+0.11/12/11+19 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1253 @38.966s 2%: 0.040+0+0.57 ms clock, 0.65+0.11/12/11+9.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1254 @39.052s 2%: 0.024+3.2+0.28 ms clock, 0.39+0.14/8.3/11+4.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1255 @39.168s 2%: 0.019+4.9+0.29 ms clock, 0.31+0.16/17/22+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1256 @39.292s 2%: 0.020+3.0+0.32 ms clock, 0.33+0.20/9.8/11+5.1 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 8025616 HeapInuse: 8282112 HeapObjects: 104810 HeapIdle 8691712 HeapReleased 0 HeapSys 16973824
gc 1257 @39.359s 2%: 0.006+0+0.90 ms clock, 0.10+0.20/9.8/11+14 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1258 @39.360s 2%: 0.059+0+0.49 ms clock, 0.95+0.20/9.8/11+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1259 @39.438s 2%: 0.017+2.6+0.30 ms clock, 0.27+0.15/8.5/10+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1260 @39.539s 2%: 0.021+5.3+0.31 ms clock, 0.33+0.15/18/30+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1261 @39.666s 2%: 0.022+3.5+0.33 ms clock, 0.36+0.17/10/14+5.2 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7974320 HeapInuse: 8232960 HeapObjects: 104048 HeapIdle 8740864 HeapReleased 0 HeapSys 16973824
gc 1262 @39.732s 2%: 0.004+0+1.0 ms clock, 0.066+0.17/10/14+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1263 @39.733s 2%: 0.038+0+0.51 ms clock, 0.62+0.17/10/14+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1264 @39.811s 2%: 0.020+2.6+0.29 ms clock, 0.32+0.21/8.1/10+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1265 @39.915s 2%: 0.020+5.8+0.28 ms clock, 0.32+0.21/18/28+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1266 @40.039s 2%: 0.020+6.0+0.39 ms clock, 0.33+0.14/11/7.7+6.3 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7991376 HeapInuse: 8249344 HeapObjects: 104302 HeapIdle 8724480 HeapReleased 0 HeapSys 16973824
gc 1267 @40.107s 2%: 0.002+0+1.0 ms clock, 0.043+0.14/11/7.7+16 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1268 @40.108s 2%: 0.070+0+0.51 ms clock, 1.1+0.14/11/7.7+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1269 @40.185s 2%: 0.017+2.7+0.32 ms clock, 0.27+0.17/9.1/11+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1270 @40.288s 2%: 0.018+6.2+0.34 ms clock, 0.29+0/18/25+5.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1271 @40.415s 2%: 0.021+3.4+0.28 ms clock, 0.33+0.17/10/15+4.5 ms cpu, 13->13->4 MB, 14 MB goal, 16 P
HeapAlloc: 7785840 HeapInuse: 8044544 HeapObjects: 101244 HeapIdle 8929280 HeapReleased 0 HeapSys 16973824
gc 1272 @40.481s 2%: 0.002+0+0.94 ms clock, 0.040+0.17/10/15+15 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
  100000	      1608 ns/op	      96 B/op	       2 allocs/op
gc 1273 @40.482s 2%: 0.005+0+0.45 ms clock, 0.089+0.17/10/15+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1274 @40.545s 2%: 0.018+0.72+0.37 ms clock, 0.29+0.21/1.1/0.94+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1275 @40.600s 2%: 0.015+1.1+0.30 ms clock, 0.25+0.15/2.5/0.97+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1276 @40.648s 2%: 0.014+0.57+0.40 ms clock, 0.23+0.15/1.1/0.60+6.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1537424 HeapInuse: 1794048 HeapObjects: 8259 HeapIdle 15147008 HeapReleased 0 HeapSys 16941056
gc 1277 @40.654s 2%: 0.004+0+0.55 ms clock, 0.075+0.15/1.1/0.60+8.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndErase1E5-16               	gc 1278 @40.655s 2%: 0.007+0+0.50 ms clock, 0.12+0.15/1.1/0.60+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1279 @40.717s 2%: 0.017+0.71+0.39 ms clock, 0.28+0.12/1.5/0.68+6.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1280 @40.771s 2%: 0.014+0.91+0.32 ms clock, 0.23+0.12/2.7/2.5+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1281 @40.820s 2%: 0.016+0.70+0.69 ms clock, 0.27+0.15/1.0/0.83+11 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554704 HeapInuse: 1794048 HeapObjects: 8517 HeapIdle 15179776 HeapReleased 0 HeapSys 16973824
gc 1282 @40.826s 2%: 0.002+0+0.55 ms clock, 0.035+0.15/1.0/0.83+8.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1283 @40.827s 2%: 0.075+0+0.52 ms clock, 1.2+0.15/1.0/0.83+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1284 @40.891s 2%: 0.020+0.64+0.26 ms clock, 0.32+0.16/1.4/0.88+4.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1285 @40.945s 2%: 0.016+0.88+0.34 ms clock, 0.26+0.17/1.9/2.3+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1286 @40.993s 2%: 0.014+0.69+0.31 ms clock, 0.23+0.15/1.1/0.98+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554688 HeapInuse: 1818624 HeapObjects: 8516 HeapIdle 15155200 HeapReleased 0 HeapSys 16973824
gc 1287 @40.999s 2%: 0.004+0+0.62 ms clock, 0.077+0.15/1.1/0.98+9.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1288 @40.999s 2%: 0.074+0+0.43 ms clock, 1.1+0.15/1.1/0.98+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1289 @41.061s 2%: 0.018+0.71+0.31 ms clock, 0.29+0.19/1.3/0.59+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1290 @41.116s 2%: 0.020+0.89+0.37 ms clock, 0.32+0.16/2.0/2.0+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1291 @41.163s 2%: 0.012+0.58+0.32 ms clock, 0.20+0.15/1.3/0.45+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554720 HeapInuse: 1810432 HeapObjects: 8518 HeapIdle 15163392 HeapReleased 0 HeapSys 16973824
gc 1292 @41.169s 2%: 0.053+0+0.54 ms clock, 0.85+0.15/1.3/0.45+8.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1293 @41.169s 2%: 0.072+0+0.44 ms clock, 1.1+0.15/1.3/0.45+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1294 @41.231s 2%: 0.014+0.65+0.38 ms clock, 0.23+0.17/1.4/0.80+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1295 @41.285s 2%: 0.015+1.0+0.34 ms clock, 0.24+0.13/2.9/0.75+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1296 @41.334s 2%: 0.015+0.60+0.31 ms clock, 0.24+0.11/1.2/0.77+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554704 HeapInuse: 1810432 HeapObjects: 8517 HeapIdle 15163392 HeapReleased 0 HeapSys 16973824
gc 1297 @41.339s 2%: 0.002+0+0.54 ms clock, 0.035+0.11/1.2/0.77+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1298 @41.340s 2%: 0.10+0+0.46 ms clock, 1.7+0.11/1.2/0.77+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1299 @41.401s 2%: 0.018+0.66+0.28 ms clock, 0.29+0.17/1.4/1.1+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1300 @41.457s 2%: 0.017+1.1+0.35 ms clock, 0.27+0.16/2.7/1.9+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1301 @41.505s 2%: 0.014+0.56+0.41 ms clock, 0.23+0.10/1.1/1.1+6.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1537472 HeapInuse: 1810432 HeapObjects: 8260 HeapIdle 15163392 HeapReleased 0 HeapSys 16973824
gc 1302 @41.510s 2%: 0.003+0+0.52 ms clock, 0.062+0.10/1.1/1.1+8.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1303 @41.511s 2%: 0.072+0+0.44 ms clock, 1.1+0.10/1.1/1.1+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1304 @41.572s 2%: 0.019+0.67+0.31 ms clock, 0.31+0.16/1.4/0.53+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1305 @41.628s 2%: 0.017+1.2+0.40 ms clock, 0.28+0.15/3.5/1.0+6.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1306 @41.677s 2%: 0.020+0.67+0.29 ms clock, 0.33+0.16/1.4/0.66+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554720 HeapInuse: 1818624 HeapObjects: 8518 HeapIdle 15155200 HeapReleased 0 HeapSys 16973824
gc 1307 @41.683s 2%: 0.002+0+0.51 ms clock, 0.036+0.16/1.4/0.66+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1308 @41.683s 2%: 0.004+0+0.43 ms clock, 0.073+0.16/1.4/0.66+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1309 @41.745s 2%: 0.016+0.65+0.37 ms clock, 0.26+0.17/1.4/0.57+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1310 @41.800s 2%: 0.017+1.1+0.34 ms clock, 0.27+0.11/2.6/1.7+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1311 @41.850s 2%: 0.014+0.63+0.31 ms clock, 0.23+0.16/1.0/0.50+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1537488 HeapInuse: 1785856 HeapObjects: 8261 HeapIdle 15187968 HeapReleased 0 HeapSys 16973824
gc 1312 @41.856s 2%: 0.003+0+0.56 ms clock, 0.058+0.16/1.0/0.50+8.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1313 @41.857s 2%: 0.019+0+0.72 ms clock, 0.31+0.16/1.0/0.50+11 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1314 @41.919s 2%: 0.019+0.75+0.31 ms clock, 0.30+0.15/1.3/1.3+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1315 @41.975s 2%: 0.015+1.2+0.39 ms clock, 0.25+0.17/2.5/1.4+6.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1316 @42.023s 2%: 0.015+1.3+0.31 ms clock, 0.25+0.15/1.1/0.57+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1537488 HeapInuse: 1785856 HeapObjects: 8261 HeapIdle 15187968 HeapReleased 0 HeapSys 16973824
gc 1317 @42.029s 2%: 0.003+0+0.61 ms clock, 0.060+0.15/1.1/0.57+9.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1318 @42.029s 2%: 0.073+0+0.39 ms clock, 1.1+0.15/1.1/0.57+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1319 @42.091s 2%: 0.017+0.69+0.34 ms clock, 0.28+0.17/1.6/0.24+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1320 @42.146s 2%: 0.015+1.0+0.40 ms clock, 0.25+0.16/2.1/2.4+6.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1321 @42.194s 2%: 0.015+0.62+0.32 ms clock, 0.24+0.15/1.1/0.86+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554704 HeapInuse: 1826816 HeapObjects: 8517 HeapIdle 15147008 HeapReleased 0 HeapSys 16973824
gc 1322 @42.200s 2%: 0.004+0+0.59 ms clock, 0.070+0.15/1.1/0.86+9.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1323 @42.200s 2%: 0.070+0+0.52 ms clock, 1.1+0.15/1.1/0.86+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1324 @42.262s 2%: 0.015+0.92+0.41 ms clock, 0.24+0.088/1.8/0.54+6.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1325 @42.318s 2%: 0.015+1.0+0.37 ms clock, 0.25+0.13/2.5/2.1+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1326 @42.367s 2%: 0.015+0.65+0.29 ms clock, 0.24+0.17/1.3/0.61+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1537568 HeapInuse: 1794048 HeapObjects: 8262 HeapIdle 15179776 HeapReleased 0 HeapSys 16973824
gc 1327 @42.372s 2%: 0.002+0+0.82 ms clock, 0.038+0.17/1.3/0.61+13 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1328 @42.373s 2%: 0.073+0+0.48 ms clock, 1.1+0.17/1.3/0.61+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1329 @42.435s 2%: 0.022+0.72+0.35 ms clock, 0.35+0.18/1.5/0.55+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1330 @42.492s 2%: 0.015+1.0+0.30 ms clock, 0.24+0.11/2.9/1.8+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1331 @42.540s 2%: 0.021+0.70+0.29 ms clock, 0.34+0.18/1.3/0.60+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554720 HeapInuse: 1818624 HeapObjects: 8518 HeapIdle 15155200 HeapReleased 0 HeapSys 16973824
gc 1332 @42.546s 2%: 0.026+0+0.58 ms clock, 0.42+0.18/1.3/0.60+9.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1333 @42.546s 2%: 0.004+0+0.42 ms clock, 0.076+0.18/1.3/0.60+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1334 @42.609s 2%: 0.021+0.73+0.34 ms clock, 0.34+0.19/1.6/0.76+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1335 @42.665s 2%: 0.018+0.96+0.30 ms clock, 0.29+0.16/2.3/2.1+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1336 @42.713s 2%: 0.012+0.59+0.33 ms clock, 0.20+0.15/1.0/0.53+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554704 HeapInuse: 1802240 HeapObjects: 8517 HeapIdle 15171584 HeapReleased 0 HeapSys 16973824
gc 1337 @42.719s 2%: 0.032+0+0.56 ms clock, 0.52+0.15/1.0/0.53+9.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1338 @42.720s 2%: 0.038+0+0.46 ms clock, 0.62+0.15/1.0/0.53+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1339 @42.782s 2%: 0.018+0.70+0.38 ms clock, 0.29+0.16/1.3/0.81+6.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1340 @42.838s 2%: 0.017+1.3+0.35 ms clock, 0.27+0.17/2.7/0.71+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1341 @42.891s 2%: 0.021+0.63+0.32 ms clock, 0.34+0.27/1.2/1.0+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554720 HeapInuse: 1818624 HeapObjects: 8518 HeapIdle 15155200 HeapReleased 0 HeapSys 16973824
gc 1342 @42.896s 2%: 0.004+0+0.55 ms clock, 0.072+0.27/1.2/1.0+8.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1343 @42.897s 2%: 0.005+0+0.58 ms clock, 0.081+0.27/1.2/1.0+9.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1344 @42.960s 2%: 0.017+0.84+0.34 ms clock, 0.27+0.24/1.1/0.92+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1345 @43.016s 2%: 0.020+0.88+0.31 ms clock, 0.32+0.15/2.5/1.5+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1346 @43.066s 2%: 0.018+0.75+0.33 ms clock, 0.29+0.20/1.3/0.30+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1537488 HeapInuse: 1785856 HeapObjects: 8261 HeapIdle 15187968 HeapReleased 0 HeapSys 16973824
gc 1347 @43.071s 2%: 0.002+0+0.51 ms clock, 0.035+0.20/1.3/0.30+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1348 @43.072s 2%: 0.004+0+0.47 ms clock, 0.079+0.20/1.3/0.30+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1349 @43.134s 2%: 0.018+0.74+0.37 ms clock, 0.28+0.19/1.5/0.62+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1350 @43.191s 2%: 0.017+1.1+0.34 ms clock, 0.27+0.19/2.7/0.64+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 1351 @43.240s 2%: 0.014+0.54+0.39 ms clock, 0.23+0.16/1.2/0.37+6.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1554704 HeapInuse: 1810432 HeapObjects: 8517 HeapIdle 15163392 HeapReleased 0 HeapSys 16973824
gc 1352 @43.246s 2%: 0.004+0+0.61 ms clock, 0.071+0.16/1.2/0.37+9.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
  100000	      1738 ns/op	      96 B/op	       2 allocs/op
gc 1353 @43.247s 2%: 0.007+0+0.59 ms clock, 0.11+0.16/1.2/0.37+9.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1354 @43.385s 2%: 0.018+0.42+0.37 ms clock, 0.29+0.14/0.82/0.52+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904456 HeapInuse: 3170304 HeapObjects: 38947 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1355 @43.464s 2%: 0.024+0+0.73 ms clock, 0.39+0.14/0.82/0.52+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndEraseWithPool1E5-16       	gc 1356 @43.465s 2%: 0.003+0+0.47 ms clock, 0.053+0.14/0.82/0.52+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1357 @43.602s 2%: 0.016+0.47+0.38 ms clock, 0.27+0.19/0.87/0.31+6.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904536 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1358 @43.652s 2%: 0.002+0+0.54 ms clock, 0.035+0.19/0.87/0.31+8.7 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1359 @43.652s 2%: 0.045+0+0.49 ms clock, 0.72+0.19/0.87/0.31+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1360 @43.761s 2%: 0.017+0.60+0.37 ms clock, 0.28+0.16/1.1/0.55+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904536 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1361 @43.841s 2%: 0.005+0+0.74 ms clock, 0.080+0.16/1.1/0.55+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1362 @43.842s 2%: 0.046+0+0.49 ms clock, 0.74+0.16/1.1/0.55+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1363 @43.982s 2%: 0.023+0.51+0.32 ms clock, 0.37+0.18/0.85/0.57+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904520 HeapInuse: 3170304 HeapObjects: 38949 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1364 @44.061s 2%: 0.005+0+0.64 ms clock, 0.087+0.18/0.85/0.57+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1365 @44.062s 2%: 0.042+0+0.53 ms clock, 0.68+0.18/0.85/0.57+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1366 @44.199s 2%: 0.016+0.55+0.33 ms clock, 0.26+0.16/0.80/0.54+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904552 HeapInuse: 3170304 HeapObjects: 38951 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1367 @44.282s 2%: 0.005+0+0.70 ms clock, 0.083+0.16/0.80/0.54+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1368 @44.283s 2%: 0.007+0+0.65 ms clock, 0.12+0.16/0.80/0.54+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1369 @44.423s 2%: 0.019+0.66+0.34 ms clock, 0.31+0.19/0.92/0.46+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904536 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1370 @44.503s 2%: 0.003+0+0.66 ms clock, 0.048+0.19/0.92/0.46+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1371 @44.504s 2%: 0.028+0+0.55 ms clock, 0.45+0.19/0.92/0.46+8.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1372 @44.642s 2%: 0.019+0.53+0.42 ms clock, 0.31+0.17/1.0/0.47+6.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904552 HeapInuse: 3170304 HeapObjects: 38951 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1373 @44.722s 2%: 0.003+0+0.69 ms clock, 0.049+0.17/1.0/0.47+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1374 @44.723s 2%: 0.078+0+0.53 ms clock, 1.2+0.17/1.0/0.47+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1375 @44.861s 2%: 0.028+0.67+0.45 ms clock, 0.46+0.22/0.95/0.78+7.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904536 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1376 @44.943s 2%: 0.010+0+0.67 ms clock, 0.16+0.22/0.95/0.78+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1377 @44.944s 2%: 0.075+0+0.45 ms clock, 1.2+0.22/0.95/0.78+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1378 @45.084s 2%: 0.018+0.54+0.33 ms clock, 0.29+0.22/0.87/0.57+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904552 HeapInuse: 3170304 HeapObjects: 38951 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1379 @45.164s 2%: 0.018+0+0.62 ms clock, 0.29+0.22/0.87/0.57+9.9 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1380 @45.164s 2%: 0.073+0+0.44 ms clock, 1.1+0.22/0.87/0.57+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1381 @45.302s 2%: 0.016+0.57+0.28 ms clock, 0.26+0.16/1.0/0.89+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904536 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1382 @45.383s 2%: 0.005+0+0.71 ms clock, 0.095+0.16/1.0/0.89+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1383 @45.384s 2%: 0.030+0+0.50 ms clock, 0.49+0.16/1.0/0.89+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1384 @45.522s 1%: 0.019+0.51+0.42 ms clock, 0.31+0.16/1.1/0.40+6.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904520 HeapInuse: 3178496 HeapObjects: 38949 HeapIdle 13795328 HeapReleased 0 HeapSys 16973824
gc 1385 @45.602s 1%: 0.034+0+0.65 ms clock, 0.55+0.16/1.1/0.40+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1386 @45.603s 1%: 0.068+0+0.53 ms clock, 1.0+0.16/1.1/0.40+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1387 @45.741s 1%: 0.018+0.61+0.37 ms clock, 0.30+0.18/1.0/0.46+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904536 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1388 @45.822s 1%: 0.005+0+0.71 ms clock, 0.093+0.18/1.0/0.46+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1389 @45.823s 1%: 0.076+0+0.47 ms clock, 1.2+0.18/1.0/0.46+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1390 @45.970s 1%: 0.021+0.54+0.54 ms clock, 0.34+0.11/1.3/0.27+8.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904600 HeapInuse: 3178496 HeapObjects: 38951 HeapIdle 13795328 HeapReleased 0 HeapSys 16973824
gc 1391 @46.052s 1%: 0.029+0+0.63 ms clock, 0.47+0.11/1.3/0.27+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1392 @46.053s 1%: 0.030+0+0.54 ms clock, 0.48+0.11/1.3/0.27+8.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1393 @46.191s 1%: 0.018+0.53+0.36 ms clock, 0.29+0.17/0.83/0.68+5.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904584 HeapInuse: 3170304 HeapObjects: 38950 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1394 @46.271s 1%: 0.004+0+0.68 ms clock, 0.073+0.17/0.83/0.68+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1395 @46.272s 1%: 0.037+0+0.43 ms clock, 0.59+0.17/0.83/0.68+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1396 @46.408s 1%: 0.021+0.59+0.53 ms clock, 0.35+0.16/1.2/0.41+8.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904600 HeapInuse: 3178496 HeapObjects: 38951 HeapIdle 13795328 HeapReleased 0 HeapSys 16973824
gc 1397 @46.489s 1%: 0.004+0+0.65 ms clock, 0.071+0.16/1.2/0.41+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1398 @46.490s 1%: 0.10+0+0.46 ms clock, 1.7+0.16/1.2/0.41+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1399 @46.630s 1%: 0.020+0.51+0.40 ms clock, 0.32+0.14/0.88/0.80+6.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2904600 HeapInuse: 3170304 HeapObjects: 38951 HeapIdle 13803520 HeapReleased 0 HeapSys 16973824
gc 1400 @46.710s 1%: 0.003+0+0.63 ms clock, 0.056+0.14/0.88/0.80+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
  100000	      2201 ns/op	      56 B/op	       1 allocs/op
gc 1401 @46.711s 1%: 0.034+0+0.58 ms clock, 0.55+0.14/0.88/0.80+9.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 3860152 HeapInuse: 4218880 HeapObjects: 6625 HeapIdle 12754944 HeapReleased 0 HeapSys 16973824
gc 1402 @46.742s 1%: 0.004+0+0.63 ms clock, 0.076+0.14/0.88/0.80+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E5-16                   	gc 1403 @46.743s 1%: 0.007+0+0.53 ms clock, 0.12+0.14/0.88/0.80+8.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3832496 HeapInuse: 4210688 HeapObjects: 6525 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1404 @46.774s 1%: 0.048+0+0.61 ms clock, 0.78+0.14/0.88/0.80+9.8 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1405 @46.775s 1%: 0.032+0+0.54 ms clock, 0.52+0.14/0.88/0.80+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3826408 HeapInuse: 4202496 HeapObjects: 6459 HeapIdle 12771328 HeapReleased 0 HeapSys 16973824
gc 1406 @46.805s 1%: 0.002+0+0.60 ms clock, 0.043+0.14/0.88/0.80+9.6 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1407 @46.806s 1%: 0.025+0+0.48 ms clock, 0.41+0.14/0.88/0.80+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3832080 HeapInuse: 4210688 HeapObjects: 6520 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1408 @46.837s 1%: 0.005+0+0.57 ms clock, 0.082+0.14/0.88/0.80+9.2 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1409 @46.838s 1%: 0.074+0+0.59 ms clock, 1.1+0.14/0.88/0.80+9.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3825992 HeapInuse: 4202496 HeapObjects: 6453 HeapIdle 12771328 HeapReleased 0 HeapSys 16973824
gc 1410 @46.870s 1%: 0.005+0+0.63 ms clock, 0.081+0.14/0.88/0.80+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1411 @46.871s 1%: 0.004+0+0.48 ms clock, 0.078+0.14/0.88/0.80+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3832104 HeapInuse: 4210688 HeapObjects: 6519 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1412 @46.902s 1%: 0.014+0+0.64 ms clock, 0.23+0.14/0.88/0.80+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1413 @46.903s 1%: 0.014+0+0.45 ms clock, 0.23+0.14/0.88/0.80+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3851328 HeapInuse: 4210688 HeapObjects: 6537 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1414 @46.933s 1%: 0.036+0+0.61 ms clock, 0.58+0.14/0.88/0.80+9.7 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1415 @46.934s 1%: 0.085+0+0.47 ms clock, 1.3+0.14/0.88/0.80+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3852208 HeapInuse: 4210688 HeapObjects: 6543 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1416 @46.964s 1%: 0.046+0+0.63 ms clock, 0.74+0.14/0.88/0.80+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1417 @46.965s 1%: 0.041+0+0.51 ms clock, 0.67+0.14/0.88/0.80+8.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3829256 HeapInuse: 4210688 HeapObjects: 6489 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1418 @46.996s 1%: 0.004+0+0.55 ms clock, 0.072+0.14/0.88/0.80+8.9 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1419 @46.997s 1%: 0.048+0+0.46 ms clock, 0.76+0.14/0.88/0.80+7.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3850136 HeapInuse: 4210688 HeapObjects: 6523 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1420 @47.027s 1%: 0.002+0+0.56 ms clock, 0.037+0.14/0.88/0.80+9.0 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1421 @47.028s 1%: 0.040+0+0.50 ms clock, 0.65+0.14/0.88/0.80+8.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3850488 HeapInuse: 4210688 HeapObjects: 6526 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1422 @47.059s 1%: 0.004+0+0.61 ms clock, 0.073+0.14/0.88/0.80+9.8 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1423 @47.060s 1%: 0.045+0+0.49 ms clock, 0.72+0.14/0.88/0.80+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3823376 HeapInuse: 4202496 HeapObjects: 6425 HeapIdle 12771328 HeapReleased 0 HeapSys 16973824
gc 1424 @47.091s 1%: 0.049+0+0.56 ms clock, 0.79+0.14/0.88/0.80+9.0 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1425 @47.091s 1%: 0.036+0+0.45 ms clock, 0.57+0.14/0.88/0.80+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3826928 HeapInuse: 4210688 HeapObjects: 6462 HeapIdle 12763136 HeapReleased 0 HeapSys 16973824
gc 1426 @47.122s 1%: 0.034+0+0.53 ms clock, 0.55+0.14/0.88/0.80+8.6 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1427 @47.123s 1%: 0.037+0+0.49 ms clock, 0.59+0.14/0.88/0.80+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3854184 HeapInuse: 4218880 HeapObjects: 6566 HeapIdle 12754944 HeapReleased 0 HeapSys 16973824
gc 1428 @47.154s 1%: 0.004+0+0.59 ms clock, 0.075+0.14/0.88/0.80+9.5 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
  100000	       313 ns/op	      36 B/op	       0 allocs/op
gc 1429 @47.155s 1%: 0.008+0+0.49 ms clock, 0.12+0.14/0.88/0.80+7.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1430 @47.172s 1%: 0.015+0.37+0.37 ms clock, 0.25+0.097/0.21/0.22+6.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3679048 HeapInuse: 4022272 HeapObjects: 4671 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1431 @47.198s 1%: 0.005+0+0.63 ms clock, 0.087+0.097/0.21/0.22+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E5-16                    	gc 1432 @47.199s 1%: 0.005+0+0.54 ms clock, 0.083+0.097/0.21/0.22+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1433 @47.216s 1%: 0.015+0.25+0.31 ms clock, 0.25+0.10/0.12/0.32+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3690216 HeapInuse: 4022272 HeapObjects: 4747 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1434 @47.243s 1%: 0.024+0+0.73 ms clock, 0.38+0.10/0.12/0.32+11 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1435 @47.244s 1%: 0.013+0+0.49 ms clock, 0.21+0.10/0.12/0.32+7.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1436 @47.260s 1%: 0.014+0.30+0.36 ms clock, 0.23+0.10/0.22/0.27+5.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3679080 HeapInuse: 4014080 HeapObjects: 4631 HeapIdle 12959744 HeapReleased 0 HeapSys 16973824
gc 1437 @47.287s 1%: 0.026+0+0.60 ms clock, 0.42+0.10/0.22/0.27+9.7 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1438 @47.287s 1%: 0.068+0+0.65 ms clock, 1.0+0.10/0.22/0.27+10 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1439 @47.304s 1%: 0.012+0.29+0.36 ms clock, 0.20+0.10/0.090/0.45+5.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3694056 HeapInuse: 4030464 HeapObjects: 4787 HeapIdle 12943360 HeapReleased 0 HeapSys 16973824
gc 1440 @47.331s 1%: 0.029+0+0.67 ms clock, 0.46+0.10/0.090/0.45+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1441 @47.332s 1%: 0.045+0+0.45 ms clock, 0.73+0.10/0.090/0.45+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1442 @47.349s 1%: 0.015+0.31+0.46 ms clock, 0.24+0.10/0.23/0.30+7.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3692792 HeapInuse: 4030464 HeapObjects: 4774 HeapIdle 12943360 HeapReleased 0 HeapSys 16973824
gc 1443 @47.376s 1%: 0.002+0+0.64 ms clock, 0.042+0.10/0.23/0.30+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1444 @47.377s 1%: 0.005+0+0.54 ms clock, 0.081+0.10/0.23/0.30+8.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1445 @47.394s 1%: 0.012+0.25+0.36 ms clock, 0.20+0.097/0.21/0.15+5.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3679864 HeapInuse: 4022272 HeapObjects: 4682 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1446 @47.421s 1%: 0.004+0+0.59 ms clock, 0.064+0.097/0.21/0.15+9.5 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1447 @47.422s 2%: 0.045+0+0.56 ms clock, 0.72+0.097/0.21/0.15+9.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1448 @47.439s 2%: 0.014+0.44+0.34 ms clock, 0.22+0.15/0.23/0.42+5.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3682168 HeapInuse: 4022272 HeapObjects: 4706 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1449 @47.466s 2%: 0.002+0+0.58 ms clock, 0.038+0.15/0.23/0.42+9.3 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1450 @47.466s 2%: 0.048+0+0.47 ms clock, 0.77+0.15/0.23/0.42+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1451 @47.483s 2%: 0.014+0.28+0.32 ms clock, 0.22+0.10/0.15/0.34+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3688856 HeapInuse: 4022272 HeapObjects: 4733 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1452 @47.511s 2%: 0.004+0+0.61 ms clock, 0.077+0.10/0.15/0.34+9.8 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1453 @47.512s 2%: 0.044+0+0.49 ms clock, 0.71+0.10/0.15/0.34+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1454 @47.528s 2%: 0.015+0.39+0.31 ms clock, 0.25+0.10/0.20/0.32+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3690968 HeapInuse: 4030464 HeapObjects: 4755 HeapIdle 12943360 HeapReleased 0 HeapSys 16973824
gc 1455 @47.554s 2%: 0.004+0+0.65 ms clock, 0.072+0.10/0.20/0.32+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1456 @47.555s 2%: 0.007+0+0.59 ms clock, 0.12+0.10/0.20/0.32+9.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1457 @47.572s 2%: 0.015+0.30+0.31 ms clock, 0.24+0.12/0.17/0.52+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3690200 HeapInuse: 4022272 HeapObjects: 4747 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1458 @47.598s 2%: 0.027+0+0.58 ms clock, 0.43+0.12/0.17/0.52+9.3 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1459 @47.599s 2%: 0.044+0+0.53 ms clock, 0.71+0.12/0.17/0.52+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1460 @47.616s 2%: 0.016+0.29+0.32 ms clock, 0.25+0.10/0.082/0.30+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3686264 HeapInuse: 4022272 HeapObjects: 4706 HeapIdle 12951552 HeapReleased 0 HeapSys 16973824
gc 1461 @47.642s 2%: 0.004+0+0.58 ms clock, 0.078+0.10/0.082/0.30+9.3 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 1462 @47.643s 2%: 0.042+0+0.51 ms clock, 0.68+0.10/0.082/0.30+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1463 @47.659s 2%: 0.015+0.28+0.28 ms clock, 0.24+0.10/0.098/0.51+4.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3692600 HeapInuse: 4030464 HeapObjects: 4772 HeapIdle 12943360 HeapReleased 0 HeapSys 16973824
gc 1464 @47.687s 2%: 0.002+0+0.74 ms clock, 0.040+0.10/0.098/0.51+11 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
  100000	       119 ns/op	       8 B/op	       0 allocs/op
gc 1465 @47.688s 2%: 0.006+0+0.49 ms clock, 0.097+0.10/0.098/0.51+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 771576 HeapInuse: 1089536 HeapObjects: 1389 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1466 @47.704s 2%: 0.002+0+0.48 ms clock, 0.039+0.10/0.098/0.51+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E5-16           	gc 1467 @47.705s 2%: 0.040+0+0.45 ms clock, 0.64+0.10/0.098/0.51+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 773696 HeapInuse: 1097728 HeapObjects: 1415 HeapIdle 15876096 HeapReleased 0 HeapSys 16973824
gc 1468 @47.720s 2%: 0.016+0+0.66 ms clock, 0.26+0.10/0.098/0.51+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1469 @47.721s 2%: 0.047+0+0.44 ms clock, 0.75+0.10/0.098/0.51+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 769064 HeapInuse: 1089536 HeapObjects: 1369 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1470 @47.737s 2%: 0.072+0+0.50 ms clock, 1.1+0.10/0.098/0.51+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1471 @47.737s 2%: 0.061+0+0.45 ms clock, 0.97+0.10/0.098/0.51+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 769392 HeapInuse: 1089536 HeapObjects: 1370 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1472 @47.752s 2%: 0.049+0+0.46 ms clock, 0.78+0.10/0.098/0.51+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1473 @47.753s 2%: 0.059+0+0.42 ms clock, 0.94+0.10/0.098/0.51+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 768968 HeapInuse: 1089536 HeapObjects: 1364 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1474 @47.769s 2%: 0.050+0+0.52 ms clock, 0.80+0.10/0.098/0.51+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1475 @47.769s 2%: 0.043+0+0.44 ms clock, 0.69+0.10/0.098/0.51+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 768112 HeapInuse: 1089536 HeapObjects: 1356 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1476 @47.785s 2%: 0.002+0+0.48 ms clock, 0.038+0.10/0.098/0.51+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1477 @47.785s 2%: 0.040+0+0.46 ms clock, 0.65+0.10/0.098/0.51+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 767144 HeapInuse: 1089536 HeapObjects: 1345 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1478 @47.801s 2%: 0.014+0+0.67 ms clock, 0.23+0.10/0.098/0.51+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1479 @47.802s 2%: 0.004+0+0.46 ms clock, 0.072+0.10/0.098/0.51+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 769056 HeapInuse: 1089536 HeapObjects: 1368 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1480 @47.817s 2%: 0.014+0+0.56 ms clock, 0.22+0.10/0.098/0.51+8.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1481 @47.818s 2%: 0.044+0+0.41 ms clock, 0.71+0.10/0.098/0.51+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 767856 HeapInuse: 1089536 HeapObjects: 1354 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1482 @47.833s 2%: 0.003+0+0.49 ms clock, 0.054+0.10/0.098/0.51+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1483 @47.834s 2%: 0.049+0+0.43 ms clock, 0.78+0.10/0.098/0.51+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 766376 HeapInuse: 1089536 HeapObjects: 1337 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1484 @47.849s 2%: 0.042+0+0.45 ms clock, 0.68+0.10/0.098/0.51+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1485 @47.850s 2%: 0.047+0+0.39 ms clock, 0.75+0.10/0.098/0.51+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 765456 HeapInuse: 1089536 HeapObjects: 1329 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1486 @47.866s 2%: 0.004+0+1.0 ms clock, 0.069+0.10/0.098/0.51+16 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1487 @47.868s 2%: 0.041+0+0.53 ms clock, 0.67+0.10/0.098/0.51+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 767088 HeapInuse: 1089536 HeapObjects: 1346 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1488 @47.884s 2%: 0.002+0+0.41 ms clock, 0.038+0.10/0.098/0.51+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1489 @47.884s 2%: 0.044+0+0.47 ms clock, 0.71+0.10/0.098/0.51+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 771216 HeapInuse: 1089536 HeapObjects: 1384 HeapIdle 15884288 HeapReleased 0 HeapSys 16973824
gc 1490 @47.900s 2%: 0.022+0+0.47 ms clock, 0.36+0.10/0.098/0.51+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
  100000	       153 ns/op	       5 B/op	       0 allocs/op
gc 1491 @47.900s 2%: 0.032+0+0.44 ms clock, 0.51+0.10/0.098/0.51+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1724176 HeapInuse: 1974272 HeapObjects: 310 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1492 @47.913s 2%: 0.050+0+0.52 ms clock, 0.81+0.10/0.098/0.51+8.3 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E5-16    	gc 1493 @47.914s 2%: 0.003+0+0.44 ms clock, 0.063+0.10/0.098/0.51+7.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1494 @47.927s 2%: 0.027+0+0.49 ms clock, 0.43+0.10/0.098/0.51+7.8 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1495 @47.927s 2%: 0.075+0+0.48 ms clock, 1.2+0.10/0.098/0.51+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1496 @47.940s 2%: 0.002+0+0.42 ms clock, 0.035+0.10/0.098/0.51+6.8 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1497 @47.940s 2%: 0.084+0+0.47 ms clock, 1.3+0.10/0.098/0.51+7.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1498 @47.953s 2%: 0.001+0+0.46 ms clock, 0.029+0.10/0.098/0.51+7.3 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1499 @47.954s 2%: 0.076+0+0.44 ms clock, 1.2+0.10/0.098/0.51+7.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1500 @47.967s 2%: 0.044+0+0.45 ms clock, 0.70+0.10/0.098/0.51+7.3 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1501 @47.967s 2%: 0.078+0+0.43 ms clock, 1.2+0.10/0.098/0.51+6.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1502 @47.980s 2%: 0.005+0+0.44 ms clock, 0.083+0.10/0.098/0.51+7.1 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1503 @47.980s 2%: 0.045+0+0.48 ms clock, 0.72+0.10/0.098/0.51+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1504 @47.993s 2%: 0.005+0+0.46 ms clock, 0.086+0.10/0.098/0.51+7.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1505 @47.994s 2%: 0.053+0+0.49 ms clock, 0.85+0.10/0.098/0.51+7.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1506 @48.006s 2%: 0.038+0+0.45 ms clock, 0.62+0.10/0.098/0.51+7.2 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1507 @48.007s 2%: 0.075+0+0.69 ms clock, 1.2+0.10/0.098/0.51+11 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1508 @48.020s 2%: 0.002+0+0.46 ms clock, 0.039+0.10/0.098/0.51+7.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1509 @48.020s 2%: 0.045+0+0.42 ms clock, 0.72+0.10/0.098/0.51+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1510 @48.033s 2%: 0.002+0+0.43 ms clock, 0.037+0.10/0.098/0.51+6.9 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1511 @48.034s 2%: 0.070+0+0.48 ms clock, 1.1+0.10/0.098/0.51+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1512 @48.046s 2%: 0.002+0+0.47 ms clock, 0.036+0.10/0.098/0.51+7.5 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1513 @48.047s 2%: 0.062+0+0.50 ms clock, 0.99+0.10/0.098/0.51+8.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1514 @48.059s 2%: 0.008+0+0.47 ms clock, 0.13+0.10/0.098/0.51+7.5 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 1515 @48.060s 2%: 0.048+0+0.45 ms clock, 0.77+0.10/0.098/0.51+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1724256 HeapInuse: 1974272 HeapObjects: 313 HeapIdle 14999552 HeapReleased 0 HeapSys 16973824
gc 1516 @48.073s 2%: 0.001+0+0.56 ms clock, 0.031+0.10/0.098/0.51+8.9 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
  100000	       127 ns/op	      15 B/op	       0 allocs/op
gc 1517 @48.074s 2%: 0.040+0+0.56 ms clock, 0.64+0.10/0.098/0.51+9.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 1518 @48.076s 2%: 0.016+28+0.38 ms clock, 0.26+0/0.14/28+6.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204080 HeapInuse: 8454144 HeapObjects: 310 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1519 @48.546s 2%: 0.005+0+0.52 ms clock, 0.080+0/0.14/28+8.4 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
#BenchmarkSort1E6-16                               	gc 1520 @48.546s 2%: 0.004+0+0.54 ms clock, 0.075+0/0.14/28+8.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1521 @48.548s 2%: 0.014+29+0.33 ms clock, 0.22+0/0.13/29+5.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1522 @49.022s 2%: 0.002+0+0.58 ms clock, 0.047+0/0.13/29+9.2 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1523 @49.023s 2%: 0.072+0+0.52 ms clock, 1.1+0/0.13/29+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1524 @49.025s 2%: 0.013+29+0.34 ms clock, 0.22+0/0.015/29+5.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1525 @49.500s 2%: 0.003+0+0.61 ms clock, 0.055+0/0.015/29+9.8 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1526 @49.501s 2%: 0.003+0+0.55 ms clock, 0.062+0/0.015/29+8.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1527 @49.503s 2%: 0.016+29+0.36 ms clock, 0.26+0/0.15/29+5.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1528 @49.976s 1%: 0.003+0+0.56 ms clock, 0.052+0/0.15/29+9.0 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1529 @49.976s 1%: 0.072+0+0.49 ms clock, 1.1+0/0.15/29+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1530 @49.979s 1%: 0.015+29+0.40 ms clock, 0.25+0/0.15/29+6.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1531 @50.456s 1%: 0.003+0+0.75 ms clock, 0.053+0/0.15/29+12 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1532 @50.457s 1%: 0.066+0+0.47 ms clock, 1.0+0/0.15/29+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1533 @50.460s 1%: 0.015+29+0.36 ms clock, 0.24+0/0.12/29+5.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1534 @50.936s 1%: 0.005+0+0.56 ms clock, 0.082+0/0.12/29+9.0 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1535 @50.937s 1%: 0.052+0+0.46 ms clock, 0.83+0/0.12/29+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1536 @50.939s 1%: 0.015+28+0.92 ms clock, 0.24+0/0.16/28+14 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1537 @51.412s 1%: 0.003+0+0.61 ms clock, 0.061+0/0.16/28+9.7 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1538 @51.413s 1%: 0.040+0+0.46 ms clock, 0.64+0/0.16/28+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1539 @51.415s 1%: 0.013+29+0.37 ms clock, 0.21+0/29/0.16+6.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1540 @51.889s 1%: 0.005+0+0.54 ms clock, 0.082+0/29/0.16+8.6 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1541 @51.889s 1%: 0.006+0+0.46 ms clock, 0.10+0/29/0.16+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1542 @51.891s 1%: 0.013+29+0.44 ms clock, 0.21+0/0.024/29+7.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1543 @52.364s 1%: 0.005+0+0.51 ms clock, 0.080+0/0.024/29+8.2 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1544 @52.365s 1%: 0.014+0+0.65 ms clock, 0.23+0/0.024/29+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1545 @52.367s 1%: 0.016+29+0.38 ms clock, 0.25+0/0.10/29+6.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1546 @52.846s 1%: 0.003+0+0.57 ms clock, 0.059+0/0.10/29+9.1 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1547 @52.847s 1%: 0.004+0+0.47 ms clock, 0.076+0/0.10/29+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1548 @52.849s 1%: 0.014+29+0.31 ms clock, 0.23+0/29/0.21+5.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1549 @53.320s 1%: 0.002+0+0.60 ms clock, 0.046+0/29/0.21+9.6 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1550 @53.321s 1%: 0.071+0+0.52 ms clock, 1.1+0/29/0.21+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1551 @53.323s 1%: 0.017+28+0.35 ms clock, 0.27+0/0.25/28+5.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1552 @53.801s 1%: 0.005+0+0.48 ms clock, 0.082+0/0.25/28+7.7 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1553 @53.802s 1%: 0.048+0+0.50 ms clock, 0.77+0/0.25/28+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1554 @53.804s 1%: 0.016+28+0.40 ms clock, 0.25+0/0.10/28+6.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1555 @54.274s 1%: 0.004+0+0.53 ms clock, 0.065+0/0.10/28+8.5 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1556 @54.275s 1%: 0.043+0+0.51 ms clock, 0.69+0/0.10/28+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1557 @54.277s 1%: 0.013+29+0.30 ms clock, 0.22+0/0.15/29+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1558 @54.751s 1%: 0.003+0+0.50 ms clock, 0.056+0/0.15/29+8.0 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1559 @54.752s 1%: 0.025+0+0.73 ms clock, 0.41+0/0.15/29+11 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1560 @54.754s 1%: 0.015+29+0.37 ms clock, 0.24+0/0.045/29+6.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8204144 HeapInuse: 8454144 HeapObjects: 313 HeapIdle 8519680 HeapReleased 0 HeapSys 16973824
gc 1561 @55.232s 1%: 0.005+0+0.63 ms clock, 0.085+0/0.045/29+10 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
 1000000	       479 ns/op	       8 B/op	       0 allocs/op
gc 1562 @55.233s 1%: 0.004+0+0.51 ms clock, 0.074+0/0.045/29+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1563 @55.401s 1%: 0.020+3.4+0.29 ms clock, 0.32+0.16/11/14+4.7 ms cpu, 4->4->3 MB, 7 MB goal, 16 P
gc 1564 @55.502s 1%: 0.019+4.6+0.31 ms clock, 0.30+0.14/15/29+5.0 ms cpu, 5->6->6 MB, 7 MB goal, 16 P
gc 1565 @55.649s 1%: 0.021+9.1+0.29 ms clock, 0.34+0.17/31/53+4.7 ms cpu, 10->10->10 MB, 12 MB goal, 16 P
gc 1566 @55.940s 1%: 0.020+23+0.31 ms clock, 0.32+0.17/57/101+4.9 ms cpu, 19->19->19 MB, 21 MB goal, 16 P
gc 1567 @56.581s 1%: 0.020+32+0.31 ms clock, 0.32+0.16/115/211+5.0 ms cpu, 37->38->38 MB, 39 MB goal, 16 P
gc 1568 @57.696s 1%: 0.036+46+0.27 ms clock, 0.57+0.20/182/477+4.4 ms cpu, 74->74->74 MB, 76 MB goal, 16 P
HeapAlloc: 88200864 HeapInuse: 88743936 HeapObjects: 1500320 HeapIdle 581632 HeapReleased 0 HeapSys 89325568
gc 1569 @57.981s 1%: 0.002+0+4.2 ms clock, 0.040+0.20/182/477+68 ms cpu, 84->84->0 MB, 84 MB goal, 16 P (forced)
#BenchmarkSetInsert1E6-16                          	 1000000	      2751 ns/op	      88 B/op	       2 allocs/op
gc 1570 @57.986s 1%: 0.005+0+0.72 ms clock, 0.090+0.20/182/477+11 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1571 @57.988s 1%: 0.012+0.21+0.46 ms clock, 0.20+0/0.20/0.23+7.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1572 @58.029s 1%: 0.016+1.2+0.30 ms clock, 0.26+0.16/3.3/2.1+4.8 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1573 @58.149s 1%: 0.015+3.5+0.43 ms clock, 0.25+0.17/12/19+7.0 ms cpu, 12->13->12 MB, 17 MB goal, 16 P
gc 1574 @58.434s 1%: 0.022+10+0.32 ms clock, 0.35+0.18/39/88+5.2 ms cpu, 22->22->22 MB, 25 MB goal, 16 P
gc 1575 @59.080s 1%: 0.021+27+0.31 ms clock, 0.34+0.24/98/222+5.0 ms cpu, 41->42->42 MB, 45 MB goal, 16 P
gc 1576 @60.209s 1%: 0.015+45+0.32 ms clock, 0.24+0.12/176/396+5.2 ms cpu, 80->81->81 MB, 84 MB goal, 16 P
gc 1577 @62.369s 1%: 0.033+14+0.29 ms clock, 0.53+0/45/80+4.7 ms cpu, 159->159->24 MB, 163 MB goal, 16 P
HeapAlloc: 43710048 HeapInuse: 44072960 HeapObjects: 605529 HeapIdle 123895808 HeapReleased 0 HeapSys 167968768
gc 1578 @62.691s 1%: 0.003+0+3.0 ms clock, 0.061+0/45/80+48 ms cpu, 41->41->0 MB, 41 MB goal, 16 P (forced)
#BenchmarkSetErase1E6-16                           	 1000000	      2009 ns/op	      96 B/op	       2 allocs/op
gc 1579 @62.695s 1%: 0.017+0+1.5 ms clock, 0.27+0/45/80+24 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1580 @62.698s 1%: 0.011+0.49+0.34 ms clock, 0.19+0/0.57/0.34+5.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1581 @62.724s 1%: 0.009+4.0+0.45 ms clock, 0.14+0.23/3.8/0.50+7.2 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1582 @62.847s 1%: 0.015+4.4+0.43 ms clock, 0.24+0.19/12/15+6.9 ms cpu, 13->13->13 MB, 17 MB goal, 16 P
gc 1583 @63.043s 1%: 0.015+2.3+0.29 ms clock, 0.24+0.20/7.1/11+4.6 ms cpu, 22->22->9 MB, 26 MB goal, 16 P
gc 1584 @63.408s 1%: 0.021+5.8+0.34 ms clock, 0.33+0.079/20/33+5.4 ms cpu, 18->18->14 MB, 19 MB goal, 16 P
gc 1585 @63.788s 1%: 0.018+2.7+0.27 ms clock, 0.29+0.23/8.3/13+4.4 ms cpu, 27->27->10 MB, 28 MB goal, 16 P
gc 1586 @64.035s 1%: 0.016+4.2+0.27 ms clock, 0.27+0.15/13/1.8+4.3 ms cpu, 20->21->11 MB, 21 MB goal, 16 P
gc 1587 @64.497s 1%: 0.022+5.7+0.28 ms clock, 0.36+0.19/21/49+4.5 ms cpu, 22->22->15 MB, 23 MB goal, 16 P
gc 1588 @65.017s 1%: 0.021+3.7+0.36 ms clock, 0.34+0.28/12/30+5.8 ms cpu, 29->29->12 MB, 30 MB goal, 16 P
HeapAlloc: 25804448 HeapInuse: 26099712 HeapObjects: 300317 HeapIdle 141869056 HeapReleased 0 HeapSys 167968768
gc 1589 @65.408s 1%: 0.003+0+2.2 ms clock, 0.055+0.28/12/30+36 ms cpu, 24->24->0 MB, 24 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E6-16                  	 1000000	      2714 ns/op	      96 B/op	       2 allocs/op
gc 1590 @65.411s 1%: 0.015+0+1.0 ms clock, 0.24+0.28/12/30+17 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1591 @65.414s 1%: 0.017+0.40+0.34 ms clock, 0.27+0.053/0.34/0.46+5.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1592 @65.468s 1%: 0.017+1.4+0.28 ms clock, 0.28+0.23/3.9/3.8+4.6 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1593 @65.690s 1%: 0.019+4.3+0.31 ms clock, 0.31+0.24/15/22+5.0 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1594 @66.385s 1%: 0.023+5.1+0.30 ms clock, 0.37+0.22/13/24+4.8 ms cpu, 21->21->12 MB, 25 MB goal, 16 P
gc 1595 @67.186s 1%: 0.018+6.1+0.30 ms clock, 0.29+0/21/19+4.8 ms cpu, 22->22->13 MB, 24 MB goal, 16 P
gc 1596 @68.051s 1%: 0.017+2.2+0.29 ms clock, 0.27+0/5.6/5.4+4.7 ms cpu, 26->26->9 MB, 27 MB goal, 16 P
gc 1597 @68.563s 1%: 0.020+5.0+0.32 ms clock, 0.33+0/18/24+5.1 ms cpu, 18->18->13 MB, 19 MB goal, 16 P
HeapAlloc: 14696936 HeapInuse: 15032320 HeapObjects: 135362 HeapIdle 152936448 HeapReleased 0 HeapSys 167968768
gc 1598 @68.724s 1%: 0.023+0+1.6 ms clock, 0.37+0/18/24+26 ms cpu, 14->14->0 MB, 14 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E6-16          	 1000000	      3313 ns/op	      56 B/op	       1 allocs/op
gc 1599 @68.726s 1%: 0.007+0+1.3 ms clock, 0.12+0/18/24+22 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1600 @68.824s 1%: 0.019+3.6+0.29 ms clock, 0.31+0.26/10/11+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1601 @68.922s 1%: 0.020+5.6+0.27 ms clock, 0.33+0.18/19/35+4.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1602 @69.039s 1%: 0.022+11+0.36 ms clock, 0.36+0.18/38/60+5.7 ms cpu, 13->13->13 MB, 14 MB goal, 16 P
gc 1603 @69.271s 1%: 0.021+15+0.29 ms clock, 0.34+0.34/60/148+4.7 ms cpu, 25->26->25 MB, 26 MB goal, 16 P
gc 1604 @69.780s 1%: 0.022+49+0.33 ms clock, 0.36+0.28/187/96+5.3 ms cpu, 48->49->46 MB, 50 MB goal, 16 P
HeapAlloc: 83348112 HeapInuse: 83894272 HeapObjects: 1197019 HeapIdle 84074496 HeapReleased 0 HeapSys 167968768
gc 1605 @70.593s 1%: 0.017+0+6.2 ms clock, 0.28+0.28/187/96+100 ms cpu, 79->79->0 MB, 79 MB goal, 16 P (forced)
#BenchmarkIntSetInsert1E6-16                       	 1000000	      1871 ns/op	      88 B/op	       2 allocs/op
gc 1606 @70.599s 1%: 0.056+0+1.2 ms clock, 0.90+0.28/187/96+19 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1607 @70.603s 1%: 0.018+0.34+0.38 ms clock, 0.29+0.026/0.42/0.51+6.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1608 @70.629s 1%: 0.014+1.4+0.28 ms clock, 0.23+0.26/3.2/3.2+4.4 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1609 @70.735s 1%: 0.017+3.8+0.31 ms clock, 0.27+0.25/12/16+5.0 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1610 @70.895s 1%: 0.019+8.3+0.32 ms clock, 0.30+0.32/29/72+5.1 ms cpu, 21->21->20 MB, 24 MB goal, 16 P
gc 1611 @71.206s 1%: 0.020+19+0.29 ms clock, 0.33+0.25/71/148+4.7 ms cpu, 38->38->36 MB, 41 MB goal, 16 P
gc 1612 @71.932s 1%: 0.022+47+0.38 ms clock, 0.36+0.57/173/308+6.0 ms cpu, 70->71->68 MB, 73 MB goal, 16 P
gc 1613 @73.113s 1%: 0.019+21+0.31 ms clock, 0.31+0.42/71/164+5.0 ms cpu, 133->133->41 MB, 137 MB goal, 16 P
HeapAlloc: 82009440 HeapInuse: 82526208 HeapObjects: 1098965 HeapIdle 85442560 HeapReleased 0 HeapSys 167968768
gc 1614 @73.725s 1%: 0.003+0+3.5 ms clock, 0.057+0.42/71/164+56 ms cpu, 78->78->0 MB, 78 MB goal, 16 P (forced)
#BenchmarkIntSetErase1E6-16                        	 1000000	      1360 ns/op	      96 B/op	       2 allocs/op
gc 1615 @73.729s 1%: 0.022+0+1.1 ms clock, 0.35+0.42/71/164+19 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1616 @73.732s 1%: 0.013+0.30+0.42 ms clock, 0.21+0/0.40/0.44+6.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1617 @73.748s 1%: 0.012+0.94+0.29 ms clock, 0.19+0.17/2.5/4.1+4.7 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1618 @73.818s 1%: 0.016+3.6+0.31 ms clock, 0.26+0.17/11/21+4.9 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1619 @73.932s 1%: 0.016+2.3+0.29 ms clock, 0.26+0.19/7.4/8.5+4.6 ms cpu, 21->21->10 MB, 24 MB goal, 16 P
gc 1620 @74.155s 1%: 0.023+3.7+0.27 ms clock, 0.37+0.21/12/27+4.4 ms cpu, 19->19->13 MB, 20 MB goal, 16 P
gc 1621 @74.380s 1%: 0.020+2.3+0.31 ms clock, 0.33+0.24/6.3/9.7+5.0 ms cpu, 26->26->9 MB, 27 MB goal, 16 P
gc 1622 @74.604s 1%: 0.023+5.0+0.36 ms clock, 0.38+0.30/16/16+5.8 ms cpu, 18->18->13 MB, 19 MB goal, 16 P
gc 1623 @74.917s 1%: 0.021+4.8+0.31 ms clock, 0.34+0.18/17/32+4.9 ms cpu, 25->25->14 MB, 26 MB goal, 16 P
gc 1624 @75.206s 1%: 0.022+2.7+0.28 ms clock, 0.36+0.29/8.4/14+4.5 ms cpu, 27->27->11 MB, 28 MB goal, 16 P
gc 1625 @75.446s 1%: 0.024+2.7+0.30 ms clock, 0.38+0.28/8.7/12+4.9 ms cpu, 21->21->10 MB, 22 MB goal, 16 P
HeapAlloc: 13667856 HeapInuse: 13926400 HeapObjects: 81616 HeapIdle 154042368 HeapReleased 0 HeapSys 167968768
gc 1626 @75.498s 1%: 0.004+0+1.6 ms clock, 0.074+0.28/8.7/12+26 ms cpu, 13->13->0 MB, 13 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndErase1E6-16               	 1000000	      1769 ns/op	      96 B/op	       2 allocs/op
gc 1627 @75.500s 1%: 0.069+0+1.2 ms clock, 1.0+0.28/8.7/12+18 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1628 @75.503s 1%: 0.014+0.30+0.36 ms clock, 0.23+0.006/0.36/0.32+5.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1629 @75.541s 1%: 0.018+1.3+0.79 ms clock, 0.29+0.18/3.5/4.0+12 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1630 @75.696s 1%: 0.020+3.7+0.30 ms clock, 0.33+0.19/13/27+4.8 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1631 @76.147s 1%: 0.020+3.3+0.42 ms clock, 0.32+0.26/11/17+6.8 ms cpu, 21->21->11 MB, 24 MB goal, 16 P
gc 1632 @76.660s 1%: 0.019+5.1+0.31 ms clock, 0.30+0/15/18+4.9 ms cpu, 22->22->12 MB, 23 MB goal, 16 P
gc 1633 @77.269s 1%: 0.021+5.6+0.29 ms clock, 0.35+0.31/20/36+4.7 ms cpu, 24->24->15 MB, 25 MB goal, 16 P
HeapAlloc: 26636024 HeapInuse: 26968064 HeapObjects: 352101 HeapIdle 141000704 HeapReleased 0 HeapSys 167968768
gc 1634 @77.914s 1%: 0.005+0+2.6 ms clock, 0.084+0.31/20/36+42 ms cpu, 25->25->0 MB, 25 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndEraseWithPool1E6-16       	 1000000	      2415 ns/op	      56 B/op	       1 allocs/op
gc 1635 @77.917s 1%: 0.063+0+1.1 ms clock, 1.0+0.31/20/36+18 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1636 @77.948s 1%: 0.017+0.35+0.27 ms clock, 0.28+0.18/0.20/0.63+4.4 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 1637 @77.986s 1%: 0.018+0.46+0.39 ms clock, 0.29+0.40/0.025/0.71+6.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1638 @78.043s 1%: 0.021+0.49+0.32 ms clock, 0.34+0.42/0.047/1.0+5.1 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1639 @78.063s 1%: 0.021+1.1+0.39 ms clock, 0.34+0.93/0.29/1.3+6.3 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1640 @78.169s 1%: 0.022+0.57+0.61 ms clock, 0.35+0.41/0.059/0.88+9.8 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1641 @78.236s 1%: 0.023+1.8+0.37 ms clock, 0.37+1.7/0.15/2.5+5.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38249784 HeapInuse: 38617088 HeapObjects: 32041 HeapIdle 129351680 HeapReleased 0 HeapSys 167968768
gc 1642 @78.389s 1%: 0.005+0+1.6 ms clock, 0.090+1.7/0.15/2.5+26 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E6-16                   	gc 1643 @78.391s 1%: 0.004+0+1.1 ms clock, 0.079+1.7/0.15/2.5+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1644 @78.425s 1%: 0.021+0.45+0.37 ms clock, 0.33+0.20/0.20/0.80+5.9 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1645 @78.460s 1%: 0.019+0.70+0.37 ms clock, 0.31+0.59/0.023/1.3+5.9 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1646 @78.521s 1%: 0.023+0.49+0.40 ms clock, 0.38+0.42/0.28/0.85+6.5 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1647 @78.540s 1%: 0.021+0.99+0.36 ms clock, 0.34+0.85/0.28/1.5+5.8 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1648 @78.649s 1%: 0.022+0.83+0.46 ms clock, 0.35+0.62/0.26/1.3+7.3 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1649 @78.715s 1%: 0.065+2.0+0.42 ms clock, 1.0+1.8/0.27/2.5+6.8 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38263192 HeapInuse: 38633472 HeapObjects: 32183 HeapIdle 129335296 HeapReleased 0 HeapSys 167968768
gc 1650 @78.891s 1%: 0.002+0+1.7 ms clock, 0.044+1.8/0.27/2.5+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1651 @78.893s 1%: 0.048+0+1.3 ms clock, 0.77+1.8/0.27/2.5+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1652 @78.927s 1%: 0.018+0.39+1.1 ms clock, 0.28+0.24/0.19/0.97+18 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1653 @78.965s 1%: 0.020+0.54+0.29 ms clock, 0.32+0.40/0.11/1.0+4.7 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1654 @79.024s 1%: 0.020+0.58+0.41 ms clock, 0.33+0.51/0.48/0.81+6.6 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1655 @79.041s 1%: 0.021+0.98+0.34 ms clock, 0.34+0.92/0.25/1.3+5.5 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1656 @79.148s 1%: 0.021+0.59+0.31 ms clock, 0.34+0.43/0.45/0.88+5.0 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1657 @79.214s 1%: 0.022+2.0+0.42 ms clock, 0.35+1.9/0.34/2.6+6.7 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38233240 HeapInuse: 38600704 HeapObjects: 31871 HeapIdle 129368064 HeapReleased 0 HeapSys 167968768
gc 1658 @79.371s 1%: 0.022+0+1.8 ms clock, 0.35+1.9/0.34/2.6+30 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1659 @79.374s 1%: 0.045+0+1.3 ms clock, 0.72+1.9/0.34/2.6+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1660 @79.408s 1%: 0.021+0.52+0.38 ms clock, 0.33+0.20/0.20/0.96+6.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1661 @79.446s 1%: 0.021+0.64+0.39 ms clock, 0.34+0.45/0.13/1.3+6.3 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1662 @79.507s 1%: 0.022+0.49+0.37 ms clock, 0.36+0.41/0.37/1.1+5.9 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1663 @79.527s 1%: 0.023+1.3+0.32 ms clock, 0.37+1.1/0.28/2.1+5.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1664 @79.634s 1%: 0.020+0.47+0.44 ms clock, 0.32+0.42/0.28/0.97+7.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1665 @79.702s 1%: 0.021+1.7+0.33 ms clock, 0.33+1.6/0.13/2.2+5.3 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38258104 HeapInuse: 38625280 HeapObjects: 32130 HeapIdle 129343488 HeapReleased 0 HeapSys 167968768
gc 1666 @79.855s 1%: 0.003+0+1.4 ms clock, 0.057+1.6/0.13/2.2+23 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1667 @79.857s 1%: 0.025+0+1.1 ms clock, 0.40+1.6/0.13/2.2+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1668 @79.893s 1%: 0.022+0.51+0.32 ms clock, 0.35+0.23/0.27/0.87+5.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1669 @79.929s 1%: 0.021+0.54+0.39 ms clock, 0.35+0.42/0.19/0.95+6.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1670 @79.987s 1%: 0.019+0.99+1.1 ms clock, 0.31+0.73/0.19/1.3+18 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1671 @80.008s 1%: 0.029+0.90+0.38 ms clock, 0.47+0.81/0.26/1.2+6.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1672 @80.117s 1%: 0.026+0.73+0.39 ms clock, 0.42+0.65/0.14/1.1+6.3 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1673 @80.180s 1%: 0.022+1.9+0.44 ms clock, 0.36+0/0.18/2.1+7.1 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38262520 HeapInuse: 38625280 HeapObjects: 32176 HeapIdle 129343488 HeapReleased 0 HeapSys 167968768
gc 1674 @80.350s 1%: 0.005+0+1.9 ms clock, 0.083+0/0.18/2.1+31 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1675 @80.352s 1%: 0.045+0+1.4 ms clock, 0.73+0/0.18/2.1+22 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1676 @80.386s 1%: 0.019+0.36+0.31 ms clock, 0.30+0.23/0.28/0.85+5.1 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1677 @80.425s 1%: 0.020+0.52+0.34 ms clock, 0.33+0.41/0.14/0.94+5.5 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1678 @80.485s 1%: 0.019+0.68+0.41 ms clock, 0.31+0.57/0.32/1.1+6.6 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1679 @80.506s 1%: 0.023+1.0+0.38 ms clock, 0.38+0.88/0.24/1.2+6.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1680 @80.611s 1%: 0.019+0.62+0.36 ms clock, 0.31+0.42/0.18/0.99+5.7 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1681 @80.679s 1%: 0.024+1.7+0.36 ms clock, 0.39+1.6/0.20/2.4+5.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38253784 HeapInuse: 38617088 HeapObjects: 32085 HeapIdle 129351680 HeapReleased 0 HeapSys 167968768
gc 1682 @80.849s 1%: 0.007+0+1.8 ms clock, 0.11+1.6/0.20/2.4+29 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1683 @80.851s 1%: 0.032+0+1.8 ms clock, 0.52+1.6/0.20/2.4+30 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1684 @80.888s 1%: 0.021+0.40+0.33 ms clock, 0.34+0.20/0.23/0.68+5.4 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1685 @80.923s 1%: 0.023+0.59+0.36 ms clock, 0.37+0.51/0.056/1.1+5.8 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1686 @80.984s 1%: 0.019+0.48+0.34 ms clock, 0.31+0.40/0.23/0.84+5.5 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1687 @81.004s 1%: 0.019+1.0+0.34 ms clock, 0.30+0.83/0.18/1.8+5.5 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1688 @81.117s 1%: 0.026+0.72+4.1 ms clock, 0.43+0.59/0.092/1.2+65 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1689 @81.213s 1%: 0.026+1.9+0.31 ms clock, 0.42+1.8/0.18/2.6+4.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38250712 HeapInuse: 38617088 HeapObjects: 32053 HeapIdle 129351680 HeapReleased 0 HeapSys 167968768
gc 1690 @81.380s 1%: 0.016+0+1.4 ms clock, 0.26+1.8/0.18/2.6+23 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1691 @81.382s 1%: 0.004+0+1.1 ms clock, 0.074+1.8/0.18/2.6+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1692 @81.416s 1%: 0.018+0.44+0.34 ms clock, 0.30+0.20/0.16/0.86+5.4 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1693 @81.457s 1%: 0.021+0.59+0.34 ms clock, 0.34+0.42/0.11/1.0+5.5 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1694 @81.527s 1%: 0.021+0.63+0.39 ms clock, 0.33+0.55/0.11/1.2+6.2 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1695 @81.548s 1%: 0.024+1.3+0.33 ms clock, 0.39+1.1/0.085/2.1+5.3 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1696 @81.658s 1%: 0.021+0.49+0.29 ms clock, 0.34+0.43/0.15/1.1+4.7 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1697 @81.722s 1%: 0.019+1.7+0.36 ms clock, 0.31+1.6/0.25/1.9+5.8 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38233432 HeapInuse: 38600704 HeapObjects: 31873 HeapIdle 129368064 HeapReleased 0 HeapSys 167968768
gc 1698 @81.894s 1%: 0.006+0+1.8 ms clock, 0.099+1.6/0.25/1.9+29 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1699 @81.896s 1%: 0.078+0+1.1 ms clock, 1.2+1.6/0.25/1.9+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1700 @81.929s 1%: 0.015+0.48+0.42 ms clock, 0.24+0.20/0.20/0.64+6.7 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1701 @81.968s 1%: 0.029+0.59+0.34 ms clock, 0.47+0.48/0.11/0.99+5.5 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1702 @82.027s 1%: 0.018+0.50+0.38 ms clock, 0.29+0.46/0.39/0.94+6.1 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1703 @82.049s 1%: 0.025+1.6+0.38 ms clock, 0.40+1.2/0.25/1.9+6.2 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1704 @82.161s 1%: 0.025+1.9+0.38 ms clock, 0.40+0/0.37/2.2+6.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1705 @82.226s 1%: 0.022+1.8+0.36 ms clock, 0.36+1.7/0.25/2.5+5.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38239576 HeapInuse: 38600704 HeapObjects: 31937 HeapIdle 129368064 HeapReleased 0 HeapSys 167968768
gc 1706 @82.381s 1%: 0.003+0+1.5 ms clock, 0.062+1.7/0.25/2.5+24 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1707 @82.383s 1%: 0.030+0+1.0 ms clock, 0.48+1.7/0.25/2.5+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1708 @82.417s 1%: 0.022+0.36+0.34 ms clock, 0.36+0.23/0.30/0.85+5.4 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1709 @82.454s 1%: 0.025+0.90+0.32 ms clock, 0.40+0/0.51/1.1+5.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1710 @82.516s 1%: 0.020+2.4+0.35 ms clock, 0.32+0/0.48/2.5+5.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1711 @82.535s 1%: 0.022+0.90+0.35 ms clock, 0.35+0.82/0.25/1.0+5.7 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1712 @82.645s 1%: 0.019+0.50+0.39 ms clock, 0.31+0.42/0.16/1.0+6.3 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1713 @82.714s 1%: 0.023+1.9+0.37 ms clock, 0.37+1.6/0.16/2.3+5.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38252632 HeapInuse: 38617088 HeapObjects: 32073 HeapIdle 129351680 HeapReleased 0 HeapSys 167968768
gc 1714 @82.888s 1%: 0.005+0+1.9 ms clock, 0.082+1.6/0.16/2.3+31 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1715 @82.890s 1%: 0.032+0+1.3 ms clock, 0.51+1.6/0.16/2.3+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1716 @82.924s 1%: 0.015+0.44+0.57 ms clock, 0.24+0.20/0.22/0.68+9.1 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1717 @82.963s 1%: 0.021+0.93+0.34 ms clock, 0.34+0.74/0.031/1.5+5.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1718 @83.025s 1%: 0.019+0.57+0.40 ms clock, 0.31+0.52/0.39/0.98+6.4 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1719 @83.046s 1%: 0.022+1.0+0.37 ms clock, 0.35+0.89/0.10/2.1+5.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1720 @83.154s 1%: 0.020+2.6+0.32 ms clock, 0.32+0/2.4/0.86+5.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1721 @83.219s 1%: 0.021+1.9+0.39 ms clock, 0.33+1.8/0.30/2.3+6.2 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38272888 HeapInuse: 38633472 HeapObjects: 32284 HeapIdle 129335296 HeapReleased 0 HeapSys 167968768
gc 1722 @83.377s 1%: 0.004+0+1.7 ms clock, 0.075+1.8/0.30/2.3+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1723 @83.379s 1%: 0.055+0+1.3 ms clock, 0.88+1.8/0.30/2.3+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1724 @83.415s 1%: 0.024+0.39+0.39 ms clock, 0.39+0.21/0.14/0.89+6.3 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1725 @83.453s 1%: 0.020+0.72+0.32 ms clock, 0.32+0.51/0.19/1.4+5.1 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1726 @83.519s 1%: 0.019+2.3+0.40 ms clock, 0.31+0/0.57/2.3+6.5 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1727 @83.541s 1%: 0.023+1.0+0.43 ms clock, 0.37+1.0/0.28/1.5+7.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1728 @83.655s 1%: 0.021+0.63+0.36 ms clock, 0.34+0.52/0.28/1.2+5.8 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1729 @83.725s 1%: 0.024+2.5+0.40 ms clock, 0.38+2.4/0.17/3.1+6.5 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38246568 HeapInuse: 38608896 HeapObjects: 32010 HeapIdle 129359872 HeapReleased 0 HeapSys 167968768
gc 1730 @83.900s 1%: 0.016+0+1.6 ms clock, 0.26+2.4/0.17/3.1+26 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1731 @83.902s 1%: 0.047+0+1.6 ms clock, 0.75+2.4/0.17/3.1+26 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1732 @83.936s 1%: 0.019+0.46+0.32 ms clock, 0.30+0.30/0.079/1.1+5.1 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1733 @83.975s 1%: 0.019+0.49+0.32 ms clock, 0.30+0.41/0.21/0.82+5.1 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1734 @84.040s 1%: 0.022+0.66+0.35 ms clock, 0.35+0.50/0.11/0.93+5.6 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1735 @84.060s 1%: 0.024+1.3+0.35 ms clock, 0.39+1.2/0.14/1.9+5.7 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1736 @84.168s 1%: 0.022+0.56+0.38 ms clock, 0.35+0.51/0.19/1.1+6.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1737 @84.234s 1%: 0.019+2.0+0.38 ms clock, 0.31+1.6/0.19/2.0+6.1 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38237544 HeapInuse: 38600704 HeapObjects: 31916 HeapIdle 129368064 HeapReleased 0 HeapSys 167968768
gc 1738 @84.389s 1%: 0.004+0+1.8 ms clock, 0.076+1.6/0.19/2.0+29 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1739 @84.391s 1%: 0.031+0+1.2 ms clock, 0.49+1.6/0.19/2.0+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1740 @84.424s 1%: 0.018+0.34+0.37 ms clock, 0.29+0.21/0.19/0.83+6.0 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1741 @84.461s 1%: 0.021+0.64+0.39 ms clock, 0.33+0.51/0.21/0.99+6.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1742 @84.521s 1%: 0.019+0.60+0.38 ms clock, 0.31+0.41/0.19/0.74+6.1 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1743 @84.543s 1%: 0.026+1.1+0.30 ms clock, 0.43+1.0/0.31/1.5+4.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1744 @84.652s 1%: 0.024+0.58+0.39 ms clock, 0.38+0.52/0.58/0.81+6.2 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1745 @84.719s 1%: 0.023+2.8+0.34 ms clock, 0.37+2.5/0.002/3.0+5.4 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38230248 HeapInuse: 38592512 HeapObjects: 31840 HeapIdle 129376256 HeapReleased 0 HeapSys 167968768
gc 1746 @84.878s 1%: 0.004+0+1.8 ms clock, 0.079+2.5/0.002/3.0+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1747 @84.880s 1%: 0.006+0+1.2 ms clock, 0.10+2.5/0.002/3.0+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1748 @84.914s 1%: 0.021+0.39+0.41 ms clock, 0.34+0.22/0.36/0.70+6.7 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1749 @84.950s 1%: 0.019+0.88+0.31 ms clock, 0.31+0.57/0.24/0.99+5.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1750 @85.012s 1%: 0.021+2.7+0.41 ms clock, 0.33+0/0.66/2.7+6.5 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1751 @85.034s 1%: 0.021+0.98+0.43 ms clock, 0.34+0.82/0.10/1.5+6.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1752 @85.143s 1%: 0.022+0.55+0.46 ms clock, 0.36+0.48/0.15/1.1+7.3 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1753 @85.209s 1%: 0.021+1.9+0.40 ms clock, 0.34+1.6/0.16/2.2+6.4 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38244072 HeapInuse: 38617088 HeapObjects: 31984 HeapIdle 129351680 HeapReleased 0 HeapSys 167968768
gc 1754 @85.363s 1%: 0.002+0+1.5 ms clock, 0.042+1.6/0.16/2.2+24 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
 1000000	       483 ns/op	      53 B/op	       0 allocs/op
gc 1755 @85.365s 1%: 0.052+0+1.1 ms clock, 0.83+1.6/0.16/2.2+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1756 @85.368s 1%: 0.015+0.27+0.32 ms clock, 0.25+0/0.32/0.54+5.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1757 @85.378s 1%: 0.013+0.25+0.30 ms clock, 0.22+0.042/0.15/0.42+4.9 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1758 @85.405s 1%: 0.019+0.42+0.59 ms clock, 0.31+0.20/0.12/0.70+9.4 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1759 @85.503s 1%: 0.022+0.56+0.33 ms clock, 0.36+0.48/0.13/0.91+5.4 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1760 @85.520s 1%: 0.019+0.90+0.36 ms clock, 0.30+0.83/0.75/0.73+5.8 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1761 @85.685s 1%: 0.021+2.1+0.42 ms clock, 0.34+2.1/0.13/3.0+6.7 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46248152 HeapInuse: 46612480 HeapObjects: 31988 HeapIdle 121356288 HeapReleased 0 HeapSys 167968768
gc 1762 @86.055s 1%: 0.003+0+1.6 ms clock, 0.054+2.1/0.13/3.0+27 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E6-16                    	gc 1763 @86.057s 1%: 0.008+0+1.5 ms clock, 0.13+2.1/0.13/3.0+25 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1764 @86.060s 1%: 0.016+0.36+0.36 ms clock, 0.25+0/0.54/0.33+5.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1765 @86.071s 1%: 0.015+0.34+0.78 ms clock, 0.24+0.042/0.20/0.67+12 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1766 @86.098s 1%: 0.020+0.38+0.37 ms clock, 0.32+0.25/0.11/1.0+5.9 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1767 @86.195s 1%: 0.022+2.5+0.54 ms clock, 0.35+0.56/0.26/2.5+8.7 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1768 @86.210s 1%: 0.019+1.1+0.47 ms clock, 0.31+0.83/0.26/1.0+7.6 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1769 @86.374s 1%: 0.022+1.8+0.34 ms clock, 0.36+1.7/0.33/2.4+5.5 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46252424 HeapInuse: 46612480 HeapObjects: 32035 HeapIdle 121356288 HeapReleased 0 HeapSys 167968768
gc 1770 @86.699s 1%: 0.003+0+1.7 ms clock, 0.056+1.7/0.33/2.4+27 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1771 @86.701s 1%: 0.041+0+1.0 ms clock, 0.67+1.7/0.33/2.4+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1772 @86.704s 1%: 0.018+0.32+0.29 ms clock, 0.29+0/0.41/0.41+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1773 @86.715s 1%: 0.015+0.35+0.27 ms clock, 0.24+0.040/0.28/0.61+4.3 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1774 @86.742s 1%: 0.017+0.45+0.44 ms clock, 0.28+0.24/0.23/0.71+7.1 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1775 @86.842s 1%: 0.023+0.71+0.38 ms clock, 0.36+0.66/0.14/1.4+6.1 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1776 @86.855s 1%: 0.021+1.3+0.26 ms clock, 0.34+0.40/1.0/1.0+4.3 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1777 @87.022s 1%: 0.024+1.6+0.42 ms clock, 0.38+1.6/0.007/2.3+6.8 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46227544 HeapInuse: 46587904 HeapObjects: 31776 HeapIdle 121380864 HeapReleased 0 HeapSys 167968768
gc 1778 @87.388s 1%: 0.003+0+1.9 ms clock, 0.050+1.6/0.007/2.3+31 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1779 @87.390s 1%: 0.013+0+1.4 ms clock, 0.22+1.6/0.007/2.3+23 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1780 @87.394s 1%: 0.020+0.32+0.39 ms clock, 0.32+0/0.46/0.66+6.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1781 @87.403s 1%: 0.013+0.32+0.29 ms clock, 0.21+0.040/0.16/0.59+4.7 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1782 @87.432s 1%: 0.021+0.47+0.43 ms clock, 0.34+0.29/0.083/0.98+6.9 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1783 @87.532s 1%: 0.021+1.1+0.43 ms clock, 0.35+0.84/0.42/1.3+7.0 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1784 @87.564s 1%: 0.018+1.1+0.32 ms clock, 0.30+0.024/1.0/0.93+5.2 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1785 @87.720s 1%: 0.025+1.8+0.39 ms clock, 0.41+1.7/0.21/2.4+6.3 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46252888 HeapInuse: 46612480 HeapObjects: 32040 HeapIdle 121356288 HeapReleased 0 HeapSys 167968768
gc 1786 @88.040s 1%: 0.003+0+1.7 ms clock, 0.053+1.7/0.21/2.4+27 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1787 @88.042s 1%: 0.007+0+1.2 ms clock, 0.11+1.7/0.21/2.4+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1788 @88.045s 1%: 0.025+0.40+0.28 ms clock, 0.40+0/0.37/0.40+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1789 @88.056s 1%: 0.015+0.57+0.33 ms clock, 0.24+0.042/0.029/0.67+5.3 ms cpu, 8->9->8 MB, 15 MB goal, 16 P
gc 1790 @88.083s 1%: 0.017+0.63+0.45 ms clock, 0.27+0.31/0.082/1.0+7.2 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1791 @88.188s 1%: 0.021+4.0+0.35 ms clock, 0.34+0.63/0/4.0+5.6 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1792 @88.290s 1%: 0.022+0.45+0.36 ms clock, 0.36+0.31/0.24/0.77+5.8 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1793 @88.361s 1%: 0.022+1.7+0.40 ms clock, 0.36+1.7/0.20/2.3+6.4 ms cpu, 44->44->43 MB, 45 MB goal, 16 P
HeapAlloc: 46269672 HeapInuse: 46637056 HeapObjects: 32215 HeapIdle 121331712 HeapReleased 0 HeapSys 167968768
gc 1794 @88.680s 1%: 0.003+0+1.5 ms clock, 0.060+1.7/0.20/2.3+24 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1795 @88.682s 1%: 0.064+0+1.2 ms clock, 1.0+1.7/0.20/2.3+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1796 @88.685s 1%: 0.016+0.26+0.29 ms clock, 0.26+0/0.45/0.38+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1797 @88.694s 1%: 0.015+0.31+0.29 ms clock, 0.24+0.064/0.17/0.58+4.6 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1798 @88.723s 1%: 0.021+0.44+0.41 ms clock, 0.34+0.20/0.19/0.80+6.5 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1799 @88.826s 1%: 0.019+3.6+0.32 ms clock, 0.31+0/0.69/3.6+5.2 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1800 @88.844s 1%: 0.022+1.2+0.34 ms clock, 0.36+1.1/0.31/1.8+5.5 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1801 @89.032s 1%: 0.022+1.7+0.37 ms clock, 0.36+1.7/0.30/2.6+6.0 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46255160 HeapInuse: 46628864 HeapObjects: 32064 HeapIdle 121339904 HeapReleased 0 HeapSys 167968768
gc 1802 @89.393s 1%: 0.006+0+1.8 ms clock, 0.10+1.7/0.30/2.6+29 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1803 @89.395s 1%: 0.063+0+1.1 ms clock, 1.0+1.7/0.30/2.6+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1804 @89.398s 1%: 0.017+0.33+0.31 ms clock, 0.28+0/0.43/0.51+5.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1805 @89.407s 1%: 0.015+0.46+0.31 ms clock, 0.24+0.35/0.27/0.98+5.0 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1806 @89.435s 1%: 0.045+0.36+0.32 ms clock, 0.72+0.20/0.27/0.77+5.2 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1807 @89.534s 1%: 0.023+3.0+0.45 ms clock, 0.37+0/0.82/2.9+7.2 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1808 @89.549s 1%: 0.022+1.0+0.39 ms clock, 0.35+0.83/0.065/1.5+6.3 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1809 @89.727s 1%: 0.020+1.9+0.45 ms clock, 0.32+1.7/0.22/2.5+7.2 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46244872 HeapInuse: 46604288 HeapObjects: 31957 HeapIdle 121364480 HeapReleased 0 HeapSys 167968768
gc 1810 @90.043s 1%: 0.005+0+1.7 ms clock, 0.080+1.7/0.22/2.5+27 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1811 @90.045s 1%: 0.032+0+1.0 ms clock, 0.51+1.7/0.22/2.5+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1812 @90.048s 1%: 0.016+0.31+0.28 ms clock, 0.26+0/0.48/0.26+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1813 @90.058s 1%: 0.016+0.42+0.34 ms clock, 0.25+0.039/0.15/0.87+5.4 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1814 @90.086s 1%: 0.016+0.39+0.32 ms clock, 0.25+0.19/0.23/0.62+5.2 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1815 @90.184s 1%: 0.020+0.66+0.47 ms clock, 0.32+0.60/0.027/1.4+7.5 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1816 @90.213s 1%: 0.025+1.1+0.37 ms clock, 0.41+0.032/1.1/0.84+6.0 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1817 @90.361s 1%: 0.023+1.8+0.37 ms clock, 0.37+1.7/0.36/2.2+6.0 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46250328 HeapInuse: 46612480 HeapObjects: 32014 HeapIdle 121356288 HeapReleased 0 HeapSys 167968768
gc 1818 @90.725s 1%: 0.003+0+1.6 ms clock, 0.051+1.7/0.36/2.2+25 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1819 @90.727s 1%: 0.076+0+1.1 ms clock, 1.2+1.7/0.36/2.2+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1820 @90.730s 1%: 0.014+0.39+0.33 ms clock, 0.23+0/0.19/0.53+5.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1821 @90.740s 1%: 0.012+0.38+0.33 ms clock, 0.19+0.041/0.19/0.47+5.3 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1822 @90.766s 1%: 0.015+0.78+0.35 ms clock, 0.24+0.27/0.26/0.64+5.6 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1823 @90.865s 1%: 0.019+0.79+0.34 ms clock, 0.30+0.60/0.18/1.5+5.4 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1824 @90.938s 1%: 0.019+1.2+0.39 ms clock, 0.31+0.082/1.0/1.0+6.2 ms cpu, 26->26->25 MB, 28 MB goal, 16 P
gc 1825 @91.052s 1%: 0.022+1.8+0.38 ms clock, 0.35+1.8/0.038/3.0+6.2 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46245144 HeapInuse: 46604288 HeapObjects: 31960 HeapIdle 121364480 HeapReleased 0 HeapSys 167968768
gc 1826 @91.412s 1%: 0.005+0+1.8 ms clock, 0.083+1.8/0.038/3.0+29 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1827 @91.414s 1%: 0.091+0+1.0 ms clock, 1.4+1.8/0.038/3.0+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1828 @91.417s 1%: 0.015+0.31+0.42 ms clock, 0.24+0/0.31/0.57+6.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1829 @91.427s 1%: 0.015+1.6+0.33 ms clock, 0.24+0.043/0.16/0.53+5.3 ms cpu, 8->9->8 MB, 15 MB goal, 16 P
gc 1830 @91.453s 1%: 0.018+0.42+0.41 ms clock, 0.29+0.28/0.22/0.77+6.6 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1831 @91.559s 1%: 0.022+0.93+0.40 ms clock, 0.35+0.73/0.18/1.0+6.4 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1832 @91.657s 1%: 0.024+0.44+0.35 ms clock, 0.38+0.34/0.15/0.89+5.7 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1833 @91.733s 1%: 0.027+1.8+0.38 ms clock, 0.43+1.6/0.21/2.0+6.1 ms cpu, 44->44->43 MB, 45 MB goal, 16 P
HeapAlloc: 46257896 HeapInuse: 46628864 HeapObjects: 32093 HeapIdle 121339904 HeapReleased 0 HeapSys 167968768
gc 1834 @92.101s 1%: 0.003+0+1.8 ms clock, 0.053+1.6/0.21/2.0+30 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1835 @92.103s 1%: 0.044+0+1.3 ms clock, 0.71+1.6/0.21/2.0+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1836 @92.106s 1%: 0.014+0.29+0.28 ms clock, 0.22+0/0.41/0.34+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1837 @92.116s 1%: 0.011+0.29+0.71 ms clock, 0.18+0.042/0.19/0.47+11 ms cpu, 8->9->8 MB, 15 MB goal, 16 P
gc 1838 @92.142s 1%: 0.017+0.43+0.39 ms clock, 0.27+0.21/0.16/0.72+6.3 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1839 @92.243s 1%: 0.022+4.7+0.34 ms clock, 0.36+0/0.80/4.6+5.5 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1840 @92.349s 1%: 0.022+0.60+0.39 ms clock, 0.36+0.42/0.27/0.92+6.3 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1841 @92.428s 1%: 0.022+2.4+0.35 ms clock, 0.36+2.3/0.15/3.1+5.6 ms cpu, 44->44->43 MB, 45 MB goal, 16 P
HeapAlloc: 46256728 HeapInuse: 46620672 HeapObjects: 32081 HeapIdle 121348096 HeapReleased 0 HeapSys 167968768
gc 1842 @92.790s 1%: 0.002+0+1.9 ms clock, 0.043+2.3/0.15/3.1+31 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1843 @92.792s 1%: 0.039+0+1.3 ms clock, 0.63+2.3/0.15/3.1+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1844 @92.795s 1%: 0.015+0.41+0.32 ms clock, 0.24+0/0.46/0.25+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1845 @92.805s 1%: 0.016+0.45+0.33 ms clock, 0.26+0.043/0.23/0.57+5.3 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1846 @92.832s 1%: 0.018+0.41+0.34 ms clock, 0.30+0.29/0.087/1.1+5.4 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1847 @92.937s 1%: 0.021+3.4+0.35 ms clock, 0.35+0/0.70/3.5+5.6 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1848 @92.955s 1%: 0.022+1.6+0.39 ms clock, 0.36+1.5/0.23/2.3+6.3 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1849 @93.129s 1%: 0.024+2.5+0.35 ms clock, 0.38+2.3/0.25/3.0+5.7 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46236840 HeapInuse: 46596096 HeapObjects: 31874 HeapIdle 121372672 HeapReleased 0 HeapSys 167968768
gc 1850 @93.445s 1%: 0.005+0+1.6 ms clock, 0.081+2.3/0.25/3.0+26 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1851 @93.447s 1%: 0.073+0+1.2 ms clock, 1.1+2.3/0.25/3.0+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1852 @93.451s 1%: 0.016+0.33+0.30 ms clock, 0.26+0/0.40/0.55+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1853 @93.461s 1%: 0.014+0.33+0.29 ms clock, 0.23+0.053/0.14/0.59+4.7 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1854 @93.488s 1%: 0.020+0.40+0.29 ms clock, 0.32+0.20/0.19/0.68+4.7 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1855 @93.590s 1%: 0.021+0.65+0.39 ms clock, 0.34+0.58/0.19/0.93+6.3 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1856 @93.630s 1%: 0.022+1.2+0.47 ms clock, 0.35+0.041/1.0/1.0+7.5 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1857 @93.776s 1%: 0.022+1.7+0.35 ms clock, 0.36+1.6/0.23/2.3+5.6 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46250664 HeapInuse: 46612480 HeapObjects: 32018 HeapIdle 121356288 HeapReleased 0 HeapSys 167968768
gc 1858 @94.107s 1%: 0.006+0+1.6 ms clock, 0.10+1.6/0.23/2.3+27 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
 1000000	       173 ns/op	       8 B/op	       0 allocs/op
gc 1859 @94.109s 1%: 0.043+0+1.2 ms clock, 0.70+1.6/0.23/2.3+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1860 @94.129s 1%: 0.019+0.43+0.37 ms clock, 0.31+0.10/0.59/0.54+5.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4671880 HeapInuse: 5005312 HeapObjects: 11821 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1861 @94.285s 1%: 0.005+0+2.0 ms clock, 0.082+0.10/0.59/0.54+33 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E6-16           	gc 1862 @94.288s 1%: 0.011+0+1.3 ms clock, 0.18+0.10/0.59/0.54+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1863 @94.306s 1%: 0.015+0.36+0.40 ms clock, 0.25+0.13/0.41/0.62+6.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4669560 HeapInuse: 5005312 HeapObjects: 11799 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1864 @94.463s 1%: 0.003+0+1.7 ms clock, 0.056+0.13/0.41/0.62+28 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1865 @94.465s 1%: 0.022+0+1.3 ms clock, 0.35+0.13/0.41/0.62+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1866 @94.483s 1%: 0.015+0.41+0.40 ms clock, 0.24+0.10/0.23/0.88+6.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4673976 HeapInuse: 5013504 HeapObjects: 11845 HeapIdle 162955264 HeapReleased 0 HeapSys 167968768
gc 1867 @94.641s 1%: 0.005+0+1.8 ms clock, 0.092+0.10/0.23/0.88+30 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1868 @94.643s 1%: 0.007+0+1.5 ms clock, 0.11+0.10/0.23/0.88+24 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1869 @94.661s 1%: 0.016+0.37+0.31 ms clock, 0.26+0.12/0.22/0.69+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4670232 HeapInuse: 5005312 HeapObjects: 11806 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1870 @94.816s 1%: 0.005+0+1.8 ms clock, 0.080+0.12/0.22/0.69+29 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1871 @94.818s 1%: 0.005+0+1.5 ms clock, 0.080+0.12/0.22/0.69+24 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1872 @94.837s 1%: 0.018+0.42+0.44 ms clock, 0.29+0.10/0.20/0.63+7.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4664952 HeapInuse: 5005312 HeapObjects: 11751 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1873 @94.993s 1%: 0.005+0+1.9 ms clock, 0.083+0.10/0.20/0.63+31 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1874 @94.996s 1%: 0.056+0+1.4 ms clock, 0.89+0.10/0.20/0.63+23 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1875 @95.014s 1%: 0.017+0.40+0.59 ms clock, 0.28+0.10/0.27/0.56+9.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4659768 HeapInuse: 5013504 HeapObjects: 11697 HeapIdle 162955264 HeapReleased 0 HeapSys 167968768
gc 1876 @95.171s 1%: 0.004+0+1.9 ms clock, 0.077+0.10/0.27/0.56+30 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1877 @95.173s 1%: 0.044+0+1.3 ms clock, 0.70+0.10/0.27/0.56+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1878 @95.191s 1%: 0.022+0.57+0.32 ms clock, 0.35+0.10/0.20/0.80+5.2 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4661880 HeapInuse: 4997120 HeapObjects: 11719 HeapIdle 162971648 HeapReleased 0 HeapSys 167968768
gc 1879 @95.348s 1%: 0.002+0+1.9 ms clock, 0.046+0.10/0.20/0.80+30 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1880 @95.351s 1%: 0.048+0+1.1 ms clock, 0.77+0.10/0.20/0.80+19 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1881 @95.368s 1%: 0.016+0.44+0.41 ms clock, 0.25+0.10/0.22/0.64+6.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4672056 HeapInuse: 5013504 HeapObjects: 11825 HeapIdle 162955264 HeapReleased 0 HeapSys 167968768
gc 1882 @95.523s 1%: 0.006+0+1.7 ms clock, 0.10+0.10/0.22/0.64+28 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1883 @95.525s 1%: 0.040+0+1.1 ms clock, 0.64+0.10/0.22/0.64+18 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1884 @95.543s 1%: 0.016+0.43+0.34 ms clock, 0.26+0.096/0.21/0.83+5.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4650488 HeapInuse: 4988928 HeapObjects: 11643 HeapIdle 162979840 HeapReleased 0 HeapSys 167968768
gc 1885 @95.698s 1%: 0.004+0+1.9 ms clock, 0.077+0.096/0.21/0.83+30 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1886 @95.700s 1%: 0.045+0+1.4 ms clock, 0.72+0.096/0.21/0.83+23 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1887 @95.718s 1%: 0.017+0.50+0.42 ms clock, 0.27+0.10/0.072/0.73+6.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4671384 HeapInuse: 5005312 HeapObjects: 11818 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1888 @95.877s 1%: 0.003+0+1.9 ms clock, 0.060+0.10/0.072/0.73+31 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1889 @95.879s 1%: 0.004+0+1.3 ms clock, 0.078+0.10/0.072/0.73+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1890 @95.898s 1%: 0.016+0.31+0.31 ms clock, 0.25+0.12/0.22/0.72+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4670520 HeapInuse: 5021696 HeapObjects: 11809 HeapIdle 162947072 HeapReleased 0 HeapSys 167968768
gc 1891 @96.052s 1%: 0.003+0+1.6 ms clock, 0.054+0.12/0.22/0.72+26 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1892 @96.054s 1%: 0.049+0+1.1 ms clock, 0.79+0.12/0.22/0.72+18 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1893 @96.072s 1%: 0.020+0.47+0.34 ms clock, 0.33+0.10/0.19/0.86+5.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4669080 HeapInuse: 5005312 HeapObjects: 11794 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1894 @96.228s 1%: 0.006+0+2.0 ms clock, 0.11+0.10/0.19/0.86+32 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1895 @96.231s 1%: 0.018+0+1.4 ms clock, 0.28+0.10/0.19/0.86+23 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1896 @96.249s 1%: 0.016+0.47+0.34 ms clock, 0.25+0.14/0.35/0.49+5.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4671960 HeapInuse: 5005312 HeapObjects: 11824 HeapIdle 162963456 HeapReleased 0 HeapSys 167968768
gc 1897 @96.407s 1%: 0.021+0+1.7 ms clock, 0.33+0.14/0.35/0.49+27 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
 1000000	       176 ns/op	       5 B/op	       0 allocs/op
gc 1898 @96.409s 1%: 0.006+0+1.3 ms clock, 0.099+0.14/0.35/0.49+22 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1899 @96.415s 1%: 0.023+0.47+0.32 ms clock, 0.37+0/0.066/0.89+5.2 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073120 HeapInuse: 24322048 HeapObjects: 325 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1900 @96.598s 1%: 0.004+0+1.4 ms clock, 0.079+0/0.066/0.89+23 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E6-16    	gc 1901 @96.600s 1%: 0.010+0+1.2 ms clock, 0.16+0/0.066/0.89+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1902 @96.606s 1%: 0.017+0.41+0.33 ms clock, 0.27+0/0.16/0.85+5.4 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1903 @96.789s 1%: 0.005+0+1.6 ms clock, 0.085+0/0.16/0.85+25 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1904 @96.791s 1%: 0.013+0+1.0 ms clock, 0.21+0/0.16/0.85+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1905 @96.797s 1%: 0.019+0.55+0.41 ms clock, 0.31+0/0.55/0.39+6.6 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1906 @96.987s 1%: 0.003+0+1.1 ms clock, 0.057+0/0.55/0.39+18 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1907 @96.989s 1%: 0.043+0+1.2 ms clock, 0.69+0/0.55/0.39+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1908 @96.995s 1%: 0.021+0.42+0.49 ms clock, 0.34+0/0.29/0.89+7.9 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1909 @97.179s 1%: 0.004+0+1.2 ms clock, 0.075+0/0.29/0.89+20 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1910 @97.181s 1%: 0.040+0+1.1 ms clock, 0.64+0/0.29/0.89+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1911 @97.187s 1%: 0.021+0.46+0.41 ms clock, 0.34+0/0.39/0.66+6.7 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1912 @97.371s 1%: 0.003+0+1.5 ms clock, 0.059+0/0.39/0.66+25 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1913 @97.373s 1%: 0.025+0+1.5 ms clock, 0.40+0/0.39/0.66+24 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1914 @97.380s 1%: 0.022+0.44+0.40 ms clock, 0.35+0/0.16/0.73+6.4 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1915 @97.565s 1%: 0.003+0+1.2 ms clock, 0.061+0/0.16/0.73+19 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1916 @97.566s 1%: 0.066+0+1.0 ms clock, 1.0+0/0.16/0.73+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1917 @97.572s 1%: 0.020+0.53+0.39 ms clock, 0.33+0/0.029/1.0+6.2 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1918 @97.760s 1%: 0.005+0+1.4 ms clock, 0.082+0/0.029/1.0+23 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1919 @97.761s 1%: 0.070+0+1.1 ms clock, 1.1+0/0.029/1.0+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1920 @97.767s 1%: 0.018+0.38+0.37 ms clock, 0.29+0/0.41/0.61+5.9 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1921 @97.958s 1%: 0.003+0+1.3 ms clock, 0.055+0/0.41/0.61+20 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1922 @97.959s 1%: 0.046+0+1.2 ms clock, 0.74+0/0.41/0.61+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1923 @97.966s 1%: 0.021+0.62+0.46 ms clock, 0.33+0/0.030/1.0+7.3 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1924 @98.151s 1%: 0.004+0+1.2 ms clock, 0.075+0/0.030/1.0+20 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1925 @98.152s 1%: 0.004+0+1.2 ms clock, 0.073+0/0.030/1.0+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1926 @98.158s 1%: 0.021+0.48+0.36 ms clock, 0.34+0/0.29/0.69+5.8 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1927 @98.352s 1%: 0.005+0+1.4 ms clock, 0.080+0/0.29/0.69+22 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1928 @98.354s 1%: 0.045+0+1.2 ms clock, 0.72+0/0.29/0.69+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1929 @98.360s 1%: 0.022+0.42+0.40 ms clock, 0.36+0/0.23/0.81+6.4 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1930 @98.549s 1%: 0.004+0+1.4 ms clock, 0.073+0/0.23/0.81+23 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1931 @98.550s 1%: 0.045+0+1.2 ms clock, 0.73+0/0.23/0.81+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1932 @98.556s 1%: 0.021+0.45+0.31 ms clock, 0.34+0/0.069/1.6+5.0 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1933 @98.768s 1%: 0.002+0+1.2 ms clock, 0.046+0/0.069/1.6+19 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1934 @98.769s 1%: 0.006+0+1.1 ms clock, 0.10+0/0.069/1.6+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1935 @98.776s 1%: 0.020+0.42+0.29 ms clock, 0.33+0/0.49/0.73+4.7 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24073200 HeapInuse: 24322048 HeapObjects: 328 HeapIdle 143646720 HeapReleased 0 HeapSys 167968768
gc 1936 @98.967s 1%: 0.005+0+1.4 ms clock, 0.081+0/0.49/0.73+22 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
 1000000	       197 ns/op	      23 B/op	       0 allocs/op
gc 1937 @98.969s 1%: 0.044+0+1.2 ms clock, 0.71+0/0.49/0.73+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1938 @98.988s 1%: 0.023+289+0.45 ms clock, 0.37+0/0.15/289+7.2 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
HeapAlloc: 80204736 HeapInuse: 80445440 HeapObjects: 325 HeapIdle 87523328 HeapReleased 0 HeapSys 167968768
gc 1939 @102.851s 1%: 0.002+0+0.94 ms clock, 0.044+0/0.15/289+15 ms cpu, 76->76->0 MB, 76 MB goal, 16 P (forced)
#BenchmarkSort1E7-16                               	10000000	       388 ns/op	       8 B/op	       0 allocs/op
gc 1940 @102.852s 1%: 0.050+0+0.97 ms clock, 0.80+0/0.15/289+15 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1941 @102.959s 1%: 0.018+4.1+0.32 ms clock, 0.30+0/11/15+5.2 ms cpu, 4->4->3 MB, 8 MB goal, 16 P
gc 1942 @103.057s 1%: 0.017+4.3+0.31 ms clock, 0.27+0.26/13/28+5.0 ms cpu, 5->6->5 MB, 7 MB goal, 16 P
gc 1943 @103.150s 1%: 0.018+7.1+0.32 ms clock, 0.29+0/26/45+5.1 ms cpu, 10->10->10 MB, 11 MB goal, 16 P
gc 1944 @103.432s 1%: 0.020+48+0.46 ms clock, 0.32+0.61/83/1.4+7.4 ms cpu, 19->19->19 MB, 20 MB goal, 16 P
gc 1945 @104.062s 1%: 0.022+49+0.33 ms clock, 0.36+0.21/184/6.2+5.4 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
gc 1946 @105.144s 1%: 0.018+48+0.30 ms clock, 0.29+0.25/189/430+4.9 ms cpu, 72->73->73 MB, 74 MB goal, 16 P
gc 1947 @106.940s 1%: 0.018+91+0.31 ms clock, 0.29+0.51/359/919+5.0 ms cpu, 143->144->144 MB, 146 MB goal, 16 P
gc 1948 @111.172s 1%: 0.037+170+0.32 ms clock, 0.59+0/673/1820+5.2 ms cpu, 281->284->283 MB, 289 MB goal, 16 P
GC forced
gc 3 @121.117s 0%: 0.25+12+0.44 ms clock, 2.3+0/1.8/2.5+3.9 ms cpu, 3->3->0 MB, 4 MB goal, 16 P
gc 1949 @119.704s 1%: 0.054+328+0.77 ms clock, 0.87+0.89/1311/3590+12 ms cpu, 553->557->556 MB, 567 MB goal, 16 P
HeapAlloc: 880202208 HeapInuse: 883564544 HeapObjects: 15000346 HeapIdle 581632 HeapReleased 0 HeapSys 884146176
gc 1950 @129.108s 1%: 0.002+0+36 ms clock, 0.038+0.89/1311/3590+581 ms cpu, 840->840->0 MB, 840 MB goal, 16 P (forced)
#BenchmarkSetInsert1E7-16                          	10000000	      2629 ns/op	      88 B/op	       2 allocs/op
gc 1951 @129.146s 1%: 0.006+0+4.3 ms clock, 0.096+0.89/1311/3590+69 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1952 @129.170s 1%: 0.027+1.1+0.35 ms clock, 0.44+0.10/1.3/5.6+5.6 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1953 @129.204s 1%: 0.027+2.2+0.34 ms clock, 0.44+0.17/6.1/5.1+5.5 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1954 @130.323s 1%: 0.017+21+0.28 ms clock, 0.28+0.26/84/226+4.5 ms cpu, 115->116->116 MB, 154 MB goal, 16 P
gc 1955 @132.618s 1%: 0.019+83+0.33 ms clock, 0.30+0.19/315/739+5.4 ms cpu, 202->203->203 MB, 232 MB goal, 16 P
gc 1956 @137.451s 1%: 0.023+176+0.32 ms clock, 0.37+0.53/690/1796+5.1 ms cpu, 379->382->381 MB, 406 MB goal, 16 P
gc 1957 @147.271s 1%: 0.018+402+0.28 ms clock, 0.29+1.1/1605/4349+4.5 ms cpu, 734->739->738 MB, 762 MB goal, 16 P
scvg0: inuse: 823, idle: 20, sys: 844, released: 0, consumed: 844 (MB)
gc 1958 @166.369s 1%: 0.079+189+0.33 ms clock, 1.2+0.22/747/2028+5.3 ms cpu, 1439->1442->396 MB, 1476 MB goal, 16 P
HeapAlloc: 746296448 HeapInuse: 748912640 HeapObjects: 11354169 HeapIdle 766476288 HeapReleased 0 HeapSys 1515388928
gc 1959 @174.172s 1%: 0.017+0+35 ms clock, 0.28+0.22/747/2028+568 ms cpu, 712->712->0 MB, 712 MB goal, 16 P (forced)
#BenchmarkSetErase1E7-16                           	10000000	      2159 ns/op	      96 B/op	       2 allocs/op
gc 1960 @174.211s 1%: 0.021+0+7.5 ms clock, 0.34+0.22/747/2028+120 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1961 @174.236s 1%: 0.033+1.7+0.30 ms clock, 0.53+0.053/2.2/11+4.9 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1962 @174.292s 1%: 0.027+2.3+0.38 ms clock, 0.43+0.17/6.7/8.4+6.1 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1963 @175.441s 1%: 0.015+28+0.30 ms clock, 0.25+0.23/107/217+4.9 ms cpu, 115->116->116 MB, 155 MB goal, 16 P
gc 1964 @177.698s 1%: 0.019+29+0.33 ms clock, 0.30+0/106/257+5.3 ms cpu, 202->203->118 MB, 232 MB goal, 16 P
scvg0: inuse: 3, idle: 1, sys: 5, released: 0, consumed: 5 (MB)
gc 1965 @180.581s 1%: 0.017+47+0.32 ms clock, 0.27+0.50/164/333+5.2 ms cpu, 221->222->138 MB, 237 MB goal, 16 P
gc 1966 @183.616s 1%: 0.020+12+0.29 ms clock, 0.32+0.28/49/113+4.7 ms cpu, 267->267->99 MB, 277 MB goal, 16 P
gc 1967 @185.970s 1%: 0.022+53+0.81 ms clock, 0.36+0.32/160/320+12 ms cpu, 193->194->128 MB, 198 MB goal, 16 P
gc 1968 @188.843s 1%: 0.023+38+0.28 ms clock, 0.37+0.66/153/391+4.6 ms cpu, 250->250->148 MB, 256 MB goal, 16 P
gc 1969 @192.047s 1%: 0.021+35+0.29 ms clock, 0.33+0.46/123/218+4.6 ms cpu, 288->289->121 MB, 296 MB goal, 16 P
gc 1970 @194.601s 1%: 0.017+6.7+0.31 ms clock, 0.28+0/22/32+5.0 ms cpu, 236->236->85 MB, 242 MB goal, 16 P
HeapAlloc: 98439680 HeapInuse: 98746368 HeapObjects: 311149 HeapIdle 1416642560 HeapReleased 0 HeapSys 1515388928
gc 1971 @194.768s 1%: 0.011+0+8.4 ms clock, 0.17+0/22/32+135 ms cpu, 93->93->0 MB, 93 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E7-16                  	10000000	      2055 ns/op	      96 B/op	       2 allocs/op
gc 1972 @194.779s 1%: 0.007+0+7.0 ms clock, 0.12+0/22/32+112 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1973 @194.806s 1%: 0.016+1.8+0.33 ms clock, 0.26+0.037/3.0/10+5.4 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1974 @194.849s 1%: 0.014+2.3+0.35 ms clock, 0.22+0.16/6.4/9.0+5.7 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1975 @196.144s 1%: 0.020+28+0.30 ms clock, 0.32+0.26/92/151+4.8 ms cpu, 115->116->116 MB, 155 MB goal, 16 P
gc 1976 @200.504s 1%: 0.015+18+0.29 ms clock, 0.25+0.31/65/146+4.7 ms cpu, 202->202->110 MB, 232 MB goal, 16 P
gc 1977 @205.132s 1%: 0.019+22+0.27 ms clock, 0.30+0.41/84/193+4.3 ms cpu, 207->207->115 MB, 221 MB goal, 16 P
gc 1978 @209.961s 1%: 0.018+31+0.36 ms clock, 0.30+0.27/119/284+5.8 ms cpu, 223->224->132 MB, 231 MB goal, 16 P
gc 1979 @216.837s 1%: 0.021+13+0.33 ms clock, 0.34+0/51/104+5.3 ms cpu, 257->257->100 MB, 264 MB goal, 16 P
HeapAlloc: 107070408 HeapInuse: 107470848 HeapObjects: 559227 HeapIdle 1407918080 HeapReleased 0 HeapSys 1515388928
gc 1980 @217.357s 1%: 0.009+0+11 ms clock, 0.15+0/51/104+184 ms cpu, 102->102->0 MB, 102 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E7-16          	10000000	      2258 ns/op	      56 B/op	       1 allocs/op
gc 1981 @217.372s 1%: 0.012+0+9.2 ms clock, 0.19+0/51/104+148 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1982 @217.482s 1%: 0.021+3.5+0.32 ms clock, 0.35+0.20/10/31+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 1983 @217.591s 1%: 0.021+6.8+0.30 ms clock, 0.35+0.49/23/31+4.9 ms cpu, 7->7->6 MB, 8 MB goal, 16 P
gc 1984 @217.708s 1%: 0.018+7.7+0.31 ms clock, 0.29+0.45/28/53+5.0 ms cpu, 13->13->13 MB, 14 MB goal, 16 P
gc 1985 @217.949s 1%: 0.021+16+0.59 ms clock, 0.34+0.51/60/122+9.5 ms cpu, 25->26->24 MB, 26 MB goal, 16 P
gc 1986 @218.348s 1%: 0.021+77+0.40 ms clock, 0.34+11/136/185+6.4 ms cpu, 48->49->47 MB, 49 MB goal, 16 P
gc 1987 @219.385s 1%: 0.019+51+0.31 ms clock, 0.31+5.6/200/473+5.0 ms cpu, 92->94->90 MB, 95 MB goal, 16 P
gc 1988 @220.871s 1%: 0.017+94+0.32 ms clock, 0.27+0.58/372/968+5.2 ms cpu, 176->178->170 MB, 180 MB goal, 16 P
gc 1989 @223.895s 1%: 0.020+177+0.30 ms clock, 0.32+1.0/707/1740+4.9 ms cpu, 332->336->321 MB, 340 MB goal, 16 P
gc 1990 @228.779s 1%: 0.019+304+0.31 ms clock, 0.31+1.0/1216/3333+4.9 ms cpu, 625->636->607 MB, 642 MB goal, 16 P
HeapAlloc: 817637808 HeapInuse: 821002240 HeapObjects: 11090064 HeapIdle 694386688 HeapReleased 0 HeapSys 1515388928
###gc 1991 @231.587s 1%: 0.006+0+40 ms clock, 0.10+1.0/1216/3333+640 ms cpu, 780->780->0 MB, 780 MB goal, 16 P (forced)
#BenchmarkIntSetInsert1E7-16                       	10000000	      1424 ns/op	      88 B/op	       2 allocs/op
gc 1992 @231.630s 1%: 0.015+0+7.0 ms clock, 0.25+1.0/1216/3333+113 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1993 @231.655s 1%: 0.026+1.5+0.30 ms clock, 0.42+0.007/2.1/9.5+4.8 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1994 @231.683s 1%: 0.020+2.1+0.33 ms clock, 0.32+0.11/6.5/6.2+5.4 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1995 @232.348s 1%: 0.019+25+0.45 ms clock, 0.31+0.20/83/175+7.2 ms cpu, 115->116->112 MB, 154 MB goal, 16 P
gc 1996 @233.647s 1%: 0.019+55+0.34 ms clock, 0.31+0.67/215/566+5.4 ms cpu, 196->198->190 MB, 225 MB goal, 16 P
gc 1997 @236.267s 1%: 0.016+140+0.29 ms clock, 0.26+1.0/540/1403+4.7 ms cpu, 353->357->341 MB, 380 MB goal, 16 P
GC forced
gc 4 @241.132s 0%: 0.13+0.59+0.28 ms clock, 1.5+0/1.6/2.1+3.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P
gc 1998 @241.517s 1%: 0.025+282+0.41 ms clock, 0.40+0/1127/3095+6.6 ms cpu, 656->665->635 MB, 682 MB goal, 16 P
gc 1999 @250.429s 1%: 0.019+220+0.33 ms clock, 0.31+0/878/2356+5.3 ms cpu, 1238->1243->500 MB, 1271 MB goal, 16 P
HeapAlloc: 1006226096 HeapInuse: 1009926144 HeapObjects: 13786773 HeapIdle 505462784 HeapReleased 0 HeapSys 1515388928
###gc 2000 @257.518s 1%: 0.011+0+46 ms clock, 0.17+0/878/2356+748 ms cpu, 961->961->0 MB, 961 MB goal, 16 P (forced)
#BenchmarkIntSetErase1E7-16                        	10000000	      1256 ns/op	      96 B/op	       2 allocs/op
gc 2001 @257.568s 1%: 0.018+0+7.8 ms clock, 0.29+0/878/2356+125 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2002 @257.592s 1%: 0.035+1.3+0.31 ms clock, 0.57+0/1.9/7.4+5.0 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 2003 @257.632s 1%: 0.025+2.1+0.32 ms clock, 0.41+0.25/4.8/9.2+5.2 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 2004 @258.319s 1%: 0.017+25+0.38 ms clock, 0.28+0.22/87/179+6.1 ms cpu, 115->116->112 MB, 154 MB goal, 16 P
gc 2005 @259.908s 1%: 0.024+35+0.33 ms clock, 0.39+0/124/242+5.3 ms cpu, 196->197->117 MB, 225 MB goal, 16 P
gc 2006 @261.475s 1%: 0.020+30+0.30 ms clock, 0.32+0.75/117/273+4.9 ms cpu, 219->220->129 MB, 235 MB goal, 16 P
gc 2007 @263.294s 1%: 0.024+10+0.39 ms clock, 0.39+0.43/33/54+6.2 ms cpu, 250->250->86 MB, 259 MB goal, 16 P
gc 2008 @264.906s 1%: 0.018+42+0.41 ms clock, 0.29+0/159/396+6.6 ms cpu, 168->170->144 MB, 173 MB goal, 16 P
gc 2009 @267.182s 1%: 0.022+35+0.34 ms clock, 0.36+0.77/137/321+5.4 ms cpu, 282->283->135 MB, 289 MB goal, 16 P
gc 2010 @269.167s 1%: 0.024+16+0.34 ms clock, 0.38+0/58/134+5.5 ms cpu, 263->263->99 MB, 270 MB goal, 16 P
gc 2011 @270.631s 1%: 0.019+39+0.43 ms clock, 0.31+0.48/123/194+6.9 ms cpu, 193->194->121 MB, 198 MB goal, 16 P
HeapAlloc: 178385616 HeapInuse: 178995200 HeapObjects: 1461855 HeapIdle 1336393728 HeapReleased 0 HeapSys 1515388928
gc 2012 @271.241s 1%: 0.006+0+12 ms clock, 0.096+0.48/123/194+197 ms cpu, 170->170->0 MB, 170 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndErase1E7-16               	10000000	      1367 ns/op	      96 B/op	       2 allocs/op
gc 2013 @271.256s 1%: 0.011+0+7.2 ms clock, 0.17+0.48/123/194+116 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2014 @271.282s 1%: 0.019+1.8+0.40 ms clock, 0.31+0.026/1.7/8.4+6.4 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 2015 @271.317s 1%: 0.019+1.9+0.30 ms clock, 0.31+0.13/6.2/9.9+4.9 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 2016 @272.551s 1%: 0.021+27+0.51 ms clock, 0.35+0.24/99/209+8.2 ms cpu, 115->116->112 MB, 154 MB goal, 16 P
gc 2017 @275.415s 1%: 0.018+27+0.33 ms clock, 0.30+0/71/125+5.3 ms cpu, 196->196->105 MB, 225 MB goal, 16 P
gc 2018 @278.451s 1%: 0.018+17+0.40 ms clock, 0.29+0.51/60/143+6.4 ms cpu, 196->197->105 MB, 210 MB goal, 16 P
gc 2019 @281.608s 1%: 0.018+20+0.34 ms clock, 0.29+0.31/79/163+5.5 ms cpu, 203->203->111 MB, 210 MB goal, 16 P
gc 2020 @285.120s 1%: 0.026+29+0.40 ms clock, 0.42+0.20/108/222+6.4 ms cpu, 216->217->123 MB, 222 MB goal, 16 P
HeapAlloc: 171280424 HeapInuse: 171851776 HeapObjects: 1692354 HeapIdle 1343537152 HeapReleased 0 HeapSys 1515388928
gc 2021 @287.052s 1%: 0.007+0+13 ms clock, 0.11+0.20/108/222+213 ms cpu, 163->163->0 MB, 163 MB goal, 16 P (forced)
#BenchmarkIntSetInsertAndEraseWithPool1E7-16       	10000000	      1579 ns/op	      56 B/op	       1 allocs/op
gc 2022 @287.068s 1%: 0.017+0+8.1 ms clock, 0.28+0.20/108/222+131 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2023 @287.099s 1%: 0.016+1.2+0.41 ms clock, 0.26+0.11/1.6/8.9+6.6 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 2024 @287.129s 1%: 0.018+1.3+0.33 ms clock, 0.30+0.33/0.83/9.6+5.3 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 2025 @287.185s 1%: 0.017+1.3+0.59 ms clock, 0.27+0.30/0.70/8.8+9.4 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 2026 @287.214s 1%: 0.022+1.3+0.33 ms clock, 0.35+1.1/0.79/10+5.3 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 2027 @287.328s 1%: 0.019+1.8+0.38 ms clock, 0.31+0/1.0/10+6.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 2028 @287.398s 1%: 0.021+2.5+0.31 ms clock, 0.35+2.3/0.82/13+5.1 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
gc 2029 @287.586s 1%: 0.020+1.3+0.54 ms clock, 0.33+0.53/0.75/10+8.7 ms cpu, 36->36->22 MB, 71 MB goal, 16 P
gc 2030 @287.762s 1%: 0.022+3.6+0.36 ms clock, 0.36+3.5/0.86/13+5.8 ms cpu, 73->73->71 MB, 74 MB goal, 16 P
gc 2031 @288.105s 1%: 0.021+4.6+0.32 ms clock, 0.34+0.62/4.3/10+5.2 ms cpu, 72->72->72 MB, 143 MB goal, 16 P
gc 2032 @288.563s 1%: 0.022+7.1+0.32 ms clock, 0.35+6.9/1.2/18+5.2 ms cpu, 174->174->143 MB, 175 MB goal, 16 P
gc 2033 @289.760s 1%: 0.018+54+0.92 ms clock, 0.30+0/10/53+14 ms cpu, 169->169->107 MB, 286 MB goal, 16 P
gc 2034 @289.929s 1%: 0.022+14+0.36 ms clock, 0.35+14/0.56/22+5.9 ms cpu, 286->286->286 MB, 287 MB goal, 16 P
HeapAlloc: 315467992 HeapInuse: 315940864 HeapObjects: 332525 HeapIdle 1199448064 HeapReleased 0 HeapSys 1515388928
gc 2035 @291.669s 1%: 0.003+0+15 ms clock, 0.053+14/0.56/22+250 ms cpu, 300->300->187 MB, 300 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E7-16                   	10000000	       460 ns/op	      44 B/op	       0 allocs/op
gc 2036 @291.688s 1%: 0.003+0+9.2 ms clock, 0.056+14/0.56/22+148 ms cpu, 187->187->0 MB, 187 MB goal, 16 P (forced)
gc 2037 @291.714s 1%: 0.019+1.3+0.27 ms clock, 0.31+0/1.3/7.2+4.4 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 2038 @291.743s 1%: 0.021+1.4+0.38 ms clock, 0.34+0.038/0.16/8.6+6.0 ms cpu, 77->78->77 MB, 152 MB goal, 16 P
gc 2039 @292.062s 1%: 0.025+2.0+0.33 ms clock, 0.41+1.8/3.0/10+5.4 ms cpu, 127->127->112 MB, 155 MB goal, 16 P
gc 2040 @293.267s 1%: 0.019+5.9+0.31 ms clock, 0.30+5.7/0.19/14+5.0 ms cpu, 266->266->219 MB, 267 MB goal, 16 P
gc 2041 @294.803s 1%: 0.016+8.4+0.32 ms clock, 0.26+8.4/1.2/17+5.1 ms cpu, 424->424->362 MB, 438 MB goal, 16 P
HeapAlloc: 395531816 HeapInuse: 396025856 HeapObjects: 333159 HeapIdle 1119363072 HeapReleased 0 HeapSys 1515388928
gc 2042 @298.707s 1%: 0.004+0+13 ms clock, 0.071+8.4/1.2/17+217 ms cpu, 377->377->187 MB, 377 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E7-16                    	10000000	       205 ns/op	       8 B/op	       0 allocs/op
gc 2043 @298.723s 1%: 0.007+0+8.4 ms clock, 0.12+8.4/1.2/17+135 ms cpu, 187->187->0 MB, 187 MB goal, 16 P (forced)
gc 2044 @298.736s 1%: 0.016+1.3+0.32 ms clock, 0.26+0/2.4/8.1+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 2045 @298.758s 1%: 0.018+1.2+0.60 ms clock, 0.29+0.052/0.99/8.4+9.7 ms cpu, 9->9->8 MB, 15 MB goal, 16 P
gc 2046 @298.779s 1%: 0.014+1.2+0.37 ms clock, 0.23+0.17/0.13/9.7+5.9 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 2047 @298.880s 1%: 0.020+1.3+0.42 ms clock, 0.33+0.56/0.63/9.6+6.8 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 2048 @298.908s 1%: 0.017+2.4+0.30 ms clock, 0.28+0.021/1.1/10+4.9 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 2049 @299.014s 1%: 0.024+1.5+0.31 ms clock, 0.38+1.3/1.5/11+5.0 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
scvg1: inuse: 45, idle: 1400, sys: 1446, released: 0, consumed: 1446 (MB)
HeapAlloc: 48533256 HeapInuse: 48906240 HeapObjects: 48461 HeapIdle 1466449920 HeapReleased 0 HeapSys 1515356160
gc 2050 @300.999s 1%: 0.013+0+8.2 ms clock, 0.21+1.3/1.5/11+131 ms cpu, 46->46->24 MB, 46 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E7-16           	10000000	       227 ns/op	       6 B/op	       0 allocs/op
gc 2051 @301.010s 1%: 0.004+0+7.6 ms clock, 0.069+1.3/1.5/11+122 ms cpu, 24->24->0 MB, 24 MB goal, 16 P (forced)
gc 2052 @301.021s 1%: 0.013+1.3+0.38 ms clock, 0.21+0/1.3/6.3+6.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 2053 @301.073s 1%: 0.021+1.5+0.35 ms clock, 0.34+0/0.17/12+5.6 ms cpu, 183->183->183 MB, 184 MB goal, 16 P
HeapAlloc: 192754880 HeapInuse: 193003520 HeapObjects: 329 HeapIdle 1322385408 HeapReleased 0 HeapSys 1515388928
gc 2054 @303.054s 1%: 0.004+0+7.5 ms clock, 0.067+0/0.17/12+120 ms cpu, 183->183->176 MB, 183 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E7-16    	10000000	       204 ns/op	      19 B/op	       0 allocs/op
gc 2055 @303.064s 1%: 0.003+0+7.5 ms clock, 0.054+0/0.17/12+120 ms cpu, 176->176->0 MB, 176 MB goal, 16 P (forced)
#BenchmarkIntSetInsert-16                          	gc 2056 @303.077s 1%: 0.002+0+7.8 ms clock, 0.044+0/0.17/12+125 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2057 @303.090s 1%: 0.042+0+7.9 ms clock, 0.68+0/0.17/12+127 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2058 @303.114s 1%: 0.072+0+7.6 ms clock, 1.1+0/0.17/12+122 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 2059 @303.183s 1%: 0.011+3.0+0.32 ms clock, 0.17+0.035/9.2/26+5.1 ms cpu, 4->4->3 MB, 8 MB goal, 16 P
gc 2060 @303.229s 1%: 0.017+4.7+0.34 ms clock, 0.28+0.38/13/25+5.5 ms cpu, 5->5->5 MB, 7 MB goal, 16 P
gc 2061 @303.284s 1%: 0.014+6.7+0.31 ms clock, 0.22+0.25/24/42+4.9 ms cpu, 9->9->9 MB, 10 MB goal, 16 P
gc 2062 @303.431s 1%: 0.021+10+0.37 ms clock, 0.34+0.46/38/82+6.0 ms cpu, 16->17->16 MB, 18 MB goal, 16 P
gc 2063 @303.620s 1%: 0.019+18+0.31 ms clock, 0.30+0.62/66/144+5.0 ms cpu, 30->31->29 MB, 32 MB goal, 16 P
gc 2064 @304.220s 1%: 0.018+30+0.46 ms clock, 0.30+0.45/119/298+7.4 ms cpu, 57->57->55 MB, 59 MB goal, 16 P
 1000000	      1661 ns/op	      88 B/op	       2 allocs/op
PASS
ok  	github.com/cdongyang/library/rbtree	304.972s
