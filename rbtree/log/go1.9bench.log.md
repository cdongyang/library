gc 1 @0.077s 0%: 0.25+0.91+0.11 ms clock, 0.25+0/1.5/0.36+0.11 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
gc 2 @0.103s 0%: 0.12+0.59+0.20 ms clock, 0.97+0.28/1.2/0.94+1.6 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
gc 3 @0.126s 0%: 0.16+1.0+0.11 ms clock, 1.8+0.095/1.7/1.7+1.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
# github.com/cdongyang/library/rbtree
gc 1 @0.008s 2%: 0.25+1.4+0.15 ms clock, 0.50+0.20/3.7/1.7+0.30 ms cpu, 4->4->2 MB, 5 MB goal, 16 P
gc 2 @0.041s 1%: 0.076+1.6+0.10 ms clock, 0.99+0.17/2.5/6.7+1.4 ms cpu, 5->5->3 MB, 6 MB goal, 16 P
gc 3 @0.081s 1%: 0.015+2.0+0.20 ms clock, 0.25+0.17/5.0/8.4+3.3 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 4 @0.131s 1%: 0.020+2.1+0.11 ms clock, 0.32+0.32/5.7/8.8+1.8 ms cpu, 7->8->4 MB, 8 MB goal, 16 P
# github.com/cdongyang/library/rbtree_test
gc 1 @0.006s 3%: 0.080+2.4+0.14 ms clock, 0.56+0.13/4.4/2.7+1.0 ms cpu, 4->4->2 MB, 5 MB goal, 16 P
gc 2 @0.020s 4%: 0.065+3.1+0.13 ms clock, 0.98+0.21/7.4/10+2.0 ms cpu, 5->5->4 MB, 6 MB goal, 16 P
gc 3 @0.071s 2%: 0.024+2.9+0.13 ms clock, 0.38+0.22/7.7/15+2.0 ms cpu, 8->8->5 MB, 9 MB goal, 16 P
gc 4 @0.159s 1%: 0.020+3.6+0.13 ms clock, 0.32+0.34/11/16+2.1 ms cpu, 11->11->6 MB, 12 MB goal, 16 P
# testmain
gc 1 @0.010s 2%: 0.066+1.7+0.086 ms clock, 0.20+0.10/4.6/4.0+0.26 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
# testmain
gc 1 @0.002s 4%: 0.22+4.6+0.079 ms clock, 0.44+0.15/4.7/0.18+0.15 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 2 @0.020s 3%: 0.063+7.7+0.16 ms clock, 0.56+0/6.9/4.3+1.4 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 3 @0.055s 4%: 0.069+14+0.12 ms clock, 0.76+0.19/29/0.35+1.3 ms cpu, 14->15->13 MB, 15 MB goal, 16 P
gc 4 @0.121s 4%: 0.078+12+0.14 ms clock, 1.0+0.21/37/1.4+1.9 ms cpu, 24->24->23 MB, 26 MB goal, 16 P
=== RUN   TestSet
--- PASS: TestSet (0.00s)
gc 1 @0.001s 1%: 0.082+0.18+0.053 ms clock, 0.16+0/0.32/0.17+0.10 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 2 @0.002s 4%: 0.11+0.20+0.097 ms clock, 0.71+0/0.13/0.26+0.58 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 188808 HeapInuse: 589824 HeapObjects: 305 HeapIdle 917504 HeapReleased 0 HeapSys 1507328
gc 3 @0.003s 6%: 0.064+0.10+0.053 ms clock, 0.77+0/0.22/0.11+0.64 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
goos: linux
goarch: amd64
pkg: github.com/cdongyang/library/rbtree
#BenchmarkSort1E3-16                               	gc 4 @0.003s 8%: 0.058+0.11+0.091 ms clock, 0.88+0/0.24/0.19+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 198968 HeapInuse: 589824 HeapObjects: 327 HeapIdle 655360 HeapReleased 0 HeapSys 1245184
gc 5 @0.004s 10%: 0.009+0.17+0.094 ms clock, 0.14+0/0.29/0.18+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 6 @0.004s 11%: 0.010+0.13+0.10 ms clock, 0.16+0/0.21/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 200616 HeapInuse: 589824 HeapObjects: 329 HeapIdle 622592 HeapReleased 0 HeapSys 1212416
gc 7 @0.005s 12%: 0.022+0.23+0.085 ms clock, 0.35+0/0.21/0.19+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 8 @0.006s 12%: 0.014+0.16+0.080 ms clock, 0.23+0/0.23/0.29+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 202472 HeapInuse: 589824 HeapObjects: 334 HeapIdle 589824 HeapReleased 0 HeapSys 1179648
gc 9 @0.006s 13%: 0.008+0.14+0.079 ms clock, 0.13+0/0.21/0.27+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 10 @0.007s 13%: 0.012+0.14+0.068 ms clock, 0.20+0/0.17/0.19+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 204328 HeapInuse: 589824 HeapObjects: 339 HeapIdle 524288 HeapReleased 0 HeapSys 1114112
gc 11 @0.007s 13%: 0.009+0.14+0.074 ms clock, 0.15+0/0.23/0.22+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 12 @0.008s 14%: 0.010+0.14+0.078 ms clock, 0.16+0/0.24/0.18+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 207816 HeapInuse: 589824 HeapObjects: 347 HeapIdle 458752 HeapReleased 0 HeapSys 1048576
gc 13 @0.008s 14%: 0.011+0.12+0.080 ms clock, 0.18+0/0.27/0.19+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 14 @0.009s 15%: 0.010+0.12+0.086 ms clock, 0.16+0/0.24/0.20+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 211240 HeapInuse: 589824 HeapObjects: 354 HeapIdle 360448 HeapReleased 0 HeapSys 950272
gc 15 @0.010s 14%: 0.009+0.14+0.067 ms clock, 0.15+0/0.22/0.26+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 16 @0.010s 15%: 0.008+0.17+0.074 ms clock, 0.14+0/0.25/0.22+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 211400 HeapInuse: 589824 HeapObjects: 356 HeapIdle 360448 HeapReleased 0 HeapSys 950272
gc 17 @0.011s 15%: 0.008+0.15+0.073 ms clock, 0.12+0/0.32/0.18+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 18 @0.011s 15%: 0.012+0.13+0.084 ms clock, 0.19+0/0.22/0.28+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 213064 HeapInuse: 589824 HeapObjects: 359 HeapIdle 327680 HeapReleased 0 HeapSys 917504
gc 19 @0.012s 15%: 0.012+0.16+0.089 ms clock, 0.20+0/0.23/0.24+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 20 @0.012s 16%: 0.008+0.13+0.066 ms clock, 0.13+0/0.26/0.14+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 214920 HeapInuse: 589824 HeapObjects: 364 HeapIdle 294912 HeapReleased 0 HeapSys 884736
gc 21 @0.013s 15%: 0.008+0.24+0.067 ms clock, 0.14+0/0.24/0.18+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 22 @0.013s 16%: 0.008+0.18+0.087 ms clock, 0.13+0/0.20/0.26+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 216584 HeapInuse: 589824 HeapObjects: 367 HeapIdle 262144 HeapReleased 0 HeapSys 851968
gc 23 @0.014s 16%: 0.008+0.13+0.075 ms clock, 0.14+0/0.24/0.17+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 24 @0.014s 16%: 0.011+0.20+0.065 ms clock, 0.19+0/0.21/0.27+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 216680 HeapInuse: 589824 HeapObjects: 368 HeapIdle 262144 HeapReleased 0 HeapSys 851968
gc 25 @0.015s 16%: 0.007+0.14+0.075 ms clock, 0.12+0/0.22/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 26 @0.015s 16%: 0.010+0.17+0.079 ms clock, 0.16+0/0.25/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 216840 HeapInuse: 589824 HeapObjects: 370 HeapIdle 262144 HeapReleased 0 HeapSys 851968
gc 27 @0.016s 16%: 0.008+0.13+0.093 ms clock, 0.13+0/0.21/0.25+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 28 @0.016s 18%: 0.009+0.13+0.36 ms clock, 0.15+0/0.21/0.17+5.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 216936 HeapInuse: 589824 HeapObjects: 371 HeapIdle 262144 HeapReleased 0 HeapSys 851968
gc 29 @0.017s 18%: 0.010+0.16+0.086 ms clock, 0.16+0/0.23/0.37+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 30 @0.018s 18%: 0.008+0.15+0.078 ms clock, 0.13+0/0.21/0.23+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 217128 HeapInuse: 589824 HeapObjects: 373 HeapIdle 262144 HeapReleased 0 HeapSys 851968
gc 31 @0.018s 18%: 0.007+0.21+0.077 ms clock, 0.12+0/0.12/0.32+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	       586 ns/op	      10 B/op	       0 allocs/op
gc 32 @0.019s 18%: 0.009+0.20+0.068 ms clock, 0.15+0/0.19/0.18+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 297160 HeapInuse: 663552 HeapObjects: 1875 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 33 @0.021s 16%: 0.008+0.19+0.093 ms clock, 0.14+0/0.22/0.15+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsert1E3-16                          	gc 34 @0.022s 16%: 0.009+0.14+0.099 ms clock, 0.15+0/0.23/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 297288 HeapInuse: 663552 HeapObjects: 1877 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 35 @0.025s 15%: 0.014+0.20+0.085 ms clock, 0.23+0/0.27/0.32+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 36 @0.026s 15%: 0.008+0.17+0.10 ms clock, 0.13+0/0.26/0.24+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 297528 HeapInuse: 663552 HeapObjects: 1879 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 37 @0.028s 14%: 0.012+0.13+0.087 ms clock, 0.19+0/0.20/0.19+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 38 @0.028s 14%: 0.010+0.14+0.068 ms clock, 0.16+0/0.28/0.30+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 297784 HeapInuse: 663552 HeapObjects: 1882 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 39 @0.031s 13%: 0.009+0.29+0.074 ms clock, 0.14+0/0.26/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 40 @0.031s 13%: 0.012+0.12+0.092 ms clock, 0.19+0/0.23/0.14+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 297976 HeapInuse: 663552 HeapObjects: 1884 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 41 @0.033s 13%: 0.008+0.15+0.087 ms clock, 0.13+0/0.23/0.16+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 42 @0.034s 13%: 0.008+0.15+0.093 ms clock, 0.14+0/0.22/0.24+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 298168 HeapInuse: 663552 HeapObjects: 1886 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 43 @0.036s 12%: 0.010+0.14+0.072 ms clock, 0.16+0/0.21/0.17+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 44 @0.037s 13%: 0.007+0.15+0.079 ms clock, 0.12+0/0.27/0.19+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 298264 HeapInuse: 663552 HeapObjects: 1887 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 45 @0.039s 12%: 0.008+0.22+0.087 ms clock, 0.13+0/0.097/0.31+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 46 @0.040s 12%: 0.011+0.20+0.077 ms clock, 0.17+0/0.25/0.20+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 298264 HeapInuse: 663552 HeapObjects: 1887 HeapIdle 188416 HeapReleased 0 HeapSys 851968
gc 47 @0.042s 12%: 0.009+0.20+0.085 ms clock, 0.15+0/0.29/0.16+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 48 @0.043s 12%: 0.009+0.19+0.084 ms clock, 0.15+0/0.34/0.21+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 299928 HeapInuse: 663552 HeapObjects: 1890 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 49 @0.045s 11%: 0.013+0.13+0.083 ms clock, 0.20+0/0.21/0.24+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 50 @0.045s 12%: 0.006+0.17+0.090 ms clock, 0.10+0/0.26/0.19+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 299928 HeapInuse: 663552 HeapObjects: 1890 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 51 @0.048s 11%: 0.008+0.16+0.10 ms clock, 0.13+0/0.27/0.20+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 52 @0.048s 11%: 0.013+0.18+0.088 ms clock, 0.21+0/0.27/0.30+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 300024 HeapInuse: 663552 HeapObjects: 1891 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 53 @0.050s 11%: 0.008+0.27+0.091 ms clock, 0.14+0/0.25/0.18+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 54 @0.051s 11%: 0.013+0.15+0.090 ms clock, 0.21+0/0.24/0.27+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 300120 HeapInuse: 663552 HeapObjects: 1892 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 55 @0.053s 11%: 0.008+0.16+0.088 ms clock, 0.13+0/0.24/0.18+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 56 @0.054s 11%: 0.019+0.14+0.094 ms clock, 0.30+0/0.21/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 300216 HeapInuse: 663552 HeapObjects: 1893 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 57 @0.056s 11%: 0.028+0.16+0.11 ms clock, 0.46+0/0.20/0.22+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 58 @0.056s 11%: 0.017+0.15+0.10 ms clock, 0.28+0/0.22/0.18+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 300216 HeapInuse: 663552 HeapObjects: 1893 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 59 @0.059s 11%: 0.009+0.18+0.098 ms clock, 0.15+0/0.23/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 60 @0.059s 11%: 0.012+0.14+0.097 ms clock, 0.19+0/0.27/0.15+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 300408 HeapInuse: 663552 HeapObjects: 1895 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 61 @0.062s 11%: 0.011+0.17+0.088 ms clock, 0.18+0/0.12/0.40+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 62 @0.062s 11%: 0.011+0.15+0.10 ms clock, 0.18+0/0.18/0.24+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 300408 HeapInuse: 663552 HeapObjects: 1895 HeapIdle 122880 HeapReleased 0 HeapSys 786432
gc 63 @0.064s 11%: 0.008+0.17+0.083 ms clock, 0.13+0/0.32/0.18+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	      2342 ns/op	      90 B/op	       2 allocs/op
gc 64 @0.065s 11%: 0.009+0.15+0.086 ms clock, 0.15+0/0.23/0.30+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 396776 HeapInuse: 761856 HeapObjects: 3398 HeapIdle 24576 HeapReleased 0 HeapSys 786432
gc 65 @0.069s 10%: 0.008+0.16+0.11 ms clock, 0.14+0/0.22/0.23+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetErase1E3-16                           	gc 66 @0.069s 10%: 0.013+0.14+0.085 ms clock, 0.20+0/0.23/0.18+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 396904 HeapInuse: 761856 HeapObjects: 3400 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 67 @0.073s 10%: 0.011+0.18+0.10 ms clock, 0.18+0/0.27/0.20+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 68 @0.073s 10%: 0.008+0.17+0.095 ms clock, 0.13+0/0.25/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 396888 HeapInuse: 761856 HeapObjects: 3399 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 69 @0.078s 10%: 0.011+0.17+0.10 ms clock, 0.18+0/0.25/0.17+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 70 @0.078s 10%: 0.008+0.13+0.085 ms clock, 0.13+0/0.20/0.20+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 396888 HeapInuse: 761856 HeapObjects: 3399 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 71 @0.082s 9%: 0.009+0.13+0.11 ms clock, 0.14+0/0.22/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 72 @0.082s 9%: 0.009+0.23+0.073 ms clock, 0.15+0/0.23/0.18+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 396888 HeapInuse: 761856 HeapObjects: 3399 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 73 @0.086s 9%: 0.008+0.15+0.069 ms clock, 0.12+0/0.26/0.19+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 74 @0.086s 9%: 0.010+0.13+0.066 ms clock, 0.16+0/0.24/0.24+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 396984 HeapInuse: 761856 HeapObjects: 3400 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 75 @0.090s 9%: 0.007+0.13+0.082 ms clock, 0.12+0/0.19/0.19+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 76 @0.090s 9%: 0.008+0.14+0.064 ms clock, 0.13+0/0.24/0.26+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397176 HeapInuse: 761856 HeapObjects: 3402 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 77 @0.094s 9%: 0.008+0.17+0.10 ms clock, 0.13+0/0.22/0.16+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 78 @0.095s 9%: 0.010+0.16+0.15 ms clock, 0.17+0/0.25/0.20+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397368 HeapInuse: 761856 HeapObjects: 3404 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 79 @0.099s 9%: 0.011+0.18+0.060 ms clock, 0.18+0/0.20/0.23+0.97 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 80 @0.100s 9%: 0.022+0.14+0.067 ms clock, 0.36+0/0.27/0.19+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397464 HeapInuse: 761856 HeapObjects: 3405 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 81 @0.103s 8%: 0.008+0.15+0.088 ms clock, 0.12+0/0.14/0.21+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 82 @0.104s 8%: 0.012+0.16+0.11 ms clock, 0.20+0/0.32/0.20+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397560 HeapInuse: 761856 HeapObjects: 3406 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 83 @0.107s 8%: 0.009+0.26+0.077 ms clock, 0.14+0/0.46/0.13+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 84 @0.108s 8%: 0.008+0.15+0.10 ms clock, 0.12+0/0.21/0.17+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397656 HeapInuse: 761856 HeapObjects: 3407 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 85 @0.112s 8%: 0.008+0.18+0.088 ms clock, 0.12+0/0.20/0.31+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 86 @0.112s 8%: 0.009+0.16+0.072 ms clock, 0.15+0/0.28/0.16+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397848 HeapInuse: 761856 HeapObjects: 3409 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 87 @0.116s 8%: 0.007+0.14+0.11 ms clock, 0.12+0/0.21/0.25+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 88 @0.116s 8%: 0.010+0.21+0.088 ms clock, 0.16+0/0.26/0.27+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 397944 HeapInuse: 761856 HeapObjects: 3410 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 89 @0.120s 8%: 0.013+0.17+0.12 ms clock, 0.21+0/0.31/0.23+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 90 @0.121s 8%: 0.010+0.30+0.080 ms clock, 0.16+0/0.31/0.21+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 398040 HeapInuse: 761856 HeapObjects: 3411 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 91 @0.124s 8%: 0.008+0.20+0.10 ms clock, 0.13+0/0.29/0.28+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 92 @0.125s 8%: 0.011+0.15+0.065 ms clock, 0.17+0/0.23/0.15+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 398040 HeapInuse: 761856 HeapObjects: 3411 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 93 @0.129s 8%: 0.008+0.13+0.076 ms clock, 0.12+0/0.25/0.16+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 94 @0.129s 8%: 0.007+0.15+0.080 ms clock, 0.12+0/0.24/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 398136 HeapInuse: 761856 HeapObjects: 3412 HeapIdle 1073152 HeapReleased 0 HeapSys 1835008
gc 95 @0.133s 8%: 0.008+0.17+0.088 ms clock, 0.13+0/0.23/0.18+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	      1745 ns/op	      99 B/op	       2 allocs/op
gc 96 @0.133s 8%: 0.014+0.17+0.12 ms clock, 0.23+0/0.26/0.19+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310216 HeapInuse: 671744 HeapObjects: 1913 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 97 @0.135s 8%: 0.010+0.14+0.10 ms clock, 0.16+0/0.24/0.21+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E3-16                  	gc 98 @0.136s 8%: 0.009+0.15+0.094 ms clock, 0.15+0/0.16/0.22+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310248 HeapInuse: 671744 HeapObjects: 1914 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 99 @0.137s 8%: 0.008+0.18+0.11 ms clock, 0.13+0/0.16/0.33+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 100 @0.138s 8%: 0.014+0.13+0.084 ms clock, 0.22+0/0.26/0.22+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310232 HeapInuse: 671744 HeapObjects: 1913 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 101 @0.139s 8%: 0.008+0.16+0.10 ms clock, 0.12+0/0.23/0.16+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 102 @0.140s 8%: 0.011+0.14+0.087 ms clock, 0.18+0/0.21/0.22+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310328 HeapInuse: 671744 HeapObjects: 1914 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 103 @0.142s 8%: 0.008+0.16+0.078 ms clock, 0.13+0/0.26/0.22+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 104 @0.142s 8%: 0.010+0.16+0.093 ms clock, 0.16+0/0.28/0.23+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310520 HeapInuse: 671744 HeapObjects: 1916 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 105 @0.144s 8%: 0.008+0.15+0.093 ms clock, 0.12+0/0.23/0.18+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 106 @0.144s 8%: 0.010+0.15+0.099 ms clock, 0.17+0/0.26/0.27+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310712 HeapInuse: 671744 HeapObjects: 1918 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 107 @0.146s 8%: 0.022+0.25+0.089 ms clock, 0.36+0/0.28/0.21+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 108 @0.146s 8%: 0.009+0.16+0.11 ms clock, 0.15+0/0.25/0.32+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310776 HeapInuse: 671744 HeapObjects: 1919 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 109 @0.148s 8%: 0.008+0.15+0.094 ms clock, 0.13+0/0.23/0.24+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 110 @0.149s 8%: 0.012+0.15+0.097 ms clock, 0.20+0/0.26/0.25+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 310968 HeapInuse: 671744 HeapObjects: 1921 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 111 @0.151s 8%: 0.009+0.14+0.073 ms clock, 0.14+0/0.25/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 112 @0.151s 8%: 0.008+0.16+0.092 ms clock, 0.13+0/0.27/0.25+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311224 HeapInuse: 671744 HeapObjects: 1924 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 113 @0.153s 8%: 0.008+0.17+0.078 ms clock, 0.13+0/0.19/0.22+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 114 @0.153s 8%: 0.009+0.15+0.10 ms clock, 0.15+0/0.25/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311224 HeapInuse: 671744 HeapObjects: 1924 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 115 @0.155s 8%: 0.008+0.13+0.078 ms clock, 0.14+0/0.24/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 116 @0.155s 8%: 0.010+0.17+0.11 ms clock, 0.16+0/0.19/0.29+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311224 HeapInuse: 671744 HeapObjects: 1924 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 117 @0.157s 8%: 0.018+0.15+0.085 ms clock, 0.29+0/0.24/0.23+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 118 @0.158s 8%: 0.029+0.21+0.14 ms clock, 0.46+0/0.21/0.33+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311416 HeapInuse: 671744 HeapObjects: 1926 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 119 @0.159s 8%: 0.009+0.14+0.090 ms clock, 0.15+0/0.24/0.19+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 120 @0.160s 8%: 0.025+0.22+0.096 ms clock, 0.40+0/0.12/0.29+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311512 HeapInuse: 671744 HeapObjects: 1927 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 121 @0.162s 8%: 0.008+0.20+0.071 ms clock, 0.13+0/0.30/0.16+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 122 @0.162s 8%: 0.010+0.31+0.074 ms clock, 0.17+0/0.36/0.17+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311512 HeapInuse: 671744 HeapObjects: 1927 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 123 @0.164s 8%: 0.008+0.34+0.070 ms clock, 0.13+0/0.50/0.10+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 124 @0.165s 8%: 0.009+0.18+0.086 ms clock, 0.14+0/0.30/0.12+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311704 HeapInuse: 671744 HeapObjects: 1929 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 125 @0.166s 8%: 0.008+0.15+0.088 ms clock, 0.13+0/0.21/0.28+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 126 @0.167s 8%: 0.012+0.15+0.066 ms clock, 0.20+0/0.27/0.25+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 311896 HeapInuse: 671744 HeapObjects: 1931 HeapIdle 1163264 HeapReleased 0 HeapSys 1835008
gc 127 @0.168s 8%: 0.009+0.20+0.11 ms clock, 0.15+0/0.27/0.24+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	      1703 ns/op	      99 B/op	       2 allocs/op
gc 128 @0.169s 8%: 0.020+0.20+0.090 ms clock, 0.32+0/0.20/0.39+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274080 HeapInuse: 638976 HeapObjects: 1437 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 129 @0.171s 8%: 0.009+0.23+0.085 ms clock, 0.14+0/0.25/0.48+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E3-16          	gc 130 @0.172s 8%: 0.008+0.16+0.081 ms clock, 0.13+0/0.22/0.33+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274224 HeapInuse: 638976 HeapObjects: 1439 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 131 @0.174s 8%: 0.009+0.16+0.094 ms clock, 0.14+0/0.27/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 132 @0.174s 8%: 0.011+0.45+0.39 ms clock, 0.19+0/0.56/0.35+6.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274304 HeapInuse: 638976 HeapObjects: 1439 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 133 @0.178s 8%: 0.015+0.17+0.069 ms clock, 0.25+0/0.25/0.27+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 134 @0.178s 8%: 0.006+0.19+0.067 ms clock, 0.10+0/0.18/0.32+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274496 HeapInuse: 638976 HeapObjects: 1441 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 135 @0.180s 8%: 0.009+0.14+0.087 ms clock, 0.15+0/0.24/0.23+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 136 @0.181s 8%: 0.023+0.23+0.066 ms clock, 0.38+0/0.13/0.31+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274592 HeapInuse: 638976 HeapObjects: 1442 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 137 @0.183s 8%: 0.008+0.17+0.091 ms clock, 0.12+0/0.31/0.19+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 138 @0.183s 8%: 0.014+0.19+0.079 ms clock, 0.22+0/0.27/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274592 HeapInuse: 638976 HeapObjects: 1442 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 139 @0.185s 8%: 0.010+0.17+0.081 ms clock, 0.16+0/0.27/0.23+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 140 @0.186s 8%: 0.012+0.16+0.20 ms clock, 0.19+0/0.25/0.33+3.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274688 HeapInuse: 638976 HeapObjects: 1443 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 141 @0.188s 8%: 0.009+0.14+0.078 ms clock, 0.15+0/0.22/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 142 @0.188s 8%: 0.008+0.22+0.082 ms clock, 0.12+0/0.22/0.16+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274880 HeapInuse: 638976 HeapObjects: 1445 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 143 @0.191s 8%: 0.010+0.16+0.074 ms clock, 0.17+0/0.32/0.32+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 144 @0.191s 8%: 0.010+0.20+0.10 ms clock, 0.16+0/0.31/0.19+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274880 HeapInuse: 638976 HeapObjects: 1445 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 145 @0.193s 8%: 0.011+0.17+0.095 ms clock, 0.18+0/0.29/0.21+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 146 @0.194s 8%: 0.013+0.19+0.071 ms clock, 0.22+0/0.33/0.20+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 274976 HeapInuse: 638976 HeapObjects: 1446 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 147 @0.196s 8%: 0.010+0.18+0.074 ms clock, 0.16+0/0.24/0.28+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 148 @0.196s 8%: 0.024+0.18+0.084 ms clock, 0.38+0/0.27/0.39+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275072 HeapInuse: 638976 HeapObjects: 1447 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 149 @0.198s 8%: 0.011+0.16+0.075 ms clock, 0.18+0/0.30/0.16+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 150 @0.199s 8%: 0.011+0.15+0.071 ms clock, 0.17+0/0.25/0.23+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275168 HeapInuse: 638976 HeapObjects: 1448 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 151 @0.201s 8%: 0.009+0.18+0.083 ms clock, 0.15+0/0.20/0.22+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 152 @0.201s 8%: 0.010+0.18+0.087 ms clock, 0.16+0/0.31/0.23+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275168 HeapInuse: 638976 HeapObjects: 1448 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 153 @0.204s 8%: 0.009+0.21+0.10 ms clock, 0.15+0/0.16/0.35+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 154 @0.204s 9%: 0.015+0.15+1.2 ms clock, 0.25+0/0.27/0.25+19 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275360 HeapInuse: 638976 HeapObjects: 1450 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 155 @0.208s 9%: 0.010+0.20+0.097 ms clock, 0.16+0/0.094/0.45+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 156 @0.208s 9%: 0.010+0.18+0.070 ms clock, 0.16+0/0.33/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275360 HeapInuse: 638976 HeapObjects: 1450 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 157 @0.211s 9%: 0.009+0.18+0.091 ms clock, 0.15+0/0.18/0.24+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 158 @0.211s 9%: 0.012+0.29+0.096 ms clock, 0.20+0/0.25/0.18+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275456 HeapInuse: 638976 HeapObjects: 1451 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 159 @0.214s 9%: 0.012+0.13+0.080 ms clock, 0.19+0/0.22/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 160 @0.214s 9%: 0.010+0.15+0.084 ms clock, 0.17+0/0.21/0.21+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 275552 HeapInuse: 638976 HeapObjects: 1452 HeapIdle 1196032 HeapReleased 0 HeapSys 1835008
gc 161 @0.216s 9%: 0.007+0.15+0.084 ms clock, 0.12+0/0.25/0.21+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	      1995 ns/op	      61 B/op	       1 allocs/op
gc 162 @0.216s 9%: 0.015+0.17+0.082 ms clock, 0.25+0/0.27/0.25+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 269824 HeapInuse: 696320 HeapObjects: 510 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 163 @0.217s 9%: 0.009+0.20+0.089 ms clock, 0.15+0/0.28/0.26+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E3-16                   	gc 164 @0.218s 9%: 0.011+0.17+0.11 ms clock, 0.18+0/0.30/0.20+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 270912 HeapInuse: 696320 HeapObjects: 520 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 165 @0.219s 9%: 0.015+0.16+0.086 ms clock, 0.25+0/0.24/0.35+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 166 @0.219s 9%: 0.008+0.18+0.088 ms clock, 0.13+0/0.16/0.30+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 270200 HeapInuse: 696320 HeapObjects: 512 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 167 @0.220s 9%: 0.011+0.15+0.083 ms clock, 0.19+0/0.22/0.29+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 168 @0.221s 9%: 0.008+0.12+0.064 ms clock, 0.13+0/0.25/0.18+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271848 HeapInuse: 696320 HeapObjects: 531 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 169 @0.221s 9%: 0.008+0.17+0.087 ms clock, 0.12+0/0.32/0.29+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 170 @0.222s 9%: 0.011+0.16+0.083 ms clock, 0.18+0/0.27/0.25+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271248 HeapInuse: 696320 HeapObjects: 525 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 171 @0.223s 9%: 0.008+0.19+0.099 ms clock, 0.14+0/0.26/0.32+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 172 @0.223s 9%: 0.007+0.17+0.086 ms clock, 0.12+0/0.23/0.27+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271240 HeapInuse: 696320 HeapObjects: 524 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 173 @0.224s 9%: 0.011+0.18+0.13 ms clock, 0.18+0/0.34/0.23+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 174 @0.225s 9%: 0.012+0.16+0.082 ms clock, 0.20+0/0.30/0.28+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271864 HeapInuse: 696320 HeapObjects: 534 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 175 @0.226s 9%: 0.011+0.20+0.091 ms clock, 0.18+0/0.38/0.21+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 176 @0.226s 9%: 0.008+0.17+0.29 ms clock, 0.13+0/0.26/0.30+4.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 269736 HeapInuse: 696320 HeapObjects: 506 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 177 @0.227s 9%: 0.009+0.20+0.083 ms clock, 0.15+0/0.24/0.33+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 178 @0.228s 9%: 0.006+0.19+0.068 ms clock, 0.098+0/0.22/0.22+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271136 HeapInuse: 696320 HeapObjects: 522 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 179 @0.228s 9%: 0.010+0.18+0.081 ms clock, 0.16+0/0.27/0.25+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 180 @0.229s 9%: 0.008+0.17+0.072 ms clock, 0.13+0/0.24/0.26+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271344 HeapInuse: 696320 HeapObjects: 526 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 181 @0.230s 9%: 0.008+0.17+0.071 ms clock, 0.13+0/0.24/0.28+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 182 @0.230s 9%: 0.010+0.15+0.087 ms clock, 0.16+0/0.30/0.34+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271280 HeapInuse: 696320 HeapObjects: 526 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 183 @0.231s 9%: 0.007+0.15+0.087 ms clock, 0.12+0/0.25/0.27+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 184 @0.231s 9%: 0.008+0.14+0.077 ms clock, 0.13+0/0.29/0.20+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271448 HeapInuse: 696320 HeapObjects: 529 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 185 @0.232s 9%: 0.009+0.17+0.076 ms clock, 0.15+0/0.20/0.33+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 186 @0.233s 9%: 0.010+0.15+0.077 ms clock, 0.16+0/0.28/0.20+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 272544 HeapInuse: 696320 HeapObjects: 537 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 187 @0.233s 9%: 0.009+0.18+0.077 ms clock, 0.14+0/0.26/0.32+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 188 @0.234s 9%: 0.007+0.24+0.077 ms clock, 0.11+0/0.23/0.20+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 272112 HeapInuse: 696320 HeapObjects: 535 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 189 @0.235s 9%: 0.007+0.15+0.073 ms clock, 0.12+0/0.23/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 190 @0.235s 9%: 0.018+0.17+0.080 ms clock, 0.28+0/0.27/0.31+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 271376 HeapInuse: 696320 HeapObjects: 523 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 191 @0.236s 10%: 0.009+0.20+0.10 ms clock, 0.15+0/0.37/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 192 @0.237s 10%: 0.010+0.20+0.087 ms clock, 0.17+0/0.36/0.17+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 273304 HeapInuse: 696320 HeapObjects: 547 HeapIdle 1138688 HeapReleased 0 HeapSys 1835008
gc 193 @0.237s 10%: 0.009+0.15+0.083 ms clock, 0.15+0/0.29/0.21+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	       747 ns/op	      57 B/op	       0 allocs/op
gc 194 @0.238s 10%: 0.011+0.14+0.090 ms clock, 0.18+0/0.32/0.27+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 280880 HeapInuse: 704512 HeapObjects: 541 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 195 @0.239s 10%: 0.015+0.17+0.097 ms clock, 0.24+0/0.27/0.33+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E3-16                    	gc 196 @0.240s 10%: 0.013+0.15+0.072 ms clock, 0.21+0/0.24/0.29+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 279968 HeapInuse: 704512 HeapObjects: 529 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 197 @0.241s 10%: 0.009+0.14+0.074 ms clock, 0.14+0/0.27/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 198 @0.241s 10%: 0.009+0.17+0.070 ms clock, 0.14+0/0.22/0.22+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281792 HeapInuse: 704512 HeapObjects: 553 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 199 @0.242s 10%: 0.008+0.14+0.093 ms clock, 0.12+0/0.26/0.31+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 200 @0.242s 10%: 0.009+0.43+0.085 ms clock, 0.14+0/0.57/0.24+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281160 HeapInuse: 704512 HeapObjects: 543 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 201 @0.244s 10%: 0.009+0.14+0.073 ms clock, 0.15+0/0.31/0.18+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 202 @0.244s 10%: 0.007+0.16+0.077 ms clock, 0.11+0/0.22/0.32+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 280736 HeapInuse: 704512 HeapObjects: 539 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 203 @0.245s 10%: 0.009+0.20+0.069 ms clock, 0.14+0/0.29/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 204 @0.246s 10%: 0.016+0.19+0.069 ms clock, 0.26+0/0.30/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281264 HeapInuse: 704512 HeapObjects: 546 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 205 @0.247s 10%: 0.009+0.16+0.074 ms clock, 0.14+0/0.34/0.31+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 206 @0.247s 10%: 0.007+0.20+0.077 ms clock, 0.12+0/0.22/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 280376 HeapInuse: 704512 HeapObjects: 534 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 207 @0.248s 10%: 0.009+0.17+0.073 ms clock, 0.14+0/0.28/0.39+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 208 @0.248s 10%: 0.008+0.15+0.074 ms clock, 0.14+0/0.25/0.26+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 280872 HeapInuse: 704512 HeapObjects: 540 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 209 @0.249s 10%: 0.008+0.14+0.070 ms clock, 0.14+0/0.22/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 210 @0.250s 10%: 0.009+0.15+0.11 ms clock, 0.15+0/0.25/0.27+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 280536 HeapInuse: 704512 HeapObjects: 535 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 211 @0.251s 10%: 0.011+0.16+0.068 ms clock, 0.18+0/0.29/0.19+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 212 @0.251s 10%: 0.007+0.20+0.090 ms clock, 0.12+0/0.28/0.28+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281472 HeapInuse: 704512 HeapObjects: 547 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 213 @0.252s 10%: 0.008+0.16+0.074 ms clock, 0.13+0/0.26/0.33+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 214 @0.252s 10%: 0.010+0.19+0.067 ms clock, 0.16+0/0.26/0.31+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281168 HeapInuse: 704512 HeapObjects: 545 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 215 @0.253s 10%: 0.009+0.20+0.080 ms clock, 0.15+0/0.25/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 216 @0.254s 10%: 0.014+0.16+0.099 ms clock, 0.22+0/0.24/0.21+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 280864 HeapInuse: 704512 HeapObjects: 540 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 217 @0.255s 10%: 0.008+0.16+0.082 ms clock, 0.12+0/0.29/0.31+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 218 @0.255s 10%: 0.008+0.17+0.069 ms clock, 0.13+0/0.24/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281688 HeapInuse: 704512 HeapObjects: 548 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 219 @0.256s 10%: 0.008+0.19+0.080 ms clock, 0.14+0/0.24/0.29+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 220 @0.257s 10%: 0.010+0.18+0.089 ms clock, 0.16+0/0.29/0.29+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 282104 HeapInuse: 704512 HeapObjects: 557 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 221 @0.258s 10%: 0.012+0.15+0.13 ms clock, 0.19+0/0.26/0.20+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 222 @0.258s 10%: 0.011+0.15+0.086 ms clock, 0.17+0/0.28/0.27+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 281968 HeapInuse: 704512 HeapObjects: 552 HeapIdle 1130496 HeapReleased 0 HeapSys 1835008
gc 223 @0.259s 10%: 0.009+0.19+0.064 ms clock, 0.15+0/0.29/0.33+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	       504 ns/op	      10 B/op	       0 allocs/op
gc 224 @0.260s 10%: 0.009+0.14+0.088 ms clock, 0.15+0/0.22/0.25+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 225544 HeapInuse: 630784 HeapObjects: 504 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 225 @0.260s 10%: 0.009+0.17+0.072 ms clock, 0.15+0/0.28/0.29+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E3-16           	gc 226 @0.261s 10%: 0.015+0.16+0.11 ms clock, 0.25+0/0.33/0.35+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 225800 HeapInuse: 630784 HeapObjects: 508 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 227 @0.261s 10%: 0.011+0.21+0.082 ms clock, 0.19+0/0.24/0.29+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 228 @0.262s 10%: 0.011+0.31+0.073 ms clock, 0.18+0/0.37/0.20+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226096 HeapInuse: 630784 HeapObjects: 511 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 229 @0.263s 10%: 0.007+0.18+0.081 ms clock, 0.12+0/0.30/0.14+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 230 @0.263s 10%: 0.011+0.20+0.075 ms clock, 0.18+0/0.32/0.22+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226024 HeapInuse: 622592 HeapObjects: 513 HeapIdle 1212416 HeapReleased 0 HeapSys 1835008
gc 231 @0.264s 10%: 0.008+0.14+0.085 ms clock, 0.13+0/0.29/0.20+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 232 @0.264s 10%: 0.010+0.14+0.079 ms clock, 0.16+0/0.29/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226608 HeapInuse: 630784 HeapObjects: 518 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 233 @0.265s 10%: 0.009+0.15+0.072 ms clock, 0.14+0/0.23/0.31+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 234 @0.265s 10%: 0.009+0.15+0.13 ms clock, 0.15+0/0.28/0.23+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226784 HeapInuse: 630784 HeapObjects: 519 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 235 @0.266s 10%: 0.013+0.30+0.079 ms clock, 0.21+0/0.14/0.43+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 236 @0.267s 10%: 0.010+1.2+0.098 ms clock, 0.16+0/0.32/0.20+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226160 HeapInuse: 622592 HeapObjects: 512 HeapIdle 1212416 HeapReleased 0 HeapSys 1835008
gc 237 @0.268s 10%: 0.011+0.19+0.081 ms clock, 0.18+0/0.25/0.34+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 238 @0.269s 10%: 0.013+0.16+0.082 ms clock, 0.21+0/0.27/0.21+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226464 HeapInuse: 622592 HeapObjects: 516 HeapIdle 1212416 HeapReleased 0 HeapSys 1835008
gc 239 @0.269s 10%: 0.011+0.17+0.076 ms clock, 0.18+0/0.32/0.23+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 240 @0.270s 10%: 0.012+0.17+0.10 ms clock, 0.19+0/0.24/0.38+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 227056 HeapInuse: 630784 HeapObjects: 521 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 241 @0.271s 10%: 0.013+0.18+0.075 ms clock, 0.22+0/0.33/0.33+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 242 @0.271s 10%: 0.007+0.18+0.11 ms clock, 0.12+0/0.32/0.30+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226672 HeapInuse: 630784 HeapObjects: 517 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 243 @0.272s 10%: 0.013+0.20+0.082 ms clock, 0.21+0/0.34/0.25+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 244 @0.273s 10%: 0.008+0.17+0.10 ms clock, 0.13+0/0.24/0.26+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 227112 HeapInuse: 630784 HeapObjects: 524 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 245 @0.273s 10%: 0.009+0.16+0.084 ms clock, 0.14+0/0.23/0.29+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 246 @0.274s 10%: 0.010+0.23+0.075 ms clock, 0.16+0/0.22/0.19+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 226640 HeapInuse: 622592 HeapObjects: 517 HeapIdle 1212416 HeapReleased 0 HeapSys 1835008
gc 247 @0.274s 10%: 0.010+0.16+0.079 ms clock, 0.16+0/0.27/0.22+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 248 @0.275s 10%: 0.012+0.29+0.089 ms clock, 0.19+0/0.29/0.20+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 227208 HeapInuse: 630784 HeapObjects: 521 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 249 @0.276s 10%: 0.009+0.15+0.070 ms clock, 0.14+0/0.25/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 250 @0.276s 10%: 0.008+0.15+0.081 ms clock, 0.13+0/0.25/0.22+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 227168 HeapInuse: 630784 HeapObjects: 523 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 251 @0.277s 10%: 0.008+0.15+0.073 ms clock, 0.13+0/0.24/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 252 @0.277s 10%: 0.008+0.37+0.088 ms clock, 0.13+0/0.49/0.26+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 227264 HeapInuse: 630784 HeapObjects: 524 HeapIdle 1204224 HeapReleased 0 HeapSys 1835008
gc 253 @0.278s 11%: 0.009+0.16+0.084 ms clock, 0.14+0/0.30/0.22+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	       575 ns/op	       7 B/op	       0 allocs/op
gc 254 @0.278s 11%: 0.009+0.15+0.096 ms clock, 0.15+0/0.27/0.27+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 247608 HeapInuse: 614400 HeapObjects: 499 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 255 @0.279s 11%: 0.008+0.15+0.078 ms clock, 0.13+0/0.27/0.16+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E3-16    	gc 256 @0.279s 11%: 0.009+0.14+0.060 ms clock, 0.15+0/0.23/0.19+0.97 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 247848 HeapInuse: 614400 HeapObjects: 502 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 257 @0.280s 11%: 0.012+0.24+0.072 ms clock, 0.20+0/0.27/0.16+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 258 @0.280s 11%: 0.009+0.13+0.079 ms clock, 0.14+0/0.26/0.28+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 247832 HeapInuse: 614400 HeapObjects: 501 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 259 @0.281s 11%: 0.007+0.15+0.072 ms clock, 0.12+0/0.28/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 260 @0.281s 11%: 0.010+0.25+0.067 ms clock, 0.16+0/0.24/0.19+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 247928 HeapInuse: 614400 HeapObjects: 502 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 261 @0.282s 11%: 0.007+0.18+0.078 ms clock, 0.12+0/0.26/0.26+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 262 @0.282s 11%: 0.012+0.18+0.075 ms clock, 0.19+0/0.28/0.29+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 247928 HeapInuse: 614400 HeapObjects: 502 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 263 @0.283s 11%: 0.007+0.16+0.069 ms clock, 0.12+0/0.20/0.22+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 264 @0.283s 11%: 0.008+0.18+0.094 ms clock, 0.13+0/0.30/0.25+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248024 HeapInuse: 614400 HeapObjects: 503 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 265 @0.284s 11%: 0.009+0.14+0.079 ms clock, 0.14+0/0.26/0.26+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 266 @0.284s 11%: 0.013+0.18+0.069 ms clock, 0.20+0/0.11/0.43+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248024 HeapInuse: 614400 HeapObjects: 503 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 267 @0.285s 11%: 0.008+0.14+0.098 ms clock, 0.14+0/0.27/0.22+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 268 @0.285s 11%: 0.011+0.16+0.075 ms clock, 0.19+0/0.24/0.28+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248120 HeapInuse: 614400 HeapObjects: 504 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 269 @0.286s 11%: 0.009+0.18+0.075 ms clock, 0.15+0/0.34/0.30+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 270 @0.286s 11%: 0.008+0.23+0.069 ms clock, 0.13+0/0.22/0.20+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248120 HeapInuse: 614400 HeapObjects: 504 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 271 @0.287s 11%: 0.008+0.15+0.071 ms clock, 0.12+0/0.10/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 272 @0.287s 11%: 0.009+0.17+0.087 ms clock, 0.14+0/0.29/0.12+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248120 HeapInuse: 614400 HeapObjects: 504 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 273 @0.288s 11%: 0.012+0.17+0.083 ms clock, 0.19+0/0.30/0.25+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 274 @0.289s 11%: 0.007+0.16+0.074 ms clock, 0.12+0/0.28/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248120 HeapInuse: 614400 HeapObjects: 504 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 275 @0.289s 11%: 0.010+0.22+0.11 ms clock, 0.16+0/0.24/0.28+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 276 @0.290s 11%: 0.009+0.19+0.078 ms clock, 0.15+0/0.31/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248216 HeapInuse: 614400 HeapObjects: 505 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 277 @0.290s 11%: 0.012+0.20+0.094 ms clock, 0.19+0/0.36/0.20+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 278 @0.291s 11%: 0.013+0.15+0.078 ms clock, 0.21+0/0.27/0.21+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248312 HeapInuse: 614400 HeapObjects: 506 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 279 @0.291s 11%: 0.009+0.41+0.077 ms clock, 0.14+0/0.22/0.35+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 280 @0.292s 11%: 0.008+0.16+0.063 ms clock, 0.13+0/0.25/0.31+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248504 HeapInuse: 614400 HeapObjects: 508 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 281 @0.293s 11%: 0.009+0.15+0.086 ms clock, 0.14+0/0.20/0.27+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 282 @0.293s 11%: 0.011+0.16+0.078 ms clock, 0.17+0/0.30/0.27+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 248696 HeapInuse: 614400 HeapObjects: 510 HeapIdle 1220608 HeapReleased 0 HeapSys 1835008
gc 283 @0.293s 11%: 0.007+0.17+0.062 ms clock, 0.12+0/0.31/0.17+0.99 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
    1000	       458 ns/op	      28 B/op	       0 allocs/op
