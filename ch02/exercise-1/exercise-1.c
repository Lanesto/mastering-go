#include <stdio.h>
#include "exercise-1.h"

void helloFromC(char* additionalMessage) {
	printf("C got %s from Go!\n", additionalMessage);
}
