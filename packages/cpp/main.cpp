#include <iostream>
#include <string>

class User {
public:
    std::string name;
    User(std::string n) : name(n) {}
};

int main() {
    User* user = new User("John");
    delete user;
    // Use-after-free vulnerability
    std::cout << "User name: " << user->name << std::endl;
    return 0;
} 