gc 284 @0.294s 11%: 0.010+0.20+0.089 ms clock, 0.16+0/0.10/0.34+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 305304 HeapInuse: 663552 HeapObjects: 511 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 285 @0.297s 11%: 0.008+0.18+0.099 ms clock, 0.13+0/0.30/0.25+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSort1E4-16                               	gc 286 @0.297s 11%: 0.035+0.17+0.073 ms clock, 0.57+0/0.29/0.27+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 305528 HeapInuse: 663552 HeapObjects: 514 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 287 @0.300s 11%: 0.010+0.18+0.37 ms clock, 0.17+0/0.29/0.30+5.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 288 @0.300s 11%: 0.010+0.20+0.066 ms clock, 0.16+0/0.35/0.25+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 305704 HeapInuse: 663552 HeapObjects: 515 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 289 @0.303s 11%: 0.009+0.16+0.087 ms clock, 0.15+0/0.29/0.21+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 290 @0.304s 11%: 0.011+0.21+0.074 ms clock, 0.18+0/0.33/0.18+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 305704 HeapInuse: 663552 HeapObjects: 515 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 291 @0.306s 11%: 0.009+0.18+0.086 ms clock, 0.15+0/0.13/0.36+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 292 @0.307s 11%: 0.009+0.17+0.086 ms clock, 0.15+0/0.17/0.27+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 305896 HeapInuse: 663552 HeapObjects: 517 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 293 @0.309s 11%: 0.008+0.14+0.095 ms clock, 0.13+0/0.28/0.22+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 294 @0.310s 11%: 0.012+0.18+0.097 ms clock, 0.20+0/0.29/0.19+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306088 HeapInuse: 663552 HeapObjects: 519 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 295 @0.312s 11%: 0.010+0.17+0.069 ms clock, 0.17+0/0.18/0.34+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 296 @0.313s 11%: 0.013+0.25+0.086 ms clock, 0.21+0/0.083/0.34+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306184 HeapInuse: 663552 HeapObjects: 520 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 297 @0.315s 11%: 0.011+0.21+0.061 ms clock, 0.18+0/0.29/0.34+0.98 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 298 @0.316s 11%: 0.008+0.17+0.062 ms clock, 0.13+0/0.27/0.26+0.99 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306184 HeapInuse: 663552 HeapObjects: 520 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 299 @0.319s 11%: 0.008+0.24+0.090 ms clock, 0.13+0/0.24/0.14+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 300 @0.319s 11%: 0.006+0.22+0.10 ms clock, 0.10+0/0.15/0.28+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306184 HeapInuse: 663552 HeapObjects: 520 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 301 @0.322s 11%: 0.008+0.19+0.071 ms clock, 0.12+0/0.19/0.26+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 302 @0.322s 11%: 0.010+0.17+0.10 ms clock, 0.16+0/0.19/0.24+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306184 HeapInuse: 663552 HeapObjects: 520 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 303 @0.325s 11%: 0.009+0.16+0.096 ms clock, 0.15+0/0.28/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 304 @0.325s 11%: 0.011+0.24+0.065 ms clock, 0.18+0/0.26/0.38+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306184 HeapInuse: 663552 HeapObjects: 520 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 305 @0.328s 11%: 0.008+0.19+0.065 ms clock, 0.14+0/0.26/0.22+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 306 @0.329s 11%: 0.016+0.17+0.063 ms clock, 0.26+0/0.35/0.26+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306280 HeapInuse: 663552 HeapObjects: 521 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 307 @0.331s 11%: 0.009+0.20+0.11 ms clock, 0.15+0/0.27/0.31+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 308 @0.332s 11%: 0.013+0.15+0.067 ms clock, 0.20+0/0.25/0.33+1.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306376 HeapInuse: 663552 HeapObjects: 522 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 309 @0.334s 11%: 0.010+0.17+0.070 ms clock, 0.16+0/0.26/0.31+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 310 @0.335s 11%: 0.007+0.16+0.078 ms clock, 0.12+0/0.27/0.23+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 306376 HeapInuse: 663552 HeapObjects: 522 HeapIdle 1171456 HeapReleased 0 HeapSys 1835008
gc 311 @0.337s 10%: 0.008+0.18+0.072 ms clock, 0.13+0/0.26/0.26+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
   10000	       249 ns/op	       8 B/op	       0 allocs/op
gc 312 @0.338s 10%: 0.009+1.3+0.069 ms clock, 0.15+0/0.32/0.27+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1104872 HeapInuse: 1458176 HeapObjects: 15526 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 313 @0.363s 10%: 0.039+0.18+0.074 ms clock, 0.62+0/0.19/0.19+1.1 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsert1E4-16                          	gc 314 @0.364s 10%: 0.009+0.19+0.10 ms clock, 0.14+0/0.37/0.20+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1104904 HeapInuse: 1458176 HeapObjects: 15527 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 315 @0.388s 9%: 0.033+0.27+0.10 ms clock, 0.53+0/0.35/0.35+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 316 @0.388s 9%: 0.028+0.18+0.10 ms clock, 0.45+0/0.42/0.093+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1104984 HeapInuse: 1458176 HeapObjects: 15527 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 317 @0.412s 9%: 0.010+0.23+0.096 ms clock, 0.17+0/0.35/0.25+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 318 @0.412s 9%: 0.014+0.24+0.15 ms clock, 0.23+0/0.28/0.31+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105080 HeapInuse: 1458176 HeapObjects: 15528 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 319 @0.436s 8%: 0.011+0.25+0.22 ms clock, 0.18+0/0.46/0.12+3.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 320 @0.437s 8%: 0.024+0.28+0.10 ms clock, 0.38+0/0.27/0.19+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105080 HeapInuse: 1458176 HeapObjects: 15528 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 321 @0.461s 8%: 0.013+0.21+0.12 ms clock, 0.21+0/0.34/0.32+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 322 @0.461s 8%: 0.010+0.18+0.10 ms clock, 0.16+0/0.27/0.20+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105080 HeapInuse: 1458176 HeapObjects: 15528 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 323 @0.486s 8%: 0.009+0.20+0.10 ms clock, 0.15+0/0.36/0.24+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 324 @0.486s 8%: 0.011+0.20+0.12 ms clock, 0.18+0/0.35/0.36+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105080 HeapInuse: 1458176 HeapObjects: 15528 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 325 @0.512s 7%: 0.011+0.18+0.10 ms clock, 0.17+0/0.33/0.14+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 326 @0.513s 7%: 0.013+0.20+0.12 ms clock, 0.20+0/0.28/0.31+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105080 HeapInuse: 1458176 HeapObjects: 15528 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 327 @0.536s 7%: 0.040+0.22+0.17 ms clock, 0.64+0/0.31/0.33+2.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 328 @0.537s 7%: 0.11+0.26+0.13 ms clock, 1.7+0/0.29/0.15+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105176 HeapInuse: 1458176 HeapObjects: 15529 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 329 @0.561s 7%: 0.008+0.16+0.12 ms clock, 0.14+0/0.30/0.18+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 330 @0.561s 7%: 0.022+0.19+0.13 ms clock, 0.36+0/0.29/0.19+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105272 HeapInuse: 1458176 HeapObjects: 15530 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 331 @0.586s 6%: 0.028+0.21+0.099 ms clock, 0.45+0/0.27/0.23+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 332 @0.587s 6%: 0.010+0.20+0.10 ms clock, 0.17+0/0.21/0.34+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105368 HeapInuse: 1458176 HeapObjects: 15531 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 333 @0.611s 6%: 0.058+0.20+0.15 ms clock, 0.93+0/0.35/0.16+2.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 334 @0.611s 6%: 0.020+0.19+0.11 ms clock, 0.32+0/0.25/0.28+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105464 HeapInuse: 1458176 HeapObjects: 15532 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 335 @0.635s 6%: 0.008+0.23+0.11 ms clock, 0.14+0/0.23/0.26+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 336 @0.636s 6%: 0.010+0.19+0.13 ms clock, 0.17+0/0.29/0.23+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105560 HeapInuse: 1458176 HeapObjects: 15533 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 337 @0.659s 6%: 0.023+0.23+0.18 ms clock, 0.37+0/0.34/0.26+2.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 338 @0.660s 6%: 0.027+0.24+0.11 ms clock, 0.44+0/0.21/0.29+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105656 HeapInuse: 1458176 HeapObjects: 15534 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 339 @0.684s 6%: 0.037+0.21+0.12 ms clock, 0.59+0/0.37/0.15+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 340 @0.684s 6%: 0.012+0.22+0.10 ms clock, 0.20+0/0.36/0.22+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105656 HeapInuse: 1458176 HeapObjects: 15534 HeapIdle 376832 HeapReleased 0 HeapSys 1835008
gc 341 @0.708s 5%: 0.010+0.26+0.13 ms clock, 0.16+0/0.36/0.23+2.1 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 342 @0.709s 5%: 0.012+0.32+0.11 ms clock, 0.19+0/0.30/0.28+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1105752 HeapInuse: 1466368 HeapObjects: 15535 HeapIdle 368640 HeapReleased 0 HeapSys 1835008
gc 343 @0.733s 5%: 0.028+0.21+0.14 ms clock, 0.45+0/0.36/0.20+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
   10000	      2454 ns/op	      88 B/op	       2 allocs/op
