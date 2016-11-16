package extractor

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"fmt"
)

type FeatureExtractor struct{
	Name string
	Data ListSensor

	// Temporal Features
	__avgX__, __avgY__, __avgZ__, __avgTotal__ float64
	__deviationX__,__deviationY__,__deviationZ__,__deviationTotal__ float64
	__skewnessX__,__skewnessY__,__skewnessZ__,__skewnessTotal__ float64
	__kurtosisX__,__kurtosisY__,__kurtosisZ__,__kurtosisTotal__ float64

	__correlationXY__,__correlationXZ__,__correlationXTotal__,
	__correlationYZ__,__correlationYTotal__,__correlationZTotal__ float64

	__zeroX__,__zeroY__,__zeroZ__,__zeroTotal__ float64

	// Frequential Features
	__fft_x__,__fft_y__,__fft_z__ []complex128
	__dcX__,__dcY__,__dcZ__	float64
	__energyX__,__energyY__,__energyZ__ float64
	__entropyX__,__entropyY__,__entropyZ__ float64
}

// Temporal features calculation
func (l *FeatureExtractor) CalcTemporalFeatures(){
	l.__avgX__ = l.avgX()
	l.__avgY__ = l.avgY()
	l.__avgZ__ = l.avgZ()
	l.__avgTotal__ = l.avgTotal()

	l.__deviationX__ = l.deviationX()
	l.__deviationY__ = l.deviationY()
	l.__deviationZ__ = l.deviationZ()
	l.__deviationTotal__ = l.deviationTotal()

	l.__skewnessX__ = l.skewnessX()
	l.__skewnessY__ = l.skewnessY()
	l.__skewnessZ__ = l.skewnessZ()
	l.__skewnessTotal__ = l.skewnessTotal()

	l.__kurtosisX__ = l.kurtosisX()
	l.__kurtosisY__ = l.kurtosisY()
	l.__kurtosisZ__ = l.kurtosisZ()
	l.__kurtosisTotal__ = l.kurtosisTotal()

	l.__correlationXY__ = l.corr_x_y()
	l.__correlationXZ__ = l.corr_x_z()
	l.__correlationXTotal__ = l.corr_x_total()
	l.__correlationYZ__ = l.corr_y_z()
	l.__correlationYTotal__ = l.corr_y_total()
	l.__correlationZTotal__ = l.corr_z_total()

	l.__zeroX__ = l.zero_crossing_rate_x()
	l.__zeroY__ = l.zero_crossing_rate_y()
	l.__zeroZ__ = l.zero_crossing_rate_z()
	l.__zeroTotal__ = l.zero_crossing_rate_avg()
}
// Frequential features calculation
func (l *FeatureExtractor) CalcFrequentialFeatures(){
	l.__fft_x__ = l.fftX()
	l.__fft_y__ = l.fftY()
	l.__fft_z__ = l.fftZ()

	l.__dcX__ = l.dcX()
	l.__dcY__ = l.dcY()
	l.__dcZ__ = l.dcZ()

	l.__energyX__ = l.energyX()
	l.__energyY__ = l.energyY()
	l.__energyZ__ = l.energyZ()

	l.__entropyX__ = l.entropyX()
	l.__entropyY__ = l.entropyY()
	l.__entropyZ__ = l.entropyZ()

}

// Method to return number of sensors registered
func (l *FeatureExtractor) length() int{
	return len(l.Data.GetSensors())
}

// Average of X-Axis
func (l *FeatureExtractor) avgX() float64{
	var avg float64 = 0
	for _,elt := range(l.Data.GetSensors()) {
		avg += elt.x
	}
	return avg / float64(l.length())
}
// Average of Y-Axis
func (l *FeatureExtractor) avgY() float64{
	var avg float64 = 0
	for _,elt := range(l.Data.GetSensors()) {
		avg += elt.y
	}
	return avg / float64(l.length())
}
// Average of Z-Axis
func (l *FeatureExtractor) avgZ() float64{
	var avg float64 = 0
	for _,elt := range(l.Data.GetSensors()) {
		avg += elt.z
	}
	return avg / float64(l.length())
}
// Average of All combined axis
func (l *FeatureExtractor) avgTotal() float64{
	return (l.__avgX__+l.__avgY__+l.__avgZ__)/3
}


