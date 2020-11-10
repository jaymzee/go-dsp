## Filter package

### direct form 1 realization

      a[0] is assumed to be 1.0

    start with the Z transform of a filter with transfer function H(z):

      Y(z) = H(z)·X(z)
      Y(z) = N(z)/D(z)·X(z)
      D(z)·Y(z) = N(z)·X(z)

    Taking the inverse Z transform and collecting terms to one side gives:

      y[n] =   b[0]x[n]   + b[1]x[n-1] + ... + b[L]x[n-L]
             - a[1]y[n-1] - a[2]y[n-2] - ... - a[M]y[n-M]

                   b0
      x[n] ----┬---|>-→(+)-------┬---→ y[n]
               ↓        ↑        ↓
              [z]  b1   |  -a1  [z]
               ├---|>-→(+)←-<|---┤
               ↓        ↑        ↓
              [z]  b2   |  -a2  [z]
               └---|>-→(+)←-<|---┘

### BiQuad
BiQuad is a direct form I filter where the a and b polynomials are quadratic
``` go
b := [3]Float64{0.0045, 1.0, -0.0045} // numerator (filter zeros)
a := [3]Float64{1.0, -1.98, 0.991}    // denominator (filter poles)
y := BiQuad(b, a, x)
```

### DirectForm1
DirectForm1 is a direct form I filter where the a and b polynomials are
of any degree
``` go
b := []Float64{0.0045, 1.0, -0.0045} // numerator (filter zeros)
a := []Float64{1.0, -1.98, 0.991}    // denominator (filter poles)
y := BiQuad(b, a, x)
```