gc 344 @0.734s 5%: 0.007+0.19+0.099 ms clock, 0.11+0/0.36/0.24+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069320 HeapInuse: 2433024 HeapObjects: 30539 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 345 @0.774s 5%: 0.016+0.26+0.11 ms clock, 0.25+0/0.31/0.36+1.8 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetErase1E4-16                           	gc 346 @0.775s 5%: 0.011+0.18+0.10 ms clock, 0.19+0/0.33/0.16+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069352 HeapInuse: 2433024 HeapObjects: 30540 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 347 @0.815s 5%: 0.015+0.34+0.17 ms clock, 0.25+0/0.36/0.24+2.8 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 348 @0.816s 5%: 0.027+0.20+0.10 ms clock, 0.44+0/0.30/0.26+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069432 HeapInuse: 2433024 HeapObjects: 30540 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 349 @0.856s 5%: 0.010+0.21+0.10 ms clock, 0.17+0/0.35/0.23+1.7 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 350 @0.856s 5%: 0.007+0.17+0.072 ms clock, 0.12+0/0.28/0.20+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069528 HeapInuse: 2433024 HeapObjects: 30541 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 351 @0.896s 4%: 0.011+0.20+0.099 ms clock, 0.18+0/0.31/0.16+1.5 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 352 @0.896s 4%: 0.021+0.20+0.097 ms clock, 0.34+0/0.26/0.25+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069624 HeapInuse: 2433024 HeapObjects: 30542 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 353 @0.938s 4%: 0.048+0.21+0.14 ms clock, 0.78+0/0.42/0.18+2.2 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 354 @0.939s 4%: 0.008+0.21+0.14 ms clock, 0.14+0/0.35/0.16+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069720 HeapInuse: 2433024 HeapObjects: 30543 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 355 @0.978s 4%: 0.009+0.23+0.19 ms clock, 0.15+0/0.33/0.23+3.1 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 356 @0.979s 4%: 0.022+0.23+0.091 ms clock, 0.35+0/0.32/0.37+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069816 HeapInuse: 2433024 HeapObjects: 30544 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 357 @1.018s 4%: 0.010+0.20+0.18 ms clock, 0.16+0/0.32/0.15+3.0 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 358 @1.019s 4%: 0.013+0.32+0.14 ms clock, 0.20+0/0.41/0.18+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2069912 HeapInuse: 2433024 HeapObjects: 30545 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 359 @1.061s 4%: 0.011+0.20+0.14 ms clock, 0.18+0/0.37/0.24+2.3 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 360 @1.062s 4%: 0.020+0.20+0.097 ms clock, 0.33+0/0.28/0.24+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070008 HeapInuse: 2433024 HeapObjects: 30546 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 361 @1.102s 4%: 0.023+0.25+0.10 ms clock, 0.37+0/0.30/0.38+1.6 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 362 @1.103s 4%: 0.021+0.18+0.092 ms clock, 0.34+0/0.35/0.20+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070104 HeapInuse: 2433024 HeapObjects: 30547 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 363 @1.143s 3%: 0.009+0.22+0.14 ms clock, 0.14+0/0.27/0.18+2.3 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 364 @1.144s 4%: 0.042+0.21+0.14 ms clock, 0.68+0/0.26/0.17+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070104 HeapInuse: 2433024 HeapObjects: 30547 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 365 @1.183s 3%: 0.023+0.19+0.16 ms clock, 0.38+0/0.36/0.14+2.6 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 366 @1.184s 3%: 0.029+0.18+0.12 ms clock, 0.47+0/0.22/0.16+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070104 HeapInuse: 2433024 HeapObjects: 30547 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 367 @1.223s 3%: 0.049+0.21+0.14 ms clock, 0.78+0/0.30/0.13+2.2 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 368 @1.223s 3%: 0.021+0.18+0.11 ms clock, 0.33+0/0.34/0.13+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070200 HeapInuse: 2433024 HeapObjects: 30548 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 369 @1.263s 3%: 0.049+0.20+0.15 ms clock, 0.79+0/0.30/0.28+2.4 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 370 @1.264s 3%: 0.026+0.17+0.12 ms clock, 0.42+0/0.31/0.24+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070200 HeapInuse: 2433024 HeapObjects: 30548 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 371 @1.304s 3%: 0.031+0.23+0.094 ms clock, 0.50+0/0.36/0.21+1.5 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 372 @1.305s 3%: 0.009+0.18+0.12 ms clock, 0.15+0/0.19/0.34+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070296 HeapInuse: 2433024 HeapObjects: 30549 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 373 @1.344s 3%: 0.010+0.24+0.18 ms clock, 0.16+0/0.30/0.13+2.9 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 374 @1.344s 3%: 0.009+0.22+0.14 ms clock, 0.14+0/0.34/0.18+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 2070392 HeapInuse: 2433024 HeapObjects: 30550 HeapIdle 417792 HeapReleased 0 HeapSys 2850816
gc 375 @1.384s 3%: 0.009+0.26+0.16 ms clock, 0.14+0/0.30/0.27+2.6 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
   10000	      1645 ns/op	      96 B/op	       2 allocs/op
gc 376 @1.385s 3%: 0.050+0.25+0.17 ms clock, 0.80+0/0.29/0.10+2.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190376 HeapInuse: 1548288 HeapObjects: 15550 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 377 @1.402s 3%: 0.049+0.21+0.11 ms clock, 0.79+0/0.25/0.11+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E4-16                  	gc 378 @1.402s 3%: 0.056+0.22+0.11 ms clock, 0.90+0/0.28/0.20+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190504 HeapInuse: 1548288 HeapObjects: 15552 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 379 @1.419s 3%: 0.021+0.23+0.13 ms clock, 0.33+0/0.45/0.14+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 380 @1.420s 3%: 0.022+0.19+0.12 ms clock, 0.36+0/0.29/0.17+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190584 HeapInuse: 1548288 HeapObjects: 15552 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 381 @1.436s 3%: 0.060+0.19+0.17 ms clock, 0.97+0/0.31/0.17+2.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 382 @1.437s 3%: 0.025+0.16+0.16 ms clock, 0.41+0/0.32/0.17+2.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190584 HeapInuse: 1548288 HeapObjects: 15552 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 383 @1.453s 3%: 0.047+0.20+0.13 ms clock, 0.75+0/0.30/0.16+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 384 @1.454s 3%: 0.012+0.19+0.10 ms clock, 0.19+0/0.27/0.18+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190584 HeapInuse: 1548288 HeapObjects: 15552 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 385 @1.470s 3%: 0.020+0.21+0.12 ms clock, 0.33+0/0.30/0.28+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 386 @1.470s 3%: 0.008+0.18+0.13 ms clock, 0.14+0/0.31/0.14+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190680 HeapInuse: 1548288 HeapObjects: 15553 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 387 @1.488s 3%: 0.011+0.24+0.085 ms clock, 0.18+0/0.36/0.28+1.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 388 @1.489s 3%: 0.012+0.17+0.078 ms clock, 0.19+0/0.30/0.30+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190776 HeapInuse: 1548288 HeapObjects: 15554 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 389 @1.505s 3%: 0.008+0.18+0.11 ms clock, 0.13+0/0.33/0.14+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 390 @1.506s 3%: 0.013+0.18+0.097 ms clock, 0.20+0/0.31/0.24+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190872 HeapInuse: 1548288 HeapObjects: 15555 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 391 @1.522s 3%: 0.011+0.22+0.11 ms clock, 0.19+0/0.29/0.30+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 392 @1.523s 3%: 0.013+0.20+0.10 ms clock, 0.22+0/0.34/0.22+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1190968 HeapInuse: 1548288 HeapObjects: 15556 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 393 @1.539s 3%: 0.054+0.19+0.17 ms clock, 0.87+0/0.32/0.22+2.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 394 @1.540s 3%: 0.025+0.21+0.15 ms clock, 0.41+0/0.29/0.23+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191064 HeapInuse: 1548288 HeapObjects: 15557 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 395 @1.557s 3%: 0.011+0.19+0.12 ms clock, 0.18+0/0.25/0.16+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 396 @1.557s 3%: 0.010+0.29+0.10 ms clock, 0.17+0/0.27/0.32+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191160 HeapInuse: 1548288 HeapObjects: 15558 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 397 @1.574s 3%: 0.013+0.25+0.12 ms clock, 0.21+0/0.25/0.24+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 398 @1.575s 3%: 0.007+0.27+0.11 ms clock, 0.11+0/0.29/0.13+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191256 HeapInuse: 1548288 HeapObjects: 15559 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 399 @1.591s 3%: 0.051+0.22+0.12 ms clock, 0.82+0/0.28/0.15+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 400 @1.591s 3%: 0.020+0.20+0.12 ms clock, 0.32+0/0.34/0.23+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191256 HeapInuse: 1548288 HeapObjects: 15559 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 401 @1.608s 3%: 0.008+0.28+0.15 ms clock, 0.13+0/0.38/0.17+2.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 402 @1.609s 3%: 0.023+0.19+0.17 ms clock, 0.38+0/0.28/0.19+2.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191352 HeapInuse: 1548288 HeapObjects: 15560 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 403 @1.625s 3%: 0.009+0.20+0.15 ms clock, 0.15+0/0.33/0.24+2.4 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 404 @1.626s 3%: 0.11+0.23+0.11 ms clock, 1.9+0/0.35/0.19+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191352 HeapInuse: 1548288 HeapObjects: 15560 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 405 @1.642s 3%: 0.013+0.26+0.12 ms clock, 0.21+0/0.35/0.18+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 406 @1.643s 3%: 0.038+0.23+0.12 ms clock, 0.61+0/0.43/0.097+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1191352 HeapInuse: 1548288 HeapObjects: 15560 HeapIdle 1302528 HeapReleased 0 HeapSys 2850816
gc 407 @1.660s 3%: 0.015+0.19+0.15 ms clock, 0.24+0/0.36/0.071+2.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
   10000	      1644 ns/op	      96 B/op	       2 allocs/op
gc 408 @1.660s 3%: 0.009+0.32+0.16 ms clock, 0.15+0/0.29/0.29+2.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793536 HeapInuse: 1155072 HeapObjects: 10566 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 409 @1.681s 3%: 0.009+0.20+0.11 ms clock, 0.15+0/0.21/0.20+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E4-16          	gc 410 @1.681s 3%: 0.015+0.16+0.10 ms clock, 0.25+0/0.23/0.17+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793584 HeapInuse: 1155072 HeapObjects: 10567 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 411 @1.702s 3%: 0.010+0.18+0.14 ms clock, 0.17+0/0.38/0.21+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 412 @1.702s 3%: 0.014+0.20+0.12 ms clock, 0.22+0/0.35/0.16+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793664 HeapInuse: 1155072 HeapObjects: 10567 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 413 @1.723s 3%: 0.043+0.23+0.19 ms clock, 0.70+0/0.29/0.22+3.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 414 @1.723s 3%: 0.013+0.23+0.14 ms clock, 0.21+0/0.33/0.20+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793760 HeapInuse: 1155072 HeapObjects: 10568 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 415 @1.744s 3%: 0.053+0.25+0.12 ms clock, 0.84+0/0.50/0.24+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 416 @1.744s 3%: 0.009+0.22+0.13 ms clock, 0.15+0/0.28/0.13+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793760 HeapInuse: 1155072 HeapObjects: 10568 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 417 @1.765s 3%: 0.010+0.17+0.11 ms clock, 0.16+0/0.42/0.11+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 418 @1.766s 3%: 0.015+0.21+0.091 ms clock, 0.24+0/0.21/0.21+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793760 HeapInuse: 1155072 HeapObjects: 10568 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 419 @1.786s 3%: 0.008+0.16+0.13 ms clock, 0.13+0/0.29/0.12+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 420 @1.787s 3%: 0.042+0.22+0.16 ms clock, 0.68+0/0.35/0.11+2.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793760 HeapInuse: 1155072 HeapObjects: 10568 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 421 @1.807s 3%: 0.059+0.25+0.17 ms clock, 0.94+0/0.32/0.20+2.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 422 @1.808s 3%: 0.021+0.19+0.14 ms clock, 0.35+0/0.27/0.32+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793856 HeapInuse: 1155072 HeapObjects: 10569 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 423 @1.828s 3%: 0.041+0.19+0.16 ms clock, 0.67+0/0.33/0.15+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 424 @1.829s 3%: 0.018+0.20+0.12 ms clock, 0.29+0/0.23/0.23+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793952 HeapInuse: 1155072 HeapObjects: 10570 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 425 @1.849s 3%: 0.052+0.20+0.14 ms clock, 0.84+0/0.29/0.14+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 426 @1.850s 3%: 0.038+0.20+0.10 ms clock, 0.61+0/0.28/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793952 HeapInuse: 1155072 HeapObjects: 10570 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 427 @1.870s 3%: 0.008+0.21+0.12 ms clock, 0.13+0/0.21/0.23+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 428 @1.870s 3%: 0.011+0.19+0.097 ms clock, 0.18+0/0.30/0.23+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793952 HeapInuse: 1155072 HeapObjects: 10570 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 429 @1.890s 3%: 0.028+0.19+0.15 ms clock, 0.45+0/0.30/0.18+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 430 @1.891s 3%: 0.011+0.20+0.11 ms clock, 0.17+0/0.26/0.26+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793952 HeapInuse: 1155072 HeapObjects: 10570 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 431 @1.912s 3%: 0.008+0.19+0.13 ms clock, 0.13+0/0.26/0.23+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 432 @1.912s 3%: 0.076+0.20+0.14 ms clock, 1.2+0/0.24/0.33+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 793952 HeapInuse: 1155072 HeapObjects: 10570 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 433 @1.934s 3%: 0.010+0.21+0.11 ms clock, 0.16+0/0.34/0.17+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 434 @1.934s 3%: 0.040+0.20+0.095 ms clock, 0.64+0/0.32/0.15+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 794048 HeapInuse: 1155072 HeapObjects: 10571 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 435 @1.955s 2%: 0.048+0.21+0.15 ms clock, 0.78+0/0.30/0.19+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 436 @1.955s 3%: 0.021+0.20+0.11 ms clock, 0.34+0/0.25/0.15+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 794144 HeapInuse: 1155072 HeapObjects: 10572 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 437 @1.976s 2%: 0.056+0.26+0.19 ms clock, 0.89+0/0.36/0.19+3.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 438 @1.976s 3%: 0.095+0.29+0.22 ms clock, 1.5+0/0.37/0.14+3.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 794144 HeapInuse: 1155072 HeapObjects: 10572 HeapIdle 1695744 HeapReleased 0 HeapSys 2850816
gc 439 @1.997s 2%: 0.010+0.20+0.15 ms clock, 0.16+0/0.34/0.28+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
   10000	      2021 ns/op	      56 B/op	       1 allocs/op
gc 440 @1.997s 2%: 0.016+0.19+0.10 ms clock, 0.26+0/0.25/0.23+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 657016 HeapInuse: 1081344 HeapObjects: 886 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 441 @2.001s 2%: 0.007+0.20+0.11 ms clock, 0.12+0/0.30/0.21+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E4-16                   	gc 442 @2.002s 2%: 0.021+0.24+0.19 ms clock, 0.33+0/0.33/0.26+3.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 657984 HeapInuse: 1089536 HeapObjects: 896 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 443 @2.005s 3%: 0.022+0.17+0.12 ms clock, 0.35+0/0.30/0.14+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 444 @2.006s 3%: 0.015+0.17+0.11 ms clock, 0.25+0/0.32/0.17+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 660248 HeapInuse: 1089536 HeapObjects: 898 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 445 @2.009s 3%: 0.026+0.18+0.10 ms clock, 0.42+0/0.20/0.21+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 446 @2.010s 3%: 0.006+0.17+0.12 ms clock, 0.10+0/0.29/0.18+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 658952 HeapInuse: 1081344 HeapObjects: 883 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 447 @2.013s 3%: 0.009+0.17+0.093 ms clock, 0.14+0/0.36/0.18+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 448 @2.014s 3%: 0.007+0.18+0.13 ms clock, 0.12+0/0.29/0.11+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 657048 HeapInuse: 1081344 HeapObjects: 865 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 449 @2.017s 3%: 0.009+0.24+0.096 ms clock, 0.14+0/0.39/0.27+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 450 @2.018s 3%: 0.045+0.19+0.11 ms clock, 0.73+0/0.15/0.22+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 659224 HeapInuse: 1089536 HeapObjects: 906 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 451 @2.022s 3%: 0.022+0.29+0.11 ms clock, 0.36+0/0.27/0.30+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 452 @2.022s 3%: 0.016+0.23+0.11 ms clock, 0.26+0/0.30/0.24+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 658464 HeapInuse: 1081344 HeapObjects: 877 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 453 @2.026s 3%: 0.026+0.24+0.11 ms clock, 0.41+0/0.27/0.18+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 454 @2.026s 3%: 0.006+0.17+0.12 ms clock, 0.098+0/0.26/0.23+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 659216 HeapInuse: 1089536 HeapObjects: 904 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 455 @2.030s 3%: 0.009+0.21+0.18 ms clock, 0.15+0/0.34/0.25+2.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 456 @2.030s 3%: 0.019+0.20+0.11 ms clock, 0.31+0/0.30/0.24+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 658384 HeapInuse: 1089536 HeapObjects: 898 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 457 @2.034s 3%: 0.032+0.17+0.10 ms clock, 0.51+0/0.31/0.12+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 458 @2.034s 3%: 0.009+0.18+0.10 ms clock, 0.15+0/0.31/0.14+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 656224 HeapInuse: 1081344 HeapObjects: 878 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 459 @2.038s 3%: 0.027+0.18+0.12 ms clock, 0.44+0/0.19/0.23+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 460 @2.038s 3%: 0.011+0.26+0.10 ms clock, 0.17+0/0.24/0.29+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 658104 HeapInuse: 1081344 HeapObjects: 896 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 461 @2.042s 3%: 0.008+0.27+0.084 ms clock, 0.12+0/0.31/0.12+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 462 @2.043s 3%: 0.019+0.20+0.094 ms clock, 0.31+0/0.36/0.10+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 661400 HeapInuse: 1089536 HeapObjects: 910 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 463 @2.046s 3%: 0.009+0.18+0.10 ms clock, 0.14+0/0.24/0.16+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 464 @2.047s 3%: 0.040+0.22+0.10 ms clock, 0.64+0/0.30/0.27+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 662024 HeapInuse: 1089536 HeapObjects: 915 HeapIdle 1761280 HeapReleased 0 HeapSys 2850816
gc 465 @2.050s 3%: 0.042+0.21+0.10 ms clock, 0.67+0/0.32/0.28+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 466 @2.051s 3%: 0.010+0.26+0.13 ms clock, 0.16+0/0.32/0.23+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 656720 HeapInuse: 1081344 HeapObjects: 880 HeapIdle 1769472 HeapReleased 0 HeapSys 2850816
gc 467 @2.054s 3%: 0.008+0.16+0.10 ms clock, 0.13+0/0.30/0.17+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
   10000	       330 ns/op	      42 B/op	       0 allocs/op
gc 468 @2.055s 3%: 0.010+0.20+0.095 ms clock, 0.16+0/0.28/0.17+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 741368 HeapInuse: 1171456 HeapObjects: 911 HeapIdle 1679360 HeapReleased 0 HeapSys 2850816
gc 469 @2.060s 3%: 0.008+0.20+0.097 ms clock, 0.13+0/0.33/0.26+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E4-16                    	gc 470 @2.060s 3%: 0.018+0.32+0.20 ms clock, 0.29+0/0.29/0.20+3.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 740424 HeapInuse: 1163264 HeapObjects: 902 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 471 @2.065s 3%: 0.008+0.20+0.12 ms clock, 0.13+0/0.35/0.16+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 472 @2.066s 3%: 0.019+0.18+0.11 ms clock, 0.30+0/0.35/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 737400 HeapInuse: 1163264 HeapObjects: 870 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 473 @2.070s 3%: 0.026+0.19+0.10 ms clock, 0.43+0/0.36/0.16+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 474 @2.071s 3%: 0.015+0.17+0.16 ms clock, 0.24+0/0.28/0.21+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 739064 HeapInuse: 1163264 HeapObjects: 883 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 475 @2.076s 3%: 0.025+0.22+0.17 ms clock, 0.40+0/0.37/0.19+2.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 476 @2.076s 3%: 0.018+0.21+0.13 ms clock, 0.30+0/0.37/0.23+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 740632 HeapInuse: 1163264 HeapObjects: 903 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 477 @2.081s 3%: 0.030+0.20+0.11 ms clock, 0.49+0/0.30/0.17+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 478 @2.081s 3%: 0.014+0.20+0.098 ms clock, 0.23+0/0.29/0.27+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 743544 HeapInuse: 1171456 HeapObjects: 913 HeapIdle 1679360 HeapReleased 0 HeapSys 2850816
gc 479 @2.086s 3%: 0.010+0.18+0.12 ms clock, 0.17+0/0.32/0.25+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 480 @2.086s 3%: 0.057+0.17+0.10 ms clock, 0.92+0/0.43/0.082+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 741480 HeapInuse: 1163264 HeapObjects: 891 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 481 @2.091s 3%: 0.008+0.24+0.099 ms clock, 0.14+0/0.27/0.35+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 482 @2.091s 3%: 0.010+0.24+0.11 ms clock, 0.16+0/0.37/0.27+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 739640 HeapInuse: 1163264 HeapObjects: 893 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 483 @2.096s 3%: 0.009+0.19+0.11 ms clock, 0.15+0/0.29/0.24+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 484 @2.096s 3%: 0.011+0.20+0.10 ms clock, 0.18+0/0.26/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 740008 HeapInuse: 1163264 HeapObjects: 896 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 485 @2.101s 3%: 0.009+0.18+0.11 ms clock, 0.14+0/0.32/0.22+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 486 @2.101s 3%: 0.006+0.23+0.10 ms clock, 0.096+0/0.32/0.27+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 741208 HeapInuse: 1163264 HeapObjects: 908 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 487 @2.106s 3%: 0.010+0.19+0.11 ms clock, 0.16+0/0.31/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 488 @2.106s 3%: 0.010+0.20+0.11 ms clock, 0.16+0/0.37/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 740112 HeapInuse: 1163264 HeapObjects: 897 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 489 @2.111s 3%: 0.034+0.19+0.12 ms clock, 0.55+0/0.25/0.24+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 490 @2.111s 3%: 0.011+0.19+0.13 ms clock, 0.18+0/0.35/0.12+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 738824 HeapInuse: 1163264 HeapObjects: 883 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 491 @2.116s 3%: 0.023+0.16+0.11 ms clock, 0.37+0/0.30/0.26+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 492 @2.117s 3%: 0.016+0.18+0.12 ms clock, 0.25+0/0.25/0.22+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 742832 HeapInuse: 1163264 HeapObjects: 904 HeapIdle 1687552 HeapReleased 0 HeapSys 2850816
gc 493 @2.121s 3%: 0.007+0.17+0.11 ms clock, 0.12+0/0.28/0.27+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
   10000	       137 ns/op	       8 B/op	       0 allocs/op
gc 494 @2.121s 3%: 0.055+0.17+0.13 ms clock, 0.89+0/0.31/0.35+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294368 HeapInuse: 712704 HeapObjects: 666 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 495 @2.123s 3%: 0.009+0.20+0.10 ms clock, 0.15+0/0.33/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E4-16           	gc 496 @2.124s 3%: 0.041+0.21+0.084 ms clock, 0.66+0/0.30/0.31+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294928 HeapInuse: 712704 HeapObjects: 673 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 497 @2.126s 3%: 0.021+0.24+0.092 ms clock, 0.35+0/0.25/0.25+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 498 @2.126s 3%: 0.009+0.19+0.11 ms clock, 0.14+0/0.32/0.26+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294024 HeapInuse: 712704 HeapObjects: 660 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 499 @2.128s 3%: 0.008+0.26+0.11 ms clock, 0.13+0/0.32/0.22+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 500 @2.129s 3%: 0.020+0.18+0.092 ms clock, 0.32+0/0.38/0.074+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 295664 HeapInuse: 712704 HeapObjects: 680 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 501 @2.131s 3%: 0.008+0.19+0.082 ms clock, 0.13+0/0.30/0.26+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 502 @2.131s 3%: 0.015+0.19+0.13 ms clock, 0.24+0/0.33/0.19+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 293656 HeapInuse: 712704 HeapObjects: 660 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 503 @2.133s 3%: 0.008+0.21+0.11 ms clock, 0.13+0/0.31/0.18+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 504 @2.133s 3%: 0.017+0.18+0.085 ms clock, 0.28+0/0.32/0.33+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294792 HeapInuse: 712704 HeapObjects: 668 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 505 @2.135s 3%: 0.009+0.16+0.090 ms clock, 0.15+0/0.24/0.30+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 506 @2.136s 3%: 0.006+0.17+0.093 ms clock, 0.11+0/0.31/0.27+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294936 HeapInuse: 712704 HeapObjects: 672 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 507 @2.138s 3%: 0.008+0.16+0.13 ms clock, 0.12+0/0.30/0.26+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 508 @2.138s 3%: 0.016+0.22+0.098 ms clock, 0.26+0/0.36/0.25+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 295288 HeapInuse: 712704 HeapObjects: 674 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 509 @2.140s 3%: 0.009+0.16+0.16 ms clock, 0.14+0/0.27/0.13+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 510 @2.140s 3%: 0.007+0.26+0.088 ms clock, 0.12+0/0.42/0.15+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 295192 HeapInuse: 712704 HeapObjects: 674 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 511 @2.142s 3%: 0.009+0.17+0.075 ms clock, 0.15+0/0.25/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 512 @2.143s 3%: 0.014+0.24+0.073 ms clock, 0.22+0/0.28/0.15+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294848 HeapInuse: 712704 HeapObjects: 671 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 513 @2.145s 3%: 0.007+0.18+0.13 ms clock, 0.12+0/0.29/0.34+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 514 @2.145s 3%: 0.057+0.17+0.080 ms clock, 0.92+0/0.33/0.25+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 295248 HeapInuse: 712704 HeapObjects: 676 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 515 @2.148s 3%: 0.010+0.21+0.073 ms clock, 0.16+0/0.36/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 516 @2.148s 3%: 0.006+0.18+0.070 ms clock, 0.10+0/0.24/0.25+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 294992 HeapInuse: 712704 HeapObjects: 672 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 517 @2.150s 3%: 0.008+0.20+0.074 ms clock, 0.13+0/0.29/0.24+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 518 @2.150s 3%: 0.008+0.25+0.077 ms clock, 0.12+0/0.25/0.15+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 295288 HeapInuse: 712704 HeapObjects: 674 HeapIdle 2138112 HeapReleased 0 HeapSys 2850816
gc 519 @2.152s 3%: 0.008+0.15+0.075 ms clock, 0.12+0/0.34/0.28+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
   10000	       167 ns/op	       6 B/op	       0 allocs/op
gc 520 @2.153s 3%: 0.007+0.18+0.082 ms clock, 0.12+0/0.30/0.28+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 438200 HeapInuse: 794624 HeapObjects: 602 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 521 @2.154s 3%: 0.010+0.18+0.075 ms clock, 0.16+0/0.32/0.27+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E4-16    	gc 522 @2.155s 3%: 0.011+0.19+0.072 ms clock, 0.18+0/0.32/0.22+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 438440 HeapInuse: 794624 HeapObjects: 605 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 523 @2.156s 3%: 0.008+0.16+0.075 ms clock, 0.14+0/0.28/0.40+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 524 @2.156s 3%: 0.009+0.16+0.083 ms clock, 0.15+0/0.33/0.23+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 438520 HeapInuse: 794624 HeapObjects: 605 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 525 @2.158s 3%: 0.007+0.21+0.089 ms clock, 0.12+0/0.34/0.30+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 526 @2.158s 3%: 0.011+0.26+0.11 ms clock, 0.18+0/0.41/0.20+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 438712 HeapInuse: 794624 HeapObjects: 607 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 527 @2.160s 3%: 0.008+0.17+0.091 ms clock, 0.13+0/0.30/0.22+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 528 @2.160s 3%: 0.011+0.20+0.069 ms clock, 0.18+0/0.24/0.27+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 438904 HeapInuse: 794624 HeapObjects: 609 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 529 @2.162s 3%: 0.008+0.16+0.091 ms clock, 0.13+0/0.30/0.23+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 530 @2.162s 3%: 0.007+0.19+0.091 ms clock, 0.12+0/0.31/0.26+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439096 HeapInuse: 794624 HeapObjects: 611 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 531 @2.164s 3%: 0.007+0.20+0.081 ms clock, 0.12+0/0.31/0.26+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 532 @2.164s 3%: 0.007+0.17+0.098 ms clock, 0.12+0/0.31/0.22+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439096 HeapInuse: 794624 HeapObjects: 611 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 533 @2.166s 3%: 0.008+0.16+0.076 ms clock, 0.13+0/0.24/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 534 @2.166s 3%: 0.007+0.27+0.097 ms clock, 0.12+0/0.29/0.17+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439192 HeapInuse: 794624 HeapObjects: 612 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 535 @2.168s 3%: 0.008+0.17+0.073 ms clock, 0.13+0/0.32/0.27+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 536 @2.168s 3%: 0.020+0.18+0.10 ms clock, 0.33+0/0.31/0.25+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439384 HeapInuse: 794624 HeapObjects: 614 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 537 @2.170s 3%: 0.007+0.17+0.079 ms clock, 0.12+0/0.29/0.21+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 538 @2.170s 3%: 0.008+0.17+0.077 ms clock, 0.13+0/0.28/0.27+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439576 HeapInuse: 794624 HeapObjects: 616 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 539 @2.171s 3%: 0.007+0.19+0.075 ms clock, 0.12+0/0.28/0.24+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 540 @2.172s 3%: 0.007+0.21+0.098 ms clock, 0.12+0/0.45/0.26+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439768 HeapInuse: 794624 HeapObjects: 618 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 541 @2.173s 3%: 0.008+0.20+0.11 ms clock, 0.12+0/0.39/0.26+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 542 @2.174s 3%: 0.010+0.19+0.085 ms clock, 0.16+0/0.26/0.30+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 439960 HeapInuse: 794624 HeapObjects: 620 HeapIdle 2056192 HeapReleased 0 HeapSys 2850816
gc 543 @2.175s 3%: 0.008+0.16+0.10 ms clock, 0.12+0/0.30/0.23+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 544 @2.176s 3%: 0.006+0.19+0.074 ms clock, 0.096+0/0.30/0.14+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 440152 HeapInuse: 802816 HeapObjects: 622 HeapIdle 2048000 HeapReleased 0 HeapSys 2850816
gc 545 @2.177s 3%: 0.009+0.28+0.076 ms clock, 0.14+0/0.071/0.34+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
   10000	       145 ns/op	      20 B/op	       0 allocs/op