// l'écart-type For X-Axis
func (l *FeatureExtractor) deviationX() float64{
	var avgX float64 = l.__avgX__
	var length = float64(l.length())
	var deviation float64 = 0

	for _,elt := range l.Data.GetSensors(){
		var x float64 = elt.x
		deviation += ((x-avgX)*(x-avgX))
	}
	return math.Sqrt(deviation/length)
}
// l'écart-type For Y-Axis
func (l *FeatureExtractor) deviationY() float64{
	var avgY float64 = l.__avgY__
	var length = float64(l.length())
	var deviation float64 = 0

	for _,elt := range l.Data.GetSensors(){
		var y float64 = elt.y
		deviation += ((y-avgY)*(y-avgY))
	}
	return math.Sqrt(deviation/length)
}
// l'écart-type For Z-Axis
func (l *FeatureExtractor) deviationZ() float64{
	var avgZ float64 = l.__avgZ__
	var length = float64(l.length())
	var deviation float64 = 0

	for _,elt := range l.Data.GetSensors(){
		var z float64 = elt.z
		deviation += ((z-avgZ)*(z-avgZ))
	}
	return math.Sqrt(deviation/length)
}
// Méthode retournant l'écart-type de l'ensemble des axes
func (l *FeatureExtractor) deviationTotal() float64{
	return (l.__deviationX__+l.__deviationY__+l.__deviationZ__)/3
}


// Skewness For X-Axis
func (l *FeatureExtractor) skewnessX() float64{
	var avg float64 = l.__avgX__
	var length float64 = float64(l.length())
	var deviation float64 = l.__deviationX__

	var skew float64 = 0

	for _, elt := range l.Data.GetSensors(){
		skew += math.Pow((elt.x - avg),3)/math.Pow(deviation,3)
	}
	skew *= (length / ((length - 1)*(length - 2)))
	return skew
}
// Skewness For Y-Axis
func (l *FeatureExtractor) skewnessY() float64{
	var avg float64 = l.__avgY__
	var length float64 = float64(l.length())
	var deviation float64 = l.__deviationY__

	var skew float64 = 0

	for _, elt := range l.Data.GetSensors(){
		skew += math.Pow((elt.y - avg),3)/math.Pow(deviation,3)
	}
	skew *= (length / ((length - 1)*(length - 2)))
	return skew
}
// Skewness For Y-Axis
func (l *FeatureExtractor) skewnessZ() float64{
	var avg float64 = l.__avgZ__
	var length float64 = float64(l.length())
	var deviation float64 = l.__deviationZ__

	var skew float64 = 0

	for _, elt := range l.Data.GetSensors(){
		skew += math.Pow((elt.z - avg),3)/math.Pow(deviation,3)
	}
	skew *= (length / ((length - 1)*(length - 2)))
	return skew
}
// Skewness For All combined axis
func (l *FeatureExtractor) skewnessTotal() float64{
	return (l.skewnessX()+l.skewnessY()+l.skewnessZ())/3
}


