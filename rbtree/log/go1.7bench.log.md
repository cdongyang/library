gc 1 @0.080s 0%: 0.13+0.90+0.32 ms clock, 0.65+0.096/1.0/1.0+1.6 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
gc 2 @0.107s 0%: 0.11+0.73+0.48 ms clock, 1.0+0.017/1.8/0.60+4.3 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
# github.com/cdongyang/library/rbtree
gc 1 @0.005s 4%: 0.074+4.9+0.26 ms clock, 0.22+0.34/6.0/1.0+0.78 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @0.045s 2%: 0.068+3.6+0.30 ms clock, 0.75+0.37/6.4/6.6+3.3 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 3 @0.110s 1%: 0.022+4.3+0.30 ms clock, 0.36+0.45/4.3/13+4.9 ms cpu, 9->9->5 MB, 10 MB goal, 16 P
gc 4 @0.178s 1%: 0.025+4.8+0.29 ms clock, 0.40+0.38/8.8/13+4.7 ms cpu, 10->10->6 MB, 11 MB goal, 16 P
# github.com/cdongyang/library/rbtree_test
gc 1 @0.006s 3%: 0.044+3.9+0.29 ms clock, 0.17+0.094/5.0/0.98+1.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @0.026s 3%: 0.042+4.1+0.28 ms clock, 0.47+0.15/5.4/12+3.1 ms cpu, 6->7->5 MB, 7 MB goal, 16 P
gc 3 @0.084s 1%: 0.021+4.5+0.29 ms clock, 0.34+0.23/7.4/21+4.7 ms cpu, 10->10->7 MB, 11 MB goal, 16 P
gc 4 @0.174s 1%: 0.018+10+0.36 ms clock, 0.30+0/18/8.8+5.8 ms cpu, 13->13->8 MB, 14 MB goal, 16 P
# testmain
gc 1 @0.006s 4%: 0.10+3.9+0.45 ms clock, 0.30+0.009/5.2/1.5+1.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 2 @0.025s 4%: 0.040+4.6+0.31 ms clock, 0.44+0.15/10/4.5+3.4 ms cpu, 6->7->5 MB, 7 MB goal, 16 P
# testmain
gc 1 @0.002s 4%: 0.12+4.6+0.27 ms clock, 0.24+3.8/0.97/0.15+0.55 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 2 @0.020s 3%: 0.046+6.1+0.29 ms clock, 0.32+0/9.0/3.5+2.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 3 @0.056s 3%: 0.064+7.6+0.30 ms clock, 0.57+0.22/19/2.4+2.7 ms cpu, 13->13->12 MB, 14 MB goal, 16 P
gc 4 @0.115s 4%: 0.062+19+0.29 ms clock, 0.69+0.45/44/1.3+3.2 ms cpu, 24->25->23 MB, 25 MB goal, 16 P
=== RUN   TestSet
--- PASS: TestSet (0.00s)
gc 1 @0.002s 4%: 0.005+0+0.59 ms clock, 0.021+0/0/0+2.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 2 @0.003s 7%: 0.007+0+0.48 ms clock, 0.037+0/0/0+2.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 182480 HeapInuse: 466944 HeapObjects: 246 HeapIdle 155648 HeapReleased 0 HeapSys 622592
gc 3 @0.004s 10%: 0.006+0+0.43 ms clock, 0.055+0/0/0+3.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSort1E3-16                               	gc 4 @0.005s 12%: 0.017+0+0.32 ms clock, 0.17+0/0/0+3.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 191888 HeapInuse: 466944 HeapObjects: 265 HeapIdle 974848 HeapReleased 0 HeapSys 1441792
gc 5 @0.006s 16%: 0.005+0+0.46 ms clock, 0.077+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 6 @0.006s 20%: 0.004+0+0.38 ms clock, 0.066+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 7 @0.007s 24%: 0.002+0+0.54 ms clock, 0.046+0/0/0+8.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 8 @0.008s 26%: 0.003+0+0.35 ms clock, 0.058+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 9 @0.009s 28%: 0.002+0+0.42 ms clock, 0.038+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 10 @0.009s 30%: 0.004+0+0.32 ms clock, 0.071+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 11 @0.010s 31%: 0.002+0+0.40 ms clock, 0.036+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 12 @0.011s 33%: 0.004+0+0.31 ms clock, 0.066+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 13 @0.011s 33%: 0.004+0+0.32 ms clock, 0.068+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 14 @0.012s 34%: 0.004+0+0.29 ms clock, 0.071+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 15 @0.012s 35%: 0.005+0+0.32 ms clock, 0.082+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 16 @0.013s 36%: 0.004+0+0.31 ms clock, 0.062+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 17 @0.014s 36%: 0.002+0+0.31 ms clock, 0.040+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 18 @0.014s 37%: 0.005+0+0.39 ms clock, 0.085+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 19 @0.015s 38%: 0.002+0+0.31 ms clock, 0.038+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 20 @0.015s 38%: 0.005+0+0.32 ms clock, 0.080+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 21 @0.016s 39%: 0.002+0+0.43 ms clock, 0.041+0/0/0+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 22 @0.017s 41%: 0.004+0+0.75 ms clock, 0.069+0/0/0+11 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 23 @0.018s 41%: 0.001+0+0.37 ms clock, 0.027+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 24 @0.018s 42%: 0.004+0+0.55 ms clock, 0.071+0/0/0+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 25 @0.019s 43%: 0.005+0+0.47 ms clock, 0.084+0/0/0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 26 @0.020s 43%: 0.046+0+0.38 ms clock, 0.70+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 27 @0.021s 44%: 0.002+0+0.53 ms clock, 0.037+0/0/0+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 28 @0.022s 45%: 0.042+0+0.43 ms clock, 0.63+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 29 @0.022s 45%: 0.016+0+0.35 ms clock, 0.25+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 30 @0.023s 46%: 0.045+0+0.43 ms clock, 0.68+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 31 @0.024s 45%: 0.002+0+0.31 ms clock, 0.036+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 32 @0.024s 46%: 0.004+0+0.36 ms clock, 0.064+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 197456 HeapInuse: 475136 HeapObjects: 274 HeapIdle 868352 HeapReleased 0 HeapSys 1343488
gc 33 @0.025s 46%: 0.017+0+0.37 ms clock, 0.25+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       768 ns/op	      11 B/op	       0 allocs/op
gc 34 @0.026s 47%: 0.005+0+0.53 ms clock, 0.078+0/0/0+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277440 HeapInuse: 532480 HeapObjects: 1773 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 35 @0.029s 43%: 0.002+0+0.38 ms clock, 0.036+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsert1E3-16                          	gc 36 @0.030s 43%: 0.007+0+0.39 ms clock, 0.11+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 37 @0.033s 40%: 0.002+0+0.33 ms clock, 0.033+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 38 @0.033s 40%: 0.004+0+0.36 ms clock, 0.068+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 39 @0.037s 38%: 0.002+0+0.41 ms clock, 0.035+0/0/0+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 40 @0.037s 38%: 0.004+0+0.49 ms clock, 0.066+0/0/0+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 41 @0.041s 36%: 0.018+0+0.38 ms clock, 0.27+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 42 @0.041s 37%: 0.028+0+0.33 ms clock, 0.42+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 43 @0.044s 35%: 0.013+0+0.52 ms clock, 0.19+0/0/0+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 44 @0.045s 35%: 0.013+0+0.45 ms clock, 0.20+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 45 @0.048s 34%: 0.015+0+0.33 ms clock, 0.22+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 46 @0.049s 34%: 0.042+0+0.33 ms clock, 0.64+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 47 @0.052s 33%: 0.019+0+0.37 ms clock, 0.29+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 48 @0.053s 33%: 0.045+0+0.31 ms clock, 0.67+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 49 @0.056s 32%: 0.002+0+0.34 ms clock, 0.036+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 50 @0.056s 32%: 0.045+0+0.35 ms clock, 0.68+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 51 @0.060s 31%: 0.004+0+0.48 ms clock, 0.073+0/0/0+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 52 @0.060s 31%: 0.004+0+0.33 ms clock, 0.066+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 53 @0.063s 30%: 0.004+0+0.32 ms clock, 0.073+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 54 @0.064s 31%: 0.016+0+0.45 ms clock, 0.24+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 55 @0.068s 29%: 0.002+0+0.33 ms clock, 0.036+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 56 @0.069s 30%: 0.015+0+0.38 ms clock, 0.23+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 57 @0.072s 29%: 0.014+0+0.38 ms clock, 0.21+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 58 @0.072s 29%: 0.045+0+0.37 ms clock, 0.68+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 59 @0.075s 28%: 0.002+0+0.34 ms clock, 0.037+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 60 @0.076s 29%: 0.004+0+0.37 ms clock, 0.066+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 61 @0.079s 28%: 0.016+0+0.27 ms clock, 0.25+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 62 @0.079s 28%: 0.004+0+0.38 ms clock, 0.060+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 63 @0.082s 28%: 0.001+0+0.42 ms clock, 0.028+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 64 @0.083s 28%: 0.004+0+0.27 ms clock, 0.060+0/0/0+4.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 65 @0.086s 27%: 0.008+0+0.36 ms clock, 0.12+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 66 @0.086s 28%: 0.042+0+0.36 ms clock, 0.63+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 277504 HeapInuse: 532480 HeapObjects: 1776 HeapIdle 811008 HeapReleased 0 HeapSys 1343488
gc 67 @0.090s 27%: 0.004+0+0.35 ms clock, 0.069+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      3067 ns/op	      91 B/op	       2 allocs/op
gc 68 @0.090s 27%: 0.007+0+0.31 ms clock, 0.11+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373632 HeapInuse: 630784 HeapObjects: 3274 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 69 @0.095s 26%: 0.004+0+0.36 ms clock, 0.067+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetErase1E3-16                           	gc 70 @0.096s 26%: 0.014+0+0.38 ms clock, 0.21+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 71 @0.101s 25%: 0.004+0+0.34 ms clock, 0.068+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 72 @0.102s 25%: 0.084+0+0.36 ms clock, 1.2+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 73 @0.106s 25%: 0.016+0+0.44 ms clock, 0.25+0/0/0+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 74 @0.107s 25%: 0.047+0+0.34 ms clock, 0.70+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 75 @0.112s 24%: 0.001+0+0.41 ms clock, 0.029+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 76 @0.112s 24%: 0.10+0+0.34 ms clock, 1.5+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 77 @0.118s 24%: 0.020+0+0.38 ms clock, 0.30+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 78 @0.118s 24%: 0.027+0+0.33 ms clock, 0.41+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 79 @0.123s 23%: 0.013+0+0.39 ms clock, 0.20+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 80 @0.123s 23%: 0.041+0+0.32 ms clock, 0.62+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 81 @0.128s 23%: 0.015+0+0.36 ms clock, 0.22+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 82 @0.129s 23%: 0.053+0+0.30 ms clock, 0.80+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 83 @0.134s 22%: 0.005+0+0.42 ms clock, 0.083+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 84 @0.134s 23%: 0.045+0+0.38 ms clock, 0.68+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 85 @0.139s 22%: 0.002+0+0.34 ms clock, 0.036+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 86 @0.140s 22%: 0.026+0+0.37 ms clock, 0.40+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 87 @0.145s 22%: 0.024+0+0.36 ms clock, 0.37+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 88 @0.145s 22%: 0.034+0+0.36 ms clock, 0.51+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 89 @0.150s 21%: 0.004+0+0.36 ms clock, 0.067+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 90 @0.151s 22%: 0.047+0+0.37 ms clock, 0.71+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 91 @0.156s 21%: 0.030+0+0.50 ms clock, 0.45+0/0/0+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 92 @0.156s 21%: 0.045+0+0.36 ms clock, 0.68+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 93 @0.162s 21%: 0.004+0+0.31 ms clock, 0.064+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 94 @0.162s 21%: 0.004+0+0.37 ms clock, 0.060+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 95 @0.167s 21%: 0.052+0+0.33 ms clock, 0.78+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 96 @0.167s 21%: 0.075+0+0.38 ms clock, 1.1+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 97 @0.172s 20%: 0.014+0+0.35 ms clock, 0.22+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 98 @0.173s 21%: 0.027+0+0.37 ms clock, 0.40+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 373696 HeapInuse: 630784 HeapObjects: 3277 HeapIdle 712704 HeapReleased 0 HeapSys 1343488
gc 99 @0.178s 20%: 0.015+0+0.40 ms clock, 0.22+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2271 ns/op	      99 B/op	       2 allocs/op
gc 100 @0.178s 20%: 0.007+0+0.38 ms clock, 0.10+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285632 HeapInuse: 540672 HeapObjects: 1774 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 101 @0.180s 20%: 0.004+0+0.37 ms clock, 0.067+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E3-16                  	gc 102 @0.181s 20%: 0.051+0+0.36 ms clock, 0.76+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 103 @0.183s 20%: 0.004+0+0.33 ms clock, 0.064+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 104 @0.184s 21%: 0.048+0+0.35 ms clock, 0.72+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 105 @0.186s 20%: 0.042+0+0.33 ms clock, 0.64+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 106 @0.186s 21%: 0.061+0+0.36 ms clock, 0.91+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 107 @0.189s 21%: 0.017+0+0.34 ms clock, 0.26+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 108 @0.189s 21%: 0.049+0+0.32 ms clock, 0.74+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 109 @0.191s 21%: 0.006+0+0.40 ms clock, 0.098+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 110 @0.192s 21%: 0.040+0+0.33 ms clock, 0.61+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 111 @0.194s 21%: 0.016+0+0.34 ms clock, 0.24+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 112 @0.195s 21%: 0.027+0+0.37 ms clock, 0.41+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 113 @0.197s 21%: 0.005+0+0.39 ms clock, 0.083+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 114 @0.198s 21%: 0.045+0+0.95 ms clock, 0.68+0/0/0+14 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 115 @0.200s 21%: 0.002+0+0.35 ms clock, 0.035+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 116 @0.201s 21%: 0.029+0+0.35 ms clock, 0.44+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 117 @0.203s 21%: 0.002+0+0.34 ms clock, 0.038+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 118 @0.203s 21%: 0.044+0+0.31 ms clock, 0.67+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 119 @0.205s 21%: 0.039+0+0.35 ms clock, 0.59+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 120 @0.206s 21%: 0.070+0+0.30 ms clock, 1.0+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 121 @0.208s 21%: 0.027+0+0.40 ms clock, 0.40+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 122 @0.209s 21%: 0.004+0+0.37 ms clock, 0.071+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 123 @0.211s 21%: 0.016+0+0.33 ms clock, 0.25+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 124 @0.211s 21%: 0.045+0+0.38 ms clock, 0.68+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 125 @0.213s 21%: 0.015+0+0.33 ms clock, 0.23+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 126 @0.214s 22%: 0.028+0+0.36 ms clock, 0.42+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 127 @0.216s 22%: 0.002+0+0.34 ms clock, 0.045+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 128 @0.216s 22%: 0.049+0+0.39 ms clock, 0.74+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 129 @0.219s 22%: 0.006+0+0.31 ms clock, 0.091+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 130 @0.219s 22%: 0.005+0+0.31 ms clock, 0.076+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 285696 HeapInuse: 540672 HeapObjects: 1777 HeapIdle 802816 HeapReleased 0 HeapSys 1343488
gc 131 @0.221s 22%: 0.018+0+0.57 ms clock, 0.27+0/0/0+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2278 ns/op	      99 B/op	       2 allocs/op
gc 132 @0.222s 22%: 0.007+0+0.32 ms clock, 0.11+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248600 HeapInuse: 524288 HeapObjects: 1280 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 133 @0.224s 22%: 0.004+0+0.33 ms clock, 0.068+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E3-16          	gc 134 @0.225s 22%: 0.005+0+0.37 ms clock, 0.085+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 135 @0.228s 22%: 0.002+0+0.34 ms clock, 0.041+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 136 @0.228s 22%: 0.043+0+0.36 ms clock, 0.64+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 137 @0.231s 22%: 0.016+0+0.57 ms clock, 0.24+0/0/0+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 138 @0.232s 22%: 0.027+0+0.32 ms clock, 0.41+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 139 @0.234s 22%: 0.015+0+0.30 ms clock, 0.22+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 140 @0.235s 22%: 0.046+0+0.29 ms clock, 0.69+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 141 @0.237s 22%: 0.002+0+0.34 ms clock, 0.043+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 142 @0.238s 22%: 0.004+0+0.30 ms clock, 0.069+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 143 @0.240s 22%: 0.019+0+0.36 ms clock, 0.29+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 144 @0.241s 22%: 0.043+0+0.31 ms clock, 0.65+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 145 @0.243s 22%: 0.002+0+0.32 ms clock, 0.040+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 146 @0.244s 22%: 0.030+0+0.29 ms clock, 0.46+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 147 @0.246s 22%: 0.014+0+1.4 ms clock, 0.21+0/0/0+21 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 148 @0.248s 22%: 0.004+0+0.27 ms clock, 0.068+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 149 @0.251s 22%: 0.021+0+0.35 ms clock, 0.31+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 150 @0.251s 22%: 0.045+0+0.29 ms clock, 0.68+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 151 @0.254s 22%: 0.002+0+0.34 ms clock, 0.046+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 152 @0.254s 22%: 0.042+0+0.33 ms clock, 0.64+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 153 @0.257s 22%: 0.004+0+0.39 ms clock, 0.071+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 154 @0.257s 22%: 0.004+0+0.37 ms clock, 0.064+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 155 @0.260s 22%: 0.021+0+0.34 ms clock, 0.31+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 156 @0.260s 22%: 0.042+0+0.33 ms clock, 0.64+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 157 @0.263s 22%: 0.025+0+0.35 ms clock, 0.37+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 158 @0.263s 22%: 0.037+0+0.34 ms clock, 0.55+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 159 @0.266s 22%: 0.020+0+0.33 ms clock, 0.30+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 160 @0.266s 22%: 0.003+0+0.38 ms clock, 0.057+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 161 @0.269s 22%: 0.019+0+0.41 ms clock, 0.28+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 162 @0.270s 22%: 0.039+0+0.34 ms clock, 0.59+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 163 @0.272s 22%: 0.002+0+0.36 ms clock, 0.039+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 164 @0.273s 22%: 0.041+0+0.31 ms clock, 0.62+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 248680 HeapInuse: 524288 HeapObjects: 1283 HeapIdle 819200 HeapReleased 0 HeapSys 1343488
gc 165 @0.275s 22%: 0.002+0+0.48 ms clock, 0.042+0/0/0+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	      2659 ns/op	      62 B/op	       1 allocs/op
gc 166 @0.276s 22%: 0.011+0+0.38 ms clock, 0.17+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244656 HeapInuse: 598016 HeapObjects: 367 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 167 @0.277s 22%: 0.004+0+0.36 ms clock, 0.078+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E3-16                   	gc 168 @0.278s 22%: 0.009+0+0.39 ms clock, 0.13+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245872 HeapInuse: 598016 HeapObjects: 384 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 169 @0.279s 22%: 0.048+0+0.37 ms clock, 0.73+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 170 @0.279s 22%: 0.043+0+0.40 ms clock, 0.65+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244488 HeapInuse: 598016 HeapObjects: 370 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 171 @0.280s 22%: 0.022+0+0.37 ms clock, 0.33+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 172 @0.281s 23%: 0.044+0+0.31 ms clock, 0.66+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 243408 HeapInuse: 598016 HeapObjects: 356 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 173 @0.282s 23%: 0.019+0+0.33 ms clock, 0.29+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 174 @0.282s 23%: 0.045+0+0.25 ms clock, 0.68+0/0/0+3.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244024 HeapInuse: 598016 HeapObjects: 364 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 175 @0.283s 23%: 0.013+0+0.34 ms clock, 0.20+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 176 @0.284s 23%: 0.045+0+0.32 ms clock, 0.68+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244000 HeapInuse: 598016 HeapObjects: 363 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 177 @0.285s 23%: 0.004+0+0.37 ms clock, 0.066+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 178 @0.285s 23%: 0.004+0+0.28 ms clock, 0.066+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244176 HeapInuse: 598016 HeapObjects: 364 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 179 @0.286s 23%: 0.001+0+0.27 ms clock, 0.031+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 180 @0.287s 23%: 0.046+0+0.28 ms clock, 0.69+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244736 HeapInuse: 598016 HeapObjects: 369 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 181 @0.288s 23%: 0.004+0+0.27 ms clock, 0.070+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 182 @0.288s 23%: 0.049+0+0.36 ms clock, 0.73+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244432 HeapInuse: 598016 HeapObjects: 367 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 183 @0.289s 23%: 0.002+0+0.31 ms clock, 0.034+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 184 @0.290s 23%: 0.004+0+0.35 ms clock, 0.066+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244152 HeapInuse: 598016 HeapObjects: 366 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 185 @0.291s 23%: 0.014+0+0.34 ms clock, 0.21+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 186 @0.291s 23%: 0.045+0+0.32 ms clock, 0.68+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 243992 HeapInuse: 598016 HeapObjects: 364 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 187 @0.292s 23%: 0.002+0+0.28 ms clock, 0.038+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 188 @0.292s 23%: 0.048+0+0.33 ms clock, 0.72+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244856 HeapInuse: 598016 HeapObjects: 373 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 189 @0.293s 23%: 0.027+0+0.29 ms clock, 0.41+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 190 @0.294s 24%: 0.047+0+0.29 ms clock, 0.70+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245024 HeapInuse: 598016 HeapObjects: 378 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 191 @0.295s 24%: 0.004+0+0.34 ms clock, 0.074+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 192 @0.295s 24%: 0.043+0+0.33 ms clock, 0.64+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 245216 HeapInuse: 598016 HeapObjects: 378 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 193 @0.296s 24%: 0.014+0+0.30 ms clock, 0.21+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 194 @0.297s 24%: 0.050+0+0.36 ms clock, 0.75+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244728 HeapInuse: 598016 HeapObjects: 371 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 195 @0.298s 24%: 0.002+0+0.34 ms clock, 0.036+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 196 @0.298s 24%: 0.046+0+0.41 ms clock, 0.70+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 244120 HeapInuse: 598016 HeapObjects: 365 HeapIdle 745472 HeapReleased 0 HeapSys 1343488
gc 197 @0.299s 24%: 0.002+0+0.27 ms clock, 0.045+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       850 ns/op	      58 B/op	       0 allocs/op
gc 198 @0.300s 24%: 0.007+0+0.27 ms clock, 0.11+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252032 HeapInuse: 606208 HeapObjects: 360 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 199 @0.301s 24%: 0.004+0+0.34 ms clock, 0.066+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E3-16                    	gc 200 @0.301s 24%: 0.005+0+0.29 ms clock, 0.077+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252704 HeapInuse: 606208 HeapObjects: 372 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 201 @0.303s 24%: 0.022+0+0.34 ms clock, 0.33+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 202 @0.303s 24%: 0.045+0+0.28 ms clock, 0.67+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252352 HeapInuse: 606208 HeapObjects: 367 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 203 @0.304s 24%: 0.002+0+0.38 ms clock, 0.038+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 204 @0.305s 24%: 0.046+0+0.26 ms clock, 0.69+0/0/0+3.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 251512 HeapInuse: 606208 HeapObjects: 357 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 205 @0.306s 24%: 0.004+0+0.29 ms clock, 0.067+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 206 @0.306s 24%: 0.043+0+0.45 ms clock, 0.65+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252192 HeapInuse: 606208 HeapObjects: 364 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 207 @0.308s 24%: 0.020+0+0.33 ms clock, 0.31+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 208 @0.308s 24%: 0.045+0+0.32 ms clock, 0.67+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252016 HeapInuse: 606208 HeapObjects: 361 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 209 @0.309s 24%: 0.016+0+0.28 ms clock, 0.25+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 210 @0.310s 24%: 0.004+0+0.39 ms clock, 0.073+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252688 HeapInuse: 606208 HeapObjects: 372 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 211 @0.311s 25%: 0.002+0+0.38 ms clock, 0.034+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 212 @0.311s 25%: 0.046+0+0.28 ms clock, 0.69+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 253080 HeapInuse: 606208 HeapObjects: 376 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 213 @0.312s 25%: 0.002+0+0.37 ms clock, 0.037+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 214 @0.313s 25%: 0.048+0+0.28 ms clock, 0.73+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252416 HeapInuse: 606208 HeapObjects: 368 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 215 @0.314s 25%: 0.002+0+0.34 ms clock, 0.038+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 216 @0.314s 25%: 0.045+0+0.33 ms clock, 0.68+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252744 HeapInuse: 606208 HeapObjects: 371 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 217 @0.316s 25%: 0.002+0+0.29 ms clock, 0.038+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 218 @0.316s 25%: 0.004+0+0.44 ms clock, 0.065+0/0/0+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252056 HeapInuse: 606208 HeapObjects: 359 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 219 @0.317s 25%: 0.004+0+0.30 ms clock, 0.070+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 220 @0.318s 25%: 0.042+0+0.35 ms clock, 0.63+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252792 HeapInuse: 606208 HeapObjects: 371 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 221 @0.319s 25%: 0.002+0+0.57 ms clock, 0.033+0/0/0+9.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 222 @0.320s 25%: 0.043+0+0.27 ms clock, 0.65+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 252792 HeapInuse: 606208 HeapObjects: 371 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 223 @0.321s 25%: 0.001+0+0.39 ms clock, 0.030+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 224 @0.321s 25%: 0.047+0+0.30 ms clock, 0.71+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 253504 HeapInuse: 606208 HeapObjects: 378 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 225 @0.322s 25%: 0.015+0+0.38 ms clock, 0.22+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 226 @0.323s 25%: 0.043+0+0.32 ms clock, 0.64+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 251984 HeapInuse: 606208 HeapObjects: 360 HeapIdle 737280 HeapReleased 0 HeapSys 1343488
gc 227 @0.324s 25%: 0.003+0+0.32 ms clock, 0.048+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       586 ns/op	      11 B/op	       0 allocs/op
gc 228 @0.325s 25%: 0.008+0+0.37 ms clock, 0.12+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194016 HeapInuse: 507904 HeapObjects: 295 HeapIdle 835584 HeapReleased 0 HeapSys 1343488
gc 229 @0.325s 25%: 0.004+0+0.36 ms clock, 0.068+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E3-16           	gc 230 @0.326s 25%: 0.014+0+0.30 ms clock, 0.22+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194520 HeapInuse: 516096 HeapObjects: 303 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 231 @0.327s 25%: 0.004+0+0.30 ms clock, 0.074+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 232 @0.327s 26%: 0.040+0+0.36 ms clock, 0.60+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194080 HeapInuse: 516096 HeapObjects: 296 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 233 @0.328s 26%: 0.005+0+0.36 ms clock, 0.080+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 234 @0.328s 26%: 0.004+0+1.3 ms clock, 0.064+0/0/0+20 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194312 HeapInuse: 516096 HeapObjects: 300 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 235 @0.330s 26%: 0.004+0+0.33 ms clock, 0.068+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 236 @0.331s 26%: 0.003+0+0.37 ms clock, 0.058+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194312 HeapInuse: 516096 HeapObjects: 300 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 237 @0.331s 26%: 0.017+0+0.47 ms clock, 0.25+0/0/0+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 238 @0.332s 26%: 0.042+0+0.33 ms clock, 0.63+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194384 HeapInuse: 516096 HeapObjects: 300 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 239 @0.333s 26%: 0.015+0+0.34 ms clock, 0.23+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 240 @0.333s 26%: 0.045+0+0.31 ms clock, 0.67+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194560 HeapInuse: 516096 HeapObjects: 305 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 241 @0.334s 26%: 0.019+0+0.38 ms clock, 0.28+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 242 @0.335s 26%: 0.045+0+0.32 ms clock, 0.68+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194328 HeapInuse: 516096 HeapObjects: 301 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 243 @0.335s 26%: 0.002+0+0.33 ms clock, 0.036+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 244 @0.336s 26%: 0.004+0+0.29 ms clock, 0.064+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194384 HeapInuse: 516096 HeapObjects: 300 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 245 @0.336s 26%: 0.015+0+0.37 ms clock, 0.23+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 246 @0.337s 27%: 0.041+0+0.41 ms clock, 0.62+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 193760 HeapInuse: 507904 HeapObjects: 293 HeapIdle 835584 HeapReleased 0 HeapSys 1343488
gc 247 @0.338s 27%: 0.004+0+0.35 ms clock, 0.073+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 248 @0.338s 27%: 0.004+0+0.29 ms clock, 0.068+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194104 HeapInuse: 507904 HeapObjects: 299 HeapIdle 835584 HeapReleased 0 HeapSys 1343488
gc 249 @0.339s 27%: 0.004+0+0.31 ms clock, 0.076+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 250 @0.339s 27%: 0.047+0+0.35 ms clock, 0.70+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194216 HeapInuse: 516096 HeapObjects: 299 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 251 @0.340s 27%: 0.002+0+0.28 ms clock, 0.042+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 252 @0.341s 27%: 0.047+0+0.28 ms clock, 0.71+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194520 HeapInuse: 516096 HeapObjects: 303 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 253 @0.341s 27%: 0.014+0+0.33 ms clock, 0.22+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 254 @0.342s 27%: 0.045+0+0.39 ms clock, 0.68+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194504 HeapInuse: 516096 HeapObjects: 302 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 255 @0.343s 27%: 0.002+0+0.33 ms clock, 0.036+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 256 @0.343s 27%: 0.047+0+0.33 ms clock, 0.71+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194352 HeapInuse: 516096 HeapObjects: 302 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 257 @0.344s 27%: 0.004+0+0.35 ms clock, 0.069+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 258 @0.344s 27%: 0.004+0+0.33 ms clock, 0.071+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 194672 HeapInuse: 516096 HeapObjects: 307 HeapIdle 827392 HeapReleased 0 HeapSys 1343488
gc 259 @0.345s 27%: 0.004+0+0.50 ms clock, 0.072+0/0/0+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       876 ns/op	       8 B/op	       0 allocs/op
gc 260 @0.346s 27%: 0.006+0+0.31 ms clock, 0.098+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214640 HeapInuse: 491520 HeapObjects: 271 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 261 @0.346s 27%: 0.002+0+0.29 ms clock, 0.035+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E3-16    	gc 262 @0.347s 27%: 0.006+0+0.37 ms clock, 0.092+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 263 @0.348s 27%: 0.004+0+0.29 ms clock, 0.071+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 264 @0.348s 27%: 0.043+0+0.35 ms clock, 0.65+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 265 @0.349s 27%: 0.002+0+0.29 ms clock, 0.041+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 266 @0.349s 28%: 0.044+0+0.27 ms clock, 0.67+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 267 @0.350s 28%: 0.011+0+0.32 ms clock, 0.17+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 268 @0.350s 28%: 0.004+0+0.27 ms clock, 0.066+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 269 @0.351s 28%: 0.004+0+0.42 ms clock, 0.071+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 270 @0.351s 28%: 0.043+0+0.42 ms clock, 0.65+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 271 @0.352s 28%: 0.002+0+0.29 ms clock, 0.034+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 272 @0.353s 28%: 0.046+0+0.36 ms clock, 0.69+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 273 @0.353s 28%: 0.017+0+0.26 ms clock, 0.25+0/0/0+4.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 274 @0.354s 28%: 0.003+0+0.27 ms clock, 0.056+0/0/0+4.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 275 @0.354s 28%: 0.002+0+0.28 ms clock, 0.037+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 276 @0.355s 28%: 0.004+0+0.26 ms clock, 0.070+0/0/0+4.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 277 @0.355s 28%: 0.004+0+0.27 ms clock, 0.070+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 278 @0.356s 28%: 0.003+0+0.31 ms clock, 0.059+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 279 @0.356s 28%: 0.002+0+0.25 ms clock, 0.040+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 280 @0.357s 28%: 0.004+0+0.26 ms clock, 0.061+0/0/0+4.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 281 @0.357s 28%: 0.002+0+0.33 ms clock, 0.035+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 282 @0.358s 28%: 0.044+0+0.34 ms clock, 0.66+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 283 @0.359s 28%: 0.015+0+0.34 ms clock, 0.23+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 284 @0.359s 28%: 0.027+0+0.31 ms clock, 0.41+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 285 @0.360s 28%: 0.011+0+0.30 ms clock, 0.17+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 286 @0.360s 28%: 0.047+0+0.34 ms clock, 0.71+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 287 @0.361s 28%: 0.018+0+0.33 ms clock, 0.27+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 288 @0.361s 29%: 0.031+0+0.41 ms clock, 0.46+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 289 @0.362s 29%: 0.002+0+0.43 ms clock, 0.036+0/0/0+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 290 @0.363s 29%: 0.043+0+0.31 ms clock, 0.65+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 214720 HeapInuse: 491520 HeapObjects: 274 HeapIdle 851968 HeapReleased 0 HeapSys 1343488
gc 291 @0.363s 29%: 0.017+0+0.33 ms clock, 0.25+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
    1000	       640 ns/op	      28 B/op	       0 allocs/op