gc 546 @2.178s 3%: 0.006+0.18+0.089 ms clock, 0.11+0/0.33/0.28+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038136 HeapInuse: 1400832 HeapObjects: 621 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 547 @2.211s 3%: 0.020+0.21+0.19 ms clock, 0.32+0/0.28/0.25+3.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSort1E5-16                               	gc 548 @2.211s 3%: 0.038+0.20+0.15 ms clock, 0.62+0/0.29/0.16+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038264 HeapInuse: 1400832 HeapObjects: 623 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 549 @2.244s 3%: 0.010+0.19+0.16 ms clock, 0.17+0/0.32/0.17+2.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 550 @2.245s 3%: 0.010+0.24+0.12 ms clock, 0.17+0/0.39/0.15+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038344 HeapInuse: 1400832 HeapObjects: 623 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 551 @2.278s 3%: 0.071+0.22+0.21 ms clock, 1.1+0/0.23/0.19+3.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 552 @2.279s 3%: 0.073+0.21+0.14 ms clock, 1.1+0/0.35/0.11+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038344 HeapInuse: 1400832 HeapObjects: 623 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 553 @2.312s 3%: 0.043+0.20+0.17 ms clock, 0.70+0/0.24/0.22+2.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 554 @2.313s 3%: 0.016+0.25+0.13 ms clock, 0.25+0/0.23/0.28+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038440 HeapInuse: 1400832 HeapObjects: 624 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 555 @2.345s 3%: 0.034+0.22+0.13 ms clock, 0.54+0/0.36/0.22+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 556 @2.346s 3%: 0.099+0.23+0.15 ms clock, 1.5+0/0.48/0.061+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038440 HeapInuse: 1400832 HeapObjects: 624 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 557 @2.379s 3%: 0.009+0.27+0.12 ms clock, 0.15+0/0.42/0.21+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 558 @2.379s 3%: 0.075+0.18+0.096 ms clock, 1.2+0/0.27/0.26+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038632 HeapInuse: 1400832 HeapObjects: 626 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 559 @2.413s 3%: 0.038+0.24+0.12 ms clock, 0.62+0/0.21/0.33+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 560 @2.414s 3%: 0.010+0.20+0.14 ms clock, 0.16+0/0.24/0.20+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038632 HeapInuse: 1400832 HeapObjects: 626 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 561 @2.447s 3%: 0.014+0.20+0.13 ms clock, 0.23+0/0.25/0.25+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 562 @2.447s 3%: 0.040+0.34+0.14 ms clock, 0.65+0/0.38/0.19+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038632 HeapInuse: 1400832 HeapObjects: 626 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 563 @2.480s 3%: 0.010+0.24+0.14 ms clock, 0.16+0/0.33/0.17+2.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 564 @2.481s 3%: 0.014+0.26+0.092 ms clock, 0.22+0/0.35/0.16+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038728 HeapInuse: 1400832 HeapObjects: 627 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 565 @2.515s 3%: 0.012+0.23+0.095 ms clock, 0.19+0/0.33/0.21+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 566 @2.515s 3%: 0.008+0.24+0.11 ms clock, 0.14+0/0.42/0.17+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038824 HeapInuse: 1400832 HeapObjects: 628 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 567 @2.549s 3%: 0.010+0.27+0.12 ms clock, 0.16+0/0.45/0.11+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 568 @2.550s 3%: 0.048+0.21+0.10 ms clock, 0.78+0/0.42/0.10+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038824 HeapInuse: 1400832 HeapObjects: 628 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 569 @2.583s 3%: 0.010+0.21+0.11 ms clock, 0.17+0/0.32/0.18+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 570 @2.583s 3%: 0.019+0.20+0.13 ms clock, 0.31+0/0.35/0.24+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038824 HeapInuse: 1400832 HeapObjects: 628 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 571 @2.616s 3%: 0.050+0.23+0.11 ms clock, 0.81+0/0.27/0.20+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 572 @2.616s 3%: 0.059+0.22+0.12 ms clock, 0.95+0/0.25/0.28+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1038920 HeapInuse: 1400832 HeapObjects: 629 HeapIdle 1449984 HeapReleased 0 HeapSys 2850816
gc 573 @2.649s 3%: 0.012+0.19+0.11 ms clock, 0.19+0/0.26/0.15+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
  100000	       324 ns/op	       8 B/op	       0 allocs/op
gc 574 @2.650s 3%: 0.011+0.21+0.12 ms clock, 0.18+0/0.25/0.26+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 575 @2.771s 2%: 0.016+4.5+0.17 ms clock, 0.26+0.18/12/10+2.8 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 576 @2.894s 2%: 0.015+6.0+0.11 ms clock, 0.24+0.16/21/45+1.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036392 HeapInuse: 9420800 HeapObjects: 150635 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 577 @2.916s 2%: 0.008+0.25+0.18 ms clock, 0.13+0/0.40/0.22+2.9 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
#BenchmarkSetInsert1E5-16                          	gc 578 @2.917s 2%: 0.057+0.20+0.10 ms clock, 0.92+0/0.33/0.13+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 579 @3.034s 2%: 0.017+3.0+0.19 ms clock, 0.27+0.23/10/20+3.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 580 @3.154s 2%: 0.017+5.6+0.16 ms clock, 0.28+0.18/20/42+2.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036488 HeapInuse: 9420800 HeapObjects: 150635 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 581 @3.175s 2%: 0.009+0.25+0.15 ms clock, 0.15+0/0.42/0.29+2.4 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 582 @3.177s 2%: 0.009+0.31+0.14 ms clock, 0.14+0/0.28/0.30+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 583 @3.294s 2%: 0.015+3.1+0.13 ms clock, 0.25+0.21/10/18+2.2 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 584 @3.415s 2%: 0.015+7.6+0.18 ms clock, 0.25+0.25/23/47+2.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036472 HeapInuse: 9420800 HeapObjects: 150634 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 585 @3.436s 2%: 0.038+0.27+0.19 ms clock, 0.61+0/0.38/0.26+3.1 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 586 @3.438s 2%: 0.026+0.23+0.11 ms clock, 0.41+0/0.43/0.16+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 587 @3.556s 2%: 0.021+3.7+0.12 ms clock, 0.33+0.26/12/20+2.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 588 @3.678s 2%: 0.018+6.3+0.13 ms clock, 0.29+0.42/21/42+2.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036504 HeapInuse: 9420800 HeapObjects: 150636 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 589 @3.699s 2%: 0.038+0.23+0.16 ms clock, 0.62+0/0.41/0.22+2.5 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 590 @3.700s 2%: 0.021+0.23+0.20 ms clock, 0.34+0/0.33/0.29+3.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 591 @3.818s 2%: 0.016+3.6+0.12 ms clock, 0.26+0.25/11/16+2.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 592 @3.944s 2%: 0.018+6.5+0.13 ms clock, 0.30+0.24/22/41+2.1 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036504 HeapInuse: 9420800 HeapObjects: 150636 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 593 @3.963s 2%: 0.013+0.25+0.21 ms clock, 0.21+0/0.35/0.15+3.5 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 594 @3.965s 2%: 0.020+0.28+0.12 ms clock, 0.33+0/0.27/0.28+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 595 @4.082s 2%: 0.018+3.0+0.094 ms clock, 0.29+0.19/9.9/20+1.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 596 @4.204s 2%: 0.018+6.2+0.11 ms clock, 0.29+0.41/20/46+1.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036552 HeapInuse: 9420800 HeapObjects: 150639 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 597 @4.224s 2%: 0.011+0.21+0.10 ms clock, 0.18+0/0.35/0.24+1.6 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 598 @4.227s 2%: 0.053+0.22+0.16 ms clock, 0.85+0/0.29/0.27+2.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 599 @4.343s 2%: 0.018+2.8+0.12 ms clock, 0.30+0.22/9.7/20+1.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 600 @4.466s 2%: 0.016+5.0+0.16 ms clock, 0.26+0.21/17/43+2.5 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036488 HeapInuse: 9420800 HeapObjects: 150635 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 601 @4.486s 2%: 0.010+0.25+0.080 ms clock, 0.16+0/0.27/0.39+1.2 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 602 @4.488s 2%: 0.009+0.22+0.072 ms clock, 0.15+0/0.38/0.32+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 603 @4.606s 2%: 0.017+3.3+0.12 ms clock, 0.27+0.24/11/17+1.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 604 @4.729s 2%: 0.017+5.2+0.31 ms clock, 0.28+0.24/19/41+5.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036488 HeapInuse: 9420800 HeapObjects: 150635 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 605 @4.748s 2%: 0.010+0.29+0.13 ms clock, 0.17+0/0.36/0.25+2.1 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 606 @4.750s 2%: 0.008+0.20+0.12 ms clock, 0.14+0/0.35/0.19+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 607 @4.869s 2%: 0.018+2.9+0.11 ms clock, 0.30+0.24/9.8/20+1.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 608 @4.991s 2%: 0.017+5.4+0.10 ms clock, 0.27+0.25/19/42+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036504 HeapInuse: 9420800 HeapObjects: 150636 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 609 @5.011s 2%: 0.011+0.23+0.12 ms clock, 0.18+0/0.41/0.20+2.0 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 610 @5.012s 2%: 0.008+0.27+0.11 ms clock, 0.14+0/0.38/0.22+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 611 @5.129s 2%: 0.017+3.2+0.14 ms clock, 0.28+0.16/11/18+2.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 612 @5.250s 2%: 0.015+8.6+0.21 ms clock, 0.24+0.26/22/43+3.3 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036616 HeapInuse: 9420800 HeapObjects: 150638 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 613 @5.271s 2%: 0.024+0.24+0.15 ms clock, 0.38+0/0.43/0.25+2.4 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 614 @5.273s 2%: 0.010+0.22+0.17 ms clock, 0.16+0/0.34/0.25+2.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 615 @5.390s 1%: 0.016+3.0+0.16 ms clock, 0.26+0.23/9.4/18+2.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 616 @5.516s 1%: 0.017+5.4+0.11 ms clock, 0.27+0.14/19/45+1.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036568 HeapInuse: 9420800 HeapObjects: 150635 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 617 @5.535s 1%: 0.011+0.26+0.12 ms clock, 0.17+0/0.42/0.34+1.9 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 618 @5.537s 1%: 0.010+0.25+0.15 ms clock, 0.17+0/0.39/0.14+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 619 @5.656s 1%: 0.017+3.3+0.13 ms clock, 0.28+0.18/11/19+2.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 620 @5.779s 1%: 0.023+6.0+0.13 ms clock, 0.37+0.15/21/44+2.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036600 HeapInuse: 9420800 HeapObjects: 150637 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 621 @5.800s 1%: 0.010+0.27+0.13 ms clock, 0.16+0/0.39/0.41+2.1 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 622 @5.802s 1%: 0.008+0.24+0.13 ms clock, 0.14+0/0.43/0.24+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 623 @5.919s 1%: 0.018+3.7+0.13 ms clock, 0.29+0.31/12/17+2.1 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 624 @6.041s 1%: 0.018+5.2+0.11 ms clock, 0.29+0.20/18/40+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036584 HeapInuse: 9420800 HeapObjects: 150636 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 625 @6.060s 1%: 0.010+0.31+0.16 ms clock, 0.16+0/0.43/0.19+2.6 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 626 @6.063s 1%: 0.011+0.23+0.12 ms clock, 0.18+0/0.41/0.17+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 627 @6.181s 1%: 0.017+3.5+0.10 ms clock, 0.27+0.19/11/17+1.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 628 @6.303s 1%: 0.018+4.6+0.10 ms clock, 0.29+0.16/17/42+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036600 HeapInuse: 9420800 HeapObjects: 150637 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 629 @6.322s 1%: 0.009+0.23+0.20 ms clock, 0.14+0/0.37/0.12+3.3 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 630 @6.324s 1%: 0.023+0.22+0.16 ms clock, 0.37+0/0.42/0.21+2.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 631 @6.442s 1%: 0.018+2.9+0.18 ms clock, 0.29+0.21/10/22+2.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 632 @6.567s 1%: 0.019+10+0.14 ms clock, 0.30+0.45/26/26+2.2 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036584 HeapInuse: 9420800 HeapObjects: 150636 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 633 @6.589s 1%: 0.044+0.23+0.16 ms clock, 0.70+0/0.30/0.25+2.6 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 634 @6.591s 1%: 0.009+0.23+0.12 ms clock, 0.14+0/0.34/0.22+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 635 @6.714s 1%: 0.018+2.8+0.12 ms clock, 0.29+0.42/10/16+2.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 636 @6.837s 1%: 0.018+6.7+0.18 ms clock, 0.30+0.21/24/35+2.8 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036600 HeapInuse: 9420800 HeapObjects: 150637 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 637 @6.857s 1%: 0.009+0.22+0.079 ms clock, 0.14+0/0.42/0.14+1.2 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
gc 638 @6.858s 1%: 0.006+0.19+0.12 ms clock, 0.10+0/0.34/0.15+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 639 @6.975s 1%: 0.018+3.2+0.21 ms clock, 0.30+0.17/11/15+3.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
gc 640 @7.097s 1%: 0.018+6.1+0.11 ms clock, 0.28+0.23/22/36+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 9036664 HeapInuse: 9420800 HeapObjects: 150636 HeapIdle 770048 HeapReleased 0 HeapSys 10190848
gc 641 @7.117s 1%: 0.008+0.22+0.14 ms clock, 0.14+0/0.33/0.19+2.3 ms cpu, 8->8->0 MB, 15 MB goal, 16 P (forced)
  100000	      2591 ns/op	      88 B/op	       2 allocs/op
gc 642 @7.118s 1%: 0.079+0.21+0.20 ms clock, 1.2+0/0.28/0.21+3.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 643 @7.211s 1%: 0.017+3.8+0.15 ms clock, 0.27+0.44/10/8.8+2.5 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 644 @7.345s 1%: 0.015+6.1+0.17 ms clock, 0.24+0.24/20/27+2.7 ms cpu, 7->8->8 MB, 8 MB goal, 16 P
gc 645 @7.517s 1%: 0.020+2.3+0.14 ms clock, 0.32+0.24/7.8/12+2.3 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5685624 HeapInuse: 6053888 HeapObjects: 79834 HeapIdle 11476992 HeapReleased 0 HeapSys 17530880
gc 646 @7.565s 1%: 0.011+0.27+0.20 ms clock, 0.18+0/0.23/0.19+3.2 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
#BenchmarkSetErase1E5-16                           	gc 647 @7.566s 1%: 0.039+0.23+0.11 ms clock, 0.62+0/0.36/0.20+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 648 @7.659s 1%: 0.014+2.7+0.12 ms clock, 0.23+0.25/9.2/16+2.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 649 @7.791s 1%: 0.017+5.3+0.14 ms clock, 0.27+0.50/18/41+2.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 650 @7.950s 1%: 0.020+3.3+0.12 ms clock, 0.32+0.36/10/10+1.9 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6475688 HeapInuse: 6848512 HeapObjects: 93303 HeapIdle 10682368 HeapReleased 0 HeapSys 17530880
gc 651 @8.007s 1%: 0.034+0.27+0.19 ms clock, 0.54+0/0.42/0.38+3.1 ms cpu, 6->6->0 MB, 7 MB goal, 16 P (forced)
gc 652 @8.008s 1%: 0.010+0.23+0.17 ms clock, 0.16+0/0.40/0.15+2.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 653 @8.100s 1%: 0.017+2.9+0.11 ms clock, 0.28+0.44/9.8/14+1.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 654 @8.241s 1%: 0.019+5.4+0.10 ms clock, 0.31+0.15/19/35+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 655 @8.404s 1%: 0.023+2.6+0.23 ms clock, 0.37+0.19/8.7/9.6+3.7 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6154520 HeapInuse: 6520832 HeapObjects: 87826 HeapIdle 11010048 HeapReleased 0 HeapSys 17530880
gc 656 @8.457s 1%: 0.042+0.22+0.14 ms clock, 0.67+0/0.36/0.16+2.2 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 657 @8.458s 1%: 0.072+0.22+0.18 ms clock, 1.1+0/0.39/0.16+3.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 658 @8.552s 1%: 0.018+2.5+0.11 ms clock, 0.29+0.37/8.8/16+1.8 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 659 @8.687s 1%: 0.020+5.9+0.18 ms clock, 0.32+0.19/21/34+3.0 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 660 @8.853s 1%: 0.021+2.0+0.10 ms clock, 0.34+0.41/6.9/10+1.6 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5793704 HeapInuse: 6160384 HeapObjects: 81675 HeapIdle 11370496 HeapReleased 0 HeapSys 17530880
gc 661 @8.903s 1%: 0.011+0.27+0.11 ms clock, 0.17+0/0.43/0.28+1.7 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 662 @8.904s 1%: 0.017+0.27+0.094 ms clock, 0.28+0/0.36/0.15+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 663 @9.000s 1%: 0.019+4.1+0.12 ms clock, 0.31+0.12/11/12+2.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 664 @9.128s 1%: 0.017+5.4+0.12 ms clock, 0.28+0.30/20/27+2.0 ms cpu, 7->8->8 MB, 8 MB goal, 16 P
gc 665 @9.234s 1%: 0.016+1.7+0.12 ms clock, 0.26+0.14/6.1/10+1.9 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5703832 HeapInuse: 6070272 HeapObjects: 80145 HeapIdle 11460608 HeapReleased 0 HeapSys 17530880
gc 666 @9.265s 1%: 0.010+0.24+0.10 ms clock, 0.17+0/0.34/0.20+1.6 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 667 @9.266s 1%: 0.011+0.19+0.11 ms clock, 0.17+0/0.25/0.18+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 668 @9.363s 1%: 0.018+3.1+0.092 ms clock, 0.28+0.12/10/9.9+1.4 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 669 @9.447s 1%: 0.016+5.3+0.10 ms clock, 0.25+0.12/17/33+1.6 ms cpu, 7->8->8 MB, 8 MB goal, 16 P
gc 670 @9.555s 1%: 0.018+1.9+0.079 ms clock, 0.29+0.11/6.4/11+1.2 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5794616 HeapInuse: 6160384 HeapObjects: 81690 HeapIdle 11370496 HeapReleased 0 HeapSys 17530880
gc 671 @9.604s 1%: 0.009+0.26+0.19 ms clock, 0.15+0/0.39/0.19+3.1 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 672 @9.605s 1%: 0.009+0.21+0.13 ms clock, 0.15+0/0.39/0.11+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 673 @9.697s 1%: 0.015+2.9+0.29 ms clock, 0.24+0.25/9.7/10+4.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 674 @9.831s 1%: 0.018+5.2+0.11 ms clock, 0.30+0.17/19/26+1.8 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 675 @9.992s 1%: 0.019+3.0+0.16 ms clock, 0.30+0.24/9.3/8.8+2.5 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5973304 HeapInuse: 6340608 HeapObjects: 84737 HeapIdle 11190272 HeapReleased 0 HeapSys 17530880
gc 676 @10.043s 1%: 0.012+0.25+0.22 ms clock, 0.19+0/0.38/0.18+3.5 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 677 @10.044s 1%: 0.050+0.23+0.11 ms clock, 0.80+0/0.46/0.10+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 678 @10.138s 1%: 0.017+2.7+0.13 ms clock, 0.27+0.47/9.2/18+2.1 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 679 @10.272s 1%: 0.017+5.9+0.11 ms clock, 0.28+0.38/20/41+1.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 680 @10.436s 1%: 0.020+2.1+0.11 ms clock, 0.33+0.17/7.1/13+1.7 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5974936 HeapInuse: 6340608 HeapObjects: 84763 HeapIdle 11190272 HeapReleased 0 HeapSys 17530880
gc 681 @10.487s 1%: 0.016+0.27+0.11 ms clock, 0.26+0/0.40/0.37+1.8 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 682 @10.488s 1%: 0.012+0.20+0.088 ms clock, 0.20+0/0.35/0.28+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 683 @10.586s 1%: 0.021+2.8+0.12 ms clock, 0.34+0.50/10/10+2.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 684 @10.725s 1%: 0.019+9.2+0.11 ms clock, 0.31+0.064/26/19+1.8 ms cpu, 7->8->8 MB, 8 MB goal, 16 P
gc 685 @10.899s 1%: 0.019+2.2+0.090 ms clock, 0.31+0.18/7.3/13+1.4 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5650136 HeapInuse: 6021120 HeapObjects: 79227 HeapIdle 11509760 HeapReleased 0 HeapSys 17530880
gc 686 @10.950s 1%: 0.013+0.31+0.13 ms clock, 0.20+0/0.44/0.29+2.1 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 687 @10.951s 1%: 0.012+0.22+0.11 ms clock, 0.19+0/0.29/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 688 @11.046s 1%: 0.019+2.8+0.12 ms clock, 0.30+0.22/9.7/13+2.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 689 @11.178s 1%: 0.018+6.1+0.12 ms clock, 0.29+0.32/21/26+2.0 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 690 @11.343s 1%: 0.020+5.7+0.13 ms clock, 0.32+0.71/5.9/2.5+2.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5974904 HeapInuse: 6340608 HeapObjects: 84761 HeapIdle 11190272 HeapReleased 0 HeapSys 17530880
gc 691 @11.396s 1%: 0.012+0.23+0.15 ms clock, 0.19+0/0.35/0.27+2.5 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 692 @11.397s 1%: 0.017+0.24+0.076 ms clock, 0.27+0/0.31/0.20+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 693 @11.489s 1%: 0.016+3.4+0.10 ms clock, 0.26+0.50/11/11+1.6 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 694 @11.624s 1%: 0.018+7.2+0.14 ms clock, 0.29+4.1/25/24+2.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 695 @11.793s 1%: 0.017+3.0+0.13 ms clock, 0.27+0.29/9.6/9.3+2.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6081016 HeapInuse: 6455296 HeapObjects: 86573 HeapIdle 11075584 HeapReleased 0 HeapSys 17530880
gc 696 @11.844s 1%: 0.014+0.35+0.17 ms clock, 0.23+0/0.41/0.24+2.8 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 697 @11.845s 1%: 0.012+0.24+0.11 ms clock, 0.20+0/0.38/0.23+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 698 @11.939s 1%: 0.014+2.6+0.14 ms clock, 0.23+0.24/9.1/17+2.3 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 699 @12.075s 1%: 0.019+5.0+0.10 ms clock, 0.31+0.32/18/38+1.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 700 @12.239s 1%: 0.017+2.3+0.13 ms clock, 0.28+0.18/7.7/12+2.1 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5937368 HeapInuse: 6307840 HeapObjects: 84123 HeapIdle 11223040 HeapReleased 0 HeapSys 17530880
gc 701 @12.288s 1%: 0.031+0.25+0.20 ms clock, 0.50+0/0.35/0.14+3.3 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 702 @12.290s 1%: 0.011+0.20+0.19 ms clock, 0.17+0/0.40/0.18+3.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 703 @12.382s 1%: 0.015+2.6+0.21 ms clock, 0.24+0.24/8.8/9.6+3.4 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 704 @12.516s 1%: 0.017+4.4+0.11 ms clock, 0.27+0.16/16/22+1.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 705 @12.677s 1%: 0.017+2.9+0.12 ms clock, 0.28+0.25/8.8/11+2.0 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6188728 HeapInuse: 6561792 HeapObjects: 88409 HeapIdle 10969088 HeapReleased 0 HeapSys 17530880
gc 706 @12.731s 1%: 0.029+0.24+0.17 ms clock, 0.47+0/0.30/0.21+2.8 ms cpu, 5->5->0 MB, 7 MB goal, 16 P (forced)
gc 707 @12.733s 1%: 0.021+0.22+0.14 ms clock, 0.34+0/0.34/0.22+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 708 @12.826s 1%: 0.013+2.4+0.18 ms clock, 0.21+0.20/8.4/17+3.0 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 709 @12.961s 1%: 0.018+7.4+0.17 ms clock, 0.28+0.21/23/42+2.8 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 710 @13.128s 1%: 0.020+2.0+0.080 ms clock, 0.32+0.20/7.1/12+1.2 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5794712 HeapInuse: 6160384 HeapObjects: 81691 HeapIdle 11370496 HeapReleased 0 HeapSys 17530880
gc 711 @13.176s 1%: 0.009+0.23+0.11 ms clock, 0.15+0/0.39/0.24+1.8 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 712 @13.177s 1%: 0.007+0.24+0.18 ms clock, 0.11+0/0.40/0.13+2.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 713 @13.270s 1%: 0.019+2.8+0.11 ms clock, 0.31+0.24/8.4/12+1.7 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 714 @13.405s 1%: 0.020+6.3+0.10 ms clock, 0.33+0.42/23/39+1.6 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 715 @13.576s 1%: 0.021+3.1+0.12 ms clock, 0.34+0.26/8.5/13+1.9 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 6027240 HeapInuse: 6397952 HeapObjects: 85655 HeapIdle 11132928 HeapReleased 0 HeapSys 17530880
gc 716 @13.627s 1%: 0.014+0.33+0.14 ms clock, 0.22+0/0.43/0.20+2.2 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
gc 717 @13.629s 1%: 0.023+0.22+0.11 ms clock, 0.36+0/0.29/0.24+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 718 @13.723s 1%: 0.016+3.7+0.13 ms clock, 0.25+0.27/12/5.8+2.1 ms cpu, 4->4->4 MB, 5 MB goal, 16 P
gc 719 @13.857s 1%: 0.016+6.5+0.13 ms clock, 0.27+0.25/19/38+2.1 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 720 @14.026s 1%: 0.018+2.6+0.14 ms clock, 0.28+0.13/8.5/13+2.3 ms cpu, 15->15->3 MB, 16 MB goal, 16 P
HeapAlloc: 5794168 HeapInuse: 6160384 HeapObjects: 81679 HeapIdle 11370496 HeapReleased 0 HeapSys 17530880
gc 721 @14.081s 1%: 0.036+0.22+0.13 ms clock, 0.57+0/0.40/0.25+2.1 ms cpu, 5->5->0 MB, 6 MB goal, 16 P (forced)
  100000	      1921 ns/op	      96 B/op	       2 allocs/op
