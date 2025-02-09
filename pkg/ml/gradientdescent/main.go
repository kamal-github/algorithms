package main

import (
	"fmt"
	"math"
)

type GradientDescent struct{}

// Train applies gradient descent to find the best w and b
func (_ GradientDescent) Train(trainingSet [][2]float64) (float64, float64) {
	if len(trainingSet) == 0 {
		return -1, -1
	}

	w, b := 0.01, 0.0  // weight and bias in ŷ = wx + b
	alpha := 0.01      // learning rate
	epsilon := 0.00001 // tolerance rate - upto which we are ok to get the prediction, as getting 0 is not always practical or necessary.

	for i := 0; i < 20000; i++ {
		derivOverWeight, derivOverBias := derivativeOfCostFn(trainingSet, w, b)
		if math.IsNaN(derivOverWeight) || math.IsNaN(derivOverBias) || math.IsInf(derivOverWeight, 0) || math.IsInf(derivOverBias, 0) {
			fmt.Printf(
				"ALERT: w and m became NaN: w=%f, b=%f, derivOverBias=%f, derivOverWeight=%f",
				w, b, derivOverBias, derivOverWeight,
			)
			break
		}

		// Stop when gradients are very small
		if math.Abs(derivOverWeight) < epsilon || math.Abs(derivOverBias) < epsilon {
			break
		}

		// Update step
		w = w - alpha*derivOverWeight
		b = b - alpha*derivOverBias
	}

	return w, b
}

// derivativeOfCostFn calculates the partial derivatives of the cost function
func derivativeOfCostFn(trainingSet [][2]float64, w, b float64) (float64, float64) {
	dw, db := 0.0, 0.0
	n := float64(len(trainingSet))

	for _, set := range trainingSet {
		prediction := w*set[0] + b
		costFn := prediction - set[1]
		dw += costFn * set[0]
		db += costFn
	}

	return dw / n, db / n
}

func (_ GradientDescent) PredictPrice(plotSize float64, w float64, b float64) float64 {
	return w*plotSize + b
}

func normalize(ts [][2]float64) ([][2]float64, float64, float64, float64, float64) {
	maxX, minX, minY, maxY := ts[0][0], ts[0][0], ts[0][1], ts[0][1]

	for _, xi := range ts {
		if maxX < xi[0] {
			maxX = xi[0]
		}
		if minX > xi[0] {
			minX = xi[0]
		}

		if maxY < xi[1] {
			maxY = xi[1]
		}
		if minY > xi[1] {
			minY = xi[1]
		}

	}

	normalizedTS := make([][2]float64, len(ts))
	for i, xi := range ts {
		normalizedTS[i] = [2]float64{(xi[0] - minX) / (maxX - minX), (xi[1] - minY) / (maxY - minY)}
	}

	return normalizedTS, minX, maxX, minY, maxY
}

func minMaxNormalize(val, min, max float64) float64 {
	return (val - min) / (max - min)
}

func minMaxDeNormalize(normalizedVal, min, max float64) float64 {
	return (normalizedVal * (max - min)) + min
}

func main() {
	var trainingSet = [][2]float64{
		{523, 115.3}, {1842, 443.1}, {3765, 1015.6}, {2894, 678.4}, {2510, 589.2},
		{4521, 1158.3}, {1342, 289.6}, {4930, 1320.5}, {3742, 996.7}, {2911, 701.2},
		{1734, 390.4}, {1123, 268.6}, {2645, 608.3}, {4389, 1140.2}, {1492, 339.3},
		{2875, 675.8}, {1984, 479.1}, {3165, 752.7}, {2201, 521.9}, {3489, 876.4},
		{2711, 624.1}, {4783, 1265.8}, {1584, 366.3}, {4012, 995.7}, {1998, 466.2},
		{3678, 927.6}, {2483, 575.4}, {4211, 1102.3}, {1323, 293.7}, {4921, 1333.4},
		{3456, 870.2}, {1992, 474.8}, {1784, 403.2}, {3125, 765.4}, {2389, 553.1},
		{4621, 1192.9}, {1348, 314.6}, {2932, 713.2}, {1532, 358.7}, {2564, 605.9},
		{4841, 1299.2}, {1723, 398.2}, {3172, 759.5}, {2658, 618.3}, {1821, 419.6},
		{4762, 1247.1}, {3284, 813.2}, {1981, 482.5}, {3892, 1029.6}, {2341, 543.4},
		{1452, 335.1}, {3621, 912.3}, {2874, 678.5}, {1738, 391.7}, {2432, 562.8},
		{4627, 1189.4}, {1384, 323.9}, {3748, 995.2}, {1923, 472.1}, {3981, 1052.8},
		{2678, 629.3}, {1729, 402.6}, {3162, 762.9}, {2491, 580.7}, {1829, 416.5},
		{4723, 1233.4}, {3521, 885.2}, {2012, 488.1}, {3827, 1005.9}, {2293, 530.2},
		{1402, 328.1}, {3648, 925.4}, {2745, 641.3}, {1782, 399.2}, {2487, 571.4},
		{4592, 1174.8}, {1345, 309.7}, {2842, 669.1}, {1592, 371.6}, {2475, 568.2},
		{4812, 1293.5}, {1698, 390.5}, {3284, 805.3}, {2714, 627.4}, {1856, 426.7},
		{4785, 1260.8}, {3462, 864.7}, {1989, 478.3}, {3723, 986.4}, {2304, 539.7},
		{1482, 338.9}, {3581, 899.3}, {2897, 685.2}, {1764, 410.2}, {2532, 589.8},
		{4692, 1219.3}, {1338, 306.4}, {2764, 650.5}, {1678, 389.2}, {2412, 560.7},
	}

	normalizedTS, minX, maxX, minY, maxY := normalize(trainingSet)

	g := GradientDescent{}
	w, b := g.Train(normalizedTS)
	fmt.Printf("Optimal Weight: %f, Optimal Bias: %f\n", w, b)

	scaledPlotSize := minMaxNormalize(3500.0, minX, maxX)
	scaledPredictedPrice := g.PredictPrice(scaledPlotSize, w, b)
	predictedPrice := minMaxDeNormalize(scaledPredictedPrice, minY, maxY)
	fmt.Printf("Predicted price for 3500 sq ft: $%.2f (in 100k USD)\n", predictedPrice)
}