// Kurtosis For X-Axis
func (l *FeatureExtractor) kurtosisX() float64{
	var avg float64 = l.__avgX__
	var length float64 = float64(l.length())
	var deviation float64 = l.__deviationX__

	var kurtosis float64 = 0

	for _, elt := range l.Data.GetSensors(){
		kurtosis += math.Pow((elt.x - avg),4)/math.Pow(deviation,4)
	}

	kurtosis *= (((length)*(length+1))/((length-1)*(length-2)*(length-3)))
	kurtosis -= (3*math.Pow((length-1),2))/((length-2)*(length-3))
	return kurtosis
}
// Kurtosis For Y-Axis
func (l *FeatureExtractor) kurtosisY() float64{
	var avg float64 = l.__avgY__
	var length float64 = float64(l.length())
	var deviation float64 = l.__deviationY__

	var kurtosis float64 = 0

	for _, elt := range l.Data.GetSensors(){
		kurtosis += math.Pow((elt.y - avg),4)/math.Pow(deviation,4)
	}

	kurtosis *= (((length)*(length+1))/((length-1)*(length-2)*(length-3)))
	kurtosis -= (3*math.Pow((length-1),2))/((length-2)*(length-3))
	return kurtosis
}
// Kurtosis For Z-Axis
func (l *FeatureExtractor) kurtosisZ() float64{
	var avg float64 = l.__avgZ__
	var length float64 = float64(l.length())
	var deviation float64 = l.__deviationZ__

	var kurtosis float64 = 0

	for _, elt := range l.Data.GetSensors(){
		kurtosis += math.Pow((elt.z - avg),4)/math.Pow(deviation,4)
	}

	kurtosis *= (((length)*(length+1))/((length-1)*(length-2)*(length-3)))
	kurtosis -= (3*math.Pow((length-1),2))/((length-2)*(length-3))
	return kurtosis
}
// Kurtosis For Z-Axis
func (l *FeatureExtractor) kurtosisTotal() float64{
	return (l.kurtosisX()+l.kurtosisY()+l.kurtosisZ())/3
}


func (l *FeatureExtractor) __avg_corr_total() float64{
	var total float64 = 0
	for _, elt := range l.Data.GetSensors(){
		total += elt.x
		total += elt.y
		total += elt.z
	}
	total /= float64(l.length())
	return total
}
// Correlation between X and Y
func (l *FeatureExtractor) corr_x_y() float64{
	var mult float64 = 0
	for _, elt := range l.Data.GetSensors(){
		mult += (elt.x * elt.y)
	}
	mult /= float64(l.length())

	var cov float64 = mult - l.__avgX__*l.__avgY__
	var std float64 = l.__deviationX__ * l.__deviationY__
	return cov / std
}
// Correlation between X and Z
func (l *FeatureExtractor) corr_x_z() float64{
	var mult float64 = 0
	for _, elt := range l.Data.GetSensors(){
		mult += (elt.x * elt.z)
	}
	mult /= float64(l.length())

	var cov float64 = mult - l.__avgX__*l.__avgZ__
	var std float64 = l.__deviationX__ * l.__deviationZ__
	return cov / std
}
// Correlation between Y and Z
func (l *FeatureExtractor) corr_y_z() float64{
	var mult float64 = 0
	for _, elt := range l.Data.GetSensors(){
		mult += (elt.y * elt.z)
	}
	mult /= float64(l.length())

	var cov float64 = mult - l.__avgY__*l.__avgZ__
	var std float64 = l.__deviationY__ * l.__deviationZ__
	return cov / std
}
// Correlation between X and TOTAL
func (l *FeatureExtractor) corr_x_total() float64{
	var mult float64 = 0
	for _, elt := range l.Data.GetSensors(){
		mult += (elt.x *(elt.x + elt.y + elt.z))
	}
	mult /= float64(l.length())

	var cov float64 = mult - l.__avgX__*l.__avg_corr_total()
	var std float64 = l.__deviationX__ * l.__deviationTotal__
	return cov / std
}
// Correlation between Y and TOTAL
func (l *FeatureExtractor) corr_y_total() float64{
	var mult float64 = 0
	for _, elt := range l.Data.GetSensors(){
		mult += (elt.y *(elt.x + elt.y + elt.z))
	}
	mult /= float64(l.length())

	var cov float64 = mult - l.__avgY__*l.__avg_corr_total()
	var std float64 = l.__deviationY__ * l.__deviationTotal__
	return cov / std
}
// Correlation between Z and TOTAL
func (l *FeatureExtractor) corr_z_total() float64{
	var mult float64 = 0
	for _, elt := range l.Data.GetSensors(){
		mult += (elt.z *(elt.x + elt.y + elt.z))
	}
	mult /= float64(l.length())

	var cov float64 = mult - l.__avgZ__*l.__avg_corr_total()
	var std float64 = l.__deviationZ__ * l.__deviationTotal__
	return cov / std
}