gc 722 @14.083s 1%: 0.013+0.23+0.16 ms clock, 0.22+0/0.40/0.26+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 723 @14.139s 1%: 0.017+0.70+0.12 ms clock, 0.27+0.14/1.7/1.1+1.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 724 @14.201s 1%: 0.013+0.66+0.083 ms clock, 0.21+0.12/1.8/1.9+1.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 725 @14.259s 1%: 0.013+0.62+0.12 ms clock, 0.21+0.13/1.5/1.2+1.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797368 HeapInuse: 2187264 HeapObjects: 13551 HeapIdle 15343616 HeapReleased 0 HeapSys 17530880
gc 726 @14.266s 1%: 0.012+0.24+0.12 ms clock, 0.20+0/0.39/0.26+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E5-16                  	gc 727 @14.267s 1%: 0.014+0.21+0.12 ms clock, 0.23+0/0.36/0.18+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 728 @14.340s 1%: 0.014+0.76+0.088 ms clock, 0.23+0.27/1.8/1.3+1.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 729 @14.401s 1%: 0.016+0.76+0.10 ms clock, 0.26+0.23/1.7/2.4+1.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 730 @14.458s 1%: 0.013+0.62+0.17 ms clock, 0.20+0.19/1.3/1.6+2.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797496 HeapInuse: 2179072 HeapObjects: 13553 HeapIdle 15351808 HeapReleased 0 HeapSys 17530880
gc 731 @14.466s 1%: 0.010+0.24+0.10 ms clock, 0.16+0/0.41/0.20+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 732 @14.466s 1%: 0.029+0.16+0.11 ms clock, 0.46+0/0.19/0.13+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 733 @14.540s 1%: 0.018+0.73+0.11 ms clock, 0.28+0.16/1.5/0.89+1.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 734 @14.583s 1%: 0.018+0.79+0.10 ms clock, 0.29+0.16/2.0/1.7+1.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 735 @14.627s 1%: 0.019+0.72+0.12 ms clock, 0.31+0.23/1.5/1.1+2.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797608 HeapInuse: 2170880 HeapObjects: 13555 HeapIdle 15360000 HeapReleased 0 HeapSys 17530880
gc 736 @14.635s 1%: 0.012+0.25+0.095 ms clock, 0.20+0/0.37/0.17+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 737 @14.635s 1%: 0.012+0.28+0.084 ms clock, 0.19+0/0.32/0.16+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 738 @14.708s 1%: 0.017+0.83+0.11 ms clock, 0.28+0.10/2.2/1.3+1.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 739 @14.769s 1%: 0.014+1.7+0.11 ms clock, 0.22+0.12/1.7/4.4+1.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 740 @14.827s 1%: 0.018+0.83+0.14 ms clock, 0.29+0.10/1.9/1.6+2.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797688 HeapInuse: 2187264 HeapObjects: 13555 HeapIdle 15343616 HeapReleased 0 HeapSys 17530880
gc 741 @14.838s 1%: 0.010+0.27+0.13 ms clock, 0.17+0/0.44/0.23+2.1 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 742 @14.838s 1%: 0.011+0.29+0.14 ms clock, 0.19+0/0.36/0.30+2.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 743 @14.913s 1%: 0.023+0.79+0.11 ms clock, 0.37+0.16/2.0/1.6+1.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 744 @14.979s 1%: 0.018+0.86+0.14 ms clock, 0.30+0.10/2.6/1.4+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 745 @15.040s 1%: 0.019+0.72+0.18 ms clock, 0.31+0.15/1.8/0.80+2.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797704 HeapInuse: 2170880 HeapObjects: 13556 HeapIdle 15360000 HeapReleased 0 HeapSys 17530880
gc 746 @15.048s 1%: 0.011+0.20+0.27 ms clock, 0.17+0/0.35/0.15+4.4 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 747 @15.049s 1%: 0.011+0.19+0.087 ms clock, 0.18+0/0.34/0.28+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 748 @15.122s 1%: 0.018+0.75+0.11 ms clock, 0.29+0.20/1.8/1.4+1.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 749 @15.185s 1%: 0.018+1.0+0.096 ms clock, 0.29+0.24/2.7/1.4+1.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 750 @15.242s 1%: 0.013+0.74+0.20 ms clock, 0.20+0.25/1.7/1.1+3.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797784 HeapInuse: 2170880 HeapObjects: 13556 HeapIdle 15360000 HeapReleased 0 HeapSys 17530880
gc 751 @15.249s 1%: 0.008+0.27+0.17 ms clock, 0.13+0/0.36/0.14+2.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 752 @15.249s 1%: 0.016+0.20+0.11 ms clock, 0.26+0/0.27/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 753 @15.321s 1%: 0.014+0.73+0.14 ms clock, 0.23+0.17/1.7/1.6+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 754 @15.377s 1%: 0.013+0.79+0.15 ms clock, 0.21+0.28/1.8/2.1+2.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 755 @15.436s 1%: 0.015+0.73+0.14 ms clock, 0.24+0.38/1.6/1.3+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1797880 HeapInuse: 2162688 HeapObjects: 13557 HeapIdle 15368192 HeapReleased 0 HeapSys 17530880
gc 756 @15.443s 1%: 0.038+0.25+0.14 ms clock, 0.60+0/0.44/0.23+2.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 757 @15.444s 1%: 0.009+0.23+0.17 ms clock, 0.14+0/0.42/0.23+2.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 758 @15.517s 1%: 0.018+0.84+0.14 ms clock, 0.28+0.14/2.0/1.3+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 759 @15.579s 1%: 0.015+0.91+0.099 ms clock, 0.24+0.37/2.5/2.4+1.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 760 @15.635s 1%: 0.017+0.70+0.13 ms clock, 0.27+0.10/1.8/1.5+2.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798008 HeapInuse: 2179072 HeapObjects: 13560 HeapIdle 15351808 HeapReleased 0 HeapSys 17530880
gc 761 @15.642s 1%: 0.010+0.22+0.10 ms clock, 0.17+0/0.34/0.24+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 762 @15.643s 1%: 0.007+0.19+0.11 ms clock, 0.11+0/0.32/0.22+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 763 @15.714s 1%: 0.015+1.0+0.12 ms clock, 0.24+0.085/2.2/0.75+2.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 764 @15.778s 1%: 0.019+0.83+0.12 ms clock, 0.30+0.37/2.2/2.2+1.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 765 @15.837s 1%: 0.016+0.78+0.16 ms clock, 0.25+0.24/1.7/1.4+2.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798072 HeapInuse: 2179072 HeapObjects: 13559 HeapIdle 15351808 HeapReleased 0 HeapSys 17530880
gc 766 @15.845s 1%: 0.037+0.24+0.12 ms clock, 0.59+0/0.37/0.18+1.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 767 @15.845s 1%: 0.016+0.20+0.11 ms clock, 0.25+0/0.36/0.23+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 768 @15.916s 1%: 0.018+0.73+0.14 ms clock, 0.29+0.20/1.8/1.6+2.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 769 @15.977s 1%: 0.015+0.89+0.20 ms clock, 0.25+0.13/2.2/2.0+3.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 770 @16.035s 1%: 0.016+0.78+0.13 ms clock, 0.26+0.17/1.7/1.1+2.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798184 HeapInuse: 2195456 HeapObjects: 13561 HeapIdle 15335424 HeapReleased 0 HeapSys 17530880
gc 771 @16.042s 1%: 0.036+0.24+0.18 ms clock, 0.58+0/0.39/0.16+2.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 772 @16.043s 1%: 0.019+0.21+0.12 ms clock, 0.31+0/0.36/0.10+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 773 @16.114s 1%: 0.018+0.76+0.12 ms clock, 0.30+0.34/1.7/1.5+2.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 774 @16.175s 1%: 0.014+0.99+0.18 ms clock, 0.22+0.26/1.7/2.3+2.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 775 @16.233s 1%: 0.014+0.76+0.19 ms clock, 0.23+0.18/1.7/1.5+3.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798264 HeapInuse: 2203648 HeapObjects: 13561 HeapIdle 15327232 HeapReleased 0 HeapSys 17530880
gc 776 @16.240s 1%: 0.010+0.25+0.16 ms clock, 0.16+0/0.32/0.24+2.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 777 @16.241s 1%: 0.011+0.23+0.11 ms clock, 0.17+0/0.18/0.25+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 778 @16.314s 1%: 0.017+0.73+0.23 ms clock, 0.27+0.13/2.1/0.83+3.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 779 @16.375s 1%: 0.014+1.0+0.16 ms clock, 0.22+0.25/1.9/2.0+2.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 780 @16.433s 1%: 0.013+0.75+0.21 ms clock, 0.21+0.32/1.6/1.1+3.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798360 HeapInuse: 2187264 HeapObjects: 13562 HeapIdle 15343616 HeapReleased 0 HeapSys 17530880
gc 781 @16.440s 1%: 0.008+0.23+0.16 ms clock, 0.14+0/0.41/0.12+2.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 782 @16.441s 1%: 0.020+0.22+0.16 ms clock, 0.32+0/0.29/0.10+2.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 783 @16.513s 1%: 0.017+0.80+0.13 ms clock, 0.28+0.10/2.0/1.5+2.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 784 @16.575s 1%: 0.014+0.81+0.16 ms clock, 0.22+0.31/1.8/2.2+2.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 785 @16.632s 1%: 0.015+0.68+0.16 ms clock, 0.24+0.18/1.7/1.2+2.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798376 HeapInuse: 2179072 HeapObjects: 13563 HeapIdle 15351808 HeapReleased 0 HeapSys 17530880
gc 786 @16.639s 1%: 0.010+0.25+0.10 ms clock, 0.17+0/0.43/0.15+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 787 @16.640s 1%: 0.021+0.20+0.12 ms clock, 0.33+0/0.31/0.16+2.0 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 788 @16.711s 1%: 0.013+0.83+0.15 ms clock, 0.22+0.21/1.8/1.2+2.5 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 789 @16.771s 1%: 0.012+0.86+0.14 ms clock, 0.20+0.45/2.3/1.3+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 790 @16.829s 1%: 0.012+0.74+0.20 ms clock, 0.20+0.35/1.6/1.4+3.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798360 HeapInuse: 2170880 HeapObjects: 13562 HeapIdle 15360000 HeapReleased 0 HeapSys 17530880
gc 791 @16.836s 1%: 0.009+0.20+0.14 ms clock, 0.15+0/0.37/0.22+2.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 792 @16.837s 1%: 0.021+0.22+0.11 ms clock, 0.33+0/0.40/0.18+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 793 @16.908s 1%: 0.013+0.72+0.11 ms clock, 0.21+0.19/1.5/1.4+1.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 794 @16.969s 1%: 0.013+0.83+0.15 ms clock, 0.21+0.43/2.3/2.3+2.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 795 @17.028s 1%: 0.016+0.77+0.21 ms clock, 0.26+0.20/1.9/1.0+3.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798456 HeapInuse: 2162688 HeapObjects: 13563 HeapIdle 15368192 HeapReleased 0 HeapSys 17530880
gc 796 @17.035s 1%: 0.011+0.28+0.15 ms clock, 0.18+0/0.45/0.22+2.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 797 @17.036s 1%: 0.011+0.23+0.15 ms clock, 0.18+0/0.39/0.32+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 798 @17.109s 1%: 0.017+0.70+0.18 ms clock, 0.27+0.19/1.9/1.9+2.9 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 799 @17.171s 1%: 0.018+1.0+0.12 ms clock, 0.30+0.31/2.3/2.1+2.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
gc 800 @17.231s 1%: 0.018+0.66+0.10 ms clock, 0.30+0.22/1.5/1.6+1.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 1798552 HeapInuse: 2170880 HeapObjects: 13564 HeapIdle 15360000 HeapReleased 0 HeapSys 17530880
gc 801 @17.238s 1%: 0.011+0.20+0.10 ms clock, 0.17+0/0.30/0.27+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
  100000	      2020 ns/op	      96 B/op	       2 allocs/op
gc 802 @17.239s 1%: 0.053+0.24+0.092 ms clock, 0.85+0/0.45/0.19+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 803 @17.389s 1%: 0.018+0.51+0.10 ms clock, 0.29+0.19/0.88/0.71+1.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2963568 HeapInuse: 3334144 HeapObjects: 40662 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 804 @17.479s 1%: 0.009+0.27+0.14 ms clock, 0.15+0/0.36/0.15+2.3 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E5-16          	gc 805 @17.480s 1%: 0.018+0.24+0.12 ms clock, 0.29+0/0.36/0.24+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 806 @17.633s 1%: 0.017+1.1+0.078 ms clock, 0.28+0.14/1.2/0.84+1.2 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2963584 HeapInuse: 3334144 HeapObjects: 40661 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 807 @17.720s 1%: 0.039+0.22+0.16 ms clock, 0.62+0/0.39/0.21+2.7 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 808 @17.721s 1%: 0.014+0.27+0.13 ms clock, 0.23+0/0.38/0.22+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 809 @17.872s 1%: 0.017+0.50+0.17 ms clock, 0.28+0.18/0.96/0.86+2.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2963584 HeapInuse: 3334144 HeapObjects: 40661 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 810 @17.959s 1%: 0.043+0.28+0.079 ms clock, 0.70+0/0.37/0.40+1.2 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 811 @17.960s 1%: 0.009+0.20+0.10 ms clock, 0.14+0/0.35/0.23+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 812 @18.110s 1%: 0.017+0.50+0.10 ms clock, 0.27+0.30/1.1/0.92+1.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2963680 HeapInuse: 3334144 HeapObjects: 40662 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 813 @18.198s 1%: 0.010+0.26+0.15 ms clock, 0.17+0/0.38/0.25+2.4 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 814 @18.199s 1%: 0.007+0.27+0.11 ms clock, 0.12+0/0.30/0.13+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 815 @18.347s 1%: 0.017+0.49+0.19 ms clock, 0.27+0.29/0.96/0.52+3.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2963776 HeapInuse: 3334144 HeapObjects: 40663 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 816 @18.401s 1%: 0.009+0.22+0.15 ms clock, 0.15+0/0.33/0.26+2.4 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 817 @18.401s 1%: 0.012+0.20+0.095 ms clock, 0.19+0/0.26/0.13+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 818 @18.551s 1%: 0.014+0.54+0.14 ms clock, 0.23+0.19/1.1/0.95+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2963904 HeapInuse: 3334144 HeapObjects: 40666 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 819 @18.640s 1%: 0.010+0.25+0.14 ms clock, 0.17+0/0.41/0.31+2.3 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 820 @18.641s 1%: 0.011+0.22+0.20 ms clock, 0.17+0/0.32/0.24+3.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 821 @18.789s 1%: 0.016+0.55+0.10 ms clock, 0.25+0.14/1.0/0.28+1.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964000 HeapInuse: 3334144 HeapObjects: 40667 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 822 @18.878s 1%: 0.025+0.27+0.093 ms clock, 0.40+0/0.49/0.17+1.4 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 823 @18.879s 1%: 0.010+0.22+0.13 ms clock, 0.16+0/0.30/0.19+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 824 @19.032s 1%: 0.021+0.44+0.090 ms clock, 0.34+0.14/1.1/1.4+1.4 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964080 HeapInuse: 3334144 HeapObjects: 40667 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 825 @19.122s 1%: 0.011+0.24+0.16 ms clock, 0.18+0/0.39/0.25+2.5 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 826 @19.123s 1%: 0.015+0.24+0.10 ms clock, 0.25+0/0.36/0.19+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 827 @19.271s 1%: 0.020+0.49+0.11 ms clock, 0.33+0.11/1.0/1.4+1.7 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964176 HeapInuse: 3334144 HeapObjects: 40668 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 828 @19.359s 1%: 0.013+0.33+0.14 ms clock, 0.20+0/0.39/0.33+2.2 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 829 @19.360s 1%: 0.012+0.23+0.087 ms clock, 0.20+0/0.44/0.23+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 830 @19.507s 1%: 0.022+0.52+0.10 ms clock, 0.36+0.086/1.1/1.0+1.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964176 HeapInuse: 3334144 HeapObjects: 40668 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 831 @19.595s 1%: 0.010+0.26+0.45 ms clock, 0.16+0/0.51/0.16+7.2 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 832 @19.596s 1%: 0.009+0.20+0.12 ms clock, 0.15+0/0.36/0.22+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 833 @19.745s 1%: 0.016+0.55+0.19 ms clock, 0.26+0.11/1.1/0.90+3.0 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964192 HeapInuse: 3334144 HeapObjects: 40669 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 834 @19.832s 1%: 0.011+0.25+0.14 ms clock, 0.18+0/0.34/0.19+2.2 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 835 @19.832s 1%: 0.028+0.18+0.11 ms clock, 0.46+0/0.34/0.17+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 836 @19.978s 1%: 0.016+0.65+0.17 ms clock, 0.27+0.14/1.2/0.62+2.8 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964176 HeapInuse: 3334144 HeapObjects: 40668 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 837 @20.064s 1%: 0.011+0.23+0.17 ms clock, 0.18+0/0.41/0.097+2.8 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 838 @20.064s 1%: 0.008+0.25+0.15 ms clock, 0.13+0/0.43/0.15+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 839 @20.213s 1%: 0.018+0.85+0.14 ms clock, 0.30+0.15/1.0/0.70+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964272 HeapInuse: 3334144 HeapObjects: 40669 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 840 @20.300s 1%: 0.010+0.24+0.16 ms clock, 0.17+0/0.36/0.23+2.6 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 841 @20.301s 1%: 0.012+0.23+0.18 ms clock, 0.20+0/0.31/0.23+2.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 842 @20.446s 1%: 0.014+0.63+0.22 ms clock, 0.23+0.21/1.1/0.35+3.6 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964272 HeapInuse: 3334144 HeapObjects: 40669 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 843 @20.536s 1%: 0.010+0.24+0.14 ms clock, 0.16+0/0.34/0.15+2.3 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 844 @20.537s 1%: 0.017+0.26+0.13 ms clock, 0.28+0/0.34/0.17+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 845 @20.685s 1%: 0.017+0.59+0.14 ms clock, 0.27+0.16/1.0/0.60+2.3 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964384 HeapInuse: 3334144 HeapObjects: 40671 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 846 @20.772s 1%: 0.012+0.30+0.18 ms clock, 0.19+0/0.35/0.36+2.9 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
gc 847 @20.773s 1%: 0.059+0.23+0.15 ms clock, 0.95+0/0.35/0.26+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 848 @20.923s 0%: 0.017+0.64+0.13 ms clock, 0.27+0.15/0.97/0.81+2.1 ms cpu, 4->4->1 MB, 5 MB goal, 16 P
HeapAlloc: 2964496 HeapInuse: 3334144 HeapObjects: 40673 HeapIdle 14196736 HeapReleased 0 HeapSys 17530880
gc 849 @21.010s 0%: 0.010+0.31+0.11 ms clock, 0.16+0/0.43/0.20+1.8 ms cpu, 2->2->0 MB, 4 MB goal, 16 P (forced)
  100000	      2368 ns/op	      56 B/op	       1 allocs/op
gc 850 @21.011s 0%: 0.050+0.24+0.097 ms clock, 0.80+0/0.29/0.27+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3866592 HeapInuse: 4358144 HeapObjects: 4663 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 851 @21.040s 0%: 0.010+0.22+0.10 ms clock, 0.16+0/0.38/0.34+1.6 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E5-16                   	gc 852 @21.041s 0%: 0.012+0.25+0.14 ms clock, 0.19+0/0.47/0.26+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3846288 HeapInuse: 4358144 HeapObjects: 4644 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 853 @21.070s 0%: 0.010+0.28+0.17 ms clock, 0.17+0/0.35/0.40+2.7 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 854 @21.070s 0%: 0.13+0.25+0.096 ms clock, 2.1+0/0.29/0.17+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3869008 HeapInuse: 4358144 HeapObjects: 4688 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 855 @21.100s 0%: 0.011+0.36+0.083 ms clock, 0.17+0/0.16/0.37+1.3 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 856 @21.101s 0%: 0.007+0.20+0.13 ms clock, 0.12+0/0.39/0.19+2.1 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3840936 HeapInuse: 4349952 HeapObjects: 4589 HeapIdle 13180928 HeapReleased 0 HeapSys 17530880
gc 857 @21.130s 0%: 0.010+0.40+0.13 ms clock, 0.16+0/0.45/0.12+2.2 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 858 @21.131s 0%: 0.071+0.25+0.12 ms clock, 1.1+0/0.46/0.28+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3859472 HeapInuse: 4349952 HeapObjects: 4590 HeapIdle 13180928 HeapReleased 0 HeapSys 17530880
gc 859 @21.160s 0%: 0.051+0.35+0.11 ms clock, 0.83+0/0.45/0.090+1.8 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 860 @21.160s 0%: 0.038+0.21+0.11 ms clock, 0.61+0/0.31/0.17+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3867352 HeapInuse: 4358144 HeapObjects: 4670 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 861 @21.189s 0%: 0.009+0.35+0.19 ms clock, 0.14+0/0.62/0.19+3.1 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 862 @21.190s 0%: 0.048+0.23+0.18 ms clock, 0.77+0/0.32/0.20+2.9 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3871784 HeapInuse: 4366336 HeapObjects: 4711 HeapIdle 13164544 HeapReleased 0 HeapSys 17530880
gc 863 @21.219s 0%: 0.011+0.26+0.14 ms clock, 0.18+0/0.55/0.18+2.3 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 864 @21.219s 0%: 0.072+0.22+0.16 ms clock, 1.1+0/0.31/0.25+2.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3868640 HeapInuse: 4358144 HeapObjects: 4685 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 865 @21.248s 0%: 0.052+0.29+0.15 ms clock, 0.83+0/0.51/0.25+2.5 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 866 @21.249s 0%: 0.067+0.23+0.092 ms clock, 1.0+0/0.35/0.28+1.4 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3839096 HeapInuse: 4349952 HeapObjects: 4565 HeapIdle 13180928 HeapReleased 0 HeapSys 17530880
gc 867 @21.279s 0%: 0.036+0.25+0.17 ms clock, 0.58+0/0.26/0.35+2.7 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 868 @21.280s 0%: 0.017+0.19+0.20 ms clock, 0.28+0/0.34/0.24+3.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3865896 HeapInuse: 4358144 HeapObjects: 4654 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 869 @21.309s 0%: 0.010+0.30+0.11 ms clock, 0.17+0/0.50/0.28+1.8 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 870 @21.310s 0%: 0.021+0.22+0.11 ms clock, 0.33+0/0.31/0.18+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3880536 HeapInuse: 4374528 HeapObjects: 4807 HeapIdle 13156352 HeapReleased 0 HeapSys 17530880
gc 871 @21.342s 0%: 0.012+0.47+0.11 ms clock, 0.19+0/0.48/0.10+1.8 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 872 @21.342s 0%: 0.019+0.19+0.12 ms clock, 0.31+0/0.34/0.29+2.0 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3844536 HeapInuse: 4358144 HeapObjects: 4623 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 873 @21.373s 0%: 0.010+0.40+0.11 ms clock, 0.16+0/0.34/0.32+1.8 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 874 @21.374s 0%: 0.016+0.25+0.11 ms clock, 0.26+0/0.47/0.085+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3866528 HeapInuse: 4358144 HeapObjects: 4662 HeapIdle 13172736 HeapReleased 0 HeapSys 17530880
gc 875 @21.402s 0%: 0.008+0.28+0.11 ms clock, 0.13+0/0.52/0.23+1.9 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
gc 876 @21.403s 0%: 0.027+0.22+0.10 ms clock, 0.43+0/0.37/0.18+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 3858552 HeapInuse: 4349952 HeapObjects: 4578 HeapIdle 13180928 HeapReleased 0 HeapSys 17530880
gc 877 @21.432s 0%: 0.027+0.24+0.13 ms clock, 0.43+0/0.41/0.22+2.1 ms cpu, 3->3->1 MB, 4 MB goal, 16 P (forced)
  100000	       297 ns/op	      36 B/op	       0 allocs/op
gc 878 @21.433s 0%: 0.010+0.27+0.10 ms clock, 0.16+0/0.38/0.19+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 879 @21.449s 0%: 0.012+0.31+0.17 ms clock, 0.19+0.13/0.33/0.24+2.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3721696 HeapInuse: 4194304 HeapObjects: 3556 HeapIdle 13336576 HeapReleased 0 HeapSys 17530880
gc 880 @21.474s 0%: 0.010+0.22+0.10 ms clock, 0.16+0/0.45/0.28+1.7 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E5-16                    	gc 881 @21.475s 1%: 0.051+0.20+0.080 ms clock, 0.83+0/0.37/0.22+1.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 882 @21.490s 1%: 0.014+0.26+0.095 ms clock, 0.22+0.13/0.37/0.25+1.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3709312 HeapInuse: 4186112 HeapObjects: 3464 HeapIdle 13312000 HeapReleased 0 HeapSys 17498112
gc 883 @21.518s 0%: 0.010+0.25+0.10 ms clock, 0.16+0/0.48/0.21+1.6 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 884 @21.518s 0%: 0.009+0.22+0.080 ms clock, 0.14+0/0.36/0.34+1.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 885 @21.533s 0%: 0.011+0.30+0.11 ms clock, 0.18+0.15/0.24/0.39+1.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3722240 HeapInuse: 4194304 HeapObjects: 3548 HeapIdle 13303808 HeapReleased 0 HeapSys 17498112
gc 886 @21.561s 0%: 0.009+0.31+0.17 ms clock, 0.14+0/0.36/0.56+2.7 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 887 @21.562s 1%: 0.009+0.19+0.10 ms clock, 0.15+0/0.33/0.28+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 888 @21.577s 1%: 0.015+0.31+0.10 ms clock, 0.24+0.10/0.19/0.43+1.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3703648 HeapInuse: 4177920 HeapObjects: 3405 HeapIdle 13320192 HeapReleased 0 HeapSys 17498112
gc 889 @21.602s 0%: 0.010+0.43+0.11 ms clock, 0.16+0/0.46/0.086+1.8 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 890 @21.602s 1%: 0.035+0.25+0.10 ms clock, 0.57+0/0.38/0.21+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 891 @21.615s 1%: 0.012+0.20+0.10 ms clock, 0.19+0.093/0.24/0.29+1.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3716480 HeapInuse: 4186112 HeapObjects: 3488 HeapIdle 13312000 HeapReleased 0 HeapSys 17498112
gc 892 @21.631s 1%: 0.006+0.26+0.13 ms clock, 0.10+0/0.46/0.15+2.0 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 893 @21.632s 1%: 0.007+0.22+0.10 ms clock, 0.12+0/0.37/0.27+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 894 @21.647s 1%: 0.016+0.24+0.12 ms clock, 0.25+0/0.35/0.31+2.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3719264 HeapInuse: 4194304 HeapObjects: 3517 HeapIdle 13303808 HeapReleased 0 HeapSys 17498112
gc 895 @21.673s 1%: 0.009+0.30+0.11 ms clock, 0.15+0/0.49/0.34+1.8 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 896 @21.674s 1%: 0.017+0.17+0.14 ms clock, 0.27+0/0.29/0.23+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 897 @21.684s 1%: 0.012+0.27+0.092 ms clock, 0.19+0.092/0.23/0.29+1.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3715232 HeapInuse: 4186112 HeapObjects: 3475 HeapIdle 13312000 HeapReleased 0 HeapSys 17498112
gc 898 @21.702s 1%: 0.007+0.24+0.084 ms clock, 0.12+0/0.34/0.29+1.3 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 899 @21.703s 1%: 0.009+0.20+0.084 ms clock, 0.15+0/0.29/0.21+1.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 900 @21.718s 1%: 0.015+0.26+0.10 ms clock, 0.25+0.12/0.17/0.26+1.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3721088 HeapInuse: 4194304 HeapObjects: 3536 HeapIdle 13303808 HeapReleased 0 HeapSys 17498112
gc 901 @21.744s 1%: 0.008+0.28+0.099 ms clock, 0.14+0/0.34/0.49+1.5 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 902 @21.745s 1%: 0.009+0.20+0.096 ms clock, 0.14+0/0.37/0.27+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 903 @21.761s 1%: 0.013+0.28+0.092 ms clock, 0.21+0.092/0.33/0.29+1.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3718592 HeapInuse: 4194304 HeapObjects: 3510 HeapIdle 13303808 HeapReleased 0 HeapSys 17498112
gc 904 @21.779s 1%: 0.039+0.20+0.12 ms clock, 0.63+0/0.33/0.32+2.0 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 905 @21.779s 1%: 0.008+0.20+0.092 ms clock, 0.14+0/0.31/0.20+1.4 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 906 @21.795s 1%: 0.012+0.44+0.17 ms clock, 0.20+0.10/0.21/0.26+2.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3723584 HeapInuse: 4194304 HeapObjects: 3562 HeapIdle 13303808 HeapReleased 0 HeapSys 17498112
gc 907 @21.819s 1%: 0.010+0.26+0.12 ms clock, 0.17+0/0.46/0.27+1.9 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 908 @21.820s 1%: 0.016+0.20+0.099 ms clock, 0.25+0/0.36/0.13+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 909 @21.840s 1%: 0.016+0.23+0.097 ms clock, 0.25+0.012/0.32/0.19+1.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3719072 HeapInuse: 4202496 HeapObjects: 3515 HeapIdle 13295616 HeapReleased 0 HeapSys 17498112
gc 910 @21.857s 1%: 0.021+0.21+0.15 ms clock, 0.34+0/0.28/0.25+2.5 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
gc 911 @21.858s 1%: 0.010+0.18+0.13 ms clock, 0.17+0/0.28/0.18+2.1 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 912 @21.873s 1%: 0.011+0.29+0.16 ms clock, 0.18+0.14/0.19/0.38+2.6 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 3721952 HeapInuse: 4194304 HeapObjects: 3545 HeapIdle 13303808 HeapReleased 0 HeapSys 17498112
gc 913 @21.899s 1%: 0.012+0.24+0.098 ms clock, 0.20+0/0.31/0.33+1.5 ms cpu, 3->3->1 MB, 6 MB goal, 16 P (forced)
  100000	       118 ns/op	       8 B/op	       0 allocs/op
