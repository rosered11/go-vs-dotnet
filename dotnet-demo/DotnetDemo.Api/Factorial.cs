using System;

namespace DotnetDemo.Api;

public class Factorial
{
    public int Calculate(int n)
    {
        if (n <= 1) return 1;
        return n * Calculate(n - 1);
    }
}