gc 292 @0.364s 29%: 0.013+0+0.37 ms clock, 0.20+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271120 HeapInuse: 548864 HeapObjects: 271 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 293 @0.367s 29%: 0.002+0+0.36 ms clock, 0.035+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSort1E4-16                               	gc 294 @0.367s 29%: 0.004+0+0.31 ms clock, 0.070+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 295 @0.370s 28%: 0.002+0+0.31 ms clock, 0.035+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 296 @0.371s 29%: 0.044+0+0.31 ms clock, 0.66+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 297 @0.374s 28%: 0.005+0+0.31 ms clock, 0.081+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 298 @0.374s 28%: 0.11+0+0.33 ms clock, 1.7+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 299 @0.377s 28%: 0.006+0+0.38 ms clock, 0.090+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 300 @0.377s 28%: 0.046+0+0.37 ms clock, 0.69+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 301 @0.380s 28%: 0.002+0+0.30 ms clock, 0.035+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 302 @0.381s 28%: 0.041+0+0.32 ms clock, 0.62+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 303 @0.384s 28%: 0.002+0+0.34 ms clock, 0.038+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 304 @0.384s 28%: 0.004+0+0.26 ms clock, 0.062+0/0/0+3.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 305 @0.387s 28%: 0.002+0+0.35 ms clock, 0.035+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 306 @0.387s 28%: 0.045+0+0.36 ms clock, 0.68+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 307 @0.390s 28%: 0.002+0+0.34 ms clock, 0.035+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 308 @0.391s 28%: 0.045+0+0.27 ms clock, 0.68+0/0/0+4.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 309 @0.394s 28%: 0.002+0+0.38 ms clock, 0.041+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 310 @0.394s 28%: 0.083+0+0.42 ms clock, 1.2+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 311 @0.397s 28%: 0.002+0+0.32 ms clock, 0.038+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 312 @0.398s 28%: 0.013+0+0.43 ms clock, 0.19+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 313 @0.400s 28%: 0.002+0+0.33 ms clock, 0.047+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 314 @0.401s 28%: 0.007+0+0.31 ms clock, 0.10+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 315 @0.404s 28%: 0.002+0+0.36 ms clock, 0.035+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 316 @0.404s 28%: 0.004+0+0.39 ms clock, 0.069+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 317 @0.407s 28%: 0.002+0+0.37 ms clock, 0.036+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 318 @0.408s 28%: 0.062+0+0.34 ms clock, 0.93+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 271184 HeapInuse: 548864 HeapObjects: 274 HeapIdle 794624 HeapReleased 0 HeapSys 1343488
gc 319 @0.410s 28%: 0.018+0+0.37 ms clock, 0.27+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       274 ns/op	       8 B/op	       0 allocs/op
gc 320 @0.411s 28%: 0.047+0+0.29 ms clock, 0.70+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069440 HeapInuse: 1327104 HeapObjects: 15273 HeapIdle 16384 HeapReleased 0 HeapSys 1343488
gc 321 @0.444s 26%: 0.004+0+0.40 ms clock, 0.064+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSetInsert1E4-16                          	gc 322 @0.445s 26%: 0.005+0+0.56 ms clock, 0.075+0/0/0+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 323 @0.479s 24%: 0.047+0+0.73 ms clock, 0.70+0/0/0+11 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 324 @0.480s 24%: 0.076+0+0.39 ms clock, 1.1+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 325 @0.513s 23%: 0.004+0+0.50 ms clock, 0.069+0/0/0+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 326 @0.514s 23%: 0.044+0+0.30 ms clock, 0.66+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 327 @0.547s 21%: 0.004+0+0.50 ms clock, 0.068+0/0/0+8.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 328 @0.548s 21%: 0.015+0+0.39 ms clock, 0.23+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 329 @0.581s 20%: 0.013+0+0.42 ms clock, 0.20+0/0/0+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 330 @0.581s 20%: 0.043+0+0.30 ms clock, 0.64+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 331 @0.616s 19%: 0.022+0+0.46 ms clock, 0.34+0/0/0+7.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 332 @0.616s 19%: 0.038+0+0.35 ms clock, 0.58+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 333 @0.649s 18%: 0.016+0+0.39 ms clock, 0.24+0/0/0+5.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 334 @0.650s 18%: 0.040+0+0.38 ms clock, 0.61+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 335 @0.683s 18%: 0.040+0+0.44 ms clock, 0.61+0/0/0+6.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 336 @0.683s 18%: 0.038+0+0.32 ms clock, 0.58+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 337 @0.717s 17%: 0.002+0+0.37 ms clock, 0.040+0/0/0+5.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 338 @0.717s 17%: 0.043+0+0.41 ms clock, 0.65+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 339 @0.752s 16%: 0.034+0+0.41 ms clock, 0.52+0/0/0+6.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 340 @0.752s 16%: 0.035+0+0.41 ms clock, 0.53+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 341 @0.786s 15%: 0.004+0+0.48 ms clock, 0.074+0/0/0+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 342 @0.786s 16%: 0.070+0+0.40 ms clock, 1.0+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 343 @0.819s 15%: 0.034+0+0.51 ms clock, 0.52+0/0/0+7.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 344 @0.820s 15%: 0.005+0+0.37 ms clock, 0.083+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 345 @0.853s 14%: 0.028+0+0.45 ms clock, 0.42+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 346 @0.854s 14%: 0.072+0+0.40 ms clock, 1.0+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 347 @0.888s 14%: 0.042+0+0.44 ms clock, 0.64+0/0/0+6.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 348 @0.888s 14%: 0.084+0+0.36 ms clock, 1.2+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 349 @0.919s 14%: 0.014+0+0.45 ms clock, 0.21+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 350 @0.919s 14%: 0.068+0+0.40 ms clock, 1.0+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 351 @0.950s 13%: 0.038+0+0.40 ms clock, 0.57+0/0/0+6.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 352 @0.951s 13%: 0.049+0+0.38 ms clock, 0.73+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1069504 HeapInuse: 1327104 HeapObjects: 15276 HeapIdle 1064960 HeapReleased 0 HeapSys 2392064
gc 353 @0.981s 13%: 0.016+0+0.45 ms clock, 0.25+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
   10000	      3062 ns/op	      88 B/op	       2 allocs/op
gc 354 @0.982s 13%: 0.007+0+0.39 ms clock, 0.11+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031360 HeapInuse: 2293760 HeapObjects: 30274 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 355 @1.034s 12%: 0.004+0+0.46 ms clock, 0.076+0/0/0+7.4 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
#BenchmarkSetErase1E4-16                           	gc 356 @1.035s 12%: 0.005+0+0.42 ms clock, 0.087+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 357 @1.084s 12%: 0.041+0+0.46 ms clock, 0.62+0/0/0+7.0 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 358 @1.085s 12%: 0.075+0+0.40 ms clock, 1.1+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 359 @1.135s 11%: 0.004+0+0.52 ms clock, 0.068+0/0/0+8.4 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 360 @1.136s 11%: 0.046+0+0.39 ms clock, 0.69+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 361 @1.185s 11%: 0.012+0+0.47 ms clock, 0.19+0/0/0+7.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 362 @1.185s 11%: 0.067+0+0.36 ms clock, 1.0+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 363 @1.235s 10%: 0.003+0+0.58 ms clock, 0.056+0/0/0+8.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 364 @1.235s 10%: 0.089+0+0.35 ms clock, 1.3+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 365 @1.284s 10%: 0.002+0+0.62 ms clock, 0.035+0/0/0+9.9 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 366 @1.285s 10%: 0.045+0+0.53 ms clock, 0.68+0/0/0+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 367 @1.333s 10%: 0.040+0+0.52 ms clock, 0.61+0/0/0+7.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 368 @1.334s 10%: 0.046+0+0.43 ms clock, 0.69+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 369 @1.383s 10%: 0.004+0+0.52 ms clock, 0.068+0/0/0+8.4 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 370 @1.383s 10%: 0.004+0+0.53 ms clock, 0.069+0/0/0+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 371 @1.436s 9%: 0.004+0+0.47 ms clock, 0.071+0/0/0+7.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 372 @1.436s 9%: 0.044+0+0.33 ms clock, 0.66+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 373 @1.489s 9%: 0.004+0+0.47 ms clock, 0.067+0/0/0+7.5 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 374 @1.490s 9%: 0.080+0+0.29 ms clock, 1.2+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 375 @1.542s 9%: 0.004+0+0.51 ms clock, 0.068+0/0/0+8.2 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 376 @1.543s 9%: 0.022+0+0.45 ms clock, 0.33+0/0/0+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 377 @1.595s 8%: 0.004+0+0.42 ms clock, 0.062+0/0/0+6.3 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 378 @1.596s 8%: 0.045+0+0.39 ms clock, 0.68+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 379 @1.648s 8%: 0.002+0+0.46 ms clock, 0.039+0/0/0+7.3 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 380 @1.649s 8%: 0.040+0+0.36 ms clock, 0.60+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 381 @1.701s 8%: 0.002+0+0.50 ms clock, 0.037+0/0/0+8.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 382 @1.702s 8%: 0.053+0+0.39 ms clock, 0.80+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 383 @1.754s 8%: 0.004+0+0.43 ms clock, 0.069+0/0/0+6.9 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 384 @1.755s 8%: 0.045+0+0.35 ms clock, 0.67+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 2031424 HeapInuse: 2293760 HeapObjects: 30277 HeapIdle 98304 HeapReleased 0 HeapSys 2392064
gc 385 @1.808s 8%: 0.002+0+0.50 ms clock, 0.042+0/0/0+8.1 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
   10000	      2053 ns/op	      96 B/op	       2 allocs/op
gc 386 @1.808s 8%: 0.040+0+0.37 ms clock, 0.60+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151360 HeapInuse: 1409024 HeapObjects: 15274 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 387 @1.830s 7%: 0.004+0+0.44 ms clock, 0.069+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E4-16                  	gc 388 @1.830s 8%: 0.007+0+0.37 ms clock, 0.11+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 389 @1.851s 7%: 0.044+0+0.43 ms clock, 0.67+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 390 @1.852s 7%: 0.051+0+0.35 ms clock, 0.76+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 391 @1.873s 7%: 0.037+0+0.52 ms clock, 0.56+0/0/0+7.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 392 @1.874s 7%: 0.033+0+0.39 ms clock, 0.50+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 393 @1.891s 7%: 0.001+0+0.38 ms clock, 0.023+0/0/0+6.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 394 @1.892s 7%: 0.014+0+0.34 ms clock, 0.21+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 395 @1.905s 7%: 0.025+0+0.39 ms clock, 0.37+0/0/0+5.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 396 @1.905s 7%: 0.062+0+0.35 ms clock, 0.94+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 397 @1.918s 7%: 0.055+0+0.41 ms clock, 0.82+0/0/0+6.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 398 @1.919s 7%: 0.008+0+0.34 ms clock, 0.12+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 399 @1.932s 7%: 0.026+0+0.42 ms clock, 0.39+0/0/0+6.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 400 @1.933s 7%: 0.068+0+0.36 ms clock, 1.0+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 401 @1.945s 7%: 0.004+0+0.44 ms clock, 0.061+0/0/0+6.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 402 @1.946s 7%: 0.064+0+0.36 ms clock, 0.96+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 403 @1.959s 7%: 0.038+0+0.42 ms clock, 0.57+0/0/0+6.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 404 @1.959s 7%: 0.040+0+0.33 ms clock, 0.60+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 405 @1.972s 7%: 0.058+0+0.39 ms clock, 0.87+0/0/0+5.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 406 @1.973s 7%: 0.041+0+0.37 ms clock, 0.62+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 407 @1.986s 7%: 0.003+0+0.42 ms clock, 0.057+0/0/0+6.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 408 @1.987s 7%: 0.041+0+0.41 ms clock, 0.62+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 409 @2.001s 7%: 0.001+0+0.31 ms clock, 0.023+0/0/0+4.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 410 @2.002s 7%: 0.003+0+0.32 ms clock, 0.052+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 411 @2.015s 7%: 0.003+0+0.46 ms clock, 0.056+0/0/0+7.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 412 @2.016s 7%: 0.003+0+0.28 ms clock, 0.047+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 413 @2.028s 7%: 0.054+0+0.40 ms clock, 0.82+0/0/0+6.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 414 @2.029s 7%: 0.039+0+0.33 ms clock, 0.58+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 415 @2.042s 7%: 0.055+0+0.46 ms clock, 0.83+0/0/0+7.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 416 @2.043s 7%: 0.026+0+0.34 ms clock, 0.40+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1151424 HeapInuse: 1409024 HeapObjects: 15277 HeapIdle 983040 HeapReleased 0 HeapSys 2392064
gc 417 @2.056s 7%: 0.040+0+0.43 ms clock, 0.61+0/0/0+6.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
   10000	      1291 ns/op	      96 B/op	       2 allocs/op