gc 914 @21.899s 1%: 0.064+0.18+0.084 ms clock, 1.0+0/0.25/0.18+1.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 804104 HeapInuse: 1236992 HeapObjects: 1404 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 915 @21.908s 1%: 0.005+0.15+0.086 ms clock, 0.093+0/0.26/0.12+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E5-16           	gc 916 @21.909s 1%: 0.013+0.17+0.11 ms clock, 0.21+0/0.30/0.16+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 796752 HeapInuse: 1228800 HeapObjects: 1325 HeapIdle 16269312 HeapReleased 0 HeapSys 17498112
gc 917 @21.923s 1%: 0.031+0.19+0.11 ms clock, 0.49+0/0.34/0.19+1.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 918 @21.923s 1%: 0.016+0.28+0.11 ms clock, 0.26+0/0.31/0.28+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 800248 HeapInuse: 1236992 HeapObjects: 1362 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 919 @21.939s 1%: 0.058+0.19+0.14 ms clock, 0.94+0/0.26/0.24+2.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 920 @21.939s 1%: 0.010+0.18+0.10 ms clock, 0.16+0/0.36/0.11+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 800744 HeapInuse: 1236992 HeapObjects: 1369 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 921 @21.953s 1%: 0.008+0.19+0.18 ms clock, 0.13+0/0.21/0.20+2.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 922 @21.954s 1%: 0.044+0.16+0.081 ms clock, 0.70+0/0.29/0.18+1.2 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 798120 HeapInuse: 1228800 HeapObjects: 1341 HeapIdle 16269312 HeapReleased 0 HeapSys 17498112
gc 923 @21.963s 1%: 0.006+0.24+0.098 ms clock, 0.096+0/0.45/0.18+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 924 @21.964s 1%: 0.008+0.18+0.073 ms clock, 0.13+0/0.37/0.21+1.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 798168 HeapInuse: 1228800 HeapObjects: 1340 HeapIdle 16269312 HeapReleased 0 HeapSys 17498112
gc 925 @21.972s 1%: 0.007+0.24+0.10 ms clock, 0.11+0/0.35/0.19+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 926 @21.973s 1%: 0.010+0.15+0.10 ms clock, 0.16+0/0.28/0.20+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 800344 HeapInuse: 1236992 HeapObjects: 1362 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 927 @21.987s 1%: 0.010+0.22+0.10 ms clock, 0.17+0/0.38/0.21+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 928 @21.988s 1%: 0.012+0.17+0.087 ms clock, 0.19+0/0.29/0.16+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 801544 HeapInuse: 1236992 HeapObjects: 1375 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 929 @21.997s 1%: 0.007+0.22+0.11 ms clock, 0.11+0/0.36/0.11+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 930 @21.998s 1%: 0.039+0.16+0.090 ms clock, 0.63+0/0.26/0.31+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 803048 HeapInuse: 1236992 HeapObjects: 1391 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 931 @22.006s 1%: 0.008+0.18+0.16 ms clock, 0.13+0/0.37/0.19+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 932 @22.007s 1%: 0.007+0.20+0.10 ms clock, 0.12+0/0.27/0.10+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 806392 HeapInuse: 1236992 HeapObjects: 1422 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 933 @22.021s 1%: 0.008+0.30+0.11 ms clock, 0.13+0/0.41/0.16+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 934 @22.021s 1%: 0.014+0.16+0.10 ms clock, 0.23+0/0.30/0.30+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 797688 HeapInuse: 1228800 HeapObjects: 1336 HeapIdle 16269312 HeapReleased 0 HeapSys 17498112
gc 935 @22.039s 1%: 0.021+0.23+0.10 ms clock, 0.34+0/0.39/0.17+1.7 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 936 @22.039s 1%: 0.007+0.19+0.094 ms clock, 0.12+0/0.29/0.13+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 800816 HeapInuse: 1236992 HeapObjects: 1372 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 937 @22.054s 1%: 0.010+0.19+0.10 ms clock, 0.17+0/0.43/0.27+1.6 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 938 @22.054s 1%: 0.012+0.21+0.091 ms clock, 0.19+0/0.37/0.29+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 801016 HeapInuse: 1236992 HeapObjects: 1369 HeapIdle 16261120 HeapReleased 0 HeapSys 17498112
gc 939 @22.069s 1%: 0.010+0.26+0.094 ms clock, 0.16+0/0.36/0.29+1.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
  100000	       144 ns/op	       5 B/op	       0 allocs/op
gc 940 @22.070s 1%: 0.022+0.21+0.093 ms clock, 0.35+0/0.29/0.21+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1857624 HeapInuse: 2220032 HeapObjects: 696 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 941 @22.081s 1%: 0.008+0.26+0.089 ms clock, 0.14+0/0.42/0.37+1.4 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E5-16    	gc 942 @22.082s 1%: 0.031+0.22+0.10 ms clock, 0.49+0/0.18/0.29+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1857768 HeapInuse: 2220032 HeapObjects: 698 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 943 @22.093s 1%: 0.010+0.16+0.085 ms clock, 0.17+0/0.25/0.24+1.3 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 944 @22.093s 1%: 0.007+0.16+0.11 ms clock, 0.12+0/0.34/0.32+1.8 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1857752 HeapInuse: 2220032 HeapObjects: 697 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 945 @22.110s 1%: 0.014+0.25+0.11 ms clock, 0.23+0/0.43/0.45+1.8 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 946 @22.111s 1%: 0.008+0.23+0.078 ms clock, 0.14+0/0.41/0.31+1.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1857848 HeapInuse: 2220032 HeapObjects: 698 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 947 @22.124s 1%: 0.010+0.27+0.10 ms clock, 0.16+0/0.42/0.29+1.6 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 948 @22.124s 1%: 0.012+0.18+0.097 ms clock, 0.19+0/0.36/0.27+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858040 HeapInuse: 2220032 HeapObjects: 700 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 949 @22.136s 1%: 0.011+0.22+0.090 ms clock, 0.18+0/0.28/0.29+1.4 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 950 @22.137s 1%: 0.008+0.22+0.095 ms clock, 0.14+0/0.37/0.32+1.5 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858136 HeapInuse: 2220032 HeapObjects: 701 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 951 @22.148s 1%: 0.009+0.31+0.13 ms clock, 0.15+0/0.35/0.14+2.1 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 952 @22.149s 1%: 0.016+0.19+0.10 ms clock, 0.26+0/0.34/0.13+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858136 HeapInuse: 2220032 HeapObjects: 701 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 953 @22.161s 1%: 0.010+0.24+0.099 ms clock, 0.16+0/0.36/0.28+1.5 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 954 @22.162s 1%: 0.015+0.24+0.10 ms clock, 0.24+0/0.25/0.25+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858232 HeapInuse: 2220032 HeapObjects: 702 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 955 @22.173s 1%: 0.018+0.23+0.14 ms clock, 0.30+0/0.34/0.27+2.3 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 956 @22.174s 1%: 0.035+0.20+0.10 ms clock, 0.57+0/0.33/0.21+1.6 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858424 HeapInuse: 2220032 HeapObjects: 704 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 957 @22.185s 1%: 0.060+0.25+0.085 ms clock, 0.96+0/0.41/0.10+1.3 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 958 @22.186s 1%: 0.018+0.25+0.13 ms clock, 0.30+0/0.35/0.29+2.1 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858424 HeapInuse: 2220032 HeapObjects: 704 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 959 @22.197s 1%: 0.057+0.24+0.089 ms clock, 0.91+0/0.42/0.22+1.4 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 960 @22.198s 1%: 0.012+0.19+0.075 ms clock, 0.20+0/0.33/0.29+1.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858616 HeapInuse: 2220032 HeapObjects: 706 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 961 @22.209s 1%: 0.023+0.19+0.13 ms clock, 0.37+0/0.35/0.13+2.1 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 962 @22.210s 1%: 0.010+0.44+0.13 ms clock, 0.17+0/0.15/0.42+2.2 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858616 HeapInuse: 2220032 HeapObjects: 706 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 963 @22.223s 1%: 0.045+0.23+0.11 ms clock, 0.72+0/0.43/0.099+1.8 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
gc 964 @22.223s 1%: 0.011+0.29+0.10 ms clock, 0.18+0/0.38/0.14+1.7 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
HeapAlloc: 1858808 HeapInuse: 2220032 HeapObjects: 708 HeapIdle 15278080 HeapReleased 0 HeapSys 17498112
gc 965 @22.235s 1%: 0.012+0.21+0.14 ms clock, 0.20+0/0.36/0.16+2.3 ms cpu, 1->1->1 MB, 4 MB goal, 16 P (forced)
  100000	       114 ns/op	      16 B/op	       0 allocs/op
gc 966 @22.236s 1%: 0.051+0.27+0.14 ms clock, 0.81+0/0.44/0.15+2.3 ms cpu, 1->1->0 MB, 4 MB goal, 16 P (forced)
gc 967 @22.238s 1%: 0.016+47+5.8 ms clock, 0.26+0/0.28/31+94 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8248616 HeapInuse: 8617984 HeapObjects: 707 HeapIdle 8880128 HeapReleased 0 HeapSys 17498112
gc 968 @22.703s 1%: 0.011+0.29+0.13 ms clock, 0.18+0/0.41/0.27+2.0 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
#BenchmarkSort1E6-16                               	gc 969 @22.704s 1%: 0.061+0.27+0.11 ms clock, 0.98+0/0.41/0.25+1.7 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 970 @22.706s 1%: 0.016+45+6.0 ms clock, 0.26+0/0.36/29+96 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8248648 HeapInuse: 8617984 HeapObjects: 708 HeapIdle 8880128 HeapReleased 0 HeapSys 17498112
gc 971 @23.158s 1%: 0.028+0.27+0.19 ms clock, 0.45+0/0.46/0.25+3.0 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 972 @23.159s 1%: 0.048+0.20+0.097 ms clock, 0.77+0/0.35/0.15+1.5 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 973 @23.161s 1%: 0.013+44+5.7 ms clock, 0.21+0/0.043/29+91 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8248744 HeapInuse: 8626176 HeapObjects: 709 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 974 @23.613s 1%: 0.010+0.31+0.19 ms clock, 0.17+0/0.41/0.23+3.0 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 975 @23.614s 1%: 0.025+0.25+0.17 ms clock, 0.41+0/0.35/0.20+2.8 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 976 @23.616s 1%: 0.014+44+5.8 ms clock, 0.22+0/0.19/29+93 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8248744 HeapInuse: 8626176 HeapObjects: 709 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 977 @24.065s 1%: 0.011+0.32+0.12 ms clock, 0.18+0/0.37/0.39+2.0 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 978 @24.066s 1%: 0.016+0.22+0.13 ms clock, 0.26+0/0.41/0.19+2.1 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 979 @24.068s 1%: 0.012+44+5.7 ms clock, 0.20+0/0.24/29+91 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8248840 HeapInuse: 8626176 HeapObjects: 710 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 980 @24.515s 1%: 0.012+0.21+0.11 ms clock, 0.19+0/0.36/0.28+1.9 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 981 @24.516s 1%: 0.010+0.23+0.11 ms clock, 0.16+0/0.33/0.29+1.9 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 982 @24.518s 1%: 0.015+29+15 ms clock, 0.24+0/0.14/30+245 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249032 HeapInuse: 8626176 HeapObjects: 712 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 983 @24.967s 1%: 0.013+0.23+0.16 ms clock, 0.21+0/0.31/0.22+2.6 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 984 @24.967s 1%: 0.042+0.24+0.11 ms clock, 0.67+0/0.48/0.094+1.8 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 985 @24.969s 1%: 0.013+30+15 ms clock, 0.20+0/0.22/29+243 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249128 HeapInuse: 8626176 HeapObjects: 713 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 986 @25.417s 1%: 0.013+0.29+0.16 ms clock, 0.21+0/0.37/0.21+2.7 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 987 @25.417s 1%: 0.006+0.23+0.19 ms clock, 0.11+0/0.34/0.22+3.0 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 988 @25.420s 1%: 0.012+29+15 ms clock, 0.19+0/0.042/29+243 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249224 HeapInuse: 8626176 HeapObjects: 714 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 989 @25.881s 1%: 0.011+0.26+0.090 ms clock, 0.18+0/0.49/0.26+1.4 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 990 @25.881s 1%: 0.008+0.26+0.57 ms clock, 0.13+0/0.30/0.28+9.1 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 991 @25.884s 1%: 0.018+32+16 ms clock, 0.29+0/0.23/32+261 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249416 HeapInuse: 8626176 HeapObjects: 716 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 992 @26.344s 1%: 0.011+0.30+0.19 ms clock, 0.18+0/0.28/0.21+3.1 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 993 @26.345s 1%: 0.070+0.22+0.17 ms clock, 1.1+0/0.40/0.10+2.8 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 994 @26.347s 1%: 0.011+44+5.7 ms clock, 0.18+0/0.23/29+92 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249512 HeapInuse: 8626176 HeapObjects: 717 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 995 @26.800s 1%: 0.012+0.24+0.13 ms clock, 0.20+0/0.37/0.23+2.1 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 996 @26.801s 1%: 0.009+0.23+0.11 ms clock, 0.15+0/0.34/0.28+1.8 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 997 @26.803s 1%: 0.012+44+5.8 ms clock, 0.20+0/0.18/29+92 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249608 HeapInuse: 8626176 HeapObjects: 718 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 998 @27.251s 1%: 0.012+0.32+0.19 ms clock, 0.19+0/0.47/0.089+3.0 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 999 @27.252s 1%: 0.012+0.31+0.16 ms clock, 0.19+0/0.30/0.49+2.7 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 1000 @27.254s 1%: 0.012+44+5.7 ms clock, 0.19+0/0.18/29+92 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249608 HeapInuse: 8626176 HeapObjects: 718 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 1001 @27.701s 1%: 0.016+0.26+0.13 ms clock, 0.26+0/0.43/0.29+2.1 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 1002 @27.702s 1%: 0.009+0.39+0.13 ms clock, 0.14+0/0.42/0.20+2.2 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 1003 @27.705s 1%: 0.014+45+5.7 ms clock, 0.22+0/0.22/30+92 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249608 HeapInuse: 8626176 HeapObjects: 718 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 1004 @28.157s 1%: 0.039+0.24+0.10 ms clock, 0.62+0/0.33/0.32+1.7 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 1005 @28.157s 1%: 0.065+0.28+0.18 ms clock, 1.0+0/0.40/0.31+2.9 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 1006 @28.159s 1%: 0.011+44+5.7 ms clock, 0.17+0/0.37/29+91 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249608 HeapInuse: 8626176 HeapObjects: 718 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 1007 @28.615s 1%: 0.013+0.29+0.12 ms clock, 0.21+0/0.43/0.16+2.0 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
gc 1008 @28.615s 1%: 0.008+0.23+0.16 ms clock, 0.12+0/0.39/0.20+2.6 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 1009 @28.617s 1%: 0.013+44+5.7 ms clock, 0.20+0/0.23/29+92 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
HeapAlloc: 8249704 HeapInuse: 8626176 HeapObjects: 719 HeapIdle 8871936 HeapReleased 0 HeapSys 17498112
gc 1010 @29.069s 1%: 0.012+0.37+0.17 ms clock, 0.19+0/0.42/0.18+2.8 ms cpu, 7->7->0 MB, 15 MB goal, 16 P (forced)
 1000000	       454 ns/op	       8 B/op	       0 allocs/op
gc 1011 @29.070s 1%: 0.048+0.25+0.13 ms clock, 0.77+0/0.24/0.27+2.1 ms cpu, 0->0->0 MB, 7 MB goal, 16 P (forced)
gc 1012 @29.191s 1%: 0.021+3.6+0.15 ms clock, 0.34+0.25/12/15+2.4 ms cpu, 4->4->4 MB, 7 MB goal, 16 P
gc 1013 @29.264s 1%: 0.015+4.9+0.14 ms clock, 0.24+0.23/16/38+2.3 ms cpu, 5->6->6 MB, 8 MB goal, 16 P
gc 1014 @29.372s 1%: 0.019+6.1+0.10 ms clock, 0.30+0.39/23/53+1.6 ms cpu, 10->10->10 MB, 12 MB goal, 16 P
gc 1015 @29.585s 1%: 0.027+13+0.11 ms clock, 0.43+0.18/53/124+1.8 ms cpu, 19->20->19 MB, 21 MB goal, 16 P
gc 1016 @30.041s 1%: 0.020+27+0.17 ms clock, 0.32+0.23/102/218+2.7 ms cpu, 37->38->38 MB, 39 MB goal, 16 P
gc 1017 @30.899s 1%: 0.017+46+0.16 ms clock, 0.28+0.43/175/427+2.6 ms cpu, 74->75->75 MB, 76 MB goal, 16 P
HeapAlloc: 88246584 HeapInuse: 88915968 HeapObjects: 1500731 HeapIdle 933888 HeapReleased 0 HeapSys 89849856
gc 1018 @31.124s 1%: 0.023+0.42+0.17 ms clock, 0.37+0/0.58/0.71+2.8 ms cpu, 84->84->0 MB, 150 MB goal, 16 P (forced)
#BenchmarkSetInsert1E6-16                          	 1000000	      2070 ns/op	      88 B/op	       2 allocs/op
gc 1019 @31.141s 1%: 0.023+0.27+0.13 ms clock, 0.38+0/0.47/0.20+2.1 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1020 @31.144s 1%: 0.013+0.29+0.14 ms clock, 0.21+0.20/0.38/0.10+2.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1021 @31.172s 1%: 0.014+1.3+0.15 ms clock, 0.23+0.17/4.0/4.7+2.5 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1022 @31.302s 1%: 0.015+3.7+0.22 ms clock, 0.24+0.46/13/24+3.5 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1023 @31.520s 1%: 0.019+11+0.13 ms clock, 0.31+0.50/41/62+2.1 ms cpu, 21->21->21 MB, 25 MB goal, 16 P
gc 1024 @31.957s 1%: 0.016+31+0.15 ms clock, 0.27+0.39/89/169+2.4 ms cpu, 40->40->40 MB, 43 MB goal, 16 P
gc 1025 @32.834s 1%: 0.020+38+0.14 ms clock, 0.32+0.40/153/398+2.2 ms cpu, 77->78->78 MB, 81 MB goal, 16 P
gc 1026 @34.409s 1%: 0.018+15+0.11 ms clock, 0.29+1.0/58/115+1.7 ms cpu, 152->152->32 MB, 156 MB goal, 16 P
HeapAlloc: 58208152 HeapInuse: 58761216 HeapObjects: 852286 HeapIdle 102391808 HeapReleased 0 HeapSys 161153024
gc 1027 @34.854s 1%: 0.050+0.31+0.11 ms clock, 0.81+0/0.47/0.72+1.8 ms cpu, 55->55->0 MB, 64 MB goal, 16 P (forced)
#BenchmarkSetErase1E6-16                           	 1000000	      1584 ns/op	      96 B/op	       2 allocs/op
gc 1028 @34.859s 1%: 0.010+0.18+0.082 ms clock, 0.16+0/0.33/0.15+1.3 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1029 @34.861s 1%: 0.017+0.23+0.11 ms clock, 0.28+0.065/0.27/0.30+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1030 @34.890s 1%: 0.015+1.1+0.12 ms clock, 0.25+0.33/3.5/4.7+2.0 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1031 @34.970s 1%: 0.015+3.1+0.12 ms clock, 0.25+0.29/11/17+2.0 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1032 @35.192s 1%: 0.018+2.7+0.10 ms clock, 0.29+0.44/9.5/12+1.7 ms cpu, 21->21->10 MB, 25 MB goal, 16 P
gc 1033 @35.362s 1%: 0.017+4.1+0.11 ms clock, 0.28+0.17/14/20+1.8 ms cpu, 20->20->14 MB, 21 MB goal, 16 P
gc 1034 @35.553s 1%: 0.017+2.0+0.13 ms clock, 0.27+0.15/6.5/10+2.2 ms cpu, 27->27->10 MB, 28 MB goal, 16 P
gc 1035 @35.835s 1%: 0.018+4.3+0.12 ms clock, 0.29+0.23/13/25+2.0 ms cpu, 20->20->12 MB, 21 MB goal, 16 P
gc 1036 @36.190s 1%: 0.019+5.2+0.11 ms clock, 0.30+0.51/18/27+1.8 ms cpu, 24->24->15 MB, 25 MB goal, 16 P
gc 1037 @36.588s 1%: 0.018+3.4+0.12 ms clock, 0.30+0.26/12/26+1.9 ms cpu, 29->30->13 MB, 30 MB goal, 16 P
HeapAlloc: 25850056 HeapInuse: 26271744 HeapObjects: 300725 HeapIdle 134881280 HeapReleased 0 HeapSys 161153024
gc 1038 @36.870s 1%: 0.010+0.30+0.13 ms clock, 0.17+0/0.38/0.30+2.1 ms cpu, 24->24->0 MB, 26 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E6-16                  	 1000000	      2012 ns/op	      96 B/op	       2 allocs/op
gc 1039 @36.873s 1%: 0.011+0.23+0.15 ms clock, 0.19+0/0.34/0.29+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1040 @36.875s 1%: 0.013+0.25+0.11 ms clock, 0.22+0.074/0.35/0.19+1.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1041 @36.900s 1%: 0.015+1.1+0.10 ms clock, 0.24+0.17/3.1/4.3+1.6 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1042 @37.067s 1%: 0.017+4.3+0.20 ms clock, 0.27+0.33/14/1.2+3.2 ms cpu, 12->12->12 MB, 17 MB goal, 16 P
gc 1043 @37.585s 1%: 0.054+4.0+0.13 ms clock, 0.86+0.18/13/15+2.0 ms cpu, 21->21->12 MB, 25 MB goal, 16 P
gc 1044 @38.134s 1%: 0.016+3.6+0.11 ms clock, 0.26+0.34/12/21+1.8 ms cpu, 23->23->13 MB, 24 MB goal, 16 P
gc 1045 @38.640s 1%: 0.016+0.76+0.096 ms clock, 0.26+0.28/1.9/1.6+1.5 ms cpu, 26->26->8 MB, 27 MB goal, 16 P
gc 1046 @38.884s 1%: 0.018+4.2+0.12 ms clock, 0.29+0.12/13/24+2.0 ms cpu, 16->16->14 MB, 17 MB goal, 16 P
HeapAlloc: 15367360 HeapInuse: 15810560 HeapObjects: 148960 HeapIdle 145342464 HeapReleased 0 HeapSys 161153024
gc 1047 @38.966s 1%: 0.009+0.24+0.13 ms clock, 0.15+0/0.34/0.18+2.1 ms cpu, 14->14->0 MB, 28 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E6-16          	 1000000	      2093 ns/op	      56 B/op	       1 allocs/op
gc 1048 @38.967s 1%: 0.009+0.22+0.093 ms clock, 0.14+0/0.25/0.17+1.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1049 @38.986s 1%: 0.011+0.29+0.13 ms clock, 0.18+0.10/0.21/0.22+2.1 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 1050 @39.009s 1%: 0.015+0.27+0.16 ms clock, 0.24+0.21/0.20/0.28+2.6 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1051 @39.052s 1%: 0.015+2.5+0.12 ms clock, 0.24+0/0.65/2.3+1.9 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1052 @39.064s 1%: 0.016+0.57+0.13 ms clock, 0.26+0/0.58/0.23+2.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1053 @39.157s 1%: 0.021+6.7+0.23 ms clock, 0.33+0.53/0.53/5.8+3.7 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1054 @39.205s 1%: 0.021+1.4+0.13 ms clock, 0.34+0.75/0.94/0.28+2.1 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39255616 HeapInuse: 39755776 HeapObjects: 19917 HeapIdle 121397248 HeapReleased 0 HeapSys 161153024
gc 1055 @39.366s 1%: 0.013+0.42+0.12 ms clock, 0.21+0/0.50/0.46+2.0 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E6-16                   	gc 1056 @39.367s 1%: 0.009+0.18+0.12 ms clock, 0.14+0/0.36/0.15+2.0 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1057 @39.398s 1%: 0.014+0.34+0.11 ms clock, 0.23+0.23/0.27/0.32+1.8 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1058 @39.429s 1%: 0.018+0.44+0.18 ms clock, 0.29+0.40/0.30/0.25+2.9 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1059 @39.488s 1%: 0.018+2.4+0.13 ms clock, 0.30+0/1.0/1.6+2.2 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1060 @39.501s 1%: 0.016+0.57+0.097 ms clock, 0.26+0.47/0.38/0.51+1.5 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1061 @39.596s 1%: 0.015+3.2+0.11 ms clock, 0.24+0.38/0.60/2.5+1.8 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1062 @39.643s 1%: 0.020+1.4+0.20 ms clock, 0.33+0.78/0.70/0.52+3.3 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39242688 HeapInuse: 39747584 HeapObjects: 19783 HeapIdle 121405440 HeapReleased 0 HeapSys 161153024
gc 1063 @39.745s 1%: 0.009+0.34+0.15 ms clock, 0.15+0/0.41/0.11+2.5 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1064 @39.746s 1%: 0.027+0.23+0.10 ms clock, 0.44+0/0.19/0.18+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1065 @39.777s 1%: 0.014+0.34+0.13 ms clock, 0.23+0.21/0.33/0.29+2.1 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1066 @39.809s 1%: 0.016+0.44+0.20 ms clock, 0.27+0.41/0.26/0.27+3.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1067 @39.866s 1%: 0.019+2.4+0.21 ms clock, 0.30+0/1.1/1.8+3.3 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1068 @39.879s 1%: 0.018+0.71+0.12 ms clock, 0.30+0.66/0.23/0.21+2.0 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1069 @39.998s 1%: 0.019+5.8+0.17 ms clock, 0.31+0/1.0/5.2+2.7 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1070 @40.046s 1%: 0.023+1.6+0.14 ms clock, 0.37+0/1.5/0.34+2.2 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39232896 HeapInuse: 39731200 HeapObjects: 19681 HeapIdle 121421824 HeapReleased 0 HeapSys 161153024
gc 1071 @40.201s 1%: 0.012+0.51+0.15 ms clock, 0.19+0/0.35/0.56+2.4 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1072 @40.202s 1%: 0.029+0.15+0.11 ms clock, 0.46+0/0.24/0.19+1.8 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1073 @40.232s 1%: 0.030+0.46+0.16 ms clock, 0.48+0.17/0.31/0.28+2.7 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1074 @40.266s 1%: 0.019+0.38+0.13 ms clock, 0.31+0.34/0.26/0.35+2.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1075 @40.327s 1%: 0.022+16+0.14 ms clock, 0.35+0.81/0.16/3.6+2.3 ms cpu, 10->11->7 MB, 18 MB goal, 16 P
gc 1076 @40.346s 1%: 0.017+0.71+0.12 ms clock, 0.28+0.66/0.25/0.23+1.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1077 @40.461s 1%: 0.020+5.4+0.15 ms clock, 0.33+0/0.97/4.9+2.4 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1078 @40.511s 1%: 0.023+1.4+0.098 ms clock, 0.37+0/0.86/1.0+1.5 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39253344 HeapInuse: 39755776 HeapObjects: 19894 HeapIdle 121397248 HeapReleased 0 HeapSys 161153024
gc 1079 @40.664s 1%: 0.012+0.34+0.20 ms clock, 0.20+0/0.37/0.49+3.2 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1080 @40.665s 1%: 0.018+0.21+0.12 ms clock, 0.29+0/0.38/0.33+2.0 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1081 @40.696s 1%: 0.014+0.26+0.096 ms clock, 0.22+0.16/0.27/0.23+1.5 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1082 @40.728s 1%: 0.018+0.48+0.19 ms clock, 0.30+0.32/0.12/0.24+3.1 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1083 @40.785s 1%: 0.019+4.2+0.16 ms clock, 0.30+0.42/0.39/3.5+2.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1084 @40.801s 1%: 0.014+0.57+0.094 ms clock, 0.23+0.52/0.23/0.36+1.5 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1085 @40.883s 1%: 0.015+3.2+0.19 ms clock, 0.25+0/1.3/2.1+3.1 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1086 @40.916s 1%: 0.014+0.91+0.16 ms clock, 0.23+0.52/0.48/0.39+2.5 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39266784 HeapInuse: 39763968 HeapObjects: 20034 HeapIdle 121389056 HeapReleased 0 HeapSys 161153024
gc 1087 @41.015s 1%: 0.075+0.46+0.11 ms clock, 1.2+0/0.35/0.38+1.8 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1088 @41.016s 1%: 0.032+0.21+0.10 ms clock, 0.51+0/0.29/0.17+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1089 @41.046s 1%: 0.017+0.30+0.18 ms clock, 0.27+0.17/0.19/0.21+2.9 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1090 @41.069s 1%: 0.014+0.40+0.13 ms clock, 0.23+0.35/0.25/0.29+2.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1091 @41.106s 1%: 0.014+2.8+0.18 ms clock, 0.23+0/1.0/2.2+3.0 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1092 @41.119s 1%: 0.018+0.75+0.12 ms clock, 0.29+0.67/0.33/0.53+1.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1093 @41.238s 1%: 0.018+4.2+0.12 ms clock, 0.30+0/1.3/2.9+1.9 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1094 @41.280s 1%: 0.017+1.5+0.14 ms clock, 0.28+0.90/0.88/0.40+2.3 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39255648 HeapInuse: 39755776 HeapObjects: 19918 HeapIdle 121397248 HeapReleased 0 HeapSys 161153024
gc 1095 @41.443s 1%: 0.034+0.34+0.15 ms clock, 0.55+0/0.63/0.36+2.5 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1096 @41.443s 1%: 0.014+0.23+0.11 ms clock, 0.22+0/0.42/0.22+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1097 @41.474s 1%: 0.014+0.33+0.14 ms clock, 0.23+0.25/0.41/0.42+2.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1098 @41.507s 1%: 0.018+0.45+0.092 ms clock, 0.29+0.38/0.36/0.47+1.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1099 @41.567s 1%: 0.018+2.4+0.21 ms clock, 0.29+0/0.90/1.9+3.4 ms cpu, 10->10->6 MB, 19 MB goal, 16 P
gc 1100 @41.582s 1%: 0.019+0.75+0.13 ms clock, 0.31+0.67/0.19/0.31+2.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1101 @41.674s 1%: 0.018+2.1+0.14 ms clock, 0.28+0.33/0.31/1.6+2.2 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1102 @41.718s 1%: 0.017+1.5+0.12 ms clock, 0.28+0.92/0.80/0.27+1.9 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39251232 HeapInuse: 39755776 HeapObjects: 19872 HeapIdle 121397248 HeapReleased 0 HeapSys 161153024
gc 1103 @41.864s 1%: 0.010+0.48+0.16 ms clock, 0.17+0/0.33/0.39+2.5 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1104 @41.865s 1%: 0.010+0.21+0.12 ms clock, 0.16+0/0.27/0.20+1.9 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1105 @41.896s 1%: 0.013+0.27+0.18 ms clock, 0.22+0.18/0.17/0.34+2.9 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1106 @41.928s 1%: 0.017+1.0+0.19 ms clock, 0.28+0.96/0.24/0.33+3.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1107 @41.987s 1%: 0.017+15+0.14 ms clock, 0.28+0.81/0/3.2+2.2 ms cpu, 10->11->7 MB, 19 MB goal, 16 P
gc 1108 @42.006s 1%: 0.017+0.85+0.11 ms clock, 0.28+0.72/0.32/0.33+1.8 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1109 @42.122s 1%: 0.018+5.0+0.16 ms clock, 0.29+0/0.94/4.4+2.7 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1110 @42.166s 1%: 0.020+1.3+0.11 ms clock, 0.32+0.78/0.79/0.33+1.9 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39257952 HeapInuse: 39755776 HeapObjects: 19942 HeapIdle 121397248 HeapReleased 0 HeapSys 161153024
gc 1111 @42.284s 1%: 0.010+0.28+0.22 ms clock, 0.17+0/0.43/0.36+3.6 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1112 @42.285s 1%: 0.034+0.19+0.10 ms clock, 0.55+0/0.29/0.15+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1113 @42.315s 1%: 0.016+0.33+0.13 ms clock, 0.25+0.16/0.27/0.45+2.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1114 @42.348s 1%: 0.018+0.50+0.15 ms clock, 0.28+0.47/0.35/0.41+2.4 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1115 @42.405s 1%: 0.016+3.6+0.12 ms clock, 0.26+0/0.79/2.9+2.0 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1116 @42.419s 1%: 0.016+0.74+0.12 ms clock, 0.27+0.71/0.087/0.28+1.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1117 @42.536s 1%: 0.018+4.1+0.13 ms clock, 0.29+0.54/0.55/2.4+2.1 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1118 @42.574s 1%: 0.017+0.90+0.15 ms clock, 0.27+0.49/0.61/0.32+2.4 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39244896 HeapInuse: 39747584 HeapObjects: 19806 HeapIdle 121405440 HeapReleased 0 HeapSys 161153024
gc 1119 @42.717s 1%: 0.010+0.51+0.15 ms clock, 0.17+0/0.53/0.20+2.5 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1120 @42.718s 1%: 0.007+0.20+0.10 ms clock, 0.11+0/0.24/0.22+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1121 @42.749s 1%: 0.016+0.35+0.15 ms clock, 0.26+0.16/0.28/0.33+2.4 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1122 @42.782s 1%: 0.015+0.58+0.19 ms clock, 0.25+0.54/0.21/0.37+3.0 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1123 @42.839s 1%: 0.017+12+0.19 ms clock, 0.27+0.44/0.57/2.9+3.0 ms cpu, 10->10->7 MB, 18 MB goal, 16 P
gc 1124 @42.856s 1%: 0.016+0.77+0.35 ms clock, 0.25+0.72/0.30/0.31+5.6 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1125 @42.972s 1%: 0.018+5.0+0.13 ms clock, 0.29+0/0.79/4.6+2.1 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1126 @43.019s 1%: 0.018+1.5+0.14 ms clock, 0.29+0/1.5/0.30+2.2 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39243552 HeapInuse: 39747584 HeapObjects: 19792 HeapIdle 121405440 HeapReleased 0 HeapSys 161153024
gc 1127 @43.137s 1%: 0.009+0.37+0.17 ms clock, 0.15+0/0.47/0.14+2.8 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1128 @43.138s 1%: 0.056+0.20+0.10 ms clock, 0.91+0/0.27/0.15+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1129 @43.168s 1%: 0.018+0.29+0.14 ms clock, 0.30+0.23/0.19/0.28+2.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1130 @43.200s 1%: 0.018+0.37+0.14 ms clock, 0.29+0.35/0.28/0.40+2.2 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1131 @43.260s 1%: 0.018+2.5+0.15 ms clock, 0.29+0/1.1/1.8+2.4 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1132 @43.272s 1%: 0.016+0.60+0.18 ms clock, 0.26+0/0.62/0.34+2.9 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1133 @43.386s 1%: 0.021+4.8+0.16 ms clock, 0.34+0/1.2/3.7+2.6 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1134 @43.425s 1%: 0.020+1.1+0.10 ms clock, 0.32+0.68/0.70/0.27+1.7 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39247200 HeapInuse: 39747584 HeapObjects: 19830 HeapIdle 121405440 HeapReleased 0 HeapSys 161153024
gc 1135 @43.570s 1%: 0.010+0.33+0.14 ms clock, 0.17+0/0.44/0.28+2.3 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1136 @43.571s 1%: 0.017+0.16+0.10 ms clock, 0.28+0/0.25/0.18+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1137 @43.601s 1%: 0.014+0.30+0.12 ms clock, 0.22+0/0.41/0.21+2.0 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1138 @43.633s 1%: 0.021+0.43+0.18 ms clock, 0.33+0.35/0.27/0.40+2.9 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1139 @43.695s 1%: 0.020+3.7+0.13 ms clock, 0.32+0/1.2/3.0+2.2 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1140 @43.712s 1%: 0.018+0.88+0.13 ms clock, 0.28+0.76/0.26/0.32+2.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1141 @43.827s 1%: 0.017+3.2+0.23 ms clock, 0.27+0/1.4/2.2+3.7 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1142 @43.869s 1%: 0.020+1.5+0.18 ms clock, 0.32+0.75/1.1/0.59+2.8 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39265152 HeapInuse: 39763968 HeapObjects: 20017 HeapIdle 121389056 HeapReleased 0 HeapSys 161153024
gc 1143 @44.030s 1%: 0.013+0.46+0.093 ms clock, 0.21+0/0.57/0.14+1.5 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1144 @44.031s 1%: 0.036+0.26+0.10 ms clock, 0.58+0/0.33/0.19+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1145 @44.062s 1%: 0.021+0.38+0.16 ms clock, 0.34+0.26/0.33/0.32+2.6 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1146 @44.096s 1%: 0.022+0.56+0.13 ms clock, 0.35+0/0.77/0.33+2.1 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1147 @44.142s 1%: 0.020+0.78+0.19 ms clock, 0.32+0.11/0.62/0.12+3.1 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1148 @44.151s 1%: 0.023+0.86+0.13 ms clock, 0.38+0.74/0.26/0.28+2.1 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1149 @44.235s 1%: 0.020+3.7+0.19 ms clock, 0.32+0/1.3/2.5+3.1 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1150 @44.280s 1%: 0.020+2.7+0.18 ms clock, 0.33+0/2.7/0.33+2.9 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39238656 HeapInuse: 39739392 HeapObjects: 19741 HeapIdle 121413632 HeapReleased 0 HeapSys 161153024
gc 1151 @44.427s 1%: 0.011+0.33+0.28 ms clock, 0.18+0/0.32/0.39+4.5 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1152 @44.429s 1%: 0.009+0.26+0.17 ms clock, 0.15+0/0.34/0.23+2.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1153 @44.459s 1%: 0.016+0.29+0.16 ms clock, 0.26+0.16/0.21/0.38+2.6 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1154 @44.491s 1%: 0.014+0.41+0.14 ms clock, 0.23+0.38/0.26/0.37+2.3 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1155 @44.552s 1%: 0.018+12+0.17 ms clock, 0.28+0.41/0.39/3.1+2.8 ms cpu, 10->10->7 MB, 19 MB goal, 16 P
gc 1156 @44.569s 1%: 0.017+0.71+0.088 ms clock, 0.27+0.66/0.31/0.35+1.4 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1157 @44.684s 0%: 0.017+4.6+0.15 ms clock, 0.27+0/0.82/4.2+2.5 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1158 @44.724s 0%: 0.022+0.99+0.16 ms clock, 0.35+0.57/0.58/0.34+2.5 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39280032 HeapInuse: 39780352 HeapObjects: 20172 HeapIdle 121372672 HeapReleased 0 HeapSys 161153024
gc 1159 @44.868s 0%: 0.026+0.30+0.17 ms clock, 0.41+0/0.45/0.31+2.8 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
gc 1160 @44.869s 0%: 0.017+0.20+0.11 ms clock, 0.28+0/0.39/0.10+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1161 @44.899s 0%: 0.014+0.49+0.13 ms clock, 0.22+0.19/0.28/0.27+2.2 ms cpu, 6->6->4 MB, 8 MB goal, 16 P
gc 1162 @44.931s 0%: 0.015+0.58+0.17 ms clock, 0.24+0.56/0.36/0.32+2.7 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1163 @44.992s 0%: 0.017+2.6+0.11 ms clock, 0.28+0.41/0.66/1.7+1.8 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1164 @45.005s 0%: 0.020+0.71+0.11 ms clock, 0.33+0.68/0.34/0.38+1.7 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1165 @45.122s 0%: 0.019+4.5+0.15 ms clock, 0.30+0/0.91/3.9+2.4 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1166 @45.166s 0%: 0.023+1.5+0.15 ms clock, 0.37+0.90/0.72/0.37+2.5 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
HeapAlloc: 39245952 HeapInuse: 39747584 HeapObjects: 19817 HeapIdle 121405440 HeapReleased 0 HeapSys 161153024
gc 1167 @45.312s 0%: 0.011+0.36+0.10 ms clock, 0.18+0/0.57/0.30+1.6 ms cpu, 37->37->23 MB, 74 MB goal, 16 P (forced)
 1000000	       443 ns/op	      54 B/op	       0 allocs/op
