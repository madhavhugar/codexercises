package main

// https://leetcode.com/contest/biweekly-contest-32/problems/kth-missing-positive-number/

func findKthPositive(arr []int, k int) int {
	originalK := k
	first := arr[0]
	if first != 1 {
		k = k - first + 1
	}
	if k-1 < 0 {
		return originalK
	}
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		currK := k
		if diff > 1 {
			k = k - diff + 1
		}
		if k <= 0 {
			return arr[i] + currK
		}
	}
	return arr[len(arr)-1] + k
}

/* Other better solutions

// O(nlog(n))

 int findKthPositive(vector<int>& arr, int k) {
        // if first element is greater than k return k
        if(arr.empty() || arr[0] > k)
            return k;
        int n = arr.size();
        // if last element is smaller than (length of arr + k) return (length of arr + k)
        if(arr[n-1] < n+k)
            return n+k;

        int l = 0, r = n-1;
        while(l < r)
        {
            int m = (l+r)/2;
            if((arr[m]-(m+1)) < k)
                l = m+1;
            else
                r = m;
        }
        return k+l;
    }


	// O(n)
    var findKthPositive = function(arr, k) {
    let index = 0;
    let current = 1;
    let missing = 0;

    while (index < arr.length) {
        if (arr[index] === current) {
            ++index;
        } else {
            ++missing;
        }

        if (missing === k) return current;

        ++current;
    }

    return current + k - 1 - missing;
};

*/
