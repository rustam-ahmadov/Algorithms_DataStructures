using System.Linq;

namespace MyNamespace
{
    public class Solution
    {
        //O(n)
            public static int RomanToInt(string s)
            {
                int res = 0;

                for (int i = 0; i < s.Length; i++)
                {
                    if (i < (s.Length - 1) && RomanCharToInt(s[i]) < RomanCharToInt(s[i + 1]))
                    {
                        res -= RomanCharToInt(s[i]);
                    }
                    else
                    {
                        res += RomanCharToInt(s[i]);
                    }
                }

                return res;
            }
            public static int RomanCharToInt(char c)
            {
                switch (c)
                {
                    case 'I': return 1;
                    case 'V': return 5;
                    case 'X': return 10;
                    case 'L': return 50;
                    case 'C': return 100;
                    case 'D': return 500;
                    case 'M': return 1000;
                    default: return 0;
                }
            }


        static void Main()
        {
            int num = RomanToInt("MCMXCIV");
            System.Console.WriteLine(num);
        }
    }

}