gc 1168 @45.313s 0%: 0.061+0.21+0.10 ms clock, 0.97+0/0.30/0.15+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1169 @45.316s 0%: 0.013+0.24+0.11 ms clock, 0.21+0/0.30/0.18+1.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1170 @45.324s 0%: 0.013+0.27+0.16 ms clock, 0.21+0.11/0.21/0.092+2.6 ms cpu, 9->9->9 MB, 15 MB goal, 16 P
gc 1171 @45.347s 0%: 0.012+0.45+0.14 ms clock, 0.20+0.17/0.20/0.27+2.3 ms cpu, 14->14->12 MB, 18 MB goal, 16 P
gc 1172 @45.413s 0%: 0.016+0.48+0.17 ms clock, 0.26+0.46/0.23/0.30+2.7 ms cpu, 20->20->14 MB, 25 MB goal, 16 P
gc 1173 @45.422s 0%: 0.017+0.75+0.13 ms clock, 0.27+0.72/0.35/0.23+2.1 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1174 @45.573s 0%: 0.019+1.4+0.20 ms clock, 0.30+0.76/0.94/0.29+3.2 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47266016 HeapInuse: 47808512 HeapObjects: 19989 HeapIdle 113344512 HeapReleased 0 HeapSys 161153024
gc 1175 @45.881s 0%: 0.013+0.50+0.29 ms clock, 0.22+0/0.50/0.16+4.6 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E6-16                    	gc 1176 @45.882s 0%: 0.062+0.23+0.11 ms clock, 0.99+0/0.37/0.14+1.8 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1177 @45.884s 0%: 0.017+0.25+0.092 ms clock, 0.27+0/0.45/0.18+1.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1178 @45.892s 0%: 0.012+0.23+0.11 ms clock, 0.20+0.14/0.32/0.12+1.8 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1179 @45.915s 0%: 0.014+0.30+0.11 ms clock, 0.23+0.23/0.33/0.32+1.9 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1180 @46.005s 0%: 0.018+2.8+0.10 ms clock, 0.28+0/1.2/2.0+1.7 ms cpu, 19->19->14 MB, 25 MB goal, 16 P
gc 1181 @46.022s 0%: 0.020+0.85+0.11 ms clock, 0.33+0.79/0.24/0.51+1.8 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1182 @46.182s 0%: 0.020+1.5+0.12 ms clock, 0.32+0.94/0.96/0.29+1.9 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47252800 HeapInuse: 47792128 HeapObjects: 19852 HeapIdle 113360896 HeapReleased 0 HeapSys 161153024
gc 1183 @46.495s 0%: 0.011+0.56+0.11 ms clock, 0.18+0/0.59/0.12+1.7 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1184 @46.496s 0%: 0.009+0.20+0.10 ms clock, 0.15+0/0.32/0.35+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1185 @46.498s 0%: 0.016+0.23+0.11 ms clock, 0.26+0/0.25/0.34+1.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1186 @46.506s 0%: 0.012+0.22+0.079 ms clock, 0.19+0.15/0.31/0.36+1.2 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1187 @46.521s 0%: 0.015+0.27+0.13 ms clock, 0.24+0.12/0.29/0.27+2.1 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1188 @46.584s 0%: 0.017+0.60+0.15 ms clock, 0.27+0.066/0.25/0.68+2.5 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1189 @46.597s 0%: 0.017+0.85+0.14 ms clock, 0.27+0/0.31/0.95+2.3 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1190 @46.732s 0%: 0.027+1.2+0.11 ms clock, 0.44+0.71/0.66/0.37+1.8 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47251360 HeapInuse: 47783936 HeapObjects: 19837 HeapIdle 113369088 HeapReleased 0 HeapSys 161153024
gc 1191 @47.038s 0%: 0.014+0.30+0.12 ms clock, 0.23+0/0.24/0.45+2.0 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1192 @47.039s 0%: 0.020+0.18+0.11 ms clock, 0.32+0/0.33/0.18+1.8 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1193 @47.041s 0%: 0.015+0.20+0.16 ms clock, 0.24+0/0.24/0.18+2.6 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1194 @47.048s 0%: 0.011+0.24+0.14 ms clock, 0.19+0.13/0.21/0.32+2.2 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1195 @47.072s 0%: 0.016+0.35+0.13 ms clock, 0.27+0.30/0.28/0.37+2.1 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1196 @47.163s 0%: 0.018+2.3+0.19 ms clock, 0.29+0.41/0.69/1.5+3.1 ms cpu, 19->19->14 MB, 25 MB goal, 16 P
gc 1197 @47.176s 0%: 0.017+0.54+0.13 ms clock, 0.27+0.49/0.27/0.34+2.2 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1198 @47.306s 0%: 0.018+0.90+0.16 ms clock, 0.30+0.48/0.61/0.32+2.6 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47245408 HeapInuse: 47783936 HeapObjects: 19775 HeapIdle 113369088 HeapReleased 0 HeapSys 161153024
gc 1199 @47.616s 0%: 0.012+0.29+0.16 ms clock, 0.19+0/0.35/0.41+2.6 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1200 @47.617s 0%: 0.052+0.23+0.16 ms clock, 0.84+0/0.38/0.18+2.5 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1201 @47.619s 0%: 0.012+0.18+0.20 ms clock, 0.19+0/0.26/0.11+3.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1202 @47.627s 0%: 0.014+0.30+0.13 ms clock, 0.22+0.14/0.46/0.20+2.1 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1203 @47.651s 0%: 0.016+0.42+0.12 ms clock, 0.26+0.30/0.38/0.23+2.0 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1204 @47.743s 0%: 0.018+2.7+0.18 ms clock, 0.29+0.47/0.62/1.6+3.0 ms cpu, 19->19->14 MB, 25 MB goal, 16 P
gc 1205 @47.756s 0%: 0.016+0.51+0.15 ms clock, 0.27+0.49/0.29/0.24+2.4 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1206 @47.884s 0%: 0.020+0.93+0.14 ms clock, 0.32+0.55/0.63/0.27+2.3 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47252800 HeapInuse: 47783936 HeapObjects: 19852 HeapIdle 113369088 HeapReleased 0 HeapSys 161153024
gc 1207 @48.169s 0%: 0.013+0.25+0.16 ms clock, 0.21+0/0.48/0.19+2.6 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1208 @48.170s 0%: 0.056+0.20+0.20 ms clock, 0.89+0/0.29/0.23+3.3 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1209 @48.172s 0%: 0.015+0.28+0.17 ms clock, 0.25+0/0.38/0.20+2.8 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1210 @48.180s 0%: 0.012+0.33+0.35 ms clock, 0.19+0.14/0.14/0.26+5.6 ms cpu, 8->9->9 MB, 15 MB goal, 16 P
gc 1211 @48.195s 0%: 0.013+0.29+0.16 ms clock, 0.21+0.24/0.27/0.20+2.5 ms cpu, 14->14->12 MB, 18 MB goal, 16 P
gc 1212 @48.260s 0%: 0.017+2.3+0.16 ms clock, 0.28+0.48/0.33/1.8+2.7 ms cpu, 20->20->14 MB, 25 MB goal, 16 P
gc 1213 @48.268s 0%: 0.015+0.67+0.13 ms clock, 0.25+0/0.80/0.18+2.1 ms cpu, 26->26->26 MB, 29 MB goal, 16 P
gc 1214 @48.418s 0%: 0.017+1.5+0.20 ms clock, 0.28+0.89/0.85/0.33+3.2 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47261728 HeapInuse: 47792128 HeapObjects: 19945 HeapIdle 113360896 HeapReleased 0 HeapSys 161153024
gc 1215 @48.707s 0%: 0.012+0.36+0.13 ms clock, 0.19+0/0.34/0.24+2.0 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1216 @48.708s 0%: 0.007+0.15+0.090 ms clock, 0.11+0/0.24/0.11+1.4 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1217 @48.710s 0%: 0.014+0.20+0.091 ms clock, 0.23+0/0.29/0.17+1.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1218 @48.715s 0%: 0.010+0.39+0.093 ms clock, 0.16+0.089/0.27/0.23+1.5 ms cpu, 8->9->9 MB, 15 MB goal, 16 P
gc 1219 @48.730s 0%: 0.016+0.22+0.10 ms clock, 0.26+0.17/0.31/0.16+1.7 ms cpu, 14->14->12 MB, 18 MB goal, 16 P
gc 1220 @48.797s 0%: 0.019+2.2+0.15 ms clock, 0.31+0/1.0/1.7+2.5 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1221 @48.857s 0%: 0.018+1.1+0.11 ms clock, 0.29+0/1.1/0.37+1.9 ms cpu, 26->26->26 MB, 29 MB goal, 16 P
gc 1222 @48.945s 0%: 0.018+0.98+0.097 ms clock, 0.30+0.56/0.60/0.43+1.5 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47270944 HeapInuse: 47808512 HeapObjects: 20041 HeapIdle 113344512 HeapReleased 0 HeapSys 161153024
gc 1223 @49.234s 0%: 0.009+0.43+0.16 ms clock, 0.15+0/0.50/0.14+2.5 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1224 @49.235s 0%: 0.064+0.24+0.12 ms clock, 1.0+0/0.25/0.23+1.9 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1225 @49.237s 0%: 0.012+0.22+0.15 ms clock, 0.19+0/0.19/0.23+2.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1226 @49.245s 0%: 0.012+0.22+0.11 ms clock, 0.19+0.12/0.093/0.26+1.7 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1227 @49.269s 0%: 0.012+0.33+0.11 ms clock, 0.20+0.17/0.22/0.31+1.7 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1228 @49.331s 0%: 0.084+0.72+0.13 ms clock, 1.3+0.15/0.64/0.14+2.2 ms cpu, 19->20->14 MB, 25 MB goal, 16 P
gc 1229 @49.344s 0%: 0.020+0.78+0.15 ms clock, 0.32+0.71/0.25/0.25+2.4 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1230 @49.499s 0%: 0.018+1.5+0.090 ms clock, 0.29+0.97/1.0/0.49+1.4 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47265952 HeapInuse: 47800320 HeapObjects: 19989 HeapIdle 113352704 HeapReleased 0 HeapSys 161153024
gc 1231 @49.794s 0%: 0.010+0.56+0.12 ms clock, 0.17+0/0.54/0.096+2.0 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1232 @49.795s 0%: 0.013+0.32+0.11 ms clock, 0.21+0/0.39/0.24+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1233 @49.797s 0%: 0.013+0.23+0.11 ms clock, 0.22+0/0.32/0.15+1.7 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1234 @49.805s 0%: 0.012+0.33+0.25 ms clock, 0.19+0.15/0.30/0.20+4.0 ms cpu, 8->9->9 MB, 15 MB goal, 16 P
gc 1235 @49.829s 0%: 0.015+0.39+0.15 ms clock, 0.24+0.34/0.27/0.33+2.4 ms cpu, 14->14->12 MB, 18 MB goal, 16 P
gc 1236 @49.924s 0%: 0.016+2.2+0.15 ms clock, 0.26+0.16/0.12/0.75+2.4 ms cpu, 20->20->14 MB, 24 MB goal, 16 P
gc 1237 @49.931s 0%: 0.016+0.61+0.13 ms clock, 0.27+0.52/0.058/0.42+2.1 ms cpu, 26->26->26 MB, 29 MB goal, 16 P
gc 1238 @50.060s 0%: 0.024+0.93+0.12 ms clock, 0.38+0.50/0.68/0.32+1.9 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47273632 HeapInuse: 47808512 HeapObjects: 20069 HeapIdle 113344512 HeapReleased 0 HeapSys 161153024
gc 1239 @50.363s 0%: 0.040+0.39+0.11 ms clock, 0.64+0/0.11/0.45+1.8 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1240 @50.364s 0%: 0.035+0.18+0.15 ms clock, 0.56+0/0.20/0.10+2.5 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1241 @50.366s 0%: 0.013+0.25+0.12 ms clock, 0.22+0/0.35/0.095+1.9 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1242 @50.373s 0%: 0.013+0.20+0.15 ms clock, 0.21+0.072/0.15/0.19+2.5 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1243 @50.396s 0%: 0.014+0.28+0.15 ms clock, 0.23+0.16/0.27/0.23+2.4 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1244 @50.461s 0%: 0.017+0.62+0.15 ms clock, 0.27+0.12/0.63/0.18+2.4 ms cpu, 19->20->14 MB, 24 MB goal, 16 P
gc 1245 @50.474s 0%: 0.020+0.93+0.12 ms clock, 0.32+0/1.0/0.29+2.0 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1246 @50.633s 0%: 0.020+1.5+0.13 ms clock, 0.33+0/1.7/0.41+2.1 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47249824 HeapInuse: 47792128 HeapObjects: 19821 HeapIdle 113360896 HeapReleased 0 HeapSys 161153024
gc 1247 @50.940s 0%: 0.011+0.34+0.18 ms clock, 0.18+0/0.32/0.29+2.8 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1248 @50.941s 0%: 0.054+0.25+0.14 ms clock, 0.87+0/0.38/0.20+2.2 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1249 @50.943s 0%: 0.017+0.28+0.14 ms clock, 0.28+0/0.46/0.14+2.2 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1250 @50.949s 0%: 0.010+0.24+0.16 ms clock, 0.16+0.11/0.25/0.13+2.7 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1251 @50.972s 0%: 0.015+0.37+0.12 ms clock, 0.25+0.26/0.26/0.36+1.9 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1252 @51.047s 0%: 0.016+3.1+0.14 ms clock, 0.26+0/1.1/2.4+2.3 ms cpu, 19->19->14 MB, 25 MB goal, 16 P
gc 1253 @51.063s 0%: 0.018+0.80+0.17 ms clock, 0.28+0.75/0.37/0.31+2.8 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1254 @51.212s 0%: 0.017+4.4+0.21 ms clock, 0.28+0.75/0.92/0.17+3.4 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47269120 HeapInuse: 47808512 HeapObjects: 20022 HeapIdle 113344512 HeapReleased 0 HeapSys 161153024
gc 1255 @51.516s 0%: 0.010+0.49+0.084 ms clock, 0.17+0/0.53/0.14+1.3 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1256 @51.517s 0%: 0.007+0.21+0.086 ms clock, 0.12+0/0.36/0.26+1.3 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1257 @51.519s 0%: 0.012+0.18+0.073 ms clock, 0.19+0/0.34/0.33+1.1 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1258 @51.524s 0%: 0.011+0.26+0.073 ms clock, 0.18+0.079/0.26/0.17+1.1 ms cpu, 8->9->9 MB, 15 MB goal, 16 P
gc 1259 @51.548s 0%: 0.015+0.31+0.11 ms clock, 0.25+0.21/0.18/0.36+1.7 ms cpu, 14->14->12 MB, 18 MB goal, 16 P
gc 1260 @51.645s 0%: 0.017+8.6+0.12 ms clock, 0.28+1.0/0/4.7+2.0 ms cpu, 20->20->14 MB, 25 MB goal, 16 P
gc 1261 @51.722s 0%: 0.016+0.44+0.20 ms clock, 0.25+0/0.35/0.45+3.2 ms cpu, 26->26->19 MB, 29 MB goal, 16 P
gc 1262 @51.807s 0%: 0.018+1.5+0.12 ms clock, 0.29+0/1.5/0.38+2.0 ms cpu, 45->45->44 MB, 46 MB goal, 16 P
HeapAlloc: 47265280 HeapInuse: 47775744 HeapObjects: 19982 HeapIdle 113377280 HeapReleased 0 HeapSys 161153024
gc 1263 @52.113s 0%: 0.011+0.50+0.13 ms clock, 0.17+0/0.54/0.17+2.0 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
gc 1264 @52.114s 0%: 0.015+0.20+0.12 ms clock, 0.25+0/0.26/0.23+1.9 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1265 @52.116s 0%: 0.016+0.27+0.086 ms clock, 0.25+0/0.38/0.20+1.3 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1266 @52.124s 0%: 0.012+0.26+0.12 ms clock, 0.19+0.079/0.21/0.26+2.0 ms cpu, 8->8->8 MB, 15 MB goal, 16 P
gc 1267 @52.139s 0%: 0.016+0.26+0.12 ms clock, 0.26+0.10/0.15/0.32+1.9 ms cpu, 14->14->12 MB, 16 MB goal, 16 P
gc 1268 @52.234s 0%: 0.017+0.53+0.19 ms clock, 0.28+0.15/0.13/0.51+3.0 ms cpu, 20->20->14 MB, 25 MB goal, 16 P
gc 1269 @52.243s 0%: 0.018+0.73+0.12 ms clock, 0.28+0.71/0.17/0.28+2.0 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1270 @52.389s 0%: 0.018+1.3+0.12 ms clock, 0.30+0.74/0.94/0.41+1.9 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 47259520 HeapInuse: 47800320 HeapObjects: 19922 HeapIdle 113352704 HeapReleased 0 HeapSys 161153024
gc 1271 @52.698s 0%: 0.024+0.48+0.12 ms clock, 0.38+0/0.51/0.14+1.9 ms cpu, 45->45->23 MB, 89 MB goal, 16 P (forced)
 1000000	       166 ns/op	       8 B/op	       0 allocs/op