gc 418 @2.056s 7%: 0.033+0+0.39 ms clock, 0.50+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754328 HeapInuse: 1032192 HeapObjects: 10280 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 419 @2.072s 7%: 0.056+0+0.44 ms clock, 0.84+0/0/0+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E4-16          	gc 420 @2.073s 7%: 0.011+0+0.34 ms clock, 0.17+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 421 @2.089s 7%: 0.041+0+0.36 ms clock, 0.61+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 422 @2.089s 7%: 0.046+0+0.42 ms clock, 0.70+0/0/0+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 423 @2.106s 7%: 0.003+0+0.41 ms clock, 0.062+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 424 @2.106s 7%: 0.049+0+0.42 ms clock, 0.73+0/0/0+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 425 @2.122s 7%: 0.060+0+0.40 ms clock, 0.90+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 426 @2.123s 7%: 0.012+0+0.36 ms clock, 0.18+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 427 @2.139s 7%: 0.023+0+0.39 ms clock, 0.34+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 428 @2.139s 7%: 0.037+0+0.35 ms clock, 0.56+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 429 @2.155s 7%: 0.016+0+0.36 ms clock, 0.25+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 430 @2.156s 7%: 0.030+0+0.34 ms clock, 0.45+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 431 @2.172s 7%: 0.032+0+0.36 ms clock, 0.49+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 432 @2.172s 7%: 0.010+0+0.34 ms clock, 0.15+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 433 @2.188s 7%: 0.041+0+0.47 ms clock, 0.62+0/0/0+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 434 @2.188s 7%: 0.046+0+0.32 ms clock, 0.70+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 435 @2.205s 7%: 0.004+0+0.40 ms clock, 0.069+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 436 @2.206s 7%: 0.050+0+0.45 ms clock, 0.75+0/0/0+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 437 @2.222s 7%: 0.001+0+0.36 ms clock, 0.024+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 438 @2.222s 7%: 0.008+0+0.59 ms clock, 0.12+0/0/0+8.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 439 @2.239s 7%: 0.057+0+0.36 ms clock, 0.85+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 440 @2.239s 7%: 0.10+0+0.33 ms clock, 1.5+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 441 @2.256s 7%: 0.093+0+0.35 ms clock, 1.3+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 442 @2.256s 7%: 0.037+0+0.35 ms clock, 0.55+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 443 @2.272s 7%: 0.037+0+0.35 ms clock, 0.56+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 444 @2.272s 7%: 0.051+0+0.36 ms clock, 0.77+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 445 @2.288s 7%: 0.025+0+0.39 ms clock, 0.38+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 446 @2.289s 7%: 0.052+0+0.38 ms clock, 0.79+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 447 @2.305s 7%: 0.026+0+0.39 ms clock, 0.39+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 448 @2.306s 7%: 0.070+0+0.32 ms clock, 1.0+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 754408 HeapInuse: 1032192 HeapObjects: 10283 HeapIdle 1359872 HeapReleased 0 HeapSys 2392064
gc 449 @2.322s 7%: 0.001+0+0.35 ms clock, 0.029+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	      1596 ns/op	      56 B/op	       1 allocs/op
gc 450 @2.322s 7%: 0.003+0+0.36 ms clock, 0.047+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 621264 HeapInuse: 991232 HeapObjects: 898 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 451 @2.325s 7%: 0.028+0+0.36 ms clock, 0.42+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E4-16                   	gc 452 @2.325s 7%: 0.019+0+0.34 ms clock, 0.29+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619568 HeapInuse: 991232 HeapObjects: 887 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 453 @2.328s 7%: 0.030+0+0.37 ms clock, 0.45+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 454 @2.328s 7%: 0.009+0+0.34 ms clock, 0.14+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 620688 HeapInuse: 991232 HeapObjects: 900 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 455 @2.331s 7%: 0.030+0+0.34 ms clock, 0.46+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 456 @2.331s 7%: 0.045+0+0.37 ms clock, 0.67+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 617888 HeapInuse: 991232 HeapObjects: 871 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 457 @2.334s 7%: 0.063+0+0.37 ms clock, 0.95+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 458 @2.334s 7%: 0.063+0+0.33 ms clock, 0.94+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 624472 HeapInuse: 991232 HeapObjects: 918 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 459 @2.337s 7%: 0.003+0+0.44 ms clock, 0.057+0/0/0+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 460 @2.337s 7%: 0.002+0+0.33 ms clock, 0.044+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619872 HeapInuse: 991232 HeapObjects: 890 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 461 @2.340s 7%: 0.044+0+0.33 ms clock, 0.66+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 462 @2.341s 7%: 0.049+0+0.33 ms clock, 0.73+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 622448 HeapInuse: 991232 HeapObjects: 918 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 463 @2.343s 7%: 0.019+0+0.31 ms clock, 0.29+0/0/0+4.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 464 @2.344s 7%: 0.057+0+0.33 ms clock, 0.85+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 622072 HeapInuse: 991232 HeapObjects: 909 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 465 @2.346s 7%: 0.001+0+0.33 ms clock, 0.018+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 466 @2.347s 7%: 0.045+0+0.33 ms clock, 0.68+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619968 HeapInuse: 991232 HeapObjects: 891 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 467 @2.349s 7%: 0.034+0+0.33 ms clock, 0.51+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 468 @2.350s 7%: 0.010+0+0.36 ms clock, 0.15+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 621384 HeapInuse: 991232 HeapObjects: 909 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 469 @2.352s 7%: 0.020+0+0.34 ms clock, 0.31+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 470 @2.353s 7%: 0.030+0+0.34 ms clock, 0.46+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 623624 HeapInuse: 991232 HeapObjects: 919 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 471 @2.355s 7%: 0.004+0+0.35 ms clock, 0.067+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 472 @2.356s 7%: 0.052+0+0.32 ms clock, 0.78+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 618928 HeapInuse: 991232 HeapObjects: 881 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 473 @2.358s 7%: 0.034+0+0.34 ms clock, 0.51+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 474 @2.359s 7%: 0.003+0+0.32 ms clock, 0.053+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 619856 HeapInuse: 991232 HeapObjects: 890 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 475 @2.361s 7%: 0.018+0+0.28 ms clock, 0.28+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 476 @2.362s 7%: 0.003+0+0.37 ms clock, 0.058+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 622216 HeapInuse: 991232 HeapObjects: 913 HeapIdle 1400832 HeapReleased 0 HeapSys 2392064
gc 477 @2.365s 7%: 0.029+0+0.40 ms clock, 0.43+0/0/0+6.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       284 ns/op	      43 B/op	       0 allocs/op
gc 478 @2.365s 7%: 0.028+0+0.39 ms clock, 0.43+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 705296 HeapInuse: 1073152 HeapObjects: 901 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 479 @2.369s 7%: 0.017+0+0.33 ms clock, 0.25+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E4-16                    	gc 480 @2.369s 7%: 0.003+0+0.34 ms clock, 0.048+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 700504 HeapInuse: 1073152 HeapObjects: 876 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 481 @2.372s 7%: 0.030+0+0.39 ms clock, 0.45+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 482 @2.373s 7%: 0.079+0+0.38 ms clock, 1.1+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 705408 HeapInuse: 1073152 HeapObjects: 908 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 483 @2.376s 7%: 0.024+0+0.30 ms clock, 0.37+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 484 @2.377s 7%: 0.004+0+0.31 ms clock, 0.062+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 702944 HeapInuse: 1073152 HeapObjects: 903 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 485 @2.380s 7%: 0.047+0+0.38 ms clock, 0.70+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 486 @2.381s 7%: 0.044+0+0.41 ms clock, 0.66+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 707240 HeapInuse: 1073152 HeapObjects: 929 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 487 @2.384s 7%: 0.024+0+0.31 ms clock, 0.36+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 488 @2.385s 7%: 0.064+0+0.36 ms clock, 0.96+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704560 HeapInuse: 1073152 HeapObjects: 920 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 489 @2.389s 7%: 0.030+0+0.41 ms clock, 0.45+0/0/0+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 490 @2.389s 7%: 0.047+0+0.35 ms clock, 0.70+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704280 HeapInuse: 1073152 HeapObjects: 919 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 491 @2.393s 7%: 0.025+0+0.36 ms clock, 0.38+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 492 @2.394s 7%: 0.043+0+0.38 ms clock, 0.65+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 700176 HeapInuse: 1073152 HeapObjects: 871 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 493 @2.398s 7%: 0.018+0+0.32 ms clock, 0.28+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 494 @2.398s 7%: 0.003+0+0.36 ms clock, 0.054+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704152 HeapInuse: 1073152 HeapObjects: 892 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 495 @2.402s 7%: 0.035+0+0.33 ms clock, 0.53+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 496 @2.403s 7%: 0.044+0+0.39 ms clock, 0.66+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 704896 HeapInuse: 1073152 HeapObjects: 902 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 497 @2.406s 7%: 0.022+0+0.32 ms clock, 0.33+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 498 @2.407s 7%: 0.056+0+0.32 ms clock, 0.84+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 701328 HeapInuse: 1073152 HeapObjects: 887 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 499 @2.411s 7%: 0.018+0+0.35 ms clock, 0.27+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 500 @2.411s 7%: 0.042+0+0.34 ms clock, 0.64+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 699936 HeapInuse: 1073152 HeapObjects: 871 HeapIdle 1318912 HeapReleased 0 HeapSys 2392064
gc 501 @2.415s 7%: 0.023+0+0.34 ms clock, 0.35+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       124 ns/op	       8 B/op	       0 allocs/op
gc 502 @2.416s 7%: 0.005+0+0.64 ms clock, 0.084+0/0/0+9.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254696 HeapInuse: 614400 HeapObjects: 387 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 503 @2.418s 7%: 0.022+0+0.34 ms clock, 0.34+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E4-16           	gc 504 @2.418s 7%: 0.010+0+0.34 ms clock, 0.15+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255192 HeapInuse: 614400 HeapObjects: 395 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 505 @2.420s 7%: 0.030+0+0.80 ms clock, 0.46+0/0/0+12 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 506 @2.421s 7%: 0.004+0+0.24 ms clock, 0.063+0/0/0+3.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 253696 HeapInuse: 614400 HeapObjects: 377 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 507 @2.422s 7%: 0.001+0+0.35 ms clock, 0.029+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 508 @2.423s 7%: 0.022+0+0.35 ms clock, 0.34+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254368 HeapInuse: 614400 HeapObjects: 384 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 509 @2.424s 7%: 0.003+0+0.35 ms clock, 0.059+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 510 @2.425s 7%: 0.041+0+0.30 ms clock, 0.61+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254328 HeapInuse: 614400 HeapObjects: 382 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 511 @2.426s 7%: 0.022+0+0.31 ms clock, 0.34+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 512 @2.427s 8%: 0.003+0+0.37 ms clock, 0.056+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255192 HeapInuse: 614400 HeapObjects: 395 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 513 @2.428s 8%: 0.021+0+0.36 ms clock, 0.32+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 514 @2.429s 8%: 0.065+0+0.49 ms clock, 0.98+0/0/0+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 254000 HeapInuse: 614400 HeapObjects: 385 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 515 @2.431s 8%: 0.019+0+0.33 ms clock, 0.29+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 516 @2.431s 8%: 0.009+0+0.34 ms clock, 0.13+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255384 HeapInuse: 614400 HeapObjects: 397 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 517 @2.433s 8%: 0.026+0+0.29 ms clock, 0.39+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 518 @2.433s 8%: 0.040+0+0.30 ms clock, 0.60+0/0/0+4.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255744 HeapInuse: 614400 HeapObjects: 402 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 519 @2.435s 8%: 0.001+0+0.41 ms clock, 0.027+0/0/0+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 520 @2.435s 8%: 0.003+0+0.33 ms clock, 0.049+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255032 HeapInuse: 614400 HeapObjects: 394 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 521 @2.437s 8%: 0.001+0+0.31 ms clock, 0.025+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 522 @2.437s 8%: 0.043+0+0.35 ms clock, 0.65+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255544 HeapInuse: 614400 HeapObjects: 399 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 523 @2.439s 8%: 0.003+0+0.36 ms clock, 0.060+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 524 @2.440s 8%: 0.043+0+0.38 ms clock, 0.65+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255648 HeapInuse: 614400 HeapObjects: 402 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 525 @2.441s 8%: 0.018+0+0.36 ms clock, 0.27+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 526 @2.442s 8%: 0.003+0+0.34 ms clock, 0.056+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 255096 HeapInuse: 614400 HeapObjects: 394 HeapIdle 1777664 HeapReleased 0 HeapSys 2392064
gc 527 @2.443s 8%: 0.015+0+0.35 ms clock, 0.23+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       159 ns/op	       6 B/op	       0 allocs/op
gc 528 @2.444s 8%: 0.003+0+0.37 ms clock, 0.057+0/0/0+5.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377584 HeapInuse: 663552 HeapObjects: 271 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 529 @2.445s 8%: 0.003+0+0.30 ms clock, 0.062+0/0/0+4.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E4-16    	gc 530 @2.446s 8%: 0.010+0+0.34 ms clock, 0.15+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 531 @2.447s 8%: 0.001+0+0.32 ms clock, 0.030+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 532 @2.448s 8%: 0.004+0+0.54 ms clock, 0.064+0/0/0+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 533 @2.449s 8%: 0.001+0+0.26 ms clock, 0.029+0/0/0+4.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 534 @2.450s 8%: 0.004+0+0.24 ms clock, 0.067+0/0/0+3.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 535 @2.451s 8%: 0.002+0+0.33 ms clock, 0.032+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 536 @2.451s 8%: 0.003+0+0.25 ms clock, 0.045+0/0/0+3.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 537 @2.452s 8%: 0.001+0+0.27 ms clock, 0.031+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 538 @2.453s 8%: 0.011+0+0.35 ms clock, 0.16+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 539 @2.454s 8%: 0.030+0+0.36 ms clock, 0.45+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 540 @2.455s 8%: 0.004+0+0.33 ms clock, 0.061+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 541 @2.456s 8%: 0.001+0+0.50 ms clock, 0.021+0/0/0+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 542 @2.457s 8%: 0.046+0+0.35 ms clock, 0.69+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 543 @2.458s 8%: 0.034+0+0.39 ms clock, 0.51+0/0/0+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 544 @2.459s 8%: 0.028+0+0.29 ms clock, 0.42+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 545 @2.460s 8%: 0.019+0+0.29 ms clock, 0.29+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 546 @2.460s 8%: 0.047+0+0.29 ms clock, 0.70+0/0/0+4.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 547 @2.462s 8%: 0.001+0+0.32 ms clock, 0.029+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 548 @2.462s 8%: 0.004+0+0.31 ms clock, 0.073+0/0/0+4.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 549 @2.463s 8%: 0.020+0+0.34 ms clock, 0.30+0/0/0+5.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 550 @2.464s 8%: 0.046+0+0.37 ms clock, 0.70+0/0/0+5.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 551 @2.465s 8%: 0.001+0+0.33 ms clock, 0.026+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 552 @2.466s 8%: 0.017+0+0.33 ms clock, 0.26+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 377664 HeapInuse: 663552 HeapObjects: 274 HeapIdle 1728512 HeapReleased 0 HeapSys 2392064
gc 553 @2.467s 8%: 0.018+0+0.32 ms clock, 0.28+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
   10000	       130 ns/op	      19 B/op	       0 allocs/op
gc 554 @2.468s 8%: 0.006+0+0.33 ms clock, 0.091+0/0/0+5.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992016 HeapInuse: 1277952 HeapObjects: 271 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 555 @2.494s 8%: 0.030+0+0.38 ms clock, 0.45+0/0/0+5.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSort1E5-16                               	gc 556 @2.495s 8%: 0.016+0+0.46 ms clock, 0.24+0/0/0+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 557 @2.520s 8%: 0.039+0+0.36 ms clock, 0.58+0/0/0+5.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 558 @2.521s 8%: 0.040+0+0.32 ms clock, 0.61+0/0/0+4.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 559 @2.547s 8%: 0.018+0+0.36 ms clock, 0.27+0/0/0+5.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 560 @2.547s 8%: 0.053+0+0.36 ms clock, 0.79+0/0/0+5.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 561 @2.573s 8%: 0.029+0+0.41 ms clock, 0.44+0/0/0+6.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 562 @2.574s 8%: 0.046+0+0.38 ms clock, 0.70+0/0/0+5.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 563 @2.600s 8%: 0.047+0+0.40 ms clock, 0.70+0/0/0+6.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 564 @2.600s 8%: 0.004+0+0.35 ms clock, 0.064+0/0/0+5.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 565 @2.626s 8%: 0.036+0+0.33 ms clock, 0.54+0/0/0+5.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 566 @2.627s 8%: 0.057+0+0.39 ms clock, 0.85+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 567 @2.653s 8%: 0.016+0+0.35 ms clock, 0.25+0/0/0+5.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 568 @2.653s 8%: 0.012+0+0.34 ms clock, 0.19+0/0/0+5.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 569 @2.679s 8%: 0.042+0+0.29 ms clock, 0.63+0/0/0+4.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 570 @2.680s 8%: 0.046+0+0.29 ms clock, 0.69+0/0/0+4.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 571 @2.705s 7%: 0.037+0+0.38 ms clock, 0.56+0/0/0+5.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 572 @2.706s 7%: 0.046+0+0.41 ms clock, 0.70+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 573 @2.732s 7%: 0.058+0+0.34 ms clock, 0.88+0/0/0+5.1 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 574 @2.732s 7%: 0.003+0+0.39 ms clock, 0.051+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 575 @2.758s 7%: 0.019+0+0.44 ms clock, 0.28+0/0/0+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 576 @2.759s 7%: 0.039+0+0.39 ms clock, 0.59+0/0/0+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 577 @2.785s 7%: 0.051+0+0.32 ms clock, 0.76+0/0/0+4.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 578 @2.785s 7%: 0.066+0+0.41 ms clock, 0.99+0/0/0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 579 @2.811s 7%: 0.022+0+0.36 ms clock, 0.33+0/0/0+5.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 580 @2.811s 7%: 0.003+0+0.70 ms clock, 0.051+0/0/0+10 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 992080 HeapInuse: 1277952 HeapObjects: 274 HeapIdle 1114112 HeapReleased 0 HeapSys 2392064
gc 581 @2.838s 7%: 0.011+0+0.35 ms clock, 0.17+0/0/0+5.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
  100000	       257 ns/op	       8 B/op	       0 allocs/op
gc 582 @2.838s 7%: 0.004+0+0.40 ms clock, 0.076+0/0/0+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 583 @2.960s 7%: 0.036+3.2+0.30 ms clock, 0.55+0.10/10/17+4.5 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 584 @3.084s 7%: 0.016+4.6+0.30 ms clock, 0.25+0.11/17/24+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996768 HeapInuse: 9289728 HeapObjects: 150307 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 585 @3.105s 7%: 0.001+0+0.90 ms clock, 0.030+0.11/17/24+14 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
#BenchmarkSetInsert1E5-16                          	gc 586 @3.106s 7%: 0.041+0+0.39 ms clock, 0.61+0.11/17/24+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 587 @3.224s 7%: 0.014+3.8+0.33 ms clock, 0.23+0.11/11/15+5.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 588 @3.396s 6%: 0.021+6.0+0.28 ms clock, 0.34+0.13/22/27+4.5 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996912 HeapInuse: 9289728 HeapObjects: 150312 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 589 @3.420s 6%: 0.004+0+0.98 ms clock, 0.076+0.13/22/27+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 590 @3.421s 6%: 0.080+0+0.57 ms clock, 1.2+0.13/22/27+8.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 591 @3.588s 6%: 0.019+3.9+0.32 ms clock, 0.30+0.14/12/15+5.1 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 592 @3.764s 6%: 0.019+5.7+0.29 ms clock, 0.31+0.15/21/47+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996928 HeapInuse: 9289728 HeapObjects: 150313 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 593 @3.789s 6%: 0.003+0+1.8 ms clock, 0.061+0.15/21/47+29 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 594 @3.791s 6%: 0.037+0+0.45 ms clock, 0.56+0.15/21/47+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 595 @3.958s 6%: 0.020+5.3+0.29 ms clock, 0.32+0.19/13/14+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 596 @4.132s 5%: 0.020+6.9+0.30 ms clock, 0.32+0.20/25/40+4.8 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996928 HeapInuse: 9289728 HeapObjects: 150313 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 597 @4.156s 5%: 0.002+0+1.0 ms clock, 0.040+0.20/25/40+16 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 598 @4.158s 5%: 0.041+0+0.41 ms clock, 0.62+0.20/25/40+6.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 599 @4.329s 5%: 0.023+5.1+0.33 ms clock, 0.37+0.14/14/18+5.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 600 @4.502s 5%: 0.019+8.0+0.31 ms clock, 0.30+0.18/29/25+5.0 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996896 HeapInuse: 9289728 HeapObjects: 150311 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 601 @4.526s 5%: 0.004+0+1.0 ms clock, 0.071+0.18/29/25+17 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 602 @4.527s 5%: 0.080+0+0.51 ms clock, 1.2+0.18/29/25+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 603 @4.695s 5%: 0.017+5.3+0.31 ms clock, 0.27+0.19/18/1.8+5.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 604 @4.869s 5%: 0.020+9.6+0.36 ms clock, 0.32+0.14/28/27+5.9 ms cpu, 7->8->8 MB, 8 MB goal, 16 P
HeapAlloc: 8996912 HeapInuse: 9289728 HeapObjects: 150312 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 605 @4.894s 5%: 0.004+0+1.0 ms clock, 0.069+0.14/28/27+17 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 606 @4.895s 5%: 0.041+0+0.47 ms clock, 0.62+0.14/28/27+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 607 @5.063s 5%: 0.019+3.2+0.28 ms clock, 0.31+0.17/11/16+4.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 608 @5.238s 5%: 0.021+8.9+0.34 ms clock, 0.34+0.15/33/10+5.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996912 HeapInuse: 9289728 HeapObjects: 150312 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 609 @5.266s 5%: 0.003+0+1.1 ms clock, 0.058+0.15/33/10+17 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 610 @5.267s 5%: 0.059+0+0.44 ms clock, 0.88+0.15/33/10+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 611 @5.435s 5%: 0.020+3.9+0.29 ms clock, 0.32+0.13/12/20+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 612 @5.607s 4%: 0.018+6.4+0.33 ms clock, 0.29+0.13/22/44+5.2 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996960 HeapInuse: 9289728 HeapObjects: 150315 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 613 @5.633s 4%: 0.004+0+1.0 ms clock, 0.069+0.13/22/44+16 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 614 @5.634s 5%: 0.013+0+0.39 ms clock, 0.20+0.13/22/44+5.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 615 @5.807s 4%: 0.017+3.3+0.44 ms clock, 0.28+0.21/11/21+7.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 616 @5.978s 4%: 0.019+8.5+0.29 ms clock, 0.31+0.14/24/29+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8996912 HeapInuse: 9289728 HeapObjects: 150312 HeapIdle 409600 HeapReleased 0 HeapSys 9699328
gc 617 @6.008s 4%: 0.004+0+0.97 ms clock, 0.075+0.14/24/29+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 618 @6.009s 4%: 0.079+0+0.42 ms clock, 1.1+0.14/24/29+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 619 @6.176s 4%: 0.019+3.2+0.29 ms clock, 0.30+0.13/11/22+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 620 @6.348s 4%: 0.019+5.6+0.33 ms clock, 0.30+0.14/20/43+5.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998768 HeapInuse: 9306112 HeapObjects: 150315 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 621 @6.373s 4%: 0.002+0+0.96 ms clock, 0.043+0.14/20/43+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 622 @6.374s 4%: 0.002+0+0.46 ms clock, 0.040+0.14/20/43+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 623 @6.544s 4%: 0.021+3.9+0.28 ms clock, 0.34+0.13/11/20+4.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 624 @6.719s 4%: 0.020+6.2+0.26 ms clock, 0.32+0.19/22/48+4.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998800 HeapInuse: 9306112 HeapObjects: 150317 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 625 @6.745s 4%: 0.002+0+2.0 ms clock, 0.034+0.19/22/48+33 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 626 @6.747s 4%: 0.005+0+1.4 ms clock, 0.081+0.19/22/48+23 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 627 @6.921s 4%: 0.020+4.2+0.25 ms clock, 0.32+0.19/11/17+4.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 628 @7.096s 4%: 0.019+5.5+0.30 ms clock, 0.31+0.17/20/42+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998768 HeapInuse: 9306112 HeapObjects: 150315 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 629 @7.121s 4%: 0.002+0+0.94 ms clock, 0.034+0.17/20/42+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 630 @7.122s 4%: 0.046+0+0.39 ms clock, 0.73+0.17/20/42+6.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 631 @7.288s 4%: 0.019+3.7+0.30 ms clock, 0.31+0.17/12/16+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 632 @7.459s 4%: 0.019+6.8+0.29 ms clock, 0.31+0.16/23/47+4.7 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998800 HeapInuse: 9306112 HeapObjects: 150317 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 633 @7.486s 4%: 0.003+0+1.0 ms clock, 0.062+0.16/23/47+17 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 634 @7.487s 4%: 0.073+0+0.41 ms clock, 1.1+0.16/23/47+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 635 @7.655s 4%: 0.019+4.1+0.29 ms clock, 0.31+0.20/12/19+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 636 @7.830s 4%: 0.022+7.7+0.28 ms clock, 0.35+0.12/27/25+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998784 HeapInuse: 9306112 HeapObjects: 150316 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 637 @7.856s 4%: 0.002+0+0.95 ms clock, 0.036+0.12/27/25+15 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 638 @7.857s 4%: 0.004+0+0.49 ms clock, 0.076+0.12/27/25+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 639 @8.026s 4%: 0.021+3.9+0.25 ms clock, 0.34+0.092/12/17+4.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 640 @8.201s 4%: 0.020+7.1+0.27 ms clock, 0.32+0.16/24/40+4.3 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998752 HeapInuse: 9306112 HeapObjects: 150314 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 641 @8.226s 4%: 0.003+0+1.0 ms clock, 0.063+0.16/24/40+16 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 642 @8.227s 4%: 0.055+0+0.43 ms clock, 0.89+0.16/24/40+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 643 @8.396s 4%: 0.022+2.8+0.25 ms clock, 0.36+0.15/9.9/21+4.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 644 @8.567s 3%: 0.022+6.6+0.31 ms clock, 0.35+0.19/21/44+5.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998784 HeapInuse: 9306112 HeapObjects: 150316 HeapIdle 327680 HeapReleased 0 HeapSys 9633792
gc 645 @8.594s 3%: 0.002+0+0.93 ms clock, 0.033+0.19/21/44+14 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
gc 646 @8.595s 3%: 0.073+0+0.41 ms clock, 1.1+0.19/21/44+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 647 @8.762s 3%: 0.021+3.7+0.31 ms clock, 0.34+0.19/12/19+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 648 @8.934s 3%: 0.019+7.8+0.32 ms clock, 0.31+0/29/21+5.1 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 8998816 HeapInuse: 9314304 HeapObjects: 150318 HeapIdle 319488 HeapReleased 0 HeapSys 9633792
gc 649 @8.961s 3%: 0.009+0+1.0 ms clock, 0.15+0/29/21+16 ms cpu, 8->8->0 MB, 8 MB goal, 16 P (forced)
  100000	      3671 ns/op	      88 B/op	       2 allocs/op
