package extractor



// Definition of Sensor type
type Sensor struct {
	x,y,z float64
}

// Definition of ListSensor type, which contains all Sensors for a dataset
//	Example: A dataset of 500 rows will have a ListSensor of 500 Sensors instances
//	Example2:A 9-DOF dataset of 500 rows will have 3 ListSensor of 500Sensors each (accel, gyro, mag)
type ListSensor struct {
	sensors []Sensor
}
func(sensors *ListSensor) AddSensor(sensor Sensor){
	sensors.sensors = append(sensors.sensors, sensor)
}
func(sensors *ListSensor) GetSensors() []Sensor{
	return sensors.sensors
}

/*
	Method to create ListSensor
	data:=  		is the dataset in float64
	column_index:=  is a slice with indexes of columns for the sensor
						==> column_index[0] will be "x"
						==> column_index[1] will be "y"
						==> column_index[2] will be "z"

	return: ListSensor 
 */


func GetSensorFromData(data [][]float64, column_index [3]int8) ListSensor{
	var sensors ListSensor
	for _, row := range(data){
		var sensor Sensor
		var c = 0
		for _, index := range(column_index){
			switch(c){
				case 0:sensor.x = row[index];	break;
				case 1:sensor.y = row[index];	break;
				case 2:sensor.z = row[index];	break;
			}
			c++
		}
		sensors.AddSensor(sensor)
	}
	return sensors
}
func GetSensorsData(float_data [][]float64) (ListSensor,ListSensor,ListSensor){
	return GetSensorFromData(float_data,[3]int8{0,1,2}),
	GetSensorFromData(float_data,[3]int8{3,4,5}),
	GetSensorFromData(float_data,[3]int8{6,7,8});
}