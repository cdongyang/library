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
	/*var benchRand = randint.Rand{First: 23456, Add: 12345, Mod: 1e9 + 7}*/
	randint benchRand=randint{23456,12345,int(1e9+7)};
	for(int i=100;i<=1e7;i*=10) {
		randint rand = benchRand;
		clock_t allStart=clock();

		double allInsert = 0,allErase = 0;
		int times = 0;
		for(times=0;times<=max(1,int(1e6)/i);times++) {
			clock_t start = clock();
			int *arr = new int[int(1e7)];
			for(int j=0;j<i;j++) {
				arr[i]=rand.Int();
				s.insert(arr[i]);
			}
			clock_t end = clock();
			allInsert += end-start;
			
			//set<int>ts;
			//ts.insert(0);
			//s.erase(s.end());
			//set<int>::iterator it;
			//s.erase(it);
			//s.erase(ts.begin());

			start=clock();
			for(int j=0;j<i;j++) {
				s.erase(arr[i]);
			}
			end=clock();
			allErase += end-start;
		}
		allInsert=allInsert * 1e9/(CLOCKS_PER_SEC*times);
		allErase=allErase * 1e9/(CLOCKS_PER_SEC*times);
		printf("insert %d elem use time %.0fns %.0fns/op\n",i,allInsert,allInsert/i);
		printf("erase %d elem use time %.0fns %.0fns/op\n",i,allErase,allErase/i);

		clock_t allEnd=clock();
		double allInsertAndErase = double(allEnd-allStart)*1e9/CLOCKS_PER_SEC/times;
		printf("insert %d elem and erase %d elem use time %.0fns %.0fns/op\n",i,i,allInsertAndErase,allInsertAndErase/i);
	}
}