gc 650 @8.963s 3%: 0.006+0+0.47 ms clock, 0.10+0/29/21+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 651 @9.097s 3%: 0.021+2.7+0.29 ms clock, 0.33+0.18/9.7/18+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 652 @9.289s 3%: 0.017+5.1+0.30 ms clock, 0.28+0.12/18/36+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 653 @9.511s 3%: 0.017+2.7+0.27 ms clock, 0.27+0.10/8.6/10+4.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6096096 HeapInuse: 6389760 HeapObjects: 87152 HeapIdle 10584064 HeapReleased 0 HeapSys 16973824
gc 654 @9.580s 3%: 0.002+0+0.85 ms clock, 0.040+0.10/8.6/10+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
#BenchmarkSetErase1E5-16                           	gc 655 @9.581s 3%: 0.032+0+0.53 ms clock, 0.51+0.10/8.6/10+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 656 @9.714s 3%: 0.019+4.5+0.28 ms clock, 0.30+0.15/12/14+4.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 657 @9.907s 3%: 0.021+6.4+0.29 ms clock, 0.34+0.23/23/25+4.7 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 658 @10.127s 3%: 0.018+2.8+0.28 ms clock, 0.29+0.17/9.4/15+4.6 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5934704 HeapInuse: 6225920 HeapObjects: 84404 HeapIdle 10715136 HeapReleased 0 HeapSys 16941056
gc 659 @10.195s 3%: 0.004+0+0.90 ms clock, 0.074+0.17/9.4/15+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 660 @10.196s 3%: 0.051+0+0.56 ms clock, 0.82+0.17/9.4/15+9.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 661 @10.332s 3%: 0.019+2.8+0.30 ms clock, 0.31+0.20/9.6/13+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 662 @10.526s 3%: 0.022+6.0+0.30 ms clock, 0.36+0.20/20/25+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 663 @10.745s 3%: 0.018+3.2+0.28 ms clock, 0.29+0.088/10/11+4.5 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6060352 HeapInuse: 6356992 HeapObjects: 86545 HeapIdle 10584064 HeapReleased 0 HeapSys 16941056
gc 664 @10.814s 3%: 0.004+0+0.80 ms clock, 0.073+0.088/10/11+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 665 @10.815s 3%: 0.074+0+0.56 ms clock, 1.1+0.088/10/11+9.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 666 @10.949s 3%: 0.019+2.8+0.29 ms clock, 0.31+0.16/9.2/11+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 667 @11.143s 3%: 0.023+5.8+0.32 ms clock, 0.38+0.16/21/38+5.2 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 668 @11.368s 3%: 0.021+3.2+0.30 ms clock, 0.34+0.15/9.3/8.8+4.9 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5934704 HeapInuse: 6225920 HeapObjects: 84404 HeapIdle 10747904 HeapReleased 0 HeapSys 16973824
gc 669 @11.435s 3%: 0.003+0+0.89 ms clock, 0.059+0.15/9.3/8.8+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 670 @11.436s 3%: 0.073+0+0.52 ms clock, 1.1+0.15/9.3/8.8+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 671 @11.569s 3%: 0.020+2.1+0.30 ms clock, 0.32+0.12/7.7/17+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 672 @11.757s 3%: 0.019+4.6+0.30 ms clock, 0.31+0.14/16/37+4.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 673 @11.977s 3%: 0.021+2.9+0.32 ms clock, 0.34+0.20/9.0/14+5.2 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6168000 HeapInuse: 6463488 HeapObjects: 88381 HeapIdle 10510336 HeapReleased 0 HeapSys 16973824
gc 674 @12.048s 3%: 0.004+0+0.87 ms clock, 0.068+0.20/9.0/14+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 675 @12.049s 3%: 0.063+0+0.57 ms clock, 1.0+0.20/9.0/14+9.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 676 @12.183s 3%: 0.018+2.7+0.33 ms clock, 0.30+0.19/8.9/15+5.4 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 677 @12.371s 3%: 0.019+5.2+0.28 ms clock, 0.31+0.15/18/42+4.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 678 @12.590s 3%: 0.021+3.3+0.27 ms clock, 0.35+0.15/8.9/13+4.3 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6042240 HeapInuse: 6340608 HeapObjects: 86237 HeapIdle 10633216 HeapReleased 0 HeapSys 16973824
gc 679 @12.658s 3%: 0.004+0+0.86 ms clock, 0.072+0.15/8.9/13+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 680 @12.659s 3%: 0.068+0+0.44 ms clock, 1.0+0.15/8.9/13+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 681 @12.793s 3%: 0.022+3.0+0.26 ms clock, 0.35+0.14/9.5/16+4.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 682 @12.991s 3%: 0.020+6.0+0.31 ms clock, 0.32+0.16/19/40+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 683 @13.215s 2%: 0.023+2.9+0.29 ms clock, 0.37+0.16/9.5/10+4.7 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5862896 HeapInuse: 6160384 HeapObjects: 83180 HeapIdle 10780672 HeapReleased 0 HeapSys 16941056
gc 684 @13.280s 2%: 0.003+0+0.81 ms clock, 0.061+0.16/9.5/10+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 685 @13.281s 2%: 0.043+0+0.51 ms clock, 0.69+0.16/9.5/10+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 686 @13.413s 2%: 0.019+2.6+0.28 ms clock, 0.31+0.18/8.8/16+4.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 687 @13.603s 2%: 0.020+5.4+0.28 ms clock, 0.32+0.15/19/36+4.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 688 @13.823s 2%: 0.020+3.5+0.40 ms clock, 0.32+0/11/11+6.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6319216 HeapInuse: 6619136 HeapObjects: 90960 HeapIdle 10354688 HeapReleased 0 HeapSys 16973824
gc 689 @13.896s 2%: 0.004+0+0.88 ms clock, 0.065+0/11/11+14 ms cpu, 6->6->0 MB, 6 MB goal, 16 P (forced)
gc 690 @13.897s 2%: 0.077+0+0.52 ms clock, 1.2+0/11/11+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 691 @14.031s 2%: 0.019+2.5+0.33 ms clock, 0.31+0.16/8.6/13+5.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 692 @14.222s 2%: 0.022+6.1+0.30 ms clock, 0.36+0.15/22/29+4.8 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 693 @14.450s 2%: 0.018+2.8+0.29 ms clock, 0.29+0.11/8.4/14+4.7 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5970608 HeapInuse: 6266880 HeapObjects: 85016 HeapIdle 10706944 HeapReleased 0 HeapSys 16973824
gc 694 @14.519s 2%: 0.002+0+0.79 ms clock, 0.041+0.11/8.4/14+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 695 @14.520s 2%: 0.039+0+0.46 ms clock, 0.62+0.11/8.4/14+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 696 @14.653s 2%: 0.016+2.9+0.30 ms clock, 0.26+0.21/8.8/15+4.9 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 697 @14.848s 2%: 0.021+4.4+0.32 ms clock, 0.35+0.19/15/34+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 698 @15.068s 2%: 0.021+2.4+0.27 ms clock, 0.34+0.18/7.8/15+4.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6096176 HeapInuse: 6389760 HeapObjects: 87156 HeapIdle 10551296 HeapReleased 0 HeapSys 16941056
gc 699 @15.137s 2%: 0.002+0+0.84 ms clock, 0.033+0.18/7.8/15+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 700 @15.138s 2%: 0.046+0+0.51 ms clock, 0.74+0.18/7.8/15+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 701 @15.270s 2%: 0.019+2.7+0.28 ms clock, 0.31+0.19/9.1/11+4.4 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 702 @15.460s 2%: 0.021+6.6+0.29 ms clock, 0.33+0.15/21/37+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 703 @15.681s 2%: 0.022+3.9+0.26 ms clock, 0.35+0.14/8.2/9.7+4.2 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5898768 HeapInuse: 6193152 HeapObjects: 83790 HeapIdle 10780672 HeapReleased 0 HeapSys 16973824
gc 704 @15.748s 2%: 0.002+0+0.80 ms clock, 0.046+0.14/8.2/9.7+12 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 705 @15.749s 2%: 0.017+0+0.44 ms clock, 0.27+0.14/8.2/9.7+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 706 @15.884s 2%: 0.021+2.7+0.25 ms clock, 0.34+0.15/8.9/18+4.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 707 @16.075s 2%: 0.021+4.8+1.5 ms clock, 0.34+0.16/17/32+24 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 708 @16.300s 2%: 0.026+2.9+0.32 ms clock, 0.42+0.13/9.3/10+5.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6203984 HeapInuse: 6496256 HeapObjects: 88994 HeapIdle 10477568 HeapReleased 0 HeapSys 16973824
gc 709 @16.372s 2%: 0.003+0+0.85 ms clock, 0.050+0.13/9.3/10+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 710 @16.373s 2%: 0.005+0+0.54 ms clock, 0.091+0.13/9.3/10+8.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 711 @16.505s 2%: 0.023+2.6+0.27 ms clock, 0.37+0.15/8.7/10+4.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 712 @16.694s 2%: 0.021+5.4+0.28 ms clock, 0.35+0.15/19/32+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 713 @16.914s 2%: 0.018+3.5+0.38 ms clock, 0.29+0.18/8.9/10+6.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6060272 HeapInuse: 6356992 HeapObjects: 86544 HeapIdle 10616832 HeapReleased 0 HeapSys 16973824
gc 714 @16.984s 2%: 0.003+0+0.83 ms clock, 0.056+0.18/8.9/10+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 715 @16.985s 2%: 0.034+0+0.46 ms clock, 0.55+0.18/8.9/10+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 716 @17.123s 2%: 0.021+3.5+0.28 ms clock, 0.35+0.11/11/12+4.5 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 717 @17.313s 2%: 0.021+10+0.33 ms clock, 0.33+0.15/25/20+5.3 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 718 @17.537s 2%: 0.022+3.3+0.28 ms clock, 0.35+0.18/9.4/7.1+4.6 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5791104 HeapInuse: 6086656 HeapObjects: 81957 HeapIdle 10854400 HeapReleased 0 HeapSys 16941056
gc 719 @17.602s 2%: 0.004+0+0.87 ms clock, 0.079+0.18/9.4/7.1+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 720 @17.603s 2%: 0.037+0+0.53 ms clock, 0.59+0.18/9.4/7.1+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 721 @17.734s 2%: 0.021+2.7+0.35 ms clock, 0.34+0.17/9.1/13+5.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 722 @17.925s 2%: 0.023+6.1+0.32 ms clock, 0.38+0.11/21/34+5.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 723 @18.146s 2%: 0.021+3.5+0.30 ms clock, 0.33+0.20/9.1/10+4.8 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5952592 HeapInuse: 6250496 HeapObjects: 84710 HeapIdle 10690560 HeapReleased 0 HeapSys 16941056
gc 724 @18.212s 2%: 0.002+0+0.94 ms clock, 0.042+0.20/9.1/10+15 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 725 @18.214s 2%: 0.047+0+0.44 ms clock, 0.76+0.20/9.1/10+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 726 @18.346s 2%: 0.021+2.6+0.29 ms clock, 0.34+0.16/8.2/13+4.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 727 @18.533s 2%: 0.018+6.1+0.29 ms clock, 0.30+0.18/20/31+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 728 @18.757s 2%: 0.020+2.5+0.29 ms clock, 0.32+0.14/8.4/8.1+4.7 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5952640 HeapInuse: 6250496 HeapObjects: 84709 HeapIdle 10723328 HeapReleased 0 HeapSys 16973824
gc 729 @18.823s 2%: 0.002+0+0.87 ms clock, 0.039+0.14/8.4/8.1+13 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
gc 730 @18.824s 2%: 0.038+0+0.48 ms clock, 0.61+0.14/8.4/8.1+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 731 @18.958s 2%: 0.023+3.6+0.38 ms clock, 0.37+0.14/12/12+6.1 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 732 @19.154s 2%: 0.023+6.6+0.30 ms clock, 0.37+0.073/23/30+4.8 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 733 @19.375s 2%: 0.022+2.5+0.30 ms clock, 0.35+0.19/7.8/13+4.9 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5916768 HeapInuse: 6209536 HeapObjects: 84099 HeapIdle 10764288 HeapReleased 0 HeapSys 16973824
gc 734 @19.441s 2%: 0.005+0+0.87 ms clock, 0.090+0.19/7.8/13+14 ms cpu, 5->5->0 MB, 5 MB goal, 16 P (forced)
  100000	      2436 ns/op	      96 B/op	       2 allocs/op
gc 735 @19.442s 2%: 0.050+0+0.77 ms clock, 0.80+0.19/7.8/13+12 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 736 @19.540s 2%: 0.033+0.91+0.36 ms clock, 0.53+0.14/2.2/1.0+5.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 737 @19.625s 2%: 0.017+1.1+0.28 ms clock, 0.28+0.13/3.0/2.5+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 738 @19.700s 2%: 0.018+0.75+0.42 ms clock, 0.29+0.15/1.8/1.5+6.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686688 HeapInuse: 1990656 HeapObjects: 11992 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 739 @19.708s 2%: 0.022+0+0.45 ms clock, 0.36+0.15/1.8/1.5+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E5-16                  	gc 740 @19.709s 2%: 0.066+0+0.48 ms clock, 1.0+0.15/1.8/1.5+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 741 @19.806s 2%: 0.018+0.70+0.30 ms clock, 0.29+0.16/1.6/1.1+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 742 @19.891s 2%: 0.013+0.96+0.31 ms clock, 0.20+0.11/2.5/2.1+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 743 @19.967s 2%: 0.021+0.80+0.37 ms clock, 0.33+0.17/1.8/0.95+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686768 HeapInuse: 2007040 HeapObjects: 11996 HeapIdle 14934016 HeapReleased 0 HeapSys 16941056
gc 744 @19.975s 2%: 0.022+0+0.58 ms clock, 0.36+0.17/1.8/0.95+9.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 745 @19.976s 2%: 0.073+0+0.47 ms clock, 1.1+0.17/1.8/0.95+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 746 @20.075s 2%: 0.020+0.94+0.37 ms clock, 0.33+0.20/1.6/0.54+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 747 @20.160s 2%: 0.016+1.0+0.28 ms clock, 0.26+0.15/2.6/3.0+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 748 @20.234s 2%: 0.014+0.72+0.35 ms clock, 0.22+0.15/1.7/1.0+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 2007040 HeapObjects: 11995 HeapIdle 14966784 HeapReleased 0 HeapSys 16973824
gc 749 @20.242s 2%: 0.022+0+0.68 ms clock, 0.35+0.15/1.7/1.0+10 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 750 @20.243s 2%: 0.042+0+0.51 ms clock, 0.68+0.15/1.7/1.0+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 751 @20.341s 2%: 0.017+0.95+0.52 ms clock, 0.28+0.19/1.8/0.46+8.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 752 @20.427s 2%: 0.015+1.0+0.36 ms clock, 0.24+0.16/3.0/2.6+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 753 @20.503s 2%: 0.021+0.78+0.39 ms clock, 0.33+0.16/1.6/1.1+6.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 2007040 HeapObjects: 11995 HeapIdle 14966784 HeapReleased 0 HeapSys 16973824
gc 754 @20.511s 2%: 0.030+0+0.60 ms clock, 0.49+0.16/1.6/1.1+9.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 755 @20.512s 2%: 0.078+0+0.48 ms clock, 1.2+0.16/1.6/1.1+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 756 @20.611s 2%: 0.019+0.83+0.32 ms clock, 0.31+0.13/1.8/1.0+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 757 @20.697s 2%: 0.015+1.3+0.39 ms clock, 0.25+0.12/3.9/0.75+6.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 758 @20.772s 2%: 0.020+0.72+0.27 ms clock, 0.33+0.16/1.6/1.5+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686736 HeapInuse: 1990656 HeapObjects: 11994 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 759 @20.781s 2%: 0.002+0+0.52 ms clock, 0.037+0.16/1.6/1.5+8.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 760 @20.781s 2%: 0.015+0+0.49 ms clock, 0.24+0.16/1.6/1.5+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 761 @20.880s 2%: 0.021+0.76+0.28 ms clock, 0.34+0.066/1.8/1.3+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 762 @20.965s 2%: 0.022+1.2+0.34 ms clock, 0.35+0.12/2.9/1.5+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 763 @21.040s 2%: 0.019+0.67+0.28 ms clock, 0.31+0.12/1.5/1.5+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 1998848 HeapObjects: 11995 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 764 @21.052s 2%: 0.004+0+0.61 ms clock, 0.079+0.12/1.5/1.5+9.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 765 @21.053s 2%: 0.007+0+0.48 ms clock, 0.11+0.12/1.5/1.5+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 766 @21.154s 2%: 0.021+0.86+0.33 ms clock, 0.34+0.20/1.5/0.82+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 767 @21.239s 2%: 0.021+1.1+0.41 ms clock, 0.34+0.13/2.4/3.0+6.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 768 @21.314s 2%: 0.019+1.1+0.37 ms clock, 0.31+0.20/1.7/0.81+5.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 1990656 HeapObjects: 11995 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 769 @21.323s 2%: 0.003+0+0.57 ms clock, 0.057+0.20/1.7/0.81+9.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 770 @21.324s 2%: 0.043+0+0.49 ms clock, 0.69+0.20/1.7/0.81+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 771 @21.421s 2%: 0.018+0.84+0.36 ms clock, 0.29+0.19/1.3/1.1+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 772 @21.507s 2%: 0.014+1.2+0.27 ms clock, 0.23+0.14/2.5/3.0+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 773 @21.582s 2%: 0.016+0.77+0.38 ms clock, 0.26+0.12/1.7/1.2+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686736 HeapInuse: 1998848 HeapObjects: 11994 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 774 @21.591s 2%: 0.025+0+0.54 ms clock, 0.41+0.12/1.7/1.2+8.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 775 @21.592s 2%: 0.005+0+0.43 ms clock, 0.088+0.12/1.7/1.2+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 776 @21.689s 2%: 0.016+0.88+0.29 ms clock, 0.27+0.12/2.2/0.72+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 777 @21.775s 2%: 0.021+1.0+0.38 ms clock, 0.34+0.12/3.1/0.84+6.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 778 @21.851s 2%: 0.020+1.7+0.28 ms clock, 0.32+0.17/2.8/0.50+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 1998848 HeapObjects: 11995 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 779 @21.860s 2%: 0.036+0+0.57 ms clock, 0.58+0.17/2.8/0.50+9.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 780 @21.860s 2%: 0.004+0+0.47 ms clock, 0.072+0.17/2.8/0.50+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 781 @21.960s 2%: 0.018+0.79+0.29 ms clock, 0.29+0.17/1.8/1.5+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 782 @22.048s 2%: 0.019+1.1+0.28 ms clock, 0.30+0.14/2.6/2.8+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 783 @22.130s 2%: 0.019+0.77+0.34 ms clock, 0.31+0.23/1.6/1.0+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 1998848 HeapObjects: 11995 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 784 @22.138s 2%: 0.003+0+0.53 ms clock, 0.048+0.23/1.6/1.0+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 785 @22.139s 2%: 0.082+0+0.47 ms clock, 1.3+0.23/1.6/1.0+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 786 @22.238s 2%: 0.022+0.91+0.29 ms clock, 0.35+0.21/2.0/0.60+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 787 @22.326s 2%: 0.020+1.2+0.35 ms clock, 0.32+0.13/3.2/2.5+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 788 @22.400s 2%: 0.022+0.71+0.32 ms clock, 0.35+0.13/1.4/1.1+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 1998848 HeapObjects: 11995 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 789 @22.409s 2%: 0.004+0+0.59 ms clock, 0.064+0.13/1.4/1.1+9.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 790 @22.410s 2%: 0.044+0+0.47 ms clock, 0.71+0.13/1.4/1.1+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 791 @22.509s 2%: 0.019+0.81+0.35 ms clock, 0.30+0.21/1.5/0.73+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 792 @22.595s 2%: 0.021+1.1+0.28 ms clock, 0.35+0.16/3.0/2.7+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 793 @22.670s 2%: 0.014+0.69+0.29 ms clock, 0.22+0.13/1.4/1.0+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686736 HeapInuse: 1990656 HeapObjects: 11994 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 794 @22.679s 2%: 0.001+0+0.56 ms clock, 0.029+0.13/1.4/1.0+8.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 795 @22.680s 2%: 0.043+0+0.38 ms clock, 0.69+0.13/1.4/1.0+6.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 796 @22.781s 2%: 0.021+1.8+0.36 ms clock, 0.34+0.26/1.5/0.67+5.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 797 @22.866s 2%: 0.018+1.1+0.26 ms clock, 0.29+0.093/3.2/2.0+4.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 798 @22.940s 2%: 0.012+0.66+0.40 ms clock, 0.20+0.17/1.3/0.65+6.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686752 HeapInuse: 1998848 HeapObjects: 11995 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 799 @22.949s 2%: 0.004+0+0.58 ms clock, 0.079+0.17/1.3/0.65+9.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 800 @22.949s 2%: 0.083+0+0.51 ms clock, 1.3+0.17/1.3/0.65+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 801 @23.047s 2%: 0.022+0.75+0.32 ms clock, 0.35+0.13/1.5/1.0+5.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 802 @23.132s 2%: 0.018+1.1+0.32 ms clock, 0.30+0.12/3.0/2.3+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 803 @23.208s 2%: 0.024+0.78+0.47 ms clock, 0.38+0.13/2.0/0.81+7.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686736 HeapInuse: 1990656 HeapObjects: 11994 HeapIdle 14983168 HeapReleased 0 HeapSys 16973824
gc 804 @23.217s 2%: 0.001+0+0.58 ms clock, 0.031+0.13/2.0/0.81+9.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 805 @23.218s 2%: 0.041+0+0.56 ms clock, 0.66+0.13/2.0/0.81+8.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 806 @23.315s 2%: 0.020+0.83+0.29 ms clock, 0.32+0.18/1.5/0.89+4.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 807 @23.401s 2%: 0.017+1.2+0.36 ms clock, 0.27+0.18/2.6/1.8+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 808 @23.477s 2%: 0.021+0.82+0.41 ms clock, 0.33+0.16/1.6/0.68+6.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686768 HeapInuse: 2023424 HeapObjects: 11996 HeapIdle 14950400 HeapReleased 0 HeapSys 16973824
gc 809 @23.486s 2%: 0.018+0+0.58 ms clock, 0.28+0.16/1.6/0.68+9.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 810 @23.487s 2%: 0.068+0+0.44 ms clock, 1.0+0.16/1.6/0.68+7.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 811 @23.585s 2%: 0.015+0.72+0.28 ms clock, 0.25+0.17/1.6/0.79+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 812 @23.670s 2%: 0.015+1.1+0.30 ms clock, 0.24+0.15/3.1/2.4+4.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 813 @23.747s 2%: 0.019+0.70+0.35 ms clock, 0.31+0.13/1.4/1.5+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686736 HeapInuse: 1998848 HeapObjects: 11994 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 814 @23.755s 2%: 0.002+0+0.54 ms clock, 0.037+0.13/1.4/1.5+8.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 815 @23.756s 2%: 0.042+0+0.47 ms clock, 0.67+0.13/1.4/1.5+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 816 @23.855s 2%: 0.021+0.83+0.28 ms clock, 0.35+0.14/1.7/0.54+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 817 @23.942s 2%: 0.018+1.1+0.38 ms clock, 0.29+0.11/2.8/2.1+6.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 818 @24.018s 2%: 0.021+0.73+0.25 ms clock, 0.33+0.059/1.5/1.0+4.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1686736 HeapInuse: 1998848 HeapObjects: 11994 HeapIdle 14974976 HeapReleased 0 HeapSys 16973824
gc 819 @24.026s 2%: 0.002+0+0.45 ms clock, 0.040+0.059/1.5/1.0+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
  100000	      2700 ns/op	      96 B/op	       2 allocs/op
