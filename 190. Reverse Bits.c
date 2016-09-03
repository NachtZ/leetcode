uint32_t reverseBits(uint32_t n) {
    	int b[40] = {0};
	int i = 0;
	uint32_t c = 1;
	while (n / 2)
	{
		b[i++] = n % 2;
		n = n / 2;
	}
	b[i++] = n % 2;
	n = 0;
	for (i = 31; i >=0; i--)
	{
		n += b[i] * c;
		c *= 2;
	}
	return n;
}