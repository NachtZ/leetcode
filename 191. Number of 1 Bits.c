int hammingWeight(uint32_t n) {
    	int count = 0;
	while (n / 2)
	{
		count += n % 2;
		n /= 2;
	}
	count += n % 2;
	return count;
}