gc 820 @24.027s 2%: 0.031+0+0.44 ms clock, 0.49+0.059/1.5/1.0+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 821 @24.228s 2%: 0.020+0.55+0.33 ms clock, 0.32+0.16/1.2/0.68+5.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924456 HeapInuse: 3235840 HeapObjects: 40317 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 822 @24.343s 2%: 0.045+0+0.73 ms clock, 0.73+0.16/1.2/0.68+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E5-16          	gc 823 @24.344s 2%: 0.013+0+0.49 ms clock, 0.20+0.16/1.2/0.68+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 824 @24.541s 2%: 0.021+0.61+0.37 ms clock, 0.33+0.23/0.83/0.53+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924552 HeapInuse: 3235840 HeapObjects: 40321 HeapIdle 13705216 HeapReleased 0 HeapSys 16941056
gc 825 @24.657s 2%: 0.003+0+0.63 ms clock, 0.048+0.23/0.83/0.53+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 826 @24.657s 2%: 0.044+0+0.48 ms clock, 0.71+0.23/0.83/0.53+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 827 @24.855s 2%: 0.020+0.61+0.28 ms clock, 0.32+0.24/1.3/0.79+4.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924520 HeapInuse: 3235840 HeapObjects: 40319 HeapIdle 13705216 HeapReleased 0 HeapSys 16941056
gc 828 @24.971s 2%: 0.003+0+0.61 ms clock, 0.052+0.24/1.3/0.79+9.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 829 @24.972s 2%: 0.005+0+0.51 ms clock, 0.091+0.24/1.3/0.79+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 830 @25.171s 2%: 0.019+0.61+0.34 ms clock, 0.31+0.24/1.3/0.43+5.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 831 @25.289s 2%: 0.026+0+0.72 ms clock, 0.42+0.24/1.3/0.43+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 832 @25.290s 2%: 0.028+0+0.47 ms clock, 0.44+0.24/1.3/0.43+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 833 @25.494s 2%: 0.019+0.59+0.32 ms clock, 0.31+0.15/1.4/0.36+5.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924552 HeapInuse: 3235840 HeapObjects: 40321 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 834 @25.609s 1%: 0.003+0+0.67 ms clock, 0.048+0.15/1.4/0.36+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 835 @25.610s 2%: 0.045+0+0.42 ms clock, 0.72+0.15/1.4/0.36+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 836 @25.806s 1%: 0.019+0.56+0.30 ms clock, 0.31+0.17/0.99/0.68+4.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924552 HeapInuse: 3235840 HeapObjects: 40321 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 837 @25.922s 1%: 0.025+0+0.65 ms clock, 0.41+0.17/0.99/0.68+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 838 @25.923s 1%: 0.078+0+0.49 ms clock, 1.2+0.17/0.99/0.68+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 839 @26.124s 1%: 0.020+0.62+0.38 ms clock, 0.33+0.20/1.1/0.56+6.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 840 @26.242s 1%: 0.026+0+0.71 ms clock, 0.42+0.20/1.1/0.56+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 841 @26.243s 1%: 0.073+0+0.57 ms clock, 1.1+0.20/1.1/0.56+9.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 842 @26.443s 1%: 0.021+0.62+0.31 ms clock, 0.35+0.21/1.0/0.27+5.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924552 HeapInuse: 3235840 HeapObjects: 40321 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 843 @26.560s 1%: 0.003+0+0.68 ms clock, 0.052+0.21/1.0/0.27+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 844 @26.561s 1%: 0.040+0+0.50 ms clock, 0.64+0.21/1.0/0.27+8.0 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 845 @26.761s 1%: 0.020+0.56+0.29 ms clock, 0.32+0.17/0.91/0.87+4.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 846 @26.877s 1%: 0.016+0+0.69 ms clock, 0.26+0.17/0.91/0.87+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 847 @26.878s 1%: 0.044+0+0.36 ms clock, 0.71+0.17/0.91/0.87+5.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 848 @27.077s 1%: 0.019+0.56+0.35 ms clock, 0.31+0.19/0.81/0.63+5.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924520 HeapInuse: 3235840 HeapObjects: 40319 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 849 @27.194s 1%: 0.028+0+0.69 ms clock, 0.45+0.19/0.81/0.63+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 850 @27.195s 1%: 0.076+0+0.50 ms clock, 1.2+0.19/0.81/0.63+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 851 @27.397s 1%: 0.026+0.66+0.38 ms clock, 0.43+0.18/1.0/0.74+6.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 852 @27.514s 1%: 0.029+0+0.63 ms clock, 0.46+0.18/1.0/0.74+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 853 @27.515s 1%: 0.012+0+0.49 ms clock, 0.20+0.18/1.0/0.74+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 854 @27.719s 1%: 0.025+0.59+0.40 ms clock, 0.40+0.17/1.2/0.28+6.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924552 HeapInuse: 3235840 HeapObjects: 40321 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 855 @27.837s 1%: 0.023+0+0.69 ms clock, 0.38+0.17/1.2/0.28+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 856 @27.837s 1%: 0.039+0+0.40 ms clock, 0.62+0.17/1.2/0.28+6.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 857 @28.035s 1%: 0.020+0.50+0.33 ms clock, 0.32+0.14/1.0/0.81+5.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 858 @28.152s 1%: 0.024+0+0.73 ms clock, 0.38+0.14/1.0/0.81+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 859 @28.153s 1%: 0.058+0+1.6 ms clock, 0.94+0.14/1.0/0.81+25 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 860 @28.356s 1%: 0.021+0.62+0.38 ms clock, 0.34+0.18/1.0/0.62+6.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924552 HeapInuse: 3235840 HeapObjects: 40321 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 861 @28.471s 1%: 0.044+0+0.63 ms clock, 0.70+0.18/1.0/0.62+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 862 @28.472s 1%: 0.049+0+0.51 ms clock, 0.79+0.18/1.0/0.62+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 863 @28.670s 1%: 0.017+0.57+0.35 ms clock, 0.28+0.19/0.63/0.88+5.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 864 @28.787s 1%: 0.023+0+0.66 ms clock, 0.37+0.19/0.63/0.88+10 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 865 @28.788s 1%: 0.039+0+0.41 ms clock, 0.63+0.19/0.63/0.88+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 866 @28.987s 1%: 0.020+0.61+0.28 ms clock, 0.32+0.17/1.2/0.60+4.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 867 @29.104s 1%: 0.003+0+0.69 ms clock, 0.056+0.17/1.2/0.60+11 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 868 @29.105s 1%: 0.047+0+0.41 ms clock, 0.76+0.17/1.2/0.60+6.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 869 @29.306s 1%: 0.020+0.59+0.40 ms clock, 0.33+0.17/1.1/0.39+6.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2924536 HeapInuse: 3235840 HeapObjects: 40320 HeapIdle 13737984 HeapReleased 0 HeapSys 16973824
gc 870 @29.425s 1%: 0.004+0+0.61 ms clock, 0.069+0.17/1.1/0.39+9.8 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
  100000	      3207 ns/op	      56 B/op	       1 allocs/op
gc 871 @29.426s 1%: 0.041+0+0.43 ms clock, 0.66+0.17/1.1/0.39+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 3849592 HeapInuse: 4268032 HeapObjects: 6534 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 872 @29.459s 1%: 0.003+0+0.56 ms clock, 0.059+0.17/1.1/0.39+8.9 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E5-16                   	gc 873 @29.460s 1%: 0.005+0+0.57 ms clock, 0.088+0.17/1.1/0.39+9.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3851768 HeapInuse: 4276224 HeapObjects: 6560 HeapIdle 12697600 HeapReleased 0 HeapSys 16973824
gc 874 @29.492s 1%: 0.003+0+0.66 ms clock, 0.057+0.17/1.1/0.39+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 875 @29.493s 1%: 0.044+0+0.47 ms clock, 0.71+0.17/1.1/0.39+7.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3850552 HeapInuse: 4276224 HeapObjects: 6543 HeapIdle 12697600 HeapReleased 0 HeapSys 16973824
gc 876 @29.526s 1%: 0.024+0+0.68 ms clock, 0.39+0.17/1.1/0.39+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 877 @29.526s 1%: 0.040+0+0.53 ms clock, 0.64+0.17/1.1/0.39+8.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3844160 HeapInuse: 4268032 HeapObjects: 6478 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 878 @29.558s 1%: 0.037+0+0.65 ms clock, 0.60+0.17/1.1/0.39+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 879 @29.559s 1%: 0.042+0+0.46 ms clock, 0.67+0.17/1.1/0.39+7.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3850328 HeapInuse: 4268032 HeapObjects: 6544 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 880 @29.590s 1%: 0.041+0+0.62 ms clock, 0.65+0.17/1.1/0.39+9.9 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 881 @29.590s 1%: 0.020+0+0.50 ms clock, 0.32+0.17/1.1/0.39+8.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3843896 HeapInuse: 4268032 HeapObjects: 6478 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 882 @29.621s 1%: 0.005+0+0.62 ms clock, 0.088+0.17/1.1/0.39+9.9 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 883 @29.622s 1%: 0.043+0+0.45 ms clock, 0.68+0.17/1.1/0.39+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3847368 HeapInuse: 4268032 HeapObjects: 6513 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 884 @29.653s 1%: 0.049+0+0.51 ms clock, 0.78+0.17/1.1/0.39+8.2 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 885 @29.654s 1%: 0.073+0+0.47 ms clock, 1.1+0.17/1.1/0.39+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3825976 HeapInuse: 4268032 HeapObjects: 6470 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 886 @29.684s 1%: 0.015+0+0.50 ms clock, 0.24+0.17/1.1/0.39+8.0 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 887 @29.685s 1%: 0.054+0+0.54 ms clock, 0.87+0.17/1.1/0.39+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3859520 HeapInuse: 4284416 HeapObjects: 6642 HeapIdle 12689408 HeapReleased 0 HeapSys 16973824
gc 888 @29.717s 1%: 0.003+0+0.67 ms clock, 0.063+0.17/1.1/0.39+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 889 @29.717s 1%: 0.025+0+0.51 ms clock, 0.40+0.17/1.1/0.39+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3854864 HeapInuse: 4276224 HeapObjects: 6590 HeapIdle 12697600 HeapReleased 0 HeapSys 16973824
gc 890 @29.749s 1%: 0.046+0+0.59 ms clock, 0.74+0.17/1.1/0.39+9.5 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 891 @29.750s 1%: 0.046+0+0.47 ms clock, 0.74+0.17/1.1/0.39+7.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3853880 HeapInuse: 4276224 HeapObjects: 6580 HeapIdle 12697600 HeapReleased 0 HeapSys 16973824
gc 892 @29.781s 1%: 0.003+0+0.66 ms clock, 0.063+0.17/1.1/0.39+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 893 @29.782s 1%: 0.020+0+0.48 ms clock, 0.32+0.17/1.1/0.39+7.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3849960 HeapInuse: 4268032 HeapObjects: 6540 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 894 @29.813s 1%: 0.004+0+0.80 ms clock, 0.066+0.17/1.1/0.39+12 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 895 @29.814s 1%: 0.045+0+0.57 ms clock, 0.73+0.17/1.1/0.39+9.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3824456 HeapInuse: 4259840 HeapObjects: 6457 HeapIdle 12713984 HeapReleased 0 HeapSys 16973824
gc 896 @29.846s 1%: 0.038+0+0.64 ms clock, 0.60+0.17/1.1/0.39+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 897 @29.846s 1%: 0.064+0+0.56 ms clock, 1.0+0.17/1.1/0.39+9.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 3848448 HeapInuse: 4268032 HeapObjects: 6522 HeapIdle 12705792 HeapReleased 0 HeapSys 16973824
gc 898 @29.878s 1%: 0.005+0+0.75 ms clock, 0.086+0.17/1.1/0.39+12 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
  100000	       313 ns/op	      36 B/op	       0 allocs/op
gc 899 @29.879s 1%: 0.005+0+0.59 ms clock, 0.089+0.17/1.1/0.39+9.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 900 @29.895s 1%: 0.013+0.32+0.39 ms clock, 0.22+0.090/0.25/0.22+6.2 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3676664 HeapInuse: 4079616 HeapObjects: 4665 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 901 @29.919s 1%: 0.003+0+0.63 ms clock, 0.057+0.090/0.25/0.22+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E5-16                    	gc 902 @29.920s 1%: 0.050+0+0.54 ms clock, 0.80+0.090/0.25/0.22+8.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 903 @29.936s 1%: 0.013+0.26+0.36 ms clock, 0.20+0.094/0.091/0.41+5.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3685336 HeapInuse: 4079616 HeapObjects: 4715 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 904 @29.963s 1%: 0.002+0+0.55 ms clock, 0.041+0.094/0.091/0.41+8.9 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 905 @29.964s 1%: 0.029+0+0.54 ms clock, 0.46+0.094/0.091/0.41+8.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 906 @29.980s 1%: 0.013+0.29+0.33 ms clock, 0.21+0.10/0.20/0.33+5.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3687064 HeapInuse: 4079616 HeapObjects: 4733 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 907 @30.007s 1%: 0.004+0+0.66 ms clock, 0.074+0.10/0.20/0.33+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 908 @30.008s 1%: 0.005+0+1.2 ms clock, 0.093+0.10/0.20/0.33+19 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 909 @30.025s 1%: 0.016+0.22+0.26 ms clock, 0.25+0.14/0.12/0.44+4.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3689752 HeapInuse: 4087808 HeapObjects: 4761 HeapIdle 12886016 HeapReleased 0 HeapSys 16973824
gc 910 @30.053s 1%: 0.004+0+0.53 ms clock, 0.075+0.14/0.12/0.44+8.4 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 911 @30.053s 1%: 0.046+0+0.47 ms clock, 0.73+0.14/0.12/0.44+7.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 912 @30.070s 1%: 0.016+0.31+0.33 ms clock, 0.26+0.099/0.16/0.29+5.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3680088 HeapInuse: 4079616 HeapObjects: 4703 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 913 @30.097s 1%: 0.002+0+0.57 ms clock, 0.041+0.099/0.16/0.29+9.1 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 914 @30.098s 1%: 0.005+0+0.59 ms clock, 0.084+0.099/0.16/0.29+9.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 915 @30.114s 1%: 0.014+0.33+0.30 ms clock, 0.23+0.11/0.11/0.26+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3683992 HeapInuse: 4079616 HeapObjects: 4701 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 916 @30.140s 1%: 0.003+0+0.69 ms clock, 0.057+0.11/0.11/0.26+11 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 917 @30.141s 1%: 0.054+0+0.51 ms clock, 0.86+0.11/0.11/0.26+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 918 @30.156s 1%: 0.013+0.29+0.32 ms clock, 0.21+0.14/0.25/0.25+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3685336 HeapInuse: 4079616 HeapObjects: 4715 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 919 @30.180s 1%: 0.006+0+0.75 ms clock, 0.10+0.14/0.25/0.25+12 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 920 @30.181s 1%: 0.074+0+0.47 ms clock, 1.1+0.14/0.25/0.25+7.6 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 921 @30.197s 1%: 0.013+0.26+0.49 ms clock, 0.20+0.096/0.21/0.18+7.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3686392 HeapInuse: 4096000 HeapObjects: 4726 HeapIdle 12877824 HeapReleased 0 HeapSys 16973824
gc 922 @30.224s 1%: 0.006+0+0.60 ms clock, 0.098+0.096/0.21/0.18+9.6 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 923 @30.225s 1%: 0.004+0+0.53 ms clock, 0.069+0.096/0.21/0.18+8.5 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 924 @30.242s 1%: 0.016+0.29+0.39 ms clock, 0.26+0.10/0.17/0.27+6.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3687928 HeapInuse: 4079616 HeapObjects: 4742 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 925 @30.268s 1%: 0.035+0+0.54 ms clock, 0.56+0.10/0.17/0.27+8.6 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 926 @30.269s 1%: 0.035+0+0.51 ms clock, 0.56+0.10/0.17/0.27+8.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 927 @30.284s 1%: 0.014+0.30+0.41 ms clock, 0.23+0.090/0.18/0.22+6.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3676536 HeapInuse: 4079616 HeapObjects: 4666 HeapIdle 12894208 HeapReleased 0 HeapSys 16973824
gc 928 @30.309s 1%: 0.005+0+0.58 ms clock, 0.080+0.090/0.18/0.22+9.4 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 929 @30.309s 1%: 0.042+0+0.55 ms clock, 0.68+0.090/0.18/0.22+8.8 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 930 @30.325s 1%: 0.013+0.28+0.29 ms clock, 0.22+0.089/0.054/0.33+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3684696 HeapInuse: 4087808 HeapObjects: 4751 HeapIdle 12886016 HeapReleased 0 HeapSys 16973824
gc 931 @30.352s 1%: 0.002+0+0.55 ms clock, 0.046+0.089/0.054/0.33+8.8 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
gc 932 @30.352s 1%: 0.072+0+0.51 ms clock, 1.1+0.089/0.054/0.33+8.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 933 @30.369s 1%: 0.015+0.28+0.30 ms clock, 0.25+0.10/0.17/0.37+4.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3690904 HeapInuse: 4087808 HeapObjects: 4773 HeapIdle 12886016 HeapReleased 0 HeapSys 16973824
gc 934 @30.396s 1%: 0.020+0+0.66 ms clock, 0.32+0.10/0.17/0.37+10 ms cpu, 3->3->1 MB, 3 MB goal, 16 P (forced)
  100000	       119 ns/op	       8 B/op	       0 allocs/op
gc 935 @30.397s 1%: 0.004+0+0.61 ms clock, 0.075+0.10/0.17/0.37+9.9 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 765616 HeapInuse: 1146880 HeapObjects: 1344 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 936 @30.412s 1%: 0.055+0+0.49 ms clock, 0.88+0.10/0.17/0.37+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E5-16           	gc 937 @30.413s 1%: 0.040+0+0.44 ms clock, 0.64+0.10/0.17/0.37+7.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 764600 HeapInuse: 1146880 HeapObjects: 1338 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 938 @30.428s 1%: 0.043+0+0.47 ms clock, 0.70+0.10/0.17/0.37+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 939 @30.429s 1%: 0.077+0+0.50 ms clock, 1.2+0.10/0.17/0.37+8.1 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 766552 HeapInuse: 1146880 HeapObjects: 1360 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 940 @30.444s 1%: 0.018+0+0.48 ms clock, 0.29+0.10/0.17/0.37+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 941 @30.445s 1%: 0.053+0+0.45 ms clock, 0.84+0.10/0.17/0.37+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 768544 HeapInuse: 1146880 HeapObjects: 1381 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 942 @30.460s 1%: 0.003+0+0.46 ms clock, 0.057+0.10/0.17/0.37+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 943 @30.461s 1%: 0.044+0+0.41 ms clock, 0.70+0.10/0.17/0.37+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 765224 HeapInuse: 1146880 HeapObjects: 1348 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 944 @30.475s 1%: 0.002+0+0.47 ms clock, 0.037+0.10/0.17/0.37+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 945 @30.476s 1%: 0.042+0+0.52 ms clock, 0.68+0.10/0.17/0.37+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 763960 HeapInuse: 1146880 HeapObjects: 1333 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 946 @30.491s 1%: 0.016+0+0.47 ms clock, 0.26+0.10/0.17/0.37+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 947 @30.492s 1%: 0.074+0+0.42 ms clock, 1.1+0.10/0.17/0.37+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 769408 HeapInuse: 1146880 HeapObjects: 1390 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 948 @30.507s 1%: 0.016+0+0.53 ms clock, 0.26+0.10/0.17/0.37+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 949 @30.508s 1%: 0.052+0+0.41 ms clock, 0.83+0.10/0.17/0.37+6.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 768288 HeapInuse: 1146880 HeapObjects: 1378 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 950 @30.523s 1%: 0.046+0+0.52 ms clock, 0.74+0.10/0.17/0.37+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 951 @30.524s 1%: 0.069+0+0.49 ms clock, 1.1+0.10/0.17/0.37+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 770632 HeapInuse: 1146880 HeapObjects: 1400 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 952 @30.539s 1%: 0.039+0+0.45 ms clock, 0.62+0.10/0.17/0.37+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 953 @30.540s 1%: 0.072+0+0.46 ms clock, 1.1+0.10/0.17/0.37+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 765544 HeapInuse: 1146880 HeapObjects: 1347 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 954 @30.556s 1%: 0.043+0+0.53 ms clock, 0.70+0.10/0.17/0.37+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 955 @30.557s 1%: 0.040+0+0.48 ms clock, 0.64+0.10/0.17/0.37+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 766024 HeapInuse: 1146880 HeapObjects: 1356 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 956 @30.572s 1%: 0.046+0+0.49 ms clock, 0.74+0.10/0.17/0.37+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 957 @30.572s 1%: 0.076+0+0.45 ms clock, 1.2+0.10/0.17/0.37+7.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 766832 HeapInuse: 1146880 HeapObjects: 1362 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 958 @30.588s 1%: 0.042+0+0.53 ms clock, 0.68+0.10/0.17/0.37+8.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 959 @30.588s 1%: 0.074+0+0.46 ms clock, 1.1+0.10/0.17/0.37+7.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 767296 HeapInuse: 1146880 HeapObjects: 1367 HeapIdle 15826944 HeapReleased 0 HeapSys 16973824
gc 960 @30.604s 1%: 0.052+0+0.48 ms clock, 0.84+0.10/0.17/0.37+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
  100000	       156 ns/op	       5 B/op	       0 allocs/op
gc 961 @30.605s 1%: 0.006+0+0.51 ms clock, 0.10+0.10/0.17/0.37+8.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
HeapAlloc: 1722096 HeapInuse: 2031616 HeapObjects: 307 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 962 @30.617s 1%: 0.002+0+0.47 ms clock, 0.039+0.10/0.17/0.37+7.5 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E5-16    	gc 963 @30.618s 1%: 0.004+0+0.58 ms clock, 0.074+0.10/0.17/0.37+9.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 964 @30.631s 1%: 0.041+0+0.46 ms clock, 0.66+0.10/0.17/0.37+7.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 965 @30.631s 1%: 0.044+0+0.45 ms clock, 0.71+0.10/0.17/0.37+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 966 @30.644s 1%: 0.043+0+0.47 ms clock, 0.70+0.10/0.17/0.37+7.5 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 967 @30.644s 1%: 0.044+0+0.44 ms clock, 0.70+0.10/0.17/0.37+7.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 968 @30.657s 1%: 0.004+0+0.49 ms clock, 0.065+0.10/0.17/0.37+7.9 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 969 @30.657s 1%: 0.045+0+0.42 ms clock, 0.73+0.10/0.17/0.37+6.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 970 @30.670s 1%: 0.037+0+0.46 ms clock, 0.59+0.10/0.17/0.37+7.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 971 @30.670s 1%: 0.041+0+0.52 ms clock, 0.66+0.10/0.17/0.37+8.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 972 @30.683s 1%: 0.003+0+0.54 ms clock, 0.056+0.10/0.17/0.37+8.6 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 973 @30.683s 1%: 0.032+0+0.46 ms clock, 0.51+0.10/0.17/0.37+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 974 @30.696s 1%: 0.038+0+0.50 ms clock, 0.61+0.10/0.17/0.37+8.0 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 975 @30.696s 1%: 0.003+0+0.48 ms clock, 0.062+0.10/0.17/0.37+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 976 @30.709s 1%: 0.002+0+0.46 ms clock, 0.038+0.10/0.17/0.37+7.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 977 @30.709s 1%: 0.043+0+0.50 ms clock, 0.69+0.10/0.17/0.37+8.0 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 978 @30.722s 1%: 0.017+0+0.50 ms clock, 0.27+0.10/0.17/0.37+8.0 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 979 @30.723s 1%: 0.046+0+0.45 ms clock, 0.73+0.10/0.17/0.37+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 980 @30.735s 1%: 0.038+0+0.44 ms clock, 0.62+0.10/0.17/0.37+7.0 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 981 @30.736s 1%: 0.044+0+0.45 ms clock, 0.71+0.10/0.17/0.37+7.3 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 982 @30.748s 1%: 0.003+0+0.48 ms clock, 0.058+0.10/0.17/0.37+7.8 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 983 @30.749s 1%: 0.042+0+0.45 ms clock, 0.67+0.10/0.17/0.37+7.2 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 984 @30.761s 1%: 0.003+0+0.52 ms clock, 0.056+0.10/0.17/0.37+8.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
gc 985 @30.762s 1%: 0.070+0+0.40 ms clock, 1.1+0.10/0.17/0.37+6.4 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
HeapAlloc: 1722176 HeapInuse: 2031616 HeapObjects: 310 HeapIdle 14942208 HeapReleased 0 HeapSys 16973824
gc 986 @30.776s 1%: 0.002+0+0.46 ms clock, 0.042+0.10/0.17/0.37+7.4 ms cpu, 1->1->1 MB, 1 MB goal, 16 P (forced)
  100000	       147 ns/op	      15 B/op	       0 allocs/op
