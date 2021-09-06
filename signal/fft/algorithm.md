Based on the Data reordering, bit reversal, and in-place algorithms section of
[Cooley-Tukey FFT](https://en.wikipedia.org/wiki/Cooley-Tukey_FFT_algorithm)

    algorithm iterative-fft is
        input: Array a of n complex values where n is a power of 2.
        output: Array A the DFT of a.

        bit-reverse-copy(a, A)
        n ← a.length
        for s = 1 to log(n) do
            m ← 2^s
            ωm ← exp(−2πi/m)
            for k = 0 to n-1 by m do
                ω ← 1
                for j = 0 to m/2 – 1 do
                    t ← ω A[k + j + m/2]
                    u ← A[k + j]
                    A[k + j] ← u + t
                    A[k + j + m/2] ← u – t
                    ω ← ω ωm
        return A