// Zero-Crossing Rate on X-Axis
func (l *FeatureExtractor) zero_crossing_rate_x() float64{
	var zero_count float64 = 0
	for index,elt := range l.Data.GetSensors(){
		if(index == 0){
			continue
		}
		if(l.Data.GetSensors()[index-1].x*elt.x < 0){
			zero_count ++
		}
	}
	zero_count /= float64(l.length())
	return zero_count
}
// Zero-Crossing Rate on Y-Axis
func (l *FeatureExtractor) zero_crossing_rate_y() float64{
	var zero_count float64 = 0
	for index,elt := range l.Data.GetSensors(){
		if(index == 0){
			continue
		}
		if(l.Data.GetSensors()[index-1].y*elt.y < 0){
			zero_count ++
		}
	}
	zero_count /= float64(l.length())
	return zero_count
}
// Zero-Crossing Rate on Z-Axis
func (l *FeatureExtractor) zero_crossing_rate_z() float64{
	var zero_count float64 = 0
	for index,elt := range l.Data.GetSensors(){
		if(index == 0){
			continue
		}
		if(l.Data.GetSensors()[index-1].z*elt.z < 0){
			zero_count ++
		}
	}
	zero_count /= float64(l.length())
	return zero_count
}
// Zero-Crossing Rate on All combined axis
func (l *FeatureExtractor) zero_crossing_rate_avg() float64{
	return (l.zero_crossing_rate_x() + l.zero_crossing_rate_y() + l.zero_crossing_rate_z()) / 3
}


// FFT For X-Axis
func (l *FeatureExtractor) fftX() []complex128{
	var list []float64
	for _, elt := range l.Data.GetSensors(){
		list = append(list,elt.x)
	}
	return fft.FFTReal(list)
}
// FFT For Y-Axis
func (l *FeatureExtractor) fftY() []complex128{
	var list []float64
	for _, elt := range l.Data.GetSensors(){
		list = append(list,elt.y)
	}
	return fft.FFTReal(list)
}
// FFT For Z-Axis
func (l *FeatureExtractor) fftZ() []complex128{
	var list []float64
	for _, elt := range l.Data.GetSensors(){
		list = append(list,elt.z)
	}
	return fft.FFTReal(list)
}


// DC Component For X-Axis
func (l *FeatureExtractor) dcX() float64{
	var result float64 = 0
	for _,elt := range l.__fft_x__{
		result += math.Pow(real(elt),2)
	}
	result /= float64(len(l.__fft_x__))
	return result
}
// DC Component For Y-Axis
func (l *FeatureExtractor) dcY() float64{
	var result float64 = 0
	for _,elt := range l.__fft_y__{
		result += math.Pow(real(elt),2)
	}
	result /= float64(len(l.__fft_y__))
	return result
}
// DC Component For Z-Axis
func (l *FeatureExtractor) dcZ() float64{
	var result float64 = 0
	for _,elt := range l.__fft_z__{
		result += math.Pow(real(elt),2)
	}
	result /= float64(len(l.__fft_z__))
	return result
}


// Energy for X-Axis
func (l *FeatureExtractor) energyX() float64{
	var result float64 = 0
	for _,elt := range l.__fft_x__{
		result += (math.Pow(real(elt),2) + math.Pow(imag(elt),2))
	}
	result /= float64(len(l.__fft_x__))
	return result
}
// Energy for Y-Axis
func (l *FeatureExtractor) energyY() float64{
	var result float64 = 0
	for _,elt := range l.__fft_y__{
		result += (math.Pow(real(elt),2) + math.Pow(imag(elt),2))
	}
	result /= float64(len(l.__fft_y__))
	return result
}
// Energy for Z-Axis
func (l *FeatureExtractor) energyZ() float64{
	var result float64 = 0
	for _,elt := range l.__fft_z__{
		result += (math.Pow(real(elt),2) + math.Pow(imag(elt),2))
	}
	result /= float64(len(l.__fft_z__))
	return result
}