gc 987 @30.777s 1%: 0.031+0+0.48 ms clock, 0.50+0.10/0.17/0.37+7.7 ms cpu, 1->1->0 MB, 1 MB goal, 16 P (forced)
gc 988 @30.780s 1%: 0.015+28+0.39 ms clock, 0.24+0/0.087/28+6.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202000 HeapInuse: 8511488 HeapObjects: 307 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 989 @31.253s 1%: 0.002+0+0.52 ms clock, 0.045+0/0.087/28+8.4 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
#BenchmarkSort1E6-16                               	gc 990 @31.254s 1%: 0.029+0+0.49 ms clock, 0.47+0/0.087/28+7.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 991 @31.257s 1%: 0.014+28+0.41 ms clock, 0.23+0/0.14/28+6.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202064 HeapInuse: 8511488 HeapObjects: 310 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 992 @31.727s 1%: 0.003+0+0.92 ms clock, 0.059+0/0.14/28+14 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 993 @31.728s 1%: 0.043+0+0.43 ms clock, 0.69+0/0.14/28+6.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 994 @31.731s 1%: 0.015+29+0.36 ms clock, 0.24+0/0.25/29+5.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 995 @32.209s 1%: 0.005+0+0.49 ms clock, 0.084+0/0.25/29+7.8 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 996 @32.210s 1%: 0.039+0+0.47 ms clock, 0.63+0/0.25/29+7.5 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 997 @32.212s 1%: 0.015+29+0.39 ms clock, 0.25+0/0.18/29+6.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 998 @32.684s 1%: 0.002+0+0.49 ms clock, 0.045+0/0.18/29+7.9 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 999 @32.684s 1%: 0.005+0+0.52 ms clock, 0.087+0/0.18/29+8.4 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1000 @32.686s 1%: 0.011+28+0.32 ms clock, 0.19+0/0.17/29+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1001 @33.166s 1%: 0.004+0+0.60 ms clock, 0.078+0/0.17/29+9.7 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1002 @33.166s 1%: 0.044+0+0.47 ms clock, 0.70+0/0.17/29+7.6 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1003 @33.168s 1%: 0.012+29+0.32 ms clock, 0.20+0/0.071/29+5.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1004 @33.643s 1%: 0.004+0+0.59 ms clock, 0.077+0/0.071/29+9.5 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1005 @33.644s 1%: 0.005+0+0.43 ms clock, 0.080+0/0.071/29+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1006 @33.647s 1%: 0.020+29+0.34 ms clock, 0.32+0/0.072/29+5.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1007 @34.124s 1%: 0.004+0+0.55 ms clock, 0.075+0/0.072/29+8.9 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1008 @34.125s 1%: 0.042+0+0.43 ms clock, 0.68+0/0.072/29+6.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1009 @34.127s 1%: 0.015+29+0.32 ms clock, 0.24+0/0.13/29+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1010 @34.600s 1%: 0.002+0+0.55 ms clock, 0.043+0/0.13/29+8.8 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1011 @34.600s 1%: 0.073+0+0.45 ms clock, 1.1+0/0.13/29+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1012 @34.602s 1%: 0.014+29+0.38 ms clock, 0.23+0/0.32/29+6.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1013 @35.074s 1%: 0.017+0+0.53 ms clock, 0.27+0/0.32/29+8.6 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1014 @35.074s 1%: 0.073+0+0.51 ms clock, 1.1+0/0.32/29+8.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1015 @35.077s 1%: 0.016+28+0.37 ms clock, 0.25+0/0.043/29+6.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1016 @35.548s 1%: 0.003+0+0.58 ms clock, 0.057+0/0.043/29+9.4 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1017 @35.548s 1%: 0.073+0+0.55 ms clock, 1.1+0/0.043/29+8.9 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1018 @35.551s 1%: 0.013+29+0.41 ms clock, 0.21+0/0.24/29+6.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1019 @36.026s 1%: 0.005+0+0.48 ms clock, 0.082+0/0.24/29+7.6 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1020 @36.027s 1%: 0.005+0+0.58 ms clock, 0.095+0/0.24/29+9.3 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1021 @36.029s 1%: 0.015+29+0.37 ms clock, 0.24+0/0.24/29+6.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1022 @36.508s 1%: 0.003+0+0.51 ms clock, 0.059+0/0.24/29+8.2 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1023 @36.509s 1%: 0.043+0+0.42 ms clock, 0.69+0/0.24/29+6.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1024 @36.511s 1%: 0.014+28+0.35 ms clock, 0.22+0/0.20/29+5.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1025 @36.989s 1%: 0.003+0+0.48 ms clock, 0.052+0/0.20/29+7.7 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1026 @36.989s 1%: 0.031+0+0.45 ms clock, 0.50+0/0.20/29+7.2 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1027 @36.992s 1%: 0.016+29+0.30 ms clock, 0.26+0/0.18/28+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1028 @37.462s 1%: 0.005+0+0.51 ms clock, 0.081+0/0.18/28+8.2 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
gc 1029 @37.463s 1%: 0.087+0+0.48 ms clock, 1.4+0/0.18/28+7.7 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1030 @37.465s 1%: 0.020+28+0.33 ms clock, 0.32+0/0.11/29+5.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8202144 HeapInuse: 8511488 HeapObjects: 311 HeapIdle 8462336 HeapReleased 0 HeapSys 16973824
gc 1031 @37.937s 1%: 0.005+0+0.60 ms clock, 0.090+0/0.11/29+9.6 ms cpu, 7->7->0 MB, 7 MB goal, 16 P (forced)
 1000000	       473 ns/op	       8 B/op	       0 allocs/op
gc 1032 @37.937s 1%: 0.021+0+0.49 ms clock, 0.34+0/0.11/29+7.8 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1033 @38.108s 1%: 0.021+3.7+0.29 ms clock, 0.34+0.14/12/16+4.7 ms cpu, 4->4->4 MB, 7 MB goal, 16 P
gc 1034 @38.211s 1%: 0.024+4.8+0.27 ms clock, 0.39+0.14/18/28+4.4 ms cpu, 5->6->6 MB, 8 MB goal, 16 P
gc 1035 @38.360s 1%: 0.020+11+0.35 ms clock, 0.33+0/30/55+5.7 ms cpu, 10->10->10 MB, 12 MB goal, 16 P
gc 1036 @38.651s 1%: 0.019+16+0.29 ms clock, 0.30+0.19/59/97+4.7 ms cpu, 19->19->19 MB, 21 MB goal, 16 P
gc 1037 @39.274s 1%: 0.020+27+0.29 ms clock, 0.32+0.30/103/185+4.6 ms cpu, 37->38->37 MB, 39 MB goal, 16 P
gc 1038 @40.701s 1%: 0.045+58+0.26 ms clock, 0.72+0.48/230/486+4.2 ms cpu, 73->74->74 MB, 75 MB goal, 16 P
HeapAlloc: 88198928 HeapInuse: 88784896 HeapObjects: 1500322 HeapIdle 540672 HeapReleased 0 HeapSys 89325568
gc 1039 @41.175s 1%: 0.003+0+5.8 ms clock, 0.053+0.48/230/486+93 ms cpu, 84->84->0 MB, 84 MB goal, 16 P (forced)
#BenchmarkSetInsert1E6-16                          	 1000000	      3243 ns/op	      88 B/op	       2 allocs/op
gc 1040 @41.181s 1%: 0.004+0+1.0 ms clock, 0.078+0.48/230/486+17 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1041 @41.185s 1%: 0.016+0.26+0.37 ms clock, 0.27+0/0.33/0.24+5.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1042 @41.226s 1%: 0.015+1.7+0.56 ms clock, 0.25+0.20/4.8/1.1+8.9 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1043 @41.422s 1%: 0.020+3.7+0.30 ms clock, 0.32+0.16/13/27+4.8 ms cpu, 12->13->12 MB, 17 MB goal, 16 P
gc 1044 @41.734s 1%: 0.020+12+0.27 ms clock, 0.32+0.028/41/88+4.3 ms cpu, 22->22->22 MB, 25 MB goal, 16 P
gc 1045 @42.384s 1%: 0.021+20+0.29 ms clock, 0.33+0.36/79/195+4.7 ms cpu, 41->42->42 MB, 45 MB goal, 16 P
gc 1046 @43.297s 1%: 0.016+47+0.29 ms clock, 0.26+0.16/176/429+4.6 ms cpu, 81->81->81 MB, 84 MB goal, 16 P
gc 1047 @45.335s 1%: 0.035+11+0.28 ms clock, 0.56+0.13/39/84+4.5 ms cpu, 159->159->24 MB, 163 MB goal, 16 P
HeapAlloc: 43503776 HeapInuse: 43917312 HeapObjects: 602044 HeapIdle 124051456 HeapReleased 0 HeapSys 167968768
gc 1048 @45.854s 1%: 0.003+0+3.7 ms clock, 0.049+0.13/39/84+59 ms cpu, 41->41->0 MB, 41 MB goal, 16 P (forced)
#BenchmarkSetErase1E6-16                           	 1000000	      2096 ns/op	      96 B/op	       2 allocs/op
gc 1049 @45.858s 1%: 0.032+0+1.7 ms clock, 0.52+0.13/39/84+28 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1050 @45.862s 1%: 0.017+0.31+0.29 ms clock, 0.27+0/0.42/0.48+4.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1051 @45.902s 1%: 0.017+1.4+0.35 ms clock, 0.28+0.24/4.2/3.8+5.6 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1052 @46.096s 1%: 0.021+3.5+0.29 ms clock, 0.34+0.23/12/24+4.7 ms cpu, 12->13->12 MB, 17 MB goal, 16 P
gc 1053 @46.402s 1%: 0.021+3.3+0.29 ms clock, 0.35+0.26/7.9/11+4.6 ms cpu, 22->22->10 MB, 25 MB goal, 16 P
gc 1054 @46.744s 1%: 0.021+5.4+0.32 ms clock, 0.34+0.18/19/39+5.1 ms cpu, 18->19->14 MB, 20 MB goal, 16 P
gc 1055 @47.166s 1%: 0.022+3.1+0.29 ms clock, 0.36+0.25/10/18+4.7 ms cpu, 27->27->10 MB, 28 MB goal, 16 P
gc 1056 @47.568s 1%: 0.021+4.3+0.30 ms clock, 0.33+0.18/14/16+4.8 ms cpu, 20->21->11 MB, 21 MB goal, 16 P
gc 1057 @48.032s 1%: 0.020+7.1+0.27 ms clock, 0.32+0.20/26/30+4.3 ms cpu, 22->22->15 MB, 23 MB goal, 16 P
gc 1058 @48.565s 1%: 0.020+6.3+0.31 ms clock, 0.33+0/17/26+5.0 ms cpu, 29->29->12 MB, 30 MB goal, 16 P
HeapAlloc: 25802352 HeapInuse: 26140672 HeapObjects: 300313 HeapIdle 141828096 HeapReleased 0 HeapSys 167968768
gc 1059 @48.951s 1%: 0.005+0+2.4 ms clock, 0.092+0/17/26+39 ms cpu, 24->24->0 MB, 24 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E6-16                  	 1000000	      3093 ns/op	      96 B/op	       2 allocs/op
gc 1060 @48.954s 1%: 0.017+0+1.5 ms clock, 0.28+0/17/26+24 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1061 @48.958s 1%: 0.017+0.36+0.32 ms clock, 0.28+0/0.33/0.43+5.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1062 @49.012s 1%: 0.020+1.5+0.30 ms clock, 0.32+0.069/4.3/6.6+4.8 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1063 @49.254s 1%: 0.021+7.1+0.27 ms clock, 0.33+0/19/18+4.4 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1064 @49.935s 1%: 0.016+4.0+0.30 ms clock, 0.26+0.18/13/20+4.8 ms cpu, 22->22->12 MB, 25 MB goal, 16 P
gc 1065 @50.491s 1%: 0.020+6.4+0.28 ms clock, 0.32+0.16/19/37+4.6 ms cpu, 23->24->14 MB, 25 MB goal, 16 P
gc 1066 @51.661s 1%: 0.021+2.5+0.29 ms clock, 0.33+0.25/8.0/10+4.6 ms cpu, 28->28->10 MB, 29 MB goal, 16 P
HeapAlloc: 17805400 HeapInuse: 18137088 HeapObjects: 200321 HeapIdle 149831680 HeapReleased 0 HeapSys 167968768
gc 1067 @52.270s 1%: 0.005+0+1.7 ms clock, 0.090+0.25/8.0/10+28 ms cpu, 17->17->0 MB, 17 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E6-16          	 1000000	      3316 ns/op	      56 B/op	       1 allocs/op
gc 1068 @52.272s 1%: 0.049+0+1.2 ms clock, 0.79+0.25/8.0/10+20 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1069 @52.306s 1%: 0.019+0.38+0.31 ms clock, 0.30+0.21/0.42/0.64+4.9 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 1070 @52.343s 1%: 0.023+0.59+0.35 ms clock, 0.37+0.45/0.11/0.99+5.7 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1071 @52.406s 1%: 0.021+0.59+0.36 ms clock, 0.33+0.51/0.19/1.1+5.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1072 @52.425s 1%: 0.020+0.97+0.36 ms clock, 0.32+0.88/0.061/1.5+5.8 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1073 @52.541s 1%: 0.024+0.58+0.37 ms clock, 0.39+0.49/0.16/1.1+6.0 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1074 @52.605s 1%: 0.019+1.7+0.43 ms clock, 0.31+1.6/0.16/2.2+6.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38239688 HeapInuse: 38666240 HeapObjects: 31954 HeapIdle 129302528 HeapReleased 0 HeapSys 167968768
gc 1075 @52.766s 1%: 0.004+0+1.5 ms clock, 0.079+1.6/0.16/2.2+25 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E6-16                   	gc 1076 @52.768s 1%: 0.009+0+1.1 ms clock, 0.15+1.6/0.16/2.2+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1077 @52.801s 1%: 0.020+0.38+0.41 ms clock, 0.32+0.20/0.19/0.73+6.5 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1078 @52.840s 1%: 0.020+0.63+0.40 ms clock, 0.33+0.51/0.17/1.0+6.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1079 @52.905s 1%: 0.022+0.58+0.33 ms clock, 0.35+0.51/0.17/1.0+5.4 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1080 @52.925s 1%: 0.021+1.0+0.34 ms clock, 0.34+0.82/0.001/1.4+5.5 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1081 @53.037s 1%: 0.022+1.8+0.29 ms clock, 0.36+0/0.57/2.1+4.7 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1082 @53.103s 1%: 0.019+1.8+0.35 ms clock, 0.30+1.6/0.19/1.9+5.6 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38247240 HeapInuse: 38674432 HeapObjects: 32035 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1083 @53.269s 1%: 0.004+0+2.2 ms clock, 0.069+1.6/0.19/1.9+36 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1084 @53.272s 1%: 0.032+0+1.3 ms clock, 0.51+1.6/0.19/1.9+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1085 @53.307s 1%: 0.019+0.38+0.28 ms clock, 0.31+0.20/0.17/0.72+4.6 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1086 @53.343s 1%: 0.020+0.67+0.33 ms clock, 0.32+0.49/0.20/1.1+5.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1087 @53.404s 1%: 0.021+0.57+0.31 ms clock, 0.34+0.50/0.31/1.0+4.9 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1088 @53.423s 1%: 0.021+1.5+0.32 ms clock, 0.33+1.2/0.31/1.3+5.2 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1089 @53.534s 1%: 0.026+1.7+0.36 ms clock, 0.42+0/0.44/1.9+5.8 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1090 @53.598s 1%: 0.022+1.8+0.63 ms clock, 0.36+1.6/0.16/2.1+10 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38243784 HeapInuse: 38666240 HeapObjects: 31999 HeapIdle 129302528 HeapReleased 0 HeapSys 167968768
gc 1091 @53.766s 1%: 0.008+0+1.7 ms clock, 0.13+1.6/0.16/2.1+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1092 @53.768s 1%: 0.072+0+1.0 ms clock, 1.1+1.6/0.16/2.1+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1093 @53.803s 1%: 0.020+0.39+0.40 ms clock, 0.33+0.20/0.23/0.72+6.5 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1094 @53.840s 1%: 0.022+0.53+0.40 ms clock, 0.35+0.46/0.31/0.85+6.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1095 @53.903s 1%: 0.021+0.46+0.43 ms clock, 0.35+0.40/0.17/1.0+7.0 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1096 @53.922s 1%: 0.024+1.0+0.30 ms clock, 0.39+0.88/0.89/0.76+4.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1097 @54.032s 1%: 0.023+0.58+0.35 ms clock, 0.37+0.51/0.13/1.3+5.6 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1098 @54.098s 1%: 0.020+1.9+0.34 ms clock, 0.33+1.7/0.20/2.3+5.5 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38245800 HeapInuse: 38666240 HeapObjects: 32020 HeapIdle 129302528 HeapReleased 0 HeapSys 167968768
gc 1099 @54.264s 1%: 0.005+0+1.8 ms clock, 0.081+1.7/0.20/2.3+29 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1100 @54.266s 1%: 0.061+0+1.2 ms clock, 0.98+1.7/0.20/2.3+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1101 @54.301s 1%: 0.019+0.37+0.30 ms clock, 0.31+0.20/0.085/0.95+4.8 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1102 @54.337s 1%: 0.020+0.60+0.56 ms clock, 0.33+0.51/0.14/1.2+9.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1103 @54.396s 1%: 0.018+0.57+0.37 ms clock, 0.29+0.52/0.17/1.1+6.0 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1104 @54.415s 1%: 0.020+0.89+0.31 ms clock, 0.32+0.81/0.080/1.2+5.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1105 @54.524s 1%: 0.021+0.87+0.33 ms clock, 0.33+0.58/0.001/1.5+5.3 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1106 @54.589s 1%: 0.021+2.1+0.37 ms clock, 0.35+1.8/0.19/2.9+5.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38254920 HeapInuse: 38674432 HeapObjects: 32115 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1107 @54.748s 1%: 0.005+0+1.7 ms clock, 0.087+1.8/0.19/2.9+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1108 @54.750s 1%: 0.024+0+1.1 ms clock, 0.39+1.8/0.19/2.9+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1109 @54.784s 1%: 0.018+0.52+0.33 ms clock, 0.29+0.20/0.10/0.84+5.3 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1110 @54.821s 1%: 0.022+0.66+0.37 ms clock, 0.35+0.59/0.53/0.88+6.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1111 @54.883s 1%: 0.022+0.70+0.36 ms clock, 0.35+0.60/0.002/1.5+5.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1112 @54.902s 1%: 0.025+1.1+0.44 ms clock, 0.40+0.88/0.002/1.4+7.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1113 @55.013s 1%: 0.027+1.8+0.32 ms clock, 0.43+0/0.61/2.5+5.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1114 @55.078s 1%: 0.019+1.6+0.41 ms clock, 0.31+1.6/0.20/1.9+6.6 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38255112 HeapInuse: 38674432 HeapObjects: 32117 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1115 @55.244s 1%: 0.003+0+1.6 ms clock, 0.059+1.6/0.20/1.9+27 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1116 @55.246s 1%: 0.082+0+1.2 ms clock, 1.3+1.6/0.20/1.9+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1117 @55.280s 1%: 0.021+0.51+0.31 ms clock, 0.34+0.30/0.60/0.49+5.0 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1118 @55.317s 1%: 0.020+0.62+0.44 ms clock, 0.32+0.51/0.16/0.96+7.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1119 @55.377s 1%: 0.019+0.64+0.37 ms clock, 0.30+0.42/0.008/0.86+6.0 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1120 @55.399s 1%: 0.021+1.0+0.30 ms clock, 0.34+0.91/0.075/1.6+4.8 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1121 @55.504s 1%: 0.021+0.58+0.42 ms clock, 0.34+0.51/0.47/0.70+6.7 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1122 @55.570s 1%: 0.022+2.0+0.32 ms clock, 0.35+1.6/0.023/2.3+5.1 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38238888 HeapInuse: 38658048 HeapObjects: 31948 HeapIdle 129310720 HeapReleased 0 HeapSys 167968768
gc 1123 @55.729s 1%: 0.002+0+1.7 ms clock, 0.044+1.6/0.023/2.3+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1124 @55.731s 1%: 0.077+0+1.1 ms clock, 1.2+1.6/0.023/2.3+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1125 @55.765s 1%: 0.019+0.38+0.34 ms clock, 0.30+0.20/0.14/0.74+5.5 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1126 @55.804s 1%: 0.020+0.62+0.39 ms clock, 0.33+0.40/0.18/0.73+6.3 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1127 @55.864s 1%: 0.021+1.7+0.38 ms clock, 0.34+0.35/0.23/1.6+6.1 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1128 @55.883s 1%: 0.023+1.1+0.37 ms clock, 0.37+0.96/0.13/1.7+5.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1129 @55.993s 1%: 0.020+0.67+0.39 ms clock, 0.32+0.56/0.069/1.1+6.2 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1130 @56.062s 1%: 0.024+1.8+0.31 ms clock, 0.38+1.7/0.094/2.5+4.9 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38249064 HeapInuse: 38674432 HeapObjects: 32054 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1131 @56.256s 1%: 0.002+0+1.7 ms clock, 0.045+1.7/0.094/2.5+27 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1132 @56.258s 1%: 0.047+0+1.3 ms clock, 0.75+1.7/0.094/2.5+22 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1133 @56.294s 1%: 0.021+0.41+0.38 ms clock, 0.33+0.22/0.39/1.1+6.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1134 @56.337s 1%: 0.022+0.56+0.42 ms clock, 0.35+0.50/0.21/1.3+6.8 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1135 @56.412s 1%: 0.022+0.50+0.37 ms clock, 0.35+0.44/0.22/1.0+6.0 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1136 @56.435s 1%: 0.022+1.2+0.32 ms clock, 0.35+0.82/0.066/1.7+5.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1137 @56.569s 1%: 0.033+0.89+0.36 ms clock, 0.54+0.65/0.20/1.2+5.8 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1138 @56.644s 1%: 0.023+2.5+0.30 ms clock, 0.37+1.6/0.29/2.9+4.8 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38235720 HeapInuse: 38658048 HeapObjects: 31915 HeapIdle 129310720 HeapReleased 0 HeapSys 167968768
gc 1139 @56.821s 1%: 0.003+0+1.7 ms clock, 0.062+1.6/0.29/2.9+28 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1140 @56.823s 1%: 0.062+0+1.1 ms clock, 0.99+1.6/0.29/2.9+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1141 @56.857s 1%: 0.020+0.42+0.49 ms clock, 0.32+0.20/0.21/0.93+7.8 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1142 @56.894s 1%: 0.021+0.67+0.33 ms clock, 0.34+0.48/0.45/1.0+5.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1143 @56.955s 1%: 0.020+0.67+0.36 ms clock, 0.33+0.59/0.30/1.0+5.8 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1144 @56.974s 1%: 0.022+0.94+0.31 ms clock, 0.36+0.88/0.080/1.4+5.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1145 @57.083s 1%: 0.022+0.58+0.35 ms clock, 0.36+0.49/0.21/0.91+5.6 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1146 @57.134s 1%: 0.019+1.2+0.36 ms clock, 0.31+1.1/0.22/1.8+5.8 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38243112 HeapInuse: 38666240 HeapObjects: 31992 HeapIdle 129302528 HeapReleased 0 HeapSys 167968768
gc 1147 @57.255s 1%: 0.004+0+1.2 ms clock, 0.067+1.1/0.22/1.8+20 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1148 @57.256s 1%: 0.046+0+0.81 ms clock, 0.75+1.1/0.22/1.8+13 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1149 @57.278s 1%: 0.022+0.49+0.29 ms clock, 0.36+0.13/0.17/0.70+4.7 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1150 @57.303s 1%: 0.018+0.36+0.32 ms clock, 0.29+0.26/0.13/0.69+5.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1151 @57.366s 1%: 0.020+0.77+0.36 ms clock, 0.33+0.71/0.046/1.4+5.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1152 @57.381s 1%: 0.018+0.83+0.37 ms clock, 0.29+0.77/0.11/1.1+6.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1153 @57.459s 1%: 0.020+1.3+0.35 ms clock, 0.33+0/0.24/1.5+5.7 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1154 @57.525s 1%: 0.028+1.8+0.38 ms clock, 0.45+1.6/0.15/2.2+6.2 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38248584 HeapInuse: 38674432 HeapObjects: 32049 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1155 @57.641s 1%: 0.002+0+1.3 ms clock, 0.034+1.6/0.15/2.2+21 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1156 @57.642s 1%: 0.078+0+1.2 ms clock, 1.2+1.6/0.15/2.2+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1157 @57.664s 1%: 0.016+0.39+0.54 ms clock, 0.25+0.13/0.36/0.58+8.6 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1158 @57.690s 1%: 0.020+0.54+0.29 ms clock, 0.32+0.29/0.16/0.69+4.7 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1159 @57.753s 1%: 0.021+1.8+0.29 ms clock, 0.33+0/0.72/1.7+4.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1160 @57.773s 1%: 0.021+1.2+0.33 ms clock, 0.33+1.0/0.010/2.0+5.4 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1161 @57.884s 1%: 0.021+0.77+0.37 ms clock, 0.34+0.60/0.60/1.0+6.0 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1162 @57.935s 1%: 0.021+1.2+0.36 ms clock, 0.35+1.0/0.13/1.9+5.7 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38248968 HeapInuse: 38674432 HeapObjects: 32053 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1163 @58.105s 1%: 0.004+0+1.8 ms clock, 0.073+1.0/0.13/1.9+29 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1164 @58.107s 1%: 0.040+0+1.1 ms clock, 0.65+1.0/0.13/1.9+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1165 @58.142s 1%: 0.017+0.41+0.36 ms clock, 0.28+0.20/0.35/0.46+5.8 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1166 @58.179s 1%: 0.019+0.64+0.37 ms clock, 0.31+0.52/0.16/1.1+6.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1167 @58.222s 1%: 0.018+0.43+0.32 ms clock, 0.29+0.27/0.11/0.85+5.1 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1168 @58.243s 1%: 0.020+1.1+0.83 ms clock, 0.32+0.97/0/1.9+13 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1169 @58.355s 1%: 0.021+0.47+0.31 ms clock, 0.34+0.41/0.10/0.78+5.0 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1170 @58.405s 1%: 0.022+1.1+0.32 ms clock, 0.36+1.1/0.047/1.8+5.1 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38260200 HeapInuse: 38682624 HeapObjects: 32170 HeapIdle 129286144 HeapReleased 0 HeapSys 167968768
gc 1171 @58.570s 1%: 0.005+0+1.7 ms clock, 0.082+1.1/0.047/1.8+27 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1172 @58.572s 1%: 0.042+0+1.2 ms clock, 0.67+1.1/0.047/1.8+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1173 @58.607s 1%: 0.018+0.36+0.29 ms clock, 0.29+0.20/0.26/0.85+4.7 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1174 @58.644s 1%: 0.019+0.56+0.42 ms clock, 0.31+0.49/0.18/0.95+6.8 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1175 @58.685s 1%: 0.017+0.61+0.33 ms clock, 0.28+0.31/0.22/0.70+5.2 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1176 @58.707s 1%: 0.022+1.0+0.38 ms clock, 0.36+0.87/0.16/1.3+6.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1177 @58.823s 1%: 0.021+0.47+0.39 ms clock, 0.33+0.42/0.12/0.94+6.2 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1178 @58.891s 1%: 0.022+1.6+0.41 ms clock, 0.36+1.6/1.5/0.83+6.6 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38251944 HeapInuse: 38674432 HeapObjects: 32084 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1179 @59.058s 1%: 0.004+0+1.5 ms clock, 0.076+1.6/1.5/0.83+24 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
gc 1180 @59.060s 1%: 0.005+0+1.4 ms clock, 0.093+1.6/1.5/0.83+23 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1181 @59.094s 1%: 0.017+0.37+0.41 ms clock, 0.28+0.20/0.25/0.51+6.6 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1182 @59.120s 1%: 0.019+0.45+0.32 ms clock, 0.31+0.32/0.37/0.76+5.1 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1183 @59.184s 1%: 0.020+0.47+0.35 ms clock, 0.32+0.41/0.17/0.84+5.6 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1184 @59.205s 1%: 0.024+1.0+0.31 ms clock, 0.38+0.83/0.044/1.3+4.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1185 @59.317s 1%: 0.022+2.1+0.32 ms clock, 0.35+0/0.64/2.1+5.1 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1186 @59.383s 1%: 0.021+2.0+0.31 ms clock, 0.34+1.9/2.0/0.77+5.0 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
HeapAlloc: 38250984 HeapInuse: 38674432 HeapObjects: 32074 HeapIdle 129294336 HeapReleased 0 HeapSys 167968768
gc 1187 @59.551s 1%: 0.003+0+1.5 ms clock, 0.055+1.9/2.0/0.77+25 ms cpu, 36->36->22 MB, 36 MB goal, 16 P (forced)
 1000000	       491 ns/op	      54 B/op	       0 allocs/op