gc 1272 @52.699s 0%: 0.059+0.16+0.099 ms clock, 0.94+0/0.27/0.13+1.5 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1273 @52.715s 0%: 0.018+0.34+0.14 ms clock, 0.29+0.12/0.22/0.20+2.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4710272 HeapInuse: 5185536 HeapObjects: 10687 HeapIdle 155967488 HeapReleased 0 HeapSys 161153024
gc 1274 @52.860s 0%: 0.030+0.61+0.20 ms clock, 0.48+0/0.77/0.27+3.2 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E6-16           	gc 1275 @52.861s 0%: 0.006+0.18+0.13 ms clock, 0.10+0/0.23/0.21+2.2 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1276 @52.877s 0%: 0.013+0.28+0.17 ms clock, 0.21+0.12/0.22/0.12+2.7 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4694128 HeapInuse: 5177344 HeapObjects: 10570 HeapIdle 155975680 HeapReleased 0 HeapSys 161153024
gc 1277 @52.970s 0%: 0.009+0.62+0.13 ms clock, 0.15+0/0.62/0.13+2.1 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1278 @52.971s 0%: 0.056+0.19+0.10 ms clock, 0.89+0/0.33/0.18+1.7 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1279 @52.982s 0%: 0.014+0.35+0.12 ms clock, 0.22+0.19/0.021/0.31+1.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4703600 HeapInuse: 5185536 HeapObjects: 10618 HeapIdle 155967488 HeapReleased 0 HeapSys 161153024
gc 1280 @53.126s 0%: 0.017+0.60+0.11 ms clock, 0.27+0/0.75/0.23+1.8 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1281 @53.127s 0%: 0.008+0.19+0.13 ms clock, 0.12+0/0.23/0.21+2.1 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1282 @53.137s 0%: 0.017+0.27+0.22 ms clock, 0.28+0.11/0.30/0.20+3.5 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4690928 HeapInuse: 5169152 HeapObjects: 10486 HeapIdle 155983872 HeapReleased 0 HeapSys 161153024
gc 1283 @53.280s 0%: 0.010+0.84+0.14 ms clock, 0.16+0/0.83/0.083+2.3 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1284 @53.282s 0%: 0.060+0.27+0.11 ms clock, 0.96+0/0.28/0.22+1.7 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1285 @53.292s 0%: 0.011+0.33+0.17 ms clock, 0.19+0.16/0.29/0.24+2.8 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4700624 HeapInuse: 5177344 HeapObjects: 10587 HeapIdle 155975680 HeapReleased 0 HeapSys 161153024
gc 1286 @53.435s 0%: 0.030+0.64+0.12 ms clock, 0.48+0/0.94/0.17+2.0 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1287 @53.436s 0%: 0.010+0.21+0.16 ms clock, 0.16+0/0.29/0.12+2.6 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1288 @53.446s 0%: 0.014+0.26+0.14 ms clock, 0.23+0.10/0.23/0.20+2.3 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4710128 HeapInuse: 5185536 HeapObjects: 10686 HeapIdle 155967488 HeapReleased 0 HeapSys 161153024
gc 1289 @53.543s 0%: 0.016+0.48+0.12 ms clock, 0.25+0/0.70/0.27+1.9 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1290 @53.544s 0%: 0.015+0.16+0.12 ms clock, 0.24+0/0.21/0.14+2.0 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1291 @53.559s 0%: 0.012+1.4+0.12 ms clock, 0.20+0.32/0.032/0.47+1.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4694608 HeapInuse: 5185536 HeapObjects: 10575 HeapIdle 155967488 HeapReleased 0 HeapSys 161153024
gc 1292 @53.706s 0%: 0.010+0.47+0.10 ms clock, 0.17+0/0.71/0.28+1.6 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1293 @53.707s 0%: 0.009+0.19+0.12 ms clock, 0.15+0/0.27/0.15+1.9 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1294 @53.723s 0%: 0.013+0.28+0.13 ms clock, 0.21+0.11/0.31/0.20+2.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4690320 HeapInuse: 5177344 HeapObjects: 10516 HeapIdle 155942912 HeapReleased 0 HeapSys 161120256
gc 1295 @53.865s 0%: 0.031+0.56+0.13 ms clock, 0.49+0/0.30/0.46+2.1 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1296 @53.866s 0%: 0.027+0.20+0.12 ms clock, 0.44+0/0.26/0.17+2.0 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1297 @53.876s 0%: 0.016+0.27+0.13 ms clock, 0.26+0.12/0.19/0.26+2.1 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4695856 HeapInuse: 5177344 HeapObjects: 10523 HeapIdle 155942912 HeapReleased 0 HeapSys 161120256
gc 1298 @54.023s 0%: 0.010+0.61+0.10 ms clock, 0.17+0/0.79/0.28+1.7 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1299 @54.024s 0%: 0.017+0.27+0.077 ms clock, 0.27+0/0.18/0.20+1.2 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1300 @54.039s 0%: 0.012+0.27+0.090 ms clock, 0.19+0.10/0.24/0.20+1.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4685232 HeapInuse: 5177344 HeapObjects: 10463 HeapIdle 155942912 HeapReleased 0 HeapSys 161120256
gc 1301 @54.183s 0%: 0.041+0.47+0.14 ms clock, 0.66+0/0.40/0.41+2.3 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1302 @54.184s 0%: 0.11+0.31+0.13 ms clock, 1.8+0/0.45/0.072+2.1 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1303 @54.200s 0%: 0.012+0.31+0.092 ms clock, 0.20+0.14/0.40/0.33+1.4 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4695120 HeapInuse: 5185536 HeapObjects: 10566 HeapIdle 155934720 HeapReleased 0 HeapSys 161120256
gc 1304 @54.303s 0%: 0.013+0.53+0.11 ms clock, 0.21+0/0.29/0.46+1.8 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1305 @54.304s 0%: 0.039+0.20+0.12 ms clock, 0.63+0/0.34/0.24+1.9 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1306 @54.314s 0%: 0.012+0.27+0.18 ms clock, 0.20+0.093/0.17/0.19+2.9 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4705840 HeapInuse: 5193728 HeapObjects: 10627 HeapIdle 155926528 HeapReleased 0 HeapSys 161120256
gc 1307 @54.458s 0%: 0.011+0.84+0.15 ms clock, 0.17+0/0.84/0.20+2.5 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
gc 1308 @54.459s 0%: 0.006+0.19+0.11 ms clock, 0.10+0/0.25/0.17+1.8 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1309 @54.475s 0%: 0.013+0.26+0.12 ms clock, 0.20+0.087/0.11/0.26+2.0 ms cpu, 4->4->3 MB, 5 MB goal, 16 P
HeapAlloc: 4695280 HeapInuse: 5177344 HeapObjects: 10517 HeapIdle 155942912 HeapReleased 0 HeapSys 161120256
gc 1310 @54.571s 0%: 0.009+0.53+0.16 ms clock, 0.15+0/0.70/0.21+2.6 ms cpu, 4->4->2 MB, 6 MB goal, 16 P (forced)
 1000000	       112 ns/op	       5 B/op	       0 allocs/op
gc 1311 @54.572s 0%: 0.056+0.21+0.13 ms clock, 0.90+0/0.34/0.15+2.1 ms cpu, 2->2->0 MB, 5 MB goal, 16 P (forced)
gc 1312 @54.577s 0%: 0.018+0.26+0.15 ms clock, 0.28+0.097/0.27/0.23+2.4 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25562472 HeapInuse: 25935872 HeapObjects: 737 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1313 @54.752s 0%: 0.011+0.28+0.12 ms clock, 0.17+0/0.45/0.20+1.9 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E6-16    	gc 1314 @54.753s 0%: 0.041+0.24+0.077 ms clock, 0.66+0/0.34/0.21+1.2 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1315 @54.758s 0%: 0.018+1.7+0.13 ms clock, 0.30+0.20/0.31/0.10+2.1 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25562712 HeapInuse: 25935872 HeapObjects: 740 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1316 @54.947s 0%: 0.013+0.31+0.23 ms clock, 0.22+0/0.46/0.13+3.7 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1317 @54.948s 0%: 0.011+0.22+0.13 ms clock, 0.18+0/0.38/0.22+2.1 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1318 @54.953s 0%: 0.018+0.37+0.32 ms clock, 0.29+0.13/0.16/0.32+5.1 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25562904 HeapInuse: 25935872 HeapObjects: 742 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1319 @55.126s 0%: 0.011+0.27+0.16 ms clock, 0.18+0/0.42/0.19+2.5 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1320 @55.127s 0%: 0.008+0.21+0.14 ms clock, 0.12+0/0.37/0.27+2.3 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1321 @55.132s 0%: 0.021+0.25+0.17 ms clock, 0.34+0.11/0.25/0.17+2.7 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563000 HeapInuse: 25935872 HeapObjects: 743 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1322 @55.316s 0%: 0.013+0.29+0.18 ms clock, 0.22+0/0.42/0.28+2.9 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1323 @55.317s 0%: 0.074+0.20+0.13 ms clock, 1.1+0/0.42/0.17+2.0 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1324 @55.322s 0%: 0.022+0.32+0.16 ms clock, 0.35+0.11/0.27/0.20+2.6 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563000 HeapInuse: 25935872 HeapObjects: 743 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1325 @55.506s 0%: 0.019+0.29+0.14 ms clock, 0.30+0/0.51/0.28+2.3 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1326 @55.507s 0%: 0.012+0.30+0.093 ms clock, 0.20+0/0.41/0.24+1.5 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1327 @55.512s 0%: 0.019+0.28+0.097 ms clock, 0.31+0.086/0.33/0.28+1.5 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563000 HeapInuse: 25935872 HeapObjects: 743 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1328 @55.682s 0%: 0.012+0.24+0.16 ms clock, 0.20+0/0.25/0.23+2.6 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1329 @55.682s 0%: 0.010+0.23+0.12 ms clock, 0.17+0/0.36/0.19+1.9 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1330 @55.688s 0%: 0.020+0.31+0.13 ms clock, 0.32+0.16/0.21/0.20+2.2 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563000 HeapInuse: 25935872 HeapObjects: 743 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1331 @55.857s 0%: 0.013+0.34+0.17 ms clock, 0.21+0/0.37/0.29+2.8 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1332 @55.858s 0%: 0.013+0.24+0.098 ms clock, 0.21+0/0.50/0.36+1.5 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1333 @55.863s 0%: 0.018+0.30+0.21 ms clock, 0.29+0.090/0.33/0.25+3.4 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563096 HeapInuse: 25935872 HeapObjects: 744 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1334 @56.034s 0%: 0.012+0.25+0.16 ms clock, 0.20+0/0.41/0.30+2.6 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1335 @56.035s 0%: 0.012+0.25+0.11 ms clock, 0.19+0/0.32/0.37+1.7 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1336 @56.040s 0%: 0.017+0.27+0.13 ms clock, 0.28+0.14/0.26/0.22+2.1 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563192 HeapInuse: 25935872 HeapObjects: 745 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1337 @56.214s 0%: 0.012+0.24+0.13 ms clock, 0.19+0/0.32/0.20+2.1 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1338 @56.215s 0%: 0.10+0.23+0.14 ms clock, 1.7+0/0.42/0.12+2.2 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1339 @56.220s 0%: 0.025+0.27+0.18 ms clock, 0.40+0.15/0.25/0.14+2.9 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563288 HeapInuse: 25935872 HeapObjects: 746 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1340 @56.388s 0%: 0.014+0.29+0.15 ms clock, 0.23+0/0.52/0.22+2.5 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1341 @56.389s 0%: 0.073+0.24+0.14 ms clock, 1.1+0/0.36/0.22+2.2 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1342 @56.395s 0%: 0.019+0.25+0.17 ms clock, 0.31+0.15/0.23/0.13+2.8 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563288 HeapInuse: 25935872 HeapObjects: 746 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1343 @56.571s 0%: 0.019+0.27+0.19 ms clock, 0.30+0/0.31/0.27+3.1 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1344 @56.572s 0%: 0.10+0.27+0.14 ms clock, 1.6+0/0.34/0.15+2.2 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1345 @56.577s 0%: 0.016+0.26+0.16 ms clock, 0.26+0.072/0.37/0.21+2.5 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563288 HeapInuse: 25935872 HeapObjects: 746 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1346 @56.759s 0%: 0.014+0.28+0.099 ms clock, 0.22+0/0.50/0.17+1.5 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
gc 1347 @56.760s 0%: 0.012+0.24+0.087 ms clock, 0.19+0/0.37/0.19+1.4 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1348 @56.765s 0%: 0.016+0.27+0.10 ms clock, 0.27+0.15/0.26/0.31+1.6 ms cpu, 24->24->24 MB, 25 MB goal, 16 P
HeapAlloc: 25563384 HeapInuse: 25935872 HeapObjects: 747 HeapIdle 135184384 HeapReleased 0 HeapSys 161120256
gc 1349 @56.954s 0%: 0.012+0.28+0.10 ms clock, 0.19+0/0.48/0.16+1.7 ms cpu, 24->24->23 MB, 48 MB goal, 16 P (forced)
 1000000	       194 ns/op	      25 B/op	       0 allocs/op
gc 1350 @56.955s 0%: 0.044+0.24+0.10 ms clock, 0.71+0/0.30/0.27+1.6 ms cpu, 23->23->0 MB, 47 MB goal, 16 P (forced)
gc 1351 @56.973s 0%: 0.022+454+65 ms clock, 0.36+0/0.22/298+1040 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
HeapAlloc: 80253128 HeapInuse: 80617472 HeapObjects: 745 HeapIdle 80502784 HeapReleased 0 HeapSys 161120256
gc 1352 @61.187s 0%: 0.010+0.25+0.14 ms clock, 0.17+0/0.29/0.23+2.3 ms cpu, 76->76->0 MB, 153 MB goal, 16 P (forced)
#BenchmarkSort1E7-16                               	10000000	       423 ns/op	       8 B/op	       0 allocs/op
gc 1353 @61.187s 0%: 0.013+0.14+0.13 ms clock, 0.21+0/0.24/0.11+2.1 ms cpu, 0->0->0 MB, 8 MB goal, 16 P (forced)
gc 1354 @61.260s 0%: 0.016+2.1+0.10 ms clock, 0.25+0.12/7.8/12+1.7 ms cpu, 4->4->3 MB, 8 MB goal, 16 P
gc 1355 @61.328s 0%: 0.014+3.3+0.11 ms clock, 0.23+0.27/11/26+1.8 ms cpu, 5->6->5 MB, 7 MB goal, 16 P
gc 1356 @61.394s 0%: 0.015+7.8+0.12 ms clock, 0.25+0.16/26/33+1.9 ms cpu, 10->10->10 MB, 11 MB goal, 16 P
gc 1357 @61.524s 0%: 0.018+12+0.12 ms clock, 0.30+0.20/44/79+2.0 ms cpu, 19->19->19 MB, 21 MB goal, 16 P
gc 1358 @61.912s 0%: 0.022+27+0.21 ms clock, 0.36+0.24/105/218+3.3 ms cpu, 37->38->37 MB, 39 MB goal, 16 P
gc 1359 @62.783s 0%: 0.015+49+0.13 ms clock, 0.25+0.21/187/358+2.1 ms cpu, 73->74->74 MB, 75 MB goal, 16 P
gc 1360 @64.079s 0%: 0.027+84+0.21 ms clock, 0.44+0.46/334/895+3.4 ms cpu, 144->146->146 MB, 148 MB goal, 16 P
gc 1361 @66.951s 0%: 0.022+167+0.14 ms clock, 0.36+0.84/660/1728+2.3 ms cpu, 285->288->288 MB, 292 MB goal, 16 P
gc 1362 @74.186s 0%: 0.024+326+0.10 ms clock, 0.39+11/1295/3511+1.7 ms cpu, 561->568->567 MB, 576 MB goal, 16 P
HeapAlloc: 880250760 HeapInuse: 883744768 HeapObjects: 15000771 HeapIdle 892928 HeapReleased 0 HeapSys 884637696
gc 1363 @80.601s 0%: 0.030+1.4+0.10 ms clock, 0.48+0/1.5/11+1.6 ms cpu, 840->840->0 MB, 1134 MB goal, 16 P (forced)
#BenchmarkSetInsert1E7-16                          	10000000	      1957 ns/op	      88 B/op	       2 allocs/op
gc 1364 @80.764s 0%: 0.017+0.23+0.12 ms clock, 0.28+0/0.34/0.15+1.9 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1365 @80.783s 0%: 0.031+0.26+0.13 ms clock, 0.50+0.12/0.23/0.10+2.1 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1366 @80.811s 0%: 0.020+1.1+0.12 ms clock, 0.33+0.23/3.3/3.2+2.0 ms cpu, 77->77->77 MB, 153 MB goal, 16 P
gc 1367 @81.685s 0%: 0.018+28+0.13 ms clock, 0.30+0.25/93/168+2.1 ms cpu, 115->116->115 MB, 155 MB goal, 16 P
gc 1368 @83.498s 0%: 0.015+72+0.14 ms clock, 0.24+0.55/283/725+2.3 ms cpu, 202->204->203 MB, 231 MB goal, 16 P
gc 1369 @87.529s 0%: 0.017+181+0.12 ms clock, 0.28+0.82/723/1834+1.9 ms cpu, 379->384->383 MB, 407 MB goal, 16 P
gc 1370 @95.066s 0%: 0.022+376+0.10 ms clock, 0.35+1.0/1500/4065+1.7 ms cpu, 736->742->741 MB, 767 MB goal, 16 P
gc 1371 @109.376s 0%: 0.023+264+0.15 ms clock, 0.36+2.7/1033/2397+2.5 ms cpu, 1446->1450->391 MB, 1483 MB goal, 16 P
HeapAlloc: 733391416 HeapInuse: 736075776 HeapObjects: 11133787 HeapIdle 788193280 HeapReleased 0 HeapSys 1524269056
gc 1372 @115.450s 0%: 0.053+1.3+0.11 ms clock, 0.85+0/1.6/8.5+1.8 ms cpu, 700->700->0 MB, 783 MB goal, 16 P (forced)
#BenchmarkSetErase1E7-16                           	10000000	      1647 ns/op	      96 B/op	       2 allocs/op
gc 1373 @115.519s 0%: 0.031+0.35+0.15 ms clock, 0.51+0/0.44/0.44+2.4 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1374 @115.539s 0%: 0.037+0.29+0.090 ms clock, 0.59+0.089/0.36/0.32+1.4 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1375 @115.567s 0%: 0.017+1.1+0.14 ms clock, 0.28+0.14/3.8/3.5+2.2 ms cpu, 77->77->77 MB, 153 MB goal, 16 P
gc 1376 @116.361s 0%: 0.017+23+0.17 ms clock, 0.28+0.22/90/181+2.7 ms cpu, 115->116->116 MB, 155 MB goal, 16 P
gc 1377 @118.157s 0%: 0.014+20+0.11 ms clock, 0.23+0.35/81/220+1.8 ms cpu, 202->203->119 MB, 232 MB goal, 16 P
gc 1378 @119.812s 0%: 0.016+33+0.18 ms clock, 0.27+0.50/131/339+3.0 ms cpu, 222->223->138 MB, 238 MB goal, 16 P
GC forced
gc 4 @120.829s 0%: 0.10+0.59+0.20 ms clock, 1.1+0/1.7/2.8+2.3 ms cpu, 3->3->1 MB, 4 MB goal, 16 P
gc 1379 @122.323s 0%: 0.016+10+0.12 ms clock, 0.26+0.62/40/110+2.0 ms cpu, 268->268->100 MB, 277 MB goal, 16 P
gc 1380 @124.511s 0%: 0.018+29+0.10 ms clock, 0.30+0.32/114/284+1.7 ms cpu, 195->195->126 MB, 200 MB goal, 16 P
gc 1381 @126.613s 0%: 0.017+44+0.085 ms clock, 0.27+0.35/174/371+1.3 ms cpu, 246->247->148 MB, 252 MB goal, 16 P
gc 1382 @129.168s 0%: 0.021+23+0.13 ms clock, 0.34+0.35/88/231+2.1 ms cpu, 289->289->121 MB, 296 MB goal, 16 P
gc 1383 @131.435s 0%: 0.023+4.7+0.11 ms clock, 0.37+0.38/17/43+1.9 ms cpu, 236->236->84 MB, 242 MB goal, 16 P
HeapAlloc: 97682520 HeapInuse: 98115584 HeapObjects: 297837 HeapIdle 1426153472 HeapReleased 0 HeapSys 1524269056
gc 1384 @131.619s 0%: 0.043+0.31+0.12 ms clock, 0.69+0/0.44/0.23+2.0 ms cpu, 93->93->0 MB, 169 MB goal, 16 P (forced)
#BenchmarkSetInsertAndErase1E7-16                  	10000000	      1610 ns/op	      96 B/op	       2 allocs/op
gc 1385 @131.623s 0%: 0.026+0.22+0.17 ms clock, 0.42+0/0.38/0.16+2.8 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1386 @131.641s 0%: 0.024+0.29+0.15 ms clock, 0.39+0.17/0.32/0.11+2.4 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1387 @131.680s 0%: 0.012+1.3+0.15 ms clock, 0.20+0.17/3.9/4.8+2.4 ms cpu, 77->77->77 MB, 153 MB goal, 16 P
gc 1388 @132.852s 0%: 0.014+26+0.15 ms clock, 0.23+0.20/103/181+2.4 ms cpu, 115->116->115 MB, 155 MB goal, 16 P
gc 1389 @135.882s 0%: 0.016+18+0.12 ms clock, 0.26+0.22/70/164+2.0 ms cpu, 202->202->110 MB, 231 MB goal, 16 P
gc 1390 @139.578s 0%: 0.022+23+0.13 ms clock, 0.36+0.16/91/218+2.1 ms cpu, 207->207->115 MB, 221 MB goal, 16 P
gc 1391 @143.306s 0%: 0.015+32+0.11 ms clock, 0.25+0.39/126/298+1.8 ms cpu, 223->223->131 MB, 231 MB goal, 16 P
gc 1392 @148.909s 0%: 0.021+23+0.12 ms clock, 0.35+0.53/88/181+1.9 ms cpu, 257->257->108 MB, 263 MB goal, 16 P
HeapAlloc: 117052880 HeapInuse: 117612544 HeapObjects: 767365 HeapIdle 1406656512 HeapReleased 0 HeapSys 1524269056
gc 1393 @149.622s 0%: 0.015+0.27+0.13 ms clock, 0.25+0/0.49/0.32+2.1 ms cpu, 111->111->0 MB, 217 MB goal, 16 P (forced)
#BenchmarkSetInsertAndEraseWithPool1E7-16          	10000000	      1800 ns/op	      56 B/op	       1 allocs/op
gc 1394 @149.625s 0%: 0.039+0.22+0.16 ms clock, 0.62+0/0.32/0.26+2.5 ms cpu, 0->0->0 MB, 4 MB goal, 16 P (forced)
gc 1395 @149.644s 0%: 0.012+0.32+0.14 ms clock, 0.19+0.12/0.30/0.32+2.3 ms cpu, 6->6->4 MB, 7 MB goal, 16 P
gc 1396 @149.677s 0%: 0.019+0.47+0.17 ms clock, 0.31+0.45/0.29/0.26+2.7 ms cpu, 11->11->9 MB, 12 MB goal, 16 P
gc 1397 @149.734s 0%: 0.017+3.4+0.11 ms clock, 0.27+0/0.83/2.8+1.7 ms cpu, 10->10->6 MB, 18 MB goal, 16 P
gc 1398 @149.747s 0%: 0.015+0.59+0.15 ms clock, 0.24+0.52/0.26/0.25+2.5 ms cpu, 18->18->18 MB, 19 MB goal, 16 P
gc 1399 @149.821s 0%: 0.017+3.2+0.14 ms clock, 0.27+0/1.4/2.0+2.3 ms cpu, 19->19->12 MB, 37 MB goal, 16 P
gc 1400 @149.851s 0%: 0.022+0.84+0.16 ms clock, 0.35+0.47/0.52/0.25+2.5 ms cpu, 37->37->37 MB, 38 MB goal, 16 P
scvg0: inuse: 37, idle: 1415, sys: 1453, released: 0, consumed: 1453 (MB)
gc 1401 @150.063s 0%: 0.017+1.1+0.16 ms clock, 0.28+0.14/1.0/0.18+2.6 ms cpu, 38->38->23 MB, 74 MB goal, 16 P
gc 1402 @150.173s 0%: 0.024+2.6+0.15 ms clock, 0.38+0.69/2.2/0.32+2.4 ms cpu, 75->75->74 MB, 76 MB goal, 16 P
gc 1403 @150.566s 0%: 0.017+8.3+0.18 ms clock, 0.27+0/2.6/6.3+3.0 ms cpu, 75->75->47 MB, 148 MB goal, 16 P
gc 1404 @150.891s 0%: 0.022+5.2+0.15 ms clock, 0.35+0.67/4.7/0.33+2.5 ms cpu, 151->151->148 MB, 152 MB goal, 16 P
gc 1405 @151.578s 0%: 0.017+8.5+0.18 ms clock, 0.27+0/2.6/6.4+2.8 ms cpu, 149->149->94 MB, 296 MB goal, 16 P
gc 1406 @152.275s 0%: 0.016+6.5+0.20 ms clock, 0.25+0.49/6.2/0.32+3.2 ms cpu, 303->303->296 MB, 304 MB goal, 16 P
gc 1407 @153.279s 0%: 0.014+4.8+0.19 ms clock, 0.23+0/2.1/3.0+3.0 ms cpu, 297->297->187 MB, 593 MB goal, 16 P
HeapAlloc: 200323400 HeapInuse: 200687616 HeapObjects: 755 HeapIdle 1323581440 HeapReleased 0 HeapSys 1524269056
gc 1408 @153.967s 0%: 0.011+4.8+0.11 ms clock, 0.17+0/5.0/0.25+1.7 ms cpu, 191->191->188 MB, 375 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsert1E7-16                   	10000000	       434 ns/op	      44 B/op	       0 allocs/op
gc 1409 @153.972s 0%: 0.019+0.23+0.14 ms clock, 0.30+0/0.54/0.11+2.3 ms cpu, 188->188->0 MB, 376 MB goal, 16 P (forced)
gc 1410 @153.991s 0%: 0.021+0.28+0.19 ms clock, 0.35+0/0.42/0.23+3.0 ms cpu, 76->76->76 MB, 77 MB goal, 16 P
gc 1411 @153.998s 0%: 0.011+0.22+0.085 ms clock, 0.18+0.13/0.14/0.21+1.3 ms cpu, 77->77->76 MB, 153 MB goal, 16 P
gc 1412 @154.263s 0%: 0.021+1.3+0.23 ms clock, 0.34+0.75/0.85/0.34+3.8 ms cpu, 128->128->113 MB, 153 MB goal, 16 P
gc 1413 @155.297s 0%: 0.019+3.3+0.11 ms clock, 0.31+0.57/2.9/0.44+1.8 ms cpu, 271->271->224 MB, 272 MB goal, 16 P
gc 1414 @156.438s 0%: 0.019+6.3+0.18 ms clock, 0.31+0.43/6.0/0.41+2.9 ms cpu, 434->434->372 MB, 449 MB goal, 16 P
HeapAlloc: 395650432 HeapInuse: 396214272 HeapObjects: 154080 HeapIdle 1128054784 HeapReleased 0 HeapSys 1524269056
gc 1415 @159.947s 0%: 0.009+3.1+0.16 ms clock, 0.14+0/0.17/3.1+2.7 ms cpu, 377->377->188 MB, 745 MB goal, 16 P (forced)
#BenchmarkSysHashMapErase1E7-16                    	10000000	       193 ns/op	       8 B/op	       0 allocs/op
gc 1416 @159.953s 0%: 0.036+0.19+0.19 ms clock, 0.58+0/0.31/0.12+3.1 ms cpu, 188->188->0 MB, 376 MB goal, 16 P (forced)
gc 1417 @159.955s 0%: 0.012+0.23+0.21 ms clock, 0.20+0.058/0.24/0.20+3.4 ms cpu, 7->8->7 MB, 8 MB goal, 16 P
gc 1418 @159.960s 0%: 0.008+0.29+0.13 ms clock, 0.14+0.11/0.17/0.10+2.1 ms cpu, 9->9->9 MB, 15 MB goal, 16 P
gc 1419 @159.975s 0%: 0.010+0.29+0.17 ms clock, 0.16+0.14/0.24/0.18+2.8 ms cpu, 14->14->12 MB, 18 MB goal, 16 P
gc 1420 @160.032s 0%: 0.015+1.6+0.25 ms clock, 0.25+0/0.67/1.3+4.1 ms cpu, 19->20->14 MB, 25 MB goal, 16 P
gc 1421 @160.044s 0%: 0.014+0.48+0.16 ms clock, 0.23+0.45/0.21/0.33+2.5 ms cpu, 26->26->26 MB, 28 MB goal, 16 P
gc 1422 @160.154s 0%: 0.015+1.1+0.17 ms clock, 0.25+0.54/0.82/0.34+2.7 ms cpu, 52->52->44 MB, 53 MB goal, 16 P
HeapAlloc: 48389216 HeapInuse: 48898048 HeapObjects: 24342 HeapIdle 1475371008 HeapReleased 0 HeapSys 1524269056
gc 1423 @161.998s 0%: 0.009+0.86+0.14 ms clock, 0.15+0/0.43/0.67+2.3 ms cpu, 46->46->24 MB, 89 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndErase1E7-16           	10000000	       204 ns/op	       6 B/op	       0 allocs/op
gc 1424 @161.999s 0%: 0.008+0.21+0.095 ms clock, 0.14+0/0.29/0.10+1.5 ms cpu, 24->24->0 MB, 48 MB goal, 16 P (forced)
gc 1425 @162.001s 0%: 0.013+0.19+0.092 ms clock, 0.21+0.12/0.27/0.16+1.4 ms cpu, 7->7->7 MB, 8 MB goal, 16 P
gc 1426 @162.042s 0%: 0.016+0.29+0.098 ms clock, 0.26+0.10/0.20/0.24+1.5 ms cpu, 194->194->194 MB, 195 MB goal, 16 P
HeapAlloc: 204337544 HeapInuse: 204709888 HeapObjects: 748 HeapIdle 1319559168 HeapReleased 0 HeapSys 1524269056
gc 1427 @164.001s 0%: 0.011+0.23+0.086 ms clock, 0.17+0/0.36/0.22+1.3 ms cpu, 194->194->187 MB, 389 MB goal, 16 P (forced)
#BenchmarkSysHashMapInsertAndEraseWithBuf1E7-16    	10000000	       200 ns/op	      20 B/op	       0 allocs/op
PASS
scvg0: inuse: 2, idle: 2, sys: 5, released: 0, consumed: 5 (MB)
ok  	github.com/cdongyang/library/rbtree	164.179s
