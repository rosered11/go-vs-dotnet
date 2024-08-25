using System;

namespace DotnetDemo.Api;

public class Heavy
{
    public int CPUBurn(int n)
    {
        int result = 1;
        for(int i = 0; i < n; i++)
            result *= i;
        return result;
    }
    public long MemoryBurn(long n)
    {
        List<uint> largeList = new();
        for(int i = 0; i < n; i++)
            largeList.Add(Convert.ToUInt32(i));
        // return largeList.Sum();
        long sum = 0;
        foreach(var num in largeList)
            sum += num;
        return sum;
    }
}
