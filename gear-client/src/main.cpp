#include <Arduino.h>


int GEAR_1 = A1;
int GEAR_2 = A2;
int GEAR_3 = A3;
int GEAR_4 = A4;
int GEAR_5 = A5;
int GEAR_6 = A0;

void setup() {
    pinMode(GEAR_1, INPUT);
    pinMode(GEAR_2, INPUT);
    pinMode(GEAR_3, INPUT);
    pinMode(GEAR_4, INPUT);
    pinMode(GEAR_5, INPUT);
    pinMode(GEAR_6, INPUT);

    Serial.begin(9600);
    Serial.println("G_BG");
    Serial.println("");
}

int gear_status = 1;

int last_gear = 0;
int tick_time = 0;

void loop() {
    int gear1 = digitalRead(GEAR_1);
    int gear2 = digitalRead(GEAR_2);
    int gear3 = digitalRead(GEAR_3);
    int gear4 = digitalRead(GEAR_4);
    int gear5 = digitalRead(GEAR_5);
    int gear6 = digitalRead(GEAR_6);


    if (gear1 == HIGH) {
        gear_status = 1;
    } else if (gear2 == HIGH) {
        gear_status = 2;
    } else if (gear3 == HIGH) {
        gear_status = 3;
    } else if (gear4 == HIGH) {
        gear_status = 4;
    } else if (gear5 == HIGH) {
        gear_status = 5;
    } else if (gear6 == HIGH) {
        gear_status = 6;
    }
//    Serial.println(String(gear1 == HIGH) + " " + String(gear2 == HIGH) + " " + String(gear3 == HIGH) + " " + String(gear4 == HIGH) + " " + String(gear5 == HIGH) + " " + String(gear6 == HIGH));
//    Serial.println(String(last_gear) + " " + String(gear_status));
    if (last_gear != gear_status) {
        last_gear = gear_status;
        Serial.print("G_CH_TO_");
        Serial.println(String(gear_status));
    }
//    Serial.println("tick " + String(tick_time));
    delay(100);
    tick_time += 1;
}