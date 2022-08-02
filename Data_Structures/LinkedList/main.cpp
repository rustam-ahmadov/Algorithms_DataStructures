#include"Linkedlist.h"
#include<time.h>;
#define __CRTDBG_MAP_ALLOC
#include <crtdbg.h>





int main()
{
	srand(time(NULL));
	{
		
		List<int>l1;
		List<int>l2;
		List<int>l4;
		List<int>l5;
		l4.push_back(80);
		l4.push_back(90);
		l4.push_back(100);
		l5.push_back(110);
		l5.push_back(120);
		l5.push_back(130);

		l1.push_back(1);
		l1.push_back(2);
		l1.push_back(3);
		l1.push_back(4);	
		l1.print_list();
		l2.push_back(5);
		l2.push_back(6);
		l2.push_back(7);
		l2.push_back(8);
		l2.push_back(9);
		l2.push_back(10);
		l2.print_list();
	
		List<int>l3(l2+l1);
		l3.print_list();
		l3 = l4 + l5;
		l3.print_list();
		//l3 = l5 + l4;
		//l3.print_list();
	

		/*l2.push_back(6);
		l2.push_back(7);
		l2.push_back(8);
		l2.push_back(9);
		l2.print_list();*/


		//l1.swap(l2);

		//l1.print_list();
		//l2.print_list();
		//l1.remove(2);
		//l1.print_list();
		

		

	}
	
	_CrtDumpMemoryLeaks();

	return 0;
}