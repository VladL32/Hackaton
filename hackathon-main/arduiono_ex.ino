#include <FirebaseESP32.h> 
#include <ESP8266WiFi.h> 
#include <Wire.h>
#include <Adafruit_Sensor.h>
#include <Adafruit_BME280.h>

#define SEALEVELPRESSURE_HPA (1013.25)

// BME280 Temperature, Humidity, Pressure sensor
Adafruit_BME280 bme;

//Connection to wifi and firebase
#define FIREBASE_HOST "https://hackathon-1018f-default-rtdb.firebaseio.com" 
#define FIREBASE_AUTH "tAIzaSyB3frk5hk2p3mlt-Iww8dJJgcndjRSPsqQ" 
#define WIFI_SSID "Study" 
#define WIFI_PASSWORD "studyaitu" 

FirebaseData firebaseData; 

void wifi() { 
  Serial.begin(9600); 
  // Connect to Wi-Fi 
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD); 
  while (WiFi.status() != WL_CONNECTED) { 
    delay(500); 
    Serial.print("."); 
  } 
  Serial.println(); 
  Serial.print("Connected to WiFi. IP address: "); 
  Serial.println(WiFi.localIP()); 
  // Initialize Firebase 
  Firebase.begin(FIREBASE_HOST, FIREBASE_AUTH); 
  } 

void setup() {
  Serial.begin(9600);
  // Initialize BME280 sensor
  if (!bme.begin(0x76)) {
    Serial.println("Could not find a valid BME280 sensor, check wiring!");
    while (1);
  }
}


void loop() {
  // Read temperature, humidity, and pressure from BME280
  float temperature = bme.readTemperature();
  float humidity = bme.readHumidity();
  float pressure = bme.readPressure() / 100.0F;

  // MQ-2 Gas Sensor
  const int MQ2_PIN = A0;

  // Water Sensor
  const int WATER_SENSOR_PIN = 2;
  
  pinMode(MQ2_PIN, INPUT);
  pinMode(WATER_SENSOR_PIN, INPUT);

  int gasValue = analogRead(MQ2_PIN);
  int waterValue = digitalRead(WATER_SENSOR_PIN);

  // Print data
  Serial.print("Temperature: ");
  Serial.print(temperature);
  Serial.print(" °C, Humidity: ");
  Serial.print(humidity);
  Serial.print(" %, Pressure: ");
  Serial.print(pressure);
  Serial.println(" hPa");
  Serial.print("MQ-2 Gas Value: ");
  Serial.println(gasValue);
  Serial.print("Water Sensor Value: ");
  Serial.println(waterValue);

  // Create a JSON object to send to Firebase 
  FirebaseJson json; 
  json.add("pressure", pressure);
  json.add("temperature", temperature); 
  json.add("humidity", humidity); 
  json.add("waterValue", waterValue); 
  json.add("gasValue", gasValue); 
  
  // Send the JSON object to Firebase 
  if (Firebase.pushJSON(firebaseData, "/sensor-data", json)) { 
    Serial.println("Data sent to Firebase successfully!"); 
  } else { 
    Serial.println("Failed to send data to Firebase. Error: " + firebaseData.errorReason()); 
  } 
  delay(5000);  // Wait for 5 seconds before reading again 
  }