gc 1188 @59.553s 1%: 0.052+0+1.3 ms clock, 0.83+1.9/2.0/0.77+21 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1189 @59.556s 1%: 0.016+0.26+0.31 ms clock, 0.26+0/0.46/0.25+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1190 @59.566s 1%: 0.014+0.34+0.40 ms clock, 0.22+0.041/0.033/0.58+6.5 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1191 @59.585s 1%: 0.014+0.46+0.29 ms clock, 0.22+0.17/0.11/0.61+4.7 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1192 @59.652s 1%: 0.016+0.45+0.50 ms clock, 0.25+0.39/0.43/0.76+8.1 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1193 @59.664s 1%: 0.024+0.64+0.33 ms clock, 0.39+0.56/0.11/0.87+5.3 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1194 @59.783s 1%: 0.020+1.1+0.36 ms clock, 0.33+1.0/0.17/1.7+5.7 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46240584 HeapInuse: 46661632 HeapObjects: 31927 HeapIdle 121307136 HeapReleased 0 HeapSys 167968768
gc 1195 @60.117s 1%: 0.004+0+1.5 ms clock, 0.074+1.0/0.17/1.7+25 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E6-16                    	gc 1196 @60.118s 1%: 0.039+0+1.4 ms clock, 0.62+1.0/0.17/1.7+23 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1197 @60.122s 1%: 0.016+0.34+0.39 ms clock, 0.27+0/0.29/0.31+6.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1198 @60.129s 1%: 0.011+0.27+0.39 ms clock, 0.17+0.039/0.10/0.41+6.2 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1199 @60.156s 1%: 0.016+0.38+0.46 ms clock, 0.26+0.20/0.15/0.58+7.4 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1200 @60.228s 1%: 0.018+0.52+0.38 ms clock, 0.29+0.44/0.25/0.69+6.1 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1201 @60.239s 1%: 0.019+0.92+0.38 ms clock, 0.31+0/0.90/0.36+6.2 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1202 @60.415s 1%: 0.022+1.7+0.43 ms clock, 0.36+1.6/0.18/3.4+6.9 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46251576 HeapInuse: 46678016 HeapObjects: 32044 HeapIdle 121290752 HeapReleased 0 HeapSys 167968768
gc 1203 @60.681s 1%: 0.012+0+1.4 ms clock, 0.20+1.6/0.18/3.4+23 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1204 @60.683s 1%: 0.046+0+1.0 ms clock, 0.74+1.6/0.18/3.4+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1205 @60.685s 1%: 0.013+0.35+0.32 ms clock, 0.22+0/0.31/0.38+5.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1206 @60.695s 1%: 0.020+0.53+0.55 ms clock, 0.32+0.064/0.099/0.62+8.8 ms cpu, 8->9->8 MB, 15 MB goal, 16 P
gc 1207 @60.712s 1%: 0.011+0.39+0.31 ms clock, 0.18+0/0.28/0.66+5.0 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1208 @60.779s 1%: 0.016+0.53+0.43 ms clock, 0.25+0.48/0.16/1.1+6.9 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1209 @60.857s 1%: 0.017+0.41+0.31 ms clock, 0.28+0.32/0.17/1.0+5.0 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1210 @60.911s 1%: 0.025+2.2+0.29 ms clock, 0.40+2.1/0.13/2.9+4.7 ms cpu, 44->44->43 MB, 45 MB goal, 16 P
HeapAlloc: 46263576 HeapInuse: 46686208 HeapObjects: 32169 HeapIdle 121282560 HeapReleased 0 HeapSys 167968768
gc 1211 @61.266s 1%: 0.004+0+1.6 ms clock, 0.066+2.1/0.13/2.9+26 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1212 @61.268s 1%: 0.020+0+1.1 ms clock, 0.33+2.1/0.13/2.9+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1213 @61.271s 1%: 0.016+0.29+0.29 ms clock, 0.27+0/0.43/0.45+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1214 @61.281s 1%: 0.015+0.37+0.31 ms clock, 0.24+0.042/0.14/0.77+5.0 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1215 @61.308s 1%: 0.019+0.47+0.30 ms clock, 0.31+0.20/0.18/0.82+4.9 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1216 @61.411s 1%: 0.020+3.4+0.31 ms clock, 0.32+0/3.3/0.83+5.0 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1217 @61.427s 1%: 0.020+1.6+0.28 ms clock, 0.32+0/1.4/0.59+4.5 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1218 @61.593s 1%: 0.020+1.9+0.33 ms clock, 0.32+1.6/0.009/2.7+5.3 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46257896 HeapInuse: 46678016 HeapObjects: 32110 HeapIdle 121290752 HeapReleased 0 HeapSys 167968768
gc 1219 @61.855s 1%: 0.004+0+1.2 ms clock, 0.072+1.6/0.009/2.7+19 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1220 @61.856s 1%: 0.037+0+1.0 ms clock, 0.60+1.6/0.009/2.7+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1221 @61.859s 1%: 0.015+0.27+0.30 ms clock, 0.24+0/0.31/0.43+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1222 @61.869s 1%: 0.015+0.39+0.31 ms clock, 0.24+0.040/0.17/0.56+5.0 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1223 @61.888s 1%: 0.015+1.7+0.33 ms clock, 0.24+0.18/0.16/1.0+5.3 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1224 @61.994s 1%: 0.020+0.78+0.29 ms clock, 0.32+0.73/0.39/1.3+4.7 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1225 @62.058s 1%: 0.025+0.91+0.34 ms clock, 0.41+0.063/0.18/1.4+5.5 ms cpu, 26->26->25 MB, 28 MB goal, 16 P
gc 1226 @62.158s 1%: 0.020+1.8+0.29 ms clock, 0.32+1.7/0.13/2.4+4.7 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46269688 HeapInuse: 46694400 HeapObjects: 32233 HeapIdle 121274368 HeapReleased 0 HeapSys 167968768
gc 1227 @62.422s 1%: 0.003+0+1.6 ms clock, 0.052+1.7/0.13/2.4+26 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1228 @62.424s 1%: 0.075+0+1.2 ms clock, 1.2+1.7/0.13/2.4+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1229 @62.428s 1%: 0.016+0.26+0.38 ms clock, 0.26+0/0.45/0.38+6.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1230 @62.438s 1%: 0.018+0.35+0.28 ms clock, 0.29+0.040/0.11/0.67+4.5 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1231 @62.466s 1%: 0.017+0.46+0.33 ms clock, 0.28+0.21/0.18/0.90+5.3 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1232 @62.567s 1%: 0.023+0.81+0.39 ms clock, 0.37+0.76/0/1.5+6.3 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1233 @62.589s 1%: 0.021+1.0+0.28 ms clock, 0.34+0.008/0.16/1.5+4.5 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1234 @62.753s 1%: 0.019+1.8+0.31 ms clock, 0.31+1.7/0.19/2.5+5.0 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46236840 HeapInuse: 46661632 HeapObjects: 31891 HeapIdle 121307136 HeapReleased 0 HeapSys 167968768
gc 1235 @63.066s 1%: 0.002+0+1.3 ms clock, 0.036+1.7/0.19/2.5+22 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1236 @63.068s 1%: 0.055+0+0.81 ms clock, 0.89+1.7/0.19/2.5+13 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1237 @63.070s 1%: 0.015+0.30+0.38 ms clock, 0.24+0/0.34/0.28+6.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1238 @63.077s 1%: 0.012+0.54+0.31 ms clock, 0.20+0.055/0.12/0.80+5.0 ms cpu, 8->9->8 MB, 15 MB goal, 16 P
gc 1239 @63.105s 1%: 0.017+0.40+0.39 ms clock, 0.27+0.27/0.080/0.63+6.2 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1240 @63.181s 1%: 0.015+3.7+0.31 ms clock, 0.24+0/0.71/3.6+5.0 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1241 @63.258s 1%: 0.019+0.36+0.45 ms clock, 0.30+0.20/0.16/0.75+7.3 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1242 @63.339s 1%: 0.023+1.8+0.37 ms clock, 0.37+1.7/0.079/2.4+5.9 ms cpu, 44->44->43 MB, 45 MB goal, 16 P
HeapAlloc: 46228664 HeapInuse: 46653440 HeapObjects: 31806 HeapIdle 121315328 HeapReleased 0 HeapSys 167968768
gc 1243 @63.618s 1%: 0.018+0+1.3 ms clock, 0.29+1.7/0.079/2.4+21 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1244 @63.619s 1%: 0.042+0+1.0 ms clock, 0.68+1.7/0.079/2.4+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1245 @63.622s 1%: 0.014+0.26+0.27 ms clock, 0.23+0/0.32/0.43+4.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1246 @63.632s 1%: 0.015+0.27+0.33 ms clock, 0.24+0.046/0.16/0.58+5.4 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1247 @63.659s 1%: 0.018+0.39+0.40 ms clock, 0.29+0.29/0.20/0.84+6.4 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1248 @63.761s 1%: 0.020+2.8+0.34 ms clock, 0.32+0.69/0.005/2.7+5.5 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1249 @63.778s 1%: 0.022+1.1+0.31 ms clock, 0.35+0.87/0.96/0.85+5.0 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1250 @63.958s 1%: 0.085+1.9+0.47 ms clock, 1.3+1.6/0.035/2.6+7.5 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46250648 HeapInuse: 46669824 HeapObjects: 32035 HeapIdle 121298944 HeapReleased 0 HeapSys 167968768
gc 1251 @64.240s 1%: 0.004+0+1.3 ms clock, 0.070+1.6/0.035/2.6+21 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1252 @64.242s 1%: 0.072+0+1.1 ms clock, 1.1+1.6/0.035/2.6+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1253 @64.245s 1%: 0.014+0.29+0.44 ms clock, 0.23+0/0.42/0.32+7.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1254 @64.255s 1%: 0.015+0.37+0.30 ms clock, 0.24+0.042/0.15/0.59+4.8 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1255 @64.282s 1%: 0.016+0.40+0.42 ms clock, 0.25+0.20/0.15/0.56+6.8 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1256 @64.384s 1%: 0.020+0.91+0.33 ms clock, 0.32+0.85/0.071/1.3+5.4 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1257 @64.403s 1%: 0.018+1.1+0.32 ms clock, 0.30+0.009/0.94/0.93+5.1 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1258 @64.515s 1%: 0.019+1.3+0.38 ms clock, 0.31+1.1/0.26/1.8+6.1 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46247368 HeapInuse: 46669824 HeapObjects: 32001 HeapIdle 121298944 HeapReleased 0 HeapSys 167968768
gc 1259 @64.776s 1%: 0.004+0+1.4 ms clock, 0.069+1.1/0.26/1.8+22 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1260 @64.777s 1%: 0.016+0+1.0 ms clock, 0.26+1.1/0.26/1.8+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1261 @64.780s 1%: 0.011+0.28+0.27 ms clock, 0.18+0/0.36/0.21+4.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1262 @64.786s 1%: 0.010+0.29+0.34 ms clock, 0.17+0.025/0.10/0.48+5.4 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1263 @64.804s 1%: 0.015+0.56+0.35 ms clock, 0.25+0.14/0.26/0.59+5.6 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1264 @64.914s 1%: 0.022+0.79+0.34 ms clock, 0.36+0.71/0.14/1.4+5.5 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1265 @64.977s 1%: 0.023+1.5+0.33 ms clock, 0.36+0.088/0.24/2.1+5.3 ms cpu, 26->26->25 MB, 28 MB goal, 16 P
gc 1266 @65.121s 1%: 0.020+1.8+0.46 ms clock, 0.33+1.7/0.14/2.5+7.4 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46243704 HeapInuse: 46669824 HeapObjects: 31963 HeapIdle 121298944 HeapReleased 0 HeapSys 167968768
gc 1267 @65.473s 1%: 0.003+0+1.5 ms clock, 0.054+1.7/0.14/2.5+25 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1268 @65.475s 1%: 0.071+0+1.2 ms clock, 1.1+1.7/0.14/2.5+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1269 @65.478s 1%: 0.016+0.35+0.36 ms clock, 0.26+0/0.36/0.46+5.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1270 @65.485s 1%: 0.011+0.63+0.30 ms clock, 0.19+0.025/0.15/0.72+4.8 ms cpu, 8->9->8 MB, 15 MB goal, 16 P
gc 1271 @65.502s 1%: 0.015+0.36+0.33 ms clock, 0.25+0.13/0.16/0.65+5.3 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1272 @65.575s 1%: 0.019+3.2+0.40 ms clock, 0.31+0/0.77/2.8+6.4 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1273 @65.659s 1%: 0.017+0.47+0.31 ms clock, 0.28+0.23/0.17/0.85+5.1 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1274 @65.729s 1%: 0.023+2.0+0.33 ms clock, 0.36+1.9/1.7/1.1+5.3 ms cpu, 44->44->43 MB, 45 MB goal, 16 P
HeapAlloc: 46245992 HeapInuse: 46669824 HeapObjects: 31987 HeapIdle 121298944 HeapReleased 0 HeapSys 167968768
gc 1275 @66.015s 1%: 0.004+0+1.9 ms clock, 0.064+1.9/1.7/1.1+30 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1276 @66.018s 1%: 0.005+0+1.5 ms clock, 0.085+1.9/1.7/1.1+24 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1277 @66.021s 1%: 0.022+0.25+0.29 ms clock, 0.36+0/0.50/0.46+4.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1278 @66.031s 1%: 0.013+0.29+0.33 ms clock, 0.21+0.040/0.24/0.55+5.3 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1279 @66.058s 1%: 0.017+0.35+0.31 ms clock, 0.27+0.20/0.16/0.68+5.0 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1280 @66.159s 1%: 0.022+2.2+0.36 ms clock, 0.36+0/0.61/2.4+5.8 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1281 @66.174s 1%: 0.019+1.3+0.38 ms clock, 0.31+1.2/0.20/1.8+6.0 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1282 @66.346s 1%: 0.021+1.7+0.35 ms clock, 0.33+1.6/1.7/0.84+5.6 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46234840 HeapInuse: 46653440 HeapObjects: 31871 HeapIdle 121315328 HeapReleased 0 HeapSys 167968768
gc 1283 @66.688s 1%: 0.004+0+1.7 ms clock, 0.072+1.6/1.7/0.84+28 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
gc 1284 @66.690s 1%: 0.041+0+1.2 ms clock, 0.66+1.6/1.7/0.84+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1285 @66.694s 1%: 0.014+0.25+0.37 ms clock, 0.23+0/0.27/0.39+5.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1286 @66.704s 1%: 0.015+0.33+0.30 ms clock, 0.24+0.040/0.10/0.57+4.8 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1287 @66.732s 1%: 0.019+0.40+0.37 ms clock, 0.31+0.25/0.18/0.84+5.9 ms cpu, 13->13->12 MB, 16 MB goal, 16 P
gc 1288 @66.801s 1%: 0.019+1.5+0.31 ms clock, 0.31+0/1.4/0.61+5.0 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1289 @66.818s 1%: 0.021+1.1+0.39 ms clock, 0.34+0/0.93/0.79+6.3 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1290 @66.999s 1%: 0.022+1.7+0.39 ms clock, 0.35+1.6/0.021/2.3+6.2 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 46256824 HeapInuse: 46678016 HeapObjects: 32100 HeapIdle 121290752 HeapReleased 0 HeapSys 167968768
gc 1291 @67.341s 1%: 0.003+0+1.7 ms clock, 0.059+1.6/0.021/2.3+27 ms cpu, 44->44->22 MB, 44 MB goal, 16 P (forced)
 1000000	       179 ns/op	       8 B/op	       0 allocs/op
gc 1292 @67.343s 1%: 0.006+0+1.3 ms clock, 0.10+1.6/0.021/2.3+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1293 @67.361s 1%: 0.015+0.35+0.37 ms clock, 0.24+0.10/0.30/0.47+6.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4655896 HeapInuse: 5054464 HeapObjects: 11715 HeapIdle 162914304 HeapReleased 0 HeapSys 167968768
gc 1294 @67.516s 1%: 0.004+0+2.0 ms clock, 0.075+0.10/0.30/0.47+32 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E6-16           	gc 1295 @67.518s 1%: 0.017+0+1.3 ms clock, 0.27+0.10/0.30/0.47+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1296 @67.536s 1%: 0.016+0.34+0.29 ms clock, 0.26+0.11/0.13/0.60+4.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4670344 HeapInuse: 5070848 HeapObjects: 11825 HeapIdle 162897920 HeapReleased 0 HeapSys 167968768
gc 1297 @67.698s 1%: 0.005+0+1.8 ms clock, 0.090+0.11/0.13/0.60+28 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1298 @67.700s 1%: 0.040+0+1.2 ms clock, 0.65+0.11/0.13/0.60+19 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1299 @67.718s 1%: 0.016+0.40+0.38 ms clock, 0.26+0.10/0.20/0.75+6.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4659784 HeapInuse: 5054464 HeapObjects: 11715 HeapIdle 162914304 HeapReleased 0 HeapSys 167968768
gc 1300 @67.877s 1%: 0.003+0+1.8 ms clock, 0.050+0.10/0.20/0.75+28 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1301 @67.879s 1%: 0.071+0+1.3 ms clock, 1.1+0.10/0.20/0.75+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1302 @67.897s 1%: 0.014+0.36+0.43 ms clock, 0.23+0.098/0.12/0.63+6.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4659432 HeapInuse: 5062656 HeapObjects: 11754 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1303 @68.056s 1%: 0.018+0+1.9 ms clock, 0.28+0.098/0.12/0.63+31 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1304 @68.058s 1%: 0.045+0+1.3 ms clock, 0.73+0.098/0.12/0.63+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1305 @68.076s 1%: 0.016+0.35+0.34 ms clock, 0.26+0.18/0.13/0.54+5.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4666984 HeapInuse: 5062656 HeapObjects: 11790 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1306 @68.231s 1%: 0.006+0+1.8 ms clock, 0.10+0.18/0.13/0.54+28 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1307 @68.233s 1%: 0.075+0+1.2 ms clock, 1.2+0.18/0.13/0.54+19 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1308 @68.251s 1%: 0.016+0.42+0.31 ms clock, 0.26+0.099/0.15/0.48+5.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4660008 HeapInuse: 5062656 HeapObjects: 11760 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1309 @68.352s 1%: 0.022+0+1.3 ms clock, 0.35+0.099/0.15/0.48+21 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1310 @68.354s 1%: 0.073+0+1.1 ms clock, 1.1+0.099/0.15/0.48+18 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1311 @68.365s 1%: 0.015+0.42+0.31 ms clock, 0.24+0.065/0.11/0.87+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4666120 HeapInuse: 5062656 HeapObjects: 11781 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1312 @68.466s 1%: 0.004+0+1.8 ms clock, 0.070+0.065/0.11/0.87+29 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1313 @68.468s 1%: 0.055+0+1.0 ms clock, 0.89+0.065/0.11/0.87+16 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1314 @68.479s 1%: 0.012+0.36+0.37 ms clock, 0.20+0.069/0.10/0.51+5.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4662760 HeapInuse: 5062656 HeapObjects: 11746 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1315 @68.579s 1%: 0.002+0+1.6 ms clock, 0.033+0.069/0.10/0.51+26 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1316 @68.581s 1%: 0.072+0+1.0 ms clock, 1.1+0.069/0.10/0.51+16 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1317 @68.593s 1%: 0.015+0.40+0.31 ms clock, 0.24+0.064/0.14/0.70+5.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4663528 HeapInuse: 5062656 HeapObjects: 11754 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1318 @68.751s 1%: 0.003+0+1.4 ms clock, 0.054+0.064/0.14/0.70+23 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1319 @68.753s 1%: 0.031+0+1.1 ms clock, 0.50+0.064/0.14/0.70+18 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1320 @68.771s 1%: 0.018+0.34+0.30 ms clock, 0.28+0.10/0.095/0.61+4.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4665736 HeapInuse: 5062656 HeapObjects: 11777 HeapIdle 162906112 HeapReleased 0 HeapSys 167968768
gc 1321 @68.930s 1%: 0.003+0+1.4 ms clock, 0.062+0.10/0.095/0.61+23 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1322 @68.932s 1%: 0.065+0+1.3 ms clock, 1.0+0.10/0.095/0.61+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1323 @68.950s 1%: 0.049+0.32+0.33 ms clock, 0.79+0.10/0.10/0.76+5.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4674856 HeapInuse: 5070848 HeapObjects: 11872 HeapIdle 162897920 HeapReleased 0 HeapSys 167968768
gc 1324 @69.107s 1%: 0.007+0+1.8 ms clock, 0.12+0.10/0.10/0.76+29 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1325 @69.110s 1%: 0.045+0+1.3 ms clock, 0.72+0.10/0.10/0.76+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1326 @69.128s 1%: 0.014+0.68+0.42 ms clock, 0.23+0.13/0.28/0.49+6.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4656072 HeapInuse: 5054464 HeapObjects: 11719 HeapIdle 162914304 HeapReleased 0 HeapSys 167968768
gc 1327 @69.229s 1%: 0.002+0+1.3 ms clock, 0.035+0.13/0.28/0.49+21 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
gc 1328 @69.230s 1%: 0.041+0+1.3 ms clock, 0.66+0.13/0.28/0.49+21 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1329 @69.242s 1%: 0.016+0.39+0.34 ms clock, 0.25+0.066/0.15/0.58+5.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4659304 HeapInuse: 5054464 HeapObjects: 11710 HeapIdle 162914304 HeapReleased 0 HeapSys 167968768
gc 1330 @69.342s 1%: 0.023+0+1.4 ms clock, 0.38+0.066/0.15/0.58+23 ms cpu, 4->4->2 MB, 4 MB goal, 16 P (forced)
 1000000	       111 ns/op	       5 B/op	       0 allocs/op
