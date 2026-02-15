#include <iostream>
#include <cmath>

int main(){
    int a;
    int b;
    int j = 0;
    std::cin >> a;
    std::cin >> b;
    if ((a >= 2) && (b <= 2 * int(pow(10, 9)))){
        for (int n = 0; n < 10; n++){
            if ((int(pow(b, n)) % a) == 0){
                std::cout << "Минимальная степень числа B = " << n << std::endl;
                j = 1;
                break;
            
            }
        }

        if (j == 0){
            std::cout << "-1" << std::endl;
        }
    }

    else{
        std::cout << "Введите другие числа, они должны подходить под условие" << std::endl;
    }
    return 0;
}

    