// Entropy subcalculation
func (l *FeatureExtractor) __entropy_subcalcul(cplx complex128, N float64, Energy float64) float64{
	return (math.Pow(real(cplx),2) + math.Pow(imag(cplx),2)) / (N - Energy )
}
// Entropy for X-Axis
func (l *FeatureExtractor) entropyX() float64{
	var result float64 = 0
	var length = float64(len(l.__fft_x__))
	for _,elt := range l.__fft_x__{
		result += l.__entropy_subcalcul(elt,length,l.__energyX__)
	}
	result *= -1
	return result
}
// Entropy for Y-Axis
func (l *FeatureExtractor) entropyY() float64{
	var result float64 = 0
	var length = float64(len(l.__fft_y__))
	for _,elt := range l.__fft_y__{
		result += l.__entropy_subcalcul(elt,length,l.__energyY__)
	}
	result *= -1
	return result
}
// Entropy for Z-Axis
func (l *FeatureExtractor) entropyZ() float64{
	var result float64 = 0
	var length = float64(len(l.__fft_z__))
	for _,elt := range l.__fft_z__{
		result += l.__entropy_subcalcul(elt,length,l.__energyZ__)
	}
	result *= -1
	return result
}

// Méthode affichant un résumé
func (l *FeatureExtractor) Print() {
	/*
	start_1 := time.Now()
	l.calc_feature_temporal()
	elapsed_1 := time.Since(start_1)

	start_2 := time.Now()
	l.calc_feature_frequential()
	elapsed_2 := time.Since(start_2)

	fmt.Printf("Temporal features took %s\n", elapsed_1 )
	fmt.Printf("Frequential features took %s\n", elapsed_2 )
	*/

	var printed string = ""
	printed += "\n______________________\n"
	printed += "---------------------------------\n"
	printed += " Name : %s (%v)\n"
	printed += "\n"
	printed += "------------------------- \n"
	printed += "-- TEMPORAL DIMENSION --- \n"
	printed += "------------------------- \n"
	printed += "   1 = Mean value :\n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "   2 = Mean value (Total) : %v\n"
	printed += "   3 = Standard Deviation : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "   4 = Standard Deviation (Total) : %v\n"
	printed += "   5 = Skewness : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "   6 = Skewness (Total) : %v\n"
	printed += "   7 = Kurtosis : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "   8 = Kurtosis (Total) : %v\n"
	printed += "   11= Correlation : \n"
	printed += "     - X_Y     : %v\n"
	printed += "     - X_Z     : %v\n"
	printed += "     - X_TOTAL : %v\n"
	printed += "     - Y_Z     : %v\n"
	printed += "     - Y_TOTAL : %v\n"
	printed += "     - Z_TOTAL : %v\n"
	printed += "   12= Zero Crossing Rate : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "     - Average : %v\n"
	printed += "------------------------- \n"
	printed += "--FREQUENTIAL DIMENSION-- \n"
	printed += "------------------------- \n"
	printed += "   1 = FFT length: X:(%v) // Y:(%v) // Z:(%v)\n"
	printed += "   2 = DC Component : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "   3 = Energy : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"
	printed += "   4 = Entropy : \n"
	printed += "     - X : %v\n"
	printed += "     - Y : %v\n"
	printed += "     - Z : %v\n"


	fmt.Printf(printed,
		l.Name,l.length(),
		l.__avgX__,l.__avgY__,l.__avgZ__,l.__avgTotal__,
		l.__deviationX__,l.__deviationY__,l.__deviationZ__,l.__deviationTotal__,
		l.__skewnessX__,l.__skewnessY__,l.__skewnessZ__, l.__skewnessTotal__,
		l.__kurtosisX__,l.__kurtosisY__,l.__kurtosisZ__, l.__kurtosisTotal__,
		l.__correlationXY__,l.__correlationXZ__,l.__correlationXTotal__,
		l.__correlationYZ__,l.__correlationYTotal__,l.__correlationZTotal__,
		l.__zeroX__,l.__zeroY__,l.__zeroZ__,l.__zeroTotal__,

		len(l.__fft_x__),len(l.__fft_y__),len(l.__fft_z__),
		l.__dcX__,l.__dcY__,l.__dcZ__,
		l.__energyX__,l.__energyY__,l.__energyZ__,
		l.__entropyX__,l.__entropyY__,l.__entropyZ__)
}