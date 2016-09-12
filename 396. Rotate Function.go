func maxRotateFunction(A []int) int {
    sum,f:= 0,0
	for i:=0;i<len(A);i++{
		sum += A[i]
		f += i*A[i]
	}
	m := f
	for i:=1;i<len(A);i++{
		t :=f+sum - A[len(A)-i]*len(A)
		if m < t{
			m = t
		}
		f += sum - A[len(A)-i]*len(A)
	}
	return m
}