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
	for(int i=100;i<=1e7;i*=10) {
		randint rand = benchRand;
		clock_t allStart=clock();

		for(int times=0;times<=max(1,i/int(1e6));times++) {
			clock_t start = clock();
			int *arr = new int[int(1e7)];
			for(int j=0;j<i;j++) {
				arr[i]=rand.Int();
				s.insert(arr[i]);
			}
			clock_t end = clock();
			printf("insert %d elem use time %.3fs %.0fns/op\n",i,double(end-start)/CLOCKS_PER_SEC,double(end-start)/CLOCKS_PER_SEC*1e9/i);

			start=clock();
			for(int j=0;j<i;j++) {
				s.erase(arr[i]);
			}
			end=clock();
			printf("erase %d elem use time %.3fs %.0fns/op\n",i,double(end-start)/CLOCKS_PER_SEC,double(end-start)/CLOCKS_PER_SEC*1e9/i);
		}

		clock_t allEnd=clock();
		printf("insert and erase %d elem use time %.3fs %.0fns/op\n",i,double(allEnd-allStart)/CLOCKS_PER_SEC,double(allEnd-allStart)/CLOCKS_PER_SEC*1e9/i);
	}
}