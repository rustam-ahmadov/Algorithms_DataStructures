string longestCommonPrefix(string[] strs)
{
    if (0 == strs.Length)
        return "";

    string prefix = strs[0];

    for (int i = 1; i < strs.Length; i++)
        while (strs[i].IndexOf(prefix) != 0)
        {
            prefix = prefix.Substring(0, prefix.Length - 1);
            if (string.IsNullOrEmpty(prefix))
            {

                return "";
            }
        }

    return prefix;
}