gc 1331 @69.344s 1%: 0.020+0+1.2 ms clock, 0.32+0.066/0.15/0.58+19 ms cpu, 2->2->0 MB, 2 MB goal, 16 P (forced)
gc 1332 @69.349s 1%: 0.021+0.32+0.76 ms clock, 0.34+0/0.34/0.61+12 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070736 HeapInuse: 24379392 HeapObjects: 318 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1333 @69.544s 1%: 0.003+0+1.2 ms clock, 0.057+0/0.34/0.61+20 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E6-16    	gc 1334 @69.546s 1%: 0.012+0+1.2 ms clock, 0.20+0/0.34/0.61+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1335 @69.553s 1%: 0.020+0.53+0.46 ms clock, 0.32+0/0.11/0.89+7.3 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1336 @69.693s 1%: 0.019+0+1.0 ms clock, 0.31+0/0.11/0.89+16 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1337 @69.694s 1%: 0.074+0+0.95 ms clock, 1.1+0/0.11/0.89+15 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1338 @69.699s 1%: 0.019+0.40+0.32 ms clock, 0.30+0/0.23/0.70+5.2 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1339 @69.900s 1%: 0.003+0+1.3 ms clock, 0.057+0/0.23/0.70+21 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1340 @69.902s 1%: 0.041+0+1.2 ms clock, 0.66+0/0.23/0.70+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1341 @69.908s 1%: 0.020+0.37+0.28 ms clock, 0.33+0/0.070/0.78+4.6 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1342 @70.052s 1%: 0.002+0+1.0 ms clock, 0.038+0/0.070/0.78+16 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1343 @70.053s 1%: 0.051+0+0.86 ms clock, 0.82+0/0.070/0.78+13 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1344 @70.059s 1%: 0.016+0.58+0.28 ms clock, 0.25+0/0.14/0.81+4.5 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1345 @70.254s 1%: 0.004+0+1.4 ms clock, 0.071+0/0.14/0.81+23 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1346 @70.256s 1%: 0.056+0+1.2 ms clock, 0.90+0/0.14/0.81+20 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1347 @70.262s 1%: 0.021+0.65+0.34 ms clock, 0.33+0/0.009/0.97+5.4 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1348 @70.402s 1%: 0.002+0+0.93 ms clock, 0.039+0/0.009/0.97+14 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1349 @70.403s 1%: 0.056+0+0.84 ms clock, 0.91+0/0.009/0.97+13 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1350 @70.409s 1%: 0.024+0.47+0.30 ms clock, 0.39+0/0.011/0.93+4.8 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1351 @70.604s 1%: 0.004+0+1.5 ms clock, 0.076+0/0.011/0.93+25 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1352 @70.606s 1%: 0.037+0+1.1 ms clock, 0.60+0/0.011/0.93+19 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1353 @70.613s 1%: 0.027+0.49+0.36 ms clock, 0.43+0/0.096/0.69+5.7 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1354 @70.754s 1%: 0.016+0+1.1 ms clock, 0.26+0/0.096/0.69+18 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1355 @70.755s 1%: 0.069+0+1.0 ms clock, 1.1+0/0.096/0.69+16 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1356 @70.760s 1%: 0.021+0.54+0.29 ms clock, 0.33+0/0.15/0.68+4.6 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1357 @70.901s 1%: 0.002+0+1.0 ms clock, 0.036+0/0.15/0.68+16 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1358 @70.902s 1%: 0.072+0+1.0 ms clock, 1.1+0/0.15/0.68+17 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1359 @70.908s 1%: 0.017+0.48+0.31 ms clock, 0.28+0/0.20/0.58+4.9 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1360 @71.107s 1%: 0.003+0+1.2 ms clock, 0.056+0/0.20/0.58+19 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1361 @71.108s 1%: 0.080+0+1.3 ms clock, 1.2+0/0.20/0.58+22 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1362 @71.115s 1%: 0.020+0.44+0.39 ms clock, 0.32+0/0.080/0.90+6.3 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1363 @71.255s 1%: 0.005+0+1.0 ms clock, 0.087+0/0.080/0.90+17 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1364 @71.256s 1%: 0.078+0+0.79 ms clock, 1.2+0/0.080/0.90+12 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1365 @71.261s 1%: 0.022+0.36+0.30 ms clock, 0.36+0/0.21/0.58+4.8 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1366 @71.404s 1%: 0.012+0+1.1 ms clock, 0.19+0/0.21/0.58+19 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
gc 1367 @71.405s 1%: 0.040+0+1.1 ms clock, 0.65+0/0.21/0.58+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1368 @71.410s 1%: 0.017+0.60+0.27 ms clock, 0.27+0/0.14/0.86+4.3 ms cpu, 22->22->22 MB, 23 MB goal, 16 P
HeapAlloc: 24070816 HeapInuse: 24379392 HeapObjects: 321 HeapIdle 143589376 HeapReleased 0 HeapSys 167968768
gc 1369 @71.604s 1%: 0.004+0+1.3 ms clock, 0.077+0/0.14/0.86+21 ms cpu, 23->23->22 MB, 23 MB goal, 16 P (forced)
 1000000	       198 ns/op	      23 B/op	       0 allocs/op
gc 1370 @71.606s 1%: 0.043+0+1.1 ms clock, 0.68+0/0.14/0.86+18 ms cpu, 22->22->0 MB, 22 MB goal, 16 P (forced)
gc 1371 @71.623s 1%: 0.020+286+0.34 ms clock, 0.32+0/0.10/286+5.5 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
HeapAlloc: 80202352 HeapInuse: 80502784 HeapObjects: 318 HeapIdle 87465984 HeapReleased 0 HeapSys 167968768
gc 1372 @75.646s 1%: 0.002+0+1.2 ms clock, 0.039+0/0.10/286+19 ms cpu, 76->76->0 MB, 76 MB goal, 16 P (forced)
#BenchmarkSort1E7-16                               	10000000	       404 ns/op	       8 B/op	       0 allocs/op
gc 1373 @75.647s 1%: 0.005+0+0.82 ms clock, 0.085+0/0.10/286+13 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1374 @75.750s 1%: 0.015+3.2+0.28 ms clock, 0.25+0/9.7/10+4.6 ms cpu, 4->4->4 MB, 8 MB goal, 16 P
gc 1375 @75.813s 1%: 0.016+4.1+0.29 ms clock, 0.26+0.20/13/18+4.7 ms cpu, 5->6->6 MB, 8 MB goal, 16 P
gc 1376 @75.917s 1%: 0.020+7.2+0.29 ms clock, 0.32+0.16/26/56+4.7 ms cpu, 10->10->10 MB, 12 MB goal, 16 P
gc 1377 @76.199s 1%: 0.021+12+0.28 ms clock, 0.34+0.18/47/83+4.5 ms cpu, 19->19->19 MB, 21 MB goal, 16 P
gc 1378 @76.582s 1%: 0.018+29+0.34 ms clock, 0.29+0.18/110/159+5.4 ms cpu, 37->38->38 MB, 39 MB goal, 16 P
gc 1379 @77.474s 1%: 0.019+44+0.29 ms clock, 0.31+0.35/162/386+4.7 ms cpu, 73->74->74 MB, 76 MB goal, 16 P
gc 1380 @79.303s 1%: 0.017+71+0.29 ms clock, 0.28+0.30/285/760+4.6 ms cpu, 144->146->146 MB, 148 MB goal, 16 P
gc 1381 @83.602s 1%: 0.035+231+0.38 ms clock, 0.56+0.31/756/1464+6.1 ms cpu, 285->287->286 MB, 292 MB goal, 16 P
gc 1382 @92.508s 1%: 0.078+324+0.28 ms clock, 1.2+1.0/1293/3532+4.5 ms cpu, 558->564->563 MB, 573 MB goal, 16 P
HeapAlloc: 880199808 HeapInuse: 883605504 HeapObjects: 15000338 HeapIdle 540672 HeapReleased 0 HeapSys 884146176
###gc 1383 @100.883s 1%: 0.003+0+38 ms clock, 0.049+1.0/1293/3532+616 ms cpu, 840->840->0 MB, 840 MB goal, 16 P (forced)
#BenchmarkSetInsert1E7-16                          	10000000	      2527 ns/op	      88 B/op	       2 allocs/op
gc 1384 @100.924s 1%: 0.003+0+4.8 ms clock, 0.063+1.0/1293/3532+77 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1385 @100.945s 1%: 0.031+0.96+0.36 ms clock, 0.50+0.015/1.8/4.8+5.8 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1386 @100.977s 1%: 0.015+2.0+0.29 ms clock, 0.25+0.13/4.6/2.6+4.7 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1387 @102.130s 1%: 0.018+25+0.30 ms clock, 0.29+0.060/94/228+4.8 ms cpu, 115->116->116 MB, 155 MB goal, 16 P
gc 1388 @104.638s 1%: 0.016+94+0.32 ms clock, 0.26+0.52/343/788+5.1 ms cpu, 202->203->203 MB, 232 MB goal, 16 P
gc 1389 @109.347s 1%: 0.019+178+0.48 ms clock, 0.30+0.76/708/1910+7.7 ms cpu, 379->382->382 MB, 407 MB goal, 16 P
gc 1390 @118.852s 1%: 0.018+384+0.29 ms clock, 0.30+0.97/1537/4199+4.6 ms cpu, 734->742->740 MB, 764 MB goal, 16 P
GC forced
gc 3 @120.894s 0%: 0.20+0.81+0.54 ms clock, 1.6+0/2.0/2.5+4.3 ms cpu, 3->3->0 MB, 4 MB goal, 16 P
gc 1391 @138.360s 1%: 0.11+188+0.35 ms clock, 1.8+0.49/745/2024+5.6 ms cpu, 1444->1446->390 MB, 1481 MB goal, 16 P
HeapAlloc: 735829632 HeapInuse: 738451456 HeapObjects: 11175791 HeapIdle 782180352 HeapReleased 0 HeapSys 1520631808
gc 1392 @146.001s 1%: 0.011+0+36 ms clock, 0.17+0.49/745/2024+582 ms cpu, 702->702->0 MB, 702 MB goal, 16 P (forced)
#BenchmarkSetErase1E7-16                           	10000000	      2184 ns/op	      96 B/op	       2 allocs/op
gc 1393 @146.040s 1%: 0.018+0+9.3 ms clock, 0.28+0.49/745/2024+150 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1394 @146.066s 1%: 0.030+1.2+0.28 ms clock, 0.48+0.013/1.2/10+4.5 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1395 @146.120s 1%: 0.025+2.4+0.31 ms clock, 0.41+0.23/6.4/8.0+5.0 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1396 @147.215s 1%: 0.020+26+0.30 ms clock, 0.32+0.23/95/199+4.9 ms cpu, 115->116->116 MB, 155 MB goal, 16 P
gc 1397 @149.439s 1%: 0.016+31+0.33 ms clock, 0.26+0/98/239+5.4 ms cpu, 202->203->119 MB, 232 MB goal, 16 P
scvg0: inuse: 136, idle: 1314, sys: 1451, released: 0, consumed: 1451 (MB)
gc 1398 @151.931s 1%: 0.019+35+0.29 ms clock, 0.31+0.24/137/321+4.7 ms cpu, 222->223->138 MB, 238 MB goal, 16 P
gc 1399 @155.287s 1%: 0.019+13+0.30 ms clock, 0.31+0.66/51/104+4.8 ms cpu, 267->267->99 MB, 276 MB goal, 16 P
gc 1400 @157.588s 1%: 0.023+34+0.35 ms clock, 0.37+0.44/126/300+5.6 ms cpu, 193->194->128 MB, 198 MB goal, 16 P
gc 1401 @160.528s 1%: 0.019+44+0.38 ms clock, 0.30+0.64/175/427+6.0 ms cpu, 249->250->148 MB, 256 MB goal, 16 P
gc 1402 @163.983s 1%: 0.018+34+0.35 ms clock, 0.28+0/112/246+5.7 ms cpu, 288->289->120 MB, 296 MB goal, 16 P
gc 1403 @166.797s 1%: 0.023+6.4+0.39 ms clock, 0.38+0/21/53+6.3 ms cpu, 235->235->85 MB, 241 MB goal, 16 P
HeapAlloc: 99219712 HeapInuse: 99573760 HeapObjects: 324479 HeapIdle 1421058048 HeapReleased 0 HeapSys 1520631808
gc 1404 @166.986s 1%: 0.009+0+9.4 ms clock, 0.15+0/21/53+151 ms cpu, 94->94->0 MB, 94 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E7-16                  	10000000	      2094 ns/op	      96 B/op	       2 allocs/op
gc 1405 @166.998s 1%: 0.015+0+7.8 ms clock, 0.24+0/21/53+125 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1406 @167.023s 1%: 0.027+1.2+0.31 ms clock, 0.43+0/1.6/8.8+4.9 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1407 @167.069s 1%: 0.020+2.5+0.28 ms clock, 0.32+0.13/6.9/11+4.6 ms cpu, 77->77->77 MB, 152 MB goal, 16 P
gc 1408 @168.470s 1%: 0.015+28+0.29 ms clock, 0.24+0.22/107/156+4.7 ms cpu, 115->116->115 MB, 155 MB goal, 16 P
gc 1409 @172.853s 1%: 0.018+20+0.32 ms clock, 0.30+0.30/77/171+5.2 ms cpu, 202->202->110 MB, 231 MB goal, 16 P
gc 1410 @177.479s 1%: 0.017+24+0.41 ms clock, 0.27+0.90/94/236+6.6 ms cpu, 206->207->115 MB, 221 MB goal, 16 P
scvg0: inuse: 3, idle: 1, sys: 5, released: 0, consumed: 5 (MB)
gc 1411 @182.394s 1%: 0.017+36+0.30 ms clock, 0.28+10/131/317+4.8 ms cpu, 222->223->131 MB, 230 MB goal, 16 P
gc 1412 @188.738s 0%: 0.025+23+0.31 ms clock, 0.40+0/89/242+5.0 ms cpu, 255->255->123 MB, 262 MB goal, 16 P
HeapAlloc: 134048296 HeapInuse: 134578176 HeapObjects: 1121037 HeapIdle 1386053632 HeapReleased 0 HeapSys 1520631808
gc 1413 @189.915s 0%: 0.026+0+9.6 ms clock, 0.43+0/89/242+155 ms cpu, 127->127->0 MB, 127 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E7-16          	10000000	      2291 ns/op	      56 B/op	       1 allocs/op
gc 1414 @189.928s 0%: 0.011+0+7.4 ms clock, 0.17+0/89/242+119 ms cpu, 0->0->0 MB, 0 MB goal, 16 P (forced)
gc 1415 @189.958s 0%: 0.015+1.3+0.31 ms clock, 0.25+0.13/1.5/9.7+5.0 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 1416 @189.991s 0%: 0.018+1.4+0.29 ms clock, 0.29+0.24/0.80/10+4.7 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1417 @190.034s 0%: 0.015+1.3+0.29 ms clock, 0.24+0.28/0.67/7.9+4.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1418 @190.055s 0%: 0.016+2.0+0.43 ms clock, 0.27+0.68/1.4/15+6.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1419 @190.135s 0%: 0.047+1.8+1.2 ms clock, 0.75+0.38/0.24/11+19 ms cpu, 19->19->11 MB, 36 MB goal, 16 P
gc 1420 @190.184s 0%: 0.020+1.6+0.39 ms clock, 0.33+1.0/1.1/9.4+6.2 ms cpu, 36->36->35 MB, 37 MB goal, 16 P
gc 1421 @190.322s 0%: 0.017+1.5+0.33 ms clock, 0.27+0.23/0.18/9.0+5.4 ms cpu, 36->36->22 MB, 71 MB goal, 16 P
gc 1422 @190.512s 0%: 0.025+3.6+0.39 ms clock, 0.40+3.5/1.0/11+6.2 ms cpu, 73->73->71 MB, 74 MB goal, 16 P
gc 1423 @190.845s 0%: 0.027+4.6+0.39 ms clock, 0.43+0.49/3.5/8.8+6.2 ms cpu, 72->72->72 MB, 143 MB goal, 16 P
gc 1424 @191.320s 0%: 0.025+9.7+0.35 ms clock, 0.40+9.6/0.16/18+5.6 ms cpu, 174->174->143 MB, 175 MB goal, 16 P
gc 1425 @192.649s 0%: 0.015+32+0.41 ms clock, 0.24+0/6.2/34+6.5 ms cpu, 169->169->107 MB, 286 MB goal, 16 P
gc 1426 @192.795s 0%: 0.027+13+0.38 ms clock, 0.43+13/1.5/21+6.1 ms cpu, 286->286->286 MB, 287 MB goal, 16 P
HeapAlloc: 315476696 HeapInuse: 316014592 HeapObjects: 332634 HeapIdle 1204617216 HeapReleased 0 HeapSys 1520631808
gc 1427 @194.471s 0%: 0.002+0+15 ms clock, 0.042+13/1.5/21+250 ms cpu, 300->300->187 MB, 300 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E7-16                   	10000000	       454 ns/op	      44 B/op	       0 allocs/op
gc 1428 @194.489s 0%: 0.003+0+7.9 ms clock, 0.049+13/1.5/21+127 ms cpu, 187->187->0 MB, 187 MB goal, 16 P (forced)
gc 1429 @194.513s 0%: 0.021+1.3+0.30 ms clock, 0.34+0/1.3/8.2+4.9 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1430 @194.535s 0%: 0.016+1.8+0.34 ms clock, 0.25+0.026/0.008/8.8+5.4 ms cpu, 77->78->77 MB, 152 MB goal, 16 P
gc 1431 @194.745s 0%: 0.019+1.8+0.27 ms clock, 0.31+1.6/0.97/12+4.3 ms cpu, 127->127->112 MB, 155 MB goal, 16 P
gc 1432 @195.701s 0%: 0.025+4.2+0.35 ms clock, 0.40+4.2/0.61/13+5.6 ms cpu, 266->266->219 MB, 267 MB goal, 16 P
gc 1433 @197.389s 0%: 0.021+13+0.41 ms clock, 0.34+13/0.90/25+6.5 ms cpu, 424->424->362 MB, 438 MB goal, 16 P
HeapAlloc: 395429256 HeapInuse: 395968512 HeapObjects: 332109 HeapIdle 1124663296 HeapReleased 0 HeapSys 1520631808
gc 1434 @200.638s 0%: 0.008+0+12 ms clock, 0.13+13/0.90/25+204 ms cpu, 377->377->187 MB, 377 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E7-16                    	10000000	       160 ns/op	       8 B/op	       0 allocs/op
gc 1435 @200.653s 0%: 0.007+0+7.8 ms clock, 0.12+13/0.90/25+125 ms cpu, 187->187->0 MB, 187 MB goal, 16 P (forced)
gc 1436 @200.664s 0%: 0.012+1.2+0.33 ms clock, 0.19+0.042/2.1/7.1+5.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1437 @200.680s 0%: 0.014+1.2+0.30 ms clock, 0.23+0.030/1.0/8.4+4.9 ms cpu, 9->9->8 MB, 15 MB goal, 16 P
gc 1438 @200.710s 0%: 0.020+1.3+0.37 ms clock, 0.32+0.27/0.90/9.8+5.9 ms cpu, 13->13->12 MB, 17 MB goal, 16 P
gc 1439 @200.807s 0%: 0.022+1.9+0.63 ms clock, 0.36+0/0.55/8.0+10 ms cpu, 19->19->14 MB, 24 MB goal, 16 P
gc 1440 @200.823s 0%: 0.022+1.9+0.34 ms clock, 0.35+0/2.3/8.8+5.4 ms cpu, 25->25->25 MB, 28 MB goal, 16 P
gc 1441 @201.007s 0%: 0.025+2.0+0.33 ms clock, 0.41+1.9/2.0/10+5.2 ms cpu, 51->51->43 MB, 52 MB goal, 16 P
HeapAlloc: 48539368 HeapInuse: 48971776 HeapObjects: 48543 HeapIdle 1471660032 HeapReleased 0 HeapSys 1520631808
gc 1442 @202.763s 0%: 0.005+0+8.9 ms clock, 0.082+1.9/2.0/10+143 ms cpu, 46->46->24 MB, 46 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E7-16           	10000000	       210 ns/op	       6 B/op	       0 allocs/op
gc 1443 @202.775s 0%: 0.006+0+8.0 ms clock, 0.10+1.9/2.0/10+128 ms cpu, 24->24->0 MB, 24 MB goal, 16 P (forced)
gc 1444 @202.787s 0%: 0.013+1.2+0.30 ms clock, 0.21+0/2.3/7.0+4.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1445 @202.844s 0%: 0.020+1.7+0.32 ms clock, 0.33+0/0.24/12+5.1 ms cpu, 183->183->183 MB, 184 MB goal, 16 P
HeapAlloc: 192752544 HeapInuse: 193060864 HeapObjects: 323 HeapIdle 1327570944 HeapReleased 0 HeapSys 1520631808
gc 1446 @204.583s 0%: 0.005+0+8.1 ms clock, 0.085+0/0.24/12+130 ms cpu, 183->183->176 MB, 183 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E7-16    	10000000	       180 ns/op	      19 B/op	       0 allocs/op
PASS
ok  	github.com/cdongyang/library/rbtree	204.773s
