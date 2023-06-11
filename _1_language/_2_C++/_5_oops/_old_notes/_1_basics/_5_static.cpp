// In C++, static is a keyword or modifier that belongs to the type not instance.
// So instance is not required to access the static members.
// In C++, static can be field, method, constructor, class, properties, operator and event.

// Advantage of C++ static keyword
// Memory efficient: Now we don't need to create instance for accessing the static members,
// so it saves memory. Moreover, it belongs to the type, so it will not get memory each time when instance is create

#include <iostream>
using namespace std;

class Account
{
public:
    int accno;   //data member (also instance variable)
    string name; //data member(also instance variable)
    static float rateOfInterest;
    static int count;
    Account(int accno, string name)
    {
        this->accno = accno;
        this->name = name;
        count++;
    }
    void display()
    {
        cout << accno << " " << name << " " << rateOfInterest << endl;
    }
    static void totalObjects()
    {
        cout << "Total Objects are: " << count;
    }
};

float Account::rateOfInterest = 6.5;
int Account::count = 0;

int main(void)
{
    Account a1 = Account(201, "Sanjay"); //creating an object of Employee
    Account a2 = Account(202, "Nakul");  //creating an object of Employee
    a1.display();
    a2.display();
    Account::totalObjects();
    return 0;
}