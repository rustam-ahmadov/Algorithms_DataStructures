#include<iostream>
#include<string.h>



//O(n^2)
int gemsCount(std::string& s, std::string&j){

    int gems_coumt = 0;

    for(int i = 0; i < s.size(); i ++)
    {
        for(int k = 0; k < j.size(); k++)
        {   
            if(s[i] == j[k]){
                gems_coumt++;
                break;
            }
        }
    }
    return gems_coumt;
}

//O(n)
int gemsCountWithTable(std::string& s, std::string& j)
{
    const int ASCI_LET_START_POS = 97;
    const int JEWELERY_STONES_SIZE = j.size();
    const int ALL_STONES_SIZE = 25;
    int gems_table[ALL_STONES_SIZE];
    
    for(int i = 0; i < ALL_STONES_SIZE; i++)
        gems_table[i] = 0;

    for(int i = 0; i < JEWELERY_STONES_SIZE; i++)
    {
        gems_table[j[i]- ASCI_LET_START_POS]++;
    }


    const int COMMON_STONES_SIZE = s.size();
    int gems_count_in_stones = 0;
    for(int i = 0; i < COMMON_STONES_SIZE; i++)
    {
        int asci_sym_of_stone = gems_table[s[i] - ASCI_LET_START_POS];
        if(asci_sym_of_stone > 0)        
            gems_count_in_stones++;        
    }

    return gems_count_in_stones;
}

int main(int argc, char** argv){

    std::string j = "aeyuio";
    std::string s = "rustiammaaaaaaaatsr";

    
    int gems_count = gemsCountWithTable(s,j);
    std::cout<<gems_count;
    


    
}