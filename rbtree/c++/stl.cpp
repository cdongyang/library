#include<bits/stdc++.h>
using namespace std;

struct randint {
	int first,add,mod;
	int Int() {
		first = (first+add)%mod;
		return first;
	}
};
int main() {
	set<int>s;
	//var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}
	randint benchRand=randint{23456,12345,int(1e9+7)};
	randint rand = benchRand;
	for(int i=100;i<=1e7;i*=10) {
		s.clear();
		clock_t start = clock();
		for(int j=0;j<i;j++) {
			s.insert(rand.Int());
		}
		clock_t end = clock();
		printf("insert %d use time %.3f s %.0fns/op\n",i,double(end-start)/CLOCKS_PER_SEC,double(end-start)/CLOCKS_PER_SEC*1e9/i);
	}
}