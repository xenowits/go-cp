//xenowi`tz -- Jai Shree Ram
#include<bits/stdc++.h>
using namespace std;

#define fori(i,a,b) for (long long int i = a; i <= b ; ++i)
#define ford(i,a,b) for(long long int i = a;i >= b ; --i)
#define mk make_pair
#define mod 1000000007
#define MAXN 200005
#define pb push_back
#define vec vector<long long int>
#define rnd mt19937_64 rng(chrono::high_resolution_clock::now().time_since_epoch().count())
#define pi pair<long long int,long long int>
#define sc second
#define fs first

int main() {

        long long n, m;
        cin >> n >> m;

        vector<long long> arr(n);
        for(int i = 0; i < n; i++) {
            cin >> arr[i];
        }

        sort(arr.begin(), arr.end());

        long long no_of_times = n/m;

        long long ans = 0, seg_ans = 0;

        for (int i = 0; i < no_of_times; i += 1) {
            seg_ans = 0;
            for (int j = i*m; j < i*m+m; j++) {
                seg_ans = (seg_ans + arr[j])%mod;
            }
            //cout << seg_ans << endl;
            ans = (ans + (i+1)*seg_ans)%mod;
        }

        seg_ans = 0;
        for (int i = m*no_of_times; i < n; i++) {
            seg_ans = (seg_ans + arr[i])%mod;
        }

        ans = (ans + no_of_times*seg_ans)%mod;

        cout << ans << endl;

}
